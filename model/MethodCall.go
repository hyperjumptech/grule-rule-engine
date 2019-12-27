package model

import (
	"github.com/hyperjumptech/grule-rule-engine/context"
	"github.com/juju/errors"
	"reflect"
)

// MethodCall defines a graph struct that form a method call. It holds the method name to call and the arguments.
type MethodCall struct {
	MethodName       string
	MethodArguments  *FunctionArgument
	knowledgeContext *context.KnowledgeContext
	ruleCtx          *RuleContext
	dataCtx          *context.DataContext
}

// Initialize will initialize this graph with context.
func (methCall *MethodCall) Initialize(knowledgeContext *context.KnowledgeContext, ruleCtx *RuleContext, dataCtx *context.DataContext) {
	methCall.knowledgeContext = knowledgeContext
	methCall.ruleCtx = ruleCtx
	methCall.dataCtx = dataCtx

	if methCall.MethodArguments != nil {
		methCall.MethodArguments.Initialize(knowledgeContext, ruleCtx, dataCtx)
	}
}

// AcceptFunctionArgument will prepare this graph with the function arguments.
func (methCall *MethodCall) AcceptFunctionArgument(funcArg *FunctionArgument) error {
	methCall.MethodArguments = funcArg
	return nil
}

// Evaluate the object graph against underlined context or execute evaluation in the sub graph.
func (methCall *MethodCall) Evaluate() (reflect.Value, error) {
	var argumentValues []reflect.Value
	if methCall.MethodArguments == nil {
		argumentValues = make([]reflect.Value, 0)
	} else {
		av, err := methCall.MethodArguments.EvaluateArguments()
		if err != nil {
			return reflect.ValueOf(nil), errors.Trace(err)
		}
		argumentValues = av
	}

	return methCall.dataCtx.ExecMethod(methCall.MethodName, argumentValues)
}

// EqualsTo will compare two method signature
func (methCall *MethodCall) EqualsTo(that AlphaNode) bool {
	typ := reflect.TypeOf(that)
	if that == nil {
		return false
	}
	if typ.Kind() == reflect.Ptr {
		if typ.Elem().Name() == "MethodCall" {
			thatFuncCall := that.(*MethodCall)
			return methCall.MethodName == thatFuncCall.MethodName &&
				methCall.MethodArguments.EqualsTo(thatFuncCall.MethodArguments)
		}
	}
	return false
}
