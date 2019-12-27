package model

import (
	"github.com/hyperjumptech/grule-rule-engine/context"
	"github.com/juju/errors"
	"reflect"
)

// AssignExpressions contains list of assignment expression in the "then" scope.
type AssignExpressions struct {
	ExpressionList   []*AssignExpression
	knowledgeContext *context.KnowledgeContext
	ruleCtx          *RuleContext
	dataCtx          *context.DataContext
}

// Initialize will initialize this graph with context.
func (ae *AssignExpressions) Initialize(knowledgeContext *context.KnowledgeContext, ruleCtx *RuleContext, dataCtx *context.DataContext) {
	ae.knowledgeContext = knowledgeContext
	ae.ruleCtx = ruleCtx
	ae.dataCtx = dataCtx

	if ae.ExpressionList != nil {
		for _, val := range ae.ExpressionList {
			val.Initialize(knowledgeContext, ruleCtx, dataCtx)
		}
	}
}

// Evaluate the object graph against underlined context or execute evaluation in the sub graph.
func (ae *AssignExpressions) Evaluate() (reflect.Value, error) {
	for _, v := range ae.ExpressionList {
		_, err := v.Evaluate()
		if err != nil {
			return reflect.ValueOf(nil), errors.Trace(err)
		}
	}
	return reflect.ValueOf(nil), nil
}
