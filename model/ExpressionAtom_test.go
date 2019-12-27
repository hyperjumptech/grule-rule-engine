package model

import (
	"reflect"
	"testing"
)

func TestExpressionAtom_EqualsTo(t *testing.T) {
	testData := make([]*AlphaNodeTest, 0)

	testData = append(testData,
		&AlphaNodeTest{
			A: &ExpressionAtom{
				Variable: "AVariable",
			},
			B: &ExpressionAtom{
				Variable: "AVariable",
			},
			E: true,
		},
		&AlphaNodeTest{
			A: &ExpressionAtom{
				Constant: &Constant{ConstantValue: reflect.ValueOf(321)},
			},
			B: &ExpressionAtom{
				Constant: &Constant{ConstantValue: reflect.ValueOf(321)},
			},
			E: true,
		},
		&AlphaNodeTest{
			A: &ExpressionAtom{
				FunctionCall: &FunctionCall{
					FunctionName:      "FuncA",
					FunctionArguments: &FunctionArgument{Arguments: []*ArgumentHolder{}},
				},
			},
			B: &ExpressionAtom{
				FunctionCall: &FunctionCall{
					FunctionName:      "FuncA",
					FunctionArguments: &FunctionArgument{Arguments: []*ArgumentHolder{}},
				},
			},
			E: true,
		},
		&AlphaNodeTest{
			A: &ExpressionAtom{
				MethodCall: &MethodCall{
					MethodName:      "FuncA",
					MethodArguments: &FunctionArgument{Arguments: []*ArgumentHolder{}},
				},
			},
			B: &ExpressionAtom{
				MethodCall: &MethodCall{
					MethodName:      "FuncA",
					MethodArguments: &FunctionArgument{Arguments: []*ArgumentHolder{}},
				},
			},
			E: true,
		},
	)

	for i, data := range testData {
		if data.A.EqualsTo(data.B) != data.E {
			t.Errorf("#%d fail", i)
			t.Fail()
		}
	}
}
