package examples

import (
	"fmt"
	"github.com/hyperjumptech/grule-rule-engine/ast"
	"github.com/hyperjumptech/grule-rule-engine/builder"
	"github.com/hyperjumptech/grule-rule-engine/engine"
	"github.com/hyperjumptech/grule-rule-engine/pkg"
	"github.com/sirupsen/logrus"
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
	logrus.SetLevel(logrus.DebugLevel)
	myFact := &MyFact{
		IntAttribute:     123,
		StringAttribute:  "Some string value",
		BooleanAttribute: true,
		FloatAttribute:   1.234,
		TimeAttribute:    time.Now(),
	}
	dataCtx := ast.NewDataContext()
	err := dataCtx.Add("MF", myFact)
	if err != nil {
		panic(err)
	}

	// Prepare knowledgebase library and load it with our rule.
	lib := ast.NewKnowledgeLibrary()
	rb := builder.NewRuleBuilder(lib)

	drls := `
rule CheckValues "Check the default values" salience 10 {
    when 
        MF.IntAttribute == 123 && MF.StringAttribute == "Some string value"
    then
        MF.WhatToSay = MF.GetWhatToSay("Hello Grule");
		Retract("CheckValues");
}
`

	err = rb.BuildRuleFromResource("Tutorial", "0.0.1", pkg.NewBytesResource([]byte(drls)))
	if err != nil {
		panic(err)
	}

	knowledgeBase := lib.NewKnowledgeBaseInstance("Tutorial", "0.0.1")

	engine := engine.NewGruleEngine()
	err = engine.Execute(dataCtx, knowledgeBase)
	if err != nil {
		panic(err)
	}

	if myFact.WhatToSay != "Let say \"Hello Grule\"" {
		t.Logf("Expected - Let say \"Hello Grule\" - but %s", myFact.WhatToSay)
		t.Fail()
	} else {
		println(myFact.WhatToSay)
	}

}
