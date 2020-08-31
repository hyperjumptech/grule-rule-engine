package ast

import (
	"fmt"
	"github.com/hyperjumptech/grule-rule-engine/model"
	"reflect"
	"strconv"
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

func (tcs *TestCStruct) EchoVariad(ss ...string) int {
	for _, s := range ss {
		fmt.Println(s)
	}
	return len(ss)
}

func TestDataContext_ExecMethod(t *testing.T) {
	TCS := &TestCStruct{
		Str: "",
		It:  0,
	}

	ctx := NewDataContext()
	err := ctx.Add("C", model.NewGoValueNode(reflect.ValueOf(TCS), "C"))
	if err != nil {
		t.Fatal(err)
	}

	_, err = ctx.ExecMethod(reflect.ValueOf(TCS), "EchoMethod", []reflect.Value{reflect.ValueOf("Yahooooo")})
	if err != nil {
		t.Fatal(err)
	}

	_, err = ctx.ExecMethod(reflect.ValueOf(TCS), "EchoMethods", []reflect.Value{reflect.ValueOf("Yahooooo")})
	if err == nil {
		t.Fatal("Error should be raised since method not found")
	}

	_, err = ctx.ExecMethod(reflect.ValueOf(TCS), "EchoMethods", []reflect.Value{reflect.ValueOf(1)})
	if err == nil {
		t.Fatal("Error should be raised since argument type is not string")
	}

	_, err = ctx.ExecMethod(reflect.ValueOf(TCS), "EchoMethods", []reflect.Value{reflect.ValueOf("Yahoooo"), reflect.ValueOf("Google")})
	if err == nil {
		t.Fatal("Error should be raised since argument count is not correct")
	}

	_, err = ctx.ExecMethod(reflect.ValueOf(TCS), "EchoMethods", []reflect.Value{})
	if err == nil {
		t.Fatal("Error should be raised since method argument not provided")
	}

	ctx.Retract("C")
	if len(ctx.Retracted()) == 0 {
		t.Error("Error, Retract failed")
	}
	if !ctx.IsRetracted("C") {
		t.Fatal("Error, should fail")
	}

	_, err = ctx.ExecMethod(reflect.ValueOf(TCS), "MethodNotExist", []reflect.Value{reflect.ValueOf("Yahooooo")})
	if err == nil {
		t.Fatal("Error, method does not exist.")
	}

	ctx.Reset()
	if len(ctx.Retracted()) != 0 {
		t.Error("Error, Reset failed")
	}

	v, err := ctx.ExecMethod(reflect.ValueOf(TCS), "EchoVariad", []reflect.Value{
		reflect.ValueOf("Weeeeee!"),
		reflect.ValueOf("Woooooo!"),
		reflect.ValueOf("Waaaaaa!"),
	})
	if err != nil {
		t.Fatal("Error calling variadic function.")
	}

	t.Logf("Type %s", v.Type().String())

	if v.Int() != 3 {
		t.Fatal("Error, variadic function should have returned 3 but got " + strconv.Itoa(v.Interface().(int)))
	}

	v, err = ctx.ExecMethod(reflect.ValueOf(TCS), "EchoVariad", []reflect.Value{})
	if err != nil {
		t.Fatal("Error calling variadic function.")
	}
	if v.Interface().(int) != 0 {
		t.Fatal("Error, variadic function should have returned 0 but got " + strconv.Itoa(v.Interface().(int)))
	}

	v, err = ctx.ExecMethod(reflect.ValueOf(TCS), "EchoVariad", []reflect.Value{
		reflect.ValueOf("Weeeeee!"),
		reflect.ValueOf(42),
		reflect.ValueOf("Waaaaaa!"),
	})
	if err == nil {
		t.Fatal("Error, calling variadic function with inconsistent args should raise an error")
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

	typ, err := ctx.GetType(reflect.ValueOf(TA.BStruct.CStruct), "Str")
	if err != nil {
		t.Errorf("Got error %v", err)
		t.FailNow()
	} else if typ.Kind() != reflect.String {
		t.Errorf("Not string, but  %s", typ.Kind().String())
		t.FailNow()
	}

	ctx.Retract("ta")
	if len(ctx.Retracted()) == 0 {
		t.Error("Error, Retract failed, it should succeed")
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

	val, err := ctx.GetValue(reflect.ValueOf(TA.BStruct.CStruct), "Str")
	if err != nil {
		t.Errorf("Got error %v", err)
		t.FailNow()
	} else if pkg.ValueToInterface(val).(string) != "TestValue" {
		t.Errorf("Value is not correct")
		t.FailNow()
	}

	ctx.Retract("ta")
	if len(ctx.Retracted()) == 0 {
		t.Error("Error, Retract failed, it should succeed")
	}
}

func TestDataContext_SetValue(t *testing.T) {

	TCS := &TestCStruct{
		Str: "",
		It:  0,
	}

	ctx := NewDataContext()
	err := ctx.Add("C", TCS)
	if err != nil {
		t.Fatal(err)
	}

	err = ctx.SetValue(reflect.ValueOf(TCS), "It", reflect.ValueOf(77))
	if err != nil {
		t.Error("error fail: ", err)
	}

	ctx.Retract("C")
	if len(ctx.Retracted()) == 0 {
		t.Error("Error, Retract failed, it should succeed")
	}
}
