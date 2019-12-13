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
)

func TestValueAdd(t *testing.T) {
	for _, va := range valuesA {
		for _, vb := range valuesB {
			vc, err := ValueAdd(va, vb)
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
		vd, err := ValueAdd(va, stringVal)
		if err != nil {
			t.Errorf("Error while adding string. Got %v", err)
		} else if vd.Kind() != reflect.String {
			t.Errorf("Add to string should yield a string. Got %s", vd.Kind().String())
		} else if GetBaseKind(va) != reflect.Float64 && vd.String() != "12Text" {
			t.Errorf("Should be \"12Text\". Got \"%s\"", vd.String())
		} else if GetBaseKind(va) == reflect.Float64 && vd.String() != "12.000000Text" {
			t.Errorf("Should be \"12.000000Text\". Got \"%s\"", vd.String())
		}

		vs, err := ValueAdd(stringVal, va)
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
			vc, err := ValueSub(va, vb)
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
		_, err := ValueSub(va, stringVal)
		if err == nil {
			t.Errorf("Subtracting with string should raise an error, but its not.")
		}

	}
}

func TestValueMul(t *testing.T) {
	for _, va := range valuesA {
		for _, vb := range valuesB {
			vc, err := ValueMul(va, vb)
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
		_, err := ValueMul(va, stringVal)
		if err == nil {
			t.Errorf("Multiplication with string should raise an error, but its not.")
		}
	}
}

func TestValueDiv(t *testing.T) {
	for _, va := range valuesA {
		for _, vb := range valuesB {
			vc, err := ValueDiv(va, vb)
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
		_, err := ValueDiv(va, stringVal)
		if err == nil {
			t.Errorf("Division with string should raise an error, but its not.")
		}
	}
}
