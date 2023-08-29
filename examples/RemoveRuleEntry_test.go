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
	"github.com/hyperjumptech/grule-rule-engine/ast"
	"github.com/hyperjumptech/grule-rule-engine/builder"
	"github.com/hyperjumptech/grule-rule-engine/engine"
	"github.com/hyperjumptech/grule-rule-engine/pkg"
	"github.com/stretchr/testify/assert"
	"testing"
)

const (
	RuleA = `
rule ColorCheck1 "test 1" salience 100 {
	when
	  Color.Name.In("Grey", "Black")
	then
	  Color.Message = "Its Grey!!!";
	  Retract("ColorCheck1");
}
`
	RuleB = `
rule ColorCheck1 "test 2"  {
	when
	  Color.Name.In("Black")
	then
	  Color.Message = "Its Black!!!";
	  Retract("ColorCheck1");
}
`
)

type ColorCheck struct {
	Name    string
	Message string
}

func TestRemoveRuleEntry(t *testing.T) {
	color := &ColorCheck{
		Name:    "Black",
		Message: "",
	}

	dataContext := ast.NewDataContext()
	err := dataContext.Add("Color", color)
	assert.NoError(t, err)

	//Load RuleA into knowledgeBase
	lib := ast.NewKnowledgeLibrary()
	rb := builder.NewRuleBuilder(lib)
	err = rb.BuildRuleFromResource("Test", "0.1.1", pkg.NewBytesResource([]byte(RuleA)))
	assert.NoError(t, err)
	kb, err := lib.NewKnowledgeBaseInstance("Test", "0.1.1")
	assert.NoError(t, err)
	eng := &engine.GruleEngine{MaxCycle: 1}
	err = eng.Execute(dataContext, kb)
	assert.NoError(t, err)
	assert.Equal(t, "Its Grey!!!", color.Message)

	//Remove RuleEntry A
	kb.RemoveRuleEntry("ColorCheck1")
	lib.RemoveRuleEntry("ColorCheck1", "Test", "0.1.1")

	//Add RuleB again, which is similar to RuleA except its output
	err = rb.BuildRuleFromResource("Test", "0.1.1", pkg.NewBytesResource([]byte(RuleB)))
	assert.NoError(t, err)
	kb, err = lib.NewKnowledgeBaseInstance("Test", "0.1.1")
	assert.NoError(t, err)
	eng = &engine.GruleEngine{MaxCycle: 1}
	err = eng.Execute(dataContext, kb)
	assert.NoError(t, err)
	assert.Equal(t, "Its Black!!!", color.Message)
}
