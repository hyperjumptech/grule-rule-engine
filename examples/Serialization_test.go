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
	"io"
	"os"
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
	if err != nil && err != io.EOF {
		assert.Nil(t, err)
	}

	kb2 := cat2.BuildKnowledgeBase()
	assert.True(t, kb.IsIdentical(kb2))
}

func TestSerializationOnFile(t *testing.T) {
	testFile := "Test.GRB"

	// SAFING INTO FILE
	// First prepare our library for loading the orignal GRL
	lib := ast.NewKnowledgeLibrary()
	rb := builder.NewRuleBuilder(lib)
	err := rb.BuildRuleFromResource("Purchase Calculator", "0.0.1", pkg.NewFileResource("CashFlowRule.grl"))
	assert.NoError(t, err)

	// Lets create the file to write into
	f, err := os.OpenFile(testFile, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	assert.Nil(t, err)

	// Save the knowledge base into the file and close it.
	err = lib.StoreKnowledgeBaseToWriter(f, "load_rules_test", "0.1.1")
	assert.Nil(t, err)
	_ = f.Close()

	// LOADING FROM FILE
	// Lets assume we are using different library, so create a new one
	lib2 := ast.NewKnowledgeLibrary()

	// Open the existing safe file
	f2, err := os.Open(testFile)
	assert.Nil(t, err)

	// Load the file directly into the library and close the file
	kb2, err := lib2.LoadKnowledgeBaseFromReader(f2, true)
	assert.Nil(t, err)
	_ = f2.Close()

	// Delete the test file
	err = os.Remove(testFile)
	assert.Nil(t, err)

	// compare that the original knowledgebase is exacly the same to the loaded one.
	assert.True(t, lib.GetKnowledgeBase("load_rules_test", "0.1.1").IsIdentical(kb2))
}
