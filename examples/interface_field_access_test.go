package examples

import (
	"testing"

	"github.com/hyperjumptech/grule-rule-engine/ast"
	"github.com/hyperjumptech/grule-rule-engine/builder"
	"github.com/hyperjumptech/grule-rule-engine/engine"
	"github.com/hyperjumptech/grule-rule-engine/pkg"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// PayloadContainer represents a struct with an interface{} field
type PayloadContainer struct {
	Type    string
	Payload interface{}
}

// NestedData represents the concrete type stored in the interface
type NestedData struct {
	Status string
	Value  int
}

func TestInterfaceFieldAccess(t *testing.T) {
	t.Parallel()
	t.Run("should not fail when accessing interface fields directly", func(t *testing.T) {
		// Rule that tries to access interface field directly
		ruleText := `
rule InterfaceFieldRule "test interface field access" {
	when
		Data.Type == "test" &&
		Data.Payload.Status == "active"
	then
		Data.Type = "processed";
		Log("Success");
}`

		// Test data with interface{} field
		testData := &PayloadContainer{
			Type: "test",
			Payload: NestedData{
				Status: "active",
				Value:  42,
			},
		}

		dataCtx := ast.NewDataContext()
		require.NoError(t, dataCtx.Add("Data", testData))

		knowledgeLibrary := ast.NewKnowledgeLibrary()
		ruleBuilder := builder.NewRuleBuilder(knowledgeLibrary)

		byteArr := pkg.NewBytesResource([]byte(ruleText))
		err := ruleBuilder.BuildRuleFromResource("Test", "1.0.0", byteArr)
		require.NoError(t, err)

		knowledgeBase, err := knowledgeLibrary.NewKnowledgeBaseInstance("Test", "1.0.0")
		require.NoError(t, err)

		gruleEngine := engine.NewGruleEngine()
		gruleEngine.ReturnErrOnFailedRuleEvaluation = true
		err = gruleEngine.Execute(dataCtx, knowledgeBase)

		assert.NoError(t, err, "should not fail to access interface field")
		assert.Equal(t, "processed", testData.Type, "rule should have executed and modified the data")

		// Verify the interface field value remains unchanged since we're only reading it
		payload, ok := testData.Payload.(NestedData)
		require.True(t, ok, "type assertion should succeed")
		assert.Equal(t, "active", payload.Status, "interface field should remain unchanged")
	})

	t.Run("should allow modifying interface fields when using pointer to struct", func(t *testing.T) {
		// Rule that modifies interface field
		ruleText := `
rule InterfaceFieldModifyRule "test interface field modification" {
	when
		Data.Type == "test" &&
		Data.Payload.Status == "active"
	then
		Data.Type = "processed";
		Data.Payload.Status = "handled";
		Log("Modified interface field");
}`

		// Test data with interface{} field containing a pointer to struct (addressable)
		testData := &PayloadContainer{
			Type: "test",
			Payload: &NestedData{
				Status: "active",
				Value:  42,
			},
		}

		dataCtx := ast.NewDataContext()
		require.NoError(t, dataCtx.Add("Data", testData))

		knowledgeLibrary := ast.NewKnowledgeLibrary()
		ruleBuilder := builder.NewRuleBuilder(knowledgeLibrary)

		byteArr := pkg.NewBytesResource([]byte(ruleText))
		err := ruleBuilder.BuildRuleFromResource("Test", "1.0.0", byteArr)
		require.NoError(t, err)

		knowledgeBase, err := knowledgeLibrary.NewKnowledgeBaseInstance("Test", "1.0.0")
		require.NoError(t, err)

		gruleEngine := engine.NewGruleEngine()
		gruleEngine.ReturnErrOnFailedRuleEvaluation = true
		err = gruleEngine.Execute(dataCtx, knowledgeBase)

		assert.NoError(t, err, "should not fail to modify interface field when using pointer")
		assert.Equal(t, "processed", testData.Type, "rule should have executed and modified the data")

		// Verify the interface field was modified
		payload, ok := testData.Payload.(*NestedData)
		require.True(t, ok, "type assertion should succeed")
		assert.Equal(t, "handled", payload.Status, "interface field should have been modified")
	})
}
