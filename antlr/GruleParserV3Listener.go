//  Copyright hyperjumptech/grule-rule-engine Authors
//
//  Licensed under the Apache License, Version 2.0 (the "License");
//  you may not use this file except in compliance with the License.
//  You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
//  Unless required by applicable law or agreed to in writing, software
//  distributed under the License is distributed on an "AS IS" BASIS,
//  WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//  See the License for the specific language governing permissions and
//  limitations under the License.

package antlr

import (
	"errors"
	"fmt"
	"github.com/antlr/antlr4/runtime/Go/antlr"
	"github.com/hyperjumptech/grule-rule-engine/antlr/parser/grulev3"
	"github.com/hyperjumptech/grule-rule-engine/ast"
	"github.com/hyperjumptech/grule-rule-engine/logger"
	"github.com/hyperjumptech/grule-rule-engine/pkg"
	"github.com/sirupsen/logrus"
	"reflect"
	"strconv"
	"strings"
)

var (
	// LoggerV3 is a logrus instance twith default fields for grule
	LoggerV3 = logger.Log.WithFields(logrus.Fields{
		"lib":    "grule",
		"struct": "GruleParserV3Listener",
	})

	// ErrAstStackInvalidParentNode is an error thrown if the parent node is not as expected
	ErrAstStackInvalidParentNode = errors.New("invalid ast parent node error")

	// ErrAstStackInvalidNode is an error thrown if the node is not as expected
	ErrAstStackInvalidNode = errors.New("invalid ast node error")
)

// NewGruleV3ParserListener create new instance of GruleV3ParserListener
func NewGruleV3ParserListener(KnowledgeBase *ast.KnowledgeBase, errorCallBack *pkg.GruleErrorReporter) *GruleV3ParserListener {
	return &GruleV3ParserListener{
		PreviousNode:  make([]string, 0),
		ErrorCallback: errorCallBack,
		KnowledgeBase: KnowledgeBase,
		Stack:         newStack(),
	}
}

// GruleV3ParserListener is an implementation of logic to build the execution flow or execution graph as it
// defined within the knowledge base.
type GruleV3ParserListener struct {
	grulev3.Basegrulev3Listener
	PreviousNode []string

	Grl           *ast.Grl
	Stack         *stack
	StopParse     bool
	ErrorCallback *pkg.GruleErrorReporter
	KnowledgeBase *ast.KnowledgeBase
}

// VisitTerminal is called when a terminal node is visited.
func (s *GruleV3ParserListener) VisitTerminal(node antlr.TerminalNode) {
	if s.StopParse {
		return
	}
	s.PreviousNode = append(s.PreviousNode, node.GetText())
	if len(s.PreviousNode) > 5 {
		s.PreviousNode = s.PreviousNode[1:]
	}
}

// VisitErrorNode is called when an error node is visited.
func (s *GruleV3ParserListener) VisitErrorNode(node antlr.ErrorNode) {
	LoggerV3.Errorf("GRL error, after '%v' and then unexpected '%s'", s.PreviousNode, node.GetText())
	s.StopParse = true
}

// EnterEveryRule is called when any rule is entered.
func (s *GruleV3ParserListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *GruleV3ParserListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterGrl is called when production grl is entered.
func (s *GruleV3ParserListener) EnterGrl(ctx *grulev3.GrlContext) {
	s.Grl = ast.NewGrl()
	s.Stack.Push(s.Grl)
}

// ExitGrl is called when production root is exited. The listener will instruct working memory re-index here.
func (s *GruleV3ParserListener) ExitGrl(ctx *grulev3.GrlContext) {
	if s.StopParse {
		return
	}
	_ = s.Stack.Pop().(*ast.Grl)
	for _, re := range s.Grl.RuleEntries {
		err := s.KnowledgeBase.AddRuleEntry(re)
		if err != nil {
			s.ErrorCallback.AddError(err)
		}
	}
}

// EnterRuleEntry is called when production ruleEntry is entered.
func (s *GruleV3ParserListener) EnterRuleEntry(ctx *grulev3.RuleEntryContext) {
	if s.StopParse {
		return
	}
	entry := ast.NewRuleEntry()
	entry.GrlText = ctx.GetText()
	s.Stack.Push(entry)
}

// ExitRuleEntry is called when production ruleEntry is exited.
func (s *GruleV3ParserListener) ExitRuleEntry(ctx *grulev3.RuleEntryContext) {
	if s.StopParse {
		return
	}
	entry := s.Stack.Pop().(*ast.RuleEntry)

	itv := s.Stack.Peek()
	entryReceiver, ok := itv.(ast.RuleEntryReceiver)
	if !ok {
		s.StopParse = true
		s.ErrorCallback.AddError(fmt.Errorf("expected node of type ast.RuleEntryReceiver (Grl) but %s : %w", reflect.TypeOf(itv).String(), ErrAstStackInvalidParentNode))
		return
	}
	if ctx.RuleName() != nil {
		entry.RuleName = ctx.RuleName().GetText()
	}
	if ctx.RuleDescription() != nil {
		txt := ctx.RuleDescription().GetText()
		entry.RuleDescription = txt[1 : len(txt)-1]
	}
	err := entryReceiver.ReceiveRuleEntry(entry)
	if err != nil {
		s.ErrorCallback.AddError(err)
	} else {
		LoggerV3.Debugf("Added RuleEntry : %s", entry.RuleName)
	}
}

// EnterSalience is called when production salience is entered.
func (s *GruleV3ParserListener) EnterSalience(ctx *grulev3.SalienceContext) {
	sal := ast.NewSalience(0)
	s.Stack.Push(sal)
}

// ExitSalience is called when production salience is exited.
func (s *GruleV3ParserListener) ExitSalience(ctx *grulev3.SalienceContext) {
	if s.StopParse {
		return
	}

	salience := s.Stack.Pop().(*ast.Salience)

	itv := s.Stack.Peek()
	salienceReceiver, ok := itv.(ast.SalienceReceiver)
	if !ok {
		s.StopParse = true
		s.ErrorCallback.AddError(fmt.Errorf("expected node of type ast.SalienceReceiver (RuleEntry) but %s : %w", reflect.TypeOf(itv).String(), ErrAstStackInvalidParentNode))
		return
	}
	err := salienceReceiver.AcceptSalience(salience)
	if err != nil {
		s.StopParse = true
		s.ErrorCallback.AddError(err)
	}
}

// EnterWhenScope is called when production whenScope is entered.
func (s *GruleV3ParserListener) EnterWhenScope(ctx *grulev3.WhenScopeContext) {
	if s.StopParse {
		return
	}
	whenScope := ast.NewWhenScope()
	whenScope.GrlText = ctx.GetText()
	s.Stack.Push(whenScope)
}

// ExitWhenScope is called when production whenScope is exited.
func (s *GruleV3ParserListener) ExitWhenScope(ctx *grulev3.WhenScopeContext) {
	if s.StopParse {
		return
	}
	when := s.Stack.Pop().(*ast.WhenScope)

	itv := s.Stack.Peek()
	receiver, ok := itv.(ast.WhenScopeReceiver)
	if !ok {
		s.StopParse = true
		s.ErrorCallback.AddError(fmt.Errorf("expected node of type ast.WhenScopeReceiver (RuleEntry) but %s : %w", reflect.TypeOf(itv).String(), ErrAstStackInvalidParentNode))
		return
	}
	err := receiver.AcceptWhenScope(when)
	if err != nil {
		s.StopParse = true
		s.ErrorCallback.AddError(err)
	}
}

// EnterThenScope is called when production thenScope is entered.
func (s *GruleV3ParserListener) EnterThenScope(ctx *grulev3.ThenScopeContext) {
	if s.StopParse {
		return
	}
	then := ast.NewThenScope()
	then.GrlText = ctx.GetText()
	s.Stack.Push(then)
}

// ExitThenScope is called when production thenScope is exited.
func (s *GruleV3ParserListener) ExitThenScope(ctx *grulev3.ThenScopeContext) {
	if s.StopParse {
		return
	}
	then := s.Stack.Pop().(*ast.ThenScope)

	itv := s.Stack.Peek()
	receiver, ok := itv.(ast.ThenScopeReceiver)
	if !ok {
		s.StopParse = true
		s.ErrorCallback.AddError(fmt.Errorf("expected node of type ast.ThenScopeReceiver (RuleEntry) but %s : %w", reflect.TypeOf(itv).String(), ErrAstStackInvalidParentNode))
		return
	}
	err := receiver.AcceptThenScope(then)
	if err != nil {
		s.StopParse = true
		s.ErrorCallback.AddError(err)
	}
}

// EnterThenExpressionList is called when production thenExpressionList is entered.
func (s *GruleV3ParserListener) EnterThenExpressionList(ctx *grulev3.ThenExpressionListContext) {
	if s.StopParse {
		return
	}
	thenExpList := ast.NewThenExpressionList()
	thenExpList.GrlText = ctx.GetText()
	s.Stack.Push(thenExpList)
}

// ExitThenExpressionList is called when production thenExpressionList is exited.
func (s *GruleV3ParserListener) ExitThenExpressionList(ctx *grulev3.ThenExpressionListContext) {
	if s.StopParse {
		return
	}
	thenExpList := s.Stack.Pop().(*ast.ThenExpressionList)

	itv := s.Stack.Peek()
	receiver, ok := itv.(ast.ThenExpressionListReceiver)
	if !ok {
		s.StopParse = true
		s.ErrorCallback.AddError(fmt.Errorf("expected node of type ast.ThenExpressionListReceiver (RuleEntry) but %s : %w", reflect.TypeOf(itv).String(), ErrAstStackInvalidParentNode))
		return
	}
	err := receiver.AcceptThenExpressionList(thenExpList)
	if err != nil {
		s.StopParse = true
		s.ErrorCallback.AddError(err)
	}
}

// EnterThenExpression is called when production thenExpression is entered.
func (s *GruleV3ParserListener) EnterThenExpression(ctx *grulev3.ThenExpressionContext) {
	if s.StopParse {
		return
	}
	thenExpr := ast.NewThenExpression()
	thenExpr.GrlText = ctx.GetText()
	s.Stack.Push(thenExpr)
}

// ExitThenExpression is called when production thenExpression is exited.
func (s *GruleV3ParserListener) ExitThenExpression(ctx *grulev3.ThenExpressionContext) {
	if s.StopParse {
		return
	}
	thenExpr := s.Stack.Pop().(*ast.ThenExpression)

	itv := s.Stack.Peek()
	receiver, ok := itv.(ast.ThenExpressionReceiver)
	if !ok {
		s.StopParse = true
		s.ErrorCallback.AddError(fmt.Errorf("expected node of type ast.ThenExpressionReceiver (ThenExpressionList) but %s : %w", reflect.TypeOf(itv).String(), ErrAstStackInvalidParentNode))
		return
	}
	err := receiver.AcceptThenExpression(thenExpr)
	if err != nil {
		s.StopParse = true
		s.ErrorCallback.AddError(err)
	}
}

// EnterAssignment is called when production assignment is entered.
func (s *GruleV3ParserListener) EnterAssignment(ctx *grulev3.AssignmentContext) {
	if s.StopParse {
		return
	}
	assign := ast.NewAssignment()
	assign.GrlText = ctx.GetText()
	s.Stack.Push(assign)
}

// ExitAssignment is called when production assignment is exited.
func (s *GruleV3ParserListener) ExitAssignment(ctx *grulev3.AssignmentContext) {
	if s.StopParse {
		return
	}
	assign := s.Stack.Pop().(*ast.Assignment)

	itv := s.Stack.Peek()
	receiver, ok := itv.(ast.AssignmentReceiver)
	if !ok {
		s.StopParse = true
		s.ErrorCallback.AddError(fmt.Errorf("expected node of type ast.AssignmentReceiver (ThenExpression) but %s : %w", reflect.TypeOf(itv).String(), ErrAstStackInvalidParentNode))
		return
	}
	assign.IsAssign = ctx.ASSIGN() != nil
	assign.IsPlusAssign = ctx.PLUS_ASIGN() != nil
	assign.IsMinusAssign = ctx.MINUS_ASIGN() != nil
	assign.IsDivAssign = ctx.DIV_ASIGN() != nil
	assign.IsMulAssign = ctx.MUL_ASIGN() != nil

	err := receiver.AcceptAssignment(assign)
	if err != nil {
		s.StopParse = true
		s.ErrorCallback.AddError(err)
	}
}

// EnterExpression is called when production expression is entered.
func (s *GruleV3ParserListener) EnterExpression(ctx *grulev3.ExpressionContext) {
	if s.StopParse {
		return
	}
	expr := ast.NewExpression()
	expr.GrlText = ctx.GetText()
	s.Stack.Push(expr)
}

// ExitExpression is called when production expression is exited.
func (s *GruleV3ParserListener) ExitExpression(ctx *grulev3.ExpressionContext) {
	if s.StopParse {
		return
	}
	expr := s.Stack.Pop().(*ast.Expression)

	itv := s.Stack.Peek()
	exprRec, ok := itv.(ast.ExpressionReceiver)
	if !ok {
		s.StopParse = true
		s.ErrorCallback.AddError(fmt.Errorf("expected node of type ast.ExpressionReceiver (ArgumentList, ArrayMapSelector, Assignment, Expression, WhenScope) but %s : %w", reflect.TypeOf(itv).String(), ErrAstStackInvalidParentNode))
		return
	}

	if ctx.LR_BRACKET() != nil && ctx.RR_BRACKET() != nil && ctx.NEGATION() != nil {
		expr.Negated = ctx.NEGATION() != nil
	}

	err := exprRec.AcceptExpression(s.KnowledgeBase.WorkingMemory.AddExpression(expr))
	if err != nil {
		s.StopParse = true
		s.ErrorCallback.AddError(err)
	}
}

// EnterMulDivOperators is called when production mulDivOperators is entered.
func (s *GruleV3ParserListener) EnterMulDivOperators(ctx *grulev3.MulDivOperatorsContext) {
	if s.StopParse {
		return
	}
	expr := s.Stack.Peek().(*ast.Expression)
	switch ctx.GetText() {
	case "*":
		expr.Operator = ast.OpMul
	case "/":
		expr.Operator = ast.OpDiv
	case "%":
		expr.Operator = ast.OpMod
	}
}

// ExitMulDivOperators is called when production mulDivOperators is exited.
func (s *GruleV3ParserListener) ExitMulDivOperators(ctx *grulev3.MulDivOperatorsContext) {}

// EnterAddMinusOperators is called when production addMinusOperators is entered.
func (s *GruleV3ParserListener) EnterAddMinusOperators(ctx *grulev3.AddMinusOperatorsContext) {
	if s.StopParse {
		return
	}
	expr := s.Stack.Peek().(*ast.Expression)
	switch ctx.GetText() {
	case "+":
		expr.Operator = ast.OpAdd
	case "-":
		expr.Operator = ast.OpSub
	case "|":
		expr.Operator = ast.OpBitOr
	case "&":
		expr.Operator = ast.OpBitAnd
	}
}

// ExitAddMinusOperators is called when production addMinusOperators is exited.
func (s *GruleV3ParserListener) ExitAddMinusOperators(ctx *grulev3.AddMinusOperatorsContext) {}

// EnterComparisonOperator is called when production comparisonOperator is entered.
func (s *GruleV3ParserListener) EnterComparisonOperator(ctx *grulev3.ComparisonOperatorContext) {
	if s.StopParse {
		return
	}
	expr := s.Stack.Peek().(*ast.Expression)
	switch ctx.GetText() {
	case "<":
		expr.Operator = ast.OpLT
	case "<=":
		expr.Operator = ast.OpLTE
	case ">":
		expr.Operator = ast.OpGT
	case ">=":
		expr.Operator = ast.OpGTE
	case "==":
		expr.Operator = ast.OpEq
	case "!=":
		expr.Operator = ast.OpNEq
	}
}

// ExitComparisonOperator is called when production comparisonOperator is exited.
func (s *GruleV3ParserListener) ExitComparisonOperator(ctx *grulev3.ComparisonOperatorContext) {}

// EnterAndLogicOperator is called when production andLogicOperator is entered.
func (s *GruleV3ParserListener) EnterAndLogicOperator(ctx *grulev3.AndLogicOperatorContext) {
	if s.StopParse {
		return
	}
	expr := s.Stack.Peek().(*ast.Expression)
	expr.Operator = ast.OpAnd
}

// ExitAndLogicOperator is called when production andLogicOperator is exited.
func (s *GruleV3ParserListener) ExitAndLogicOperator(ctx *grulev3.AndLogicOperatorContext) {}

// EnterOrLogicOperator is called when production orLogicOperator is entered.
func (s *GruleV3ParserListener) EnterOrLogicOperator(ctx *grulev3.OrLogicOperatorContext) {
	if s.StopParse {
		return
	}
	expr := s.Stack.Peek().(*ast.Expression)
	expr.Operator = ast.OpOr
}

// ExitOrLogicOperator is called when production orLogicOperator is exited.
func (s *GruleV3ParserListener) ExitOrLogicOperator(ctx *grulev3.OrLogicOperatorContext) {}

// EnterExpressionAtom is called when production expressionAtom is entered.
func (s *GruleV3ParserListener) EnterExpressionAtom(ctx *grulev3.ExpressionAtomContext) {
	if s.StopParse {
		return
	}
	atm := ast.NewExpressionAtom()
	atm.GrlText = ctx.GetText()
	s.Stack.Push(atm)
}

// ExitExpressionAtom is called when production expressionAtom is exited.
func (s *GruleV3ParserListener) ExitExpressionAtom(ctx *grulev3.ExpressionAtomContext) {
	if s.StopParse {
		return
	}
	atm := s.Stack.Pop().(*ast.ExpressionAtom)

	itv := s.Stack.Peek()
	expr, ok := itv.(ast.ExpressionAtomReceiver)
	if !ok {
		s.StopParse = true
		s.ErrorCallback.AddError(fmt.Errorf("expected node of type ast.ExpressionAtomReceiver (Expression, ExpressionAtom, ThenExpression) but %s : %w", reflect.TypeOf(itv).String(), ErrAstStackInvalidParentNode))
		return
	}

	atm.Negated = ctx.NEGATION() != nil

	err := expr.AcceptExpressionAtom(s.KnowledgeBase.WorkingMemory.AddExpressionAtom(atm))
	if err != nil {
		s.StopParse = true
		s.ErrorCallback.AddError(err)
	}
}

// EnterArrayMapSelector is called when production arrayMapSelector is entered.
func (s *GruleV3ParserListener) EnterArrayMapSelector(ctx *grulev3.ArrayMapSelectorContext) {
	if s.StopParse {
		return
	}
	sel := ast.NewArrayMapSelector()
	sel.GrlText = ctx.GetText()
	s.Stack.Push(sel)
}

// ExitArrayMapSelector is called when production arrayMapSelector is exited.
func (s *GruleV3ParserListener) ExitArrayMapSelector(ctx *grulev3.ArrayMapSelectorContext) {
	if s.StopParse {
		return
	}
	sel := s.Stack.Pop().(*ast.ArrayMapSelector)

	itv := s.Stack.Peek()
	receiver, ok := itv.(ast.ArrayMapSelectorReceiver)
	if !ok {
		s.StopParse = true
		s.ErrorCallback.AddError(fmt.Errorf("expected node of type ast.ArrayMapSelectorReceiver (ExpressionAtom, Variable) but %s : %w", reflect.TypeOf(itv).String(), ErrAstStackInvalidParentNode))
		return
	}
	err := receiver.AcceptArrayMapSelector(sel)
	if err != nil {
		s.StopParse = true
		s.ErrorCallback.AddError(err)
	}
}

// EnterFunctionCall is called when production functionCall is entered.
func (s *GruleV3ParserListener) EnterFunctionCall(ctx *grulev3.FunctionCallContext) {
	if s.StopParse {
		return
	}
	fun := ast.NewFunctionCall()
	fun.FunctionName = ctx.SIMPLENAME().GetText()
	s.Stack.Push(fun)
}

// ExitFunctionCall is called when production functionCall is exited.
func (s *GruleV3ParserListener) ExitFunctionCall(ctx *grulev3.FunctionCallContext) {
	if s.StopParse {
		return
	}
	fun := s.Stack.Pop().(*ast.FunctionCall)

	itv := s.Stack.Peek()
	metRec, ok := itv.(ast.FunctionCallReceiver)
	if !ok {
		s.StopParse = true
		s.ErrorCallback.AddError(fmt.Errorf("expected node of type ast.FunctionCallReceiver (ExpressionAtom) but %s : %w", reflect.TypeOf(itv).String(), ErrAstStackInvalidParentNode))
		return
	}
	err := metRec.AcceptFunctionCall(fun)
	if err != nil {
		s.StopParse = true
		s.ErrorCallback.AddError(err)
	}
}

// EnterArgumentList is called when production argumentList is entered.
func (s *GruleV3ParserListener) EnterArgumentList(ctx *grulev3.ArgumentListContext) {
	if s.StopParse {
		return
	}
	argList := ast.NewArgumentList()
	argList.GrlText = ctx.GetText()
	s.Stack.Push(argList)
}

// ExitArgumentList is called when production argumentList is exited.
func (s *GruleV3ParserListener) ExitArgumentList(ctx *grulev3.ArgumentListContext) {
	if s.StopParse {
		return
	}
	argList := s.Stack.Pop().(*ast.ArgumentList)

	itv := s.Stack.Peek()
	argListRec, ok := itv.(ast.ArgumentListReceiver)
	if !ok {
		s.StopParse = true
		s.ErrorCallback.AddError(fmt.Errorf("expected node of type ast.ArgumentListReceiver (FunctionCall) but %s : %w", reflect.TypeOf(itv).String(), ErrAstStackInvalidParentNode))
		return
	}
	LoggerV3.Tracef("Adding Argument List To Receiver")
	err := argListRec.AcceptArgumentList(argList)
	if err != nil {
		s.StopParse = true
		s.ErrorCallback.AddError(err)
	}
}

// EnterVariable is called when production variable is entered.
func (s *GruleV3ParserListener) EnterVariable(ctx *grulev3.VariableContext) {
	if s.StopParse {
		return
	}
	vari := ast.NewVariable()
	if ctx.SIMPLENAME() != nil && len(ctx.SIMPLENAME().GetText()) > 0 {
		vari.Name = ctx.SIMPLENAME().GetText()
	}
	if ctx.MemberVariable() != nil && len(ctx.MemberVariable().GetText()) > 0 {
		vari.Name = ctx.MemberVariable().GetText()[1:]
	}
	vari.GrlText = ctx.GetText()
	s.Stack.Push(vari)
}

// ExitVariable is called when production variable is exited.
func (s *GruleV3ParserListener) ExitVariable(ctx *grulev3.VariableContext) {
	if s.StopParse {
		return
	}
	vari := s.Stack.Pop().(*ast.Variable)

	itv := s.Stack.Peek()
	variRec, ok := itv.(ast.VariableReceiver)
	if !ok {
		s.StopParse = true
		s.ErrorCallback.AddError(fmt.Errorf("expected node of type ast.VariableReceiver (Assignment, ExpressionAtom, Variable) but %s : %w", reflect.TypeOf(itv).String(), ErrAstStackInvalidParentNode))
		return
	}

	err := variRec.AcceptVariable(s.KnowledgeBase.WorkingMemory.AddVariable(vari))
	if err != nil {
		s.StopParse = true
		s.ErrorCallback.AddError(err)
	}
}

// EnterMemberVariable is called when production memberVariable is entered.
func (s *GruleV3ParserListener) EnterMemberVariable(ctx *grulev3.MemberVariableContext) {}

// ExitMemberVariable is called when production memberVariable is exited.
func (s *GruleV3ParserListener) ExitMemberVariable(ctx *grulev3.MemberVariableContext) {
	itv := s.Stack.Peek()
	vari, ok := itv.(ast.MemberVariableReceiver)
	if !ok {
		s.StopParse = true
		s.ErrorCallback.AddError(fmt.Errorf("expected node of type ast.MemberVariableReceiver (ExpressionAtom, Variable) but %s : %w", reflect.TypeOf(itv).String(), ErrAstStackInvalidParentNode))
		return
	}
	vari.AcceptMemberVariable(ctx.SIMPLENAME().GetText())
}

// EnterConstant is called when production constant is entered.
func (s *GruleV3ParserListener) EnterConstant(ctx *grulev3.ConstantContext) {
	if s.StopParse {
		return
	}
	cons := ast.NewConstant()
	cons.GrlText = ctx.GetText()
	s.Stack.Push(cons)
}

// ExitConstant is called when production constant is exited.
func (s *GruleV3ParserListener) ExitConstant(ctx *grulev3.ConstantContext) {
	if s.StopParse {
		return
	}
	cons := s.Stack.Pop().(*ast.Constant)

	itv := s.Stack.Peek()
	conRec, ok := itv.(ast.ConstantReceiver)
	if !ok {
		s.StopParse = true
		s.ErrorCallback.AddError(fmt.Errorf("expected node of type ast.ConstantReceiver (ExpressionAtom) but %s : %w", reflect.TypeOf(itv).String(), ErrAstStackInvalidParentNode))
		return
	}
	if ctx.NIL_LITERAL() != nil {
		cons.IsNil = true
	}
	err := conRec.AcceptConstant(cons)
	if err != nil {
		s.StopParse = true
		s.ErrorCallback.AddError(err)
	}
}

// EnterStringLiteral is called when production stringLiteral is entered.
func (s *GruleV3ParserListener) EnterStringLiteral(ctx *grulev3.StringLiteralContext) {
}

// ExitStringLiteral is called when production stringLiteral is exited.
func (s *GruleV3ParserListener) ExitStringLiteral(ctx *grulev3.StringLiteralContext) {
	if s.StopParse {
		return
	}
	dec, err := unquoteString(ctx.GetText())
	if err != nil {
		s.ErrorCallback.AddError(fmt.Errorf("error parsing quoted string (%s): %s", ctx.GetText(), err.Error()))
		return
	}
	itv := s.Stack.Peek()
	receiver, ok := itv.(ast.StringLiteralReceiver)
	if !ok {
		s.StopParse = true
		s.ErrorCallback.AddError(fmt.Errorf("expected node of type ast.StringLiteralReceiver (Constant) but %s : %w", reflect.TypeOf(itv).String(), ErrAstStackInvalidParentNode))
		return
	}
	receiver.AcceptStringLiteral(&ast.StringLiteral{String: dec})
}

// EnterBooleanLiteral is called when production booleanLiteral is entered.
func (s *GruleV3ParserListener) EnterBooleanLiteral(ctx *grulev3.BooleanLiteralContext) {}

// ExitBooleanLiteral is called when production booleanLiteral is exited.
func (s *GruleV3ParserListener) ExitBooleanLiteral(ctx *grulev3.BooleanLiteralContext) {
	if s.StopParse {
		return
	}

	itv := s.Stack.Peek()
	receiver, ok := itv.(ast.BooleanLiteralReceiver)
	if !ok {
		s.StopParse = true
		s.ErrorCallback.AddError(fmt.Errorf("expected node of type ast.BooleanLiteralReceiver (Constant) but %s : %w", reflect.TypeOf(itv).String(), ErrAstStackInvalidParentNode))
		return
	}
	lit := &ast.BooleanLiteral{}
	switch strings.ToLower(ctx.GetText()) {
	case "true":
		lit.Boolean = true
	case "false":
		lit.Boolean = false
	}
	receiver.AcceptBooleanLiteral(lit)
}

// EnterIntegerLiteral is called when production integerLiteral is entered.
func (s *GruleV3ParserListener) EnterIntegerLiteral(ctx *grulev3.IntegerLiteralContext) {}

// ExitIntegerLiteral is called when production integerLiteral is exited.
func (s *GruleV3ParserListener) ExitIntegerLiteral(ctx *grulev3.IntegerLiteralContext) {
	lit := &ast.IntegerLiteral{}
	i, err := strconv.ParseInt(ctx.GetText(), 0, 64)
	if err != nil {
		s.StopParse = true
		s.ErrorCallback.AddError(err)
	} else {
		lit.Integer = i
	}
	itv := s.Stack.Peek()
	receiver, ok := itv.(ast.IntegerLiteralReceiver)
	if !ok {
		s.StopParse = true
		s.ErrorCallback.AddError(fmt.Errorf("expected node of type ast.IntegerLiteralReceiver (Constant) but %s : %w", reflect.TypeOf(itv).String(), ErrAstStackInvalidParentNode))
		return
	}
	receiver.AcceptIntegerLiteral(lit)
}

// EnterFloatLiteral is called when production floatLiteral is entered.
func (s *GruleV3ParserListener) EnterFloatLiteral(ctx *grulev3.FloatLiteralContext) {}

// ExitFloatLiteral is called when production floatLiteral is exited.
func (s *GruleV3ParserListener) ExitFloatLiteral(ctx *grulev3.FloatLiteralContext) {
	lit := &ast.FloatLiteral{}
	i, err := strconv.ParseFloat(ctx.GetText(), 64)
	if err != nil {
		s.StopParse = true
		s.ErrorCallback.AddError(err)
	} else {
		lit.Float = i
	}

	itv := s.Stack.Peek()
	receiver, ok := itv.(ast.FloatLiteralReceiver)
	if !ok {
		s.StopParse = true
		s.ErrorCallback.AddError(fmt.Errorf("expected node of type ast.FloatLiteralReceiver (Constant) but %s : %w", reflect.TypeOf(itv).String(), ErrAstStackInvalidParentNode))
		return
	}
	receiver.AcceptFloatLiteral(lit)
}
