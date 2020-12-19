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

package ast

import (
	"github.com/stretchr/testify/assert"
	"reflect"
	"testing"
)

type TestStructShenScope struct {
	StringA string
	StringB string
}

func TestNewWhenScope(t *testing.T) {

	expr1 := &Expression{
		AstID:           "abc",
		LeftExpression:  &Expression{ExpressionAtom: &ExpressionAtom{Constant: &Constant{Value: reflect.ValueOf("Whooho")}}},
		RightExpression: &Expression{ExpressionAtom: &ExpressionAtom{Constant: &Constant{Value: reflect.ValueOf("Whooho")}}},
		Operator:        OpEq,
	}

	ws := NewWhenScope()
	ws.SetGrlText("a == b")
	assert.Equal(t, "a == b", ws.GetGrlText())
	assert.Nil(t, ws.AcceptExpression(expr1), "Error when first time accept expression")
	assert.NotNil(t, ws.AcceptExpression(expr1), "Not Error when second time time accept expression")

	wm := NewWorkingMemory("T", "1")
	dt := NewDataContext()
	test := &TestStructShenScope{
		StringA: "abc",
		StringB: "abc",
	}
	dt.Add("Struct", test)

	t.Logf("%s Snapshot : %s", ws.GetAstID(), ws.GetSnapshot())

	val, err := ws.Evaluate(dt, wm)
	assert.NoError(t, err)
	assert.True(t, val.Bool())
}

func TestNewWhenScopeEvaluate(t *testing.T) {
	expr1 := &Expression{
		AstID: "abc",
		SingleExpression: &Expression{
			ExpressionAtom: &ExpressionAtom{
				Constant: &Constant{
					Value: reflect.ValueOf(123),
				},
			},
		},
	}
	wm := NewWorkingMemory("T", "1")
	dt := NewDataContext()
	val, err := expr1.Evaluate(dt, wm)
	assert.NoError(t, err)
	assert.Equal(t, 123, int(val.Int()))

	ws := NewWhenScope()
	assert.Nil(t, ws.AcceptExpression(expr1))
	val, err = ws.Evaluate(dt, wm)
	assert.NoError(t, err)
	assert.Equal(t, 123, int(val.Int()))

}
