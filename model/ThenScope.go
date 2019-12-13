package model

import (
	"github.com/hyperjumptech/grule-rule-engine/context"
	"github.com/juju/errors"
)

// ThenScope holds the language graph for then expressions.
type ThenScope struct {
	AssignExpressions *AssignExpressions
	knowledgeContext  *context.KnowledgeContext
	ruleCtx           *context.RuleContext
	dataCtx           *context.DataContext
}

// Initialize will init this object graph prior execution
func (then *ThenScope) Initialize(knowledgeContext *context.KnowledgeContext, ruleCtx *context.RuleContext, dataCtx *context.DataContext) {
	then.knowledgeContext = knowledgeContext
	then.ruleCtx = ruleCtx

	if then.AssignExpressions != nil {
		then.AssignExpressions.Initialize(knowledgeContext, ruleCtx, dataCtx)
	}
}

// Execute this graph against underlying facts.
func (then *ThenScope) Execute() error {
	_, err := then.AssignExpressions.Evaluate()
	if err != nil {
		return errors.Trace(err)
	}
	return nil
}
