//  Copyright hyperjumptech/grule-rule-engine Authors
//
//  Licensed under the Apache License, Version 2.0 (the "License");
//  you may not use this file except in compliance with the License.
//  You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
//  Unless required by applicable law or agreed to in writing, software
//  distributed under the License is distributed on an "AS IS" BASIS,
//  WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//  See the License for the specific language governing permissions and
//  limitations under the License.

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

type MyḞact struct {
	IntAttribute     int64
	StringǍttribute  string
	BooleanAttribute bool
	ḞloatAttribute   float64
	TimeAttribute    time.Time
	WhatToSay        string
}

func (mf *MyḞact) GetẀhatToSay(sentence string) string {
	return fmt.Sprintf("Let say \"%s\"", sentence)
}

func TestUnicodeTutorial(t *testing.T) {
	myFact := &MyḞact{
		IntAttribute:     123,
		StringǍttribute:  "Some string vǍlue",
		BooleanAttribute: true,
		ḞloatAttribute:   1.234,
		TimeAttribute:    time.Now(),
	}
	dataCtx := ast.NewDataContext()
	err := dataCtx.Add("MḞ", myFact)
	assert.NoError(t, err)

	// Prepare knowledgebase library and load it with our rule.
	knowledgeLibrary := ast.NewKnowledgeLibrary()
	ruleBuilder := builder.NewRuleBuilder(knowledgeLibrary)

	drls := `
rule ChĕckValuĕs "Check the default values" salience 10 {
    when 
        MḞ.IntAttribute == 123 && MḞ.StringǍttribute == "Some string vǍlue"
    then
        MḞ.WhatToSay = MḞ.GetẀhatToSay("HellǑ Grule");
		Retract("ChĕckValuĕs");
}
`
	byteArr := pkg.NewBytesResource([]byte(drls))
	err = ruleBuilder.BuildRuleFromResource("Tutorial", "0.0.1", byteArr)
	assert.NoError(t, err)

	knowledgeBase := knowledgeLibrary.NewKnowledgeBaseInstance("Tutorial", "0.0.1")

	engine := engine.NewGruleEngine()
	err = engine.Execute(dataCtx, knowledgeBase)
	assert.NoError(t, err)
	assert.Equal(t, "Let say \"HellǑ Grule\"", myFact.WhatToSay)
	println(myFact.WhatToSay)
}
