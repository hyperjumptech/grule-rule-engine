package model

import (
	"fmt"
	"github.com/hyperjumptech/grule-rule-engine/context"
	"github.com/juju/errors"
	"reflect"
)

// FunctionCall defines function structure which defines its name and arguments.
type FunctionCall struct {
	FunctionName      string
	FunctionArguments *FunctionArgument
	knowledgeContext  *context.KnowledgeContext
	ruleCtx           *RuleContext
	dataCtx           *context.DataContext
}

// AcceptFunctionArgument configure this function call with sets of function arguments.
func (funcCall *FunctionCall) AcceptFunctionArgument(funcArg *FunctionArgument) error {
	funcCall.FunctionArguments = funcArg
	return nil
}

// Initialize will prepare this graph with context
func (funcCall *FunctionCall) Initialize(knowledgeContext *context.KnowledgeContext, ruleCtx *RuleContext, dataCtx *context.DataContext) {
	funcCall.knowledgeContext = knowledgeContext
	funcCall.ruleCtx = ruleCtx
	funcCall.dataCtx = dataCtx

	if funcCall.FunctionArguments != nil {
		funcCall.FunctionArguments.Initialize(knowledgeContext, ruleCtx, dataCtx)
	}
}

// Evaluate the object graph against underlined context or execute evaluation in the sub graph.
func (funcCall *FunctionCall) Evaluate() (reflect.Value, error) {
	var argumentValues []reflect.Value
	if funcCall.FunctionArguments == nil {
		argumentValues = make([]reflect.Value, 0)
	} else {
		av, err := funcCall.FunctionArguments.EvaluateArguments()
		if err != nil {
			return reflect.ValueOf(nil), errors.Trace(err)
		}
		argumentValues = av
	}

	fName := fmt.Sprintf("DEFUNC.%s", funcCall.FunctionName)
	return funcCall.dataCtx.ExecMethod(fName, argumentValues)
}

// EqualsTo will compare two function signature
func (funcCall *FunctionCall) EqualsTo(that AlphaNode) bool {
	typ := reflect.TypeOf(that)
	if that == nil {
		return false
	}
	if typ.Kind() == reflect.Ptr {
		if typ.Elem().Name() == "FunctionCall" {
			thatFuncCall := that.(*FunctionCall)
			return funcCall.FunctionName == thatFuncCall.FunctionName &&
				funcCall.FunctionArguments.EqualsTo(thatFuncCall.FunctionArguments)
		}
	}
	return false
}

// IsContainVariable should check for this Function call whether it contains a variable argument
func (funcCall *FunctionCall) IsContainVariable(varName string) bool {
	for _, fa := range funcCall.FunctionArguments.Arguments {
		if fa.IsContainVariable(varName) {
			return true
		}
	}
	return false
}
