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
	"testing"
)

func TestWorkingMemory_Add(t *testing.T) {

	a := &Variable{GrlText: "a", Name: "a"}
	b := &Variable{GrlText: "b", Name: "b"}
	aa := &Variable{GrlText: "a", Name: "a"}
	bb := &Variable{GrlText: "b", Name: "b"}
	c := &Variable{GrlText: "c", Name: "c"}
	d := &Variable{GrlText: "d", Name: "d"}

	expr1 := &Expression{
		AstID:           "abc",
		LeftExpression:  &Expression{ExpressionAtom: &ExpressionAtom{Variable: a}},
		RightExpression: &Expression{ExpressionAtom: &ExpressionAtom{Variable: b}},
		Operator:        OpMul,
	}
	expr2 := &Expression{
		AstID:           "cde",
		LeftExpression:  &Expression{ExpressionAtom: &ExpressionAtom{Variable: aa}},
		RightExpression: &Expression{ExpressionAtom: &ExpressionAtom{Variable: bb}},
		Operator:        OpMul,
	}
	expr3 := &Expression{
		AstID:           "fgh",
		LeftExpression:  &Expression{ExpressionAtom: &ExpressionAtom{Variable: c}},
		RightExpression: &Expression{ExpressionAtom: &ExpressionAtom{Variable: d}},
		Operator:        OpMul,
	}
	wm := NewWorkingMemory("T", "1")
	wm.AddVariable(a)
	wm.AddVariable(b)
	wm.AddVariable(aa)
	wm.AddVariable(bb)
	wm.AddVariable(c)
	wm.AddVariable(d)

	wm.AddExpression(expr1)
	wm.AddExpression(expr2)
	wm.AddExpression(expr3)

	wm.IndexVariables()
	assert.True(t, wm.Reset("a"))
	assert.False(t, wm.Reset("some.variable.z"))
	assert.True(t, wm.ResetAll())
}
