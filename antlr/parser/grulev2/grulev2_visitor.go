// Code generated from C:/Users/User/Laboratory/golang/src/github.com/newm4n/grule-rule-engine/antlr\grulev2.g4 by ANTLR 4.8. DO NOT EDIT.

package grulev2 // grulev2
import "github.com/antlr/antlr4/runtime/Go/antlr"

// A complete Visitor for a parse tree produced by grulev2Parser.
type grulev2Visitor interface {
	antlr.ParseTreeVisitor

	// Visit a parse tree produced by grulev2Parser#grl.
	VisitGrl(ctx *GrlContext) interface{}

	// Visit a parse tree produced by grulev2Parser#ruleEntry.
	VisitRuleEntry(ctx *RuleEntryContext) interface{}

	// Visit a parse tree produced by grulev2Parser#salience.
	VisitSalience(ctx *SalienceContext) interface{}

	// Visit a parse tree produced by grulev2Parser#ruleName.
	VisitRuleName(ctx *RuleNameContext) interface{}

	// Visit a parse tree produced by grulev2Parser#ruleDescription.
	VisitRuleDescription(ctx *RuleDescriptionContext) interface{}

	// Visit a parse tree produced by grulev2Parser#whenScope.
	VisitWhenScope(ctx *WhenScopeContext) interface{}

	// Visit a parse tree produced by grulev2Parser#thenScope.
	VisitThenScope(ctx *ThenScopeContext) interface{}

	// Visit a parse tree produced by grulev2Parser#thenExpressionList.
	VisitThenExpressionList(ctx *ThenExpressionListContext) interface{}

	// Visit a parse tree produced by grulev2Parser#thenExpression.
	VisitThenExpression(ctx *ThenExpressionContext) interface{}

	// Visit a parse tree produced by grulev2Parser#assignment.
	VisitAssignment(ctx *AssignmentContext) interface{}

	// Visit a parse tree produced by grulev2Parser#expression.
	VisitExpression(ctx *ExpressionContext) interface{}

	// Visit a parse tree produced by grulev2Parser#mulDivOperators.
	VisitMulDivOperators(ctx *MulDivOperatorsContext) interface{}

	// Visit a parse tree produced by grulev2Parser#addMinusOperators.
	VisitAddMinusOperators(ctx *AddMinusOperatorsContext) interface{}

	// Visit a parse tree produced by grulev2Parser#comparisonOperator.
	VisitComparisonOperator(ctx *ComparisonOperatorContext) interface{}

	// Visit a parse tree produced by grulev2Parser#andLogicOperator.
	VisitAndLogicOperator(ctx *AndLogicOperatorContext) interface{}

	// Visit a parse tree produced by grulev2Parser#orLogicOperator.
	VisitOrLogicOperator(ctx *OrLogicOperatorContext) interface{}

	// Visit a parse tree produced by grulev2Parser#expressionAtom.
	VisitExpressionAtom(ctx *ExpressionAtomContext) interface{}

	// Visit a parse tree produced by grulev2Parser#arrayMapSelector.
	VisitArrayMapSelector(ctx *ArrayMapSelectorContext) interface{}

	// Visit a parse tree produced by grulev2Parser#functionCall.
	VisitFunctionCall(ctx *FunctionCallContext) interface{}

	// Visit a parse tree produced by grulev2Parser#argumentList.
	VisitArgumentList(ctx *ArgumentListContext) interface{}

	// Visit a parse tree produced by grulev2Parser#variable.
	VisitVariable(ctx *VariableContext) interface{}

	// Visit a parse tree produced by grulev2Parser#constant.
	VisitConstant(ctx *ConstantContext) interface{}

	// Visit a parse tree produced by grulev2Parser#decimalLiteral.
	VisitDecimalLiteral(ctx *DecimalLiteralContext) interface{}

	// Visit a parse tree produced by grulev2Parser#realLiteral.
	VisitRealLiteral(ctx *RealLiteralContext) interface{}

	// Visit a parse tree produced by grulev2Parser#stringLiteral.
	VisitStringLiteral(ctx *StringLiteralContext) interface{}

	// Visit a parse tree produced by grulev2Parser#booleanLiteral.
	VisitBooleanLiteral(ctx *BooleanLiteralContext) interface{}
}
