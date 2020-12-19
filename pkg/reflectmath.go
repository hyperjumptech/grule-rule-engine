//  Copyright hyperjumptech/grule-rule-engine Authors
//
//  Licensed under the Apache License, Version 2.0 (the "License");
//  you may not use this file except in compliance with the License.
//  You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
//  Unless required by applicable law or agreed to in writing, software
//  distributed under the License is distributed on an "AS IS" BASIS,
//  WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//  See the License for the specific language governing permissions and
//  limitations under the License.

package pkg

import (
	"fmt"
	"reflect"
	"time"
)

// EvaluateMultiplication will evaluate multiplication operation over two value
func EvaluateMultiplication(left, right reflect.Value) (reflect.Value, error) {
	switch left.Kind() {
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		lv := left.Int()
		switch right.Kind() {
		case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
			rv := right.Int()
			return reflect.ValueOf(lv * rv), nil
		case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
			rv := right.Uint()
			return reflect.ValueOf(lv * int64(rv)), nil
		case reflect.Float32, reflect.Float64:
			rv := right.Float()
			return reflect.ValueOf(float64(lv) * rv), nil
		default:
			return reflect.ValueOf(nil), fmt.Errorf("can not multiply data type of %s", right.Kind().String())
		}
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		lv := left.Uint()
		switch right.Kind() {
		case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
			rv := right.Int()
			return reflect.ValueOf(int64(lv) * rv), nil
		case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
			rv := right.Uint()
			return reflect.ValueOf(lv * rv), nil
		case reflect.Float32, reflect.Float64:
			rv := right.Float()
			return reflect.ValueOf(float64(lv) * rv), nil
		default:
			return reflect.ValueOf(nil), fmt.Errorf("can not multiply data type of %s", right.Kind().String())
		}
	case reflect.Float32, reflect.Float64:
		lv := left.Float()
		switch right.Kind() {
		case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
			rv := right.Int()
			return reflect.ValueOf(lv * float64(rv)), nil
		case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
			rv := right.Uint()
			return reflect.ValueOf(lv * float64(rv)), nil
		case reflect.Float32, reflect.Float64:
			rv := right.Float()
			return reflect.ValueOf(lv * rv), nil
		default:
			return reflect.ValueOf(nil), fmt.Errorf("can not multiply data type of %s", right.Kind().String())
		}
	default:
		return reflect.ValueOf(nil), fmt.Errorf("can not multiply data type of %s", left.Kind().String())
	}
}

// EvaluateDivision will evaluate division operation over two value
func EvaluateDivision(left, right reflect.Value) (reflect.Value, error) {
	switch left.Kind() {
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		lv := left.Int()
		switch right.Kind() {
		case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
			rv := right.Int()
			return reflect.ValueOf(float64(lv) / float64(rv)), nil
		case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
			rv := right.Uint()
			return reflect.ValueOf(float64(lv) / float64(rv)), nil
		case reflect.Float32, reflect.Float64:
			rv := right.Float()
			return reflect.ValueOf(float64(lv) / rv), nil
		default:
			return reflect.ValueOf(nil), fmt.Errorf("can not use data type of %s in division", right.Kind().String())
		}
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		lv := left.Uint()
		switch right.Kind() {
		case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
			rv := right.Int()
			return reflect.ValueOf(float64(lv) / float64(rv)), nil
		case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
			rv := right.Uint()
			return reflect.ValueOf(float64(lv) / float64(rv)), nil
		case reflect.Float32, reflect.Float64:
			rv := right.Float()
			return reflect.ValueOf(float64(lv) / rv), nil
		default:
			return reflect.ValueOf(nil), fmt.Errorf("can not use data type of %s in division", right.Kind().String())
		}
	case reflect.Float32, reflect.Float64:
		lv := left.Float()
		switch right.Kind() {
		case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
			rv := right.Int()
			return reflect.ValueOf(lv / float64(rv)), nil
		case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
			rv := right.Uint()
			return reflect.ValueOf(lv / float64(rv)), nil
		case reflect.Float32, reflect.Float64:
			rv := right.Float()
			return reflect.ValueOf(lv / rv), nil
		default:
			return reflect.ValueOf(nil), fmt.Errorf("can not use data type of %s in division", right.Kind().String())
		}
	default:
		return reflect.ValueOf(nil), fmt.Errorf("can not use data type of %s in division", left.Kind().String())
	}
}

// EvaluateModulo will evaluate modulo operation over two value
func EvaluateModulo(left, right reflect.Value) (reflect.Value, error) {
	switch left.Kind() {
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		lv := left.Int()
		switch right.Kind() {
		case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
			rv := right.Int()
			return reflect.ValueOf(lv % rv), nil
		case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
			rv := right.Uint()
			return reflect.ValueOf(lv % int64(rv)), nil
		default:
			return reflect.ValueOf(nil), fmt.Errorf("can not use data type of %s in modulo", right.Kind().String())
		}
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		lv := left.Uint()
		switch right.Kind() {
		case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
			rv := right.Int()
			return reflect.ValueOf(int64(lv) % rv), nil
		case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
			rv := right.Uint()
			return reflect.ValueOf(int64(lv) % int64(rv)), nil
		default:
			return reflect.ValueOf(nil), fmt.Errorf("can not use data type of %s in modulo", right.Kind().String())
		}
	default:
		return reflect.ValueOf(nil), fmt.Errorf("can not use data type of %s in modulo", left.Kind().String())
	}
}

// EvaluateAddition will evaluate addition operation over two value
func EvaluateAddition(left, right reflect.Value) (reflect.Value, error) {
	switch left.Kind() {
	case reflect.String:
		lv := left.String()
		switch right.Kind() {
		case reflect.String:
			rv := right.String()
			return reflect.ValueOf(fmt.Sprintf("%s%s", lv, rv)), nil
		case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
			rv := right.Int()
			return reflect.ValueOf(fmt.Sprintf("%s%d", lv, rv)), nil
		case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
			rv := right.Uint()
			return reflect.ValueOf(fmt.Sprintf("%s%d", lv, rv)), nil
		case reflect.Float32, reflect.Float64:
			rv := right.Float()
			return reflect.ValueOf(fmt.Sprintf("%s%f", lv, rv)), nil
		case reflect.Bool:
			rv := right.Bool()
			return reflect.ValueOf(fmt.Sprintf("%s%v", lv, rv)), nil
		default:
			if right.Type().String() == "time.Time" {
				rv := right.Interface().(time.Time)
				return reflect.ValueOf(fmt.Sprintf("%s%s", lv, rv.Format(time.RFC3339))), nil
			}
			return reflect.ValueOf(nil), fmt.Errorf("can not use data type of %s in addition", right.Kind().String())
		}
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		lv := left.Int()
		switch right.Kind() {
		case reflect.String:
			rv := right.String()
			return reflect.ValueOf(fmt.Sprintf("%d%s", lv, rv)), nil
		case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
			rv := right.Int()
			return reflect.ValueOf(lv + rv), nil
		case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
			rv := right.Uint()
			return reflect.ValueOf(lv + int64(rv)), nil
		case reflect.Float32, reflect.Float64:
			rv := right.Float()
			return reflect.ValueOf(float64(lv) + rv), nil
		default:
			return reflect.ValueOf(nil), fmt.Errorf("can not use data type of %s in addition", right.Kind().String())
		}
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		lv := left.Uint()
		switch right.Kind() {
		case reflect.String:
			rv := right.String()
			return reflect.ValueOf(fmt.Sprintf("%d%s", lv, rv)), nil
		case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
			rv := right.Int()
			return reflect.ValueOf(int64(lv) + rv), nil
		case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
			rv := right.Uint()
			return reflect.ValueOf(lv + rv), nil
		case reflect.Float32, reflect.Float64:
			rv := right.Float()
			return reflect.ValueOf(float64(lv) + rv), nil
		default:
			return reflect.ValueOf(nil), fmt.Errorf("can not use data type of %s in division", right.Kind().String())
		}
	case reflect.Float32, reflect.Float64:
		lv := left.Float()
		switch right.Kind() {
		case reflect.String:
			rv := right.String()
			return reflect.ValueOf(fmt.Sprintf("%f%s", lv, rv)), nil
		case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
			rv := right.Int()
			return reflect.ValueOf(lv + float64(rv)), nil
		case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
			rv := right.Uint()
			return reflect.ValueOf(lv + float64(rv)), nil
		case reflect.Float32, reflect.Float64:
			rv := right.Float()
			return reflect.ValueOf(lv + rv), nil
		default:
			return reflect.ValueOf(nil), fmt.Errorf("can not use data type of %s in division", right.Kind().String())
		}
	default:
		return reflect.ValueOf(nil), fmt.Errorf("can not use data type of %s in division", left.Kind().String())
	}
}

// EvaluateSubtraction will evaluate subtraction operation over two value
func EvaluateSubtraction(left, right reflect.Value) (reflect.Value, error) {
	switch left.Kind() {
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		lv := left.Int()
		switch right.Kind() {
		case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
			rv := right.Int()
			return reflect.ValueOf(lv - rv), nil
		case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
			rv := right.Uint()
			return reflect.ValueOf(lv - int64(rv)), nil
		case reflect.Float32, reflect.Float64:
			rv := right.Float()
			return reflect.ValueOf(float64(lv) - rv), nil
		default:
			return reflect.ValueOf(nil), fmt.Errorf("can not use data type of %s in subtraction", right.Kind().String())
		}
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		lv := left.Uint()
		switch right.Kind() {
		case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
			rv := right.Int()
			return reflect.ValueOf(int64(lv) - rv), nil
		case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
			rv := right.Uint()
			return reflect.ValueOf(lv - rv), nil
		case reflect.Float32, reflect.Float64:
			rv := right.Float()
			return reflect.ValueOf(float64(lv) - rv), nil
		default:
			return reflect.ValueOf(nil), fmt.Errorf("can not use data type of %s in subtraction", right.Kind().String())
		}
	case reflect.Float32, reflect.Float64:
		lv := left.Float()
		switch right.Kind() {
		case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
			rv := right.Int()
			return reflect.ValueOf(lv - float64(rv)), nil
		case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
			rv := right.Uint()
			return reflect.ValueOf(lv - float64(rv)), nil
		case reflect.Float32, reflect.Float64:
			rv := right.Float()
			return reflect.ValueOf(lv - rv), nil
		default:
			return reflect.ValueOf(nil), fmt.Errorf("can not use data type of %s in subtraction", right.Kind().String())
		}
	default:
		return reflect.ValueOf(nil), fmt.Errorf("can not use data type of %s in subtraction", left.Kind().String())
	}
}

// EvaluateBitAnd will evaluate Bitwise And operation over two value
func EvaluateBitAnd(left, right reflect.Value) (reflect.Value, error) {
	switch left.Kind() {
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		lv := left.Int()
		switch right.Kind() {
		case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
			rv := right.Int()
			return reflect.ValueOf(lv & rv), nil
		case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
			rv := right.Uint()
			return reflect.ValueOf(lv & int64(rv)), nil
		default:
			return reflect.ValueOf(nil), fmt.Errorf("can not use data type of %s in bitwise AND operation", right.Kind().String())
		}
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		lv := left.Uint()
		switch right.Kind() {
		case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
			rv := right.Int()
			return reflect.ValueOf(int64(lv) & rv), nil
		case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
			rv := right.Uint()
			return reflect.ValueOf(lv & rv), nil
		default:
			return reflect.ValueOf(nil), fmt.Errorf("can not use data type of %s in bitwise AND operation", right.Kind().String())
		}
	default:
		return reflect.ValueOf(nil), fmt.Errorf("can not use data type of %s in bitwise AND operation", left.Kind().String())
	}
}

// EvaluateBitOr will evaluate Bitwise Or operation over two value
func EvaluateBitOr(left, right reflect.Value) (reflect.Value, error) {
	switch left.Kind() {
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		lv := left.Int()
		switch right.Kind() {
		case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
			rv := right.Int()
			return reflect.ValueOf(lv | rv), nil
		case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
			rv := right.Uint()
			return reflect.ValueOf(lv | int64(rv)), nil
		default:
			return reflect.ValueOf(nil), fmt.Errorf("can not use data type of %s in bitwise OR operation", right.Kind().String())
		}
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		lv := left.Uint()
		switch right.Kind() {
		case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
			rv := right.Int()
			return reflect.ValueOf(int64(lv) | rv), nil
		case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
			rv := right.Uint()
			return reflect.ValueOf(lv | rv), nil
		default:
			return reflect.ValueOf(nil), fmt.Errorf("can not use data type of %s in bitwise OR operation", right.Kind().String())
		}
	default:
		return reflect.ValueOf(nil), fmt.Errorf("can not use data type of %s in bitwise OR operation", left.Kind().String())
	}
}

// EvaluateGreaterThan will evaluate GreaterThan operation over two value
func EvaluateGreaterThan(left, right reflect.Value) (reflect.Value, error) {
	switch left.Kind() {
	case reflect.String:
		lv := left.String()
		switch right.Kind() {
		case reflect.String:
			rv := right.String()
			return reflect.ValueOf(lv > rv), nil
		default:
			return reflect.ValueOf(nil), fmt.Errorf("can not compare data type of string to %s in GT comparison", right.Kind().String())
		}
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		lv := left.Int()
		switch right.Kind() {
		case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
			rv := right.Int()
			return reflect.ValueOf(lv > rv), nil
		case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
			rv := right.Uint()
			return reflect.ValueOf(lv > int64(rv)), nil
		case reflect.Float32, reflect.Float64:
			rv := right.Float()
			return reflect.ValueOf(float64(lv) > rv), nil
		default:
			return reflect.ValueOf(nil), fmt.Errorf("can not use data type of %s in GT comparison", right.Kind().String())
		}
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		lv := left.Uint()
		switch right.Kind() {
		case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
			rv := right.Int()
			return reflect.ValueOf(int64(lv) > rv), nil
		case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
			rv := right.Uint()
			return reflect.ValueOf(lv > rv), nil
		case reflect.Float32, reflect.Float64:
			rv := right.Float()
			return reflect.ValueOf(float64(lv) > rv), nil
		default:
			return reflect.ValueOf(nil), fmt.Errorf("can not use data type of %s in GT comparison", right.Kind().String())
		}
	case reflect.Float32, reflect.Float64:
		lv := left.Float()
		switch right.Kind() {
		case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
			rv := right.Int()
			return reflect.ValueOf(lv > float64(rv)), nil
		case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
			rv := right.Uint()
			return reflect.ValueOf(lv > float64(rv)), nil
		case reflect.Float32, reflect.Float64:
			rv := right.Float()
			return reflect.ValueOf(lv > rv), nil
		default:
			return reflect.ValueOf(nil), fmt.Errorf("can not use data type of %s in GT comparison", right.Kind().String())
		}
	default:
		if left.Type().String() == "time.Time" && right.Type().String() == "time.Time" {
			lv := left.Interface().(time.Time)
			rv := right.Interface().(time.Time)
			return reflect.ValueOf(lv.After(rv)), nil
		}
		return reflect.ValueOf(nil), fmt.Errorf("can not use data type of %s in GT comparison", left.Kind().String())
	}
}

// EvaluateLesserThan will evaluate LesserThan operation over two value
func EvaluateLesserThan(left, right reflect.Value) (reflect.Value, error) {
	switch left.Kind() {
	case reflect.String:
		lv := left.String()
		switch right.Kind() {
		case reflect.String:
			rv := right.String()
			return reflect.ValueOf(lv < rv), nil
		default:
			return reflect.ValueOf(nil), fmt.Errorf("can not compare data type of string to %s in LT comparison", right.Kind().String())
		}
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		lv := left.Int()
		switch right.Kind() {
		case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
			rv := right.Int()
			return reflect.ValueOf(lv < rv), nil
		case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
			rv := right.Uint()
			return reflect.ValueOf(lv < int64(rv)), nil
		case reflect.Float32, reflect.Float64:
			rv := right.Float()
			return reflect.ValueOf(float64(lv) < rv), nil
		default:
			return reflect.ValueOf(nil), fmt.Errorf("can not use data type of %s in LT comparison", right.Kind().String())
		}
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		lv := left.Uint()
		switch right.Kind() {
		case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
			rv := right.Int()
			return reflect.ValueOf(int64(lv) < rv), nil
		case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
			rv := right.Uint()
			return reflect.ValueOf(lv < rv), nil
		case reflect.Float32, reflect.Float64:
			rv := right.Float()
			return reflect.ValueOf(float64(lv) < rv), nil
		default:
			return reflect.ValueOf(nil), fmt.Errorf("can not use data type of %s in LT comparison", right.Kind().String())
		}
	case reflect.Float32, reflect.Float64:
		lv := left.Float()
		switch right.Kind() {
		case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
			rv := right.Int()
			return reflect.ValueOf(lv < float64(rv)), nil
		case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
			rv := right.Uint()
			return reflect.ValueOf(lv < float64(rv)), nil
		case reflect.Float32, reflect.Float64:
			rv := right.Float()
			return reflect.ValueOf(lv < rv), nil
		default:
			return reflect.ValueOf(nil), fmt.Errorf("can not use data type of %s in LT comparison", right.Kind().String())
		}
	default:
		if left.Type().String() == "time.Time" && right.Type().String() == "time.Time" {
			lv := left.Interface().(time.Time)
			rv := right.Interface().(time.Time)
			return reflect.ValueOf(lv.Before(rv)), nil
		}
		return reflect.ValueOf(nil), fmt.Errorf("can not use data type of %s in LT comparison", left.Kind().String())
	}
}

// EvaluateGreaterThanEqual will evaluate GreaterThanEqual operation over two value
func EvaluateGreaterThanEqual(left, right reflect.Value) (reflect.Value, error) {
	switch left.Kind() {
	case reflect.String:
		lv := left.String()
		switch right.Kind() {
		case reflect.String:
			rv := right.String()
			return reflect.ValueOf(lv >= rv), nil
		default:
			return reflect.ValueOf(nil), fmt.Errorf("can not compare data type of string to %s in GTE comparison", right.Kind().String())
		}
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		lv := left.Int()
		switch right.Kind() {
		case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
			rv := right.Int()
			return reflect.ValueOf(lv >= rv), nil
		case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
			rv := right.Uint()
			return reflect.ValueOf(lv >= int64(rv)), nil
		case reflect.Float32, reflect.Float64:
			rv := right.Float()
			return reflect.ValueOf(float64(lv) >= rv), nil
		default:
			return reflect.ValueOf(nil), fmt.Errorf("can not use data type of %s in GTE comparison", right.Kind().String())
		}
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		lv := left.Uint()
		switch right.Kind() {
		case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
			rv := right.Int()
			return reflect.ValueOf(int64(lv) >= rv), nil
		case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
			rv := right.Uint()
			return reflect.ValueOf(lv >= rv), nil
		case reflect.Float32, reflect.Float64:
			rv := right.Float()
			return reflect.ValueOf(float64(lv) >= rv), nil
		default:
			return reflect.ValueOf(nil), fmt.Errorf("can not use data type of %s in GTE comparison", right.Kind().String())
		}
	case reflect.Float32, reflect.Float64:
		lv := left.Float()
		switch right.Kind() {
		case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
			rv := right.Int()
			return reflect.ValueOf(lv >= float64(rv)), nil
		case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
			rv := right.Uint()
			return reflect.ValueOf(lv >= float64(rv)), nil
		case reflect.Float32, reflect.Float64:
			rv := right.Float()
			return reflect.ValueOf(lv >= rv), nil
		default:
			return reflect.ValueOf(nil), fmt.Errorf("can not use data type of %s in GTE comparison", right.Kind().String())
		}
	default:
		if left.Type().String() == "time.Time" && right.Type().String() == "time.Time" {
			lv := left.Interface().(time.Time)
			rv := right.Interface().(time.Time)
			return reflect.ValueOf(lv.After(rv) || lv == rv), nil
		}
		return reflect.ValueOf(nil), fmt.Errorf("can not use data type of %s in GTE comparison", left.Kind().String())
	}
}

// EvaluateLesserThanEqual will evaluate LesserThanEqual operation over two value
func EvaluateLesserThanEqual(left, right reflect.Value) (reflect.Value, error) {
	switch left.Kind() {
	case reflect.String:
		lv := left.String()
		switch right.Kind() {
		case reflect.String:
			rv := right.String()
			return reflect.ValueOf(lv <= rv), nil
		default:
			return reflect.ValueOf(nil), fmt.Errorf("can not compare data type of string to %s in LTE comparison", right.Kind().String())
		}
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		lv := left.Int()
		switch right.Kind() {
		case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
			rv := right.Int()
			return reflect.ValueOf(lv <= rv), nil
		case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
			rv := right.Uint()
			return reflect.ValueOf(lv <= int64(rv)), nil
		case reflect.Float32, reflect.Float64:
			rv := right.Float()
			return reflect.ValueOf(float64(lv) <= rv), nil
		default:
			return reflect.ValueOf(nil), fmt.Errorf("can not use data type of %s in LTE comparison", right.Kind().String())
		}
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		lv := left.Uint()
		switch right.Kind() {
		case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
			rv := right.Int()
			return reflect.ValueOf(int64(lv) <= rv), nil
		case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
			rv := right.Uint()
			return reflect.ValueOf(lv <= rv), nil
		case reflect.Float32, reflect.Float64:
			rv := right.Float()
			return reflect.ValueOf(float64(lv) <= rv), nil
		default:
			return reflect.ValueOf(nil), fmt.Errorf("can not use data type of %s in LTE comparison", right.Kind().String())
		}
	case reflect.Float32, reflect.Float64:
		lv := left.Float()
		switch right.Kind() {
		case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
			rv := right.Int()
			return reflect.ValueOf(lv <= float64(rv)), nil
		case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
			rv := right.Uint()
			return reflect.ValueOf(lv <= float64(rv)), nil
		case reflect.Float32, reflect.Float64:
			rv := right.Float()
			return reflect.ValueOf(lv <= rv), nil
		default:
			return reflect.ValueOf(nil), fmt.Errorf("can not use data type of %s in LTE comparison", right.Kind().String())
		}
	default:
		if left.Type().String() == "time.Time" && right.Type().String() == "time.Time" {
			lv := left.Interface().(time.Time)
			rv := right.Interface().(time.Time)
			return reflect.ValueOf(lv.Before(rv) || lv == rv), nil
		}
		return reflect.ValueOf(nil), fmt.Errorf("can not use data type of %s in LTE comparison", left.Kind().String())
	}
}

// EvaluateEqual will evaluate Equal operation over two value
func EvaluateEqual(left, right reflect.Value) (reflect.Value, error) {
	switch left.Kind() {
	case reflect.String:
		lv := left.String()
		if right.Kind() == reflect.String {
			rv := right.String()
			return reflect.ValueOf(lv == rv), nil
		}
		return reflect.ValueOf(false), nil
	case reflect.Bool:
		lv := left.Bool()
		if right.Kind() == reflect.Bool {
			rv := right.Bool()
			return reflect.ValueOf(lv == rv), nil
		}
		return reflect.ValueOf(false), nil
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		lv := left.Int()
		switch right.Kind() {
		case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
			rv := right.Int()
			return reflect.ValueOf(lv == rv), nil
		case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
			rv := right.Uint()
			return reflect.ValueOf(lv == int64(rv)), nil
		case reflect.Float32, reflect.Float64:
			rv := right.Float()
			return reflect.ValueOf(float64(lv) == rv), nil
		default:
			return reflect.ValueOf(nil), fmt.Errorf("can not use data type of %s in EQ comparison", right.Kind().String())
		}
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		lv := left.Uint()
		switch right.Kind() {
		case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
			rv := right.Int()
			return reflect.ValueOf(int64(lv) == rv), nil
		case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
			rv := right.Uint()
			return reflect.ValueOf(lv == rv), nil
		case reflect.Float32, reflect.Float64:
			rv := right.Float()
			return reflect.ValueOf(float64(lv) == rv), nil
		default:
			return reflect.ValueOf(nil), fmt.Errorf("can not use data type of %s in EQ comparison", right.Kind().String())
		}
	case reflect.Float32, reflect.Float64:
		lv := left.Float()
		switch right.Kind() {
		case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
			rv := right.Int()
			return reflect.ValueOf(lv == float64(rv)), nil
		case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
			rv := right.Uint()
			return reflect.ValueOf(lv == float64(rv)), nil
		case reflect.Float32, reflect.Float64:
			rv := right.Float()
			return reflect.ValueOf(lv == rv), nil
		default:
			return reflect.ValueOf(nil), fmt.Errorf("can not use data type of %s in EQ comparison", right.Kind().String())
		}
	default:
		if left.Type().String() == "time.Time" && right.Type().String() == "time.Time" {
			lv := left.Interface().(time.Time)
			rv := right.Interface().(time.Time)
			return reflect.ValueOf(lv == rv), nil
		}
		return reflect.ValueOf(nil), fmt.Errorf("can not use data type of %s in EQ comparison", left.Kind().String())
	}
}

// EvaluateNotEqual will evaluate NotEqual operation over two value
func EvaluateNotEqual(left, right reflect.Value) (reflect.Value, error) {
	switch left.Kind() {
	case reflect.String:
		lv := left.String()
		if right.Kind() == reflect.String {
			rv := right.String()
			return reflect.ValueOf(lv != rv), nil
		}
		return reflect.ValueOf(false), nil
	case reflect.Bool:
		lv := left.Bool()
		if right.Kind() == reflect.Bool {
			rv := right.Bool()
			return reflect.ValueOf(lv != rv), nil
		}
		return reflect.ValueOf(false), nil
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		lv := left.Int()
		switch right.Kind() {
		case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
			rv := right.Int()
			return reflect.ValueOf(lv != rv), nil
		case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
			rv := right.Uint()
			return reflect.ValueOf(lv != int64(rv)), nil
		case reflect.Float32, reflect.Float64:
			rv := right.Float()
			return reflect.ValueOf(float64(lv) != rv), nil
		default:
			return reflect.ValueOf(nil), fmt.Errorf("can not use data type of %s in EQ comparison", right.Kind().String())
		}
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		lv := left.Uint()
		switch right.Kind() {
		case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
			rv := right.Int()
			return reflect.ValueOf(int64(lv) != rv), nil
		case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
			rv := right.Uint()
			return reflect.ValueOf(lv != rv), nil
		case reflect.Float32, reflect.Float64:
			rv := right.Float()
			return reflect.ValueOf(float64(lv) != rv), nil
		default:
			return reflect.ValueOf(nil), fmt.Errorf("can not use data type of %s in EQ comparison", right.Kind().String())
		}
	case reflect.Float32, reflect.Float64:
		lv := left.Float()
		switch right.Kind() {
		case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
			rv := right.Int()
			return reflect.ValueOf(lv != float64(rv)), nil
		case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
			rv := right.Uint()
			return reflect.ValueOf(lv != float64(rv)), nil
		case reflect.Float32, reflect.Float64:
			rv := right.Float()
			return reflect.ValueOf(lv != rv), nil
		default:
			return reflect.ValueOf(nil), fmt.Errorf("can not use data type of %s in EQ comparison", right.Kind().String())
		}
	default:
		if left.Type().String() == "time.Time" && right.Type().String() == "time.Time" {
			lv := left.Interface().(time.Time)
			rv := right.Interface().(time.Time)
			return reflect.ValueOf(lv != rv), nil
		}
		return reflect.ValueOf(nil), fmt.Errorf("can not use data type of %s in EQ comparison", left.Kind().String())
	}
}

// EvaluateLogicAnd will evaluate LogicalAnd operation over two value
func EvaluateLogicAnd(left, right reflect.Value) (reflect.Value, error) {
	if left.Kind() == reflect.Bool && right.Kind() == reflect.Bool {
		lv := left.Bool()
		rv := right.Bool()
		return reflect.ValueOf(lv && rv), nil
	}
	return reflect.ValueOf(nil), fmt.Errorf("can not use data type of %s in Logical AND comparison", left.Kind().String())
}

// EvaluateLogicOr will evaluate LogicalOr operation over two value
func EvaluateLogicOr(left, right reflect.Value) (reflect.Value, error) {
	if left.Kind() == reflect.Bool && right.Kind() == reflect.Bool {
		lv := left.Bool()
		rv := right.Bool()
		return reflect.ValueOf(lv || rv), nil
	}
	return reflect.ValueOf(nil), fmt.Errorf("can not use data type of %s in Logical OR comparison", left.Kind().String())
}

// EvaluateLogicSingle will evaluate single expression value
func EvaluateLogicSingle(left reflect.Value) (reflect.Value, error) {
	if left.Kind() == reflect.Bool {
		lv := left.Bool()
		return reflect.ValueOf(lv), nil
	}
	return reflect.ValueOf(nil), fmt.Errorf("can not use data type of %s in Logical AND comparison", left.Kind().String())
}
