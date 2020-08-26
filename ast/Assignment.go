package ast

import (
	"bytes"
	"errors"
	"github.com/google/uuid"
	"github.com/hyperjumptech/grule-rule-engine/pkg"
)

// NewAssignment will create new instance of Assignment AST Node
func NewAssignment() *Assignment {
	return &Assignment{
		AstID: uuid.New().String(),
	}
}

// Assignment ast node to store assigment expression.
type Assignment struct {
	AstID   string
	GrlText string

	Variable   *Variable
	Expression *Expression
}

// AssignmentReceiver must be implemented by all other ast graph that uses an assigment expression
type AssignmentReceiver interface {
	AcceptAssignment(assignment *Assignment) error
}

// Clone will clone this Assignment. The new clone will have an identical structure
func (e *Assignment) Clone(cloneTable *pkg.CloneTable) *Assignment {
	clone := &Assignment{
		AstID:   uuid.New().String(),
		GrlText: e.GrlText,
	}
	if e.Variable != nil {
		if cloneTable.IsCloned(e.Variable.AstID) {
			clone.Variable = cloneTable.Records[e.Variable.AstID].CloneInstance.(*Variable)
		} else {
			cloned := e.Variable.Clone(cloneTable)
			clone.Variable = cloned
			cloneTable.MarkCloned(e.Variable.AstID, cloned.AstID, e.Variable, cloned)
		}
	}
	if e.Expression != nil {
		if cloneTable.IsCloned(e.Expression.AstID) {
			clone.Expression = cloneTable.Records[e.Expression.AstID].CloneInstance.(*Expression)
		} else {
			cloned := e.Expression.Clone(cloneTable)
			clone.Expression = cloned
			cloneTable.MarkCloned(e.Expression.AstID, cloned.AstID, e.Expression, cloned)
		}
	}
	return clone
}

// AcceptExpression will accept an Expression AST graph into this ast graph
func (e *Assignment) AcceptExpression(exp *Expression) error {
	if e.Expression != nil {
		return errors.New("expression for assignment already assigned")
	}
	e.Expression = exp
	return nil
}

// AcceptVariable will accept an Variable AST graph into this ast graph
func (e *Assignment) AcceptVariable(vari *Variable) error {
	if e.Variable != nil {
		return errors.New("variable for assignment already assigned")
	}
	e.Variable = vari
	return nil
}

// GetAstID get the UUID asigned for this AST graph node
func (e *Assignment) GetAstID() string {
	return e.AstID
}

// GetGrlText get the expression syntax related to this graph when it wast constructed
func (e *Assignment) GetGrlText() string {
	return e.GrlText
}

// GetSnapshot will create a structure signature or AST graph
func (e *Assignment) GetSnapshot() string {
	var buff bytes.Buffer
	buff.WriteString(ASSIGMENT)
	buff.WriteString("(")
	buff.WriteString(e.Variable.GetSnapshot())
	buff.WriteString("=")
	buff.WriteString(e.Expression.GetSnapshot())
	buff.WriteString(")")
	return buff.String()
}

// SetGrlText set the expression syntax related to this graph when it was constructed. Only ANTLR4 listener should
// call this function.
func (e *Assignment) SetGrlText(grlText string) {
	e.GrlText = grlText
}

// Execute will execute this graph in the Then scope
func (e *Assignment) Execute(dataContext IDataContext, memory *WorkingMemory) error {
	exprVal, err := e.Expression.Evaluate(dataContext, memory)
	if err != nil {
		return err
	}
	return e.Variable.Assign(exprVal, dataContext, memory)
}
