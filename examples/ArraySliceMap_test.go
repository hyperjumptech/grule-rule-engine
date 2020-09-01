package examples

import (
	"github.com/hyperjumptech/grule-rule-engine/ast"
	"github.com/hyperjumptech/grule-rule-engine/builder"
	"github.com/hyperjumptech/grule-rule-engine/engine"
	"github.com/hyperjumptech/grule-rule-engine/pkg"
	"github.com/stretchr/testify/assert"
	"testing"
)

type ArrayNode struct {
	Name        string
	StringArray []string
	NumberArray []int
	ChildArray  []*ArrayNode
}

func TestArraySlice(t *testing.T) {
	Tree := &ArrayNode{
		Name:        "Node",
		StringArray: []string{"NodeString1", "NodeString2"},
		NumberArray: []int{235, 633},
		ChildArray: []*ArrayNode{
			&ArrayNode{
				Name:        "NodeChild1",
				StringArray: []string{"NodeChildString11", "NodeChildString12"},
				NumberArray: []int{578, 296},
				ChildArray:  nil,
			}, &ArrayNode{
				Name:        "NodeChild2",
				StringArray: []string{"NodeChildString21", "NodeChildString22"},
				NumberArray: []int{744, 895},
				ChildArray:  nil,
			},
		},
	}

	rule := `
rule SetTreeName "Set the top most tree name" {
	when
		Tree.Name.ToUpper() == "NODE" &&
		Tree.StringArray[0].ToUpper() == "NODESTRING1" &&
		Tree.StringArray[1].ToLower() == "nodestring2" &&
		Tree.NumberArray[0] == 235 &&
		Tree.NumberArray[1] == 633 &&
		Tree.ChildArray[0].Name == "NodeChild1" &&
		Tree.ChildArray[0].StringArray[1] == "NodeChildString12"
	then
		Tree.Name = "VERIFIED".ToLower();
		Tree.ChildArray[0].StringArray[0] = "SetSuccessful";
		Tree.NumberArray[1] = 1000;
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
	assert.Equal(t, "SetSuccessful", Tree.ChildArray[0].StringArray[0])
	assert.Equal(t, 1000, Tree.NumberArray[1])
}
