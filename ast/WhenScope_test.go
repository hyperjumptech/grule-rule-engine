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
		LeftExpression:  &Expression{ExpressionAtom: &ExpressionAtom{Variable: &Variable{Constant: &Constant{Value: reflect.ValueOf("Whooho")}}}},
		RightExpression: &Expression{ExpressionAtom: &ExpressionAtom{Variable: &Variable{Constant: &Constant{Value: reflect.ValueOf("Whooho")}}}},
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
				Variable: &Variable{
					Constant: &Constant{
						Value: reflect.ValueOf(123),
					},
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
