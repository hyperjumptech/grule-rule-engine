package examples

import (
	"github.com/kalyan-arepalle/grule-rule-engine/ast"
	"github.com/kalyan-arepalle/grule-rule-engine/builder"
	"github.com/kalyan-arepalle/grule-rule-engine/engine"
	"github.com/kalyan-arepalle/grule-rule-engine/pkg"
	"github.com/stretchr/testify/assert"
	"testing"
)

type TestData struct {
	Index         int
	Strings       []string
	Concatenation string
}

func (f *TestData) GetStrings() []string {
	return f.Strings
}

const rule = ` rule test {
when 
	Fact.Index < Fact.Strings.Len()
then
	Fact.Concatenation = Fact.Concatenation + Fact.GetStrings()[Fact.Index];
	Fact.Index = Fact.Index + 1;
}`

func TestSliceFunctionTest(t *testing.T) {
	fact := &TestData{
		Index:   0,
		Strings: []string{"1", "2", "3"},
	}

	dataContext := ast.NewDataContext()
	err := dataContext.Add("Fact", fact)
	assert.NoError(t, err)
	knowledgeLibrary := ast.NewKnowledgeLibrary()
	ruleBuilder := builder.NewRuleBuilder(knowledgeLibrary)
	err = ruleBuilder.BuildRuleFromResource("test", "0.0.1", pkg.NewBytesResource([]byte(rule)))
	assert.NoError(t, err)
	knowledgeBase := knowledgeLibrary.NewKnowledgeBaseInstance("test", "0.0.1")
	engine := engine.NewGruleEngine()

	err = engine.Execute(dataContext, knowledgeBase)
	assert.NoError(t, err)
	assert.Equal(t, "123", fact.Concatenation)
}
