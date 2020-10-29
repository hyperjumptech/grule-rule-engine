package ast

import (
	"bytes"
	"fmt"
	"github.com/google/uuid"
	"github.com/hyperjumptech/grule-rule-engine/pkg"
	"reflect"
)

// NewArrayMapSelector create a new array selector graph
func NewArrayMapSelector() *ArrayMapSelector {
	return &ArrayMapSelector{
		AstID: uuid.New().String(),
	}
}

// ArrayMapSelector an array selector graph containing an expression that act ass array or map selector
type ArrayMapSelector struct {
	AstID   string
	GrlText string

	Expression *Expression

	Value reflect.Value
}

// ArrayMapSelectorReceiver must be implemented by all other ast graph that uses map/array selector
type ArrayMapSelectorReceiver interface {
	AcceptArrayMapSelector(sel *ArrayMapSelector) error
}

// Clone will clone this ArgumentList. The new clone will have an identical structure
func (e *ArrayMapSelector) Clone(cloneTable *pkg.CloneTable) *ArrayMapSelector {
	clone := &ArrayMapSelector{
		AstID:   uuid.New().String(),
		GrlText: e.GrlText,
	}
	if e.Expression != nil {
		if cloneTable.IsCloned(e.Expression.AstID) {
			clone.Expression = cloneTable.Records[e.Expression.AstID].CloneInstance.(*Expression)
		} else {
			clonedExpr := e.Expression.Clone(cloneTable)
			clone.Expression = clonedExpr
			cloneTable.MarkCloned(e.Expression.AstID, clonedExpr.AstID, e.Expression, clonedExpr)
		}
	}
	return clone
}

// AcceptExpression will accept Expression AST graph node into this node
func (e *ArrayMapSelector) AcceptExpression(exp *Expression) error {
	if e.Expression == nil {
		e.Expression = exp
		return nil
	}
	return fmt.Errorf("expression for when scope already assigned")
}

// GetAstID get the UUID asigned for this AST graph node
func (e *ArrayMapSelector) GetAstID() string {
	return e.AstID
}

// GetGrlText get the expression syntax related to this graph when it wast constructed
func (e *ArrayMapSelector) GetGrlText() string {
	return e.GrlText
}

// GetSnapshot will create a structure signature or AST graph
func (e *ArrayMapSelector) GetSnapshot() string {
	var buff bytes.Buffer
	buff.WriteString(MAPARRAYSELECTOR)
	buff.WriteString("(")
	buff.WriteString(e.Expression.GetSnapshot())
	buff.WriteString(")")
	return buff.String()
}

// SetGrlText set the expression syntax related to this graph when it was constructed. Only ANTLR4 listener should
// call this function.
func (e *ArrayMapSelector) SetGrlText(grlText string) {
	e.GrlText = grlText
}

// Evaluate will evaluate this AST graph for when scope evaluation
func (e *ArrayMapSelector) Evaluate(dataContext IDataContext, memory *WorkingMemory) (reflect.Value, error) {
	if e.Expression != nil {
		val, err := e.Expression.Evaluate(dataContext, memory)
		if err != nil {
			return val, err
		}
		e.Value = val
		return val, err
	}
	return reflect.ValueOf(nil), fmt.Errorf("array Map Selector contains no selector expression")
}
