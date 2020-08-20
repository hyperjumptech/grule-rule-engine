package ast

import (
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
	if ws.GetGrlText() != "a == b" {
		t.Fatalf("GRL text not equal")
	}

	if ws.AcceptExpression(expr1) != nil {
		t.Fatalf("Error when first time accept expression")
	}

	if ws.AcceptExpression(expr1) == nil {
		t.Fatalf("Not Error when second time time accept expression")
	}

	wm := NewWorkingMemory("T", "1")
	dt := NewDataContext()
	test := &TestStructShenScope{
		StringA: "abc",
		StringB: "abc",
	}
	dt.Add("Struct", test)

	if ws.DataContext != nil {
		t.Fatalf("datacontext not nil")
	}
	if ws.WorkingMemory != nil {
		t.Fatalf("working memory not nil")
	}

	ws.InitializeContext(dt, wm)

	if ws.DataContext == nil {
		t.Fatalf("datacontext nil")
	}
	if ws.WorkingMemory == nil {
		t.Fatalf("working memory nil")
	}

	t.Logf("%s Snapshot : %s", ws.GetAstID(), ws.GetSnapshot())

	val, err := ws.Evaluate()
	if err != nil {
		t.Fatalf("error while evaluating constant expression. got %s", err)
	}
	if !val.Bool() {
		t.Fatalf("Value is not as expected.")
	}

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
	val, err := expr1.Evaluate()
	if err != nil {
		t.Fatalf("error while evaluating constant expression")
	}
	if val.Int() != 123 {
		t.Fatalf("Value is not as expected. %d", val.Int())
	}

	ws := NewWhenScope()
	if ws.AcceptExpression(expr1) != nil {
		t.Fatalf("error when accepting expression first time")
	}
	val, err = ws.Evaluate()
	if err != nil {
		t.Fatalf("error while evaluating constant expression")
	}
	if val.Int() != 123 {
		t.Fatalf("Value is not as expected. %d", val.Int())
	}

}
