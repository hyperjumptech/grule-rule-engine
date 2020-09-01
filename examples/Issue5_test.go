package examples

import (
	"github.com/hyperjumptech/grule-rule-engine/ast"
	"github.com/hyperjumptech/grule-rule-engine/builder"
	"github.com/hyperjumptech/grule-rule-engine/engine"
	"github.com/hyperjumptech/grule-rule-engine/pkg"
	"github.com/stretchr/testify/assert"
	"testing"
)

const (
	Rule = `
rule UserTestRule3 "test 3"  salience 10{
	when
	  User.GetName() == "Watson"
	then
	  User.SetName("FromRuleScope3");
	  Retract("UserTestRule3");
}
`
)

type AUser struct {
	Name string
}

func (u *AUser) GetName() string {
	return u.Name
}

func (u *AUser) SetName(name string) {
	u.Name = name
}

func TestMethodCall_Issue5(t *testing.T) {
	user := &AUser{
		Name: "Watson",
	}

	dataContext := ast.NewDataContext()
	err := dataContext.Add("User", user)
	assert.NoError(t, err)

	// Prepare knowledgebase library and load it with our rule.
	lib := ast.NewKnowledgeLibrary()
	rb := builder.NewRuleBuilder(lib)
	err = rb.BuildRuleFromResource("Test", "0.1.1", pkg.NewBytesResource([]byte(Rule)))
	assert.NoError(t, err)
	eng1 := &engine.GruleEngine{MaxCycle: 5}
	kb := lib.NewKnowledgeBaseInstance("Test", "0.1.1")
	err = eng1.Execute(dataContext, kb)
	assert.NoError(t, err)
	assert.Equal(t, "FromRuleScope3", user.GetName())
}
