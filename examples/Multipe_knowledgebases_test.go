package examples

import (
	"github.com/hyperjumptech/grule-rule-engine/ast"
	"github.com/hyperjumptech/grule-rule-engine/builder"
	"github.com/hyperjumptech/grule-rule-engine/engine"
	"github.com/hyperjumptech/grule-rule-engine/pkg"
	"testing"
)

//Tests for to check whether Grules support multiple KnowledgeBases
const RideRule = `
rule  RideRule  "Ride Related Rule"  salience 10 {
when
(RideFact.Distance > 5000  ||   RideFact.Duration > 120) || (RideFact.RideType == "On-Demand" && RideFact.IsFrequentCustomer == true)
Then
RideFact.NetAmount=143.320007;
RideFact.Result=true;
Retract("RideRule");
}
`

const UserRule = `
rule  UserRule  "User Related Rule"  salience 10 {
when User.IsPalindrome(User.Name) == true
Then
User.Message = "Palindrome";
Retract("UserRule");
}
`

type RideFact struct {
	Distance           int32
	Duration           int32
	RideType           string
	IsFrequentCustomer bool
	Result             bool
	NetAmount          float32
}

type UserFact struct {
	Name    string
	Message string
}

func (u *UserFact) IsPalindrome(name string) bool {
	for i := 0; i < len(name)/2; i++ {
		if name[i] != name[len(name)-i-1] {
			return false
		}
	}
	return true
}

func TestGruleEngine_Support_Multiple_KnowledgeBases(t *testing.T) {
	//Given
	user := &UserFact{
		Name:    "madam",
		Message: "Not a Palindrome", // Default
	}
	rideFact := &RideFact{
		Distance: 6000,
		Duration: 121,
	}
	userDataContext := ast.NewDataContext()
	err := userDataContext.Add("User", user)
	if err != nil {
		t.Fatal(err)
	}
	rideDataContext := ast.NewDataContext()
	err = rideDataContext.Add("RideFact", rideFact)
	if err != nil {
		t.Fatal(err)
	}
	lib := ast.NewKnowledgeLibrary()
	ruleBuilder := builder.NewRuleBuilder(lib)

	//When
	err = ruleBuilder.BuildRuleFromResource("UserRules", "0.1.1", pkg.NewBytesResource([]byte(UserRule)))
	if err != nil {
		t.Fatal(err)
	}
	err = ruleBuilder.BuildRuleFromResource("RideRules", "0.1.1", pkg.NewBytesResource([]byte(RideRule)))
	if err != nil {
		t.Fatal(err)
	}
	userKnowledgeBase := lib.NewKnowledgeBaseInstance("UserRules", "0.1.1")
	rideKnowledgeBase := lib.NewKnowledgeBaseInstance("RideRules", "0.1.1")
	eng1 := engine.NewGruleEngine()
	err = eng1.Execute(rideDataContext, rideKnowledgeBase)
	if err != nil {
		t.Fatalf("Got error %v", err)
	}
	err = eng1.Execute(userDataContext, userKnowledgeBase)
	if err != nil {
		t.Fatalf("Got error %v", err)
	}

	//Then
	if user.Message != "Palindrome" {
		t.Fatalf("Expecting Palindrome")
	}

	if rideFact.NetAmount != float32(143.32) {
		t.Fatalf("NetAmount is not as expected")
	}
}
