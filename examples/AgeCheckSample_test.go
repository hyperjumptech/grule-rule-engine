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

	memory := ast.NewWorkingMemory()
	knowledgeBase := ast.NewKnowledgeBase("Test", "0.1.1")
	ruleBuilder := builder.NewRuleBuilder(knowledgeBase, memory)

	err = ruleBuilder.BuildRuleFromResource(pkg.NewBytesResource([]byte(rule2)))
	if err != nil {
		t.Fatal(err)
	} else {
		eng1 := &engine.GruleEngine{MaxCycle: 1}
		err := eng1.Execute(dataContext, knowledgeBase, memory)
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

	memory := ast.NewWorkingMemory()
	knowledgeBase := ast.NewKnowledgeBase("Test", "0.1.1")
	ruleBuilder := builder.NewRuleBuilder(knowledgeBase, memory)

	err = ruleBuilder.BuildRuleFromResource(pkg.NewBytesResource([]byte(rule3)))
	if err != nil {
		t.Log(err)
	} else {
		eng1 := &engine.GruleEngine{MaxCycle: 100}
		err := eng1.Execute(dataContext, knowledgeBase, memory)
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
