package ast

import (
	"github.com/stretchr/testify/assert"
	"reflect"
	"testing"
)

func TestWorkingMemory_Add(t *testing.T) {

	a := &Variable{GrlText: "a", Constant: &Constant{Value: reflect.ValueOf("a")}}
	b := &Variable{GrlText: "b", Constant: &Constant{Value: reflect.ValueOf("b")}}
	aa := &Variable{GrlText: "a", Constant: &Constant{Value: reflect.ValueOf("a")}}
	bb := &Variable{GrlText: "b", Constant: &Constant{Value: reflect.ValueOf("b")}}
	c := &Variable{GrlText: "c", Constant: &Constant{Value: reflect.ValueOf("c")}}
	d := &Variable{GrlText: "d", Constant: &Constant{Value: reflect.ValueOf("d")}}

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
