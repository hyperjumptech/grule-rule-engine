// Code generated from C:/Users/User/Laboratory/golang/src/github.com/newm4n/grule-rule-engine/antlr\grulev3.g4 by ANTLR 4.8. DO NOT EDIT.

package grulev3 // grulev3
import "github.com/antlr/antlr4/runtime/Go/antlr"

// grulev3Listener is a complete listener for a parse tree produced by grulev3Parser.
type grulev3Listener interface {
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

	// EnterConstant is called when entering the constant production.
	EnterConstant(c *ConstantContext)

	// EnterVariable is called when entering the variable production.
	EnterVariable(c *VariableContext)

	// EnterArrayMapSelector is called when entering the arrayMapSelector production.
	EnterArrayMapSelector(c *ArrayMapSelectorContext)

	// EnterMemberVariable is called when entering the memberVariable production.
	EnterMemberVariable(c *MemberVariableContext)

	// EnterFunctionCall is called when entering the functionCall production.
	EnterFunctionCall(c *FunctionCallContext)

	// EnterMethodCall is called when entering the methodCall production.
	EnterMethodCall(c *MethodCallContext)

	// EnterArgumentList is called when entering the argumentList production.
	EnterArgumentList(c *ArgumentListContext)

	// EnterFloatLiteral is called when entering the floatLiteral production.
	EnterFloatLiteral(c *FloatLiteralContext)

	// EnterDecimalFloatLiteral is called when entering the decimalFloatLiteral production.
	EnterDecimalFloatLiteral(c *DecimalFloatLiteralContext)

	// EnterHexadecimalFloatLiteral is called when entering the hexadecimalFloatLiteral production.
	EnterHexadecimalFloatLiteral(c *HexadecimalFloatLiteralContext)

	// EnterIntegerLiteral is called when entering the integerLiteral production.
	EnterIntegerLiteral(c *IntegerLiteralContext)

	// EnterDecimalLiteral is called when entering the decimalLiteral production.
	EnterDecimalLiteral(c *DecimalLiteralContext)

	// EnterHexadecimalLiteral is called when entering the hexadecimalLiteral production.
	EnterHexadecimalLiteral(c *HexadecimalLiteralContext)

	// EnterOctalLiteral is called when entering the octalLiteral production.
	EnterOctalLiteral(c *OctalLiteralContext)

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

	// ExitConstant is called when exiting the constant production.
	ExitConstant(c *ConstantContext)

	// ExitVariable is called when exiting the variable production.
	ExitVariable(c *VariableContext)

	// ExitArrayMapSelector is called when exiting the arrayMapSelector production.
	ExitArrayMapSelector(c *ArrayMapSelectorContext)

	// ExitMemberVariable is called when exiting the memberVariable production.
	ExitMemberVariable(c *MemberVariableContext)

	// ExitFunctionCall is called when exiting the functionCall production.
	ExitFunctionCall(c *FunctionCallContext)

	// ExitMethodCall is called when exiting the methodCall production.
	ExitMethodCall(c *MethodCallContext)

	// ExitArgumentList is called when exiting the argumentList production.
	ExitArgumentList(c *ArgumentListContext)

	// ExitFloatLiteral is called when exiting the floatLiteral production.
	ExitFloatLiteral(c *FloatLiteralContext)

	// ExitDecimalFloatLiteral is called when exiting the decimalFloatLiteral production.
	ExitDecimalFloatLiteral(c *DecimalFloatLiteralContext)

	// ExitHexadecimalFloatLiteral is called when exiting the hexadecimalFloatLiteral production.
	ExitHexadecimalFloatLiteral(c *HexadecimalFloatLiteralContext)

	// ExitIntegerLiteral is called when exiting the integerLiteral production.
	ExitIntegerLiteral(c *IntegerLiteralContext)

	// ExitDecimalLiteral is called when exiting the decimalLiteral production.
	ExitDecimalLiteral(c *DecimalLiteralContext)

	// ExitHexadecimalLiteral is called when exiting the hexadecimalLiteral production.
	ExitHexadecimalLiteral(c *HexadecimalLiteralContext)

	// ExitOctalLiteral is called when exiting the octalLiteral production.
	ExitOctalLiteral(c *OctalLiteralContext)

	// ExitStringLiteral is called when exiting the stringLiteral production.
	ExitStringLiteral(c *StringLiteralContext)

	// ExitBooleanLiteral is called when exiting the booleanLiteral production.
	ExitBooleanLiteral(c *BooleanLiteralContext)
}
