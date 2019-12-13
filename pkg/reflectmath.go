package pkg

import (
	"fmt"
	"github.com/juju/errors"
	"reflect"
)

// ValueAdd will try to do a mathematical addition between two values.
// It will return another value as the result and an error if between the two values are not compatible for Addition
func ValueAdd(a, b reflect.Value) (reflect.Value, error) {
	aBkind := GetBaseKind(a)
	bBkind := GetBaseKind(b)

	switch aBkind {
	case reflect.Int64:
		switch bBkind {
		case reflect.Int64:
			return reflect.ValueOf(a.Int() + b.Int()), nil
		case reflect.Uint64:
			return reflect.ValueOf(a.Int() + int64(b.Uint())), nil
		case reflect.Float64:
			return reflect.ValueOf(float64(a.Int()) + b.Float()), nil
		case reflect.String:
			return reflect.ValueOf(fmt.Sprintf("%d%s", a.Int(), b.String())), nil
		default:
			return reflect.ValueOf(nil), errors.Errorf("Can not do addition math operator between %s and %s", a.Kind().String(), b.Kind().String())
		}
	case reflect.Uint64:
		switch bBkind {
		case reflect.Int64:
			return reflect.ValueOf(int64(a.Uint()) + b.Int()), nil
		case reflect.Uint64:
			return reflect.ValueOf(a.Uint() + b.Uint()), nil
		case reflect.Float64:
			return reflect.ValueOf(float64(a.Uint()) + b.Float()), nil
		case reflect.String:
			return reflect.ValueOf(fmt.Sprintf("%d%s", a.Uint(), b.String())), nil
		default:
			return reflect.ValueOf(nil), errors.Errorf("Can not do addition math operator between %s and %s", a.Kind().String(), b.Kind().String())
		}
	case reflect.Float64:
		switch bBkind {
		case reflect.Int64:
			return reflect.ValueOf(a.Float() + float64(b.Int())), nil
		case reflect.Uint64:
			return reflect.ValueOf(a.Float() + float64(b.Uint())), nil
		case reflect.Float64:
			return reflect.ValueOf(a.Float() + b.Float()), nil
		case reflect.String:
			return reflect.ValueOf(fmt.Sprintf("%f%s", a.Float(), b.String())), nil
		default:
			return reflect.ValueOf(nil), errors.Errorf("Can not do addition math operator between %s and %s", a.Kind().String(), b.Kind().String())
		}
	case reflect.String:
		switch bBkind {
		case reflect.Int64:
			return reflect.ValueOf(fmt.Sprintf("%s%d", a.String(), b.Int())), nil
		case reflect.Uint64:
			return reflect.ValueOf(fmt.Sprintf("%s%d", a.String(), b.Uint())), nil
		case reflect.Float64:
			return reflect.ValueOf(fmt.Sprintf("%s%f", a.String(), b.Float())), nil
		case reflect.String:
			return reflect.ValueOf(fmt.Sprintf("%s%s", a.String(), b.String())), nil
		case reflect.Bool:
			return reflect.ValueOf(fmt.Sprintf("%s%v", a.String(), b.Bool())), nil
		case reflect.Ptr:
			if b.CanInterface() {
				return reflect.ValueOf(fmt.Sprintf("%s%v", a.String(), b.Interface())), nil
			}
			return reflect.ValueOf(nil), errors.Errorf("Can not do addition math operator between %s and non interface-able %s", a.Kind().String(), b.Kind().String())
		default:
			return reflect.ValueOf(fmt.Sprintf("%s%v", a.String(), b.String())), nil
		}
	default:
		if bBkind == reflect.String {
			if a.CanInterface() {
				return reflect.ValueOf(fmt.Sprintf("%v%s", a.Interface(), b.String())), nil
			}
			return reflect.ValueOf(nil), errors.Errorf("Can not do addition math operator between non interface-able %s and %s", b.Kind().String(), a.Kind().String())
		}
		return reflect.ValueOf(nil), errors.Errorf("Can not do math operator between %s and %s", a.Kind().String(), b.Kind().String())
	}
}

// ValueSub will try to do a mathematical substraction between two values.
// It will return another value as the result and an error if between the two values are not compatible for substraction
func ValueSub(a, b reflect.Value) (reflect.Value, error) {
	aBkind := GetBaseKind(a)
	bBkind := GetBaseKind(b)

	switch aBkind {
	case reflect.Int64:
		switch bBkind {
		case reflect.Int64:
			return reflect.ValueOf(a.Int() - b.Int()), nil
		case reflect.Uint64:
			return reflect.ValueOf(a.Int() - int64(b.Uint())), nil
		case reflect.Float64:
			return reflect.ValueOf(float64(a.Int()) - b.Float()), nil
		default:
			return reflect.ValueOf(nil), errors.Errorf("Can not do subtraction math operator between %s and %s", a.Kind().String(), b.Kind().String())
		}
	case reflect.Uint64:
		switch bBkind {
		case reflect.Int64:
			return reflect.ValueOf(int64(a.Uint()) - b.Int()), nil
		case reflect.Uint64:
			return reflect.ValueOf(a.Uint() - b.Uint()), nil
		case reflect.Float64:
			return reflect.ValueOf(float64(a.Uint()) - b.Float()), nil
		default:
			return reflect.ValueOf(nil), errors.Errorf("Can not do subtraction math operator between %s and %s", a.Kind().String(), b.Kind().String())
		}
	case reflect.Float64:
		switch bBkind {
		case reflect.Int64:
			return reflect.ValueOf(a.Float() - float64(b.Int())), nil
		case reflect.Uint64:
			return reflect.ValueOf(a.Float() - float64(b.Uint())), nil
		case reflect.Float64:
			return reflect.ValueOf(a.Float() - b.Float()), nil
		default:
			return reflect.ValueOf(nil), errors.Errorf("Can not do subtraction math operator between %s and %s", a.Kind().String(), b.Kind().String())
		}
	default:
		return reflect.ValueOf(nil), errors.Errorf("Can not do subtraction math operator between %s and %s", a.Kind().String(), b.Kind().String())
	}
}

// ValueMul will try to do a mathematical multiplication between two values.
// It will return another value as the result and an error if between the two values are not compatible for multiplication
func ValueMul(a, b reflect.Value) (reflect.Value, error) {
	aBkind := GetBaseKind(a)
	bBkind := GetBaseKind(b)

	switch aBkind {
	case reflect.Int64:
		switch bBkind {
		case reflect.Int64:
			return reflect.ValueOf(a.Int() * b.Int()), nil
		case reflect.Uint64:
			return reflect.ValueOf(a.Int() * int64(b.Uint())), nil
		case reflect.Float64:
			return reflect.ValueOf(float64(a.Int()) * b.Float()), nil
		default:
			return reflect.ValueOf(nil), errors.Errorf("Can not do multiplication math operator between %s and %s", a.Kind().String(), b.Kind().String())
		}
	case reflect.Uint64:
		switch bBkind {
		case reflect.Int64:
			return reflect.ValueOf(int64(a.Uint()) * b.Int()), nil
		case reflect.Uint64:
			return reflect.ValueOf(a.Uint() * b.Uint()), nil
		case reflect.Float64:
			return reflect.ValueOf(float64(a.Uint()) * b.Float()), nil
		default:
			return reflect.ValueOf(nil), errors.Errorf("Can not do multiplication math operator between %s and %s", a.Kind().String(), b.Kind().String())
		}
	case reflect.Float64:
		switch bBkind {
		case reflect.Int64:
			return reflect.ValueOf(a.Float() * float64(b.Int())), nil
		case reflect.Uint64:
			return reflect.ValueOf(a.Float() * float64(b.Uint())), nil
		case reflect.Float64:
			return reflect.ValueOf(a.Float() * b.Float()), nil
		default:
			return reflect.ValueOf(nil), errors.Errorf("Can not do multiplication math operator between %s and %s", a.Kind().String(), b.Kind().String())
		}
	default:
		return reflect.ValueOf(nil), errors.Errorf("Can not do multiplication math operator between %s and %s", a.Kind().String(), b.Kind().String())
	}
}

// ValueDiv will try to do a mathematical division between two values.
// It will return another value as the result and an error if between the two values are not compatible for division
func ValueDiv(a, b reflect.Value) (reflect.Value, error) {
	aBkind := GetBaseKind(a)
	bBkind := GetBaseKind(b)

	switch aBkind {
	case reflect.Int64:
		switch bBkind {
		case reflect.Int64:
			return reflect.ValueOf(a.Int() / b.Int()), nil
		case reflect.Uint64:
			return reflect.ValueOf(a.Int() / int64(b.Uint())), nil
		case reflect.Float64:
			return reflect.ValueOf(float64(a.Int()) / b.Float()), nil
		default:
			return reflect.ValueOf(nil), errors.Errorf("Can not do division math operator between %s and %s", a.Kind().String(), b.Kind().String())
		}
	case reflect.Uint64:
		switch bBkind {
		case reflect.Int64:
			return reflect.ValueOf(int64(a.Uint()) / b.Int()), nil
		case reflect.Uint64:
			return reflect.ValueOf(a.Uint() / b.Uint()), nil
		case reflect.Float64:
			return reflect.ValueOf(float64(a.Uint()) / b.Float()), nil
		default:
			return reflect.ValueOf(nil), errors.Errorf("Can not do division math operator between %s and %s", a.Kind().String(), b.Kind().String())
		}
	case reflect.Float64:
		switch bBkind {
		case reflect.Int64:
			return reflect.ValueOf(a.Float() / float64(b.Int())), nil
		case reflect.Uint64:
			return reflect.ValueOf(a.Float() / float64(b.Uint())), nil
		case reflect.Float64:
			return reflect.ValueOf(a.Float() / b.Float()), nil
		default:
			return reflect.ValueOf(nil), errors.Errorf("Can not do division math operator between %s and %s", a.Kind().String(), b.Kind().String())
		}
	default:
		return reflect.ValueOf(nil), errors.Errorf("Can not do division math operator between %s and %s", a.Kind().String(), b.Kind().String())
	}
}
