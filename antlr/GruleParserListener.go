package antlr

import (
	"fmt"
	"github.com/antlr/antlr4/runtime/Go/antlr"
	"github.com/golang-collections/collections/stack"
	"github.com/hyperjumptech/grule-rule-engine/antlr/parser"
	"github.com/hyperjumptech/grule-rule-engine/model"
	"github.com/juju/errors"
	"github.com/sirupsen/logrus"
	"reflect"
	"strconv"
	"strings"
)

var (
	// Logger is a logrus instance twith default fields for grule
	Logger = logrus.WithFields(logrus.Fields{
		"lib":    "grule",
		"struct": "GruleParserListener",
	})
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
	SerialCounter int
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
	Logger.Errorf("GRL error, after '%v' and then unexpected '%s'", s.PreviousNode, node.GetText())
	s.StopParse = true
	s.ErrorCallback(errors.New(fmt.Sprintf("GRL error, after '%v' and then unexpected '%s'", s.PreviousNode, node.GetText())))
}

// EnterEveryRule is called when any engine is entered.
func (s *GruleParserListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any engine is exited.
func (s *GruleParserListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterRoot is called when production root is entered.
func (s *GruleParserListener) EnterRoot(ctx *parser.RootContext) {
	Logger.Tracef("Entering GRL root")
}

// ExitRoot is called when production root is exited.
func (s *GruleParserListener) ExitRoot(ctx *parser.RootContext) {
	Logger.Tracef("Exiting GRL root")
}

// EnterRuleEntry is called when production ruleEntry is entered.
func (s *GruleParserListener) EnterRuleEntry(ctx *parser.RuleEntryContext) {
	if s.StopParse {
		return
	}
	Logger.Tracef("Entering GRL rule entry")
	entry := &model.RuleEntry{}
	s.Stack.Push(entry)
}

// ExitRuleEntry is called when production ruleEntry is exited.
func (s *GruleParserListener) ExitRuleEntry(ctx *parser.RuleEntryContext) {
	if s.StopParse {
		return
	}
	entry := s.Stack.Pop().(*model.RuleEntry)
	Logger.Tracef("Exiting GRL rule entry '%s'", entry.RuleName)
	// check for duplicate engine.
	if _, ok := s.KnowledgeBase.RuleEntries[entry.RuleName]; ok {
		Logger.Errorf("Duplicate rule entry name '%s'", entry.RuleName)
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
	Logger.Tracef("Entering GRL rule name '%s'", ctx.GetText())
	ruleName := ctx.GetText()
	entry := s.Stack.Peek().(*model.RuleEntry)
	entry.RuleName = ruleName
}

// ExitRuleName is called when production ruleName is exited.
func (s *GruleParserListener) ExitRuleName(ctx *parser.RuleNameContext) {
	Logger.Tracef("Exiting GRL rule name '%s'", ctx.GetText())
}

// EnterSalience is called when production salience is entered.
func (s *GruleParserListener) EnterSalience(ctx *parser.SalienceContext) {
	// salience were set by the decimal literal
	Logger.Tracef("Entering rule salience '%s'", ctx.GetText())
}

// ExitSalience is called when production salience is exited.
func (s *GruleParserListener) ExitSalience(ctx *parser.SalienceContext) {
	Logger.Tracef("Exiting rule salience '%s'", ctx.GetText())
}

// EnterRuleDescription is called when production ruleDescription is entered.
func (s *GruleParserListener) EnterRuleDescription(ctx *parser.RuleDescriptionContext) {
	if s.StopParse {
		return
	}
	Logger.Tracef("Entering rule description '%s'", ctx.GetText())
	ruleDescription := ctx.GetText()
	entry := s.Stack.Peek().(*model.RuleEntry)
	entry.RuleDescription = ruleDescription
}

// ExitRuleDescription is called when production ruleDescription is exited.
func (s *GruleParserListener) ExitRuleDescription(ctx *parser.RuleDescriptionContext) {
	Logger.Tracef("Exiting rule description '%s'", ctx.GetText())
}

// EnterWhenScope is called when production whenScope is entered.
func (s *GruleParserListener) EnterWhenScope(ctx *parser.WhenScopeContext) {
	if s.StopParse {
		return
	}
	Logger.Tracef("Entering when scope")
	whenScope := &model.WhenScope{}
	s.Stack.Push(whenScope)
}

// ExitWhenScope is called when production whenScope is exited.
func (s *GruleParserListener) ExitWhenScope(ctx *parser.WhenScopeContext) {
	if s.StopParse {
		return
	}
	Logger.Tracef("Exiting when scope")
	whenScope := s.Stack.Pop().(*model.WhenScope)
	ruleEntry := s.Stack.Peek().(*model.RuleEntry)
	ruleEntry.WhenScope = whenScope
}

// EnterThenScope is called when production thenScope is entered.
func (s *GruleParserListener) EnterThenScope(ctx *parser.ThenScopeContext) {
	if s.StopParse {
		return
	}
	Logger.Tracef("Entering then scope")
	thenScope := &model.ThenScope{}
	s.Stack.Push(thenScope)
}

// ExitThenScope is called when production thenScope is exited.
func (s *GruleParserListener) ExitThenScope(ctx *parser.ThenScopeContext) {
	if s.StopParse {
		return
	}
	Logger.Tracef("Exiting then scope")
	thenScope := s.Stack.Pop().(*model.ThenScope)
	ruleEntry := s.Stack.Peek().(*model.RuleEntry)
	ruleEntry.ThenScope = thenScope
}

// EnterAssignExpressions is called when production assignExpressions is entered.
func (s *GruleParserListener) EnterAssignExpressions(ctx *parser.AssignExpressionsContext) {
	if s.StopParse {
		return
	}
	Logger.Tracef("Entering assignment expressions")
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
	Logger.Tracef("Exiting assignment expressions")
	assigns := s.Stack.Pop().(*model.AssignExpressions)
	thenScope := s.Stack.Peek().(*model.ThenScope)
	thenScope.AssignExpressions = assigns
}

// EnterAssignExpression is called when production assignExpression is entered.
func (s *GruleParserListener) EnterAssignExpression(ctx *parser.AssignExpressionContext) {
	if s.StopParse {
		return
	}
	Logger.Tracef("Entering assignment expression")
	assign := &model.AssignExpression{}
	s.Stack.Push(assign)
}

// ExitAssignExpression is called when production assignExpression is exited.
func (s *GruleParserListener) ExitAssignExpression(ctx *parser.AssignExpressionContext) {
	if s.StopParse {
		return
	}
	Logger.Tracef("Exiting assignment expression")
	assign := s.Stack.Pop().(*model.AssignExpression)
	assigns := s.Stack.Peek().(*model.AssignExpressions)
	assigns.ExpressionList = append(assigns.ExpressionList, assign)
}

// EnterAssignment is called when production assignment is entered.
func (s *GruleParserListener) EnterAssignment(ctx *parser.AssignmentContext) {
	if s.StopParse {
		return
	}
	Logger.Tracef("Entering assignment")
	assignment := &model.Assignment{
		Text: ctx.GetText(),
	}
	s.Stack.Push(assignment)
}

// ExitAssignment is called when production assignment is exited.
func (s *GruleParserListener) ExitAssignment(ctx *parser.AssignmentContext) {
	if s.StopParse {
		return
	}
	Logger.Tracef("Exiting assignment")
	assignment := s.Stack.Pop().(*model.Assignment)
	assign := s.Stack.Peek().(*model.AssignExpression)
	assign.Assignment = assignment
}

// EnterExpression is called when production expression is entered.
func (s *GruleParserListener) EnterExpression(ctx *parser.ExpressionContext) {
	if s.StopParse {
		return
	}
	Logger.Tracef("Entering expression")
	expression := &model.Expression{}
	s.Stack.Push(expression)
}

// ExitExpression is called when production expression is exited.
func (s *GruleParserListener) ExitExpression(ctx *parser.ExpressionContext) {
	if s.StopParse {
		return
	}
	Logger.Tracef("Exiting expression")
	expr := s.Stack.Pop().(*model.Expression)
	holder := s.Stack.Peek().(model.ExpressionHolder)
	err := holder.AcceptExpression(expr)
	if err != nil {
		Logger.Errorf("error while exiting expression. Got '%v'", err)
		s.ErrorCallback(err)
	}
}

// EnterPredicate is called when production predicate is entered.
func (s *GruleParserListener) EnterPredicate(ctx *parser.PredicateContext) {
	if s.StopParse {
		return
	}
	Logger.Tracef("Entering predicate")
	predicate := &model.Predicate{}
	s.Stack.Push(predicate)
}

// ExitPredicate is called when production predicate is exited.
func (s *GruleParserListener) ExitPredicate(ctx *parser.PredicateContext) {
	if s.StopParse {
		return
	}
	Logger.Tracef("Exiting predicate")
	predicate := s.Stack.Pop().(*model.Predicate)
	expr := s.Stack.Peek().(*model.Expression)
	expr.Predicate = predicate
}

// EnterExpressionAtom is called when production expressionAtom is entered.
func (s *GruleParserListener) EnterExpressionAtom(ctx *parser.ExpressionAtomContext) {
	if s.StopParse {
		return
	}
	Logger.Tracef("Entering expression atom '%s'", ctx.GetText())
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
	Logger.Tracef("Exiting expression atom '%s'", ctx.GetText())
	//fmt.Println(ctx.GetText())
	exprAtom := s.Stack.Pop().(*model.ExpressionAtom)
	exprAtom.SerialNumber = s.SerialCounter
	s.SerialCounter++

	theAtm := s.KnowledgeBase.RuleContext.Add(exprAtom)

	holder := s.Stack.Peek().(model.ExpressionAtomHolder)
	err := holder.AcceptExpressionAtom(theAtm)
	if err != nil {
		Logger.Errorf("error while exiting expression atom. Got '%v'", err)
		s.ErrorCallback(err)
	}
}

// EnterMethodCall is called when production methodCall is entered.
func (s *GruleParserListener) EnterMethodCall(ctx *parser.MethodCallContext) {
	if s.StopParse {
		return
	}
	Logger.Tracef("Entering method call '%s'", ctx.DOTTEDNAME().GetText())
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
	Logger.Tracef("Exiting method call '%s'", ctx.DOTTEDNAME().GetText())
	methodCall := s.Stack.Pop().(*model.MethodCall)
	holder := s.Stack.Peek().(model.MethodCallHolder)
	err := holder.AcceptMethodCall(methodCall)
	if err != nil {
		Logger.Errorf("error while exiting method call. Got '%v'", err)
		s.ErrorCallback(err)
	}
}

// EnterFunctionCall is called when production functionCall is entered.
func (s *GruleParserListener) EnterFunctionCall(ctx *parser.FunctionCallContext) {
	if s.StopParse {
		return
	}
	Logger.Tracef("Entering function call '%s'", ctx.SIMPLENAME().GetText())
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
	Logger.Tracef("Exiting function call '%s'", ctx.SIMPLENAME().GetText())
	funcCall := s.Stack.Pop().(*model.FunctionCall)
	holder := s.Stack.Peek().(model.FunctionCallHolder)
	err := holder.AcceptFunctionCall(funcCall)
	if err != nil {
		Logger.Errorf("error while exiting function call. Got '%v'", err)
		s.ErrorCallback(err)
	}
}

// EnterFunctionArgs is called when production functionArgs is entered.
func (s *GruleParserListener) EnterFunctionArgs(ctx *parser.FunctionArgsContext) {
	if s.StopParse {
		return
	}
	Logger.Tracef("Entering function args")
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
	Logger.Tracef("Exiting function args")
	funcArgs := s.Stack.Pop().(*model.FunctionArgument)
	// return immediately when there's an error
	argHolder := s.Stack.Peek().(model.FunctionArgumentHolder)
	err := argHolder.AcceptFunctionArgument(funcArgs)
	if err != nil {
		Logger.Errorf("error while exiting function args. Got '%v'", err)
		s.ErrorCallback(err)
	}
}

// EnterLogicalOperator is called when production logicalOperator is entered.
func (s *GruleParserListener) EnterLogicalOperator(ctx *parser.LogicalOperatorContext) {
	Logger.Tracef("Entering logical operator '%s'", ctx.GetText())
}

// ExitLogicalOperator is called when production logicalOperator is exited.
func (s *GruleParserListener) ExitLogicalOperator(ctx *parser.LogicalOperatorContext) {
	Logger.Tracef("Exiting logical operator '%s'", ctx.GetText())
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
		Logger.Errorf("unknown logical operator '%s'", ctx.GetText())
		s.ErrorCallback(errors.Errorf("unknown logical operator '%s'", ctx.GetText()))
	}
}

// EnterVariable is called when production variable is entered.
func (s *GruleParserListener) EnterVariable(ctx *parser.VariableContext) {
	Logger.Tracef("Entering variable '%s'", ctx.GetText())
}

// ExitVariable is called when production variable is exited.
func (s *GruleParserListener) ExitVariable(ctx *parser.VariableContext) {
	if s.StopParse {
		return
	}
	Logger.Tracef("Exiting variable '%s'", ctx.GetText())
	varName := ctx.GetText()
	holder := s.Stack.Peek().(model.VariableHolder)
	err := holder.AcceptVariable(varName)
	if err != nil {
		Logger.Errorf("error while exiting variable. Got %v", err)
		s.ErrorCallback(err)
	}
}

// EnterMathOperator is called when production mathOperator is entered.
func (s *GruleParserListener) EnterMathOperator(ctx *parser.MathOperatorContext) {
	Logger.Tracef("Entering math operator '%s'", ctx.GetText())
}

// ExitMathOperator is called when production mathOperator is exited.
func (s *GruleParserListener) ExitMathOperator(ctx *parser.MathOperatorContext) {
	if s.StopParse {
		return
	}
	Logger.Tracef("Exiting math operator '%s'", ctx.GetText())
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
		Logger.Errorf("unknown mathematic operator '%s'", ctx.GetText())
		s.ErrorCallback(errors.Errorf("unknown mathematic operator '%s'", ctx.GetText()))
	}
}

// EnterComparisonOperator is called when production comparisonOperator is entered.
func (s *GruleParserListener) EnterComparisonOperator(ctx *parser.ComparisonOperatorContext) {
	Logger.Tracef("Entering comparison operator '%s'", ctx.GetText())
}

// ExitComparisonOperator is called when production comparisonOperator is exited.
func (s *GruleParserListener) ExitComparisonOperator(ctx *parser.ComparisonOperatorContext) {
	if s.StopParse {
		return
	}
	Logger.Tracef("Exiting comparison operator '%s'", ctx.GetText())
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
		Logger.Errorf("unknown comparison operator '%s'", ctx.GetText())
		s.ErrorCallback(errors.Errorf("unknown comparison operator '%s'", ctx.GetText()))
	}
}

// EnterConstant is called when production constant is entered.
func (s *GruleParserListener) EnterConstant(ctx *parser.ConstantContext) {
	if s.StopParse {
		return
	}
	Logger.Tracef("Entering constant")
	cons := &model.Constant{}
	s.Stack.Push(cons)
}

// ExitConstant is called when production constant is exited.
func (s *GruleParserListener) ExitConstant(ctx *parser.ConstantContext) {
	if s.StopParse {
		return
	}
	Logger.Tracef("Exiting constant")
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
		Logger.Errorf("error while exiting constant. got '%v'", err)
		s.ErrorCallback(err)
	}
}

// EnterDecimalLiteral is called when production decimalLiteral is entered.
func (s *GruleParserListener) EnterDecimalLiteral(ctx *parser.DecimalLiteralContext) {
	Logger.Tracef("Entering decimal literal '%s'", ctx.GetText())
}

// ExitDecimalLiteral is called when production decimalLiteral is exited.
func (s *GruleParserListener) ExitDecimalLiteral(ctx *parser.DecimalLiteralContext) {
	if s.StopParse {
		return
	}
	Logger.Tracef("Exiting Decimal literal '%s'", ctx.GetText())
	decHold := s.Stack.Peek().(model.DecimalHolder)
	i64, err := strconv.ParseInt(ctx.GetText(), 10, 64)
	if err != nil {
		Logger.Errorf("string to integer conversion error. literal is not a decimal '%s'", ctx.GetText())
		s.ErrorCallback(errors.Errorf("string to integer conversion error. literal is not a decimal '%s'", ctx.GetText()))
	} else {
		decHold.AcceptDecimal(i64)
		//cons.ConstantValue = reflect.ValueOf(i64)
	}
}

// EnterStringLiteral is called when production stringLiteral is entered.
func (s *GruleParserListener) EnterStringLiteral(ctx *parser.StringLiteralContext) {
	Logger.Tracef("Entering string literal '%s'", ctx.GetText())
}

// ExitStringLiteral is called when production stringLiteral is exited.
func (s *GruleParserListener) ExitStringLiteral(ctx *parser.StringLiteralContext) {
	if s.StopParse {
		return
	}
	Logger.Tracef("Exiting string literal '%s'", ctx.GetText())
	cons := s.Stack.Peek().(*model.Constant)
	cons.ConstantValue = reflect.ValueOf(strings.Trim(ctx.GetText(), "\"'"))
}

// EnterBooleanLiteral is called when production booleanLiteral is entered.
func (s *GruleParserListener) EnterBooleanLiteral(ctx *parser.BooleanLiteralContext) {
	Logger.Tracef("Entering boolean literal '%s'", ctx.GetText())
}

// ExitBooleanLiteral is called when production booleanLiteral is exited.
func (s *GruleParserListener) ExitBooleanLiteral(ctx *parser.BooleanLiteralContext) {
	if s.StopParse {
		return
	}
	Logger.Tracef("Exiting boolean literal '%s'", ctx.GetText())
	cons := s.Stack.Peek().(*model.Constant)
	val := strings.ToLower(ctx.GetText())
	switch val {
	case "true":
		cons.ConstantValue = reflect.ValueOf(true)
	case "false":
		cons.ConstantValue = reflect.ValueOf(false)
	default:
		Logger.Errorf("unknown boolean literal '%s'", ctx.GetText())
		s.ErrorCallback(errors.Errorf("unknown boolean literal '%s'", ctx.GetText()))
	}
}

// EnterRealLiteral is called when production realLiteral is entered.
func (s *GruleParserListener) EnterRealLiteral(ctx *parser.RealLiteralContext) {
	Logger.Tracef("Entering real literal '%s'", ctx.GetText())
}

// ExitRealLiteral is called when production realLiteral is exited.
func (s *GruleParserListener) ExitRealLiteral(ctx *parser.RealLiteralContext) {
	if s.StopParse {
		return
	}
	Logger.Tracef("Exiting real literal '%s'", ctx.GetText())
	cons := s.Stack.Peek().(*model.Constant)
	flo, err := strconv.ParseFloat(ctx.GetText(), 64)
	if err != nil {
		Logger.Errorf("string to float conversion error. String is not real type '%s'", ctx.GetText())
		s.ErrorCallback(errors.Errorf("string to float conversion error. String is not real type '%s'", ctx.GetText()))
		return
	}
	cons.ConstantValue = reflect.ValueOf(flo)
}
