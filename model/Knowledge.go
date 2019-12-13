package model

// KnowledgeBase hold list of rule entry to be evaluated in each cycle.
type KnowledgeBase struct {
	RuleEntries map[string]*RuleEntry
}

// NewKnowledgeBase create new instance of knowledge
func NewKnowledgeBase() *KnowledgeBase {
	return &KnowledgeBase{
		RuleEntries: make(map[string]*RuleEntry),
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
