package examples

import (
	"github.com/DataWiseHQ/grule-rule-engine/ast"
	"github.com/DataWiseHQ/grule-rule-engine/builder"
	"github.com/DataWiseHQ/grule-rule-engine/pkg"
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

// TestParsingErrorDetection demonstrate how to make your own GRL script/syntax validator using the RuleBuilder.
func TestParsingErrorDetection(t *testing.T) {

	// Normal stuff, creating a library and builder for it.
	lib := ast.NewKnowledgeLibrary()
	ruleBuilder := builder.NewRuleBuilder(lib)

	// Build normally
	err := ruleBuilder.BuildRuleFromResource("Test", "0.1.1", pkg.NewBytesResource([]byte(RuleWithError)))

	// If the err != nil something is wrong.
	if err != nil {
		// Cast the error into pkg.GruleErrorReporter with type checking.
		// Type checking is necessary because the err might not only parsing error.
		if reporter, ok := err.(*pkg.GruleErrorReporter); ok {
			// Lets iterate all the error we get during parsing.
			for i, er := range reporter.Errors {
				t.Logf("detected error #%d : %s", i, er.Error())
			}
		} else {
			// Well, its an error but not GruleErrorReporter instance. could be IO error.
			t.Error("There should be GruleErrorReporter")
			t.FailNow()
		}
	}
}
