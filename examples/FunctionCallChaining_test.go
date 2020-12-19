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

type TestTree struct {
	Name  string
	Child *TestTree
}

func (tt *TestTree) GetChild() *TestTree {
	return tt.Child
}

func (tt *TestTree) CallMyName() {
	fmt.Println("Calling your name", tt.Name)
}

func TestFunctionChain(t *testing.T) {
	Tree := &TestTree{
		Name: "Test",
		Child: &TestTree{
			Name: "TestTest",
			Child: &TestTree{
				Name: "TestTestTest",
				Child: &TestTree{
					Name:  "TestTestTestTest",
					Child: nil,
				},
			},
		},
	}

	rule := `
rule SetTreeName "Set the top most tree name" {
	when
		Tree.Name.ToUpper() == "TEST" &&
		Tree.GetChild().Name.ToUpper() == "TESTTEST" &&
		Tree.GetChild().GetChild().Name.ToUpper() == "TESTTESTTEST" &&
		Tree.GetChild().GetChild().GetChild().Name.ToLower() == "   testtesttesttest  ".Trim()
	then
		Tree.Name = "VERIFIED".ToLower();
		Tree.GetChild().GetChild().CallMyName();
		Retract("SetTreeName");
}
`

	dataContext := ast.NewDataContext()
	err := dataContext.Add("Tree", Tree)
	assert.NoError(t, err)

	lib := ast.NewKnowledgeLibrary()
	ruleBuilder := builder.NewRuleBuilder(lib)
	err = ruleBuilder.BuildRuleFromResource("TestFuncChaining", "0.0.1", pkg.NewBytesResource([]byte(rule)))
	assert.NoError(t, err)
	kb := lib.NewKnowledgeBaseInstance("TestFuncChaining", "0.0.1")
	eng1 := &engine.GruleEngine{MaxCycle: 1}
	err = eng1.Execute(dataContext, kb)
	assert.NoError(t, err)
	assert.Equal(t, "verified", Tree.Name)

}

func TestVariableChain(t *testing.T) {
	Tree := &TestTree{
		Name: "Test",
		Child: &TestTree{
			Name: "TestTest",
			Child: &TestTree{
				Name: "TestTestTest",
				Child: &TestTree{
					Name:  "TestTestTestTest",
					Child: nil,
				},
			},
		},
	}

	rule := `
rule SetTreeName "Set the top most tree name" {
	when
		Tree.Name.ToUpper() == "TEST" &&
		Tree.Child.Name.ToUpper() == "TESTTEST" &&
		Tree.Child.Child.Name.ToUpper() == "TESTTESTTEST" &&
		Tree.Child.Child.Child.Name.ToLower() == "   testtesttesttest  ".Trim()
	then
		Tree.Name = "VERIFIED".ToLower();
		Tree.Child.Child.Child.Name = "SUCCESS";
		Retract("SetTreeName");
}
`

	dataContext := ast.NewDataContext()
	err := dataContext.Add("Tree", Tree)
	assert.NoError(t, err)

	lib := ast.NewKnowledgeLibrary()
	ruleBuilder := builder.NewRuleBuilder(lib)
	err = ruleBuilder.BuildRuleFromResource("TestFuncChaining", "0.0.1", pkg.NewBytesResource([]byte(rule)))
	assert.NoError(t, err)
	kb := lib.NewKnowledgeBaseInstance("TestFuncChaining", "0.0.1")
	eng1 := &engine.GruleEngine{MaxCycle: 1}
	err = eng1.Execute(dataContext, kb)
	assert.NoError(t, err)
	assert.Equal(t, "verified", Tree.Name)
	assert.Equal(t, "SUCCESS", Tree.Child.Child.Child.Name)
}
