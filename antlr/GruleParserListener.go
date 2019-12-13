package antlr

import (
	"fmt"
	"github.com/antlr/antlr4/runtime/Go/antlr"
	"github.com/golang-collections/collections/stack"
	"github.com/hyperjumptech/grule-rule-engine/antlr/parser"
	"github.com/hyperjumptech/grule-rule-engine/model"
	"github.com/juju/errors"
	log "github.com/sirupsen/logrus"
	"reflect"
	"strconv"
	"strings"
)

// NewGruleParserListener create a new instancce of GruleParserListener.
// This listener will walk in the Grule GRL file and invoke operations based on the
// context within the knowledge base.
func NewGruleParserListener(kbase *model.KnowledgeBase) *GruleParserListener {
	return &GruleParserListener{
		Stack:         stack.New(),
		KnowledgeBase: kbase,
		ParseErrors:   make([]error, 0),
	}
}

// GruleParserListener is an implementation of logic to build the execution flow or execution graph as it
// defined within the knowledge base.
type GruleParserListener struct {
	parser.BasegruleListener
	ParseErrors []error

	//RuleEntries map[string]*model.RuleEntry
	KnowledgeBase *model.KnowledgeBase
	Stack         *stack.Stack
}

// AddError will add an error while parsing and building knowledge base.
func (s *GruleParserListener) AddError(e error) {
	log.Errorf("Got error : %v", e)
	s.ParseErrors = append(s.ParseErrors, e)
}

// VisitTerminal is called when a terminal node is visited.
func (s *GruleParserListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *GruleParserListener) VisitErrorNode(node antlr.ErrorNode) {}

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
	// return immediately when there's an error
	if len(s.ParseErrors) > 0 {
		return
	}
	entry := &model.RuleEntry{}
	s.Stack.Push(entry)
}

// ExitRuleEntry is called when production ruleEntry is exited.
func (s *GruleParserListener) ExitRuleEntry(ctx *parser.RuleEntryContext) {
	entry := s.Stack.Pop().(*model.RuleEntry)
	// return immediately when there's an error
	if len(s.ParseErrors) > 0 {
		return
	}
	// check for duplicate engine.
	if _, ok := s.KnowledgeBase.RuleEntries[entry.RuleName]; ok {
		s.AddError(errors.Errorf("duplicate rule entry name '%s'", entry.RuleName))
		return
	}
	// if everything ok, add the engine entry.
	s.KnowledgeBase.RuleEntries[entry.RuleName] = entry
}

// EnterRuleName is called when production ruleName is entered.
func (s *GruleParserListener) EnterRuleName(ctx *parser.RuleNameContext) {
	// return immediately when there's an error
	if len(s.ParseErrors) > 0 {
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
	// return immediately when there's an error
	if len(s.ParseErrors) > 0 {
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
	// return immediately when there's an error
	if len(s.ParseErrors) > 0 {
		return
	}
	whenScope := &model.WhenScope{}
	s.Stack.Push(whenScope)
}

// ExitWhenScope is called when production whenScope is exited.
func (s *GruleParserListener) ExitWhenScope(ctx *parser.WhenScopeContext) {
	whenScope := s.Stack.Pop().(*model.WhenScope)
	// return immediately when there's an error
	if len(s.ParseErrors) > 0 {
		return
	}
	ruleEntry := s.Stack.Peek().(*model.RuleEntry)
	ruleEntry.WhenScope = whenScope
}

// EnterThenScope is called when production thenScope is entered.
func (s *GruleParserListener) EnterThenScope(ctx *parser.ThenScopeContext) {
	// return immediately when there's an error
	if len(s.ParseErrors) > 0 {
		return
	}
	thenScope := &model.ThenScope{}
	s.Stack.Push(thenScope)
}

// ExitThenScope is called when production thenScope is exited.
func (s *GruleParserListener) ExitThenScope(ctx *parser.ThenScopeContext) {
	thenScope := s.Stack.Pop().(*model.ThenScope)
	// return immediately when there's an error
	if len(s.ParseErrors) > 0 {
		return
	}
	ruleEntry := s.Stack.Peek().(*model.RuleEntry)
	ruleEntry.ThenScope = thenScope
}

// EnterAssignExpressions is called when production assignExpressions is entered.
func (s *GruleParserListener) EnterAssignExpressions(ctx *parser.AssignExpressionsContext) {
	// return immediately when there's an error
	if len(s.ParseErrors) > 0 {
		return
	}
	assigns := &model.AssignExpressions{
		ExpressionList: make([]*model.AssignExpression, 0),
	}
	s.Stack.Push(assigns)
}

// ExitAssignExpressions is called when production assignExpressions is exited.
func (s *GruleParserListener) ExitAssignExpressions(ctx *parser.AssignExpressionsContext) {
	assigns := s.Stack.Pop().(*model.AssignExpressions)
	// return immediately when there's an error
	if len(s.ParseErrors) > 0 {
		return
	}
	thenScope := s.Stack.Peek().(*model.ThenScope)
	thenScope.AssignExpressions = assigns
}

// EnterAssignExpression is called when production assignExpression is entered.
func (s *GruleParserListener) EnterAssignExpression(ctx *parser.AssignExpressionContext) {
	// return immediately when there's an error
	if len(s.ParseErrors) > 0 {
		return
	}
	assign := &model.AssignExpression{}
	s.Stack.Push(assign)
}

// ExitAssignExpression is called when production assignExpression is exited.
func (s *GruleParserListener) ExitAssignExpression(ctx *parser.AssignExpressionContext) {
	assign := s.Stack.Pop().(*model.AssignExpression)
	// return immediately when there's an error
	if len(s.ParseErrors) > 0 {
		return
	}
	assigns := s.Stack.Peek().(*model.AssignExpressions)
	assigns.ExpressionList = append(assigns.ExpressionList, assign)
}

// EnterAssignment is called when production assignment is entered.
func (s *GruleParserListener) EnterAssignment(ctx *parser.AssignmentContext) {
	// return immediately when there's an error
	if len(s.ParseErrors) > 0 {
		return
	}
	assignment := &model.Assignment{}
	s.Stack.Push(assignment)
}

// ExitAssignment is called when production assignment is exited.
func (s *GruleParserListener) ExitAssignment(ctx *parser.AssignmentContext) {
	assignment := s.Stack.Pop().(*model.Assignment)
	// return immediately when there's an error
	if len(s.ParseErrors) > 0 {
		return
	}
	assign := s.Stack.Peek().(*model.AssignExpression)
	assign.Assignment = assignment
}

// EnterExpression is called when production expression is entered.
func (s *GruleParserListener) EnterExpression(ctx *parser.ExpressionContext) {
	// return immediately when there's an error
	if len(s.ParseErrors) > 0 {
		return
	}
	expression := &model.Expression{}
	s.Stack.Push(expression)
}

// ExitExpression is called when production expression is exited.
func (s *GruleParserListener) ExitExpression(ctx *parser.ExpressionContext) {
	expr := s.Stack.Pop().(*model.Expression)
	// return immediately when there's an error
	if len(s.ParseErrors) > 0 {
		return
	}
	holder := s.Stack.Peek().(model.ExpressionHolder)
	err := holder.AcceptExpression(expr)
	if err != nil {
		s.AddError(err)
	}
}

// EnterPredicate is called when production predicate is entered.
func (s *GruleParserListener) EnterPredicate(ctx *parser.PredicateContext) {
	// return immediately when there's an error
	if len(s.ParseErrors) > 0 {
		return
	}
	predicate := &model.Predicate{}
	s.Stack.Push(predicate)
}

// ExitPredicate is called when production predicate is exited.
func (s *GruleParserListener) ExitPredicate(ctx *parser.PredicateContext) {
	predicate := s.Stack.Pop().(*model.Predicate)
	// return immediately when there's an error
	if len(s.ParseErrors) > 0 {
		return
	}
	expr := s.Stack.Peek().(*model.Expression)
	expr.Predicate = predicate
}

// EnterExpressionAtom is called when production expressionAtom is entered.
func (s *GruleParserListener) EnterExpressionAtom(ctx *parser.ExpressionAtomContext) {
	// return immediately when there's an error
	if len(s.ParseErrors) > 0 {
		return
	}
	exprAtom := &model.ExpressionAtom{
		Text: ctx.GetText(),
	}
	s.Stack.Push(exprAtom)
}

// ExitExpressionAtom is called when production expressionAtom is exited.
func (s *GruleParserListener) ExitExpressionAtom(ctx *parser.ExpressionAtomContext) {
	//fmt.Println(ctx.GetText())
	exprAtom := s.Stack.Pop().(*model.ExpressionAtom)
	// return immediately when there's an error
	if len(s.ParseErrors) > 0 {
		return
	}
	holder := s.Stack.Peek().(model.ExpressionAtomHolder)
	err := holder.AcceptExpressionAtom(exprAtom)
	if err != nil {
		s.AddError(err)
	}
}

// EnterMethodCall is called when production methodCall is entered.
func (s *GruleParserListener) EnterMethodCall(ctx *parser.MethodCallContext) {
	// return immediately when there's an error
	if len(s.ParseErrors) > 0 {
		return
	}
	funcCall := &model.MethodCall{
		MethodName: ctx.DOTTEDNAME().GetText(),
	}
	s.Stack.Push(funcCall)
}

// ExitMethodCall is called when production methodCall is exited.
func (s *GruleParserListener) ExitMethodCall(ctx *parser.MethodCallContext) {
	methodCall := s.Stack.Pop().(*model.MethodCall)
	// return immediately when there's an error
	if len(s.ParseErrors) > 0 {
		return
	}
	holder := s.Stack.Peek().(model.MethodCallHolder)
	err := holder.AcceptMethodCall(methodCall)
	if err != nil {
		fmt.Printf("Got error %s\n", err)
		s.AddError(err)
	}
}

// EnterFunctionCall is called when production functionCall is entered.
func (s *GruleParserListener) EnterFunctionCall(ctx *parser.FunctionCallContext) {
	// return immediately when there's an error
	if len(s.ParseErrors) > 0 {
		return
	}
	funcCall := &model.FunctionCall{
		FunctionName: ctx.SIMPLENAME().GetText(),
	}
	s.Stack.Push(funcCall)
}

// ExitFunctionCall is called when production functionCall is exited.
func (s *GruleParserListener) ExitFunctionCall(ctx *parser.FunctionCallContext) {
	funcCall := s.Stack.Pop().(*model.FunctionCall)
	// return immediately when there's an error
	if len(s.ParseErrors) > 0 {
		return
	}
	holder := s.Stack.Peek().(model.FunctionCallHolder)
	err := holder.AcceptFunctionCall(funcCall)
	if err != nil {
		s.AddError(err)
	}
}

// EnterFunctionArgs is called when production functionArgs is entered.
func (s *GruleParserListener) EnterFunctionArgs(ctx *parser.FunctionArgsContext) {
	// return immediately when there's an error
	if len(s.ParseErrors) > 0 {
		return
	}
	funcArg := &model.FunctionArgument{
		Arguments: make([]*model.ArgumentHolder, 0),
	}
	s.Stack.Push(funcArg)
}

// ExitFunctionArgs is called when production functionArgs is exited.
func (s *GruleParserListener) ExitFunctionArgs(ctx *parser.FunctionArgsContext) {
	funcArgs := s.Stack.Pop().(*model.FunctionArgument)
	// return immediately when there's an error
	if len(s.ParseErrors) > 0 {
		return
	}
	argHolder := s.Stack.Peek().(model.FunctionArgumentHolder)
	err := argHolder.AcceptFunctionArgument(funcArgs)
	if err != nil {
		s.AddError(err)
	}
}

// EnterLogicalOperator is called when production logicalOperator is entered.
func (s *GruleParserListener) EnterLogicalOperator(ctx *parser.LogicalOperatorContext) {
}

// ExitLogicalOperator is called when production logicalOperator is exited.
func (s *GruleParserListener) ExitLogicalOperator(ctx *parser.LogicalOperatorContext) {
	// return immediately when there's an error
	if len(s.ParseErrors) > 0 {
		return
	}
	expr := s.Stack.Peek().(*model.Expression)
	if ctx.GetText() == "&&" {
		expr.LogicalOperator = model.LogicalOperatorAnd
	} else if ctx.GetText() == "||" {
		expr.LogicalOperator = model.LogicalOperatorOr
	} else {
		s.AddError(errors.Errorf("unknown logical operator %s", ctx.GetText()))
	}
}

// EnterVariable is called when production variable is entered.
func (s *GruleParserListener) EnterVariable(ctx *parser.VariableContext) {}

// ExitVariable is called when production variable is exited.
func (s *GruleParserListener) ExitVariable(ctx *parser.VariableContext) {
	// return immediately when there's an error
	if len(s.ParseErrors) > 0 {
		return
	}
	varName := ctx.GetText()
	//fmt.Println("Variable Name", varName)
	// return immediately when there's an error
	if len(s.ParseErrors) > 0 {
		return
	}
	holder := s.Stack.Peek().(model.VariableHolder)
	err := holder.AcceptVariable(varName)
	if err != nil {
		s.AddError(err)
	}
}

// EnterMathOperator is called when production mathOperator is entered.
func (s *GruleParserListener) EnterMathOperator(ctx *parser.MathOperatorContext) {
}

// ExitMathOperator is called when production mathOperator is exited.
func (s *GruleParserListener) ExitMathOperator(ctx *parser.MathOperatorContext) {
	// return immediately when there's an error
	if len(s.ParseErrors) > 0 {
		return
	}
	expr := s.Stack.Peek().(*model.ExpressionAtom)
	if ctx.GetText() == "+" {
		expr.MathOperator = model.MathOperatorPlus
	} else if ctx.GetText() == "-" {
		expr.MathOperator = model.MathOperatorMinus
	} else if ctx.GetText() == "/" {
		expr.MathOperator = model.MathOperatorDiv
	} else if ctx.GetText() == "*" {
		expr.MathOperator = model.MathOperatorMul
	} else {
		s.AddError(errors.Errorf("unknown mathematic operator %s", ctx.GetText()))
	}
}

// EnterComparisonOperator is called when production comparisonOperator is entered.
func (s *GruleParserListener) EnterComparisonOperator(ctx *parser.ComparisonOperatorContext) {}

// ExitComparisonOperator is called when production comparisonOperator is exited.
func (s *GruleParserListener) ExitComparisonOperator(ctx *parser.ComparisonOperatorContext) {
	// return immediately when there's an error
	if len(s.ParseErrors) > 0 {
		return
	}
	predicate := s.Stack.Peek().(*model.Predicate)
	if ctx.GetText() == "==" {
		predicate.ComparisonOperator = model.ComparisonOperatorEQ
	} else if ctx.GetText() == "!=" {
		predicate.ComparisonOperator = model.ComparisonOperatorNEQ
	} else if ctx.GetText() == "<" {
		predicate.ComparisonOperator = model.ComparisonOperatorLT
	} else if ctx.GetText() == "<=" {
		predicate.ComparisonOperator = model.ComparisonOperatorLTE
	} else if ctx.GetText() == ">" {
		predicate.ComparisonOperator = model.ComparisonOperatorGT
	} else if ctx.GetText() == ">=" {
		predicate.ComparisonOperator = model.ComparisonOperatorGTE
	} else {
		s.AddError(errors.Errorf("unknown comparison operator %s", ctx.GetText()))
	}
}

// EnterConstant is called when production constant is entered.
func (s *GruleParserListener) EnterConstant(ctx *parser.ConstantContext) {
	// return immediately when there's an error
	if len(s.ParseErrors) > 0 {
		return
	}
	cons := &model.Constant{}
	s.Stack.Push(cons)
}

// ExitConstant is called when production constant is exited.
func (s *GruleParserListener) ExitConstant(ctx *parser.ConstantContext) {
	cons := s.Stack.Pop().(*model.Constant)
	// return immediately when there's an error
	if len(s.ParseErrors) > 0 {
		return
	}
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
		s.AddError(err)
	}
}

// EnterDecimalLiteral is called when production decimalLiteral is entered.
func (s *GruleParserListener) EnterDecimalLiteral(ctx *parser.DecimalLiteralContext) {}

// ExitDecimalLiteral is called when production decimalLiteral is exited.
func (s *GruleParserListener) ExitDecimalLiteral(ctx *parser.DecimalLiteralContext) {
	// return immediately when there's an error
	if len(s.ParseErrors) > 0 {
		return
	}
	decHold := s.Stack.Peek().(model.DecimalHolder)
	i64, err := strconv.ParseInt(ctx.GetText(), 10, 64)
	if err != nil {
		s.AddError(errors.Errorf("string to integer conversion error. literal is not a decimal '%s'", ctx.GetText()))
	} else {
		decHold.AcceptDecimal(i64)
		//cons.ConstantValue = reflect.ValueOf(i64)
	}
}

// EnterStringLiteral is called when production stringLiteral is entered.
func (s *GruleParserListener) EnterStringLiteral(ctx *parser.StringLiteralContext) {}

// ExitStringLiteral is called when production stringLiteral is exited.
func (s *GruleParserListener) ExitStringLiteral(ctx *parser.StringLiteralContext) {
	// return immediately when there's an error
	if len(s.ParseErrors) > 0 {
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
	// return immediately when there's an error
	if len(s.ParseErrors) > 0 {
		return
	}
	cons := s.Stack.Peek().(*model.Constant)
	val := strings.ToLower(ctx.GetText())
	if val == "true" {
		cons.ConstantValue = reflect.ValueOf(true)
	} else if val == "false" {
		cons.ConstantValue = reflect.ValueOf(false)
	} else {
		s.AddError(errors.Errorf("unknown boolear literal '%s'", ctx.GetText()))
	}
}

// EnterRealLiteral is called when production realLiteral is entered.
func (s *GruleParserListener) EnterRealLiteral(ctx *parser.RealLiteralContext) {}

// ExitRealLiteral is called when production realLiteral is exited.
func (s *GruleParserListener) ExitRealLiteral(ctx *parser.RealLiteralContext) {
	// return immediately when there's an error
	if len(s.ParseErrors) > 0 {
		return
	}
	cons := s.Stack.Peek().(*model.Constant)
	flo, err := strconv.ParseFloat(ctx.GetText(), 64)
	if err != nil {
		s.AddError(errors.Errorf("string to float conversion error. String is not real type '%s'", ctx.GetText()))
		return
	}
	cons.ConstantValue = reflect.ValueOf(flo)
}
