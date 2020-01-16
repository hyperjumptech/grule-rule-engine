package ast

import (
	"bytes"
	"github.com/google/uuid"
	"reflect"
)

// NewArgumentList create a new instance of ArgumentList
func NewArgumentList() *ArgumentList {
	return &ArgumentList{
		AstID:     uuid.New().String(),
		Arguments: make([]*Expression, 0),
	}
}

// ArgumentList stores AST graph for ArgumentList that contains expression.
type ArgumentList struct {
	AstID   string
	GrlText string

	DataContext   *DataContext
	WorkingMemory *WorkingMemory

	Arguments []*Expression
}

// InitializeContext will initialize this AST graph with data context and working memory before running rule on them.
func (e *ArgumentList) InitializeContext(dataCtx *DataContext, workingMemory *WorkingMemory) {
	e.DataContext = dataCtx
	e.WorkingMemory = workingMemory
	if e.Arguments != nil && len(e.Arguments) > 0 {
		for _, expr := range e.Arguments {
			expr.InitializeContext(dataCtx, workingMemory)
		}
	}
}

// AcceptExpression will accept an expression AST graph into this ast graph
func (e *ArgumentList) AcceptExpression(exp *Expression) error {
	if e.Arguments == nil {
		e.Arguments = make([]*Expression, 0)
	}
	e.Arguments = append(e.Arguments, exp)
	return nil
}

// GetAstID get the UUID asigned for this AST graph node
func (e *ArgumentList) GetAstID() string {
	return e.AstID
}

// GetGrlText get the expression syntax related to this graph when it wast constructed
func (e *ArgumentList) GetGrlText() string {
	return e.GrlText
}

// GetSnapshot will create a structure signature or AST graph
func (e *ArgumentList) GetSnapshot() string {
	var buff bytes.Buffer
	buff.WriteString("(")
	for i, v := range e.Arguments {
		if i > 0 {
			buff.WriteString(",")
		}
		buff.WriteString(v.GetSnapshot())
	}
	buff.WriteString(")")
	return buff.String()
}

// SetGrlText set the expression syntax related to this graph when it was constructed. Only ANTLR4 listener should
// call this function.
func (e *ArgumentList) SetGrlText(grlText string) {
	e.GrlText = grlText
}

// ArgumentListReceiver will accept an ArgumentList AST graph into this ast graph
type ArgumentListReceiver interface {
	AcceptArgumentList(argList *ArgumentList)
}

// Evaluate will evaluate this AST graph for when scope evaluation
func (e *ArgumentList) Evaluate() ([]reflect.Value, error) {
	values := make([]reflect.Value, len(e.Arguments))
	for i, exp := range e.Arguments {
		val, err := exp.Evaluate()
		if err != nil {
			return values, err
		}
		values[i] = val
	}
	return values, nil
}
