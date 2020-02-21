package ast

import (
	"github.com/sirupsen/logrus"
	"testing"
)

func TestWorkingMemory_Add(t *testing.T) {
	logrus.SetLevel(logrus.TraceLevel)
	expr1 := &Expression{
		AstID:           "abc",
		LeftExpression:  &Expression{ExpressionAtom: &ExpressionAtom{Variable: &Variable{Name: "some.variable.a"}}},
		RightExpression: &Expression{ExpressionAtom: &ExpressionAtom{Variable: &Variable{Name: "some.variable.b"}}},
		Operator:        OpMul,
	}
	expr2 := &Expression{
		AstID:           "cde",
		LeftExpression:  &Expression{ExpressionAtom: &ExpressionAtom{Variable: &Variable{Name: "some.variable.a"}}},
		RightExpression: &Expression{ExpressionAtom: &ExpressionAtom{Variable: &Variable{Name: "some.variable.b"}}},
		Operator:        OpMul,
	}
	expr3 := &Expression{
		AstID:           "fgh",
		LeftExpression:  &Expression{ExpressionAtom: &ExpressionAtom{Variable: &Variable{Name: "some.variable.c"}}},
		RightExpression: &Expression{ExpressionAtom: &ExpressionAtom{Variable: &Variable{Name: "some.variable.d"}}},
		Operator:        OpMul,
	}
	wm := NewWorkingMemory()
	_, added := wm.Add(expr1)
	if !added {
		t.Fatalf("Expression not added on the first expression")
	}
	_, added = wm.Add(expr2)
	if added {
		t.Fatalf("Expression added on duplicate expression")
	}
	_, added = wm.Add(expr3)
	if !added {
		t.Fatalf("Expression not added on different expression")
	}

	if !wm.IndexVar("some.variable.a") {
		t.Fatalf("Variable not indexed while it exist")
	}
	if wm.IndexVar("some.variable.z") {
		t.Fatalf("Variable indexed while it not exist")
	}

	if !wm.Reset("some.variable.a") {
		t.Fatalf("Variable not reset while it exist")
	}
	if wm.Reset("some.variable.z") {
		t.Fatalf("Variable reset while it not exist")
	}
	if !wm.ResetAll() {
		t.Fatalf("All Variable not reset")
	}
}
