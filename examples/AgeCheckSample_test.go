package examples

import (
	"github.com/hyperjumptech/grule-rule-engine/ast"
	"github.com/hyperjumptech/grule-rule-engine/builder"
	"github.com/hyperjumptech/grule-rule-engine/engine"
	"github.com/hyperjumptech/grule-rule-engine/pkg"
	"testing"
)

const (
	rule2 = `
rule AgeNameCheck "test" {
    when
      Pogo.GetStringLength("9999") > 0 
    then
      Log(User.Name);
}
`

	rule3 = `
rule AgeNameCheck "test"  salience 10{
    when
      Pogo.Compare(User.Name, "Calo")  
    then
      User.Name = "Success";
      Log(User.Name);
      Retract("AgeNameCheck");
}
`
)

func TestMyPoGo_GetStringLength(t *testing.T) {
	user := &User{
		Name: "Calo",
		Age:  0,
		Male: true,
	}

	dataContext := ast.NewDataContext()
	err := dataContext.Add("User", user)
	if err != nil {
		t.Fatal(err)
	}
	err = dataContext.Add("Pogo", &MyPoGo{})
	if err != nil {
		t.Fatal(err)
	}

	lib := ast.NewKnowledgeLibrary()
	ruleBuilder := builder.NewRuleBuilder(lib)

	err = ruleBuilder.BuildRuleFromResource("Test", "0.1.1", pkg.NewBytesResource([]byte(rule2)))
	if err != nil {
		t.Fatal(err)
	} else {
		kb := lib.NewKnowledgeBaseInstance("Test", "0.1.1")
		eng1 := &engine.GruleEngine{MaxCycle: 1}
		err := eng1.Execute(dataContext, kb)
		if err != nil {
			t.Fatalf("Got error %v", err)
		} else {
			t.Log(user)
		}
	}
}

func TestMyPoGo_Compare(t *testing.T) {
	user := &User{
		Name: "Calo",
		Age:  0,
		Male: true,
	}

	dataContext := ast.NewDataContext()
	err := dataContext.Add("User", user)
	if err != nil {
		t.Fatal(err)
	}
	err = dataContext.Add("Pogo", &MyPoGo{})
	if err != nil {
		t.Fatal(err)
	}

	lib := ast.NewKnowledgeLibrary()
	ruleBuilder := builder.NewRuleBuilder(lib)

	err = ruleBuilder.BuildRuleFromResource("Test", "0.1.1", pkg.NewBytesResource([]byte(rule3)))
	if err != nil {
		t.Fatal(err)
	} else {
		kb := lib.NewKnowledgeBaseInstance("Test", "0.1.1")
		eng1 := &engine.GruleEngine{MaxCycle: 100}
		err := eng1.Execute(dataContext, kb)
		if err != nil {
			t.Fatalf("Got error %v", err)
		} else {
			t.Log(user)
		}
		if user.Name != "Success" {
			t.Logf("User should have changed name by rule to Success, but %s", user.Name)
			t.FailNow()
		}
	}
}
