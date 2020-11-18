package pkg

import (
	"fmt"
	"reflect"
	"testing"
	"time"
)

type TestObject struct {
	A string
	B int
	C float64
	D bool
	E time.Time
	F *TestSubObject
	G int8
	H int16
	I int32
	J int64
	K uint
	L uint8
	M uint16
	N uint32
	O uint64
	P float32
	Q []int
	R map[int]string
	S TestSubObjectNoPtr
}

func TestSetAttributeTimeValue(t *testing.T) {
	to := &TestObject{
		E: time.Date(2002, 2, 2, 2, 2, 2, 2, time.Local),
	}

	vto := reflect.ValueOf(to)
	if vto.Kind() != reflect.Ptr {
		t.FailNow()
	}
	eVal := vto.Elem().FieldByName("E")
	if eVal.Kind() != reflect.Struct || eVal.Type().String() != "time.Time" {
		t.FailNow()
	}
	if !eVal.CanSet() {
		t.FailNow()
	}

	newVal := reflect.ValueOf(time.Date(2003, 3, 3, 3, 3, 3, 3, time.Local))
	eVal.Set(newVal)
	ti := ValueToInterface(eVal).(time.Time)
	fmt.Printf("Time is %v\n", ti)
}

func (to *TestObject) FunctionA(arg1 string, arg2 string) (string, error) {
	return fmt.Sprintf("A call Arg1 : %s and Arg2 : %s", arg1, arg2), nil
}

func (to *TestObject) FunctionB(arg1, arg2 string) (string, error) {
	return fmt.Sprintf("B call Arg1 : %s and Arg2 : %s", arg1, arg2), nil
}

func (to *TestObject) FunctionC(arg1 int, arg2 string) (string, error) {
	return fmt.Sprintf("C call Arg1 : %d and Arg2 : %s", arg1, arg2), nil
}

func (to TestObject) FunctionD(arg1 string, arg2 string) (string, error) {
	return fmt.Sprintf("A call Arg1 : %s and Arg2 : %s", arg1, arg2), nil
}

func (to TestObject) FunctionE(arg1, arg2 string) (string, error) {
	return fmt.Sprintf("B call Arg1 : %s and Arg2 : %s", arg1, arg2), nil
}

func (to TestObject) FunctionF(arg1 int, arg2 string) (string, error) {
	return fmt.Sprintf("C call Arg1 : %d and Arg2 : %s", arg1, arg2), nil
}

type TestSubObject struct {
	A string
	B int
	C float64
}

type TestObjectNoPtr struct {
	A string
	B int
	C float64
	D bool
	E time.Time
	F TestSubObjectNoPtr
}

type TestSubObjectNoPtr struct {
	A string
	B int
	C float64
}

func TestGetFunctionList(t *testing.T) {
	to := reflect.ValueOf(&TestObject{})
	functions, err := GetFunctionList(to)
	if err != nil {
		t.Error("Got error")
		t.FailNow()
	}
	if len(functions) != 6 {
		t.Errorf("Got %d functions", len(functions))
		t.FailNow()
	}
	if functions[0] != "FunctionA" {
		t.Errorf("1st function name %s", functions[0])
		t.FailNow()
	}
	if functions[1] != "FunctionB" {
		t.Errorf("2nd function name %s", functions[1])
		t.FailNow()
	}

	to2 := reflect.ValueOf(TestObject{})
	functions, err = GetFunctionList(to2)
	if err != nil {
		t.Error("Got error")
		t.FailNow()
	}
	if len(functions) != 3 {
		t.Errorf("Got %d functions", len(functions))
		t.FailNow()
	}
	if functions[0] != "FunctionD" {
		t.Errorf("1st function name %s", functions[0])
		t.FailNow()
	}
	if functions[1] != "FunctionE" {
		t.Errorf("2nd function name %s", functions[1])
		t.FailNow()
	}
}

func TestGetFunctionParameterTypes(t *testing.T) {
	to := reflect.ValueOf(&TestObject{})
	types, _, err := GetFunctionParameterTypes(to, "FunctionC")
	if err != nil {
		t.Errorf("Error : %v", err)
		t.FailNow()
	}
	if len(types) != 2 {
		t.Errorf("Invalid argument count : %d", len(types))
		for idx, typ := range types {
			if typ.Kind().String() == "ptr" {
				t.Errorf("#%d : %s", idx, typ.Elem().Name())
			} else {
				t.Errorf("#%d : %s", idx, typ.Kind().String())
			}
		}
		t.FailNow()
	}
	if types[0].Kind().String() != "int" || types[1].Kind().String() != "string" {
		t.Errorf("Invalid argument type kind")
		t.FailNow()
	}
}

func TestGetFunctionReturnTypes(t *testing.T) {
	to := reflect.ValueOf(&TestObject{})
	types, err := GetFunctionReturnTypes(to, "FunctionC")
	if err != nil {
		t.Errorf("Error : %v", err)
		t.FailNow()
	}
	if len(types) != 2 {
		t.Errorf("Invalid return count : %d", len(types))
		for idx, typ := range types {
			if typ.Kind().String() == "ptr" {
				t.Errorf("#%d : %s", idx, typ.Elem().Name())
			} else {
				t.Errorf("#%d : %s", idx, typ.Kind().String())
			}
		}
		t.FailNow()
	}
	if types[0].Kind().String() != "string" || types[1].Kind().String() != "interface" {
		t.Errorf("Invalid argument type kind")
		for idx, typ := range types {
			if typ.Kind().String() == "ptr" {
				t.Errorf("#%d : %s", idx, typ.Elem().Name())
			} else {
				t.Errorf("#%d : %s", idx, typ.Kind().String())
			}
		}
		t.FailNow()
	}
}

func TestInvokeFunction(t *testing.T) {
	to := reflect.ValueOf(&TestObject{})
	param := make([]reflect.Value, 2)
	param[0] = reflect.ValueOf(10)
	param[1] = reflect.ValueOf("Ten")
	rets, err := InvokeFunction(to, "FunctionC", param)
	if err != nil {
		t.Errorf("Got error : %v", err)
	} else {
		if len(rets) != 2 {
			t.Errorf("Invalid ret outs : %d", len(rets))
		}
		if rets[0].String() != "C call Arg1 : 10 and Arg2 : Ten" {
			t.Errorf("Invalid turn : %s", rets[0].String())
		}
		if !rets[1].IsValid() {
			t.Errorf("2nd return should be valid")
		}
	}
}

func TestGetAttributeList(t *testing.T) {
	to := reflect.ValueOf(&TestObject{})
	names, err := GetAttributeList(to)
	if err != nil {
		t.Errorf("Got error : %v", err)
		t.FailNow()
	}
	if len(names) != 19 {
		t.Errorf("Invalid field count : %d", len(names))
		t.FailNow()
	}
	check := []string{
		"A", "B", "C", "D", "E", "F", "G", "H", "I", "J", "K", "L", "M", "N", "O", "P", "Q", "R", "S",
	}
	for i := 0; i < len(check); i++ {
		if names[i] != check[i] {
			t.Errorf("Attribute #%d, expect %s, but actual %s", i, check[i], names[i])
			t.FailNow()
		}
	}
}

func TestGetAttributeValue(t *testing.T) {
	to := reflect.ValueOf(TestObjectNoPtr{
		A: "string data",
		B: 123,
		C: 456.789,
		D: true,
		E: time.Date(2019, time.January, 1, 1, 1, 1, 0, time.Local),
		F: TestSubObjectNoPtr{
			A: "string",
			B: 123,
			C: 456.789,
		},
	})
	itv, err := GetAttributeInterface(to, "A")
	if err != nil {
		t.Errorf("Got error %v", err)
		t.FailNow()
	}
	if itv.(string) != "string data" {
		t.FailNow()
	}
	tim, err := GetAttributeValue(to, "E")
	if err != nil {
		t.Errorf("Got error %v", err)
		t.FailNow()
	}
	if tim.Type().String() != "time.Time" {
		t.Errorf("Not time")
		t.FailNow()
	}
}

func TestSetAttributeValue(t *testing.T) {
	to := &TestObject{
		A: "string data",
		B: 123,
		C: 456.789,
		D: true,
	}
	testObject := reflect.ValueOf(to)
	err := SetAttributeInterface(testObject, "A", "strong data")
	if err != nil {
		t.Errorf("Got error %v", err)
		t.FailNow()
	}
	err = SetAttributeInterface(testObject, "B", 456)
	if err != nil {
		t.Errorf("Got error %v", err)
		t.FailNow()
	}
	err = SetAttributeInterface(testObject, "B", 456.123)
	if err != nil {
		t.Errorf("Should be able to set with different type as long as between them are numbers")
		t.FailNow()
	}
	if to.A != "strong data" && to.B != 456 {
		t.Errorf("Setting string fail")
		t.FailNow()
	}

	tso := &TestSubObject{
		A: "TSO",
		B: 2019,
		C: 2019.6,
	}
	err = SetAttributeInterface(testObject, "F", tso)
	if err != nil {
		t.Errorf("Should not be able to set with different type : %v", err)
		t.FailNow()
	}
	if to.F.A != "TSO" && to.F.B != 2019 {
		t.Errorf("Setting object fail")
		t.FailNow()
	}
}

func TestValueToInterface(t *testing.T) {
	if ValueToInterface(reflect.ValueOf(int(100))).(int) != 100 {
		t.Error("Failed to reflect value of int")
		t.Fail()
	}

	if ValueToInterface(reflect.ValueOf(byte(100))).(byte) != byte(100) {
		t.Error("Failed to reflect value of byte")
		t.Fail()
	}
	if ValueToInterface(reflect.ValueOf(int8(100))).(int8) != int8(100) {
		t.Error("Failed to reflect value of int8")
		t.Fail()
	}
	if ValueToInterface(reflect.ValueOf(int16(100))).(int16) != int16(100) {
		t.Error("Failed to reflect value of int16")
		t.Fail()
	}
	if ValueToInterface(reflect.ValueOf(int32(100))).(int32) != int32(100) {
		t.Error("Failed to reflect value of int32")
		t.Fail()
	}
	if ValueToInterface(reflect.ValueOf(int64(100))).(int64) != int64(100) {
		t.Error("Failed to reflect value of int64")
		t.Fail()
	}
	if ValueToInterface(reflect.ValueOf(uint(100))).(uint) != uint(100) {
		t.Error("Failed to reflect value of int")
		t.Fail()
	}
	if ValueToInterface(reflect.ValueOf(uint8(100))).(uint8) != uint8(100) {
		t.Error("Failed to reflect value of uint8")
		t.Fail()
	}
	if ValueToInterface(reflect.ValueOf(uint16(100))).(uint16) != uint16(100) {
		t.Error("Failed to reflect value of uint16")
	}
	if ValueToInterface(reflect.ValueOf(uint32(100))).(uint32) != uint32(100) {
		t.Error("Failed to reflect value of uint32")
	}
	if ValueToInterface(reflect.ValueOf(uint64(100))).(uint64) != uint64(100) {
		t.Error("Failed to reflect value of uint64")
	}
	if ValueToInterface(reflect.ValueOf(float32(100))).(float32) != float32(100) {
		t.Error("Failed to reflect value of float32")
		t.Fail()
	}
	if ValueToInterface(reflect.ValueOf(float64(100))).(float64) != float64(100) {
		t.Error("Failed to reflect value of float64")
		t.Fail()
	}
	if ValueToInterface(reflect.ValueOf("Some string value")).(string) != "Some string value" {
		t.Error("Failed to reflect value of float64")
		t.Fail()
	}
}

type TestStruct struct {
	ArrayAttribute []int
	MapAttribute   map[string]int
	IntType        int
	Int8Type       int8
	Int16Type      int16
	Int32Type      int32
	Int64Type      int64
	UIntType       uint
	UInt8Type      uint8
	UInt16Type     uint16
	UInt32Type     uint32
	UInt64Type     uint64
	Float32Type    float32
	Float64Type    float64
	BoolType       bool
	StringType     string
}

func TestIsAttributeArray(t *testing.T) {
	bol, err := IsAttributeArray(reflect.ValueOf(12), "something")
	if err == nil {
		t.Error("error should be raised, obj is not struct")
		t.Fail()
	}
	ts := &TestStruct{}
	bol, err = IsAttributeArray(reflect.ValueOf(ts), "ArrayAttribute")
	if err != nil {
		t.Error(err)
		t.Fail()
	}
	if bol == false {
		t.Error("Attribute is an array")
		t.Fail()
	}
}

func TestIsAttributeMap(t *testing.T) {
	bol, err := IsAttributeMap(reflect.ValueOf(12), "something")
	if err == nil {
		t.Error("error should be raised, obj is not struct")
		t.Fail()
	}
	ts := &TestStruct{}
	bol, err = IsAttributeMap(reflect.ValueOf(ts), "MapAttribute")
	if err != nil {
		t.Error(err)
		t.Fail()
	}
	if bol == false {
		t.Error("Attribute is a map")
		t.Fail()
	}
}

func TestIsAttributeNilOrZero(t *testing.T) {
	_, err := IsAttributeNilOrZero(reflect.ValueOf(12), "something")
	if err == nil {
		t.Error("error should be raised, obj is not struct")
		t.Fail()
	}
	ts := &TestStruct{}
	bol, err := IsAttributeNilOrZero(reflect.ValueOf(ts), "ArrayAttribute")
	if err != nil {
		t.Error(err)
		t.Fail()
	}
	if !bol {
		t.Error("It should be nil or zero")
		t.Fail()
	}
}

func TestGetMapArrayValue(t *testing.T) {
	amap := make(map[string]string)
	amap["abc"] = "ABC"

	ret, err := GetMapArrayValue(amap, "abc")
	if err != nil {
		t.Errorf("got %s", err.Error())
		t.Fail()
	} else {
		if ret.(string) != "ABC" {
			t.Errorf("expect ABC but %s", ret.(string))
			t.Fail()
		}
	}
	ret, err = GetMapArrayValue(amap, "cba")
	if err == nil {
		t.Errorf("key not exist but no error")
		t.Fail()
	}
	t.Logf("Emitted err : %s", err.Error())
	ret, err = GetMapArrayValue(amap, 123)
	if err == nil {
		t.Errorf("key different type but no error")
		t.Fail()
	}
	t.Logf("Emitted err : %s", err.Error())

	aarr := make([]string, 0)
	aarr = append(aarr, "ABC")
	ret, err = GetMapArrayValue(aarr, 0)
	if err != nil {
		t.Errorf("got %s", err.Error())
		t.Fail()
	} else {
		if ret.(string) != "ABC" {
			t.Errorf("expect ABC but %s", ret.(string))
			t.Fail()
		}
	}
	ret, err = GetMapArrayValue(aarr, 3)
	if err == nil {
		t.Errorf("key out of bound but no error")
		t.Fail()
	}
	t.Logf("Emitted err : %s", err.Error())
	ret, err = GetMapArrayValue(aarr, uint(0))
	if err != nil {
		t.Errorf("key out of bound but no error")
		t.Fail()
	}
}
