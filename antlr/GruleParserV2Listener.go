package antlr

import (
	"fmt"
	"github.com/antlr/antlr4/runtime/Go/antlr"
	"github.com/golang-collections/collections/stack"
	parser2 "github.com/hyperjumptech/grule-rule-engine/antlr/parser/grulev2.g4"
	"github.com/hyperjumptech/grule-rule-engine/ast"
	"github.com/sirupsen/logrus"
	"reflect"
	"strconv"
	"strings"
)

var (
	// LoggerV2 is a logrus instance twith default fields for grule
	LoggerV2 = logrus.WithFields(logrus.Fields{
		"lib":    "grule",
		"struct": "GruleParserV2Listener",
	})
)

// NewGruleV2ParserListener create new instance of GruleV2ParserListener
func NewGruleV2ParserListener(knowleedgeBase *ast.KnowledgeBase, memory *ast.WorkingMemory, errorCallBack func(e error)) *GruleV2ParserListener {
	return &GruleV2ParserListener{
		PreviousNode:  make([]string, 0),
		VarNames:      make([]string, 0),
		ErrorCallback: errorCallBack,
		KnowledgeBase: knowleedgeBase,
		WorkingMemory: memory,
	}
}

// GruleV2ParserListener is an implementation of logic to build the execution flow or execution graph as it
// defined within the knowledge base.
type GruleV2ParserListener struct {
	parser2.Basegrulev2Listener
	PreviousNode []string
	VarNames     []string

	WorkingMemory *ast.WorkingMemory
	KnowledgeBase *ast.KnowledgeBase
	Stack         *stack.Stack
	StopParse     bool
	ErrorCallback func(e error)
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

// EnterRoot is called when production root is entered.
func (s *GruleV2ParserListener) EnterRoot(ctx *parser2.RootContext) {
	s.Stack = stack.New()
	s.Stack.Push(s.KnowledgeBase)
}

// ExitRoot is called when production root is exited.
func (s *GruleV2ParserListener) ExitRoot(ctx *parser2.RootContext) {
	if s.StopParse {
		return
	}
	s.Stack.Pop()
	if len(s.VarNames) > 0 {
		for _, varN := range s.VarNames {
			s.WorkingMemory.IndexVar(varN)
		}
	}
	//s.KnowledgeBase.MakeSnapshoot()
}

// EnterRuleEntry is called when production ruleEntry is entered.
func (s *GruleV2ParserListener) EnterRuleEntry(ctx *parser2.RuleEntryContext) {
	if s.StopParse {
		return
	}
	entry := ast.NewRuleEntry()
	entry.GrlText = ctx.GetText()
	s.Stack.Push(entry)
}

// ExitRuleEntry is called when production ruleEntry is exited.
func (s *GruleV2ParserListener) ExitRuleEntry(ctx *parser2.RuleEntryContext) {
	if s.StopParse {
		return
	}
	entry := s.Stack.Pop().(*ast.RuleEntry)
	knowledgeBase := s.Stack.Peek().(*ast.KnowledgeBase)
	err := knowledgeBase.AddRuleEntry(entry)
	if err != nil {
		s.ErrorCallback(err)
	} else {
		LoggerV2.Debugf("Added RuleEntry : %s", entry.Name)
	}
}

// EnterSalience is called when production salience is entered.
func (s *GruleV2ParserListener) EnterSalience(ctx *parser2.SalienceContext) {}

// ExitSalience is called when production salience is exited.
func (s *GruleV2ParserListener) ExitSalience(ctx *parser2.SalienceContext) {
	if s.StopParse {
		return
	}
	dec := ctx.DecimalLiteral().GetText()
	salValue, _ := strconv.Atoi(dec)
	entry := s.Stack.Peek().(*ast.RuleEntry)
	entry.Salience = salValue
}

// EnterRuleName is called when production ruleName is entered.
func (s *GruleV2ParserListener) EnterRuleName(ctx *parser2.RuleNameContext) {
	if s.StopParse {
		return
	}
	entry := s.Stack.Peek().(*ast.RuleEntry)
	entry.Name = ctx.SIMPLENAME().GetText()
}

// ExitRuleName is called when production ruleName is exited.
func (s *GruleV2ParserListener) ExitRuleName(ctx *parser2.RuleNameContext) {}

// EnterRuleDescription is called when production ruleDescription is entered.
func (s *GruleV2ParserListener) EnterRuleDescription(ctx *parser2.RuleDescriptionContext) {
	if s.StopParse {
		return
	}
	entry := s.Stack.Peek().(*ast.RuleEntry)
	entry.Description = ctx.GetText()[1 : len(ctx.GetText())-1]
}

// ExitRuleDescription is called when production ruleDescription is exited.
func (s *GruleV2ParserListener) ExitRuleDescription(ctx *parser2.RuleDescriptionContext) {}

// EnterWhenScope is called when production whenScope is entered.
func (s *GruleV2ParserListener) EnterWhenScope(ctx *parser2.WhenScopeContext) {
	if s.StopParse {
		return
	}
	whenScope := ast.NewWhenScope()
	whenScope.GrlText = ctx.GetText()
	s.Stack.Push(whenScope)
}

// ExitWhenScope is called when production whenScope is exited.
func (s *GruleV2ParserListener) ExitWhenScope(ctx *parser2.WhenScopeContext) {
	if s.StopParse {
		return
	}
	when := s.Stack.Pop().(*ast.WhenScope)
	entry := s.Stack.Peek().(*ast.RuleEntry)
	entry.WhenScope = when
}

// EnterThenScope is called when production thenScope is entered.
func (s *GruleV2ParserListener) EnterThenScope(ctx *parser2.ThenScopeContext) {
	if s.StopParse {
		return
	}
	then := ast.NewThenScope()
	then.GrlText = ctx.GetText()
	s.Stack.Push(then)
}

// ExitThenScope is called when production thenScope is exited.
func (s *GruleV2ParserListener) ExitThenScope(ctx *parser2.ThenScopeContext) {
	if s.StopParse {
		return
	}
	then := s.Stack.Pop().(*ast.ThenScope)
	entry := s.Stack.Peek().(*ast.RuleEntry)
	entry.ThenScope = then
}

// EnterThenExpressionList is called when production thenExpressionList is entered.
func (s *GruleV2ParserListener) EnterThenExpressionList(ctx *parser2.ThenExpressionListContext) {
	if s.StopParse {
		return
	}
	thenExpList := ast.NewThenExpressionList()
	thenExpList.GrlText = ctx.GetText()
	s.Stack.Push(thenExpList)
}

// ExitThenExpressionList is called when production thenExpressionList is exited.
func (s *GruleV2ParserListener) ExitThenExpressionList(ctx *parser2.ThenExpressionListContext) {
	if s.StopParse {
		return
	}
	thenExpList := s.Stack.Pop().(*ast.ThenExpressionList)
	then := s.Stack.Peek().(*ast.ThenScope)
	then.ThenExpressionList = thenExpList
}

// EnterThenExpression is called when production thenExpression is entered.
func (s *GruleV2ParserListener) EnterThenExpression(ctx *parser2.ThenExpressionContext) {
	if s.StopParse {
		return
	}
	thenExpr := ast.NewThenExpression()
	thenExpr.GrlText = ctx.GetText()
	s.Stack.Push(thenExpr)
}

// ExitThenExpression is called when production thenExpression is exited.
func (s *GruleV2ParserListener) ExitThenExpression(ctx *parser2.ThenExpressionContext) {
	if s.StopParse {
		return
	}
	thenExpr := s.Stack.Pop().(*ast.ThenExpression)
	thenExprList := s.Stack.Peek().(*ast.ThenExpressionList)
	thenExprList.ThenExpressions = append(thenExprList.ThenExpressions, thenExpr)
}

// EnterAssignment is called when production assignment is entered.
func (s *GruleV2ParserListener) EnterAssignment(ctx *parser2.AssignmentContext) {
	if s.StopParse {
		return
	}
	assign := ast.NewAssignment()
	assign.GrlText = ctx.GetText()
	s.Stack.Push(assign)
}

// ExitAssignment is called when production assignment is exited.
func (s *GruleV2ParserListener) ExitAssignment(ctx *parser2.AssignmentContext) {
	if s.StopParse {
		return
	}
	assign := s.Stack.Pop().(*ast.Assignment)
	thenExpr := s.Stack.Peek().(*ast.ThenExpression)
	thenExpr.Assignment = assign
}

// EnterExpression is called when production expression is entered.
func (s *GruleV2ParserListener) EnterExpression(ctx *parser2.ExpressionContext) {
	if s.StopParse {
		return
	}
	expr := ast.NewExpression()
	expr.GrlText = ctx.GetText()
	s.Stack.Push(expr)
}

// ExitExpression is called when production expression is exited.
func (s *GruleV2ParserListener) ExitExpression(ctx *parser2.ExpressionContext) {
	if s.StopParse {
		return
	}
	expr := s.Stack.Pop().(*ast.Expression)
	nexper, added := s.WorkingMemory.Add(expr)
	if LoggerV2.Level == logrus.TraceLevel {
		LoggerV2.Tracef("expression %s added to working memory : %v", nexper.GetSnapshot(), added)
	}
	exprRec := s.Stack.Peek().(ast.ExpressionReceiver)
	err := exprRec.AcceptExpression(nexper)
	if err != nil {
		s.StopParse = true
		s.ErrorCallback(err)
	}
}

// EnterMulDivOperators is called when production mulDivOperators is entered.
func (s *GruleV2ParserListener) EnterMulDivOperators(ctx *parser2.MulDivOperatorsContext) {
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
func (s *GruleV2ParserListener) ExitMulDivOperators(ctx *parser2.MulDivOperatorsContext) {}

// EnterAddMinusOperators is called when production addMinusOperators is entered.
func (s *GruleV2ParserListener) EnterAddMinusOperators(ctx *parser2.AddMinusOperatorsContext) {
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
func (s *GruleV2ParserListener) ExitAddMinusOperators(ctx *parser2.AddMinusOperatorsContext) {}

// EnterComparisonOperator is called when production comparisonOperator is entered.
func (s *GruleV2ParserListener) EnterComparisonOperator(ctx *parser2.ComparisonOperatorContext) {
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
func (s *GruleV2ParserListener) ExitComparisonOperator(ctx *parser2.ComparisonOperatorContext) {}

// EnterAndLogicOperator is called when production andLogicOperator is entered.
func (s *GruleV2ParserListener) EnterAndLogicOperator(ctx *parser2.AndLogicOperatorContext) {
	if s.StopParse {
		return
	}
	expr := s.Stack.Peek().(*ast.Expression)
	expr.Operator = ast.OpAnd
}

// ExitAndLogicOperator is called when production andLogicOperator is exited.
func (s *GruleV2ParserListener) ExitAndLogicOperator(ctx *parser2.AndLogicOperatorContext) {}

// EnterOrLogicOperator is called when production orLogicOperator is entered.
func (s *GruleV2ParserListener) EnterOrLogicOperator(ctx *parser2.OrLogicOperatorContext) {
	if s.StopParse {
		return
	}
	expr := s.Stack.Peek().(*ast.Expression)
	expr.Operator = ast.OpOr
}

// ExitOrLogicOperator is called when production orLogicOperator is exited.
func (s *GruleV2ParserListener) ExitOrLogicOperator(ctx *parser2.OrLogicOperatorContext) {}

// EnterExpressionAtom is called when production expressionAtom is entered.
func (s *GruleV2ParserListener) EnterExpressionAtom(ctx *parser2.ExpressionAtomContext) {
	if s.StopParse {
		return
	}
	atm := ast.NewExpressionAtom()
	atm.GrlText = ctx.GetText()
	s.Stack.Push(atm)
}

// ExitExpressionAtom is called when production expressionAtom is exited.
func (s *GruleV2ParserListener) ExitExpressionAtom(ctx *parser2.ExpressionAtomContext) {
	if s.StopParse {
		return
	}
	atm := s.Stack.Pop().(*ast.ExpressionAtom)
	expr := s.Stack.Peek().(*ast.Expression)
	expr.ExpressionAtom = atm
}

// EnterMethodCall is called when production methodCall is entered.
func (s *GruleV2ParserListener) EnterMethodCall(ctx *parser2.MethodCallContext) {
	if s.StopParse {
		return
	}
	met := ast.NewMethodCall()
	met.GrlText = ctx.GetText()
	met.MethodName = ctx.DOTTEDNAME().GetText()
	s.Stack.Push(met)
}

// ExitMethodCall is called when production methodCall is exited.
func (s *GruleV2ParserListener) ExitMethodCall(ctx *parser2.MethodCallContext) {
	if s.StopParse {
		return
	}
	met := s.Stack.Pop().(*ast.MethodCall)
	metRec := s.Stack.Peek().(ast.MethodCallReceiver)
	err := metRec.AcceptMethodCall(met)
	if err != nil {
		s.StopParse = true
		s.ErrorCallback(err)
	}
}

// EnterFunctionCall is called when production functionCall is entered.
func (s *GruleV2ParserListener) EnterFunctionCall(ctx *parser2.FunctionCallContext) {
	if s.StopParse {
		return
	}
	fun := ast.NewFunctionCall()
	fun.GrlText = ctx.GetText()
	fun.FunctionName = ctx.SIMPLENAME().GetText()
	s.Stack.Push(fun)
}

// ExitFunctionCall is called when production functionCall is exited.
func (s *GruleV2ParserListener) ExitFunctionCall(ctx *parser2.FunctionCallContext) {
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
func (s *GruleV2ParserListener) EnterArgumentList(ctx *parser2.ArgumentListContext) {
	if s.StopParse {
		return
	}
	argList := ast.NewArgumentList()
	argList.GrlText = ctx.GetText()
	s.Stack.Push(argList)
}

// ExitArgumentList is called when production argumentList is exited.
func (s *GruleV2ParserListener) ExitArgumentList(ctx *parser2.ArgumentListContext) {
	if s.StopParse {
		return
	}
	argList := s.Stack.Pop().(*ast.ArgumentList)
	argListRec := s.Stack.Peek().(ast.ArgumentListReceiver)
	LoggerV2.Tracef("Adding Argument List To Receiver")
	argListRec.AcceptArgumentList(argList)
}

// EnterVariable is called when production variable is entered.
func (s *GruleV2ParserListener) EnterVariable(ctx *parser2.VariableContext) {
	if s.StopParse {
		return
	}
	varName := ctx.GetText()
	vari := ast.NewVariable(varName)
	vari.GrlText = varName
	varFound := false
	for _, vn := range s.VarNames {
		if vn == varName {
			varFound = true
			break
		}
	}
	if !varFound {
		s.VarNames = append(s.VarNames, varName)
	}
	s.Stack.Push(vari)
}

// ExitVariable is called when production variable is exited.
func (s *GruleV2ParserListener) ExitVariable(ctx *parser2.VariableContext) {
	if s.StopParse {
		return
	}
	vari := s.Stack.Pop().(*ast.Variable)
	variRec := s.Stack.Peek().(ast.VariableReceiver)
	err := variRec.AcceptVariable(vari)
	if err != nil {
		s.StopParse = true
		s.ErrorCallback(err)
	}
}

// EnterConstant is called when production constant is entered.
func (s *GruleV2ParserListener) EnterConstant(ctx *parser2.ConstantContext) {
	if s.StopParse {
		return
	}
	cons := ast.NewConstant()
	cons.GrlText = ctx.GetText()
	s.Stack.Push(cons)
}

// ExitConstant is called when production constant is exited.
func (s *GruleV2ParserListener) ExitConstant(ctx *parser2.ConstantContext) {
	if s.StopParse {
		return
	}
	cons := s.Stack.Pop().(*ast.Constant)
	conRec := s.Stack.Peek().(ast.ConstantReceiver)
	conRec.AcceptConstant(cons)
}

// EnterDecimalLiteral is called when production decimalLiteral is entered.
func (s *GruleV2ParserListener) EnterDecimalLiteral(ctx *parser2.DecimalLiteralContext) {}

// ExitDecimalLiteral is called when production decimalLiteral is exited.
func (s *GruleV2ParserListener) ExitDecimalLiteral(ctx *parser2.DecimalLiteralContext) {
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
func (s *GruleV2ParserListener) EnterRealLiteral(ctx *parser2.RealLiteralContext) {}

// ExitRealLiteral is called when production realLiteral is exited.
func (s *GruleV2ParserListener) ExitRealLiteral(ctx *parser2.RealLiteralContext) {
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
func (s *GruleV2ParserListener) EnterStringLiteral(ctx *parser2.StringLiteralContext) {
}

// ExitStringLiteral is called when production stringLiteral is exited.
func (s *GruleV2ParserListener) ExitStringLiteral(ctx *parser2.StringLiteralContext) {
	if s.StopParse {
		return
	}
	dec := ctx.GetText()[1 : len(ctx.GetText())-1]
	if reflect.TypeOf(s.Stack.Peek()).String() == "*ast.Constant" {
		cons := s.Stack.Peek().(*ast.Constant)
		cons.Value = reflect.ValueOf(dec)
	}
}

// EnterBooleanLiteral is called when production booleanLiteral is entered.
func (s *GruleV2ParserListener) EnterBooleanLiteral(ctx *parser2.BooleanLiteralContext) {}

// ExitBooleanLiteral is called when production booleanLiteral is exited.
func (s *GruleV2ParserListener) ExitBooleanLiteral(ctx *parser2.BooleanLiteralContext) {
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
