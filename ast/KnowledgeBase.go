package ast

// NewKnowledgeBase create a new instance of KnowledgeBase
func NewKnowledgeBase(name, version string) *KnowledgeBase {
	return &KnowledgeBase{
		Name:        name,
		Version:     version,
		RuleEntries: make([]*RuleEntry, 0),
	}
}

// KnowledgeBase is a collection of RuleEntries. It has a name and version.
type KnowledgeBase struct {
	Name          string
	Version       string
	DataContext   *DataContext
	WorkingMemory *WorkingMemory
	RuleEntries   []*RuleEntry
}

// InitializeContext will initialize this AST graph with data context and working memory before running rule on them.
func (e *KnowledgeBase) InitializeContext(dataCtx *DataContext, memory *WorkingMemory) {
	e.DataContext = dataCtx
	e.WorkingMemory = memory
	if e.RuleEntries != nil {
		for _, re := range e.RuleEntries {
			re.InitializeContext(dataCtx, memory)
		}
	}
}

// RetractRule will retract the selected rule for execution on the next cycle.
func (e *KnowledgeBase) RetractRule(ruleName string) {
	for _, re := range e.RuleEntries {
		if re.Name == ruleName {
			re.Retracted = true
		}
	}
}

// IsRuleRetracted will check if a certain rule denoted by its rule name is currently retracted
func (e *KnowledgeBase) IsRuleRetracted(ruleName string) bool {
	for _, re := range e.RuleEntries {
		if re.Name == ruleName {
			return re.Retracted
		}
	}
	return false
}

// Reset will restore all rule in the knowledge
func (e *KnowledgeBase) Reset() {
	for _, re := range e.RuleEntries {
		re.Retracted = false
	}
}
