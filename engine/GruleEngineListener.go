package engine

import "github.com/hyperjumptech/grule-rule-engine/ast"

// GruleEngineListener is an interface to be implemented by those who want to listen the Engine execution.
type GruleEngineListener interface {
	// EvaluateRuleEntry will be called by the engine if it evaluate a rule entry
	EvaluateRuleEntry(cycle uint64, entry *ast.RuleEntry, candidate bool)
	// ExecuteRuleEntry will be called by the engine if it execute a rule entry in a cycle
	ExecuteRuleEntry(cycle uint64, entry *ast.RuleEntry)
	// BeginCycle will be called by the engine every time it start a new evaluation cycle
	BeginCycle(cycle uint64)
}
