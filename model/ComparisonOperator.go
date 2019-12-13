package model

var (
	// ComparisonOperatorGT a comparison operator for symbol >
	ComparisonOperatorGT = ComparisonOperator(">")
	// ComparisonOperatorLT a comparison operator for symbol <
	ComparisonOperatorLT = ComparisonOperator("<")
	// ComparisonOperatorGTE a comparison operator for symbol >=
	ComparisonOperatorGTE = ComparisonOperator(">=")
	// ComparisonOperatorLTE a comparison operator for symbol <=
	ComparisonOperatorLTE = ComparisonOperator("<=")
	// ComparisonOperatorEQ a comparison operator for symbol ==
	ComparisonOperatorEQ = ComparisonOperator("==")
	// ComparisonOperatorNEQ a comparison operator for symbol !=
	ComparisonOperatorNEQ = ComparisonOperator("!=")
)

// ComparisonOperator operator symbol for mathematical comparison
type ComparisonOperator string
