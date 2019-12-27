package model

import (
	"reflect"
	"testing"
)

func TestFunctionArgument_EqualsTo(t *testing.T) {
	testData := make([]*AlphaNodeTest, 0)
	testData = append(testData,
		&AlphaNodeTest{
			A: &FunctionArgument{Arguments: []*ArgumentHolder{
				&ArgumentHolder{
					Constant: &Constant{ConstantValue: reflect.ValueOf("123")},
				},
				&ArgumentHolder{
					Constant: &Constant{ConstantValue: reflect.ValueOf(123)},
				},
			}},
			B: &FunctionArgument{Arguments: []*ArgumentHolder{
				&ArgumentHolder{
					Constant: &Constant{ConstantValue: reflect.ValueOf("123")},
				},
				&ArgumentHolder{
					Constant: &Constant{ConstantValue: reflect.ValueOf(123)},
				},
			}},
			E: true,
		},
		&AlphaNodeTest{
			A: &FunctionArgument{Arguments: []*ArgumentHolder{
				&ArgumentHolder{
					Constant: &Constant{ConstantValue: reflect.ValueOf("123")},
				},
				&ArgumentHolder{
					Constant: &Constant{ConstantValue: reflect.ValueOf(123)},
				},
			}},
			B: &FunctionArgument{Arguments: []*ArgumentHolder{
				&ArgumentHolder{
					Constant: &Constant{ConstantValue: reflect.ValueOf("123")},
				},
			}},
			E: false,
		},
		&AlphaNodeTest{
			A: &FunctionArgument{Arguments: []*ArgumentHolder{
				&ArgumentHolder{
					Constant: &Constant{ConstantValue: reflect.ValueOf("123")},
				},
				&ArgumentHolder{
					Constant: &Constant{ConstantValue: reflect.ValueOf(123)},
				},
			}},
			B: &FunctionArgument{Arguments: []*ArgumentHolder{
				&ArgumentHolder{
					Constant: &Constant{ConstantValue: reflect.ValueOf(123)},
				},
				&ArgumentHolder{
					Constant: &Constant{ConstantValue: reflect.ValueOf("123")},
				},
			}},
			E: false,
		},
	)
	for i, data := range testData {
		if data.A.EqualsTo(data.B) != data.E {
			t.Errorf("#%d fail", i)
			t.Fail()
		}
	}
}
