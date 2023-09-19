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
	"context"
	"fmt"
	"github.com/hyperjumptech/grule-rule-engine/ast/unique"
	"reflect"

	"github.com/hyperjumptech/grule-rule-engine/pkg"
)

// NewRuleEntry create new instance of RuleEntry
func NewRuleEntry() *RuleEntry {

	return &RuleEntry{
		AstID:           unique.NewID(),
		RuleName:        "No Name",
		Salience:        0,
		RuleDescription: "No Description",
	}
}

// RuleEntry AST graph node
type RuleEntry struct {
	AstID   string
	GrlText string

	RuleName        string
	RuleDescription string
	Salience        int
	WhenScope       *WhenScope
	ThenScope       *ThenScope

	Retracted bool
	Deleted   bool //If this is true, it will be ignored while execution and fetching the matching rules
}

// MakeCatalog will create a catalog entry from RuleEntry node.
func (e *RuleEntry) MakeCatalog(cat *Catalog) {
	meta := &RuleEntryMeta{
		NodeMeta: NodeMeta{
			AstID:    e.AstID,
			GrlText:  e.GrlText,
			Snapshot: e.GetSnapshot(),
		},
	}
	if cat.AddMeta(e.AstID, meta) {
		if e.WhenScope != nil {
			meta.WhenScopeID = e.WhenScope.AstID
			e.WhenScope.MakeCatalog(cat)
		}
		if e.ThenScope != nil {
			meta.ThenScopeID = e.ThenScope.AstID
			e.ThenScope.MakeCatalog(cat)
		}
		meta.RuleName = e.RuleName
		meta.RuleDescription = e.RuleDescription
		meta.Salience = e.Salience
	}
}

// RuleEntryReceiver should be implemented by any rule AST object that receive a RuleEntry
type RuleEntryReceiver interface {
	ReceiveRuleEntry(entry *RuleEntry) error
}

// AcceptSalience will accept salience value
func (e *RuleEntry) AcceptSalience(salience *Salience) error {
	e.Salience = salience.SalienceValue

	return nil
}

// AcceptWhenScope will accept WhenScope AST Graph into this AST Graph
func (e *RuleEntry) AcceptWhenScope(when *WhenScope) error {
	e.WhenScope = when

	return nil
}

// AcceptThenScope will accept ThenScope AST Graph into this AST Graph
func (e *RuleEntry) AcceptThenScope(thenScope *ThenScope) error {
	e.ThenScope = thenScope

	return nil
}

// Clone will clone this RuleEntry. The new clone will have an identical structure
func (e *RuleEntry) Clone(cloneTable *pkg.CloneTable) *RuleEntry {
	clone := &RuleEntry{
		AstID:           unique.NewID(),
		GrlText:         e.GrlText,
		RuleName:        e.RuleName,
		RuleDescription: e.RuleDescription,
		Salience:        e.Salience,
		Retracted:       false,
		Deleted:         e.Deleted,
	}
	if e.WhenScope != nil {
		if cloneTable.IsCloned(e.WhenScope.AstID) {
			clone.WhenScope = cloneTable.Records[e.WhenScope.AstID].CloneInstance.(*WhenScope)
		} else {
			clonedWhenScope := e.WhenScope.Clone(cloneTable)
			clone.WhenScope = clonedWhenScope
			cloneTable.MarkCloned(e.WhenScope.AstID, clonedWhenScope.AstID, e.WhenScope, clonedWhenScope)
		}
	}

	if e.ThenScope != nil {
		if cloneTable.IsCloned(e.ThenScope.AstID) {
			clone.ThenScope = cloneTable.Records[e.ThenScope.AstID].CloneInstance.(*ThenScope)
		} else {
			clonedThenScope := e.ThenScope.Clone(cloneTable)
			clone.ThenScope = clonedThenScope
			cloneTable.MarkCloned(e.ThenScope.AstID, clonedThenScope.AstID, e.ThenScope, clonedThenScope)
		}
	}

	return clone
}

// GetAstID get the UUID asigned for this AST graph node
func (e *RuleEntry) GetAstID() string {

	return e.AstID
}

// GetGrlText get the expression syntax related to this graph when it wast constructed
func (e *RuleEntry) GetGrlText() string {

	return e.GrlText
}

// GetSnapshot will create a structure signature or AST graph
func (e *RuleEntry) GetSnapshot() string {
	var buff bytes.Buffer
	buff.WriteString(RULEENTRY)
	buff.WriteString("(")
	buff.WriteString(fmt.Sprintf("N:%s DEC:\"%s\" SAL:%d W:%s T:%s}", e.RuleName, e.RuleDescription, e.Salience, e.WhenScope.GetSnapshot(), e.ThenScope.GetSnapshot()))
	buff.WriteString(")")

	return buff.String()
}

// SetGrlText set the expression syntax related to this graph when it was constructed. Only ANTLR4 listener should
// call this function.
func (e *RuleEntry) SetGrlText(grlText string) {
	e.GrlText = grlText
}

// Evaluate will evaluate this AST graph for when scope evaluation
func (e *RuleEntry) Evaluate(ctx context.Context, dataContext IDataContext, memory *WorkingMemory) (can bool, err error) {
	if ctx.Err() != nil {

		return false, fmt.Errorf("context error on evaluating rule %s. got %w", e.RuleName, ctx.Err())
	}
	defer func() {
		if r := recover(); r != nil {
			err = fmt.Errorf("error while evaluating rule %s, panic recovered", e.RuleName)
			can = false
		}
	}()
	if e.Retracted {

		return false, nil
	}
	val, err := e.WhenScope.Evaluate(dataContext, memory)
	if err != nil {
		AstLog.Errorf("Error while evaluating rule %s, got %v", e.RuleName, err)

		return false, fmt.Errorf("evaluating expression in rule '%s' the when raised an error. got %v", dataContext.GetRuleEntry().RuleName, err)
	}
	if val.Kind() != reflect.Bool {

		return false, fmt.Errorf("evaluating expression in rule '%s', the when is not a boolean expression : %s", dataContext.GetRuleEntry().RuleName, e.WhenScope.Expression.GetGrlText())
	}

	return val.Bool(), nil
}

// Execute will execute this graph in the Then scope
func (e *RuleEntry) Execute(ctx context.Context, dataContext IDataContext, memory *WorkingMemory) (err error) {
	if ctx.Err() != nil {

		return fmt.Errorf("context error on executing rule %s. got %w", e.RuleName, ctx.Err())
	}
	if e.ThenScope == nil {

		return fmt.Errorf("RuleEntry %s have no then scope", e.RuleName)
	}
	defer func() {
		if r := recover(); r != nil {
			err = fmt.Errorf("rule engine execute panic on rule %s ! recovered : %v", e.RuleName, r)
		}
	}()

	return e.ThenScope.Execute(dataContext, memory)
}
