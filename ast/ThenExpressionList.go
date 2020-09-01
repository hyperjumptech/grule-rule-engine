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
	AstID   string
	GrlText string

	ThenExpressions []*ThenExpression
}

// ThenExpressionListReceiver must be implemented by any AST object that hold a ThenExpression list AST object
type ThenExpressionListReceiver interface {
	AcceptThenExpressionList(list *ThenExpressionList) error
}

// AcceptThenExpression will accept ThenExpression AST graph into this ExpressionList
func (e *ThenExpressionList) AcceptThenExpression(expr *ThenExpression) error {
	if e.ThenExpressions == nil {
		e.ThenExpressions = make([]*ThenExpression, 0)
	}
	e.ThenExpressions = append(e.ThenExpressions, expr)
	return nil
}

// Clone will clone this ThenExpressionList. The new clone will have an identical structure
func (e *ThenExpressionList) Clone(cloneTable *pkg.CloneTable) *ThenExpressionList {
	clone := &ThenExpressionList{
		AstID:   uuid.New().String(),
		GrlText: e.GrlText,
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
	buff.WriteString(THENEXPRESSIONLIST)
	buff.WriteString("(")
	for idx, es := range e.ThenExpressions {
		if idx > 0 {
			buff.WriteString(",")
		}
		buff.WriteString(es.GetSnapshot())
	}
	buff.WriteString(")")
	return buff.String()
}

// SetGrlText set the expression syntax related to this graph when it was constructed. Only ANTLR4 listener should
// call this function.
func (e *ThenExpressionList) SetGrlText(grlText string) {
	e.GrlText = grlText
}

// Execute will execute this graph in the Then scope
func (e *ThenExpressionList) Execute(dataContext IDataContext, memory *WorkingMemory) error {
	for _, es := range e.ThenExpressions {
		err := es.Execute(dataContext, memory)
		if err != nil {
			return err
		}
	}
	return nil
}
