package ast

import (
	"bytes"
	"reflect"

	"github.com/google/uuid"
	"github.com/hyperjumptech/grule-rule-engine/pkg"
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

	Arguments []*Expression
}

// Clone will clone this ArgumentList. The new clone will have an identical structure
func (e *ArgumentList) Clone(cloneTable *pkg.CloneTable) *ArgumentList {
	clone := &ArgumentList{
		AstID:   uuid.New().String(),
		GrlText: e.GrlText,
	}
	if e.Arguments != nil {
		clone.Arguments = make([]*Expression, len(e.Arguments))
		for k, expr := range e.Arguments {
			if cloneTable.IsCloned(expr.AstID) {
				clone.Arguments[k] = cloneTable.Records[expr.AstID].CloneInstance.(*Expression)
			} else {
				clonedExpr := expr.Clone(cloneTable)
				clone.Arguments[k] = clonedExpr
				cloneTable.MarkCloned(expr.AstID, clonedExpr.AstID, expr, clonedExpr)
			}
		}
	}
	return clone
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
	buff.WriteString(ARGUMENTLIST)
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
	AcceptArgumentList(argList *ArgumentList) error
}

// Evaluate will evaluate this AST graph for when scope evaluation
func (e *ArgumentList) Evaluate(dataContext IDataContext, memory *WorkingMemory) ([]reflect.Value, error) {
	values := make([]reflect.Value, len(e.Arguments))
	for i, exp := range e.Arguments {
		val, err := exp.Evaluate(dataContext, memory)
		if err != nil {
			return values, err
		}
		values[i] = val
	}
	return values, nil
}
