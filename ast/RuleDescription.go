package ast

func NewRuleDescription(text string) *RuleDescription {
	return &RuleDescription{
		Text: text,
	}
}

type RuleDescription struct {
	Text string
}

type RuleDescriptionReceiver interface {
	AcceptRuleDescription(ruleDiscription *RuleDescription) error
}
