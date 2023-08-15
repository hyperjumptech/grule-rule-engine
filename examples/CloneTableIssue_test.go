package examples

import (
	"github.com/kalyan-arepalle/grule-rule-engine/ast"
	"github.com/kalyan-arepalle/grule-rule-engine/builder"
	"github.com/kalyan-arepalle/grule-rule-engine/engine"
	"github.com/kalyan-arepalle/grule-rule-engine/pkg"
	"github.com/stretchr/testify/assert"
	"testing"
)

type StructStringsData struct {
	Strings []string
}

func (f *StructStringsData) GetStrings() []string {
	return f.Strings
}

const panickingRule = ` rule test {
when 
	Fact.GetStrings()[0] == Fact.GetStrings()[1]
then
	Complete();
}`

func TestSliceFunctionPanicTest(t *testing.T) {
	fact := &StructStringsData{
		Strings: []string{"0", "0"},
	}

	dataContext := ast.NewDataContext()
	err := dataContext.Add("Fact", fact)
	assert.NoError(t, err)
	knowledgeLibrary := ast.NewKnowledgeLibrary()
	ruleBuilder := builder.NewRuleBuilder(knowledgeLibrary)
	err = ruleBuilder.BuildRuleFromResource("test", "0.0.1", pkg.NewBytesResource([]byte(panickingRule)))
	assert.NoError(t, err)
	knowledgeBase := knowledgeLibrary.NewKnowledgeBaseInstance("test", "0.0.1")
	engine := engine.NewGruleEngine()

	err = engine.Execute(dataContext, knowledgeBase)
	assert.NoError(t, err)
}
