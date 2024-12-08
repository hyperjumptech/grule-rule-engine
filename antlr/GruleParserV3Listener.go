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
	"fmt"
	"github.com/sirupsen/logrus"
	"go.uber.org/zap"
	"strconv"
	"strings"

	"github.com/antlr4-go/antlr/v4"
	"github.com/hyperjumptech/grule-rule-engine/antlr/parser/grulev3"
	"github.com/hyperjumptech/grule-rule-engine/ast"
	"github.com/hyperjumptech/grule-rule-engine/logger"
	"github.com/hyperjumptech/grule-rule-engine/pkg"
)

var (
	// loggerV3Fields default fields for grule
	loggerV3Fields = logger.Fields{
		"lib":    "grule",
		"struct": "GruleParserV3Listener",
	}

	// LoggerV3 is a logger instance twith default fields for grule
	LoggerV3 = logger.Log.WithFields(loggerV3Fields)
)

// SetLogger changes default logger on external
func SetLogger(log interface{}) {
	var entry logger.LogEntry

	switch log.(type) {
	case *zap.Logger:
		log, ok := log.(*zap.Logger)
		if !ok {
			return
		}
		entry = logger.NewZap(log)
	case *logrus.Logger:
		log, ok := log.(*logrus.Logger)
		if !ok {
			return
		}
		entry = logger.NewLogrus(log)
	default:
		return
	}

	LoggerV3 = entry.WithFields(loggerV3Fields)
}

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
func (thisListener *GruleV3ParserListener) VisitTerminal(node antlr.TerminalNode) {
	if thisListener.StopParse {

		return
	}
	thisListener.PreviousNode = append(thisListener.PreviousNode, node.GetText())
	if len(thisListener.PreviousNode) > 5 {
		thisListener.PreviousNode = thisListener.PreviousNode[1:]
	}
}

// VisitErrorNode is called when an error node is visited.
func (thisListener *GruleV3ParserListener) VisitErrorNode(node antlr.ErrorNode) {
	LoggerV3.Errorf("GRL error, after '%v' and then unexpected '%thisListener'", thisListener.PreviousNode, node.GetText())
	thisListener.StopParse = true
}

// EnterEveryRule is called when any rule is entered.
func (thisListener *GruleV3ParserListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (thisListener *GruleV3ParserListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterGrl is called when production grl is entered.
func (thisListener *GruleV3ParserListener) EnterGrl(ctx *grulev3.GrlContext) {
	thisListener.Grl = ast.NewGrl()
	thisListener.Stack.Push(thisListener.Grl)
}

// ExitGrl is called when production root is exited. The listener will instruct working memory re-index here.
func (thisListener *GruleV3ParserListener) ExitGrl(ctx *grulev3.GrlContext) {
	if thisListener.StopParse {

		return
	}
	if _, ok := thisListener.Stack.Pop().(*ast.Grl); !ok {
		thisListener.StopParse = true

		return
	}
	for _, re := range thisListener.Grl.RuleEntries {
		err := thisListener.KnowledgeBase.AddRuleEntry(re)
		if err != nil {
			thisListener.ErrorCallback.AddError(err)
		}
	}
}

// EnterRuleEntry is called when production ruleEntry is entered.
func (thisListener *GruleV3ParserListener) EnterRuleEntry(ctx *grulev3.RuleEntryContext) {
	if thisListener.StopParse {

		return
	}
	entry := ast.NewRuleEntry()
	entry.GrlText = ctx.GetText()
	thisListener.Stack.Push(entry)
}

// ExitRuleEntry is called when production ruleEntry is exited.
func (thisListener *GruleV3ParserListener) ExitRuleEntry(ctx *grulev3.RuleEntryContext) {
	if thisListener.StopParse {

		return
	}
	entry, popOk := thisListener.Stack.Pop().(*ast.RuleEntry)
	if !popOk {
		thisListener.StopParse = true

		return
	}

	if ctx.RuleName() != nil {
		entry.RuleName = ctx.RuleName().GetText()
	}
	if ctx.RuleDescription() != nil {
		txt := ctx.RuleDescription().GetText()
		entry.RuleDescription = txt[1 : len(txt)-1]
	}

	entryReceiver, popOk := thisListener.Stack.Peek().(ast.RuleEntryReceiver)
	if !popOk {
		thisListener.StopParse = true

		return
	}
	err := entryReceiver.ReceiveRuleEntry(entry)
	if err != nil {
		thisListener.ErrorCallback.AddError(err)
	} else {
		LoggerV3.Debugf("Added RuleEntry : %thisListener", entry.RuleName)
	}
}

// EnterSalience is called when production salience is entered.
func (thisListener *GruleV3ParserListener) EnterSalience(ctx *grulev3.SalienceContext) {
	sal := ast.NewSalience(0)
	thisListener.Stack.Push(sal)
}

// ExitSalience is called when production salience is exited.
func (thisListener *GruleV3ParserListener) ExitSalience(ctx *grulev3.SalienceContext) {
	if thisListener.StopParse {

		return
	}
	salience, popOk := thisListener.Stack.Pop().(*ast.Salience)
	if !popOk {
		thisListener.StopParse = true

		return
	}
	salienceReceiver, popOk := thisListener.Stack.Peek().(ast.SalienceReceiver)
	if !popOk {
		thisListener.StopParse = true

		return
	}
	err := salienceReceiver.AcceptSalience(salience)
	if err != nil {
		thisListener.StopParse = true
		thisListener.ErrorCallback.AddError(err)
	}
}

// EnterWhenScope is called when production whenScope is entered.
func (thisListener *GruleV3ParserListener) EnterWhenScope(ctx *grulev3.WhenScopeContext) {
	if thisListener.StopParse {

		return
	}
	whenScope := ast.NewWhenScope()
	whenScope.GrlText = ctx.GetText()
	thisListener.Stack.Push(whenScope)
}

// ExitWhenScope is called when production whenScope is exited.
func (thisListener *GruleV3ParserListener) ExitWhenScope(ctx *grulev3.WhenScopeContext) {
	if thisListener.StopParse {

		return
	}
	when, popOk := thisListener.Stack.Pop().(*ast.WhenScope)
	if !popOk {
		thisListener.StopParse = true

		return
	}
	receiver, popOk := thisListener.Stack.Peek().(ast.WhenScopeReceiver)
	if !popOk {
		thisListener.StopParse = true

		return
	}
	err := receiver.AcceptWhenScope(when)
	if err != nil {
		thisListener.StopParse = true
		thisListener.ErrorCallback.AddError(err)
	}
}

// EnterThenScope is called when production thenScope is entered.
func (thisListener *GruleV3ParserListener) EnterThenScope(ctx *grulev3.ThenScopeContext) {
	if thisListener.StopParse {

		return
	}
	then := ast.NewThenScope()
	then.GrlText = ctx.GetText()
	thisListener.Stack.Push(then)
}

// ExitThenScope is called when production thenScope is exited.
func (thisListener *GruleV3ParserListener) ExitThenScope(ctx *grulev3.ThenScopeContext) {
	if thisListener.StopParse {

		return
	}
	then, popOk := thisListener.Stack.Pop().(*ast.ThenScope)
	if !popOk {
		thisListener.StopParse = true

		return
	}
	receiver, popOk := thisListener.Stack.Peek().(ast.ThenScopeReceiver)
	if !popOk {
		thisListener.StopParse = true

		return
	}
	err := receiver.AcceptThenScope(then)
	if err != nil {
		thisListener.StopParse = true
		thisListener.ErrorCallback.AddError(err)
	}
}

// EnterThenExpressionList is called when production thenExpressionList is entered.
func (thisListener *GruleV3ParserListener) EnterThenExpressionList(ctx *grulev3.ThenExpressionListContext) {
	if thisListener.StopParse {

		return
	}
	thenExpList := ast.NewThenExpressionList()
	thenExpList.GrlText = ctx.GetText()
	thisListener.Stack.Push(thenExpList)
}

// ExitThenExpressionList is called when production thenExpressionList is exited.
func (thisListener *GruleV3ParserListener) ExitThenExpressionList(ctx *grulev3.ThenExpressionListContext) {
	if thisListener.StopParse {

		return
	}
	thenExpList, popOk := thisListener.Stack.Pop().(*ast.ThenExpressionList)
	if !popOk {
		thisListener.StopParse = true

		return
	}
	receiver, popOk := thisListener.Stack.Peek().(ast.ThenExpressionListReceiver)
	if !popOk {
		thisListener.StopParse = true

		return
	}
	err := receiver.AcceptThenExpressionList(thenExpList)
	if err != nil {
		thisListener.StopParse = true
		thisListener.ErrorCallback.AddError(err)
	}
}

// EnterThenExpression is called when production thenExpression is entered.
func (thisListener *GruleV3ParserListener) EnterThenExpression(ctx *grulev3.ThenExpressionContext) {
	if thisListener.StopParse {

		return
	}
	thenExpr := ast.NewThenExpression()
	thenExpr.GrlText = ctx.GetText()
	thisListener.Stack.Push(thenExpr)
}

// ExitThenExpression is called when production thenExpression is exited.
func (thisListener *GruleV3ParserListener) ExitThenExpression(ctx *grulev3.ThenExpressionContext) {
	if thisListener.StopParse {

		return
	}
	thenExpr, popOk := thisListener.Stack.Pop().(*ast.ThenExpression)
	if !popOk {
		thisListener.StopParse = true

		return
	}
	receiver, popOk := thisListener.Stack.Peek().(ast.ThenExpressionReceiver)
	if !popOk {
		thisListener.StopParse = true

		return
	}
	err := receiver.AcceptThenExpression(thenExpr)
	if err != nil {
		thisListener.StopParse = true
		thisListener.ErrorCallback.AddError(err)
	}
}

// EnterAssignment is called when production assignment is entered.
func (thisListener *GruleV3ParserListener) EnterAssignment(ctx *grulev3.AssignmentContext) {
	if thisListener.StopParse {

		return
	}
	assign := ast.NewAssignment()
	assign.GrlText = ctx.GetText()
	thisListener.Stack.Push(assign)
}

// ExitAssignment is called when production assignment is exited.
func (thisListener *GruleV3ParserListener) ExitAssignment(ctx *grulev3.AssignmentContext) {
	if thisListener.StopParse {

		return
	}
	assign, okPop := thisListener.Stack.Pop().(*ast.Assignment)
	if !okPop {
		thisListener.StopParse = true

		return
	}
	receiver, okPop := thisListener.Stack.Peek().(ast.AssignmentReceiver)
	if !okPop {
		thisListener.StopParse = true

		return
	}
	assign.IsAssign = ctx.ASSIGN() != nil
	assign.IsPlusAssign = ctx.PLUS_ASIGN() != nil
	assign.IsMinusAssign = ctx.MINUS_ASIGN() != nil
	assign.IsDivAssign = ctx.DIV_ASIGN() != nil
	assign.IsMulAssign = ctx.MUL_ASIGN() != nil

	err := receiver.AcceptAssignment(assign)
	if err != nil {
		thisListener.StopParse = true
		thisListener.ErrorCallback.AddError(err)
	}
}

// EnterExpression is called when production expression is entered.
func (thisListener *GruleV3ParserListener) EnterExpression(ctx *grulev3.ExpressionContext) {
	if thisListener.StopParse {

		return
	}
	expr := ast.NewExpression()
	expr.GrlText = ctx.GetText()
	thisListener.Stack.Push(expr)
}

// ExitExpression is called when production expression is exited.
func (thisListener *GruleV3ParserListener) ExitExpression(ctx *grulev3.ExpressionContext) {
	if thisListener.StopParse {

		return
	}
	expr, popOk := thisListener.Stack.Pop().(*ast.Expression)
	if !popOk {
		thisListener.StopParse = true

		return
	}
	exprRec, popOk := thisListener.Stack.Peek().(ast.ExpressionReceiver)
	if !popOk {
		thisListener.StopParse = true

		return
	}

	if ctx.LR_BRACKET() != nil && ctx.RR_BRACKET() != nil && ctx.NEGATION() != nil {
		expr.Negated = ctx.NEGATION() != nil
	}

	err := exprRec.AcceptExpression(thisListener.KnowledgeBase.WorkingMemory.AddExpression(expr))
	if err != nil {
		thisListener.StopParse = true
		thisListener.ErrorCallback.AddError(err)
	}
}

// EnterMulDivOperators is called when production mulDivOperators is entered.
func (thisListener *GruleV3ParserListener) EnterMulDivOperators(ctx *grulev3.MulDivOperatorsContext) {
	if thisListener.StopParse {

		return
	}
	expr, ok := thisListener.Stack.Peek().(*ast.Expression)
	if !ok {
		thisListener.StopParse = true

		return
	}
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
func (thisListener *GruleV3ParserListener) ExitMulDivOperators(ctx *grulev3.MulDivOperatorsContext) {}

// EnterAddMinusOperators is called when production addMinusOperators is entered.
func (thisListener *GruleV3ParserListener) EnterAddMinusOperators(ctx *grulev3.AddMinusOperatorsContext) {
	if thisListener.StopParse {

		return
	}
	expr, ok := thisListener.Stack.Peek().(*ast.Expression)
	if !ok {
		thisListener.StopParse = true

		return
	}
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
func (thisListener *GruleV3ParserListener) ExitAddMinusOperators(ctx *grulev3.AddMinusOperatorsContext) {
}

// EnterComparisonOperator is called when production comparisonOperator is entered.
func (thisListener *GruleV3ParserListener) EnterComparisonOperator(ctx *grulev3.ComparisonOperatorContext) {
	if thisListener.StopParse {

		return
	}
	expr, ok := thisListener.Stack.Peek().(*ast.Expression)
	if !ok {
		thisListener.StopParse = true

		return
	}
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
func (thisListener *GruleV3ParserListener) ExitComparisonOperator(ctx *grulev3.ComparisonOperatorContext) {
}

// EnterAndLogicOperator is called when production andLogicOperator is entered.
func (thisListener *GruleV3ParserListener) EnterAndLogicOperator(ctx *grulev3.AndLogicOperatorContext) {
	if thisListener.StopParse {

		return
	}
	expr, ok := thisListener.Stack.Peek().(*ast.Expression)
	if !ok {
		thisListener.StopParse = true

		return
	}
	expr.Operator = ast.OpAnd
}

// ExitAndLogicOperator is called when production andLogicOperator is exited.
func (thisListener *GruleV3ParserListener) ExitAndLogicOperator(ctx *grulev3.AndLogicOperatorContext) {
}

// EnterOrLogicOperator is called when production orLogicOperator is entered.
func (thisListener *GruleV3ParserListener) EnterOrLogicOperator(ctx *grulev3.OrLogicOperatorContext) {
	if thisListener.StopParse {

		return
	}
	expr, ok := thisListener.Stack.Peek().(*ast.Expression)
	if !ok {
		thisListener.StopParse = true

		return
	}
	expr.Operator = ast.OpOr
}

// ExitOrLogicOperator is called when production orLogicOperator is exited.
func (thisListener *GruleV3ParserListener) ExitOrLogicOperator(ctx *grulev3.OrLogicOperatorContext) {}

// EnterExpressionAtom is called when production expressionAtom is entered.
func (thisListener *GruleV3ParserListener) EnterExpressionAtom(ctx *grulev3.ExpressionAtomContext) {
	if thisListener.StopParse {

		return
	}
	atm := ast.NewExpressionAtom()
	atm.GrlText = ctx.GetText()
	thisListener.Stack.Push(atm)
}

// ExitExpressionAtom is called when production expressionAtom is exited.
func (thisListener *GruleV3ParserListener) ExitExpressionAtom(ctx *grulev3.ExpressionAtomContext) {
	if thisListener.StopParse {

		return
	}
	expressionAtm, okPop := thisListener.Stack.Pop().(*ast.ExpressionAtom)
	if !okPop {
		thisListener.StopParse = true

		return
	}
	expr, okPop := thisListener.Stack.Peek().(ast.ExpressionAtomReceiver)
	if !okPop {
		thisListener.StopParse = true

		return
	}
	expressionAtm.Negated = ctx.NEGATION() != nil

	err := expr.AcceptExpressionAtom(thisListener.KnowledgeBase.WorkingMemory.AddExpressionAtom(expressionAtm))
	if err != nil {
		thisListener.StopParse = true
		thisListener.ErrorCallback.AddError(err)
	}
}

// EnterArrayMapSelector is called when production arrayMapSelector is entered.
func (thisListener *GruleV3ParserListener) EnterArrayMapSelector(ctx *grulev3.ArrayMapSelectorContext) {
	if thisListener.StopParse {

		return
	}
	sel := ast.NewArrayMapSelector()
	sel.GrlText = ctx.GetText()
	thisListener.Stack.Push(sel)
}

// ExitArrayMapSelector is called when production arrayMapSelector is exited.
func (thisListener *GruleV3ParserListener) ExitArrayMapSelector(ctx *grulev3.ArrayMapSelectorContext) {
	if thisListener.StopParse {

		return
	}
	sel, popOk := thisListener.Stack.Pop().(*ast.ArrayMapSelector)
	if !popOk {
		thisListener.StopParse = true

		return
	}
	receiver, popOk := thisListener.Stack.Peek().(ast.ArrayMapSelectorReceiver)
	if !popOk {
		thisListener.StopParse = true

		return
	}
	err := receiver.AcceptArrayMapSelector(sel)
	if err != nil {
		thisListener.StopParse = true
		thisListener.ErrorCallback.AddError(err)
	}
}

// EnterFunctionCall is called when production functionCall is entered.
func (thisListener *GruleV3ParserListener) EnterFunctionCall(ctx *grulev3.FunctionCallContext) {
	if thisListener.StopParse {

		return
	}
	fun := ast.NewFunctionCall()
	fun.FunctionName = ctx.SIMPLENAME().GetText()
	thisListener.Stack.Push(fun)
}

// ExitFunctionCall is called when production functionCall is exited.
func (thisListener *GruleV3ParserListener) ExitFunctionCall(ctx *grulev3.FunctionCallContext) {
	if thisListener.StopParse {

		return
	}
	fun, popOk := thisListener.Stack.Pop().(*ast.FunctionCall)
	if !popOk {
		thisListener.StopParse = true

		return
	}
	metRec, popOk := thisListener.Stack.Peek().(ast.FunctionCallReceiver)
	if !popOk {
		thisListener.StopParse = true

		return
	}
	err := metRec.AcceptFunctionCall(fun)
	if err != nil {
		thisListener.StopParse = true
		thisListener.ErrorCallback.AddError(err)
	}
}

// EnterArgumentList is called when production argumentList is entered.
func (thisListener *GruleV3ParserListener) EnterArgumentList(ctx *grulev3.ArgumentListContext) {
	if thisListener.StopParse {

		return
	}
	argList := ast.NewArgumentList()
	argList.GrlText = ctx.GetText()
	thisListener.Stack.Push(argList)
}

// ExitArgumentList is called when production argumentList is exited.
func (thisListener *GruleV3ParserListener) ExitArgumentList(ctx *grulev3.ArgumentListContext) {
	if thisListener.StopParse {

		return
	}
	argList, popOk := thisListener.Stack.Pop().(*ast.ArgumentList)
	if !popOk {
		thisListener.StopParse = true

		return
	}
	argListRec, popOk := thisListener.Stack.Peek().(ast.ArgumentListReceiver)
	if !popOk {
		thisListener.StopParse = true

		return
	}
	LoggerV3.Tracef("Adding Argument List To Receiver")
	err := argListRec.AcceptArgumentList(argList)
	if err != nil {
		thisListener.StopParse = true
		thisListener.ErrorCallback.AddError(err)
	}
}

// EnterVariable is called when production variable is entered.
func (thisListener *GruleV3ParserListener) EnterVariable(ctx *grulev3.VariableContext) {
	if thisListener.StopParse {

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
	thisListener.Stack.Push(vari)
}

// ExitVariable is called when production variable is exited.
func (thisListener *GruleV3ParserListener) ExitVariable(ctx *grulev3.VariableContext) {
	if thisListener.StopParse {

		return
	}
	vari, okPop := thisListener.Stack.Pop().(*ast.Variable)
	if !okPop {
		thisListener.StopParse = true

		return
	}
	variRec, okPop := thisListener.Stack.Peek().(ast.VariableReceiver)
	if !okPop {
		thisListener.StopParse = true

		return
	}

	err := variRec.AcceptVariable(thisListener.KnowledgeBase.WorkingMemory.AddVariable(vari))
	if err != nil {
		thisListener.StopParse = true
		thisListener.ErrorCallback.AddError(err)
	}
}

// EnterMemberVariable is called when production memberVariable is entered.
func (thisListener *GruleV3ParserListener) EnterMemberVariable(ctx *grulev3.MemberVariableContext) {}

// ExitMemberVariable is called when production memberVariable is exited.
func (thisListener *GruleV3ParserListener) ExitMemberVariable(ctx *grulev3.MemberVariableContext) {
	vari, ok := thisListener.Stack.Peek().(ast.MemberVariableReceiver)
	if !ok {
		thisListener.StopParse = true

		return
	}
	vari.AcceptMemberVariable(ctx.SIMPLENAME().GetText())
}

// EnterConstant is called when production constant is entered.
func (thisListener *GruleV3ParserListener) EnterConstant(ctx *grulev3.ConstantContext) {
	if thisListener.StopParse {

		return
	}
	cons := ast.NewConstant()
	cons.GrlText = ctx.GetText()
	thisListener.Stack.Push(cons)
}

// ExitConstant is called when production constant is exited.
func (thisListener *GruleV3ParserListener) ExitConstant(ctx *grulev3.ConstantContext) {
	if thisListener.StopParse {

		return
	}
	cons, popOk := thisListener.Stack.Pop().(*ast.Constant)
	if !popOk {
		thisListener.StopParse = true

		return
	}
	conRec, popOk := thisListener.Stack.Peek().(ast.ConstantReceiver)
	if !popOk {
		thisListener.StopParse = true

		return
	}
	if ctx.NIL_LITERAL() != nil {
		cons.IsNil = true
	}
	err := conRec.AcceptConstant(cons)
	if err != nil {
		thisListener.StopParse = true
		thisListener.ErrorCallback.AddError(err)
	}
}

// EnterStringLiteral is called when production stringLiteral is entered.
func (thisListener *GruleV3ParserListener) EnterStringLiteral(ctx *grulev3.StringLiteralContext) {
}

// ExitStringLiteral is called when production stringLiteral is exited.
func (thisListener *GruleV3ParserListener) ExitStringLiteral(ctx *grulev3.StringLiteralContext) {
	if thisListener.StopParse {

		return
	}
	dec, err := unquoteString(ctx.GetText())
	if err != nil {
		thisListener.ErrorCallback.AddError(fmt.Errorf("error parsing quoted string (%s): %s", ctx.GetText(), err.Error()))

		return
	}
	receiver, ok := thisListener.Stack.Peek().(ast.StringLiteralReceiver)
	if !ok {
		thisListener.StopParse = true

		return
	}
	receiver.AcceptStringLiteral(&ast.StringLiteral{String: dec})
}

// EnterBooleanLiteral is called when production booleanLiteral is entered.
func (thisListener *GruleV3ParserListener) EnterBooleanLiteral(ctx *grulev3.BooleanLiteralContext) {}

// ExitBooleanLiteral is called when production booleanLiteral is exited.
func (thisListener *GruleV3ParserListener) ExitBooleanLiteral(ctx *grulev3.BooleanLiteralContext) {
	if thisListener.StopParse {

		return
	}
	receiver, ok := thisListener.Stack.Peek().(ast.BooleanLiteralReceiver)
	if !ok {
		thisListener.StopParse = true

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
func (thisListener *GruleV3ParserListener) EnterIntegerLiteral(ctx *grulev3.IntegerLiteralContext) {}

// ExitIntegerLiteral is called when production integerLiteral is exited.
func (thisListener *GruleV3ParserListener) ExitIntegerLiteral(ctx *grulev3.IntegerLiteralContext) {
	lit := &ast.IntegerLiteral{}
	i, err := strconv.ParseInt(ctx.GetText(), 0, 64)
	if err != nil {
		thisListener.StopParse = true
		thisListener.ErrorCallback.AddError(err)
	} else {
		lit.Integer = i
	}
	receiver, ok := thisListener.Stack.Peek().(ast.IntegerLiteralReceiver)
	if !ok {
		thisListener.StopParse = true

		return
	}
	receiver.AcceptIntegerLiteral(lit)
}

// EnterFloatLiteral is called when production floatLiteral is entered.
func (thisListener *GruleV3ParserListener) EnterFloatLiteral(ctx *grulev3.FloatLiteralContext) {}

// ExitFloatLiteral is called when production floatLiteral is exited.
func (thisListener *GruleV3ParserListener) ExitFloatLiteral(ctx *grulev3.FloatLiteralContext) {
	lit := &ast.FloatLiteral{}
	i, err := strconv.ParseFloat(ctx.GetText(), 64)
	if err != nil {
		thisListener.StopParse = true
		thisListener.ErrorCallback.AddError(err)
	} else {
		lit.Float = i
	}
	receiver, ok := thisListener.Stack.Peek().(ast.FloatLiteralReceiver)
	if !ok {
		thisListener.StopParse = true

		return
	}
	receiver.AcceptFloatLiteral(lit)
}
