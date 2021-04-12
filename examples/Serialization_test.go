//  Copyright hyperjumptech/grule-rule-engine Authors
//
//  Licensed under the Apache License, Version 2.0 (the "License");
//  you may not use this file except in compliance with the License.
//  You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
//  Unless required by applicable law or agreed to in writing, software
//  distributed under the License is distributed on an "AS IS" BASIS,
//  WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//  See the License for the specific language governing permissions and
//  limitations under the License.

package examples

import (
	"bytes"
	"github.com/hyperjumptech/grule-rule-engine/ast"
	"github.com/hyperjumptech/grule-rule-engine/builder"
	"github.com/hyperjumptech/grule-rule-engine/pkg"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSerialization(t *testing.T) {
	lib := ast.NewKnowledgeLibrary()
	rb := builder.NewRuleBuilder(lib)
	err := rb.BuildRuleFromResource("Purchase Calculator", "0.0.1", pkg.NewFileResource("CashFlowRule.grl"))
	assert.NoError(t, err)

	kb := lib.GetKnowledgeBase("Purchase Calculator", "0.0.1")
	cat := kb.MakeCatalog()

	buff1 := &bytes.Buffer{}
	err = cat.WriteCatalogToWriter(buff1)
	assert.Nil(t, err)

	buff2 := bytes.NewBuffer(buff1.Bytes())
	cat2 := &ast.Catalog{}
	err = cat2.ReadCatalogFromReader(buff2)
	assert.Nil(t, err)

	kb2 := cat2.BuildKnowledgeBase()
	assert.True(t, kb.IsIdentical(kb2))
}
