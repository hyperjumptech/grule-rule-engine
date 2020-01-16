package ast

import (
	"bytes"
	"errors"
	"github.com/google/uuid"
	"reflect"
)

// NewExpressionAtom create new instance of ExpressionAtom
func NewExpressionAtom() *ExpressionAtom {
	return &ExpressionAtom{
		AstID: uuid.New().String(),
	}
}

// ExpressionAtom AST node graph
type ExpressionAtom struct {
	AstID         string
	GrlText       string
	DataContext   *DataContext
	WorkingMemory *WorkingMemory

	Constant     *Constant
	Variable     *Variable
	FunctionCall *FunctionCall
	MethodCall   *MethodCall
	Value        reflect.Value
}

// InitializeContext will initialize this AST graph with data context and working memory before running rule on them.
func (e *ExpressionAtom) InitializeContext(dataCtx *DataContext, WorkingMemory *WorkingMemory) {
	e.DataContext = dataCtx
	e.WorkingMemory = WorkingMemory
	if e.Constant != nil {
		e.Constant.InitializeContext(dataCtx, WorkingMemory)
	}
	if e.Variable != nil {
		e.Variable.InitializeContext(dataCtx, WorkingMemory)
	}
	if e.FunctionCall != nil {
		e.FunctionCall.InitializeContext(dataCtx, WorkingMemory)
	}
	if e.MethodCall != nil {
		e.MethodCall.InitializeContext(dataCtx, WorkingMemory)
	}
}

// AcceptVariable will accept an Variable AST graph into this ast graph
func (e *ExpressionAtom) AcceptVariable(vari *Variable) error {
	if e.Variable != nil {
		return errors.New("variable for ExpressionAtom already assigned")
	}
	e.Variable = vari
	return nil
}

// AcceptConstant will accept an Constant AST graph into this ast graph
func (e *ExpressionAtom) AcceptConstant(con *Constant) error {
	if e.Constant != nil {
		return errors.New("constant for ExpressionAtom already assigned")
	}
	e.Constant = con
	return nil
}

// AcceptMethodCall will accept an MethodCall AST graph into this ast graph
func (e *ExpressionAtom) AcceptMethodCall(fun *MethodCall) error {
	if e.MethodCall != nil {
		return errors.New("constant for ExpressionAtom already assigned")
	}
	e.MethodCall = fun
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
	if e.Variable != nil {
		buff.WriteString(e.Variable.GetSnapshot())
	} else if e.Constant != nil {
		buff.WriteString(e.Constant.GetSnapshot())
	} else if e.FunctionCall != nil {
		buff.WriteString(e.FunctionCall.GetSnapshot())
	} else if e.MethodCall != nil {
		buff.WriteString(e.MethodCall.GetSnapshot())
	}
	return buff.String()
}

// SetGrlText set the expression syntax related to this graph when it was constructed. Only ANTLR4 listener should
// call this function.
func (e *ExpressionAtom) SetGrlText(grlText string) {
	e.GrlText = grlText
}

// Evaluate will evaluate this AST graph for when scope evaluation
func (e *ExpressionAtom) Evaluate() (reflect.Value, error) {
	var val reflect.Value
	var err error
	if e.Variable != nil {
		val, err = e.Variable.Evaluate()
	} else if e.FunctionCall != nil {
		val, err = e.FunctionCall.Evaluate()
	} else if e.MethodCall != nil {
		val, err = e.MethodCall.Evaluate()
	} else if e.Constant != nil {
		val, err = e.Constant.Evaluate()
	}
	if err == nil {
		e.Value = val
	}
	return val, err
}
