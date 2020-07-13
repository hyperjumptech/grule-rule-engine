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
		Salience: 0,
	}
}

// RuleEntry AST graph node
type RuleEntry struct {
	AstID         string
	GrlText       string
	DataContext   IDataContext
	WorkingMemory *WorkingMemory

	Name        string
	Description string
	Salience    int
	WhenScope   *WhenScope
	ThenScope   *ThenScope

	Retracted bool
}

// Clone will clone this RuleEntry. The new clone will have an identical structure
func (e RuleEntry) Clone(cloneTable *pkg.CloneTable) *RuleEntry {
	clone := &RuleEntry{
		AstID:         uuid.New().String(),
		GrlText:       e.GrlText,
		DataContext:   nil,
		WorkingMemory: nil,
		Name:          e.Name,
		Description:   e.Description,
		Salience:      e.Salience,
		Retracted:     false,
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

// InitializeContext will initialize this AST graph with data context and working memory before running rule on them.
func (e *RuleEntry) InitializeContext(dataCtx IDataContext, WorkingMemory *WorkingMemory) {
	e.DataContext = dataCtx
	e.WorkingMemory = WorkingMemory
	if e.WhenScope != nil {
		e.WhenScope.InitializeContext(dataCtx, WorkingMemory)
	}
	if e.ThenScope != nil {
		e.ThenScope.InitializeContext(dataCtx, WorkingMemory)
	}
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
	buff.WriteString(fmt.Sprintf("rule %s \"%s\" salience %d {%s%s}", e.Name, e.Description, e.Salience, e.WhenScope.GetSnapshot(), e.ThenScope.GetSnapshot()))
	return buff.String()
}

// SetGrlText set the expression syntax related to this graph when it was constructed. Only ANTLR4 listener should
// call this function.
func (e *RuleEntry) SetGrlText(grlText string) {
	e.GrlText = grlText
}

// Evaluate will evaluate this AST graph for when scope evaluation
func (e *RuleEntry) Evaluate() (bool, error) {
	if e.Retracted {
		return false, nil
	}
	val, err := e.WhenScope.Evaluate()
	if err != nil {
		AstLog.Errorf("Error while evaluating rule %s, got %v", e.Name, err)
		return false, err
	}
	if val.Kind() != reflect.Bool {
		return false, fmt.Errorf("expression in when is not a boolean expression : %s", e.WhenScope.Expression.GetGrlText())
	}
	return val.Bool(), nil
}

// Execute will execute this graph in the Then scope
func (e *RuleEntry) Execute() error {
	return e.ThenScope.Execute()
}
