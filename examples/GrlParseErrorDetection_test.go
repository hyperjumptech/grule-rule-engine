package examples

import (
	"github.com/hyperjumptech/grule-rule-engine/ast"
	"github.com/hyperjumptech/grule-rule-engine/builder"
	"github.com/hyperjumptech/grule-rule-engine/pkg"
	"testing"
)

const (
	RuleWithError = `
rule ErrorRule1 "Rule with error"  salience 10{
    when
      Pogo.Compare(User.Name, "Calo")  
    then
      User.Name = "Success";
      Log(User.Name)
      Retract("AgeNameCheck");
}
`
)

func TestParsingErrorDetection(t *testing.T) {
	lib := ast.NewKnowledgeLibrary()
	ruleBuilder := builder.NewRuleBuilder(lib)
	err := ruleBuilder.BuildRuleFromResource("Test", "0.1.1", pkg.NewBytesResource([]byte(RuleWithError)))

	if reporter, ok := err.(*pkg.GruleErrorReporter); ok && reporter.HasError() {
		for i, er := range reporter.Errors {
			t.Logf("detected error #%d : %s", i, er.Error())
		}
	} else {
		t.Error("There should be an error")
		t.FailNow()
	}
}
