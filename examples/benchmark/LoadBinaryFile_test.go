package benchmark

import (
	"bufio"
	"bytes"
	"github.com/hyperjumptech/grule-rule-engine/ast"
	"github.com/hyperjumptech/grule-rule-engine/builder"
	"github.com/hyperjumptech/grule-rule-engine/engine"
	"github.com/hyperjumptech/grule-rule-engine/pkg"
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"log"
	"os"
	"testing"
	"time"
)

func TestCreateBinaryFile(t *testing.T) {
	start := time.Now()
	input, _ := ioutil.ReadFile("1000_rules.grl")
	rules := string(input)
	fact := &RideFact{
		Distance: 6000,
		Duration: 121,
	}
	dctx := ast.NewDataContext()
	_ = dctx.Add("Fact", fact)

	lib := ast.NewKnowledgeLibrary()
	rb := builder.NewRuleBuilder(lib)
	_ = rb.BuildRuleFromResource("load_rules_test", "0.1.1", pkg.NewBytesResource([]byte(rules)))
	_ = lib.NewKnowledgeBaseInstance("load_rules_test", "0.1.1")
	elapsed := time.Since(start)
	log.Printf("Loading 1000 rules took with out binary %v ms", elapsed.Milliseconds())
	buf := new(bytes.Buffer)
	start = time.Now()
	err := lib.StoreKnowledgeBaseToWriter(buf, "load_rules_test", "0.1.1")
	assert.Nil(t, err)
	f, err := os.Create("kb.dat")
	assert.Nil(t, err)
	_, err = f.Write(buf.Bytes())
	assert.Nil(t, err)
	defer f.Close()
	elapsed = time.Since(start)
	log.Printf("Converting 1000 rules in to binary format took:%v ms", elapsed.Milliseconds())
}

func TestLoad1000RulesIntoKnowledgeBaseFromBinary(t *testing.T) {
	start := time.Now()
	fact := &RideFact{
		Distance: 6000,
		Duration: 121,
	}
	dctx := ast.NewDataContext()
	_ = dctx.Add("Fact", fact)

	f2, err := os.Open("kb.dat")
	assert.Nil(t, err)
	reader := bufio.NewReader(f2)
	lib := ast.NewKnowledgeLibrary()
	kb, err := lib.LoadKnowledgeBaseFromReader(reader, true)
	assert.Nil(t, err)
	e := engine.NewGruleEngine()
	err = e.Execute(dctx, kb)
	assert.Nil(t, err)
	elapsed := time.Since(start)
	log.Printf("Load 1000 rules from binary took %v ms", elapsed.Milliseconds())
}
