package model

import (
	"reflect"
	"testing"
)

type CompareTest struct {
	A reflect.Value
	B reflect.Value
	E bool
}

func TestConstant_EqualsTo(t *testing.T) {
	tl := make([]*CompareTest, 0)

	tl = append(tl, &CompareTest{
		A: reflect.ValueOf(12),
		B: reflect.ValueOf(12),
		E: true,
	}, &CompareTest{
		A: reflect.ValueOf(12.12),
		B: reflect.ValueOf(12.12),
		E: true,
	}, &CompareTest{
		A: reflect.ValueOf("12"),
		B: reflect.ValueOf("12"),
		E: true,
	}, &CompareTest{
		A: reflect.ValueOf(false),
		B: reflect.ValueOf(false),
		E: true,
	}, &CompareTest{
		A: reflect.ValueOf(byte(12)),
		B: reflect.ValueOf(byte(12)),
		E: true,
	}, &CompareTest{
		A: reflect.ValueOf(uint(12)),
		B: reflect.ValueOf(uint(12)),
		E: true,
	}, &CompareTest{
		A: reflect.ValueOf(12),
		B: reflect.ValueOf(12.0),
		E: false,
	}, &CompareTest{
		A: reflect.ValueOf(12),
		B: reflect.ValueOf("12"),
		E: false,
	}, &CompareTest{
		A: reflect.ValueOf(12),
		B: reflect.ValueOf(uint(12)),
		E: false,
	}, &CompareTest{
		A: reflect.ValueOf(nil),
		B: reflect.ValueOf(uint(12)),
		E: false,
	}, &CompareTest{
		A: reflect.ValueOf(nil),
		B: reflect.ValueOf(nil),
		E: true,
	})

	for i, tls := range tl {
		a := &Constant{ConstantValue: tls.A}
		b := &Constant{ConstantValue: tls.B}
		if a.EqualsTo(b) != tls.E {
			t.Errorf("#%d failed.", i)
			t.Fail()
		}
	}
}
