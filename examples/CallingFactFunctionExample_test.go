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

package examples

import (
	"fmt"
	"github.com/hyperjumptech/grule-rule-engine/ast"
	"github.com/hyperjumptech/grule-rule-engine/builder"
	"github.com/hyperjumptech/grule-rule-engine/engine"
	"github.com/hyperjumptech/grule-rule-engine/pkg"
	"github.com/stretchr/testify/assert"
	"testing"
)

const (
	CallFactFuncGRL = `
rule RaiseHandToClap "Raise the hand to be able to clap" {
	when 
		Clapper.ClapCount < 10 &&
		Clapper.HandsUp == false
	then
		Clapper.HandsUp = true;
		Log("RaiseHandToClap : Now hands up");
}

rule CheckIfCanClap "If hands is up we can clap" {
	when
		Clapper.HandsUp &&
		Clapper.CanClap == false
	then
		Clapper.CanClap = true;
		Log("CheckIfCanClap : Now can clap");
}

rule LetsClap "If hands are up and can clap, lets clap" {
	when
		Clapper.HandsUp && Clapper.CanClap
	then
		Log("LetsClap : Now clapping. Hands Down thus Can't clap'");
		Clapper.Clap();

		// instruct the engine to forget about variables and functions. Otherwise the engine will use the last remembered value.
		Forget("Clapper.CanClap");
		Forget("Clapper.HandsUp");
		Forget("Clapper.ClapCount");
		Forget("Clapper.Clap()");
}
`
)

type Clapper struct {
	CanClap   bool
	HandsUp   bool
	ClapCount int64
}

func (c *Clapper) Clap() {
	c.ClapCount++
	c.HandsUp = false
	c.CanClap = false
}

func TestCallingFactFunction(t *testing.T) {
	c := &Clapper{
		CanClap:   false,
		HandsUp:   false,
		ClapCount: 0,
	}

	dataContext := ast.NewDataContext()
	err := dataContext.Add("Clapper", c)
	assert.NoError(t, err)

	// Prepare knowledgebase library and load it with our rule.
	lib := ast.NewKnowledgeLibrary()
	ruleBuilder := builder.NewRuleBuilder(lib)
	err = ruleBuilder.BuildRuleFromResource("CallingFactFunction", "0.1.1", pkg.NewBytesResource([]byte(CallFactFuncGRL)))
	knowledgeBase := lib.NewKnowledgeBaseInstance("CallingFactFunction", "0.1.1")
	assert.NoError(t, err)
	eng1 := &engine.GruleEngine{MaxCycle: 500}
	err = eng1.Execute(dataContext, knowledgeBase)
	assert.NoError(t, err)
	fmt.Printf("%v", c)
	assert.Equal(t, 10, int(c.ClapCount))
}
