// Code generated from C:/Users/User/Laboratory/golang/src/github.com/newm4n/grule-rule-engine/antlr\grulev3.g4 by ANTLR 4.8. DO NOT EDIT.

package grulev3 // grulev3
import "github.com/antlr/antlr4/runtime/Go/antlr"

// Basegrulev3Listener is a complete listener for a parse tree produced by grulev3Parser.
type Basegrulev3Listener struct{}

var _ grulev3Listener = &Basegrulev3Listener{}

// VisitTerminal is called when a terminal node is visited.
func (s *Basegrulev3Listener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *Basegrulev3Listener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *Basegrulev3Listener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *Basegrulev3Listener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterGrl is called when production grl is entered.
func (s *Basegrulev3Listener) EnterGrl(ctx *GrlContext) {}

// ExitGrl is called when production grl is exited.
func (s *Basegrulev3Listener) ExitGrl(ctx *GrlContext) {}

// EnterRuleEntry is called when production ruleEntry is entered.
func (s *Basegrulev3Listener) EnterRuleEntry(ctx *RuleEntryContext) {}

// ExitRuleEntry is called when production ruleEntry is exited.
func (s *Basegrulev3Listener) ExitRuleEntry(ctx *RuleEntryContext) {}

// EnterSalience is called when production salience is entered.
func (s *Basegrulev3Listener) EnterSalience(ctx *SalienceContext) {}

// ExitSalience is called when production salience is exited.
func (s *Basegrulev3Listener) ExitSalience(ctx *SalienceContext) {}

// EnterRuleName is called when production ruleName is entered.
func (s *Basegrulev3Listener) EnterRuleName(ctx *RuleNameContext) {}

// ExitRuleName is called when production ruleName is exited.
func (s *Basegrulev3Listener) ExitRuleName(ctx *RuleNameContext) {}

// EnterRuleDescription is called when production ruleDescription is entered.
func (s *Basegrulev3Listener) EnterRuleDescription(ctx *RuleDescriptionContext) {}

// ExitRuleDescription is called when production ruleDescription is exited.
func (s *Basegrulev3Listener) ExitRuleDescription(ctx *RuleDescriptionContext) {}

// EnterWhenScope is called when production whenScope is entered.
func (s *Basegrulev3Listener) EnterWhenScope(ctx *WhenScopeContext) {}

// ExitWhenScope is called when production whenScope is exited.
func (s *Basegrulev3Listener) ExitWhenScope(ctx *WhenScopeContext) {}

// EnterThenScope is called when production thenScope is entered.
func (s *Basegrulev3Listener) EnterThenScope(ctx *ThenScopeContext) {}

// ExitThenScope is called when production thenScope is exited.
func (s *Basegrulev3Listener) ExitThenScope(ctx *ThenScopeContext) {}

// EnterThenExpressionList is called when production thenExpressionList is entered.
func (s *Basegrulev3Listener) EnterThenExpressionList(ctx *ThenExpressionListContext) {}

// ExitThenExpressionList is called when production thenExpressionList is exited.
func (s *Basegrulev3Listener) ExitThenExpressionList(ctx *ThenExpressionListContext) {}

// EnterThenExpression is called when production thenExpression is entered.
func (s *Basegrulev3Listener) EnterThenExpression(ctx *ThenExpressionContext) {}

// ExitThenExpression is called when production thenExpression is exited.
func (s *Basegrulev3Listener) ExitThenExpression(ctx *ThenExpressionContext) {}

// EnterAssignment is called when production assignment is entered.
func (s *Basegrulev3Listener) EnterAssignment(ctx *AssignmentContext) {}

// ExitAssignment is called when production assignment is exited.
func (s *Basegrulev3Listener) ExitAssignment(ctx *AssignmentContext) {}

// EnterExpression is called when production expression is entered.
func (s *Basegrulev3Listener) EnterExpression(ctx *ExpressionContext) {}

// ExitExpression is called when production expression is exited.
func (s *Basegrulev3Listener) ExitExpression(ctx *ExpressionContext) {}

// EnterMulDivOperators is called when production mulDivOperators is entered.
func (s *Basegrulev3Listener) EnterMulDivOperators(ctx *MulDivOperatorsContext) {}

// ExitMulDivOperators is called when production mulDivOperators is exited.
func (s *Basegrulev3Listener) ExitMulDivOperators(ctx *MulDivOperatorsContext) {}

// EnterAddMinusOperators is called when production addMinusOperators is entered.
func (s *Basegrulev3Listener) EnterAddMinusOperators(ctx *AddMinusOperatorsContext) {}

// ExitAddMinusOperators is called when production addMinusOperators is exited.
func (s *Basegrulev3Listener) ExitAddMinusOperators(ctx *AddMinusOperatorsContext) {}

// EnterComparisonOperator is called when production comparisonOperator is entered.
func (s *Basegrulev3Listener) EnterComparisonOperator(ctx *ComparisonOperatorContext) {}

// ExitComparisonOperator is called when production comparisonOperator is exited.
func (s *Basegrulev3Listener) ExitComparisonOperator(ctx *ComparisonOperatorContext) {}

// EnterAndLogicOperator is called when production andLogicOperator is entered.
func (s *Basegrulev3Listener) EnterAndLogicOperator(ctx *AndLogicOperatorContext) {}

// ExitAndLogicOperator is called when production andLogicOperator is exited.
func (s *Basegrulev3Listener) ExitAndLogicOperator(ctx *AndLogicOperatorContext) {}

// EnterOrLogicOperator is called when production orLogicOperator is entered.
func (s *Basegrulev3Listener) EnterOrLogicOperator(ctx *OrLogicOperatorContext) {}

// ExitOrLogicOperator is called when production orLogicOperator is exited.
func (s *Basegrulev3Listener) ExitOrLogicOperator(ctx *OrLogicOperatorContext) {}

// EnterExpressionAtom is called when production expressionAtom is entered.
func (s *Basegrulev3Listener) EnterExpressionAtom(ctx *ExpressionAtomContext) {}

// ExitExpressionAtom is called when production expressionAtom is exited.
func (s *Basegrulev3Listener) ExitExpressionAtom(ctx *ExpressionAtomContext) {}

// EnterConstant is called when production constant is entered.
func (s *Basegrulev3Listener) EnterConstant(ctx *ConstantContext) {}

// ExitConstant is called when production constant is exited.
func (s *Basegrulev3Listener) ExitConstant(ctx *ConstantContext) {}

// EnterVariable is called when production variable is entered.
func (s *Basegrulev3Listener) EnterVariable(ctx *VariableContext) {}

// ExitVariable is called when production variable is exited.
func (s *Basegrulev3Listener) ExitVariable(ctx *VariableContext) {}

// EnterArrayMapSelector is called when production arrayMapSelector is entered.
func (s *Basegrulev3Listener) EnterArrayMapSelector(ctx *ArrayMapSelectorContext) {}

// ExitArrayMapSelector is called when production arrayMapSelector is exited.
func (s *Basegrulev3Listener) ExitArrayMapSelector(ctx *ArrayMapSelectorContext) {}

// EnterMemberVariable is called when production memberVariable is entered.
func (s *Basegrulev3Listener) EnterMemberVariable(ctx *MemberVariableContext) {}

// ExitMemberVariable is called when production memberVariable is exited.
func (s *Basegrulev3Listener) ExitMemberVariable(ctx *MemberVariableContext) {}

// EnterFunctionCall is called when production functionCall is entered.
func (s *Basegrulev3Listener) EnterFunctionCall(ctx *FunctionCallContext) {}

// ExitFunctionCall is called when production functionCall is exited.
func (s *Basegrulev3Listener) ExitFunctionCall(ctx *FunctionCallContext) {}

// EnterMethodCall is called when production methodCall is entered.
func (s *Basegrulev3Listener) EnterMethodCall(ctx *MethodCallContext) {}

// ExitMethodCall is called when production methodCall is exited.
func (s *Basegrulev3Listener) ExitMethodCall(ctx *MethodCallContext) {}

// EnterArgumentList is called when production argumentList is entered.
func (s *Basegrulev3Listener) EnterArgumentList(ctx *ArgumentListContext) {}

// ExitArgumentList is called when production argumentList is exited.
func (s *Basegrulev3Listener) ExitArgumentList(ctx *ArgumentListContext) {}

// EnterFloatLiteral is called when production floatLiteral is entered.
func (s *Basegrulev3Listener) EnterFloatLiteral(ctx *FloatLiteralContext) {}

// ExitFloatLiteral is called when production floatLiteral is exited.
func (s *Basegrulev3Listener) ExitFloatLiteral(ctx *FloatLiteralContext) {}

// EnterDecimalFloatLiteral is called when production decimalFloatLiteral is entered.
func (s *Basegrulev3Listener) EnterDecimalFloatLiteral(ctx *DecimalFloatLiteralContext) {}

// ExitDecimalFloatLiteral is called when production decimalFloatLiteral is exited.
func (s *Basegrulev3Listener) ExitDecimalFloatLiteral(ctx *DecimalFloatLiteralContext) {}

// EnterHexadecimalFloatLiteral is called when production hexadecimalFloatLiteral is entered.
func (s *Basegrulev3Listener) EnterHexadecimalFloatLiteral(ctx *HexadecimalFloatLiteralContext) {}

// ExitHexadecimalFloatLiteral is called when production hexadecimalFloatLiteral is exited.
func (s *Basegrulev3Listener) ExitHexadecimalFloatLiteral(ctx *HexadecimalFloatLiteralContext) {}

// EnterIntegerLiteral is called when production integerLiteral is entered.
func (s *Basegrulev3Listener) EnterIntegerLiteral(ctx *IntegerLiteralContext) {}

// ExitIntegerLiteral is called when production integerLiteral is exited.
func (s *Basegrulev3Listener) ExitIntegerLiteral(ctx *IntegerLiteralContext) {}

// EnterDecimalLiteral is called when production decimalLiteral is entered.
func (s *Basegrulev3Listener) EnterDecimalLiteral(ctx *DecimalLiteralContext) {}

// ExitDecimalLiteral is called when production decimalLiteral is exited.
func (s *Basegrulev3Listener) ExitDecimalLiteral(ctx *DecimalLiteralContext) {}

// EnterHexadecimalLiteral is called when production hexadecimalLiteral is entered.
func (s *Basegrulev3Listener) EnterHexadecimalLiteral(ctx *HexadecimalLiteralContext) {}

// ExitHexadecimalLiteral is called when production hexadecimalLiteral is exited.
func (s *Basegrulev3Listener) ExitHexadecimalLiteral(ctx *HexadecimalLiteralContext) {}

// EnterOctalLiteral is called when production octalLiteral is entered.
func (s *Basegrulev3Listener) EnterOctalLiteral(ctx *OctalLiteralContext) {}

// ExitOctalLiteral is called when production octalLiteral is exited.
func (s *Basegrulev3Listener) ExitOctalLiteral(ctx *OctalLiteralContext) {}

// EnterStringLiteral is called when production stringLiteral is entered.
func (s *Basegrulev3Listener) EnterStringLiteral(ctx *StringLiteralContext) {}

// ExitStringLiteral is called when production stringLiteral is exited.
func (s *Basegrulev3Listener) ExitStringLiteral(ctx *StringLiteralContext) {}

// EnterBooleanLiteral is called when production booleanLiteral is entered.
func (s *Basegrulev3Listener) EnterBooleanLiteral(ctx *BooleanLiteralContext) {}

// ExitBooleanLiteral is called when production booleanLiteral is exited.
func (s *Basegrulev3Listener) ExitBooleanLiteral(ctx *BooleanLiteralContext) {}
