// Copyright DataWiseHQ/grule-rule-engine Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package examples

import (
	"bytes"
	"github.com/hyperjumptech/grule-rule-engine/ast"
	"github.com/hyperjumptech/grule-rule-engine/builder"
	"github.com/hyperjumptech/grule-rule-engine/engine"
	"github.com/hyperjumptech/grule-rule-engine/pkg"
	"github.com/rs/zerolog"
	"github.com/stretchr/testify/assert"
	"testing"
)

const zerologRule = `
rule HelloWorld "Prints hello" {
	when
		true
	then
		Log("zerolog hello from grule");
		Retract("HelloWorld");
}
`

func TestZerologIntegration(t *testing.T) {
	var buf bytes.Buffer

	zl := zerolog.New(&buf).Level(zerolog.InfoLevel).With().Timestamp().Logger()

	engine.SetLogger(&zl)

	dataContext := ast.NewDataContext()
	lib := ast.NewKnowledgeLibrary()
	ruleBuilder := builder.NewRuleBuilder(lib)

	err := ruleBuilder.BuildRuleFromResource("HelloWorld", "0.0.1", pkg.NewBytesResource([]byte(zerologRule)))
	assert.NoError(t, err)

	kb, err := lib.NewKnowledgeBaseInstance("HelloWorld", "0.0.1")
	assert.NoError(t, err)

	grule := engine.NewGruleEngine()
	err = grule.Execute(dataContext, kb)
	assert.NoError(t, err)

}
