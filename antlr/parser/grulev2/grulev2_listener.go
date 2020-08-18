// Code generated from C:/Users/User/Laboratory/golang/src/github.com/newm4n/grule-rule-engine/antlr\grulev2.g4 by ANTLR 4.8. DO NOT EDIT.

package grulev2 // grulev2
import "github.com/antlr/antlr4/runtime/Go/antlr"

// grulev2Listener is a complete listener for a parse tree produced by grulev2Parser.
type grulev2Listener interface {
	antlr.ParseTreeListener

	// EnterGrl is called when entering the grl production.
	EnterGrl(c *GrlContext)

	// EnterRuleEntry is called when entering the ruleEntry production.
	EnterRuleEntry(c *RuleEntryContext)

	// EnterSalience is called when entering the salience production.
	EnterSalience(c *SalienceContext)

	// EnterRuleName is called when entering the ruleName production.
	EnterRuleName(c *RuleNameContext)

	// EnterRuleDescription is called when entering the ruleDescription production.
	EnterRuleDescription(c *RuleDescriptionContext)

	// EnterWhenScope is called when entering the whenScope production.
	EnterWhenScope(c *WhenScopeContext)

	// EnterThenScope is called when entering the thenScope production.
	EnterThenScope(c *ThenScopeContext)

	// EnterThenExpressionList is called when entering the thenExpressionList production.
	EnterThenExpressionList(c *ThenExpressionListContext)

	// EnterThenExpression is called when entering the thenExpression production.
	EnterThenExpression(c *ThenExpressionContext)

	// EnterAssignment is called when entering the assignment production.
	EnterAssignment(c *AssignmentContext)

	// EnterExpression is called when entering the expression production.
	EnterExpression(c *ExpressionContext)

	// EnterMulDivOperators is called when entering the mulDivOperators production.
	EnterMulDivOperators(c *MulDivOperatorsContext)

	// EnterAddMinusOperators is called when entering the addMinusOperators production.
	EnterAddMinusOperators(c *AddMinusOperatorsContext)

	// EnterComparisonOperator is called when entering the comparisonOperator production.
	EnterComparisonOperator(c *ComparisonOperatorContext)

	// EnterAndLogicOperator is called when entering the andLogicOperator production.
	EnterAndLogicOperator(c *AndLogicOperatorContext)

	// EnterOrLogicOperator is called when entering the orLogicOperator production.
	EnterOrLogicOperator(c *OrLogicOperatorContext)

	// EnterExpressionAtom is called when entering the expressionAtom production.
	EnterExpressionAtom(c *ExpressionAtomContext)

	// EnterArrayMapSelector is called when entering the arrayMapSelector production.
	EnterArrayMapSelector(c *ArrayMapSelectorContext)

	// EnterFunctionCall is called when entering the functionCall production.
	EnterFunctionCall(c *FunctionCallContext)

	// EnterArgumentList is called when entering the argumentList production.
	EnterArgumentList(c *ArgumentListContext)

	// EnterVariable is called when entering the variable production.
	EnterVariable(c *VariableContext)

	// EnterConstant is called when entering the constant production.
	EnterConstant(c *ConstantContext)

	// EnterDecimalLiteral is called when entering the decimalLiteral production.
	EnterDecimalLiteral(c *DecimalLiteralContext)

	// EnterRealLiteral is called when entering the realLiteral production.
	EnterRealLiteral(c *RealLiteralContext)

	// EnterStringLiteral is called when entering the stringLiteral production.
	EnterStringLiteral(c *StringLiteralContext)

	// EnterBooleanLiteral is called when entering the booleanLiteral production.
	EnterBooleanLiteral(c *BooleanLiteralContext)

	// ExitGrl is called when exiting the grl production.
	ExitGrl(c *GrlContext)

	// ExitRuleEntry is called when exiting the ruleEntry production.
	ExitRuleEntry(c *RuleEntryContext)

	// ExitSalience is called when exiting the salience production.
	ExitSalience(c *SalienceContext)

	// ExitRuleName is called when exiting the ruleName production.
	ExitRuleName(c *RuleNameContext)

	// ExitRuleDescription is called when exiting the ruleDescription production.
	ExitRuleDescription(c *RuleDescriptionContext)

	// ExitWhenScope is called when exiting the whenScope production.
	ExitWhenScope(c *WhenScopeContext)

	// ExitThenScope is called when exiting the thenScope production.
	ExitThenScope(c *ThenScopeContext)

	// ExitThenExpressionList is called when exiting the thenExpressionList production.
	ExitThenExpressionList(c *ThenExpressionListContext)

	// ExitThenExpression is called when exiting the thenExpression production.
	ExitThenExpression(c *ThenExpressionContext)

	// ExitAssignment is called when exiting the assignment production.
	ExitAssignment(c *AssignmentContext)

	// ExitExpression is called when exiting the expression production.
	ExitExpression(c *ExpressionContext)

	// ExitMulDivOperators is called when exiting the mulDivOperators production.
	ExitMulDivOperators(c *MulDivOperatorsContext)

	// ExitAddMinusOperators is called when exiting the addMinusOperators production.
	ExitAddMinusOperators(c *AddMinusOperatorsContext)

	// ExitComparisonOperator is called when exiting the comparisonOperator production.
	ExitComparisonOperator(c *ComparisonOperatorContext)

	// ExitAndLogicOperator is called when exiting the andLogicOperator production.
	ExitAndLogicOperator(c *AndLogicOperatorContext)

	// ExitOrLogicOperator is called when exiting the orLogicOperator production.
	ExitOrLogicOperator(c *OrLogicOperatorContext)

	// ExitExpressionAtom is called when exiting the expressionAtom production.
	ExitExpressionAtom(c *ExpressionAtomContext)

	// ExitArrayMapSelector is called when exiting the arrayMapSelector production.
	ExitArrayMapSelector(c *ArrayMapSelectorContext)

	// ExitFunctionCall is called when exiting the functionCall production.
	ExitFunctionCall(c *FunctionCallContext)

	// ExitArgumentList is called when exiting the argumentList production.
	ExitArgumentList(c *ArgumentListContext)

	// ExitVariable is called when exiting the variable production.
	ExitVariable(c *VariableContext)

	// ExitConstant is called when exiting the constant production.
	ExitConstant(c *ConstantContext)

	// ExitDecimalLiteral is called when exiting the decimalLiteral production.
	ExitDecimalLiteral(c *DecimalLiteralContext)

	// ExitRealLiteral is called when exiting the realLiteral production.
	ExitRealLiteral(c *RealLiteralContext)

	// ExitStringLiteral is called when exiting the stringLiteral production.
	ExitStringLiteral(c *StringLiteralContext)

	// ExitBooleanLiteral is called when exiting the booleanLiteral production.
	ExitBooleanLiteral(c *BooleanLiteralContext)
}
