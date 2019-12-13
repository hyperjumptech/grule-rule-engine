package model

// ExpressionAtomHolder defines all graph that should be able to store an expression atom.
type ExpressionAtomHolder interface {
	AcceptExpressionAtom(exprAtom *ExpressionAtom) error
}
