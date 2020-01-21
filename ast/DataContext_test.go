package ast

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/hyperjumptech/grule-rule-engine/pkg"
)

type TestAStruct struct {
	BStruct *TestBStruct
}

type TestBStruct struct {
	CStruct *TestCStruct
}

type TestCStruct struct {
	Str string
	It  int
}

func (tcs *TestCStruct) EchoMethod(s string) {
	fmt.Println(s)
}

func TestDataContext_ExecMethod(t *testing.T) {
	TCS := &TestCStruct{
		Str: "",
		It:  0,
	}

	ctx := NewDataContext()
	err := ctx.Add("C", TCS)
	if err != nil {
		t.Fatal(err)
	}

	_, err = ctx.ExecMethod("C.EchoMethod", []reflect.Value{reflect.ValueOf("Yahooooo")})
	if err != nil {
		t.Fatal(err)
	}

	_, err = ctx.ExecMethod("C.EchoMethods", []reflect.Value{reflect.ValueOf("Yahooooo")})
	if err == nil {
		t.Fatal("Error should be raised since method not found")
	}

	_, err = ctx.ExecMethod("C.EchoMethods", []reflect.Value{reflect.ValueOf(1)})
	if err == nil {
		t.Fatal("Error should be raised since argument type is not string")
	}

	_, err = ctx.ExecMethod("C.EchoMethods", []reflect.Value{reflect.ValueOf("Yahoooo"), reflect.ValueOf("Google")})
	if err == nil {
		t.Fatal("Error should be raised since argument count is not correct")
	}

	_, err = ctx.ExecMethod("C.EchoMethods", []reflect.Value{})
	if err == nil {
		t.Fatal("Error should be raised since method argument not provided")
	}
}

func TestDataContext_GetType(t *testing.T) {
	TA := &TestAStruct{BStruct: &TestBStruct{CStruct: &TestCStruct{
		Str: "TestValue",
		It:  100,
	}}}

	ctx := NewDataContext()
	err := ctx.Add("ta", TA)
	if err != nil {
		t.Fatal(err)
	}

	typ, err := ctx.GetType("ta.BStruct.CStruct.Str")
	if err != nil {
		t.Errorf("Got error %v", err)
		t.FailNow()
	} else if typ.Kind() != reflect.String {
		t.Errorf("Not string, but  %s", typ.Kind().String())
		t.FailNow()
	}
}

func TestDataContext_GetValue(t *testing.T) {
	TA := &TestAStruct{BStruct: &TestBStruct{CStruct: &TestCStruct{
		Str: "TestValue",
		It:  100,
	}}}

	ctx := NewDataContext()
	err := ctx.Add("ta", TA)
	if err != nil {
		t.Fatal(err)
	}

	val, err := ctx.GetValue("ta.BStruct.CStruct.Str")
	if err != nil {
		t.Errorf("Got error %v", err)
		t.FailNow()
	} else if pkg.ValueToInterface(val).(string) != "TestValue" {
		t.Errorf("Value is not correct")
		t.FailNow()
	}

}
