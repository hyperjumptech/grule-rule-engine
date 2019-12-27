package model

import "testing"

func TestMethodCall_EqualsTo(t *testing.T) {
	testData := make([]*AlphaNodeTest, 0)

	testData = append(testData,
		&AlphaNodeTest{
			A: &MethodCall{
				MethodName:      "funcName",
				MethodArguments: &FunctionArgument{Arguments: []*ArgumentHolder{}},
			},
			B: &MethodCall{
				MethodName:      "funcName",
				MethodArguments: &FunctionArgument{Arguments: []*ArgumentHolder{}},
			},
			E: true,
		},
		&AlphaNodeTest{
			A: &MethodCall{
				MethodName:      "funcNameA",
				MethodArguments: &FunctionArgument{Arguments: []*ArgumentHolder{}},
			},
			B: &MethodCall{
				MethodName:      "funcNameB",
				MethodArguments: &FunctionArgument{Arguments: []*ArgumentHolder{}},
			},
			E: false,
		},
		&AlphaNodeTest{
			A: &MethodCall{
				MethodName:      "funcName",
				MethodArguments: &FunctionArgument{Arguments: []*ArgumentHolder{}},
			},
			B: &MethodCall{
				MethodName: "funcName",
				MethodArguments: &FunctionArgument{Arguments: []*ArgumentHolder{&ArgumentHolder{
					Variable: "Some Var",
				}}},
			},
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
