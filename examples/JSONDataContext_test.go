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

type ObjectResult struct {
	Result string
}

func TestSimpleNativeType(t *testing.T) {
	oresult := &ObjectResult{
		Result: "NoResult",
	}
	dataContext := ast.NewDataContext()
	err := dataContext.Add("R", oresult)
	assert.NoError(t, err)
	err = dataContext.AddJSON("int", []byte(`1000`))
	assert.NoError(t, err)
	err = dataContext.AddJSON("str", []byte(`"A String"`))
	assert.NoError(t, err)
	err = dataContext.AddJSON("arr", []byte(`["Str1","Str2","Str3"]`))
	assert.NoError(t, err)
	err = dataContext.AddJSON("map", []byte(`{ "str" : "value", "int" : 123, "flo": 12.34, "bol": true }`))
	assert.NoError(t, err)

	DoTestIntegerJSON(t, dataContext, oresult)
	oresult.Result = "NoResult"
	DoTestStringJSON(t, dataContext, oresult)
	oresult.Result = "NoResult"
	DoTestArrayJSON(t, dataContext, oresult)
	oresult.Result = "NoResult"
	DoTestMapJSON(t, dataContext, oresult)
}

func DoTestIntegerJSON(t *testing.T, dataContext ast.IDataContext, oresult *ObjectResult) {
	rule := `
rule CheckIfJSONIntWorks {
	when R.Result == "NoResult" && int == 1000 
	then R.Result = "PERFECT";
}`

	// Prepare knowledgebase library and load it with our rule.
	lib := ast.NewKnowledgeLibrary()
	rb := builder.NewRuleBuilder(lib)
	err := rb.BuildRuleFromResource("TestJSONSimple", "0.0.1", pkg.NewBytesResource([]byte(rule)))
	assert.NoError(t, err)
	eng1 := &engine.GruleEngine{MaxCycle: 5}
	kb := lib.NewKnowledgeBaseInstance("TestJSONSimple", "0.0.1")
	err = eng1.Execute(dataContext, kb)
	assert.NoError(t, err)
	assert.Equal(t, "PERFECT", oresult.Result)
}

func DoTestStringJSON(t *testing.T, dataContext ast.IDataContext, oresult *ObjectResult) {
	rule := `
rule CheckIfJSONStringWorks {
	when R.Result == "NoResult" && str.ToUpper() == "A STRING" 
	then R.Result = "PERFECT";
}`

	// Prepare knowledgebase library and load it with our rule.
	lib := ast.NewKnowledgeLibrary()
	rb := builder.NewRuleBuilder(lib)
	err := rb.BuildRuleFromResource("TestJSONSimple", "0.0.1", pkg.NewBytesResource([]byte(rule)))
	assert.NoError(t, err)
	eng1 := &engine.GruleEngine{MaxCycle: 5}
	kb := lib.NewKnowledgeBaseInstance("TestJSONSimple", "0.0.1")
	err = eng1.Execute(dataContext, kb)
	assert.NoError(t, err)
	assert.Equal(t, "PERFECT", oresult.Result)
}

func DoTestArrayJSON(t *testing.T, dataContext ast.IDataContext, oresult *ObjectResult) {
	rule := `
rule CheckIfJSONArrayWorks {
	when R.Result == "NoResult" && arr[2].ToUpper() == "STR3" 
	then R.Result = "PERFECT";
}`

	// Prepare knowledgebase library and load it with our rule.
	lib := ast.NewKnowledgeLibrary()
	rb := builder.NewRuleBuilder(lib)
	err := rb.BuildRuleFromResource("TestJSONSimple", "0.0.1", pkg.NewBytesResource([]byte(rule)))
	assert.NoError(t, err)
	eng1 := &engine.GruleEngine{MaxCycle: 5}
	kb := lib.NewKnowledgeBaseInstance("TestJSONSimple", "0.0.1")
	err = eng1.Execute(dataContext, kb)
	assert.NoError(t, err)
	assert.Equal(t, "PERFECT", oresult.Result)
}

func DoTestMapJSON(t *testing.T, dataContext ast.IDataContext, oresult *ObjectResult) {
	rule := `
rule CheckIfJSONMapWorks {
	when R.Result == "NoResult" && map["flo"] == 12.34 
	then R.Result = "PERFECT";
}`

	// Prepare knowledgebase library and load it with our rule.
	lib := ast.NewKnowledgeLibrary()
	rb := builder.NewRuleBuilder(lib)
	err := rb.BuildRuleFromResource("TestJSONSimple", "0.0.1", pkg.NewBytesResource([]byte(rule)))
	assert.NoError(t, err)
	eng1 := &engine.GruleEngine{MaxCycle: 5}
	kb := lib.NewKnowledgeBaseInstance("TestJSONSimple", "0.0.1")
	err = eng1.Execute(dataContext, kb)
	assert.NoError(t, err)
	assert.Equal(t, "PERFECT", oresult.Result)
}

func TestComplexJSONType(t *testing.T) {
	oresult := &ObjectResult{
		Result: "NoResult",
	}
	dataContext := ast.NewDataContext()
	err := dataContext.Add("R", oresult)
	assert.NoError(t, err)
	err = dataContext.AddJSON("json", []byte(`
{
	"name" : {"first" : "Clark", "last": "Kent"},
	"company" : {
		"name" : "Daily Planet",
		"address" : "1st Super Street"
	},
	"villains" : [ "blaze", "baud", "chemo", "lex luthor", "ignition", "grax" ]
}
`))

	assert.NoError(t, err)

	rule := `
rule CheckIfJSONIntWorks {
	when 
		R.Result == "NoResult" && 
		json.name.first == "Clark" &&
		json.name["last"] == "Kent" &&
		json.villains[4].ToUpper() == "IGNITION"
	then 
		R.Result = "PERFECT";
}`

	// Prepare knowledgebase library and load it with our rule.
	lib := ast.NewKnowledgeLibrary()
	rb := builder.NewRuleBuilder(lib)
	err = rb.BuildRuleFromResource("TestJSONBitComplex", "0.0.1", pkg.NewBytesResource([]byte(rule)))
	assert.NoError(t, err)
	eng1 := &engine.GruleEngine{MaxCycle: 5}
	kb := lib.NewKnowledgeBaseInstance("TestJSONBitComplex", "0.0.1")
	err = eng1.Execute(dataContext, kb)
	assert.NoError(t, err)
	assert.Equal(t, "PERFECT", oresult.Result)
}
