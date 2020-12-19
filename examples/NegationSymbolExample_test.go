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

type StructTest struct {
	StringValue string
	BoolValue   bool
}

const (
	NegationRule1 = `
rule  NegationCheck  "User Related Rule"  salience 10 {
	when 
		!(StructTest.StringValue != "ABC")
	Then
		StructTest.StringValue="ITS ABC";
		Retract("NegationCheck");
}
`

	NegationRule2 = `
rule  NegationCheck  "User Related Rule"  salience 10 {
	when 
		!StructTest.BoolValue
	Then
		StructTest.StringValue="YES ITS NOT";
		Retract("NegationCheck");
}
`
)

func TestNegationSymbolExpressionAtom(t *testing.T) {
	structTest := &StructTest{
		BoolValue: false,
	}

	dataContext := ast.NewDataContext()
	err := dataContext.Add("StructTest", structTest)
	assert.NoError(t, err)

	// Prepare knowledgebase library and load it with our rule.
	lib := ast.NewKnowledgeLibrary()
	rb := builder.NewRuleBuilder(lib)
	err = rb.BuildRuleFromResource("TestNegation", "1.0.0", pkg.NewBytesResource([]byte(NegationRule2)))
	assert.NoError(t, err)
	eng1 := &engine.GruleEngine{MaxCycle: 5}
	kb := lib.NewKnowledgeBaseInstance("TestNegation", "1.0.0")
	err = eng1.Execute(dataContext, kb)
	assert.NoError(t, err)
	assert.Equal(t, "YES ITS NOT", structTest.StringValue)
}

func TestNegationSymbolExpression(t *testing.T) {
	structTest := &StructTest{
		StringValue: "ABC",
	}

	dataContext := ast.NewDataContext()
	err := dataContext.Add("StructTest", structTest)
	assert.NoError(t, err)

	// Prepare knowledgebase library and load it with our rule.
	lib := ast.NewKnowledgeLibrary()
	rb := builder.NewRuleBuilder(lib)
	err = rb.BuildRuleFromResource("TestNegation", "1.0.0", pkg.NewBytesResource([]byte(NegationRule1)))
	assert.NoError(t, err)
	eng1 := &engine.GruleEngine{MaxCycle: 5}
	kb := lib.NewKnowledgeBaseInstance("TestNegation", "1.0.0")
	err = eng1.Execute(dataContext, kb)
	assert.NoError(t, err)
	assert.Equal(t, "ITS ABC", structTest.StringValue)
}
