//  Copyright DataWiseHQ/grule-rule-engine Authors
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

package builder

import (
	"github.com/DataWiseHQ/grule-rule-engine/ast"
	"github.com/DataWiseHQ/grule-rule-engine/pkg"
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

func TestRuleEntry_Clone(t *testing.T) {
	testRule := `rule  CloneRule  "Duplicate Rule 1"  salience 5 {
when
	(Fact.Distance > 5000  &&   Fact.Duration > 120) && (Fact.Result == false)
Then
   Fact.NetAmount=143.320007;
   Fact.Result=true;
}`
	lib := ast.NewKnowledgeLibrary()
	rb := NewRuleBuilder(lib)
	err := rb.BuildRuleFromResource("testrule", "0.1.1", pkg.NewBytesResource([]byte(testRule)))
	assert.NoError(t, err)
	kb := lib.GetKnowledgeBase("testrule", "0.1.1")
	re := kb.RuleEntries["CloneRule"]

	ct := &pkg.CloneTable{Records: make(map[string]*pkg.CloneRecord)}
	reClone := re.Clone(ct)

	assert.Equal(t, re.GetSnapshot(), reClone.GetSnapshot())
}
