package model

import (
	"github.com/hyperjumptech/grule-rule-engine/context"
	"reflect"
)

// Constant holds a constants, it holds a simple golang value.
type Constant struct {
	ConstantValue    reflect.Value
	knowledgeContext *context.KnowledgeContext
	ruleCtx          *context.RuleContext
	dataCtx          *context.DataContext
}

// Initialize will initialize this graph with context
func (cons *Constant) Initialize(knowledgeContext *context.KnowledgeContext, ruleCtx *context.RuleContext, dataCtx *context.DataContext) {
	cons.knowledgeContext = knowledgeContext
	cons.ruleCtx = ruleCtx
	cons.dataCtx = dataCtx
}

// Evaluate the object graph against underlined context or execute evaluation in the sub graph.
func (cons *Constant) Evaluate() (reflect.Value, error) {
	return cons.ConstantValue, nil
}

// AcceptDecimal prepare this graph with a decimal value.
func (cons *Constant) AcceptDecimal(val int64) error {
	cons.ConstantValue = reflect.ValueOf(val)
	return nil
}
