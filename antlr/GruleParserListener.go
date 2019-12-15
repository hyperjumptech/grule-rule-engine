package antlr

import (
	"fmt"
	"github.com/antlr/antlr4/runtime/Go/antlr"
	"github.com/golang-collections/collections/stack"
	"github.com/hyperjumptech/grule-rule-engine/antlr/parser"
	"github.com/hyperjumptech/grule-rule-engine/model"
	"github.com/juju/errors"
	"reflect"
	"strconv"
	"strings"
)

// NewGruleParserListener create a new instancce of GruleParserListener.
// This listener will walk in the Grule GRL file and invoke operations based on the
// context within the knowledge base.
func NewGruleParserListener(kbase *model.KnowledgeBase, errCallback func(e error)) *GruleParserListener {
	return &GruleParserListener{
		Stack:         stack.New(),
		KnowledgeBase: kbase,
		PreviousNode:  make([]string, 0),
		ErrorCallback: errCallback,
		StopParse:     false,
	}
}

// GruleParserListener is an implementation of logic to build the execution flow or execution graph as it
// defined within the knowledge base.
type GruleParserListener struct {
	parser.BasegruleListener
	PreviousNode []string

	KnowledgeBase *model.KnowledgeBase
	Stack         *stack.Stack
	StopParse     bool
	ErrorCallback func(e error)
}

// VisitTerminal is called when a terminal node is visited.
func (s *GruleParserListener) VisitTerminal(node antlr.TerminalNode) {
	if s.StopParse {
		return
	}
	s.PreviousNode = append(s.PreviousNode, node.GetText())
	if len(s.PreviousNode) > 5 {
		s.PreviousNode = s.PreviousNode[1:]
	}
}

// VisitErrorNode is called when an error node is visited.
func (s *GruleParserListener) VisitErrorNode(node antlr.ErrorNode) {
	s.StopParse = true
	s.ErrorCallback(errors.New(fmt.Sprintf("GRL error, after %v and then unexpected '%s'", s.PreviousNode, node.GetText())))
}

// EnterEveryRule is called when any engine is entered.
func (s *GruleParserListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any engine is exited.
func (s *GruleParserListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterRoot is called when production root is entered.
func (s *GruleParserListener) EnterRoot(ctx *parser.RootContext) {}

// ExitRoot is called when production root is exited.
func (s *GruleParserListener) ExitRoot(ctx *parser.RootContext) {}

// EnterRuleEntry is called when production ruleEntry is entered.
func (s *GruleParserListener) EnterRuleEntry(ctx *parser.RuleEntryContext) {
	if s.StopParse {
		return
	}
	entry := &model.RuleEntry{}
	s.Stack.Push(entry)
}

// ExitRuleEntry is called when production ruleEntry is exited.
func (s *GruleParserListener) ExitRuleEntry(ctx *parser.RuleEntryContext) {
	if s.StopParse {
		return
	}
	entry := s.Stack.Pop().(*model.RuleEntry)
	// check for duplicate engine.
	if _, ok := s.KnowledgeBase.RuleEntries[entry.RuleName]; ok {
		s.ErrorCallback(errors.Errorf("duplicate rule entry name '%s'", entry.RuleName))
		return
	}
	// if everything ok, add the engine entry.
	s.KnowledgeBase.RuleEntries[entry.RuleName] = entry
}

// EnterRuleName is called when production ruleName is entered.
func (s *GruleParserListener) EnterRuleName(ctx *parser.RuleNameContext) {
	if s.StopParse {
		return
	}
	ruleName := ctx.GetText()
	entry := s.Stack.Peek().(*model.RuleEntry)
	entry.RuleName = ruleName
}

// ExitRuleName is called when production ruleName is exited.
func (s *GruleParserListener) ExitRuleName(ctx *parser.RuleNameContext) {}

// EnterSalience is called when production salience is entered.
func (s *GruleParserListener) EnterSalience(ctx *parser.SalienceContext) {
	// salience were set by the decimal literal
}

// ExitSalience is called when production salience is exited.
func (s *GruleParserListener) ExitSalience(ctx *parser.SalienceContext) {}

// EnterRuleDescription is called when production ruleDescription is entered.
func (s *GruleParserListener) EnterRuleDescription(ctx *parser.RuleDescriptionContext) {
	if s.StopParse {
		return
	}
	ruleDescription := ctx.GetText()
	entry := s.Stack.Peek().(*model.RuleEntry)
	entry.RuleDescription = ruleDescription
}

// ExitRuleDescription is called when production ruleDescription is exited.
func (s *GruleParserListener) ExitRuleDescription(ctx *parser.RuleDescriptionContext) {}

// EnterWhenScope is called when production whenScope is entered.
func (s *GruleParserListener) EnterWhenScope(ctx *parser.WhenScopeContext) {
	if s.StopParse {
		return
	}
	whenScope := &model.WhenScope{}
	s.Stack.Push(whenScope)
}

// ExitWhenScope is called when production whenScope is exited.
func (s *GruleParserListener) ExitWhenScope(ctx *parser.WhenScopeContext) {
	if s.StopParse {
		return
	}
	whenScope := s.Stack.Pop().(*model.WhenScope)
	ruleEntry := s.Stack.Peek().(*model.RuleEntry)
	ruleEntry.WhenScope = whenScope
}

// EnterThenScope is called when production thenScope is entered.
func (s *GruleParserListener) EnterThenScope(ctx *parser.ThenScopeContext) {
	if s.StopParse {
		return
	}
	thenScope := &model.ThenScope{}
	s.Stack.Push(thenScope)
}

// ExitThenScope is called when production thenScope is exited.
func (s *GruleParserListener) ExitThenScope(ctx *parser.ThenScopeContext) {
	if s.StopParse {
		return
	}
	thenScope := s.Stack.Pop().(*model.ThenScope)
	ruleEntry := s.Stack.Peek().(*model.RuleEntry)
	ruleEntry.ThenScope = thenScope
}

// EnterAssignExpressions is called when production assignExpressions is entered.
func (s *GruleParserListener) EnterAssignExpressions(ctx *parser.AssignExpressionsContext) {
	if s.StopParse {
		return
	}
	assigns := &model.AssignExpressions{
		ExpressionList: make([]*model.AssignExpression, 0),
	}
	s.Stack.Push(assigns)
}

// ExitAssignExpressions is called when production assignExpressions is exited.
func (s *GruleParserListener) ExitAssignExpressions(ctx *parser.AssignExpressionsContext) {
	if s.StopParse {
		return
	}
	assigns := s.Stack.Pop().(*model.AssignExpressions)
	thenScope := s.Stack.Peek().(*model.ThenScope)
	thenScope.AssignExpressions = assigns
}

// EnterAssignExpression is called when production assignExpression is entered.
func (s *GruleParserListener) EnterAssignExpression(ctx *parser.AssignExpressionContext) {
	if s.StopParse {
		return
	}
	assign := &model.AssignExpression{}
	s.Stack.Push(assign)
}

// ExitAssignExpression is called when production assignExpression is exited.
func (s *GruleParserListener) ExitAssignExpression(ctx *parser.AssignExpressionContext) {
	if s.StopParse {
		return
	}
	assign := s.Stack.Pop().(*model.AssignExpression)
	assigns := s.Stack.Peek().(*model.AssignExpressions)
	assigns.ExpressionList = append(assigns.ExpressionList, assign)
}

// EnterAssignment is called when production assignment is entered.
func (s *GruleParserListener) EnterAssignment(ctx *parser.AssignmentContext) {
	if s.StopParse {
		return
	}
	assignment := &model.Assignment{}
	s.Stack.Push(assignment)
}

// ExitAssignment is called when production assignment is exited.
func (s *GruleParserListener) ExitAssignment(ctx *parser.AssignmentContext) {
	if s.StopParse {
		return
	}
	assignment := s.Stack.Pop().(*model.Assignment)
	assign := s.Stack.Peek().(*model.AssignExpression)
	assign.Assignment = assignment
}

// EnterExpression is called when production expression is entered.
func (s *GruleParserListener) EnterExpression(ctx *parser.ExpressionContext) {
	if s.StopParse {
		return
	}
	expression := &model.Expression{}
	s.Stack.Push(expression)
}

// ExitExpression is called when production expression is exited.
func (s *GruleParserListener) ExitExpression(ctx *parser.ExpressionContext) {
	if s.StopParse {
		return
	}
	expr := s.Stack.Pop().(*model.Expression)
	holder := s.Stack.Peek().(model.ExpressionHolder)
	err := holder.AcceptExpression(expr)
	if err != nil {
		s.ErrorCallback(err)
	}
}

// EnterPredicate is called when production predicate is entered.
func (s *GruleParserListener) EnterPredicate(ctx *parser.PredicateContext) {
	if s.StopParse {
		return
	}
	predicate := &model.Predicate{}
	s.Stack.Push(predicate)
}

// ExitPredicate is called when production predicate is exited.
func (s *GruleParserListener) ExitPredicate(ctx *parser.PredicateContext) {
	if s.StopParse {
		return
	}
	predicate := s.Stack.Pop().(*model.Predicate)
	expr := s.Stack.Peek().(*model.Expression)
	expr.Predicate = predicate
}

// EnterExpressionAtom is called when production expressionAtom is entered.
func (s *GruleParserListener) EnterExpressionAtom(ctx *parser.ExpressionAtomContext) {
	if s.StopParse {
		return
	}
	exprAtom := &model.ExpressionAtom{
		Text: ctx.GetText(),
	}
	s.Stack.Push(exprAtom)
}

// ExitExpressionAtom is called when production expressionAtom is exited.
func (s *GruleParserListener) ExitExpressionAtom(ctx *parser.ExpressionAtomContext) {
	if s.StopParse {
		return
	}
	//fmt.Println(ctx.GetText())
	exprAtom := s.Stack.Pop().(*model.ExpressionAtom)
	holder := s.Stack.Peek().(model.ExpressionAtomHolder)
	err := holder.AcceptExpressionAtom(exprAtom)
	if err != nil {
		s.ErrorCallback(err)
	}
}

// EnterMethodCall is called when production methodCall is entered.
func (s *GruleParserListener) EnterMethodCall(ctx *parser.MethodCallContext) {
	if s.StopParse {
		return
	}
	funcCall := &model.MethodCall{
		MethodName: ctx.DOTTEDNAME().GetText(),
	}
	s.Stack.Push(funcCall)
}

// ExitMethodCall is called when production methodCall is exited.
func (s *GruleParserListener) ExitMethodCall(ctx *parser.MethodCallContext) {
	if s.StopParse {
		return
	}
	methodCall := s.Stack.Pop().(*model.MethodCall)
	holder := s.Stack.Peek().(model.MethodCallHolder)
	err := holder.AcceptMethodCall(methodCall)
	if err != nil {
		fmt.Printf("Got error %s\n", err)
		s.ErrorCallback(err)
	}
}

// EnterFunctionCall is called when production functionCall is entered.
func (s *GruleParserListener) EnterFunctionCall(ctx *parser.FunctionCallContext) {
	if s.StopParse {
		return
	}
	funcCall := &model.FunctionCall{
		FunctionName: ctx.SIMPLENAME().GetText(),
	}
	s.Stack.Push(funcCall)
}

// ExitFunctionCall is called when production functionCall is exited.
func (s *GruleParserListener) ExitFunctionCall(ctx *parser.FunctionCallContext) {
	if s.StopParse {
		return
	}
	funcCall := s.Stack.Pop().(*model.FunctionCall)
	holder := s.Stack.Peek().(model.FunctionCallHolder)
	err := holder.AcceptFunctionCall(funcCall)
	if err != nil {
		s.ErrorCallback(err)
	}
}

// EnterFunctionArgs is called when production functionArgs is entered.
func (s *GruleParserListener) EnterFunctionArgs(ctx *parser.FunctionArgsContext) {
	if s.StopParse {
		return
	}
	funcArg := &model.FunctionArgument{
		Arguments: make([]*model.ArgumentHolder, 0),
	}
	s.Stack.Push(funcArg)
}

// ExitFunctionArgs is called when production functionArgs is exited.
func (s *GruleParserListener) ExitFunctionArgs(ctx *parser.FunctionArgsContext) {
	if s.StopParse {
		return
	}
	funcArgs := s.Stack.Pop().(*model.FunctionArgument)
	// return immediately when there's an error
	argHolder := s.Stack.Peek().(model.FunctionArgumentHolder)
	err := argHolder.AcceptFunctionArgument(funcArgs)
	if err != nil {
		s.ErrorCallback(err)
	}
}

// EnterLogicalOperator is called when production logicalOperator is entered.
func (s *GruleParserListener) EnterLogicalOperator(ctx *parser.LogicalOperatorContext) {
}

// ExitLogicalOperator is called when production logicalOperator is exited.
func (s *GruleParserListener) ExitLogicalOperator(ctx *parser.LogicalOperatorContext) {
	if s.StopParse {
		return
	}
	expr := s.Stack.Peek().(*model.Expression)
	switch ctx.GetText() {
	case "&&":
		expr.LogicalOperator = model.LogicalOperatorAnd
	case "||":
		expr.LogicalOperator = model.LogicalOperatorOr
	default:
		s.ErrorCallback(errors.Errorf("unknown logical operator %s", ctx.GetText()))
	}
}

// EnterVariable is called when production variable is entered.
func (s *GruleParserListener) EnterVariable(ctx *parser.VariableContext) {}

// ExitVariable is called when production variable is exited.
func (s *GruleParserListener) ExitVariable(ctx *parser.VariableContext) {
	if s.StopParse {
		return
	}
	varName := ctx.GetText()
	holder := s.Stack.Peek().(model.VariableHolder)
	err := holder.AcceptVariable(varName)
	if err != nil {
		s.ErrorCallback(err)
	}
}

// EnterMathOperator is called when production mathOperator is entered.
func (s *GruleParserListener) EnterMathOperator(ctx *parser.MathOperatorContext) {
}

// ExitMathOperator is called when production mathOperator is exited.
func (s *GruleParserListener) ExitMathOperator(ctx *parser.MathOperatorContext) {
	if s.StopParse {
		return
	}
	expr := s.Stack.Peek().(*model.ExpressionAtom)
	switch ctx.GetText() {
	case "+":
		expr.MathOperator = model.MathOperatorPlus
	case "-":
		expr.MathOperator = model.MathOperatorMinus
	case "/":
		expr.MathOperator = model.MathOperatorDiv
	case "*":
		expr.MathOperator = model.MathOperatorMul
	default:
		s.ErrorCallback(errors.Errorf("unknown mathematic operator %s", ctx.GetText()))
	}
}

// EnterComparisonOperator is called when production comparisonOperator is entered.
func (s *GruleParserListener) EnterComparisonOperator(ctx *parser.ComparisonOperatorContext) {}

// ExitComparisonOperator is called when production comparisonOperator is exited.
func (s *GruleParserListener) ExitComparisonOperator(ctx *parser.ComparisonOperatorContext) {
	if s.StopParse {
		return
	}
	predicate := s.Stack.Peek().(*model.Predicate)
	switch ctx.GetText() {
	case "==":
		predicate.ComparisonOperator = model.ComparisonOperatorEQ
	case "!=":
		predicate.ComparisonOperator = model.ComparisonOperatorNEQ
	case "<":
		predicate.ComparisonOperator = model.ComparisonOperatorLT
	case "<=":
		predicate.ComparisonOperator = model.ComparisonOperatorLTE
	case ">":
		predicate.ComparisonOperator = model.ComparisonOperatorGT
	case ">=":
		predicate.ComparisonOperator = model.ComparisonOperatorGTE
	default:
		s.ErrorCallback(errors.Errorf("unknown comparison operator %s", ctx.GetText()))
	}
}

// EnterConstant is called when production constant is entered.
func (s *GruleParserListener) EnterConstant(ctx *parser.ConstantContext) {
	if s.StopParse {
		return
	}
	cons := &model.Constant{}
	s.Stack.Push(cons)
}

// ExitConstant is called when production constant is exited.
func (s *GruleParserListener) ExitConstant(ctx *parser.ConstantContext) {
	if s.StopParse {
		return
	}
	cons := s.Stack.Pop().(*model.Constant)
	if ctx.NULL_LITERAL() != nil {
		if ctx.NOT() != nil {
			cons.ConstantValue = reflect.ValueOf("")
		} else {
			cons.ConstantValue = reflect.ValueOf(nil)
		}
	}

	holder := s.Stack.Peek().(model.ConstantHolder)
	err := holder.AcceptConstant(cons)
	if err != nil {
		s.ErrorCallback(err)
	}
}

// EnterDecimalLiteral is called when production decimalLiteral is entered.
func (s *GruleParserListener) EnterDecimalLiteral(ctx *parser.DecimalLiteralContext) {}

// ExitDecimalLiteral is called when production decimalLiteral is exited.
func (s *GruleParserListener) ExitDecimalLiteral(ctx *parser.DecimalLiteralContext) {
	if s.StopParse {
		return
	}
	decHold := s.Stack.Peek().(model.DecimalHolder)
	i64, err := strconv.ParseInt(ctx.GetText(), 10, 64)
	if err != nil {
		s.ErrorCallback(errors.Errorf("string to integer conversion error. literal is not a decimal '%s'", ctx.GetText()))
	} else {
		decHold.AcceptDecimal(i64)
		//cons.ConstantValue = reflect.ValueOf(i64)
	}
}

// EnterStringLiteral is called when production stringLiteral is entered.
func (s *GruleParserListener) EnterStringLiteral(ctx *parser.StringLiteralContext) {}

// ExitStringLiteral is called when production stringLiteral is exited.
func (s *GruleParserListener) ExitStringLiteral(ctx *parser.StringLiteralContext) {
	if s.StopParse {
		return
	}
	cons := s.Stack.Peek().(*model.Constant)
	cons.ConstantValue = reflect.ValueOf(strings.Trim(ctx.GetText(), "\"'"))
}

// EnterBooleanLiteral is called when production booleanLiteral is entered.
func (s *GruleParserListener) EnterBooleanLiteral(ctx *parser.BooleanLiteralContext) {
}

// ExitBooleanLiteral is called when production booleanLiteral is exited.
func (s *GruleParserListener) ExitBooleanLiteral(ctx *parser.BooleanLiteralContext) {
	if s.StopParse {
		return
	}
	cons := s.Stack.Peek().(*model.Constant)
	val := strings.ToLower(ctx.GetText())
	switch val {
	case "true":
		cons.ConstantValue = reflect.ValueOf(true)
	case "false":
		cons.ConstantValue = reflect.ValueOf(false)
	default:
		s.ErrorCallback(errors.Errorf("unknown boolear literal '%s'", ctx.GetText()))
	}
}

// EnterRealLiteral is called when production realLiteral is entered.
func (s *GruleParserListener) EnterRealLiteral(ctx *parser.RealLiteralContext) {}

// ExitRealLiteral is called when production realLiteral is exited.
func (s *GruleParserListener) ExitRealLiteral(ctx *parser.RealLiteralContext) {
	if s.StopParse {
		return
	}
	cons := s.Stack.Peek().(*model.Constant)
	flo, err := strconv.ParseFloat(ctx.GetText(), 64)
	if err != nil {
		s.ErrorCallback(errors.Errorf("string to float conversion error. String is not real type '%s'", ctx.GetText()))
		return
	}
	cons.ConstantValue = reflect.ValueOf(flo)
}
