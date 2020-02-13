package examples

import (
	"github.com/hyperjumptech/grule-rule-engine/ast"
	"github.com/hyperjumptech/grule-rule-engine/builder"
	"github.com/hyperjumptech/grule-rule-engine/engine"
	"github.com/hyperjumptech/grule-rule-engine/pkg"
	"testing"
)

type ValueData struct {
	IntValue int
	Rating   string
	Expect   string
}

const (
	SalienceDRL = `

// Highest salience, if IntValue is bellow 33, all rule may match but this one take precedence
rule LowRule "If its on the low range, rating is low" salience 30 {
	When 
		V.Rating == "" &&
		V.IntValue < 33
	Then
		V.Rating = "Low";
}

// Lower salience, if IntValue is bellow 66, some rule may match but this one take precedence, unless there rule with highest salience are met (LowRule).
rule MediumRule "If its on the medium range, rating is medium" salience 20 {
	When 
		V.Rating == "" &&
		V.IntValue < 66
	Then
		V.Rating = "Medium";
}

// Even lower salience
rule HighRule "If its on the high range, rating is high" salience 10 {
	When 
		V.Rating == ""  &&
		V.IntValue < 300
	Then
		V.Rating = "High";
}


// Lowest and negative salience, will win the last and ensure other higher salience take precedence
rule ImpossibleRule "If its on the very very very high range, rating is high" salience -100 {
	When 
		V.Rating == ""
	Then
		V.Rating = "This is not right";
}
`
)

func TestSalience(t *testing.T) {
	testData := []*ValueData{
		&ValueData{
			IntValue: 10,
			Expect:   "Low",
		},
		&ValueData{
			IntValue: 20,
			Expect:   "Low",
		},
		&ValueData{
			IntValue: 30,
			Expect:   "Low",
		},
		&ValueData{
			IntValue: 40,
			Expect:   "Medium",
		},
		&ValueData{
			IntValue: 50,
			Expect:   "Medium",
		},
		&ValueData{
			IntValue: 60,
			Expect:   "Medium",
		},
		&ValueData{
			IntValue: 70,
			Expect:   "High",
		},
		&ValueData{
			IntValue: 80,
			Expect:   "High",
		},
		&ValueData{
			IntValue: 90,
			Expect:   "High",
		},
		&ValueData{
			IntValue: 1000000,
			Expect:   "This is not right",
		},
	}

	workingMemory := ast.NewWorkingMemory()
	knowledgeBase := ast.NewKnowledgeBase("Tutorial", "0.0.1")
	ruleBuilder := builder.NewRuleBuilder(knowledgeBase, workingMemory)
	byteArr := pkg.NewBytesResource([]byte(SalienceDRL))
	err := ruleBuilder.BuildRuleFromResource(byteArr)
	if err != nil {
		panic(err)
	}

	engine := engine.NewGruleEngine()

	for _, td := range testData {
		dataCtx := ast.NewDataContext()
		err := dataCtx.Add("V", td)
		if err != nil {
			panic(err)
		}

		err = engine.Execute(dataCtx, knowledgeBase, workingMemory)
		if err != nil {
			panic(err)
		}

		if td.Rating != td.Expect {
			t.Logf("On intValue = %d, expect \"%s\" but \"%s\"", td.IntValue, td.Expect, td.Rating)
			t.Fail()
		}
	}

}
