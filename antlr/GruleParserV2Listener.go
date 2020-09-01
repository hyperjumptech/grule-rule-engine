package antlr

import (
	"fmt"
	"reflect"
	"strconv"
	"strings"
	"unicode/utf8"

	"github.com/antlr/antlr4/runtime/Go/antlr"
	"github.com/hyperjumptech/grule-rule-engine/antlr/parser/grulev2"
	"github.com/hyperjumptech/grule-rule-engine/ast"
	"github.com/sirupsen/logrus"
)

var (
	// LoggerV2 is a logrus instance twith default fields for grule
	LoggerV2 = logrus.WithFields(logrus.Fields{
		"lib":    "grule",
		"struct": "GruleParserV2Listener",
	})
)

// NewGruleV2ParserListener create new instance of GruleV2ParserListener
func NewGruleV2ParserListener(KnowledgeBase *ast.KnowledgeBase, errorCallBack func(e error)) *GruleV2ParserListener {
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

	Grl           *ast.Grl
	Stack         *stack
	StopParse     bool
	ErrorCallback func(e error)
	KnowledgeBase *ast.KnowledgeBase
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
	s.Grl = ast.NewGrl()
	s.Stack.Push(s.Grl)
}

// ExitGrl is called when production root is exited. The listener will instruct working memory re-index here.
func (s *GruleV2ParserListener) ExitGrl(ctx *grulev2.GrlContext) {
	if s.StopParse {
		return
	}
	_ = s.Stack.Pop().(*ast.Grl)
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
	entry := ast.NewRuleEntry()
	entry.GrlText = ctx.GetText()
	s.Stack.Push(entry)
}

// ExitRuleEntry is called when production ruleEntry is exited.
func (s *GruleV2ParserListener) ExitRuleEntry(ctx *grulev2.RuleEntryContext) {
	if s.StopParse {
		return
	}
	entry := s.Stack.Pop().(*ast.RuleEntry)
	entryReceiver := s.Stack.Peek().(ast.RuleEntryReceiver)
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
	receiver := s.Stack.Peek().(ast.SalienceReceiver)
	err := receiver.AcceptSalience(ast.NewSalience(salValue))
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
	receiver := s.Stack.Peek().(ast.RuleNameReceiver)
	err := receiver.AcceptRuleName(ast.NewRuleName(ctx.SIMPLENAME().GetText()))
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
	receiver := s.Stack.Peek().(ast.RuleDescriptionReceiver)
	err := receiver.AcceptRuleDescription(ast.NewRuleDescription(ctx.GetText()[1 : len(ctx.GetText())-1]))
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
	whenScope := ast.NewWhenScope()
	whenScope.GrlText = ctx.GetText()
	s.Stack.Push(whenScope)
}

// ExitWhenScope is called when production whenScope is exited.
func (s *GruleV2ParserListener) ExitWhenScope(ctx *grulev2.WhenScopeContext) {
	if s.StopParse {
		return
	}
	when := s.Stack.Pop().(*ast.WhenScope)
	receiver := s.Stack.Peek().(ast.WhenScopeReceiver)
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
	then := ast.NewThenScope()
	then.GrlText = ctx.GetText()
	s.Stack.Push(then)
}

// ExitThenScope is called when production thenScope is exited.
func (s *GruleV2ParserListener) ExitThenScope(ctx *grulev2.ThenScopeContext) {
	if s.StopParse {
		return
	}
	then := s.Stack.Pop().(*ast.ThenScope)
	receiver := s.Stack.Peek().(ast.ThenScopeReceiver)
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
	thenExpList := ast.NewThenExpressionList()
	thenExpList.GrlText = ctx.GetText()
	s.Stack.Push(thenExpList)
}

// ExitThenExpressionList is called when production thenExpressionList is exited.
func (s *GruleV2ParserListener) ExitThenExpressionList(ctx *grulev2.ThenExpressionListContext) {
	if s.StopParse {
		return
	}
	thenExpList := s.Stack.Pop().(*ast.ThenExpressionList)
	receiver := s.Stack.Peek().(ast.ThenExpressionListReceiver)
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
	thenExpr := ast.NewThenExpression()
	thenExpr.GrlText = ctx.GetText()
	s.Stack.Push(thenExpr)
}

// ExitThenExpression is called when production thenExpression is exited.
func (s *GruleV2ParserListener) ExitThenExpression(ctx *grulev2.ThenExpressionContext) {
	if s.StopParse {
		return
	}
	thenExpr := s.Stack.Pop().(*ast.ThenExpression)

	receiver := s.Stack.Peek().(ast.ThenExpressionReceiver)
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
	assign := ast.NewAssignment()
	assign.GrlText = ctx.GetText()
	s.Stack.Push(assign)
}

// ExitAssignment is called when production assignment is exited.
func (s *GruleV2ParserListener) ExitAssignment(ctx *grulev2.AssignmentContext) {
	if s.StopParse {
		return
	}
	assign := s.Stack.Pop().(*ast.Assignment)
	receiver := s.Stack.Peek().(ast.AssignmentReceiver)
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
	expr := ast.NewExpression()
	expr.GrlText = ctx.GetText()
	s.Stack.Push(expr)
}

// ExitExpression is called when production expression is exited.
func (s *GruleV2ParserListener) ExitExpression(ctx *grulev2.ExpressionContext) {
	if s.StopParse {
		return
	}
	expr := s.Stack.Pop().(*ast.Expression)
	exprRec := s.Stack.Peek().(ast.ExpressionReceiver)

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
func (s *GruleV2ParserListener) ExitMulDivOperators(ctx *grulev2.MulDivOperatorsContext) {}

// EnterAddMinusOperators is called when production addMinusOperators is entered.
func (s *GruleV2ParserListener) EnterAddMinusOperators(ctx *grulev2.AddMinusOperatorsContext) {
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
func (s *GruleV2ParserListener) ExitAddMinusOperators(ctx *grulev2.AddMinusOperatorsContext) {}

// EnterComparisonOperator is called when production comparisonOperator is entered.
func (s *GruleV2ParserListener) EnterComparisonOperator(ctx *grulev2.ComparisonOperatorContext) {
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
func (s *GruleV2ParserListener) ExitComparisonOperator(ctx *grulev2.ComparisonOperatorContext) {}

// EnterAndLogicOperator is called when production andLogicOperator is entered.
func (s *GruleV2ParserListener) EnterAndLogicOperator(ctx *grulev2.AndLogicOperatorContext) {
	if s.StopParse {
		return
	}
	expr := s.Stack.Peek().(*ast.Expression)
	expr.Operator = ast.OpAnd
}

// ExitAndLogicOperator is called when production andLogicOperator is exited.
func (s *GruleV2ParserListener) ExitAndLogicOperator(ctx *grulev2.AndLogicOperatorContext) {}

// EnterOrLogicOperator is called when production orLogicOperator is entered.
func (s *GruleV2ParserListener) EnterOrLogicOperator(ctx *grulev2.OrLogicOperatorContext) {
	if s.StopParse {
		return
	}
	expr := s.Stack.Peek().(*ast.Expression)
	expr.Operator = ast.OpOr
}

// ExitOrLogicOperator is called when production orLogicOperator is exited.
func (s *GruleV2ParserListener) ExitOrLogicOperator(ctx *grulev2.OrLogicOperatorContext) {}

// EnterExpressionAtom is called when production expressionAtom is entered.
func (s *GruleV2ParserListener) EnterExpressionAtom(ctx *grulev2.ExpressionAtomContext) {
	if s.StopParse {
		return
	}
	atm := ast.NewExpressionAtom()
	atm.GrlText = ctx.GetText()
	s.Stack.Push(atm)
}

// ExitExpressionAtom is called when production expressionAtom is exited.
func (s *GruleV2ParserListener) ExitExpressionAtom(ctx *grulev2.ExpressionAtomContext) {
	if s.StopParse {
		return
	}
	atm := s.Stack.Pop().(*ast.ExpressionAtom)
	expr := s.Stack.Peek().(ast.ExpressionAtomReceiver)

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
	sel := ast.NewArrayMapSelector()
	sel.GrlText = ctx.GetText()
	s.Stack.Push(sel)
}

// ExitArrayMapSelector is called when production arrayMapSelector is exited.
func (s *GruleV2ParserListener) ExitArrayMapSelector(ctx *grulev2.ArrayMapSelectorContext) {
	if s.StopParse {
		return
	}
	sel := s.Stack.Pop().(*ast.ArrayMapSelector)
	receiver := s.Stack.Peek().(ast.ArrayMapSelectorReceiver)
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
	fun := ast.NewFunctionCall()
	fun.GrlText = ctx.GetText()
	fun.FunctionName = ctx.SIMPLENAME().GetText()
	s.Stack.Push(fun)
}

// ExitFunctionCall is called when production functionCall is exited.
func (s *GruleV2ParserListener) ExitFunctionCall(ctx *grulev2.FunctionCallContext) {
	if s.StopParse {
		return
	}
	fun := s.Stack.Pop().(*ast.FunctionCall)
	metRec := s.Stack.Peek().(ast.FunctionCallReceiver)
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
	argList := ast.NewArgumentList()
	argList.GrlText = ctx.GetText()
	s.Stack.Push(argList)
}

// ExitArgumentList is called when production argumentList is exited.
func (s *GruleV2ParserListener) ExitArgumentList(ctx *grulev2.ArgumentListContext) {
	if s.StopParse {
		return
	}
	argList := s.Stack.Pop().(*ast.ArgumentList)
	argListRec := s.Stack.Peek().(ast.ArgumentListReceiver)
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
	vari := ast.NewVariable()
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
	vari := s.Stack.Pop().(*ast.Variable)
	variRec := s.Stack.Peek().(ast.VariableReceiver)

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
	cons := ast.NewConstant()
	cons.GrlText = ctx.GetText()
	s.Stack.Push(cons)
}

// ExitConstant is called when production constant is exited.
func (s *GruleV2ParserListener) ExitConstant(ctx *grulev2.ConstantContext) {
	if s.StopParse {
		return
	}
	cons := s.Stack.Pop().(*ast.Constant)
	conRec := s.Stack.Peek().(ast.ConstantReceiver)
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
		cons := s.Stack.Peek().(*ast.Constant)
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
		cons := s.Stack.Peek().(*ast.Constant)
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
		cons := s.Stack.Peek().(*ast.Constant)
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
		cons := s.Stack.Peek().(*ast.Constant)
		switch strings.ToLower(ctx.GetText()) {
		case "true":
			cons.Value = reflect.ValueOf(true)
		case "false":
			cons.Value = reflect.ValueOf(false)
		}
	}
}

func unquoteString(s string) (string, error) {
	n := len(s)
	if n < 2 {
		return "", strconv.ErrSyntax
	}
	quote := s[0]
	if quote != s[n-1] {
		return "", strconv.ErrSyntax
	}
	s = s[1 : n-1]

	if quote != '"' && quote != '\'' {
		return "", strconv.ErrSyntax
	}

	if !contains(s, '\\') && !contains(s, quote) && utf8.ValidString(s) {
		return s, nil
	}

	var runeTmp [utf8.UTFMax]byte
	buf := make([]byte, 0, 3*len(s)/2)
	for len(s) > 0 {
		c, multibyte, ss, err := strconv.UnquoteChar(s, quote)
		if err != nil {
			return "", err
		}
		s = ss
		if c < utf8.RuneSelf || !multibyte {
			buf = append(buf, byte(c))
		} else {
			n := utf8.EncodeRune(runeTmp[:], c)
			buf = append(buf, runeTmp[:n]...)
		}
	}
	return string(buf), nil
}

func contains(s string, c byte) bool {
	return strings.IndexByte(s, c) != -1
}
