package ast

// NewRuleName creates RuleName ast object
func NewRuleName(simpleName string) *RuleName {
	return &RuleName{
		SimpleName: simpleName,
	}
}

// RuleName is a simple ast object to hold rule name
type RuleName struct {
	SimpleName string
}

// RuleNameReceiver should be implemented by AST object who stores RuleName
type RuleNameReceiver interface {
	AcceptRuleName(ruleName *RuleName) error
}
