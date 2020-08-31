package ast

// NewRuleDescription will create new rule description
func NewRuleDescription(text string) *RuleDescription {
	return &RuleDescription{
		Text: text,
	}
}

// RuleDescription is a simple AST object to store rule description
type RuleDescription struct {
	Text string
}

// RuleDescriptionReceiver must be implemented by any object that will receive rule description
type RuleDescriptionReceiver interface {
	AcceptRuleDescription(ruleDiscription *RuleDescription) error
}
