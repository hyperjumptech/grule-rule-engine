//  Copyright hyperjumptech/grule-rule-engine Authors
//
//  Licensed under the Apache License, Version 2.0 (the "License");
//  you may not use this file except in compliance with the License.
//  You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
//  Unless required by applicable law or agreed to in writing, software
//  distributed under the License is distributed on an "AS IS" BASIS,
//  WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//  See the License for the specific language governing permissions and
//  limitations under the License.

package builder

import (
	"fmt"
	"github.com/hyperjumptech/grule-rule-engine/ast"
	"github.com/hyperjumptech/grule-rule-engine/logger"
	"github.com/sirupsen/logrus"
	"go.uber.org/zap"
	"time"

	"github.com/antlr/antlr4/runtime/Go/antlr"
	antlr2 "github.com/hyperjumptech/grule-rule-engine/antlr"
	parser "github.com/hyperjumptech/grule-rule-engine/antlr/parser/grulev3"
	"github.com/hyperjumptech/grule-rule-engine/pkg"
)

var (
	// builderLogFields default fields for grule
	builderLogFields = logger.Fields{
		"package": "builder",
	}

	// BuilderLog is a logger instance twith default fields for grule
	BuilderLog = logger.Log.WithFields(builderLogFields)
)

// SetLogger changes default logger on external
func SetLogger(log interface{}) {
	var entry logger.LogEntry

	switch log.(type) {
	case *zap.Logger:
		log, ok := log.(*zap.Logger)
		if !ok {
			return
		}
		entry = logger.NewZap(log)
	case *logrus.Logger:
		log, ok := log.(*logrus.Logger)
		if !ok {
			return
		}
		entry = logger.NewLogrus(log)
	default:
		return
	}

	BuilderLog = entry.WithFields(builderLogFields)
}

// NewRuleBuilder creates new RuleBuilder instance. This builder will add all loaded rules into the specified knowledgebase.
func NewRuleBuilder(KnowledgeLibrary *ast.KnowledgeLibrary) *RuleBuilder {
	return &RuleBuilder{
		KnowledgeLibrary: KnowledgeLibrary,
	}
}

// RuleBuilder builds rule from GRL script into contained KnowledgeBase
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
func (builder *RuleBuilder) BuildRulesFromBundle(name, version string, bundle pkg.ResourceBundle) error {
	bundles, err := bundle.Load()
	if err != nil {
		return err
	}
	return builder.BuildRuleFromResources(name, version, bundles)
}

// MustBuildRulesFromBundle is the same with BuildRulesFromBundle but it will panic if any error arises during loading resource and inserting it to knowledgebase
func (builder *RuleBuilder) MustBuildRulesFromBundle(name, version string, bundle pkg.ResourceBundle) {
	builder.MustBuildRuleFromResources(name, version, bundle.MustLoad())
}

// BuildRuleFromResources will load rules from multiple resources. It will return an error if it encounter an error on the first script it found.
func (builder *RuleBuilder) BuildRuleFromResources(name, version string, resource []pkg.Resource) error {
	for _, v := range resource {
		err := builder.BuildRuleFromResource(name, version, v)
		if err != nil {
			return err
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
		return err
	}

	// Immediately parse the loaded resource
	is := antlr.NewInputStream(string(data))
	lexer := parser.Newgrulev3Lexer(is)

	errReporter := &pkg.GruleErrorReporter{
		Errors: make([]error, 0),
	}

	lexer.RemoveErrorListeners()
	lexer.AddErrorListener(errReporter)

	stream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)

	kb := builder.KnowledgeLibrary.GetKnowledgeBase(name, version)
	if kb == nil {
		return fmt.Errorf("KnowledgeBase %s:%s is not in this library", name, version)
	}

	listener := antlr2.NewGruleV3ParserListener(kb, errReporter)

	psr := parser.Newgrulev3Parser(stream)

	psr.RemoveErrorListeners()
	psr.AddErrorListener(errReporter)

	psr.BuildParseTrees = true
	antlr.ParseTreeWalkerDefault.Walk(listener, psr.Grl())

	grl := listener.Grl
	for _, ruleEntry := range grl.RuleEntries {
		err := kb.AddRuleEntry(ruleEntry)
		if err != nil && err.Error() != "rule entry TestNoDesc already exist" {
			BuilderLog.Tracef("warning while adding rule entry : %s. got %s, possibly already added by antlr listener", ruleEntry.RuleName, err.Error())
		}
	}

	kb.WorkingMemory.IndexVariables()

	// Get the loading duration.
	dur := time.Now().Sub(startTime)

	if errReporter.HasError() {
		BuilderLog.Errorf("GRL syntax error. got %s", errReporter.Error())
		for i, errr := range errReporter.Errors {
			BuilderLog.Errorf("%d : %s", i, errr.Error())
		}
		return errReporter
	}

	BuilderLog.Debugf("Loading rule resource : %s success. Time taken %d ms", resource.String(), dur.Nanoseconds()/1e6)

	return nil
}
