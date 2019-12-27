package model

import (
	"reflect"
	"testing"
)

type AlphaNodeTest struct {
	A AlphaNode
	B AlphaNode
	E bool
}

func TestArgumentHolder_EqualsTo(t *testing.T) {
	testData := make([]*AlphaNodeTest, 0)
	testData = append(testData,
		&AlphaNodeTest{
			A: &ArgumentHolder{
				Variable: "abc.abc.a",
			},
			B: &ArgumentHolder{
				Variable: "abc.abc.a",
			},
			E: true,
		},
		&AlphaNodeTest{
			A: &ArgumentHolder{
				Variable: "abc.abc.a",
			},
			B: &ArgumentHolder{
				Variable: "abc.abc.c",
			},
			E: false,
		},
		&AlphaNodeTest{
			A: &ArgumentHolder{
				Constant: &Constant{ConstantValue: reflect.ValueOf(123)},
			},
			B: &ArgumentHolder{
				Constant: &Constant{ConstantValue: reflect.ValueOf(123)},
			},
			E: true,
		},
		&AlphaNodeTest{
			A: &ArgumentHolder{
				Constant: &Constant{ConstantValue: reflect.ValueOf(123)},
			},
			B: &ArgumentHolder{
				Constant: &Constant{ConstantValue: reflect.ValueOf("123")},
			},
			E: false,
		},
		&AlphaNodeTest{
			A: &ArgumentHolder{
				FunctionCall: &FunctionCall{
					FunctionName: "functionName",
					FunctionArguments: &FunctionArgument{Arguments: []*ArgumentHolder{
						&ArgumentHolder{Variable: "varA"},
					}},
				},
			},
			B: &ArgumentHolder{
				FunctionCall: &FunctionCall{
					FunctionName: "functionName",
					FunctionArguments: &FunctionArgument{Arguments: []*ArgumentHolder{
						&ArgumentHolder{Variable: "varA"},
					}},
				},
			},
			E: true,
		},
		&AlphaNodeTest{
			A: &ArgumentHolder{
				FunctionCall: &FunctionCall{
					FunctionName: "functionName",
					FunctionArguments: &FunctionArgument{Arguments: []*ArgumentHolder{
						&ArgumentHolder{Variable: "varA"},
					}},
				},
			},
			B: &ArgumentHolder{
				FunctionCall: &FunctionCall{
					FunctionName: "functionName",
					FunctionArguments: &FunctionArgument{Arguments: []*ArgumentHolder{
						&ArgumentHolder{Variable: "varB"},
					}},
				},
			},
			E: false,
		},
		&AlphaNodeTest{
			A: &ArgumentHolder{
				MethodCall: &MethodCall{
					MethodName: "functionName",
					MethodArguments: &FunctionArgument{Arguments: []*ArgumentHolder{
						&ArgumentHolder{Variable: "varA"},
					}},
				},
			},
			B: &ArgumentHolder{
				MethodCall: &MethodCall{
					MethodName: "functionName",
					MethodArguments: &FunctionArgument{Arguments: []*ArgumentHolder{
						&ArgumentHolder{Variable: "varA"},
					}},
				},
			},
			E: true,
		},
		&AlphaNodeTest{
			A: &ArgumentHolder{
				MethodCall: &MethodCall{
					MethodName: "functionName",
					MethodArguments: &FunctionArgument{Arguments: []*ArgumentHolder{
						&ArgumentHolder{Variable: "varA"},
					}},
				},
			},
			B: &ArgumentHolder{
				MethodCall: &MethodCall{
					MethodName: "functionName",
					MethodArguments: &FunctionArgument{Arguments: []*ArgumentHolder{
						&ArgumentHolder{Variable: "varB"},
					}},
				},
			},
			E: false,
		},
		&AlphaNodeTest{
			A: &ArgumentHolder{
				Constant: &Constant{ConstantValue: reflect.ValueOf(123)},
			},
			B: &ArgumentHolder{
				Constant: &Constant{ConstantValue: reflect.ValueOf("123")},
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
