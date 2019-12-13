package model

// ExpressionHolder defines a graph that should be able to hold an expression.
type ExpressionHolder interface {
	AcceptExpression(expression *Expression) error
}
