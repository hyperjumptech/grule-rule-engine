package ast

import (
	"fmt"
	"github.com/hyperjumptech/grule-rule-engine/events"
	"github.com/hyperjumptech/grule-rule-engine/pkg/eventbus"
	"sync"
)

// NewKnowledgeBase create a new instance of KnowledgeBase
func NewKnowledgeBase(name, version string) *KnowledgeBase {
	return &KnowledgeBase{
		Name:        name,
		Version:     version,
		RuleEntries: make(map[string]*RuleEntry),
		Publisher:   eventbus.DefaultBrooker.GetPublisher(events.RuleEntryEventTopic),
	}
}

// KnowledgeBase is a collection of RuleEntries. It has a name and version.
type KnowledgeBase struct {
	lock          sync.Mutex
	Name          string
	Version       string
	DataContext   *DataContext
	WorkingMemory *WorkingMemory
	RuleEntries   map[string]*RuleEntry
	Publisher     *eventbus.Publisher
}

// AddRuleEntry add ruleentry into this knowledge base.
// return an error if a rule entry with the same name already exist in this knowledge base.
func (e *KnowledgeBase) AddRuleEntry(entry *RuleEntry) error {
	e.lock.Lock()
	defer e.lock.Unlock()
	if e.ContainsRuleEntry(entry.Name) {
		return fmt.Errorf("rule entry %s already exist", entry.Name)
	}
	e.RuleEntries[entry.Name] = entry
	if e.DataContext != nil && e.WorkingMemory != nil {
		entry.InitializeContext(e.DataContext, e.WorkingMemory)
	}

	e.Publisher.Publish(&events.RuleEntryEvent{
		EventType: events.RuleEntryAddedEvent,
		RuleName:  entry.Name,
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
		if re.Name == ruleName {
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
				RuleName:  re.Name,
			})
		}
	}
}
