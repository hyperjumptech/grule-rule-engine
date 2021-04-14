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
	"github.com/hyperjumptech/grule-rule-engine/ast"
	"github.com/hyperjumptech/grule-rule-engine/builder"
	"github.com/hyperjumptech/grule-rule-engine/engine"
	"github.com/hyperjumptech/grule-rule-engine/pkg"
	"github.com/stretchr/testify/assert"
	"testing"
)

const (
	Rule4 = `
rule UserTestRule4 "test 3"  salience 10{
	when
	  User.Auth.GetEmail() == "watson@test.com"
	then
	  User.Name = "FromRuleScope4";
	  Retract("UserTestRule4");
}
`
)

type UserWithAuth struct {
	Auth *UserAuth
	Name string
}

func (user *UserWithAuth) GetName() string {
	return user.Name
}

type UserAuth struct {
	Email string
}

func (auth *UserAuth) GetEmail() string {
	return auth.Email
}

func TestMethodCall_Issue4(t *testing.T) {
	user := &UserWithAuth{
		Auth: &UserAuth{Email: "watson@test.com"},
	}

	if user.GetName() != "" {
		t.Fatal("User name not empty")
	}

	dataContext := ast.NewDataContext()
	err := dataContext.Add("User", user)
	assert.NoError(t, err)

	// Prepare knowledgebase library and load it with our rule.
	lib := ast.NewKnowledgeLibrary()
	rb := builder.NewRuleBuilder(lib)
	err = rb.BuildRuleFromResource("Test", "0.1.1", pkg.NewBytesResource([]byte(Rule4)))
	assert.NoError(t, err)
	kb := lib.NewKnowledgeBaseInstance("Test", "0.1.1")
	eng1 := &engine.GruleEngine{MaxCycle: 3}
	err = eng1.Execute(dataContext, kb)
	assert.NoError(t, err)
	assert.Equal(t, "FromRuleScope4", user.GetName())
}
