package ast

import (
	"fmt"
	"sync"

	"github.com/hyperjumptech/grule-rule-engine/events"
	"github.com/hyperjumptech/grule-rule-engine/pkg"
	"github.com/hyperjumptech/grule-rule-engine/pkg/eventbus"
)

// NewKnowledgeLibrary create a new instance KnowledgeLibrary
func NewKnowledgeLibrary() *KnowledgeLibrary {
	return &KnowledgeLibrary{
		Library: make(map[string]*KnowledgeBase),
	}
}

// KnowledgeLibrary is a knowledgebase store.
type KnowledgeLibrary struct {
	Library map[string]*KnowledgeBase
}

// GetKnowledgeBase will get the actual KnowledgeBase blue print that will be used to create instances.
// Although this KnowledgeBase blueprint works, It SHOULD NOT be used directly in the engine.
// You should obtain KnowledgeBase instance by calling NewKnowledgeBaseInstance
func (lib *KnowledgeLibrary) GetKnowledgeBase(name, version string) *KnowledgeBase {
	kb, ok := lib.Library[fmt.Sprintf("%s:%s", name, version)]
	if ok {
		return kb
	}
	kb = &KnowledgeBase{
		Name:          name,
		Version:       version,
		RuleEntries:   make(map[string]*RuleEntry),
		Publisher:     eventbus.DefaultBrooker.GetPublisher(events.RuleEntryEventTopic),
		WorkingMemory: NewWorkingMemory(name, version),
	}
	lib.Library[fmt.Sprintf("%s:%s", name, version)] = kb
	return kb
}

// NewKnowledgeBaseInstance will create a new instance based on KnowledgeBase blue print
// identified by its name and version
func (lib *KnowledgeLibrary) NewKnowledgeBaseInstance(name, version string) *KnowledgeBase {
	kb, ok := lib.Library[fmt.Sprintf("%s:%s", name, version)]
	if ok {
		cTable := pkg.NewCloneTable()
		return kb.Clone(cTable)
	}
	return nil
}

// KnowledgeBase is a collection of RuleEntries. It has a name and version.
type KnowledgeBase struct {
	lock          sync.Mutex
	Name          string
	Version       string
	DataContext   IDataContext
	WorkingMemory *WorkingMemory
	RuleEntries   map[string]*RuleEntry
	Publisher     *eventbus.Publisher
}

// Clone will clone this instance of KnowledgeBase and produce another (structure wise) identical instance.
func (e *KnowledgeBase) Clone(cloneTable *pkg.CloneTable) *KnowledgeBase {
	clone := &KnowledgeBase{
		Name:        e.Name,
		Version:     e.Version,
		RuleEntries: make(map[string]*RuleEntry),
		Publisher:   eventbus.DefaultBrooker.GetPublisher(events.RuleEntryEventTopic),
	}
	if e.RuleEntries != nil {
		for k, entry := range e.RuleEntries {
			if cloneTable.IsCloned(entry.AstID) {
				clone.RuleEntries[k] = cloneTable.Records[entry.AstID].CloneInstance.(*RuleEntry)
			} else {
				cloned := entry.Clone(cloneTable)
				clone.RuleEntries[k] = cloned
				cloneTable.MarkCloned(entry.AstID, cloned.AstID, entry, cloned)
			}
		}
	}
	if e.WorkingMemory != nil {
		clone.WorkingMemory = e.WorkingMemory.Clone(cloneTable)
	}
	return clone
}

// AddRuleEntry add ruleentry into this knowledge base.
// return an error if a rule entry with the same name already exist in this knowledge base.
func (e *KnowledgeBase) AddRuleEntry(entry *RuleEntry) error {
	e.lock.Lock()
	defer e.lock.Unlock()
	if e.ContainsRuleEntry(entry.RuleName.SimpleName) {
		return fmt.Errorf("rule entry %s already exist", entry.RuleName.SimpleName)
	}
	e.RuleEntries[entry.RuleName.SimpleName] = entry
	if e.DataContext != nil && e.WorkingMemory != nil {
		entry.InitializeContext(e.DataContext, e.WorkingMemory)
	}

	e.Publisher.Publish(&events.RuleEntryEvent{
		EventType: events.RuleEntryAddedEvent,
		RuleName:  entry.RuleName.SimpleName,
	})

	return nil
}

// ContainsRuleEntry will check if a rule with such name is already exist in this knowledge base.
func (e *KnowledgeBase) ContainsRuleEntry(name string) bool {
	_, ok := e.RuleEntries[name]
	return ok
}

// RemoveRuleEntry remove the rule entry with specified name from this knowledge base
func (e *KnowledgeBase) RemoveRuleEntry(name string) {
	e.lock.Lock()
	defer e.lock.Unlock()
	if e.ContainsRuleEntry(name) {
		delete(e.RuleEntries, name)
		// emit rule entry remove event
		e.Publisher.Publish(&events.RuleEntryEvent{
			EventType: events.RuleEntryRemovedEvent,
			RuleName:  name,
		})
	}
}

// InitializeContext will initialize this AST graph with data context and working memory before running rule on them.
func (e *KnowledgeBase) InitializeContext(dataCtx IDataContext) {
	e.DataContext = dataCtx
	if e.RuleEntries != nil {
		for _, re := range e.RuleEntries {
			re.InitializeContext(dataCtx, e.WorkingMemory)
		}
	}
}

// RetractRule will retract the selected rule for execution on the next cycle.
func (e *KnowledgeBase) RetractRule(ruleName string) {
	for _, re := range e.RuleEntries {
		if re.RuleName.SimpleName == ruleName {
			re.Retracted = true

			// emit rule entry retract event
			e.Publisher.Publish(&events.RuleEntryEvent{
				EventType: events.RuleEntryRetractedEvent,
				RuleName:  ruleName,
			})
		}
	}
}

// IsRuleRetracted will check if a certain rule denoted by its rule name is currently retracted
func (e *KnowledgeBase) IsRuleRetracted(ruleName string) bool {
	for _, re := range e.RuleEntries {
		if re.RuleName.SimpleName == ruleName {
			return re.Retracted
		}
	}
	return false
}

// Reset will restore all rule in the knowledge
func (e *KnowledgeBase) Reset() {
	for _, re := range e.RuleEntries {
		if re.Retracted {
			re.Retracted = false

			// emit rule entry reset event
			e.Publisher.Publish(&events.RuleEntryEvent{
				EventType: events.RuleEntryResetEvent,
				RuleName:  re.RuleName.SimpleName,
			})
		}
	}
}
