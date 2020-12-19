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
	"time"

	"github.com/antlr/antlr4/runtime/Go/antlr"
	antlr2 "github.com/hyperjumptech/grule-rule-engine/antlr"
	parser "github.com/hyperjumptech/grule-rule-engine/antlr/parser/grulev3"
	"github.com/hyperjumptech/grule-rule-engine/pkg"
)

var (
	// BuilderLog is a logrus instance twith default fields for grule
	BuilderLog = logger.Log.WithFields(logrus.Fields{
		"package": "builder",
	})
)

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
	stream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)

	var parseError error
	errCall := func(e error) {
		parseError = e
	}

	kb := builder.KnowledgeLibrary.GetKnowledgeBase(name, version)
	if kb == nil {
		return fmt.Errorf("KnowledgeBase %s:%s is not in this library", name, version)
	}

	listener := antlr2.NewGruleV3ParserListener(kb, errCall)

	psr := parser.Newgrulev3Parser(stream)
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

	if parseError != nil {
		BuilderLog.Errorf("Loading rule resource : %s failed. Got %v. Time take %d ms", resource.String(), parseError, dur.Nanoseconds()/1e6)
		return fmt.Errorf("error were found before builder bailing out. Got %v", parseError)
	}

	BuilderLog.Debugf("Loading rule resource : %s success. Time taken %d ms", resource.String(), dur.Nanoseconds()/1e6)

	return nil
}
