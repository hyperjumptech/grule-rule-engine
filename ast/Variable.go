package ast

import (
	"bytes"
	"github.com/google/uuid"
	"reflect"
)

// NewVariable create new instance of Variable
func NewVariable(name string) *Variable {
	return &Variable{
		AstID: uuid.New().String(),
		Name:  name,
	}
}

// Variable AST graph node
type Variable struct {
	AstID         string
	GrlText       string
	DataContext   *DataContext
	WorkingMemory *WorkingMemory

	Name  string
	Value reflect.Value
}

// InitializeContext will initialize this AST graph with data context and working memory before running rule on them.
func (e *Variable) InitializeContext(dataCtx *DataContext, WorkingMemory *WorkingMemory) {
	e.DataContext = dataCtx
	e.WorkingMemory = WorkingMemory
}

// VariableReceiver should be implemented by AST graph node to receive Variable AST graph node
type VariableReceiver interface {
	AcceptVariable(exp *Variable) error
}

// GetAstID get the UUID asigned for this AST graph node
func (e *Variable) GetAstID() string {
	return e.AstID
}

// GetGrlText get the expression syntax related to this graph when it wast constructed
func (e *Variable) GetGrlText() string {
	return e.GrlText
}

// GetSnapshot will create a structure signature or AST graph
func (e *Variable) GetSnapshot() string {
	var buff bytes.Buffer
	buff.WriteString("var:")
	buff.WriteString(e.Name)
	return buff.String()
}

// SetGrlText set the expression syntax related to this graph when it was constructed. Only ANTLR4 listener should
// call this function.
func (e *Variable) SetGrlText(grlText string) {
	e.GrlText = grlText
}

// Evaluate will evaluate this AST graph for when scope evaluation
func (e *Variable) Evaluate() (reflect.Value, error) {
	val, err := e.DataContext.GetValue(e.Name)
	if err != nil {
		return reflect.ValueOf(nil), err
	}
	e.Value = val
	return e.Value, nil
}
