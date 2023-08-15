package examples

import (
	"fmt"
	"testing"

	"github.com/kalyan-arepalle/grule-rule-engine/ast"
	"github.com/kalyan-arepalle/grule-rule-engine/builder"
	"github.com/kalyan-arepalle/grule-rule-engine/engine"
	"github.com/kalyan-arepalle/grule-rule-engine/pkg"
)

const (
	inputRule = `
	rule TestRule "" {
		when
			R.Result == 'NoResult' &&
			inputs.i_am_missing == 'abc' &&
                        inputs.name.first == 'john'
		then
			R.Result = "ok";
	}
	`
)

func TestDataContextMissingFact(t *testing.T) {

	oresult := &ObjectResult{
		Result: "NoResult",
	}

	// build rules
	lib := ast.NewKnowledgeLibrary()
	rb := builder.NewRuleBuilder(lib)
	err := rb.BuildRuleFromResource("Test", "0.0.1", pkg.NewBytesResource([]byte(inputRule)))

	// 	add JSON fact
	json := []byte(`{"blabla":"bla","name":{"first":"john","last":"doe"}}`)
	kb := lib.NewKnowledgeBaseInstance("Test", "0.0.1")
	dcx := ast.NewDataContext()

	err = dcx.Add("R", oresult)
	err = dcx.AddJSON("inputs", json)
	if err != nil {
		fmt.Println(err.Error())
	}

	// results in panic
	engine.NewGruleEngine().Execute(dcx, kb)

}
