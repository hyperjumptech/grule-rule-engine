package builder

import (
	"github.com/hyperjumptech/grule-rule-engine/ast"
	"github.com/hyperjumptech/grule-rule-engine/pkg"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNoPanic(t *testing.T) {
	GRL := `rule TestNoDesc { when true then Ok(); }`
	lib := ast.NewKnowledgeLibrary()
	ruleBuilder := NewRuleBuilder(lib)
	err := ruleBuilder.BuildRuleFromResource("CallingLog", "0.1.1", pkg.NewBytesResource([]byte(GRL)))
	assert.NoError(t, err)
}
