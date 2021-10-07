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
rule ColorCheck2 "test 2"  {
	when
	  Color.Name.In("Black")
	then
	  Color.Message = "Its Black!!!";
	  Retract("ColorCheck2");
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
	kb := lib.NewKnowledgeBaseInstance("Test", "0.1.1")
	eng := &engine.GruleEngine{MaxCycle: 1}
	err = eng.Execute(dataContext, kb)
	assert.NoError(t, err)
	assert.Equal(t, "Its Grey!!!", color.Message)

	//Remove RuleEntry A and add RuleB to correctly point the color
	kb.RemoveRuleEntry("ColorCheck1")
	lib.RemoveRuleEntry("ColorCheck1", "Test", "0.1.1")
	err = rb.BuildRuleFromResource("Test", "0.1.1", pkg.NewBytesResource([]byte(RuleB)))
	assert.NoError(t, err)
	kb = lib.NewKnowledgeBaseInstance("Test", "0.1.1")
	eng = &engine.GruleEngine{MaxCycle: 3}
	err = eng.Execute(dataContext, kb)
	assert.NoError(t, err)
	// THIS IS FAILING
	assert.True(t, kb.RuleEntries["Deleted_ColorCheck1"] != nil )
	assert.Equal(t, "Its Black!!!", color.Message)
}