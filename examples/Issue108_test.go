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

type Struct108 struct {
	Rule1Done bool
	Rule2Done bool
	Rule3Done bool
	Sequence  []string
}

var Rule108 = `
rule Conflicting1 "First conflicting rule" salience 1 {
	when
		F.Rule1Done == false 
	then
		F.Rule1Done = true;
		F.Sequence.Append("1");
}

rule Conflicting2 "Second conflicting rule" salience 3 {
	when
		F.Rule2Done == false 
	then
		F.Rule2Done = true;
		F.Sequence.Append("2");
}

rule Conflicting3 "Third conflicting rule" salience 2 {
	when
		F.Rule3Done == false 
	then
		F.Rule3Done = true;
		F.Sequence.Append("3");
}
`

func TestIssue108(t *testing.T) {
	Obj := &Struct108{
		Rule1Done: false,
		Rule2Done: false,
		Rule3Done: false,
		Sequence:  make([]string, 0),
	}

	dataContext := ast.NewDataContext()
	err := dataContext.Add("F", Obj)
	assert.NoError(t, err)

	// Prepare knowledgebase library and load it with our rule.
	lib := ast.NewKnowledgeLibrary()
	rb := builder.NewRuleBuilder(lib)
	err = rb.BuildRuleFromResource("Test108", "0.0.1", pkg.NewBytesResource([]byte(Rule108)))
	assert.NoError(t, err)
	eng1 := &engine.GruleEngine{MaxCycle: 5}
	kb := lib.NewKnowledgeBaseInstance("Test108", "0.0.1")
	err = eng1.Execute(dataContext, kb)
	assert.NoError(t, err)
	assert.Equal(t, 3, len(Obj.Sequence))
	assert.Equal(t, "2", Obj.Sequence[0])
	assert.Equal(t, "3", Obj.Sequence[1])
	assert.Equal(t, "1", Obj.Sequence[2])
}
