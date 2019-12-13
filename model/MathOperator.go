package model

var (
	// MathOperatorMul a math operator symbol for Multiplication
	MathOperatorMul = MathOperator(1)
	// MathOperatorDiv a math operator symbol for Division
	MathOperatorDiv = MathOperator(2)
	// MathOperatorPlus a math operator symbol for Addition
	MathOperatorPlus = MathOperator(3)
	// MathOperatorMinus a math operator symbol for Substraction
	MathOperatorMinus = MathOperator(4)
)

// MathOperator define a constants type of mathematical operator.
type MathOperator int
