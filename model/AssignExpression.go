package model

import (
	"github.com/hyperjumptech/grule-rule-engine/context"
	"github.com/juju/errors"
	"reflect"
)

// AssignExpression an expression for assignment, used to assign a variable with some function, constants or method  all or simply calling function.
type AssignExpression struct {
	Assignment       *Assignment
	FunctionCall     *FunctionCall
	MethodCall       *MethodCall
	knowledgeContext *context.KnowledgeContext
	ruleCtx          *context.RuleContext
	dataCtx          *context.DataContext
}

// Initialize will initiate this graph with context.
func (ae *AssignExpression) Initialize(knowledgeContext *context.KnowledgeContext, ruleCtx *context.RuleContext, dataCtx *context.DataContext) {
	ae.knowledgeContext = knowledgeContext
	ae.ruleCtx = ruleCtx
	ae.dataCtx = dataCtx

	if ae.Assignment != nil {
		ae.Assignment.Initialize(knowledgeContext, ruleCtx, dataCtx)
	}

	if ae.FunctionCall != nil {
		ae.FunctionCall.Initialize(knowledgeContext, ruleCtx, dataCtx)
	}

	if ae.MethodCall != nil {
		ae.MethodCall.Initialize(knowledgeContext, ruleCtx, dataCtx)
	}
}

// AcceptFunctionCall prepare this graph for function call.
func (ae *AssignExpression) AcceptFunctionCall(funcCall *FunctionCall) error {
	ae.FunctionCall = funcCall
	return nil
}

// AcceptMethodCall prepare this graph for method  all
func (ae *AssignExpression) AcceptMethodCall(methodCall *MethodCall) error {
	ae.MethodCall = methodCall
	return nil
}

// Evaluate the object graph against underlined context or execute evaluation in the sub graph.
func (ae *AssignExpression) Evaluate() (reflect.Value, error) {
	if ae.Assignment != nil {
		return ae.Assignment.Evaluate()
	}
	if ae.FunctionCall != nil {
		return ae.FunctionCall.Evaluate()
	}
	if ae.MethodCall != nil {
		return ae.MethodCall.Evaluate()
	}
	return reflect.ValueOf(nil), errors.Errorf("no assignment, function or method call to evaluate")

}
