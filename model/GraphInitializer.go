package model

import "github.com/hyperjumptech/grule-rule-engine/context"

// GraphInitializer defines all graph that can be initalized with context.
type GraphInitializer interface {
	Initialize(knowledgeContext *context.KnowledgeContext, ruleCtx *RuleContext, dataCtx *context.DataContext)
}
