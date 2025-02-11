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

package examples

import (
	"testing"

	"github.com/DataWiseHQ/grule-rule-engine/ast"
	"github.com/DataWiseHQ/grule-rule-engine/engine"

	"github.com/stretchr/testify/assert"
)

func Test_NoPanicOnEmptyKnowledgeBase(t *testing.T) {
	// create a new fact for user
	user := &User{
		Name: "Calo",
		Age:  0,
		Male: true,
	}
	// create an empty data context
	dataContext := ast.NewDataContext()
	// add the fact struct to the data context
	err := dataContext.Add("User", user)
	if err != nil {
		t.Fatal(err)
	}

	t.Run("with nil knowledge base in execute", func(t *testing.T) {
		eng := &engine.GruleEngine{MaxCycle: 10}
		err = eng.Execute(dataContext, nil)

		assert.NotNil(t, err)
	})

	t.Run("with nil knowledge base in FetchMatchingRules", func(t *testing.T) {
		eng := &engine.GruleEngine{MaxCycle: 10}
		_, err = eng.FetchMatchingRules(dataContext, nil)

		assert.NotNil(t, err)
	})
}
