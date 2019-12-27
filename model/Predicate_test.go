package model

import (
	"reflect"
	"testing"
)

func TestPredicate_EqualsTo(t *testing.T) {
	testData := make([]*AlphaNodeTest, 0)

	testData = append(testData, &AlphaNodeTest{
		A: &Predicate{
			ExpressionAtomLeft: &ExpressionAtom{
				Variable: "aVar",
			},
			ExpressionAtomRight: &ExpressionAtom{
				Variable: "bVar",
			},
			ComparisonOperator: ComparisonOperatorEQ,
		},
		B: &Predicate{
			ExpressionAtomLeft: &ExpressionAtom{
				Variable: "bVar",
			},
			ExpressionAtomRight: &ExpressionAtom{
				Variable: "aVar",
			},
			ComparisonOperator: ComparisonOperatorEQ,
		},
		E: true,
	},
		&AlphaNodeTest{
			A: &Predicate{
				ExpressionAtomLeft: &ExpressionAtom{
					Variable: "aVar",
				},
				ExpressionAtomRight: &ExpressionAtom{
					Constant: &Constant{ConstantValue: reflect.ValueOf(123)},
				},
				ComparisonOperator: ComparisonOperatorEQ,
			},
			B: &Predicate{
				ExpressionAtomLeft: &ExpressionAtom{
					Constant: &Constant{ConstantValue: reflect.ValueOf(123)},
				},
				ExpressionAtomRight: &ExpressionAtom{
					Variable: "aVar",
				},
				ComparisonOperator: ComparisonOperatorEQ,
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
