package ast

import (
	"bytes"
	"errors"
	"reflect"

	"github.com/google/uuid"
	"github.com/hyperjumptech/grule-rule-engine/pkg"
)

// NewThenExpression create new instance of ThenExpression
func NewThenExpression() *ThenExpression {
	return &ThenExpression{
		AstID: uuid.New().String(),
	}
}

// ThenExpression AST graph node
type ThenExpression struct {
	AstID         string
	GrlText       string
	DataContext   IDataContext
	WorkingMemory *WorkingMemory

	Assignment   *Assignment
	FunctionCall *FunctionCall
	Variable     *Variable
}

type ThenExpressionReceiver interface {
	AcceptThenExpression(expr *ThenExpression) error
}

// Clone will clone this ThenExpression. The new clone will have an identical structure
func (e ThenExpression) Clone(cloneTable *pkg.CloneTable) *ThenExpression {
	clone := &ThenExpression{
		AstID:         uuid.New().String(),
		GrlText:       e.GrlText,
		DataContext:   nil,
		WorkingMemory: nil,
	}

	if e.Assignment != nil {
		if cloneTable.IsCloned(e.Assignment.AstID) {
			clone.Assignment = cloneTable.Records[e.Assignment.AstID].CloneInstance.(*Assignment)
		} else {
			cloned := e.Assignment.Clone(cloneTable)
			clone.Assignment = cloned
			cloneTable.MarkCloned(e.Assignment.AstID, cloned.AstID, e.Assignment, cloned)
		}
	}

	if e.FunctionCall != nil {
		if cloneTable.IsCloned(e.FunctionCall.AstID) {
			clone.FunctionCall = cloneTable.Records[e.FunctionCall.AstID].CloneInstance.(*FunctionCall)
		} else {
			cloned := e.FunctionCall.Clone(cloneTable)
			clone.FunctionCall = cloned
			cloneTable.MarkCloned(e.FunctionCall.AstID, cloned.AstID, e.FunctionCall, cloned)
		}
	}

	return clone
}

// InitializeContext will initialize this AST graph with data context and working memory before running rule on them.
func (e *ThenExpression) InitializeContext(dataCtx IDataContext, WorkingMemory *WorkingMemory) {
	e.DataContext = dataCtx
	e.WorkingMemory = WorkingMemory
	if e.Assignment != nil {
		e.Assignment.InitializeContext(dataCtx, WorkingMemory)
	}
	if e.FunctionCall != nil {
		e.FunctionCall.InitializeContext(dataCtx, WorkingMemory)
	}
}

func (e *ThenExpression) AcceptAssignment(assignment *Assignment) error {
	e.Assignment = assignment
	return nil
}

// AcceptFunctionCall will accept an FunctionCall AST graph into this ast graph
func (e *ThenExpression) AcceptFunctionCall(fun *FunctionCall) error {
	if e.FunctionCall != nil {
		return errors.New("constant for ThenExpression already assigned")
	}
	e.FunctionCall = fun
	return nil
}

// GetAstID get the UUID asigned for this AST graph node
func (e *ThenExpression) GetAstID() string {
	return e.AstID
}

// GetGrlText get the expression syntax related to this graph when it wast constructed
func (e *ThenExpression) GetGrlText() string {
	return e.GrlText
}

func (e *ThenExpression) AcceptVariable(vari *Variable) error {
	e.Variable = vari
	return nil
}

// GetSnapshot will create a structure signature or AST graph
func (e *ThenExpression) GetSnapshot() string {
	var buff bytes.Buffer
	buff.WriteString(THENEXPRESSION)
	buff.WriteString("(")
	if e.Assignment != nil {
		buff.WriteString(e.Assignment.GetSnapshot())
	} else if e.FunctionCall != nil {
		buff.WriteString(e.FunctionCall.GetSnapshot())
	}
	buff.WriteString(")")
	return buff.String()
}

// SetGrlText set the expression syntax related to this graph when it was constructed. Only ANTLR4 listener should
// call this function.
func (e *ThenExpression) SetGrlText(grlText string) {
	e.GrlText = grlText
}

// Execute will execute this graph in the Then scope
func (e *ThenExpression) Execute() error {
	if e.Assignment != nil {
		return e.Assignment.Execute()
	}
	if e.FunctionCall != nil {
		v := e.DataContext.Get("DEFUNC")
		_, err := e.FunctionCall.Evaluate(reflect.ValueOf(v))
		return err
	}
	if e.Variable != nil {
		_, err := e.Variable.Evaluate()
		return err
	}
	return nil
}
