package model

var (
	// LogicalOperatorAnd serve as operator value of AND logic
	LogicalOperatorAnd = LogicalOperator(1)
	// LogicalOperatorOr serve as operator value of OR logic
	LogicalOperatorOr = LogicalOperator(2)
)

// LogicalOperator a logical operator symbol
type LogicalOperator int
