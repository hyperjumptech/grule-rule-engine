package examples

import (
	"github.com/hyperjumptech/grule-rule-engine/ast"
	"github.com/hyperjumptech/grule-rule-engine/builder"
	"github.com/hyperjumptech/grule-rule-engine/engine"
	"github.com/hyperjumptech/grule-rule-engine/pkg"
	"testing"
)

//Tests for to check whether Grules support multiple KnowledgeBases
const Rules = `
rule  Rule1  "Rule1"  salience 20 {
when
	SampleFact.Text == "test"
Then
	SampleFact.Rule1Executed = true;
	Retract("Rule1");
}
rule  Rule2  "Rule2"  salience 10 {
when
	SampleFact.Text2 == "test"
Then
	SampleFact.Rule2Executed = true;
	Retract("Rule2");
}
`

type SampleFact struct {
	Text          string
	Text2         string
	Rule1Executed bool
	Rule2Executed bool
}

func TestGruleEngine_Support_Multiple_Rules(t *testing.T) {
	//Given
	sampleFact := &SampleFact{
		Text:          "test",
		Text2:         "test",
		Rule1Executed: false,
		Rule2Executed: false,
	}
	sampleDataContext := ast.NewDataContext()
	err := sampleDataContext.Add("SampleFact", sampleFact)
	if err != nil {
		t.Fatal(err)
	}
	lib := ast.NewKnowledgeLibrary()
	ruleBuilder := builder.NewRuleBuilder(lib)

	//When
	err = ruleBuilder.BuildRuleFromResource("SampleRuleSet", "0.1.1", pkg.NewBytesResource([]byte(Rules)))
	if err != nil {
		t.Fatal(err)
	}
	sampleKnowledgeBase := lib.NewKnowledgeBaseInstance("SampleRuleSet", "0.1.1")
	eng1 := engine.NewGruleEngine()
	err = eng1.Execute(sampleDataContext, sampleKnowledgeBase)
	if err != nil {
		t.Fatalf("Got error %v", err)
	}

	// Only the first rule should be executed
	if sampleFact.Rule1Executed == false {
		t.Fatalf("Expecting Rule1Executed to be true")
	}

	if sampleFact.Rule2Executed == true {
		t.Fatalf("Expecting Rule2Executed to be false")
	}
}
