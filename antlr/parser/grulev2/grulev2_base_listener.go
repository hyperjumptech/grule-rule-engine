// Code generated from C:/Users/User/Laboratory/golang/src/github.com/newm4n/grule-rule-engine/antlr\grulev2.g4 by ANTLR 4.8. DO NOT EDIT.

package grulev2 // grulev2
import "github.com/antlr/antlr4/runtime/Go/antlr"

// Basegrulev2Listener is a complete listener for a parse tree produced by grulev2Parser.
type Basegrulev2Listener struct{}

var _ grulev2Listener = &Basegrulev2Listener{}

// VisitTerminal is called when a terminal node is visited.
func (s *Basegrulev2Listener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *Basegrulev2Listener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *Basegrulev2Listener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *Basegrulev2Listener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterGrl is called when production grl is entered.
func (s *Basegrulev2Listener) EnterGrl(ctx *GrlContext) {}

// ExitGrl is called when production grl is exited.
func (s *Basegrulev2Listener) ExitGrl(ctx *GrlContext) {}

// EnterRuleEntry is called when production ruleEntry is entered.
func (s *Basegrulev2Listener) EnterRuleEntry(ctx *RuleEntryContext) {}

// ExitRuleEntry is called when production ruleEntry is exited.
func (s *Basegrulev2Listener) ExitRuleEntry(ctx *RuleEntryContext) {}

// EnterSalience is called when production salience is entered.
func (s *Basegrulev2Listener) EnterSalience(ctx *SalienceContext) {}

// ExitSalience is called when production salience is exited.
func (s *Basegrulev2Listener) ExitSalience(ctx *SalienceContext) {}

// EnterRuleName is called when production ruleName is entered.
func (s *Basegrulev2Listener) EnterRuleName(ctx *RuleNameContext) {}

// ExitRuleName is called when production ruleName is exited.
func (s *Basegrulev2Listener) ExitRuleName(ctx *RuleNameContext) {}

// EnterRuleDescription is called when production ruleDescription is entered.
func (s *Basegrulev2Listener) EnterRuleDescription(ctx *RuleDescriptionContext) {}

// ExitRuleDescription is called when production ruleDescription is exited.
func (s *Basegrulev2Listener) ExitRuleDescription(ctx *RuleDescriptionContext) {}

// EnterWhenScope is called when production whenScope is entered.
func (s *Basegrulev2Listener) EnterWhenScope(ctx *WhenScopeContext) {}

// ExitWhenScope is called when production whenScope is exited.
func (s *Basegrulev2Listener) ExitWhenScope(ctx *WhenScopeContext) {}

// EnterThenScope is called when production thenScope is entered.
func (s *Basegrulev2Listener) EnterThenScope(ctx *ThenScopeContext) {}

// ExitThenScope is called when production thenScope is exited.
func (s *Basegrulev2Listener) ExitThenScope(ctx *ThenScopeContext) {}

// EnterThenExpressionList is called when production thenExpressionList is entered.
func (s *Basegrulev2Listener) EnterThenExpressionList(ctx *ThenExpressionListContext) {}

// ExitThenExpressionList is called when production thenExpressionList is exited.
func (s *Basegrulev2Listener) ExitThenExpressionList(ctx *ThenExpressionListContext) {}

// EnterThenExpression is called when production thenExpression is entered.
func (s *Basegrulev2Listener) EnterThenExpression(ctx *ThenExpressionContext) {}

// ExitThenExpression is called when production thenExpression is exited.
func (s *Basegrulev2Listener) ExitThenExpression(ctx *ThenExpressionContext) {}

// EnterAssignment is called when production assignment is entered.
func (s *Basegrulev2Listener) EnterAssignment(ctx *AssignmentContext) {}

// ExitAssignment is called when production assignment is exited.
func (s *Basegrulev2Listener) ExitAssignment(ctx *AssignmentContext) {}

// EnterExpression is called when production expression is entered.
func (s *Basegrulev2Listener) EnterExpression(ctx *ExpressionContext) {}

// ExitExpression is called when production expression is exited.
func (s *Basegrulev2Listener) ExitExpression(ctx *ExpressionContext) {}

// EnterMulDivOperators is called when production mulDivOperators is entered.
func (s *Basegrulev2Listener) EnterMulDivOperators(ctx *MulDivOperatorsContext) {}

// ExitMulDivOperators is called when production mulDivOperators is exited.
func (s *Basegrulev2Listener) ExitMulDivOperators(ctx *MulDivOperatorsContext) {}

// EnterAddMinusOperators is called when production addMinusOperators is entered.
func (s *Basegrulev2Listener) EnterAddMinusOperators(ctx *AddMinusOperatorsContext) {}

// ExitAddMinusOperators is called when production addMinusOperators is exited.
func (s *Basegrulev2Listener) ExitAddMinusOperators(ctx *AddMinusOperatorsContext) {}

// EnterComparisonOperator is called when production comparisonOperator is entered.
func (s *Basegrulev2Listener) EnterComparisonOperator(ctx *ComparisonOperatorContext) {}

// ExitComparisonOperator is called when production comparisonOperator is exited.
func (s *Basegrulev2Listener) ExitComparisonOperator(ctx *ComparisonOperatorContext) {}

// EnterAndLogicOperator is called when production andLogicOperator is entered.
func (s *Basegrulev2Listener) EnterAndLogicOperator(ctx *AndLogicOperatorContext) {}

// ExitAndLogicOperator is called when production andLogicOperator is exited.
func (s *Basegrulev2Listener) ExitAndLogicOperator(ctx *AndLogicOperatorContext) {}

// EnterOrLogicOperator is called when production orLogicOperator is entered.
func (s *Basegrulev2Listener) EnterOrLogicOperator(ctx *OrLogicOperatorContext) {}

// ExitOrLogicOperator is called when production orLogicOperator is exited.
func (s *Basegrulev2Listener) ExitOrLogicOperator(ctx *OrLogicOperatorContext) {}

// EnterExpressionAtom is called when production expressionAtom is entered.
func (s *Basegrulev2Listener) EnterExpressionAtom(ctx *ExpressionAtomContext) {}

// ExitExpressionAtom is called when production expressionAtom is exited.
func (s *Basegrulev2Listener) ExitExpressionAtom(ctx *ExpressionAtomContext) {}

// EnterArrayMapSelector is called when production arrayMapSelector is entered.
func (s *Basegrulev2Listener) EnterArrayMapSelector(ctx *ArrayMapSelectorContext) {}

// ExitArrayMapSelector is called when production arrayMapSelector is exited.
func (s *Basegrulev2Listener) ExitArrayMapSelector(ctx *ArrayMapSelectorContext) {}

// EnterFunctionCall is called when production functionCall is entered.
func (s *Basegrulev2Listener) EnterFunctionCall(ctx *FunctionCallContext) {}

// ExitFunctionCall is called when production functionCall is exited.
func (s *Basegrulev2Listener) ExitFunctionCall(ctx *FunctionCallContext) {}

// EnterArgumentList is called when production argumentList is entered.
func (s *Basegrulev2Listener) EnterArgumentList(ctx *ArgumentListContext) {}

// ExitArgumentList is called when production argumentList is exited.
func (s *Basegrulev2Listener) ExitArgumentList(ctx *ArgumentListContext) {}

// EnterVariable is called when production variable is entered.
func (s *Basegrulev2Listener) EnterVariable(ctx *VariableContext) {}

// ExitVariable is called when production variable is exited.
func (s *Basegrulev2Listener) ExitVariable(ctx *VariableContext) {}

// EnterConstant is called when production constant is entered.
func (s *Basegrulev2Listener) EnterConstant(ctx *ConstantContext) {}

// ExitConstant is called when production constant is exited.
func (s *Basegrulev2Listener) ExitConstant(ctx *ConstantContext) {}

// EnterDecimalLiteral is called when production decimalLiteral is entered.
func (s *Basegrulev2Listener) EnterDecimalLiteral(ctx *DecimalLiteralContext) {}

// ExitDecimalLiteral is called when production decimalLiteral is exited.
func (s *Basegrulev2Listener) ExitDecimalLiteral(ctx *DecimalLiteralContext) {}

// EnterRealLiteral is called when production realLiteral is entered.
func (s *Basegrulev2Listener) EnterRealLiteral(ctx *RealLiteralContext) {}

// ExitRealLiteral is called when production realLiteral is exited.
func (s *Basegrulev2Listener) ExitRealLiteral(ctx *RealLiteralContext) {}

// EnterStringLiteral is called when production stringLiteral is entered.
func (s *Basegrulev2Listener) EnterStringLiteral(ctx *StringLiteralContext) {}

// ExitStringLiteral is called when production stringLiteral is exited.
func (s *Basegrulev2Listener) ExitStringLiteral(ctx *StringLiteralContext) {}

// EnterBooleanLiteral is called when production booleanLiteral is entered.
func (s *Basegrulev2Listener) EnterBooleanLiteral(ctx *BooleanLiteralContext) {}

// ExitBooleanLiteral is called when production booleanLiteral is exited.
func (s *Basegrulev2Listener) ExitBooleanLiteral(ctx *BooleanLiteralContext) {}
