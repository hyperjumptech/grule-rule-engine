package ast

import (
	"bytes"
	"errors"
	"reflect"

	"github.com/google/uuid"
	"github.com/hyperjumptech/grule-rule-engine/pkg"
)

// NewExpressionAtom create new instance of ExpressionAtom
func NewExpressionAtom() *ExpressionAtom {
	return &ExpressionAtom{
		AstID: uuid.New().String(),
	}
}

// ExpressionAtom AST node graph
type ExpressionAtom struct {
	AstID   string
	GrlText string

	Variable     *Variable
	FunctionCall *FunctionCall
	Value        reflect.Value

	Evaluated bool
}

// ExpressionAtomReceiver contains function to be implemented by other AST graph to receive an ExpressionAtom AST graph
type ExpressionAtomReceiver interface {
	AcceptExpressionAtom(exp *ExpressionAtom) error
}

// Clone will clone this ExpressionAtom. The new clone will have an identical structure
func (e *ExpressionAtom) Clone(cloneTable *pkg.CloneTable) *ExpressionAtom {
	clone := &ExpressionAtom{
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

// AcceptVariable will accept an Variable AST graph into this ast graph
func (e *ExpressionAtom) AcceptVariable(vari *Variable) error {
	if e.Variable != nil {
		return errors.New("variable for ExpressionAtom already assigned")
	}
	e.Variable = vari
	return nil
}

// AcceptFunctionCall will accept an FunctionCall AST graph into this ast graph
func (e *ExpressionAtom) AcceptFunctionCall(fun *FunctionCall) error {
	if e.FunctionCall != nil {
		return errors.New("constant for ExpressionAtom already assigned")
	}
	e.FunctionCall = fun
	return nil
}

// GetAstID get the UUID asigned for this AST graph node
func (e *ExpressionAtom) GetAstID() string {
	return e.AstID
}

// GetGrlText get the expression syntax related to this graph when it wast constructed
func (e *ExpressionAtom) GetGrlText() string {
	return e.GrlText
}

// GetSnapshot will create a structure signature or AST graph
func (e *ExpressionAtom) GetSnapshot() string {
	var buff bytes.Buffer
	buff.WriteString(EXPRESSIONATOM)
	buff.WriteString("(")
	if e.Variable != nil {
		buff.WriteString(e.Variable.GetSnapshot())
	} else if e.FunctionCall != nil {
		buff.WriteString(e.FunctionCall.GetSnapshot())
	}
	buff.WriteString(")")
	return buff.String()
}

// SetGrlText set the expression syntax related to this graph when it was constructed. Only ANTLR4 listener should
// call this function.
func (e *ExpressionAtom) SetGrlText(grlText string) {
	e.GrlText = grlText
}

// Evaluate will evaluate this AST graph for when scope evaluation
func (e *ExpressionAtom) Evaluate(dataContext IDataContext, memory *WorkingMemory) (reflect.Value, error) {
	if e.Evaluated == true {
		return e.Value, nil
	}
	if e.Variable != nil {
		val, err := e.Variable.Evaluate(dataContext, memory)
		if err != nil {
			return reflect.Value{}, err
		}
		e.Value = val
		e.Evaluated = true
		return val, err
	}
	if e.FunctionCall != nil {
		valueNode := dataContext.Get("DEFUNC")
		args, err := e.FunctionCall.EvaluateArgumentList(dataContext, memory)
		if err != nil {
			return reflect.Value{}, err
		}
		ret, err := valueNode.CallFunction(e.FunctionCall.FunctionName, args...)
		if err != nil {
			return reflect.Value{}, err
		}
		e.Value = ret
		e.Evaluated = true
		return ret, err
	}
	panic("should not be reached")
}
