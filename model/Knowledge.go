package model

// KnowledgeBase hold list of rule entry to be evaluated in each cycle.
type KnowledgeBase struct {
	Version     string
	Name        string
	RuleEntries map[string]*RuleEntry
	RuleContext *RuleContext
}

// NewKnowledgeBase create new instance of knowledge
func NewKnowledgeBase(name, version string) *KnowledgeBase {
	return &KnowledgeBase{
		Version:     version,
		Name:        name,
		RuleEntries: make(map[string]*RuleEntry),
		RuleContext: NewRuleContext(),
	}
}

// Retract retract a rule entry from next evaluation cycle.
func (k *KnowledgeBase) Retract(ruleEntryName string) {
	if re, ok := k.RuleEntries[ruleEntryName]; ok {
		re.Retracted = true
	}
}

// Reset will reset the retract status of all rule entries.
func (k *KnowledgeBase) Reset() {
	for _, v := range k.RuleEntries {
		v.Retracted = false
	}
}

// RuleContextReset will reset the rule contexts, render them un-evaluated
func (k *KnowledgeBase) RuleContextReset() {
	k.RuleContext.Reset()
}
