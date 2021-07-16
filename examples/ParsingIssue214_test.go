package examples

import (
	"github.com/hyperjumptech/grule-rule-engine/ast"
	"github.com/hyperjumptech/grule-rule-engine/builder"
	"github.com/hyperjumptech/grule-rule-engine/pkg"
	"testing"
)

func TestParsingIssue(t *testing.T) {

	GRL :=
		`rule ExampleRuleName "One line rule description" salience 0  {
  when
    <FACT.SomeField == "true"
  then
    FACT.WriteMessage*( "message to write" );
}`

	lib := ast.NewKnowledgeLibrary()
	ruleBuilder := builder.NewRuleBuilder(lib)
	err := ruleBuilder.BuildRuleFromResource("SimpleTest", "1.0", pkg.NewBytesResource([]byte(GRL)))
	if err == nil {
		t.Errorf("Expected to be an error, but no error")
		t.Fail()
	} else {
		t.Log("Error as expected, instead of panic")
	}
}
