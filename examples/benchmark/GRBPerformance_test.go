package benchmark

import (
	"github.com/DataWiseHQ/grule-rule-engine/ast"
	"github.com/DataWiseHQ/grule-rule-engine/builder"
	"github.com/DataWiseHQ/grule-rule-engine/pkg"
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
	"time"
)

var (
	ruleCount = 5000

	grlFile = "5000_dummy_rules.grl"
	grbFile = "5000_dummy_rules.grb"

	knowledgeName    = "Dummy Knowledge"
	knowledgeVersion = "v1.0.0"
)

func TestSerializationPerformanceOnFile(t *testing.T) {
	if testing.Short() {
		return
	}

	t.Log("GENERATING DUMMY RULES")
	err := GenRandomRule(grlFile, ruleCount)
	assert.Nil(t, err)

	t.Log("LOADING GRL")
	timer := time.Now()

	// SAFING INTO FILE
	// First prepare our library for loading the orignal GRL
	lib := ast.NewKnowledgeLibrary()
	rb := builder.NewRuleBuilder(lib)
	err = rb.BuildRuleFromResource(knowledgeName, knowledgeVersion, pkg.NewFileResource(grlFile))
	assert.NoError(t, err)

	durationA := time.Since(timer)
	t.Log("SAVING BINARY INTO GRB")
	timer = time.Now()

	// Lets create the file to write into
	f, err := os.OpenFile(grbFile, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	assert.Nil(t, err)

	// Save the knowledge base into the file and close it.
	err = lib.StoreKnowledgeBaseToWriter(f, knowledgeName, knowledgeVersion)
	assert.Nil(t, err)
	_ = f.Close()

	durationB := time.Since(timer)
	t.Log("LOADING BINARY FROM GRB")
	timer = time.Now()

	// LOADING FROM FILE
	// Lets assume we are using different library, so create a new one
	lib2 := ast.NewKnowledgeLibrary()

	// Open the existing safe file
	f2, err := os.Open(grbFile)
	assert.Nil(t, err)

	// Load the file directly into the library and close the file
	kb2, err := lib2.LoadKnowledgeBaseFromReader(f2, true)
	assert.Nil(t, err)
	_ = f2.Close()

	durationC := time.Since(timer)

	t.Log("-------------------------------------------")
	t.Logf("Load GRL duration   : %d ms", durationA.Milliseconds())
	t.Logf("Saving GRB duration : %d ms", durationB.Milliseconds())
	t.Logf("Load GRB duration   : %d ms", durationC.Milliseconds())

	// Delete the test file
	err = os.Remove(grlFile)
	assert.Nil(t, err)
	err = os.Remove(grbFile)
	assert.Nil(t, err)

	// compare that the original knowledgebase is exacly the same to the loaded one.
	assert.True(t, lib.GetKnowledgeBase(knowledgeName, knowledgeVersion).IsIdentical(kb2))
}
