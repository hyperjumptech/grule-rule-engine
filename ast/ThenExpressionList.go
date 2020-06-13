package ast

import (
	"bytes"
	"github.com/google/uuid"
	"github.com/hyperjumptech/grule-rule-engine/pkg"
)

// NewThenExpressionList creates new instance of ThenExpressionList
func NewThenExpressionList() *ThenExpressionList {
	return &ThenExpressionList{
		AstID:           uuid.New().String(),
		ThenExpressions: make([]*ThenExpression, 0),
	}
}

// ThenExpressionList AST graph node
type ThenExpressionList struct {
	AstID         string
	GrlText       string
	DataContext   *DataContext
	WorkingMemory *WorkingMemory

	ThenExpressions []*ThenExpression
}

// Clone will clone this ThenExpressionList. The new clone will have an identical structure
func (e ThenExpressionList) Clone(cloneTable *pkg.CloneTable) *ThenExpressionList {
	clone := &ThenExpressionList{
		AstID:         uuid.New().String(),
		GrlText:       e.GrlText,
		DataContext:   nil,
		WorkingMemory: nil,
	}

	if e.ThenExpressions != nil {
		clone.ThenExpressions = make([]*ThenExpression, len(e.ThenExpressions))
		for k, expr := range e.ThenExpressions {
			if cloneTable.IsCloned(expr.AstID) {
				clone.ThenExpressions[k] = cloneTable.Records[expr.AstID].CloneInstance.(*ThenExpression)
			} else {
				cloned := expr.Clone(cloneTable)
				clone.ThenExpressions[k] = cloned
				cloneTable.MarkCloned(expr.AstID, cloned.AstID, expr, cloned)
			}
		}
	}

	return clone
}

// InitializeContext will initialize this AST graph with data context and working memory before running rule on them.
func (e *ThenExpressionList) InitializeContext(dataCtx *DataContext, WorkingMemory *WorkingMemory) {
	e.DataContext = dataCtx
	e.WorkingMemory = WorkingMemory
	if e.ThenExpressions != nil {
		for _, te := range e.ThenExpressions {
			te.InitializeContext(dataCtx, WorkingMemory)
		}
	}
}

// GetAstID get the UUID asigned for this AST graph node
func (e *ThenExpressionList) GetAstID() string {
	return e.AstID
}

// GetGrlText get the expression syntax related to this graph when it wast constructed
func (e *ThenExpressionList) GetGrlText() string {
	return e.GrlText
}

// GetSnapshot will create a structure signature or AST graph
func (e *ThenExpressionList) GetSnapshot() string {
	var buff bytes.Buffer
	for _, es := range e.ThenExpressions {
		buff.WriteString(es.GetSnapshot())
		buff.WriteString("; ")
	}
	return buff.String()
}

// SetGrlText set the expression syntax related to this graph when it was constructed. Only ANTLR4 listener should
// call this function.
func (e *ThenExpressionList) SetGrlText(grlText string) {
	e.GrlText = grlText
}

// Execute will execute this graph in the Then scope
func (e *ThenExpressionList) Execute() error {
	for _, es := range e.ThenExpressions {
		err := es.Execute()
		if err != nil {
			return err
		}
	}
	return nil
}
