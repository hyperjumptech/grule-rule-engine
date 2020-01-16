package ast

import (
	"bytes"
	"github.com/google/uuid"
)

// NewThenScope will create new instance of ThenScope
func NewThenScope() *ThenScope {
	return &ThenScope{
		AstID: uuid.New().String(),
	}
}

// ThenScope AST graph node
type ThenScope struct {
	AstID         string
	GrlText       string
	DataContext   *DataContext
	WorkingMemory *WorkingMemory

	ThenExpressionList *ThenExpressionList
}

// InitializeContext will initialize this AST graph with data context and working memory before running rule on them.
func (e *ThenScope) InitializeContext(dataCtx *DataContext, WorkingMemory *WorkingMemory) {
	e.DataContext = dataCtx
	e.WorkingMemory = WorkingMemory
	if e.ThenExpressionList != nil {
		e.ThenExpressionList.InitializeContext(dataCtx, WorkingMemory)
	}
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
	buff.WriteString(" THEN ")
	buff.WriteString(e.ThenExpressionList.GetSnapshot())
	return buff.String()
}

// SetGrlText set the expression syntax related to this graph when it was constructed. Only ANTLR4 listener should
// call this function.
func (e *ThenScope) SetGrlText(grlText string) {
	e.GrlText = grlText
}

// Execute will execute this graph in the Then scope
func (e *ThenScope) Execute() error {
	return e.ThenExpressionList.Execute()
}
