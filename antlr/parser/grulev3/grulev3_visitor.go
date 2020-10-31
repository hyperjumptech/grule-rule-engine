// Code generated from C:/Users/User/Laboratory/golang/src/github.com/newm4n/grule-rule-engine/antlr\grulev3.g4 by ANTLR 4.8. DO NOT EDIT.

package grulev3 // grulev3
import "github.com/antlr/antlr4/runtime/Go/antlr"

// A complete Visitor for a parse tree produced by grulev3Parser.
type grulev3Visitor interface {
	antlr.ParseTreeVisitor

	// Visit a parse tree produced by grulev3Parser#grl.
	VisitGrl(ctx *GrlContext) interface{}

	// Visit a parse tree produced by grulev3Parser#ruleEntry.
	VisitRuleEntry(ctx *RuleEntryContext) interface{}

	// Visit a parse tree produced by grulev3Parser#salience.
	VisitSalience(ctx *SalienceContext) interface{}

	// Visit a parse tree produced by grulev3Parser#ruleName.
	VisitRuleName(ctx *RuleNameContext) interface{}

	// Visit a parse tree produced by grulev3Parser#ruleDescription.
	VisitRuleDescription(ctx *RuleDescriptionContext) interface{}

	// Visit a parse tree produced by grulev3Parser#whenScope.
	VisitWhenScope(ctx *WhenScopeContext) interface{}

	// Visit a parse tree produced by grulev3Parser#thenScope.
	VisitThenScope(ctx *ThenScopeContext) interface{}

	// Visit a parse tree produced by grulev3Parser#thenExpressionList.
	VisitThenExpressionList(ctx *ThenExpressionListContext) interface{}

	// Visit a parse tree produced by grulev3Parser#thenExpression.
	VisitThenExpression(ctx *ThenExpressionContext) interface{}

	// Visit a parse tree produced by grulev3Parser#assignment.
	VisitAssignment(ctx *AssignmentContext) interface{}

	// Visit a parse tree produced by grulev3Parser#expression.
	VisitExpression(ctx *ExpressionContext) interface{}

	// Visit a parse tree produced by grulev3Parser#mulDivOperators.
	VisitMulDivOperators(ctx *MulDivOperatorsContext) interface{}

	// Visit a parse tree produced by grulev3Parser#addMinusOperators.
	VisitAddMinusOperators(ctx *AddMinusOperatorsContext) interface{}

	// Visit a parse tree produced by grulev3Parser#comparisonOperator.
	VisitComparisonOperator(ctx *ComparisonOperatorContext) interface{}

	// Visit a parse tree produced by grulev3Parser#andLogicOperator.
	VisitAndLogicOperator(ctx *AndLogicOperatorContext) interface{}

	// Visit a parse tree produced by grulev3Parser#orLogicOperator.
	VisitOrLogicOperator(ctx *OrLogicOperatorContext) interface{}

	// Visit a parse tree produced by grulev3Parser#expressionAtom.
	VisitExpressionAtom(ctx *ExpressionAtomContext) interface{}

	// Visit a parse tree produced by grulev3Parser#constant.
	VisitConstant(ctx *ConstantContext) interface{}

	// Visit a parse tree produced by grulev3Parser#variable.
	VisitVariable(ctx *VariableContext) interface{}

	// Visit a parse tree produced by grulev3Parser#arrayMapSelector.
	VisitArrayMapSelector(ctx *ArrayMapSelectorContext) interface{}

	// Visit a parse tree produced by grulev3Parser#memberVariable.
	VisitMemberVariable(ctx *MemberVariableContext) interface{}

	// Visit a parse tree produced by grulev3Parser#functionCall.
	VisitFunctionCall(ctx *FunctionCallContext) interface{}

	// Visit a parse tree produced by grulev3Parser#methodCall.
	VisitMethodCall(ctx *MethodCallContext) interface{}

	// Visit a parse tree produced by grulev3Parser#argumentList.
	VisitArgumentList(ctx *ArgumentListContext) interface{}

	// Visit a parse tree produced by grulev3Parser#floatLiteral.
	VisitFloatLiteral(ctx *FloatLiteralContext) interface{}

	// Visit a parse tree produced by grulev3Parser#decimalFloatLiteral.
	VisitDecimalFloatLiteral(ctx *DecimalFloatLiteralContext) interface{}

	// Visit a parse tree produced by grulev3Parser#hexadecimalFloatLiteral.
	VisitHexadecimalFloatLiteral(ctx *HexadecimalFloatLiteralContext) interface{}

	// Visit a parse tree produced by grulev3Parser#integerLiteral.
	VisitIntegerLiteral(ctx *IntegerLiteralContext) interface{}

	// Visit a parse tree produced by grulev3Parser#decimalLiteral.
	VisitDecimalLiteral(ctx *DecimalLiteralContext) interface{}

	// Visit a parse tree produced by grulev3Parser#hexadecimalLiteral.
	VisitHexadecimalLiteral(ctx *HexadecimalLiteralContext) interface{}

	// Visit a parse tree produced by grulev3Parser#octalLiteral.
	VisitOctalLiteral(ctx *OctalLiteralContext) interface{}

	// Visit a parse tree produced by grulev3Parser#stringLiteral.
	VisitStringLiteral(ctx *StringLiteralContext) interface{}

	// Visit a parse tree produced by grulev3Parser#booleanLiteral.
	VisitBooleanLiteral(ctx *BooleanLiteralContext) interface{}
}
