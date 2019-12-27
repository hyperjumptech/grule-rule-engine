package model

import "testing"

func TestFunctionCall_EqualsTo(t *testing.T) {
	testData := make([]*AlphaNodeTest, 0)

	testData = append(testData,
		&AlphaNodeTest{
			A: &FunctionCall{
				FunctionName:      "funcName",
				FunctionArguments: &FunctionArgument{Arguments: []*ArgumentHolder{}},
			},
			B: &FunctionCall{
				FunctionName:      "funcName",
				FunctionArguments: &FunctionArgument{Arguments: []*ArgumentHolder{}},
			},
			E: true,
		},
		&AlphaNodeTest{
			A: &FunctionCall{
				FunctionName:      "funcNameA",
				FunctionArguments: &FunctionArgument{Arguments: []*ArgumentHolder{}},
			},
			B: &FunctionCall{
				FunctionName:      "funcNameB",
				FunctionArguments: &FunctionArgument{Arguments: []*ArgumentHolder{}},
			},
			E: false,
		},
		&AlphaNodeTest{
			A: &FunctionCall{
				FunctionName:      "funcName",
				FunctionArguments: &FunctionArgument{Arguments: []*ArgumentHolder{}},
			},
			B: &FunctionCall{
				FunctionName: "funcName",
				FunctionArguments: &FunctionArgument{Arguments: []*ArgumentHolder{&ArgumentHolder{
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
