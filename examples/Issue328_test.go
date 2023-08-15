package examples

import (
	"testing"

	"github.com/kalyan-arepalle/grule-rule-engine/ast"
	"github.com/kalyan-arepalle/grule-rule-engine/builder"
	"github.com/kalyan-arepalle/grule-rule-engine/engine"
	"github.com/kalyan-arepalle/grule-rule-engine/pkg"
	"github.com/stretchr/testify/assert"
)

const (
	SliceOORRule = `
		rule SliceOORRule {
			when
				PriceSlice.Prices[4] > 10 // will cause panic
			then
				Log("Price number 4 is greater than 10");
				Retract("SliceOORRule");
		}`
)

type AUserSliceIssue struct {
	Prices []int
}

func TestMethodCall_SliceOOR(t *testing.T) {
	ps := &AUserSliceIssue{
		Prices: []int{1, 2, 3},
	}

	dataContext := ast.NewDataContext()
	err := dataContext.Add("PriceSlice", ps)
	assert.NoError(t, err)

	// Prepare knowledgebase library and load it with our rule.
	lib := ast.NewKnowledgeLibrary()
	rb := builder.NewRuleBuilder(lib)
	err = rb.BuildRuleFromResource("Test", "0.1.1", pkg.NewBytesResource([]byte(SliceOORRule)))
	assert.NoError(t, err)

	// expect no panic and no error (ReturnErrOnFailedRuleEvaluation = false)
	eng1 := &engine.GruleEngine{MaxCycle: 5}
	kb := lib.NewKnowledgeBaseInstance("Test", "0.1.1")
	err = eng1.Execute(dataContext, kb)
	assert.NoError(t, err)

	// expect no panic and execute to return an error here
	eng1 = &engine.GruleEngine{MaxCycle: 5, ReturnErrOnFailedRuleEvaluation: true}
	kb = lib.NewKnowledgeBaseInstance("Test", "0.1.1")
	err = eng1.Execute(dataContext, kb)
	assert.Error(t, err)
}
