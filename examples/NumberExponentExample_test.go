package examples

import (
	"github.com/hyperjumptech/grule-rule-engine/ast"
	"github.com/hyperjumptech/grule-rule-engine/builder"
	"github.com/hyperjumptech/grule-rule-engine/engine"
	"github.com/hyperjumptech/grule-rule-engine/pkg"
	"github.com/stretchr/testify/assert"
	"testing"
)

type ExponentData struct {
	Check float64
	Set   float64
}

const ExponentRule = `
rule  ExponentCheck  "User Related Rule"  salience 10 {
	when 
		ExponentData.Check == 6.67428e-11
	Then
		ExponentData.Set = .12345E+5;
		Retract("ExponentCheck");
}
`

func TestEvaluateAndAssignExponentNumber(t *testing.T) {
	exponent := &ExponentData{
		Check: 6.67428e-11,
		Set:   0,
	}

	dataContext := ast.NewDataContext()
	err := dataContext.Add("ExponentData", exponent)
	assert.NoError(t, err)

	// Prepare knowledgebase library and load it with our rule.
	lib := ast.NewKnowledgeLibrary()
	rb := builder.NewRuleBuilder(lib)
	err = rb.BuildRuleFromResource("TestExponent", "1.0.0", pkg.NewBytesResource([]byte(ExponentRule)))
	assert.NoError(t, err)
	eng1 := &engine.GruleEngine{MaxCycle: 5}
	kb := lib.NewKnowledgeBaseInstance("TestExponent", "1.0.0")
	err = eng1.Execute(dataContext, kb)
	assert.NoError(t, err)
	assert.Equal(t, .12345e+5, exponent.Set)

}
