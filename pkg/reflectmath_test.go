package pkg

import (
	"reflect"
	"testing"
)

var (
	intVal   = reflect.ValueOf(12)
	int8Val  = reflect.ValueOf(int8(12))
	int16Val = reflect.ValueOf(int16(12))
	int32Val = reflect.ValueOf(int32(12))
	int64Val = reflect.ValueOf(int64(12))

	uintVal   = reflect.ValueOf(uint(12))
	uint8Val  = reflect.ValueOf(uint8(12))
	uint16Val = reflect.ValueOf(uint16(12))
	uint32Val = reflect.ValueOf(uint32(12))
	uint64Val = reflect.ValueOf(uint64(12))

	float32Val = reflect.ValueOf(float32(12))
	float64Val = reflect.ValueOf(float64(12))

	intVal2   = reflect.ValueOf(3)
	int8Val2  = reflect.ValueOf(int8(3))
	int16Val2 = reflect.ValueOf(int16(3))
	int32Val2 = reflect.ValueOf(int32(3))
	int64Val2 = reflect.ValueOf(int64(3))

	uintVal2   = reflect.ValueOf(uint(3))
	uint8Val2  = reflect.ValueOf(uint8(3))
	uint16Val2 = reflect.ValueOf(uint16(3))
	uint32Val2 = reflect.ValueOf(uint32(3))
	uint64Val2 = reflect.ValueOf(uint64(3))

	float32Val2 = reflect.ValueOf(float32(3))
	float64Val2 = reflect.ValueOf(float64(3))

	valuesA = []reflect.Value{intVal, int8Val, int16Val, int32Val, int64Val, uintVal, uint8Val, uint16Val, uint32Val, uint64Val, float32Val, float64Val}
	valuesB = []reflect.Value{intVal2, int8Val2, int16Val2, int32Val2, int64Val2, uintVal2, uint8Val2, uint16Val2, uint32Val2, uint64Val2, float32Val2, float64Val2}

	intVal3   = reflect.ValueOf(0x55)
	int8Val3  = reflect.ValueOf(int8(0x55))
	int16Val3 = reflect.ValueOf(int16(0x55))
	int32Val3 = reflect.ValueOf(int32(0x55))
	int64Val3 = reflect.ValueOf(int64(0x55))

	uintVal3   = reflect.ValueOf(uint(0x55))
	uint8Val3  = reflect.ValueOf(uint8(0x55))
	uint16Val3 = reflect.ValueOf(uint16(0x55))
	uint32Val3 = reflect.ValueOf(uint32(0x55))
	uint64Val3 = reflect.ValueOf(uint64(0x55))

	intVal4   = reflect.ValueOf(0x7F)
	int8Val4  = reflect.ValueOf(int8(0x7F))
	int16Val4 = reflect.ValueOf(int16(0x7F))
	int32Val4 = reflect.ValueOf(int32(0x7F))
	int64Val4 = reflect.ValueOf(int64(0x7F))

	uintVal4   = reflect.ValueOf(uint(0x7F))
	uint8Val4  = reflect.ValueOf(uint8(0x7F))
	uint16Val4 = reflect.ValueOf(uint16(0x7F))
	uint32Val4 = reflect.ValueOf(uint32(0x7F))
	uint64Val4 = reflect.ValueOf(uint64(0x7F))

	StrCompareTest = []*StrCompare{
		{"A", "A", true, false, false, true, false, true},
		{"AA", "A", false, true, true, true, false, false},
		{"A", "AA", false, true, false, false, true, true},
		{" ", "  ", false, true, false, false, true, true},
		{" ", "A", false, true, false, false, true, true},
		{"A", " ", false, true, true, true, false, false},
		{"A", "aa", false, true, false, false, true, true},
		{"aa", "A", false, true, true, true, false, false},
		{"a", "AA", false, true, true, true, false, false},
		{"AA", "a", false, true, false, false, true, true},
	}
)

type StrCompare struct {
	A   string
	B   string
	Eq  bool
	Neq bool
	Gt  bool
	Gte bool
	Lt  bool
	Lte bool
}

func TestStringComparison(t *testing.T) {
	for i, v := range StrCompareTest {
		val, err := EvaluateEqual(reflect.ValueOf(v.A), reflect.ValueOf(v.B))
		if err != nil {
			t.Errorf(err.Error())
			t.Fail()
		} else if val.Bool() != v.Eq {
			t.Errorf("%d Expect \"%s\" and \"%s\" EQ expect %v but %v", i, v.A, v.B, v.Eq, !v.Eq)
		}

		val, err = EvaluateNotEqual(reflect.ValueOf(v.A), reflect.ValueOf(v.B))
		if err != nil {
			t.Errorf(err.Error())
			t.Fail()
		} else if val.Bool() != v.Neq {
			t.Errorf("%d Expect \"%s\" and \"%s\" NEQ expect %v but %v", i, v.A, v.B, v.Neq, !v.Neq)
		}

		val, err = EvaluateGreaterThan(reflect.ValueOf(v.A), reflect.ValueOf(v.B))
		if err != nil {
			t.Errorf(err.Error())
			t.Fail()
		} else if val.Bool() != v.Gt {
			t.Errorf("%d Expect \"%s\" and \"%s\" GT expect %v but %v", i, v.A, v.B, v.Gt, !v.Gt)
		}

		val, err = EvaluateGreaterThanEqual(reflect.ValueOf(v.A), reflect.ValueOf(v.B))
		if err != nil {
			t.Errorf(err.Error())
			t.Fail()
		} else if val.Bool() != v.Gte {
			t.Errorf("%d Expect \"%s\" and \"%s\" GTE expect %v but %v", i, v.A, v.B, v.Gte, !v.Gte)
		}

		val, err = EvaluateLesserThan(reflect.ValueOf(v.A), reflect.ValueOf(v.B))
		if err != nil {
			t.Errorf(err.Error())
			t.Fail()
		} else if val.Bool() != v.Lt {
			t.Errorf("%d Expect \"%s\" and \"%s\" LT expect %v but %v", i, v.A, v.B, v.Lt, !v.Lt)
		}

		val, err = EvaluateLesserThanEqual(reflect.ValueOf(v.A), reflect.ValueOf(v.B))
		if err != nil {
			t.Errorf(err.Error())
			t.Fail()
		} else if val.Bool() != v.Lte {
			t.Errorf("%d Expect \"%s\" and \"%s\" LTE expect %v but %v", i, v.A, v.B, v.Lte, !v.Lte)
		}
	}
}

func TestValueAdd(t *testing.T) {
	for _, va := range valuesA {
		for _, vb := range valuesB {
			vc, err := EvaluateAddition(va, vb)
			if err != nil {
				t.Errorf("Error %v", err)
			}
			//t.Logf("%s + %s = %s", va.Kind().String(), vb.Kind().String(), vc.Kind().String())
			if vc.Kind() == reflect.Uint64 {
				if vc.Uint() != 15 {
					t.Errorf("Expected uint 10 but %d", vc.Uint())
				}
			} else if vc.Kind() == reflect.Int64 {
				if vc.Int() != 15 {
					t.Errorf("Expected int 10 but %d", vc.Int())
				}
			} else if vc.Kind() == reflect.Float64 {
				if vc.Float() != 15 {
					t.Errorf("Expected float 10 but %d", vc.Int())
				}
			} else {
				t.Errorf("Math Add expect number types return")
			}
			if GetBaseKind(va) == reflect.Float64 || GetBaseKind(vb) == reflect.Float64 {
				if vc.Kind() != reflect.Float64 {
					t.Errorf("Any Add to float should yield Float64, but %s", vc.Kind().String())
				}
			} else if GetBaseKind(va) == reflect.Int64 || GetBaseKind(vb) == reflect.Int64 {
				if vc.Kind() != reflect.Int64 {
					t.Errorf("Any Add to int should yield int64, but %s", vc.Kind().String())
				}
			} else {
				if vc.Kind() != reflect.Uint64 {
					t.Errorf("The rest should be uint64, but %s", vc.Kind().String())
				}
			}
		}
		stringVal := reflect.ValueOf("Text")
		vd, err := EvaluateAddition(va, stringVal)
		if err != nil {
			t.Errorf("Error while adding string. Got %v", err)
		} else if vd.Kind() != reflect.String {
			t.Errorf("Add to string should yield a string. Got %s", vd.Kind().String())
		} else if GetBaseKind(va) != reflect.Float64 && vd.String() != "12Text" {
			t.Errorf("Should be \"12Text\". Got \"%s\"", vd.String())
		} else if GetBaseKind(va) == reflect.Float64 && vd.String() != "12.000000Text" {
			t.Errorf("Should be \"12.000000Text\". Got \"%s\"", vd.String())
		}

		vs, err := EvaluateAddition(stringVal, va)
		if err != nil {
			t.Errorf("Error while adding string. Got %v", err)
		} else if vs.Kind() != reflect.String {
			t.Errorf("Add to string should yield a string. Got %s", vd.Kind().String())
		} else if GetBaseKind(va) != reflect.Float64 && vs.String() != "Text12" {
			t.Errorf("Should be \"Text12\". Got \"%s\"", vd.String())
		} else if GetBaseKind(va) == reflect.Float64 && vs.String() != "Text12.000000" {
			t.Errorf("Should be \"Text12.000000\". Got \"%s\"", vd.String())
		}
	}
}

func TestValueSub(t *testing.T) {
	for _, va := range valuesA {
		for _, vb := range valuesB {
			vc, err := EvaluateSubtraction(va, vb)
			if err != nil {
				t.Errorf("Error %v", err)
			}
			//t.Logf("%s + %s = %s", va.Kind().String(), vb.Kind().String(), vc.Kind().String())
			if vc.Kind() == reflect.Uint64 {
				if vc.Uint() != 9 {
					t.Errorf("Expected uint 10 but %d", vc.Uint())
				}
			} else if vc.Kind() == reflect.Int64 {
				if vc.Int() != 9 {
					t.Errorf("Expected int 10 but %d", vc.Int())
				}
			} else if vc.Kind() == reflect.Float64 {
				if vc.Float() != 9 {
					t.Errorf("Expected float 10 but %d", vc.Int())
				}
			} else {
				t.Errorf("Math Sub expect number types return")
			}
			if GetBaseKind(va) == reflect.Float64 || GetBaseKind(vb) == reflect.Float64 {
				if vc.Kind() != reflect.Float64 {
					t.Errorf("Any Sub to float should yield Float64, but %s", vc.Kind().String())
				}
			} else if GetBaseKind(va) == reflect.Int64 || GetBaseKind(vb) == reflect.Int64 {
				if vc.Kind() != reflect.Int64 {
					t.Errorf("Any Sub to int should yield int64, but %s", vc.Kind().String())
				}
			} else {
				if vc.Kind() != reflect.Uint64 {
					t.Errorf("The rest should be uint64, but %s", vc.Kind().String())
				}
			}
		}
		stringVal := reflect.ValueOf("Text")
		_, err := EvaluateSubtraction(va, stringVal)
		if err == nil {
			t.Errorf("Subtracting with string should raise an error, but its not.")
		}

	}
}

func TestValueMul(t *testing.T) {
	for _, va := range valuesA {
		for _, vb := range valuesB {
			vc, err := EvaluateMultiplication(va, vb)
			if err != nil {
				t.Errorf("Error %v", err)
			}
			//t.Logf("%s + %s = %s", va.Kind().String(), vb.Kind().String(), vc.Kind().String())
			if vc.Kind() == reflect.Uint64 {
				if vc.Uint() != 36 {
					t.Errorf("Expected uint 10 but %d", vc.Uint())
				}
			} else if vc.Kind() == reflect.Int64 {
				if vc.Int() != 36 {
					t.Errorf("Expected int 10 but %d", vc.Int())
				}
			} else if vc.Kind() == reflect.Float64 {
				if vc.Float() != 36 {
					t.Errorf("Expected float 10 but %d", vc.Int())
				}
			} else {
				t.Errorf("Math Mul expect number types return")
			}
			if GetBaseKind(va) == reflect.Float64 || GetBaseKind(vb) == reflect.Float64 {
				if vc.Kind() != reflect.Float64 {
					t.Errorf("Any Mul to float should yield Float64, but %s", vc.Kind().String())
				}
			} else if GetBaseKind(va) == reflect.Int64 || GetBaseKind(vb) == reflect.Int64 {
				if vc.Kind() != reflect.Int64 {
					t.Errorf("Any Mul to int should yield int64, but %s", vc.Kind().String())
				}
			} else {
				if vc.Kind() != reflect.Uint64 {
					t.Errorf("The rest should be uint64, but %s", vc.Kind().String())
				}
			}
		}

		stringVal := reflect.ValueOf("Text")
		_, err := EvaluateMultiplication(va, stringVal)
		if err == nil {
			t.Errorf("Multiplication with string should raise an error, but its not.")
		}
	}
}

func TestValueDiv(t *testing.T) {
	for _, va := range valuesA {
		for _, vb := range valuesB {
			vc, err := EvaluateDivision(va, vb)
			if err != nil {
				t.Errorf("Error %v", err)
			}
			//t.Logf("%s + %s = %s", va.Kind().String(), vb.Kind().String(), vc.Kind().String())
			if vc.Kind() == reflect.Uint64 {
				if vc.Uint() != 4 {
					t.Errorf("Expected uint 10 but %d", vc.Uint())
				}
			} else if vc.Kind() == reflect.Int64 {
				if vc.Int() != 4 {
					t.Errorf("Expected int 10 but %d", vc.Int())
				}
			} else if vc.Kind() == reflect.Float64 {
				if vc.Float() != 4 {
					t.Errorf("Expected float 10 but %d", vc.Int())
				}
			} else {
				t.Errorf("Math div expect number types return")
			}
			if GetBaseKind(va) == reflect.Float64 || GetBaseKind(vb) == reflect.Float64 {
				if vc.Kind() != reflect.Float64 {
					t.Errorf("Any Div to float should yield Float64, but %s", vc.Kind().String())
				}
			} else if GetBaseKind(va) == reflect.Int64 || GetBaseKind(vb) == reflect.Int64 {
				if vc.Kind() != reflect.Float64 {
					t.Errorf("Any Div to int should yield int64, but %s", vc.Kind().String())
				}
			} else {
				if vc.Kind() != reflect.Float64 {
					t.Errorf("The rest should be float64, but %s", vc.Kind().String())
				}
			}
		}
		stringVal := reflect.ValueOf("Text")
		_, err := EvaluateDivision(va, stringVal)
		if err == nil {
			t.Errorf("Division with string should raise an error, but its not.")
		}
	}
}

func TestEvaluateModulo(t *testing.T) {
	valuesMA := []reflect.Value{intVal, int8Val, int16Val, int32Val, int64Val, uintVal, uint8Val, uint16Val, uint32Val, uint64Val}
	valuesMB := []reflect.Value{intVal2, int8Val2, int16Val2, int32Val2, int64Val2, uintVal2, uint8Val2, uint16Val2, uint32Val2, uint64Val2}

	for _, va := range valuesMA {
		for _, vb := range valuesMB {
			vc, err := EvaluateModulo(va, vb)

			if err != nil {
				t.Errorf("Error %v", err)
			}

			if vc.Int() != 0 {
				t.Errorf("12 mod 3 not 0")
			}
		}
	}
}

func TestEvaluateBitAnd(t *testing.T) {
	valuesMA := []reflect.Value{intVal3, int8Val3, int16Val3, int32Val3, int64Val3, uintVal3, uint8Val3, uint16Val3, uint32Val3, uint64Val3}
	valuesMB := []reflect.Value{intVal4, int8Val4, int16Val4, int32Val4, int64Val4, uintVal4, uint8Val4, uint16Val4, uint32Val4, uint64Val4}

	for _, va := range valuesMA {
		for _, vb := range valuesMB {
			vc, err := EvaluateBitAnd(va, vb)

			if err != nil {
				t.Errorf("Error %v", err)
			}

			if vc.Kind() == reflect.Uint64 && vc.Uint() != 0x55 {
				t.Errorf("0x55 & 0x7F != 0x55 but %d", vc.Uint())
			}
			if vc.Kind() == reflect.Int64 && vc.Int() != 0x55 {
				t.Errorf("0x55 & 0x7F != 0x55 but %d", vc.Int())
			}
		}
	}
}

func TestEvaluateBitOr(t *testing.T) {
	valuesMA := []reflect.Value{intVal3, int8Val3, int16Val3, int32Val3, int64Val3, uintVal3, uint8Val3, uint16Val3, uint32Val3, uint64Val3}
	valuesMB := []reflect.Value{intVal4, int8Val4, int16Val4, int32Val4, int64Val4, uintVal4, uint8Val4, uint16Val4, uint32Val4, uint64Val4}

	for _, va := range valuesMA {
		for _, vb := range valuesMB {
			vc, err := EvaluateBitOr(va, vb)

			if err != nil {
				t.Errorf("Error %v", err)
			}

			if vc.Kind() == reflect.Uint64 && vc.Uint() != 0x7F {
				t.Errorf("0x55 & 0x7F != 0x7F but %d", vc.Uint())
			}
			if vc.Kind() == reflect.Int64 && vc.Int() != 0x7F {
				t.Errorf("0x55 & 0x7F != 0x7F but %d", vc.Int())
			}
		}
	}
}

func TestEvaluateGreaterThan(t *testing.T) {
	for _, va := range valuesA {
		for _, vb := range valuesB {
			vc, err := EvaluateGreaterThan(va, vb)

			if err != nil {
				t.Errorf("Error %v", err)
			}

			if vc.Kind() != reflect.Bool || vc.Bool() == false {
				t.Errorf("12 > 3 == false")
			}
		}
	}
}

func TestEvaluateLesserThan(t *testing.T) {
	for _, va := range valuesA {
		for _, vb := range valuesB {
			vc, err := EvaluateLesserThan(vb, va)

			if err != nil {
				t.Errorf("Error %v", err)
			}

			if vc.Kind() != reflect.Bool || vc.Bool() == false {
				t.Errorf("3 < 12 == false")
			}
		}
	}
}

func TestEvaluateGreaterThanEqual(t *testing.T) {
	for _, va := range valuesA {
		for _, vb := range valuesB {
			vc, err := EvaluateGreaterThanEqual(va, vb)

			if err != nil {
				t.Errorf("Error %v", err)
			}

			if vc.Kind() != reflect.Bool || vc.Bool() == false {
				t.Errorf("12 >= 3 == false")
			}
		}
	}
}

func TestEvaluateLesserThanEqual(t *testing.T) {
	for _, va := range valuesA {
		for _, vb := range valuesB {
			vc, err := EvaluateLesserThanEqual(vb, va)

			if err != nil {
				t.Errorf("Error %v", err)
			}

			if vc.Kind() != reflect.Bool || vc.Bool() == false {
				t.Errorf("3 <= 12 == false")
			}
		}
	}
}

func TestEvaluateEqual(t *testing.T) {
	for _, va := range valuesA {
		for _, vb := range valuesA {
			vc, err := EvaluateEqual(vb, va)

			if err != nil {
				t.Errorf("Error %v", err)
			}

			if vc.Kind() != reflect.Bool || vc.Bool() == false {
				t.Errorf("3 == 3 == false")
			}
		}
	}
}

func TestEvaluateNotEqual(t *testing.T) {
	for _, va := range valuesA {
		for _, vb := range valuesB {
			vc, err := EvaluateNotEqual(vb, va)

			if err != nil {
				t.Errorf("Error %v", err)
			}

			if vc.Kind() != reflect.Bool || vc.Bool() == false {
				t.Errorf("3 != 12 == false")
			}
		}
	}
}
