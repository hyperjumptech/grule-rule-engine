package model

import (
	"github.com/hyperjumptech/grule-rule-engine/context"
	"github.com/juju/errors"
)

// RuleEntry represent the language graph of a single rule entry.
type RuleEntry struct {
	Salience         int64
	RuleName         string
	RuleDescription  string
	WhenScope        *WhenScope
	ThenScope        *ThenScope
	knowledgeContext *context.KnowledgeContext
	ruleCtx          *RuleContext
	dataCtx          *context.DataContext

	Retracted bool
}

// AcceptDecimal will store salience information.
func (entry *RuleEntry) AcceptDecimal(val int64) error {
	entry.Salience = val
	return nil
}

// Initialize will init this graph prior execution.
func (entry *RuleEntry) Initialize(knowledgeContext *context.KnowledgeContext, ruleCtx *RuleContext, dataCtx *context.DataContext) {
	entry.knowledgeContext = knowledgeContext
	entry.ruleCtx = ruleCtx
	entry.dataCtx = dataCtx

	if entry.WhenScope != nil {
		entry.WhenScope.Initialize(knowledgeContext, ruleCtx, dataCtx)
	}

	if entry.ThenScope != nil {
		entry.ThenScope.Initialize(knowledgeContext, ruleCtx, dataCtx)
	}
}

// CanExecute Test whether this rule entry are eligible for execution by the rule engine with the underlying data.
func (entry *RuleEntry) CanExecute() (bool, error) {
	if entry.Retracted {
		return false, nil
	}
	bol, err := entry.WhenScope.ExecuteWhen()
	if err != nil {
		return false, errors.Trace(err)
	}
	return bol, nil
}

// Execute will execute the action part of the rule entry.
func (entry *RuleEntry) Execute() error {
	println("Execute rule", entry.RuleName)
	return entry.ThenScope.Execute()
}
