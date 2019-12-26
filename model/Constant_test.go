package model

import (
	"reflect"
	"testing"
)

func TestConstant_EqualsTo(t *testing.T) {
	a := &Constant{ConstantValue: reflect.ValueOf(345)}
	b := &Constant{ConstantValue: reflect.ValueOf(345)}
	if a.EqualsTo(b) == false {
		t.Fail()
	}
	var c AlphaNode
	c = b
	println(reflect.TypeOf(c).Elem().Name())
	d := c.(*Constant).ConstantValue
	println(d.Interface())
}
