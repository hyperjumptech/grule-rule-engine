package ast

import (
	"bytes"
	"errors"
	"github.com/google/uuid"
	"reflect"
)

// NewWhenScope creates new instance of WhenScope
func NewWhenScope() *WhenScope {
	return &WhenScope{
		AstID: uuid.New().String(),
	}
}

// WhenScope AST graph node
type WhenScope struct {
	AstID         string
	GrlText       string
	DataContext   *DataContext
	WorkingMemory *WorkingMemory

	Expression *Expression
}

// InitializeContext will initialize this AST graph with data context and working memory before running rule on them.
func (e *WhenScope) InitializeContext(dataCtx *DataContext, WorkingMemory *WorkingMemory) {
	e.DataContext = dataCtx
	e.WorkingMemory = WorkingMemory
	if e.Expression != nil {
		e.Expression.InitializeContext(dataCtx, WorkingMemory)
	}
}

// AcceptExpression will accept Expression AST graph node into this node
func (e *WhenScope) AcceptExpression(exp *Expression) error {
	if e.Expression == nil {
		e.Expression = exp
		return nil
	}
	return errors.New("expression for when scope already assigned")
}

// GetAstID get the UUID asigned for this AST graph node
func (e *WhenScope) GetAstID() string {
	return e.AstID
}

// GetGrlText get the expression syntax related to this graph when it wast constructed
func (e *WhenScope) GetGrlText() string {
	return e.GrlText
}

// GetSnapshot will create a structure signature or AST graph
func (e *WhenScope) GetSnapshot() string {
	var buff bytes.Buffer
	buff.WriteString(" when ")
	buff.WriteString(e.Expression.GetSnapshot())
	return buff.String()
}

// SetGrlText set the expression syntax related to this graph when it was constructed. Only ANTLR4 listener should
// call this function.
func (e *WhenScope) SetGrlText(grlText string) {
	e.GrlText = grlText
}

// Evaluate will evaluate this AST graph for when scope evaluation
func (e *WhenScope) Evaluate() (reflect.Value, error) {
	return e.Expression.Evaluate()
}
