package model

import (
	"fmt"
	"github.com/hyperjumptech/grule-rule-engine/context"
	"github.com/hyperjumptech/grule-rule-engine/pkg"
	"github.com/juju/errors"
	"reflect"
)

// WhenScope struct hold the syntax graph for "When" expression.
type WhenScope struct {
	Expression       *Expression
	knowledgeContext *context.KnowledgeContext
	ruleCtx          *context.RuleContext
	dataCtx          *context.DataContext
}

// Initialize initialize the object graph prior execution
func (when *WhenScope) Initialize(knowledgeContext *context.KnowledgeContext, ruleCtx *context.RuleContext, dataCtx *context.DataContext) {
	when.knowledgeContext = knowledgeContext
	when.ruleCtx = ruleCtx
	when.dataCtx = dataCtx

	if when.Expression != nil {
		when.Expression.Initialize(knowledgeContext, ruleCtx, dataCtx)
	}
}

// AcceptExpression will accept any child expression underneath this Scope.
func (when *WhenScope) AcceptExpression(expression *Expression) error {
	if when.Expression != nil {
		return fmt.Errorf("expression were set twice in when scope")
	}
	when.Expression = expression
	return nil
}

// ExecuteWhen will evaluate all underneath expression.
func (when *WhenScope) ExecuteWhen() (bool, error) {
	val, err := when.Expression.Evaluate()
	if err != nil {
		return false, errors.Trace(err)
	}
	if pkg.GetBaseKind(val) != reflect.Bool {
		return false, errors.Errorf("unexpected when result... its not boolean")
	}
	return val.Bool(), nil
}
