package model

// NewRuleContext will create a new instance of `RuleContext`
func NewRuleContext() *RuleContext {
	return &RuleContext{ExpressionAtoms: make([]*ExpressionAtom, 0)}
}

// RuleContext serve a context information about some rule executions.
// for now, this rule  not yet used.
type RuleContext struct {
	ExpressionAtoms []*ExpressionAtom
}

// Reset will reset all expression atoms as if they were never evaluated before.
func (rc *RuleContext) Reset() {
	for _, eat := range rc.ExpressionAtoms {
		eat.Reset()
	}
}

// Contains will check if a provided expression atom is already in this sets of expression atom,
// this to ensure no duplication.
func (rc *RuleContext) Contains(ea *ExpressionAtom) bool {
	for _, eat := range rc.ExpressionAtoms {
		if ea.EqualsTo(eat) {
			return true
		}
	}
	return false
}

// Add will add an expression atom if its not contained withing this rule context.
// It will return the expression atom from argument if its not exist in this context
// or return the one from this context if its already exist.
// This will ensure no duplication.
func (rc *RuleContext) Add(ea *ExpressionAtom) *ExpressionAtom {
	for _, eat := range rc.ExpressionAtoms {
		if ea.EqualsTo(eat) {
			return eat
		}
	}
	rc.ExpressionAtoms = append(rc.ExpressionAtoms, ea)
	return ea
}
