package builder

import (
	"github.com/antlr/antlr4/runtime/Go/antlr"
	antlr2 "github.com/hyperjumptech/grule-rule-engine/antlr"
	"github.com/hyperjumptech/grule-rule-engine/antlr/parser"
	"github.com/hyperjumptech/grule-rule-engine/model"
	"github.com/hyperjumptech/grule-rule-engine/pkg"
	"github.com/juju/errors"
	log "github.com/sirupsen/logrus"
)

// NewRuleBuilder creates new RuleBuilder instance. This builder will add all loaded rules into the specified knowledgebase.
func NewRuleBuilder(KnowledgeBase *model.KnowledgeBase) *RuleBuilder {
	return &RuleBuilder{
		KnowledgeBase: KnowledgeBase,
	}
}

// RuleBuilder builds rule from DRL script into contained KnowledgeBase
type RuleBuilder struct {
	KnowledgeBase *model.KnowledgeBase
}

// MustBuildRuleFromResources is similar to BuildRuleFromResources, with the difference is, it will panic if rule script contains error.
func (builder *RuleBuilder) MustBuildRuleFromResources(resource []pkg.Resource) {
	for _, v := range resource {
		err := builder.BuildRuleFromResource(v)
		if err != nil {
			panic(err)
		}
	}
}

// MustBuildRuleFromResource is similar to BuildRuleFromResource, with the difference is, it will panic if rule script contains error.
func (builder *RuleBuilder) MustBuildRuleFromResource(resource pkg.Resource) {
	if err := builder.BuildRuleFromResource(resource); err != nil {
		panic(err)
	}
}

// BuildRuleFromResources will load rules from multiple resources. It will return an error if it encounter an error on the first script it found.
func (builder *RuleBuilder) BuildRuleFromResources(resource []pkg.Resource) error {
	for _, v := range resource {
		err := builder.BuildRuleFromResource(v)
		if err != nil {
			return errors.Trace(err)
		}
	}
	return nil
}

// BuildRuleFromResource will load rules from a single resource. It will return an error if it encounter an error on the specified resource.
func (builder *RuleBuilder) BuildRuleFromResource(resource pkg.Resource) error {
	data, err := resource.Load()
	if err != nil {
		return errors.Trace(err)
	}
	sdata := string(data)

	is := antlr.NewInputStream(sdata)
	lexer := parser.NewgruleLexer(is)
	stream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)

	listener := antlr2.NewGruleParserListener(builder.KnowledgeBase)

	psr := parser.NewgruleParser(stream)
	psr.BuildParseTrees = true
	antlr.ParseTreeWalkerDefault.Walk(listener, psr.Root())

	if len(listener.ParseErrors) > 0 {
		log.Errorf("Loading rule resource : %s failed. Got %d errors. 1st error : %v", resource.String(), len(listener.ParseErrors), listener.ParseErrors[0])
		return errors.Errorf("error were found before builder bailing out. %d errors. 1st error : %v", len(listener.ParseErrors), listener.ParseErrors[0])
	}
	log.Debugf("Loading rule resource : %s success", resource.String())
	return nil
}
