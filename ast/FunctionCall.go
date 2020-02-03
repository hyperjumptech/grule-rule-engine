package ast

import (
	"bytes"
	"fmt"
	"github.com/google/uuid"
	log "github.com/sirupsen/logrus"
	"reflect"
)

// NewFunctionCall creates new instance of FunctionCall
func NewFunctionCall() *FunctionCall {
	return &FunctionCall{
		AstID:        uuid.New().String(),
		ArgumentList: NewArgumentList(),
	}
}

// FunctionCall AST graph node
type FunctionCall struct {
	AstID         string
	GrlText       string
	DataContext   *DataContext
	WorkingMemory *WorkingMemory

	FunctionName string
	ArgumentList *ArgumentList
	Value        reflect.Value
}

// InitializeContext will initialize this AST graph with data context and working memory before running rule on them.
func (e *FunctionCall) InitializeContext(dataCtx *DataContext, WorkingMemory *WorkingMemory) {
	e.DataContext = dataCtx
	e.WorkingMemory = WorkingMemory
	if e.ArgumentList != nil {
		e.ArgumentList.InitializeContext(dataCtx, WorkingMemory)
	}
}

// FunctionCallReceiver should be implemented bu AST graph node to receive a FunctionCall AST graph mode
type FunctionCallReceiver interface {
	AcceptFunctionCall(fun *FunctionCall) error
}

// GetAstID get the UUID asigned for this AST graph node
func (e *FunctionCall) GetAstID() string {
	return e.AstID
}

// GetGrlText get the expression syntax related to this graph when it wast constructed
func (e *FunctionCall) GetGrlText() string {
	return e.GrlText
}

// GetSnapshot will create a structure signature or AST graph
func (e *FunctionCall) GetSnapshot() string {
	var buff bytes.Buffer
	if e.ArgumentList == nil {
		log.Errorf("Argument is nil")
	} else {
		if e.ArgumentList.Arguments == nil {
			log.Errorf("Argument.Argumeent array is nil")
		}
	}
	buff.WriteString(fmt.Sprintf("func->%s%s", e.FunctionName, e.ArgumentList.GetSnapshot()))
	return buff.String()
}

// SetGrlText set the expression syntax related to this graph when it was constructed. Only ANTLR4 listener should
// call this function.
func (e *FunctionCall) SetGrlText(grlText string) {
	e.GrlText = grlText
}

// AcceptArgumentList will accept an ArgumentList AST graph into this ast graph
func (e *FunctionCall) AcceptArgumentList(argList *ArgumentList) {
	log.Tracef("Method received argument list")
	e.ArgumentList = argList
}

// Evaluate will evaluate this AST graph for when scope evaluation
func (e *FunctionCall) Evaluate() (reflect.Value, error) {
	objName := fmt.Sprintf("DEFUNC.%s", e.FunctionName)
	args, err := e.ArgumentList.Evaluate()
	if err != nil {
		return reflect.ValueOf(nil), err
	}
	retVal, err := e.DataContext.ExecMethod(objName, args)
	if err == nil {
		e.Value = retVal
	}
	return retVal, err
}
