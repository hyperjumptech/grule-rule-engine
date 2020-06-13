package examples

import (
	"github.com/hyperjumptech/grule-rule-engine/ast"
	"github.com/hyperjumptech/grule-rule-engine/builder"
	"github.com/hyperjumptech/grule-rule-engine/engine"
	"github.com/hyperjumptech/grule-rule-engine/pkg"
	"testing"
)

const (
	Rule4 = `
rule UserTestRule4 "test 3"  salience 10{
	when
	  User.Auth.GetEmail() == "watson@test.com"
	then
	  User.Name = "FromRuleScope4";
	  Retract("UserTestRule4");
}
`
)

type UserWithAuth struct {
	Auth *UserAuth
	Name string
}

func (user *UserWithAuth) GetName() string {
	return user.Name
}

type UserAuth struct {
	Email string
}

func (auth *UserAuth) GetEmail() string {
	return auth.Email
}

func TestMethodCall_Issue4(t *testing.T) {
	user := &UserWithAuth{
		Auth: &UserAuth{Email: "watson@test.com"},
	}

	if user.GetName() != "" {
		t.Fatal("User name not empty")
	}

	dataContext := ast.NewDataContext()
	err := dataContext.Add("User", user)
	if err != nil {
		t.Fatal(err)
	}

	// Prepare knowledgebase library and load it with our rule.
	lib := ast.NewKnowledgeLibrary()
	rb := builder.NewRuleBuilder(lib)
	err = rb.BuildRuleFromResource("Test", "0.1.1", pkg.NewBytesResource([]byte(Rule4)))
	if err != nil {
		t.Fatal(err)
	} else {
		kb := lib.NewKnowledgeBaseInstance("Test", "0.1.1")
		eng1 := &engine.GruleEngine{MaxCycle: 3}
		err := eng1.Execute(dataContext, kb)
		if err != nil {
			t.Fatal(err)
		}
		if user.GetName() != "FromRuleScope4" {
			t.Fatalf("User should be FromRuleScope4 but %s", user.GetName())
		}
	}
}
