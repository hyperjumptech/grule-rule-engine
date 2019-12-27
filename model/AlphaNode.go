package model

// AlphaNode are DSL graph that should be deep-equalizable, since they may be duplicated across the rule.
// In RETE algorithm, Equals AlphaNode must not be duplicated and only use one instance for evaluation. This will improve
// the performance.
type AlphaNode interface {
	// EqualsTo check if this node is equals to the `other` node.
	EqualsTo(other AlphaNode) bool
}
