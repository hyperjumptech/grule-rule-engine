package ast

import (
	"bytes"
	"fmt"
	"reflect"

	"github.com/google/uuid"
	"github.com/hyperjumptech/grule-rule-engine/pkg"
)

// NewRuleEntry create new instance of RuleEntry
func NewRuleEntry() *RuleEntry {
	return &RuleEntry{
		AstID:    uuid.New().String(),
		Salience: NewSalience(0),
	}
}

// RuleEntry AST graph node
type RuleEntry struct {
	AstID   string
	GrlText string

	RuleName        *RuleName
	RuleDescription *RuleDescription
	Salience        *Salience
	WhenScope       *WhenScope
	ThenScope       *ThenScope

	Retracted bool
}

// RuleEntryReceiver should be implemented by any rule AST object that receive a RuleEntry
type RuleEntryReceiver interface {
	ReceiveRuleEntry(entry *RuleEntry) error
}

// Clone will clone this RuleEntry. The new clone will have an identical structure
func (e *RuleEntry) Clone(cloneTable *pkg.CloneTable) *RuleEntry {
	clone := &RuleEntry{
		AstID:           uuid.New().String(),
		GrlText:         e.GrlText,
		RuleName:        e.RuleName,
		RuleDescription: e.RuleDescription,
		Salience:        e.Salience,
		Retracted:       false,
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

	if e.GetSnapshot() != clone.GetSnapshot() {
		panic(fmt.Sprintf("ThenScope clone failed : original [%s] clone [%s]", e.GetSnapshot(), clone.GetSnapshot()))
	}

	return clone
}

// AcceptRuleName will accept rule name ast object into this rule entry
func (e *RuleEntry) AcceptRuleName(ruleName *RuleName) error {
	e.RuleName = ruleName
	return nil
}

// AcceptRuleDescription will accept rule description ast object into this rule entry
func (e *RuleEntry) AcceptRuleDescription(ruleDescription *RuleDescription) error {
	e.RuleDescription = ruleDescription
	return nil
}

// AcceptSalience will accept rule salience ast object into this rule entry
func (e *RuleEntry) AcceptSalience(salience *Salience) error {
	e.Salience = salience
	return nil
}

// AcceptWhenScope will accept when scope ast object into this rule entry
func (e *RuleEntry) AcceptWhenScope(when *WhenScope) error {
	e.WhenScope = when
	return nil
}

// AcceptThenScope will accept then scope ast object into this rule entry
func (e *RuleEntry) AcceptThenScope(then *ThenScope) error {
	e.ThenScope = then
	return nil
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
	buff.WriteString(fmt.Sprintf("N:%s DEC:\"%s\" SAL:%d W:%s T:%s}", e.RuleName.SimpleName, e.RuleDescription.Text, e.Salience, e.WhenScope.GetSnapshot(), e.ThenScope.GetSnapshot()))
	buff.WriteString(")")
	return buff.String()
}

// SetGrlText set the expression syntax related to this graph when it was constructed. Only ANTLR4 listener should
// call this function.
func (e *RuleEntry) SetGrlText(grlText string) {
	e.GrlText = grlText
}

// Evaluate will evaluate this AST graph for when scope evaluation
func (e *RuleEntry) Evaluate(dataContext IDataContext, memory *WorkingMemory) (bool, error) {
	if e.Retracted {
		return false, nil
	}
	val, err := e.WhenScope.Evaluate(dataContext, memory)
	if err != nil {
		AstLog.Errorf("Error while evaluating rule %s, got %v", e.RuleName.SimpleName, err)
		return false, err
	}
	if val.Kind() != reflect.Bool {
		return false, fmt.Errorf("expression in when is not a boolean expression : %s", e.WhenScope.Expression.GetGrlText())
	}
	return val.Bool(), nil
}

// Execute will execute this graph in the Then scope
func (e *RuleEntry) Execute(dataContext IDataContext, memory *WorkingMemory) (err error) {
	if e.ThenScope == nil {
		return fmt.Errorf("RuleEntry %s have no then scope", e.RuleName.SimpleName)
	}
	defer func() {
		if r := recover(); r != nil {
			err = fmt.Errorf("rule engine execute panic ! recovered : %v", r)
		}
	}()
	return e.ThenScope.Execute(dataContext, memory)
}
