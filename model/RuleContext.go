package model

func NewRuleContext() *RuleContext {
	return &RuleContext{ExpressionAtoms: make([]*ExpressionAtom, 0)}
}

// RuleContext serve a context information about some rule executions.
// for now, this rule  not yet used.
type RuleContext struct {
	ExpressionAtoms []*ExpressionAtom
}

func (rc *RuleContext) Reset() {
	for _, eat := range rc.ExpressionAtoms {
		eat.Reset()
	}
}

func (rc *RuleContext) Contains(ea *ExpressionAtom) bool {
	for _, eat := range rc.ExpressionAtoms {
		if ea.EqualsTo(eat) {
			return true
		}
	}
	return false
}

func (rc *RuleContext) Add(ea *ExpressionAtom) *ExpressionAtom {
	for _, eat := range rc.ExpressionAtoms {
		if ea.EqualsTo(eat) {
			return eat
		}
	}
	rc.ExpressionAtoms = append(rc.ExpressionAtoms, ea)
	return ea
}
