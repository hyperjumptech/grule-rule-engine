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
	Rule1 = `
rule ColorCheck "test 1"  {
	when
	  Color.Name.In("Black", "Yellow")
	then
	  Color.Message = "Its either black or yellow!!!";
	  Retract("ColorCheck");
}
`
)

type Color struct {
	Name string
	Message string
}

func TestStringInExample(t *testing.T) {
	color := &Color{
		Name:    "Black",
		Message: "",
	}

	dataContext := ast.NewDataContext()
	err := dataContext.Add("Color", color)
	assert.NoError(t, err)

	// Prepare knowledgebase library and load it with our rule.
	lib := ast.NewKnowledgeLibrary()
	rb := builder.NewRuleBuilder(lib)
	err = rb.BuildRuleFromResource("Test", "0.1.1", pkg.NewBytesResource([]byte(Rule1)))
	assert.NoError(t, err)
	kb := lib.NewKnowledgeBaseInstance("Test", "0.1.1")
	eng1 := &engine.GruleEngine{MaxCycle: 1}
	err = eng1.Execute(dataContext, kb)
	assert.NoError(t, err)
	assert.Equal(t, "Its either black or yellow!!!", color.Message)
}
