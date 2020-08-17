package ast

import (
	"bytes"
	"fmt"
	"reflect"

	"github.com/google/uuid"
	"github.com/hyperjumptech/grule-rule-engine/pkg"
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
	DataContext   IDataContext
	WorkingMemory *WorkingMemory

	Name             string
	Constant         *Constant
	Variable         *Variable
	FunctionCall     *FunctionCall
	ArrayMapSelector *ArrayMapSelector

	Value reflect.Value
}

// Clone will clone this Variable. The new clone will have an identical structure
func (e Variable) Clone(cloneTable *pkg.CloneTable) *Variable {
	clone := &Variable{
		AstID:         uuid.New().String(),
		GrlText:       e.GrlText,
		DataContext:   nil,
		WorkingMemory: nil,
		Name:          e.Name,
	}
	return clone
}

// InitializeContext will initialize this AST graph with data context and working memory before running rule on them.
func (e *Variable) InitializeContext(dataCtx IDataContext, WorkingMemory *WorkingMemory) {
	e.DataContext = dataCtx
	e.WorkingMemory = WorkingMemory
}

// VariableReceiver should be implemented by AST graph node to receive Variable AST graph node
type VariableReceiver interface {
	AcceptVariable(exp *Variable) error
}

func (e *Variable) AcceptVariable(vari *Variable) error {
	e.Variable = vari
	return nil
}

func (e *Variable) AcceptArrayMapSelector(sel *ArrayMapSelector) error {
	e.ArrayMapSelector = sel
	return nil
}

func (e *Variable) AcceptFunctionCall(fu *FunctionCall) error {
	e.FunctionCall = fu
	return nil
}

func (e *Variable) AcceptConstant(con *Constant) error {
	e.Constant = con
	return nil
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
	if len(e.Name) > 0 && e.Variable == nil {
		val, err := e.DataContext.GetValue(e.Name)
		if err != nil {
			return reflect.ValueOf(nil), err
		}
		e.Value = val
		return e.Value, nil
	}
	if e.Constant != nil {
		val, err := e.Constant.Evaluate()
		if err != nil {
			return reflect.ValueOf(nil), err
		}
		e.Value = val
		return e.Value, nil
	}
	if e.Variable != nil && e.FunctionCall != nil {
		varValue, err := e.Variable.Evaluate()
		if err != nil {
			return reflect.ValueOf(nil), err
		}
		val, err := e.FunctionCall.Evaluate(varValue)
		if err != nil {
			return reflect.ValueOf(nil), err
		}
		e.Value = val
		return e.Value, nil
	}
	if e.Variable != nil && len(e.Name) > 0 {
		varValue, err := e.Variable.Evaluate()
		if err != nil {
			return reflect.ValueOf(nil), err
		}
		val, err := pkg.GetAttributeValue(pkg.ValueToInterface(varValue), e.Name)
		if err != nil {
			return reflect.ValueOf(nil), err
		}
		e.Value = val
		return e.Value, nil
	}
	if e.Variable != nil && e.ArrayMapSelector != nil {
		varValue, err := e.Variable.Evaluate()
		if err != nil {
			return reflect.ValueOf(nil), err
		}
		selValue, err := e.ArrayMapSelector.Evaluate()
		if err != nil {
			return reflect.ValueOf(nil), err
		}
		val, err := pkg.GetMapArrayValue(pkg.ValueToInterface(varValue), pkg.ValueToInterface(selValue))
		if err != nil {
			return reflect.ValueOf(nil), err
		}
		e.Value = reflect.ValueOf(val)
		return e.Value, nil
	}

	return reflect.ValueOf(nil), fmt.Errorf("this code part should not be reached")
}
