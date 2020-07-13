package ast

import (
	"bytes"
	"fmt"
	"reflect"

	"github.com/google/uuid"
	"github.com/hyperjumptech/grule-rule-engine/pkg"
	log "github.com/sirupsen/logrus"
)

// NewMethodCall create new instance of MethodCall
func NewMethodCall() *MethodCall {
	return &MethodCall{
		AstID:        uuid.New().String(),
		ArgumentList: NewArgumentList(),
	}
}

// MethodCall AST node graph
type MethodCall struct {
	AstID         string
	GrlText       string
	DataContext   IDataContext
	WorkingMemory *WorkingMemory

	MethodName   string
	ArgumentList *ArgumentList

	Value reflect.Value
}

// Clone will clone this MethodCall. The new clone will have an identical structure
func (e MethodCall) Clone(cloneTable *pkg.CloneTable) *MethodCall {
	clone := &MethodCall{
		AstID:         uuid.New().String(),
		GrlText:       e.GrlText,
		DataContext:   nil,
		WorkingMemory: nil,
		MethodName:    e.MethodName,
	}

	if e.ArgumentList != nil {
		if cloneTable.IsCloned(e.ArgumentList.AstID) {
			clone.ArgumentList = cloneTable.Records[e.ArgumentList.AstID].CloneInstance.(*ArgumentList)
		} else {
			cloned := e.ArgumentList.Clone(cloneTable)
			clone.ArgumentList = cloned
			cloneTable.MarkCloned(e.ArgumentList.AstID, cloned.AstID, e.ArgumentList, cloned)
		}
	}
	return clone
}

// InitializeContext will initialize this AST graph with data context and working memory before running rule on them.
func (e *MethodCall) InitializeContext(dataCtx IDataContext, WorkingMemory *WorkingMemory) {
	e.DataContext = dataCtx
	e.WorkingMemory = WorkingMemory
	if e.ArgumentList != nil {
		e.ArgumentList.InitializeContext(dataCtx, WorkingMemory)
	}
}

// MethodCallReceiver should be implemented by AST graph node to receive MethodCall AST graph node
type MethodCallReceiver interface {
	AcceptMethodCall(fun *MethodCall) error
}

// GetAstID get the UUID asigned for this AST graph node
func (e *MethodCall) GetAstID() string {
	return e.AstID
}

// GetGrlText get the expression syntax related to this graph when it wast constructed
func (e *MethodCall) GetGrlText() string {
	return e.GrlText
}

// GetSnapshot will create a structure signature or AST graph
func (e *MethodCall) GetSnapshot() string {
	var buff bytes.Buffer
	if e.ArgumentList == nil {
		log.Errorf("Argument is nil")
	} else {
		if e.ArgumentList.Arguments == nil {
			log.Errorf("Argument.Argumeent array is nil")
		}
	}
	buff.WriteString(fmt.Sprintf("meth->%s%s", e.MethodName, e.ArgumentList.GetSnapshot()))
	return buff.String()
}

// SetGrlText set the expression syntax related to this graph when it was constructed. Only ANTLR4 listener should
// call this function.
func (e *MethodCall) SetGrlText(grlText string) {
	e.GrlText = grlText
}

// AcceptArgumentList will accept an ArgumentList AST graph into this ast graph
func (e *MethodCall) AcceptArgumentList(argList *ArgumentList) {
	log.Tracef("Method received argument list")
	e.ArgumentList = argList
}

// Evaluate will evaluate this AST graph for when scope evaluation
func (e *MethodCall) Evaluate() (reflect.Value, error) {
	args, err := e.ArgumentList.Evaluate()
	if err != nil {
		return reflect.ValueOf(nil), err
	}
	log.Tracef("Calling method %s", e.MethodName)
	retVal, err := e.DataContext.ExecMethod(e.MethodName, args)
	if err == nil {
		e.Value = retVal
	}
	return retVal, err
}
