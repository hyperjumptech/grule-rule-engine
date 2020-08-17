package ast

func NewRuleName(simpleName string) *RuleName {
	return &RuleName{
		SimpleName: simpleName,
	}
}

type RuleName struct {
	SimpleName string
}

type RuleNameReceiver interface {
	AcceptRuleName(ruleName *RuleName) error
}
