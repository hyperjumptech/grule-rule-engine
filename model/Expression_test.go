package model

import (
	"reflect"
	"testing"
)

func TestExpression_EqualsTo(t *testing.T) {
	testData := make([]*AlphaNodeTest, 0)

	testData = append(testData,
		&AlphaNodeTest{
			A: &Expression{
				LeftExpression: &Expression{
					Predicate: &Predicate{
						ExpressionAtomLeft: &ExpressionAtom{
							Variable: "varname",
						},
						ExpressionAtomRight: &ExpressionAtom{
							Constant: &Constant{ConstantValue: reflect.ValueOf(123)},
						},
						ComparisonOperator: ComparisonOperatorEQ,
					},
				},
				RightExpression: &Expression{
					Predicate: &Predicate{
						ExpressionAtomLeft: &ExpressionAtom{
							Variable: "varname_two",
						},
						ExpressionAtomRight: &ExpressionAtom{
							Constant: &Constant{ConstantValue: reflect.ValueOf(321)},
						},
						ComparisonOperator: ComparisonOperatorEQ,
					},
				},
				LogicalOperator: LogicalOperatorAnd,
			},
			B: &Expression{
				LeftExpression: &Expression{
					Predicate: &Predicate{
						ExpressionAtomLeft: &ExpressionAtom{
							Variable: "varname",
						},
						ExpressionAtomRight: &ExpressionAtom{
							Constant: &Constant{ConstantValue: reflect.ValueOf(123)},
						},
						ComparisonOperator: ComparisonOperatorEQ,
					},
				},
				RightExpression: &Expression{
					Predicate: &Predicate{
						ExpressionAtomLeft: &ExpressionAtom{
							Variable: "varname_two",
						},
						ExpressionAtomRight: &ExpressionAtom{
							Constant: &Constant{ConstantValue: reflect.ValueOf(321)},
						},
						ComparisonOperator: ComparisonOperatorEQ,
					},
				},
				LogicalOperator: LogicalOperatorAnd,
			},
			E: true,
		},
		&AlphaNodeTest{
			A: &Expression{
				LeftExpression: &Expression{
					Predicate: &Predicate{
						ExpressionAtomLeft: &ExpressionAtom{
							Variable: "varname",
						},
						ExpressionAtomRight: &ExpressionAtom{
							Constant: &Constant{ConstantValue: reflect.ValueOf(123)},
						},
						ComparisonOperator: ComparisonOperatorEQ,
					},
				},
				RightExpression: &Expression{
					Predicate: &Predicate{
						ExpressionAtomLeft: &ExpressionAtom{
							Variable: "varname_two",
						},
						ExpressionAtomRight: &ExpressionAtom{
							Constant: &Constant{ConstantValue: reflect.ValueOf(321)},
						},
						ComparisonOperator: ComparisonOperatorEQ,
					},
				},
				LogicalOperator: LogicalOperatorAnd,
			},
			B: &Expression{
				RightExpression: &Expression{
					Predicate: &Predicate{
						ExpressionAtomLeft: &ExpressionAtom{
							Variable: "varname",
						},
						ExpressionAtomRight: &ExpressionAtom{
							Constant: &Constant{ConstantValue: reflect.ValueOf(123)},
						},
						ComparisonOperator: ComparisonOperatorEQ,
					},
				},
				LeftExpression: &Expression{
					Predicate: &Predicate{
						ExpressionAtomLeft: &ExpressionAtom{
							Variable: "varname_two",
						},
						ExpressionAtomRight: &ExpressionAtom{
							Constant: &Constant{ConstantValue: reflect.ValueOf(321)},
						},
						ComparisonOperator: ComparisonOperatorEQ,
					},
				},
				LogicalOperator: LogicalOperatorAnd,
			},
			E: true,
		},
		&AlphaNodeTest{
			A: &Expression{Predicate: &Predicate{
				ExpressionAtomLeft: &ExpressionAtom{Variable: "VarA"},
			}},
			B: &Expression{Predicate: &Predicate{
				ExpressionAtomLeft: &ExpressionAtom{Variable: "VarA"},
			}},
			E: true,
		},
		&AlphaNodeTest{
			A: &Expression{Predicate: &Predicate{
				ExpressionAtomLeft: &ExpressionAtom{Variable: "VarA"},
			}},
			B: &Expression{Predicate: &Predicate{
				ExpressionAtomLeft: &ExpressionAtom{Variable: "VarB"},
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
