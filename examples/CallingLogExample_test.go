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
	DRL = `
rule CallingLog "Calling a log" {
	when
		true
	then
		Log("Hello Grule");
		Retract("CallingLog");
}
`
)

func TestCallingLog(t *testing.T) {
	dataContext := ast.NewDataContext()

	lib := ast.NewKnowledgeLibrary()
	ruleBuilder := builder.NewRuleBuilder(lib)
	err := ruleBuilder.BuildRuleFromResource("CallingLog", "0.1.1", pkg.NewBytesResource([]byte(DRL)))
	assert.NoError(t, err)

	knowledgeBase := lib.NewKnowledgeBaseInstance("CallingLog", "0.1.1")

	eng1 := &engine.GruleEngine{MaxCycle: 1}
	err = eng1.Execute(dataContext, knowledgeBase)
	assert.NoError(t, err)
}
