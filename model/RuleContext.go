package model

import (
	log "github.com/sirupsen/logrus"
)

// NewRuleContext will create a new instance of `RuleContext`
func NewRuleContext() *RuleContext {
	return &RuleContext{
		ExpressionAtoms: make([]*ExpressionAtom, 0),
	}
}

// RuleContext serve a context information about some rule executions.
// for now, this rule  not yet used.
type RuleContext struct {
	ExpressionAtoms []*ExpressionAtom
}

// Reset will reset all expression atoms as if they were never evaluated before.
func (rc *RuleContext) Reset() {

	//
	for _, eat := range rc.ExpressionAtoms {
		eat.Reset()
	}
}

// ResetVariable will reset any ExpressionAtom that contains
// a speciffic variablee, be it a variable expression or variable specified in function / method arguments.
//
// Grule will try it best to find such variable changes in the assigment expression.
// But it obviously can't detect variable change in a struct if the variable is changed from within the
// struct or modified from outside grule script. Thus, the rule developer should manually call this
// function to have relevant variable to reset-ed.
func (rc *RuleContext) ResetVariable(variable string) {
	toReset := rc.findExpressionAtomContainsVariable(variable)
	log.Tracef("Resetting %d expression atoms.\n", len(toReset))
	for _, eat := range toReset {
		log.Tracef("Resetting : %s\n", eat.Text)
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
	if ea.FunctionCall != nil && ea.FunctionCall.FunctionName == "Now" {
		return ea
	}
	for _, eat := range rc.ExpressionAtoms {
		if ea.EqualsTo(eat) {
			return eat
		}
	}
	rc.ExpressionAtoms = append(rc.ExpressionAtoms, ea)
	return ea
}

// findExpressionAtomContainsVariable will find any of the ExpressionAtoms that may relate to a speciffic variable.
// this way we can look for ExpressionAtom to reset if we see change in them and then we can reset it status so
// it can get re-evaluated.
func (rc *RuleContext) findExpressionAtomContainsVariable(varName string) []*ExpressionAtom {
	eas := make([]*ExpressionAtom, 0)
	for _, eat := range rc.ExpressionAtoms {
		if eat.IsContainVariable(eat, varName) {
			eas = append(eas, eat)
		}
	}
	return eas
}
