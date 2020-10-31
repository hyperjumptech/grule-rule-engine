package benchmark

import (
	"fmt"
	"github.com/hyperjumptech/grule-rule-engine/ast/v2"
	"github.com/hyperjumptech/grule-rule-engine/builder"
	"github.com/hyperjumptech/grule-rule-engine/engine"
	"github.com/hyperjumptech/grule-rule-engine/pkg"
	"io/ioutil"
	"testing"
)

/**
  Benchmarking `engine.Execute` function by running 100 and 1000 rules with different N values
  Please refer docs/benchmarking_en.md for more info
*/

var knowledgeBase *v2.KnowledgeBase

func Benchmark_Grule_Execution_Engine(b *testing.B) {
	rules := []struct {
		name string
		fun  func()
	}{
		{"100 rules", load100RulesIntoKnowledgebase},
		{"1000 rules", load1000RulesIntoKnowledgebase},
	}
	for _, rule := range rules {
		for k := 0.; k <= 10; k++ {
			b.Run(fmt.Sprintf("%s", rule.name), func(b *testing.B) {
				rule.fun()
				for i := 0; i < b.N; i++ {
					f1 := RideFact{
						Distance: 6000,
						Duration: 121,
					}
					e := engine.NewGruleEngine()
					//Fact1
					dataCtx := v2.NewDataContext()
					err := dataCtx.Add("Fact", &f1)
					if err != nil {
						b.Fail()
					}
					err = e.Execute(dataCtx, knowledgeBase)
					if err != nil {
						fmt.Print(err)
					}
				}
			})
		}
	}
}

func load100RulesIntoKnowledgebase() {
	input, _ := ioutil.ReadFile("100_complicated_rules.grl")
	rules := string(input)
	fact := &RideFact{
		Distance: 6000,
		Duration: 121,
	}
	dctx := v2.NewDataContext()
	_ = dctx.Add("Fact", fact)

	lib := v2.NewKnowledgeLibrary()
	rb := builder.NewRuleBuilder(lib)
	_ = rb.BuildRuleFromResource("exec_rules_test", "0.1.1", pkg.NewBytesResource([]byte(rules)))
	knowledgeBase = lib.NewKnowledgeBaseInstance("exec_rules_test", "0.1.1")
}

func load1000RulesIntoKnowledgebase() {
	input, _ := ioutil.ReadFile("1000_complicated_rules.grl")
	rules := string(input)
	fact := &RideFact{
		Distance: 6000,
		Duration: 121,
	}
	dctx := v2.NewDataContext()
	_ = dctx.Add("Fact", fact)

	lib := v2.NewKnowledgeLibrary()
	rb := builder.NewRuleBuilder(lib)
	_ = rb.BuildRuleFromResource("exec_rules_test", "0.1.1", pkg.NewBytesResource([]byte(rules)))
	knowledgeBase = lib.NewKnowledgeBaseInstance("exec_rules_test", "0.1.1")
}
