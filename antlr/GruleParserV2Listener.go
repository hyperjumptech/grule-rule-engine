package antlr

import (
	"fmt"
	"github.com/hyperjumptech/grule-rule-engine/ast/v2"
	"github.com/hyperjumptech/grule-rule-engine/logger"
	"reflect"
	"strconv"
	"strings"

	"github.com/antlr/antlr4/runtime/Go/antlr"
	"github.com/hyperjumptech/grule-rule-engine/antlr/parser/grulev2"
	"github.com/sirupsen/logrus"
)

var (
	// LoggerV2 is a logrus instance twith default fields for grule
	LoggerV2 = logger.Log.WithFields(logrus.Fields{
		"lib":    "grule",
		"struct": "GruleParserV2Listener",
	})
)

// NewGruleV2ParserListener create new instance of GruleV2ParserListener
func NewGruleV2ParserListener(KnowledgeBase *v2.KnowledgeBase, errorCallBack func(e error)) *GruleV2ParserListener {
	return &GruleV2ParserListener{
		PreviousNode:  make([]string, 0),
		ErrorCallback: errorCallBack,
		KnowledgeBase: KnowledgeBase,
		Stack:         newStack(),
	}
}

// GruleV2ParserListener is an implementation of logic to build the execution flow or execution graph as it
// defined within the knowledge base.
type GruleV2ParserListener struct {
	grulev2.Basegrulev2Listener
	PreviousNode []string

	Grl           *v2.Grl
	Stack         *stack
	StopParse     bool
	ErrorCallback func(e error)
	KnowledgeBase *v2.KnowledgeBase
}

// VisitTerminal is called when a terminal node is visited.
func (s *GruleV2ParserListener) VisitTerminal(node antlr.TerminalNode) {
	if s.StopParse {
		return
	}
	s.PreviousNode = append(s.PreviousNode, node.GetText())
	if len(s.PreviousNode) > 5 {
		s.PreviousNode = s.PreviousNode[1:]
	}
}

// VisitErrorNode is called when an error node is visited.
func (s *GruleV2ParserListener) VisitErrorNode(node antlr.ErrorNode) {
	LoggerV2.Errorf("GRL error, after '%v' and then unexpected '%s'", s.PreviousNode, node.GetText())
	s.StopParse = true
	s.ErrorCallback(fmt.Errorf("GRL error, after '%v' and then unexpected '%s'", s.PreviousNode, node.GetText()))
}

// EnterEveryRule is called when any rule is entered.
func (s *GruleV2ParserListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *GruleV2ParserListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterGrl is called when production grl is entered.
func (s *GruleV2ParserListener) EnterGrl(ctx *grulev2.GrlContext) {
	s.Grl = v2.NewGrl()
	s.Stack.Push(s.Grl)
}

// ExitGrl is called when production root is exited. The listener will instruct working memory re-index here.
func (s *GruleV2ParserListener) ExitGrl(ctx *grulev2.GrlContext) {
	if s.StopParse {
		return
	}
	_ = s.Stack.Pop().(*v2.Grl)
	for _, re := range s.Grl.RuleEntries {
		err := s.KnowledgeBase.AddRuleEntry(re)
		if err != nil {
			s.ErrorCallback(err)
		}
	}
}

// EnterRuleEntry is called when production ruleEntry is entered.
func (s *GruleV2ParserListener) EnterRuleEntry(ctx *grulev2.RuleEntryContext) {
	if s.StopParse {
		return
	}
	entry := v2.NewRuleEntry()
	entry.GrlText = ctx.GetText()
	s.Stack.Push(entry)
}

// ExitRuleEntry is called when production ruleEntry is exited.
func (s *GruleV2ParserListener) ExitRuleEntry(ctx *grulev2.RuleEntryContext) {
	if s.StopParse {
		return
	}
	entry := s.Stack.Pop().(*v2.RuleEntry)
	entryReceiver := s.Stack.Peek().(v2.RuleEntryReceiver)
	err := entryReceiver.ReceiveRuleEntry(entry)
	if err != nil {
		s.ErrorCallback(err)
	} else {
		LoggerV2.Debugf("Added RuleEntry : %s", entry.RuleName.SimpleName)
	}
}

// EnterSalience is called when production salience is entered.
func (s *GruleV2ParserListener) EnterSalience(ctx *grulev2.SalienceContext) {}

// ExitSalience is called when production salience is exited.
func (s *GruleV2ParserListener) ExitSalience(ctx *grulev2.SalienceContext) {
	if s.StopParse {
		return
	}
	dec := ctx.DecimalLiteral().GetText()
	salValue, _ := strconv.Atoi(dec)
	receiver := s.Stack.Peek().(v2.SalienceReceiver)
	err := receiver.AcceptSalience(v2.NewSalience(salValue))
	if err != nil {
		s.StopParse = true
		s.ErrorCallback(err)
	}
}

// EnterRuleName is called when production ruleName is entered.
func (s *GruleV2ParserListener) EnterRuleName(ctx *grulev2.RuleNameContext) {
}

// ExitRuleName is called when production ruleName is exited.
func (s *GruleV2ParserListener) ExitRuleName(ctx *grulev2.RuleNameContext) {
	if s.StopParse {
		return
	}
	receiver := s.Stack.Peek().(v2.RuleNameReceiver)
	err := receiver.AcceptRuleName(v2.NewRuleName(ctx.SIMPLENAME().GetText()))
	if err != nil {
		s.StopParse = true
		s.ErrorCallback(err)
	}
}

// EnterRuleDescription is called when production ruleDescription is entered.
func (s *GruleV2ParserListener) EnterRuleDescription(ctx *grulev2.RuleDescriptionContext) {
}

// ExitRuleDescription is called when production ruleDescription is exited.
func (s *GruleV2ParserListener) ExitRuleDescription(ctx *grulev2.RuleDescriptionContext) {
	if s.StopParse {
		return
	}
	receiver := s.Stack.Peek().(v2.RuleDescriptionReceiver)
	err := receiver.AcceptRuleDescription(v2.NewRuleDescription(ctx.GetText()[1 : len(ctx.GetText())-1]))
	if err != nil {
		s.StopParse = true
		s.ErrorCallback(err)
	}
}

// EnterWhenScope is called when production whenScope is entered.
func (s *GruleV2ParserListener) EnterWhenScope(ctx *grulev2.WhenScopeContext) {
	if s.StopParse {
		return
	}
	whenScope := v2.NewWhenScope()
	whenScope.GrlText = ctx.GetText()
	s.Stack.Push(whenScope)
}

// ExitWhenScope is called when production whenScope is exited.
func (s *GruleV2ParserListener) ExitWhenScope(ctx *grulev2.WhenScopeContext) {
	if s.StopParse {
		return
	}
	when := s.Stack.Pop().(*v2.WhenScope)
	receiver := s.Stack.Peek().(v2.WhenScopeReceiver)
	err := receiver.AcceptWhenScope(when)
	if err != nil {
		s.StopParse = true
		s.ErrorCallback(err)
	}
}

// EnterThenScope is called when production thenScope is entered.
func (s *GruleV2ParserListener) EnterThenScope(ctx *grulev2.ThenScopeContext) {
	if s.StopParse {
		return
	}
	then := v2.NewThenScope()
	then.GrlText = ctx.GetText()
	s.Stack.Push(then)
}

// ExitThenScope is called when production thenScope is exited.
func (s *GruleV2ParserListener) ExitThenScope(ctx *grulev2.ThenScopeContext) {
	if s.StopParse {
		return
	}
	then := s.Stack.Pop().(*v2.ThenScope)
	receiver := s.Stack.Peek().(v2.ThenScopeReceiver)
	err := receiver.AcceptThenScope(then)
	if err != nil {
		s.StopParse = true
		s.ErrorCallback(err)
	}
}

// EnterThenExpressionList is called when production thenExpressionList is entered.
func (s *GruleV2ParserListener) EnterThenExpressionList(ctx *grulev2.ThenExpressionListContext) {
	if s.StopParse {
		return
	}
	thenExpList := v2.NewThenExpressionList()
	thenExpList.GrlText = ctx.GetText()
	s.Stack.Push(thenExpList)
}

// ExitThenExpressionList is called when production thenExpressionList is exited.
func (s *GruleV2ParserListener) ExitThenExpressionList(ctx *grulev2.ThenExpressionListContext) {
	if s.StopParse {
		return
	}
	thenExpList := s.Stack.Pop().(*v2.ThenExpressionList)
	receiver := s.Stack.Peek().(v2.ThenExpressionListReceiver)
	err := receiver.AcceptThenExpressionList(thenExpList)
	if err != nil {
		s.StopParse = true
		s.ErrorCallback(err)
	}
}

// EnterThenExpression is called when production thenExpression is entered.
func (s *GruleV2ParserListener) EnterThenExpression(ctx *grulev2.ThenExpressionContext) {
	if s.StopParse {
		return
	}
	thenExpr := v2.NewThenExpression()
	thenExpr.GrlText = ctx.GetText()
	s.Stack.Push(thenExpr)
}

// ExitThenExpression is called when production thenExpression is exited.
func (s *GruleV2ParserListener) ExitThenExpression(ctx *grulev2.ThenExpressionContext) {
	if s.StopParse {
		return
	}
	thenExpr := s.Stack.Pop().(*v2.ThenExpression)

	receiver := s.Stack.Peek().(v2.ThenExpressionReceiver)
	err := receiver.AcceptThenExpression(thenExpr)
	if err != nil {
		s.StopParse = true
		s.ErrorCallback(err)
	}
}

// EnterAssignment is called when production assignment is entered.
func (s *GruleV2ParserListener) EnterAssignment(ctx *grulev2.AssignmentContext) {
	if s.StopParse {
		return
	}
	assign := v2.NewAssignment()
	assign.GrlText = ctx.GetText()
	s.Stack.Push(assign)
}

// ExitAssignment is called when production assignment is exited.
func (s *GruleV2ParserListener) ExitAssignment(ctx *grulev2.AssignmentContext) {
	if s.StopParse {
		return
	}
	assign := s.Stack.Pop().(*v2.Assignment)
	receiver := s.Stack.Peek().(v2.AssignmentReceiver)
	err := receiver.AcceptAssignment(assign)
	if err != nil {
		s.StopParse = true
		s.ErrorCallback(err)
	}
}

// EnterExpression is called when production expression is entered.
func (s *GruleV2ParserListener) EnterExpression(ctx *grulev2.ExpressionContext) {
	if s.StopParse {
		return
	}
	expr := v2.NewExpression()
	expr.GrlText = ctx.GetText()
	s.Stack.Push(expr)
}

// ExitExpression is called when production expression is exited.
func (s *GruleV2ParserListener) ExitExpression(ctx *grulev2.ExpressionContext) {
	if s.StopParse {
		return
	}
	expr := s.Stack.Pop().(*v2.Expression)
	exprRec := s.Stack.Peek().(v2.ExpressionReceiver)

	err := exprRec.AcceptExpression(s.KnowledgeBase.WorkingMemory.AddExpression(expr))
	if err != nil {
		s.StopParse = true
		s.ErrorCallback(err)
	}
}

// EnterMulDivOperators is called when production mulDivOperators is entered.
func (s *GruleV2ParserListener) EnterMulDivOperators(ctx *grulev2.MulDivOperatorsContext) {
	if s.StopParse {
		return
	}
	expr := s.Stack.Peek().(*v2.Expression)
	switch ctx.GetText() {
	case "*":
		expr.Operator = v2.OpMul
	case "/":
		expr.Operator = v2.OpDiv
	case "%":
		expr.Operator = v2.OpMod
	}
}

// ExitMulDivOperators is called when production mulDivOperators is exited.
func (s *GruleV2ParserListener) ExitMulDivOperators(ctx *grulev2.MulDivOperatorsContext) {}

// EnterAddMinusOperators is called when production addMinusOperators is entered.
func (s *GruleV2ParserListener) EnterAddMinusOperators(ctx *grulev2.AddMinusOperatorsContext) {
	if s.StopParse {
		return
	}
	expr := s.Stack.Peek().(*v2.Expression)
	switch ctx.GetText() {
	case "+":
		expr.Operator = v2.OpAdd
	case "-":
		expr.Operator = v2.OpSub
	case "|":
		expr.Operator = v2.OpBitOr
	case "&":
		expr.Operator = v2.OpBitAnd
	}
}

// ExitAddMinusOperators is called when production addMinusOperators is exited.
func (s *GruleV2ParserListener) ExitAddMinusOperators(ctx *grulev2.AddMinusOperatorsContext) {}

// EnterComparisonOperator is called when production comparisonOperator is entered.
func (s *GruleV2ParserListener) EnterComparisonOperator(ctx *grulev2.ComparisonOperatorContext) {
	if s.StopParse {
		return
	}
	expr := s.Stack.Peek().(*v2.Expression)
	switch ctx.GetText() {
	case "<":
		expr.Operator = v2.OpLT
	case "<=":
		expr.Operator = v2.OpLTE
	case ">":
		expr.Operator = v2.OpGT
	case ">=":
		expr.Operator = v2.OpGTE
	case "==":
		expr.Operator = v2.OpEq
	case "!=":
		expr.Operator = v2.OpNEq
	}
}

// ExitComparisonOperator is called when production comparisonOperator is exited.
func (s *GruleV2ParserListener) ExitComparisonOperator(ctx *grulev2.ComparisonOperatorContext) {}

// EnterAndLogicOperator is called when production andLogicOperator is entered.
func (s *GruleV2ParserListener) EnterAndLogicOperator(ctx *grulev2.AndLogicOperatorContext) {
	if s.StopParse {
		return
	}
	expr := s.Stack.Peek().(*v2.Expression)
	expr.Operator = v2.OpAnd
}

// ExitAndLogicOperator is called when production andLogicOperator is exited.
func (s *GruleV2ParserListener) ExitAndLogicOperator(ctx *grulev2.AndLogicOperatorContext) {}

// EnterOrLogicOperator is called when production orLogicOperator is entered.
func (s *GruleV2ParserListener) EnterOrLogicOperator(ctx *grulev2.OrLogicOperatorContext) {
	if s.StopParse {
		return
	}
	expr := s.Stack.Peek().(*v2.Expression)
	expr.Operator = v2.OpOr
}

// ExitOrLogicOperator is called when production orLogicOperator is exited.
func (s *GruleV2ParserListener) ExitOrLogicOperator(ctx *grulev2.OrLogicOperatorContext) {}

// EnterExpressionAtom is called when production expressionAtom is entered.
func (s *GruleV2ParserListener) EnterExpressionAtom(ctx *grulev2.ExpressionAtomContext) {
	if s.StopParse {
		return
	}
	atm := v2.NewExpressionAtom()
	atm.GrlText = ctx.GetText()
	s.Stack.Push(atm)
}

// ExitExpressionAtom is called when production expressionAtom is exited.
func (s *GruleV2ParserListener) ExitExpressionAtom(ctx *grulev2.ExpressionAtomContext) {
	if s.StopParse {
		return
	}
	atm := s.Stack.Pop().(*v2.ExpressionAtom)
	expr := s.Stack.Peek().(v2.ExpressionAtomReceiver)

	err := expr.AcceptExpressionAtom(s.KnowledgeBase.WorkingMemory.AddExpressionAtom(atm))
	if err != nil {
		s.StopParse = true
		s.ErrorCallback(err)
	}
}

// EnterArrayMapSelector is called when production arrayMapSelector is entered.
func (s *GruleV2ParserListener) EnterArrayMapSelector(ctx *grulev2.ArrayMapSelectorContext) {
	if s.StopParse {
		return
	}
	sel := v2.NewArrayMapSelector()
	sel.GrlText = ctx.GetText()
	s.Stack.Push(sel)
}

// ExitArrayMapSelector is called when production arrayMapSelector is exited.
func (s *GruleV2ParserListener) ExitArrayMapSelector(ctx *grulev2.ArrayMapSelectorContext) {
	if s.StopParse {
		return
	}
	sel := s.Stack.Pop().(*v2.ArrayMapSelector)
	receiver := s.Stack.Peek().(v2.ArrayMapSelectorReceiver)
	err := receiver.AcceptArrayMapSelector(sel)
	if err != nil {
		s.StopParse = true
		s.ErrorCallback(err)
	}
}

// EnterFunctionCall is called when production functionCall is entered.
func (s *GruleV2ParserListener) EnterFunctionCall(ctx *grulev2.FunctionCallContext) {
	if s.StopParse {
		return
	}
	fun := v2.NewFunctionCall()
	fun.GrlText = ctx.GetText()
	fun.FunctionName = ctx.SIMPLENAME().GetText()
	s.Stack.Push(fun)
}

// ExitFunctionCall is called when production functionCall is exited.
func (s *GruleV2ParserListener) ExitFunctionCall(ctx *grulev2.FunctionCallContext) {
	if s.StopParse {
		return
	}
	fun := s.Stack.Pop().(*v2.FunctionCall)
	metRec := s.Stack.Peek().(v2.FunctionCallReceiver)
	err := metRec.AcceptFunctionCall(fun)
	if err != nil {
		s.StopParse = true
		s.ErrorCallback(err)
	}
}

// EnterArgumentList is called when production argumentList is entered.
func (s *GruleV2ParserListener) EnterArgumentList(ctx *grulev2.ArgumentListContext) {
	if s.StopParse {
		return
	}
	argList := v2.NewArgumentList()
	argList.GrlText = ctx.GetText()
	s.Stack.Push(argList)
}

// ExitArgumentList is called when production argumentList is exited.
func (s *GruleV2ParserListener) ExitArgumentList(ctx *grulev2.ArgumentListContext) {
	if s.StopParse {
		return
	}
	argList := s.Stack.Pop().(*v2.ArgumentList)
	argListRec := s.Stack.Peek().(v2.ArgumentListReceiver)
	LoggerV2.Tracef("Adding Argument List To Receiver")
	err := argListRec.AcceptArgumentList(argList)
	if err != nil {
		s.StopParse = true
		s.ErrorCallback(err)
	}
}

// EnterVariable is called when production variable is entered.
func (s *GruleV2ParserListener) EnterVariable(ctx *grulev2.VariableContext) {
	if s.StopParse {
		return
	}
	vari := v2.NewVariable()
	if ctx.SIMPLENAME() != nil && len(ctx.SIMPLENAME().GetText()) > 0 {
		vari.Name = ctx.SIMPLENAME().GetText()
	}
	vari.GrlText = ctx.GetText()
	s.Stack.Push(vari)
}

// ExitVariable is called when production variable is exited.
func (s *GruleV2ParserListener) ExitVariable(ctx *grulev2.VariableContext) {
	if s.StopParse {
		return
	}
	vari := s.Stack.Pop().(*v2.Variable)
	variRec := s.Stack.Peek().(v2.VariableReceiver)

	err := variRec.AcceptVariable(s.KnowledgeBase.WorkingMemory.AddVariable(vari))
	if err != nil {
		s.StopParse = true
		s.ErrorCallback(err)
	}
}

// EnterConstant is called when production constant is entered.
func (s *GruleV2ParserListener) EnterConstant(ctx *grulev2.ConstantContext) {
	if s.StopParse {
		return
	}
	cons := v2.NewConstant()
	cons.GrlText = ctx.GetText()
	s.Stack.Push(cons)
}

// ExitConstant is called when production constant is exited.
func (s *GruleV2ParserListener) ExitConstant(ctx *grulev2.ConstantContext) {
	if s.StopParse {
		return
	}
	cons := s.Stack.Pop().(*v2.Constant)
	conRec := s.Stack.Peek().(v2.ConstantReceiver)
	err := conRec.AcceptConstant(cons)
	if err != nil {
		s.StopParse = true
		s.ErrorCallback(err)
	}
}

// EnterDecimalLiteral is called when production decimalLiteral is entered.
func (s *GruleV2ParserListener) EnterDecimalLiteral(ctx *grulev2.DecimalLiteralContext) {}

// ExitDecimalLiteral is called when production decimalLiteral is exited.
func (s *GruleV2ParserListener) ExitDecimalLiteral(ctx *grulev2.DecimalLiteralContext) {
	if s.StopParse {
		return
	}
	dec, _ := strconv.Atoi(ctx.GetText())
	if reflect.TypeOf(s.Stack.Peek()).String() == "*ast.Constant" {
		cons := s.Stack.Peek().(*v2.Constant)
		cons.Value = reflect.ValueOf(int64(dec))
	}
}

// EnterRealLiteral is called when production realLiteral is entered.
func (s *GruleV2ParserListener) EnterRealLiteral(ctx *grulev2.RealLiteralContext) {}

// ExitRealLiteral is called when production realLiteral is exited.
func (s *GruleV2ParserListener) ExitRealLiteral(ctx *grulev2.RealLiteralContext) {
	if s.StopParse {
		return
	}
	floa, _ := strconv.ParseFloat(ctx.GetText(), 64)
	if reflect.TypeOf(s.Stack.Peek()).String() == "*ast.Constant" {
		cons := s.Stack.Peek().(*v2.Constant)
		cons.Value = reflect.ValueOf(floa)
	}
}

// EnterStringLiteral is called when production stringLiteral is entered.
func (s *GruleV2ParserListener) EnterStringLiteral(ctx *grulev2.StringLiteralContext) {
}

// ExitStringLiteral is called when production stringLiteral is exited.
func (s *GruleV2ParserListener) ExitStringLiteral(ctx *grulev2.StringLiteralContext) {
	if s.StopParse {
		return
	}
	dec, err := unquoteString(ctx.GetText())
	if err != nil {
		s.ErrorCallback(fmt.Errorf("error parsing quoted string (%s): %s", ctx.GetText(), err.Error()))
		return
	}
	if reflect.TypeOf(s.Stack.Peek()).String() == "*ast.Constant" {
		cons := s.Stack.Peek().(*v2.Constant)
		cons.Value = reflect.ValueOf(dec)
	}
}

// EnterBooleanLiteral is called when production booleanLiteral is entered.
func (s *GruleV2ParserListener) EnterBooleanLiteral(ctx *grulev2.BooleanLiteralContext) {}

// ExitBooleanLiteral is called when production booleanLiteral is exited.
func (s *GruleV2ParserListener) ExitBooleanLiteral(ctx *grulev2.BooleanLiteralContext) {
	if s.StopParse {
		return
	}
	if reflect.TypeOf(s.Stack.Peek()).String() == "*ast.Constant" {
		cons := s.Stack.Peek().(*v2.Constant)
		switch strings.ToLower(ctx.GetText()) {
		case "true":
			cons.Value = reflect.ValueOf(true)
		case "false":
			cons.Value = reflect.ValueOf(false)
		}
	}
}
