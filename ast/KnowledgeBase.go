//  Copyright hyperjumptech/grule-rule-engine Authors
//
//  Licensed under the Apache License, Version 2.0 (the "License");
//  you may not use this file except in compliance with the License.
//  You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
//  Unless required by applicable law or agreed to in writing, software
//  distributed under the License is distributed on an "AS IS" BASIS,
//  WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//  See the License for the specific language governing permissions and
//  limitations under the License.

package ast

import (
	"bytes"
	"fmt"
	"sort"
	"strings"
	"sync"

	"github.com/hyperjumptech/grule-rule-engine/pkg"
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
		newClone := kb.Clone(pkg.NewCloneTable())
		if kb.IsIdentical(newClone) {
			AstLog.Debugf("Successfully create instance [%s:%s]", newClone.Name, newClone.Version)
			return newClone
		}
		AstLog.Fatalf("ORIGIN   : %s", kb.GetSnapshot())
		AstLog.Fatalf("CLONE    : %s", newClone.GetSnapshot())
		panic("The clone is not identical")
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
}

// IsIdentical will validate if two KnoledgeBase is identical. Used to validate if the origin and clone is identical.
func (e *KnowledgeBase) IsIdentical(that *KnowledgeBase) bool {
	return e.GetSnapshot() == that.GetSnapshot()
}

// GetSnapshot will create this knowledge base signature
func (e *KnowledgeBase) GetSnapshot() string {
	var buffer bytes.Buffer
	buffer.WriteString(fmt.Sprintf("%s:%s[", e.Name, e.Version))
	keys := make([]string, 0)
	for i := range e.RuleEntries {
		keys = append(keys, i)
	}
	sort.SliceStable(keys, func(i, j int) bool {
		return strings.Compare(keys[i], keys[j]) >= 0
	})
	for i, k := range keys {
		if i > 0 {
			buffer.WriteString(",")
		}
		buffer.WriteString(e.RuleEntries[k].GetSnapshot())
	}
	buffer.WriteString("]")
	return buffer.String()
}

// Clone will clone this instance of KnowledgeBase and produce another (structure wise) identical instance.
func (e *KnowledgeBase) Clone(cloneTable *pkg.CloneTable) *KnowledgeBase {
	clone := &KnowledgeBase{
		Name:        e.Name,
		Version:     e.Version,
		RuleEntries: make(map[string]*RuleEntry),
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
	if e.ContainsRuleEntry(entry.RuleName) {
		return fmt.Errorf("rule entry %s already exist", entry.RuleName)
	}
	e.RuleEntries[entry.RuleName] = entry
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
	}
}

// InitializeContext will initialize this AST graph with data context and working memory before running rule on them.
func (e *KnowledgeBase) InitializeContext(dataCtx IDataContext) {
	e.DataContext = dataCtx
}

// RetractRule will retract the selected rule for execution on the next cycle.
func (e *KnowledgeBase) RetractRule(ruleName string) {
	for _, re := range e.RuleEntries {
		if re.RuleName == ruleName {
			re.Retracted = true
		}
	}
}

// IsRuleRetracted will check if a certain rule denoted by its rule name is currently retracted
func (e *KnowledgeBase) IsRuleRetracted(ruleName string) bool {
	for _, re := range e.RuleEntries {
		if re.RuleName == ruleName {
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
		}
	}
}
