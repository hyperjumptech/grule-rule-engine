package examples

import (
	"fmt"
	"github.com/hyperjumptech/grule-rule-engine/ast"
	"github.com/hyperjumptech/grule-rule-engine/builder"
	"github.com/hyperjumptech/grule-rule-engine/engine"
	"github.com/hyperjumptech/grule-rule-engine/pkg"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

type MyFact struct {
	IntAttribute     int64
	StringAttribute  string
	BooleanAttribute bool
	FloatAttribute   float64
	TimeAttribute    time.Time
	WhatToSay        string
}

func (mf *MyFact) GetWhatToSay(sentence string) string {
	return fmt.Sprintf("Let say \"%s\"", sentence)
}

func TestTutorial(t *testing.T) {
	myFact := &MyFact{
		IntAttribute:     123,
		StringAttribute:  "Some string value",
		BooleanAttribute: true,
		FloatAttribute:   1.234,
		TimeAttribute:    time.Now(),
	}
	dataCtx := ast.NewDataContext()
	err := dataCtx.Add("MF", myFact)
	assert.NoError(t, err)

	// Prepare knowledgebase library and load it with our rule.
	knowledgeLibrary := ast.NewKnowledgeLibrary()
	ruleBuilder := builder.NewRuleBuilder(knowledgeLibrary)

	drls := `
rule CheckValues "Check the default values" salience 10 {
    when 
        MF.IntAttribute == 123 && MF.StringAttribute == "Some string value"
    then
        MF.WhatToSay = MF.GetWhatToSay("Hello Grule");
		Retract("CheckValues");
}
`
	byteArr := pkg.NewBytesResource([]byte(drls))
	err = ruleBuilder.BuildRuleFromResource("Tutorial", "0.0.1", byteArr)
	assert.NoError(t, err)

	knowledgeBase := knowledgeLibrary.NewKnowledgeBaseInstance("Tutorial", "0.0.1")

	engine := engine.NewGruleEngine()
	err = engine.Execute(dataCtx, knowledgeBase)
	assert.NoError(t, err)
	assert.Equal(t, "Let say \"Hello Grule\"", myFact.WhatToSay)
	println(myFact.WhatToSay)
}
