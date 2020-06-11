package builder

import (
	"fmt"
	"github.com/antlr/antlr4/runtime/Go/antlr"
	antlr2 "github.com/hyperjumptech/grule-rule-engine/antlr"
	parser2 "github.com/hyperjumptech/grule-rule-engine/antlr/parser/grulev2.g4"
	"github.com/hyperjumptech/grule-rule-engine/ast"
	"github.com/hyperjumptech/grule-rule-engine/pkg"
	"github.com/juju/errors"
	log "github.com/sirupsen/logrus"
	"time"
)

// NewRuleBuilder creates new RuleBuilder instance. This builder will add all loaded rules into the specified knowledgebase.
func NewRuleBuilder(KnowledgeLibrary *ast.KnowledgeLibrary) *RuleBuilder {
	return &RuleBuilder{
		KnowledgeLibrary: KnowledgeLibrary,
	}
}

// RuleBuilder builds rule from DRL script into contained KnowledgeBase
type RuleBuilder struct {
	KnowledgeLibrary *ast.KnowledgeLibrary
}

// MustBuildRuleFromResources is similar to BuildRuleFromResources, with the difference is, it will panic if rule script contains error.
func (builder *RuleBuilder) MustBuildRuleFromResources(name, version string, resource []pkg.Resource) {
	for _, v := range resource {
		err := builder.BuildRuleFromResource(name, version, v)
		if err != nil {
			panic(err)
		}
	}
}

// MustBuildRuleFromResource is similar to BuildRuleFromResource, with the difference is, it will panic if rule script contains error.
func (builder *RuleBuilder) MustBuildRuleFromResource(name, version string, resource pkg.Resource) {
	if err := builder.BuildRuleFromResource(name, version, resource); err != nil {
		panic(err)
	}
}

// BuildRulesFromBundle will load rules from a bundle into knowledge base.
func (builder *RuleBuilder) BuildRulesFromBundle(name, version string, bundle pkg.ResouceBundle) error {
	bundles, err := bundle.Load()
	if err != nil {
		return err
	}
	return builder.BuildRuleFromResources(name, version, bundles)
}

// MustBuildRulesFromBundle is the same with BuildRulesFromBundle but it will panic if any error arises during loading resource and inserting it to knowledgebase
func (builder *RuleBuilder) MustBuildRulesFromBundle(name, version string, bundle pkg.ResouceBundle) {
	builder.MustBuildRuleFromResources(name, version, bundle.MustLoad())
}

// BuildRuleFromResources will load rules from multiple resources. It will return an error if it encounter an error on the first script it found.
func (builder *RuleBuilder) BuildRuleFromResources(name, version string, resource []pkg.Resource) error {
	for _, v := range resource {
		err := builder.BuildRuleFromResource(name, version, v)
		if err != nil {
			return errors.Trace(err)
		}
	}
	return nil
}

// BuildRuleFromResource will load rules from a single resource. It will return an error if it encounter an error on the specified resource.
func (builder *RuleBuilder) BuildRuleFromResource(name, version string, resource pkg.Resource) error {
	// save the starting time, we need to see the loading time in debug log
	startTime := time.Now()

	// Load the resource
	data, err := resource.Load()
	if err != nil {
		return errors.Trace(err)
	}
	sdata := string(data)

	// Immediately parse the loaded resource
	is := antlr.NewInputStream(sdata)
	lexer := parser2.Newgrulev2Lexer(is)
	stream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)

	var parseError error
	errCall := func(e error) {
		parseError = e
	}

	kb := builder.KnowledgeLibrary.GetKnowledgeBase(name, version)
	if kb == nil {
		return fmt.Errorf("KnowledgeBase %s:%s is not in this library", name, version)
	}

	listener := antlr2.NewGruleV2ParserListener(kb, kb.WorkingMemory, errCall)

	psr := parser2.Newgrulev2Parser(stream)
	psr.BuildParseTrees = true
	antlr.ParseTreeWalkerDefault.Walk(listener, psr.Root())

	// Get the loading duration.
	dur := time.Now().Sub(startTime)

	if parseError != nil {
		log.Errorf("Loading rule resource : %s failed. Got %v. Time take %d ms", resource.String(), parseError, dur.Milliseconds())
		return errors.Errorf("error were found before builder bailing out. Got %v", parseError)
	}

	log.Debugf("Loading rule resource : %s success. Time taken %d ms", resource.String(), dur.Milliseconds())

	return nil
}
