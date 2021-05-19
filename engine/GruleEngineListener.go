package engine

import "github.com/hyperjumptech/grule-rule-engine/ast"

type GruleEngineListener interface {
	EvaluateRuleEntry(cycle uint64, entry *ast.RuleEntry, candidate bool)
	ExecuteRuleEntry(cycle uint64, entry *ast.RuleEntry)
	BeginCycle(cycle uint64)
}
