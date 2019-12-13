package model

import (
	"github.com/hyperjumptech/grule-rule-engine/context"
	"github.com/juju/errors"
	"reflect"
)

// FunctionArgument stores set of argument within a function call.
type FunctionArgument struct {
	Arguments        []*ArgumentHolder
	knowledgeContext *context.KnowledgeContext
	ruleCtx          *context.RuleContext
	dataCtx          *context.DataContext
}

// EvaluateArguments the object graph against underlined context or execute evaluation in the sub graph.
func (funcArg *FunctionArgument) EvaluateArguments() ([]reflect.Value, error) {
	if funcArg.Arguments == nil || len(funcArg.Arguments) == 0 {
		return make([]reflect.Value, 0), nil
	}
	retVal := make([]reflect.Value, len(funcArg.Arguments))
	for i, v := range funcArg.Arguments {
		rv, err := v.Evaluate()
		if err != nil {
			return retVal, errors.Trace(err)
		}
		retVal[i] = rv
	}
	return retVal, nil
}

// Initialize will prepare this set of arguments with contexts.
func (funcArg *FunctionArgument) Initialize(knowledgeContext *context.KnowledgeContext, ruleCtx *context.RuleContext, dataCtx *context.DataContext) {
	funcArg.knowledgeContext = knowledgeContext
	funcArg.ruleCtx = ruleCtx
	funcArg.dataCtx = dataCtx

	if funcArg.Arguments != nil {
		for _, val := range funcArg.Arguments {
			val.Initialize(knowledgeContext, ruleCtx, dataCtx)
		}
	}
}

// AcceptExpression add an expression into function arguments.
func (funcArg *FunctionArgument) AcceptExpression(expression *Expression) error {
	holder := &ArgumentHolder{
		Expression: expression,
	}
	funcArg.Arguments = append(funcArg.Arguments, holder)
	return nil
}

// AcceptFunctionCall add a function call into function arguments.
func (funcArg *FunctionArgument) AcceptFunctionCall(funcCall *FunctionCall) error {
	holder := &ArgumentHolder{
		FunctionCall: funcCall,
	}
	funcArg.Arguments = append(funcArg.Arguments, holder)
	return nil
}

// AcceptMethodCall add a method call into function argument.
func (funcArg *FunctionArgument) AcceptMethodCall(methodCall *MethodCall) error {
	holder := &ArgumentHolder{
		MethodCall: methodCall,
	}
	funcArg.Arguments = append(funcArg.Arguments, holder)
	return nil
}

// AcceptVariable add a variable into function argument.
func (funcArg *FunctionArgument) AcceptVariable(name string) error {
	holder := &ArgumentHolder{
		Variable: name,
	}
	funcArg.Arguments = append(funcArg.Arguments, holder)
	return nil
}

// AcceptConstant add a constant into function argument.
func (funcArg *FunctionArgument) AcceptConstant(cons *Constant) error {
	holder := &ArgumentHolder{
		Constant: cons,
	}
	funcArg.Arguments = append(funcArg.Arguments, holder)
	return nil
}
