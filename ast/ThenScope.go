package ast

import (
	"bytes"
	"github.com/google/uuid"
	"github.com/hyperjumptech/grule-rule-engine/pkg"
)

// NewThenScope will create new instance of ThenScope
func NewThenScope() *ThenScope {
	return &ThenScope{
		AstID: uuid.New().String(),
	}
}

// ThenScope AST graph node
type ThenScope struct {
	AstID   string
	GrlText string

	ThenExpressionList *ThenExpressionList
}

// ThenScopeReceiver must be implemented by any AST object that will hold a ThenScope
type ThenScopeReceiver interface {
	AcceptThenScope(thenScope *ThenScope) error
}

// Clone will clone this ThenScope. The new clone will have an identical structure
func (e *ThenScope) Clone(cloneTable *pkg.CloneTable) *ThenScope {
	clone := &ThenScope{
		AstID:   uuid.New().String(),
		GrlText: e.GrlText,
	}

	if e.ThenExpressionList != nil {
		if cloneTable.IsCloned(e.ThenExpressionList.AstID) {
			clone.ThenExpressionList = cloneTable.Records[e.ThenExpressionList.AstID].CloneInstance.(*ThenExpressionList)
		} else {
			cloned := e.ThenExpressionList.Clone(cloneTable)
			clone.ThenExpressionList = cloned
			cloneTable.MarkCloned(e.ThenExpressionList.AstID, cloned.AstID, e.ThenExpressionList, cloned)
		}
	}

	return clone
}

// AcceptThenExpressionList will accept ThenExpressionList graph into this ThenScope
func (e *ThenScope) AcceptThenExpressionList(list *ThenExpressionList) error {
	e.ThenExpressionList = list
	return nil
}

// GetAstID get the UUID asigned for this AST graph node
func (e *ThenScope) GetAstID() string {
	return e.AstID
}

// GetGrlText get the expression syntax related to this graph when it wast constructed
func (e *ThenScope) GetGrlText() string {
	return e.GrlText
}

// GetSnapshot will create a structure signature or AST graph
func (e *ThenScope) GetSnapshot() string {
	var buff bytes.Buffer
	buff.WriteString(THENSCOPE)
	buff.WriteString("(")
	buff.WriteString(e.ThenExpressionList.GetSnapshot())
	buff.WriteString(")")
	return buff.String()
}

// SetGrlText set the expression syntax related to this graph when it was constructed. Only ANTLR4 listener should
// call this function.
func (e *ThenScope) SetGrlText(grlText string) {
	e.GrlText = grlText
}

// Execute will execute this graph in the Then scope
func (e *ThenScope) Execute(dataContext IDataContext, memory *WorkingMemory) error {
	if e.ThenExpressionList == nil {
		AstLog.Warnf("Can not execute nil expression list")
	}
	return e.ThenExpressionList.Execute(dataContext, memory)
}
