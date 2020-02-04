package examples

import (
	"github.com/hyperjumptech/grule-rule-engine/ast"
	"github.com/hyperjumptech/grule-rule-engine/builder"
	"github.com/hyperjumptech/grule-rule-engine/engine"
	"github.com/hyperjumptech/grule-rule-engine/pkg"
	"github.com/sirupsen/logrus"
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
	logrus.SetLevel(logrus.DebugLevel)
	dataContext := ast.NewDataContext()
	memory := ast.NewWorkingMemory()
	knowledgeBase := ast.NewKnowledgeBase("CallingLog", "0.1.1")
	ruleBuilder := builder.NewRuleBuilder(knowledgeBase, memory)

	err := ruleBuilder.BuildRuleFromResource(pkg.NewBytesResource([]byte(DRL)))
	if err != nil {
		panic(err)
	} else {
		eng1 := &engine.GruleEngine{MaxCycle: 1}
		err := eng1.Execute(dataContext, knowledgeBase, memory)
		if err != nil {
			t.Fatalf("Got error %v", err)
		}
	}
}
