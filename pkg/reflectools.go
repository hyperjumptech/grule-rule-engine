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
	"github.com/hyperjumptech/grule-rule-engine/logger"
	"math"
	"reflect"
)

// GetFunctionList get list of functions in a struct instance
func GetFunctionList(obj reflect.Value) ([]string, error) {
	if !IsStruct(obj) {
		return nil, fmt.Errorf("GetFunctionList : param is not a struct")
	}
	ret := make([]string, 0)

	objType := obj.Type()
	for i := 0; i < objType.NumMethod(); i++ {
		ret = append(ret, objType.Method(i).Name)
	}
	return ret, nil
}

// GetFunctionParameterTypes get list of parameter types of specific function in a struct instance
func GetFunctionParameterTypes(obj reflect.Value, methodName string) ([]reflect.Type, bool, error) {
	if !IsStruct(obj) {
		return nil, false, fmt.Errorf("GetFunctionParameterTypes : param is not a struct")
	}
	ret := make([]reflect.Type, 0)
	objType := obj.Type()

	meth, found := objType.MethodByName(methodName)
	if found {
		x := meth.Type
		for i := 1; i < x.NumIn(); i++ {
			ret = append(ret, x.In(i))
		}
		return ret, meth.Type.IsVariadic(), nil
	}
	return nil, false, fmt.Errorf("function %s not found", methodName)
}

// GetFunctionReturnTypes get list of return types of specific function in a struct instance
func GetFunctionReturnTypes(obj reflect.Value, methodName string) ([]reflect.Type, error) {
	if !IsStruct(obj) {
		return nil, fmt.Errorf("GetFunctionReturnTypes : param is not a struct")
	}
	ret := make([]reflect.Type, 0)
	objType := obj.Type()

	meth, found := objType.MethodByName(methodName)
	if found {
		x := meth.Type
		for i := 0; i < x.NumOut(); i++ {
			ret = append(ret, x.Out(i))
		}
	} else {
		return nil, fmt.Errorf("function %s not found", methodName)
	}
	return ret, nil
}

// InvokeFunction invokes a specific function in a struct instance, using parameters array
func InvokeFunction(obj reflect.Value, methodName string, param []reflect.Value) (retval []reflect.Value, err error) {
	defer func() {
		if r := recover(); r != nil {
			err = fmt.Errorf("error when invoking function %s. got %s", methodName, r)
		}
	}()

	if !IsStruct(obj) {
		return nil, fmt.Errorf("InvokeFunction : param is not a struct")
	}
	funcVal := obj.MethodByName(methodName)

	if !funcVal.IsValid() {
		return nil, fmt.Errorf("invalid function %s", methodName)
	}
	retVals := funcVal.Call(param)
	return retVals, nil
}

// IsValidField validates if an instance struct have a field with such name
func IsValidField(objVal reflect.Value, fieldName string) bool {
	if !IsStruct(objVal) {
		return false
	}
	objType := objVal.Type()
	if objType.Kind() == reflect.Struct {
		fieldVal := objVal.FieldByName(fieldName)
		return fieldVal.IsValid()
	} else if objType.Kind() == reflect.Ptr {
		fieldVal := objVal.Elem().FieldByName(fieldName)
		return fieldVal.IsValid()
	} else {
		return false
	}
}

// IsStruct validates if an instance is struct or pointer to struct
func IsStruct(val reflect.Value) bool {
	if val.IsValid() {
		typ := val.Type()
		if typ.Kind() != reflect.Ptr {
			return typ.Kind() == reflect.Struct
		}
		return typ.Elem().Kind() == reflect.Struct
	}
	return false
}

// ValueToInterface will try to obtain an interface to a speciffic value.
// it will detect the value's kind.
func ValueToInterface(v reflect.Value) interface{} {
	if v.Type().Kind() == reflect.String {
		return v.String()
	}
	switch v.Type().Kind() {
	case reflect.Int:
		return int(v.Int())
	case reflect.Int8:
		return int8(v.Int())
	case reflect.Int16:
		return int16(v.Int())
	case reflect.Int32:
		return int32(v.Int())
	case reflect.Int64:
		return v.Int()
	case reflect.Uint:
		return uint(v.Uint())
	case reflect.Uint8:
		return uint8(v.Uint())
	case reflect.Uint16:
		return uint16(v.Uint())
	case reflect.Uint32:
		return uint32(v.Uint())
	case reflect.Uint64:
		return v.Uint()
	case reflect.Float32:
		return float32(v.Float())
	case reflect.Float64:
		return v.Float()
	case reflect.Bool:
		return v.Bool()
	case reflect.Ptr:
		newPtr := reflect.New(v.Elem().Type())
		newPtr.Elem().Set(v.Elem())
		return newPtr.Interface()
	case reflect.Struct:
		if v.CanInterface() {
			return v.Interface()
		}
		logger.Log.Errorf("Can't interface value of struct %v", v)
		return nil
	default:
		return nil
	}
}

// GetAttributeList will populate list of struct's public member variable.
func GetAttributeList(obj reflect.Value) ([]string, error) {
	if !IsStruct(obj) {
		return nil, fmt.Errorf("GetAttributeList : param is not a struct")
	}
	strRet := make([]string, 0)
	e := obj.Elem()
	for i := 0; i < e.Type().NumField(); i++ {
		strRet = append(strRet, e.Type().Field(i).Name)
	}
	return strRet, nil
}

// GetAttributeValue will retrieve a members variable value.
func GetAttributeValue(obj reflect.Value, fieldName string) (reflect.Value, error) {
	if !IsStruct(obj) {
		return reflect.ValueOf(nil), fmt.Errorf("GetAttributeValue : param is not a struct")
	}
	if !IsValidField(obj, fieldName) {
		return reflect.ValueOf(nil), fmt.Errorf("attribute named %s not exist in struct", fieldName)
	}
	structval := obj
	var attrVal reflect.Value
	if structval.Kind() == reflect.Ptr {
		attrVal = structval.Elem().FieldByName(fieldName)
	} else {
		attrVal = structval.FieldByName(fieldName)
	}
	return attrVal, nil
}

// GetAttributeInterface will retrieve a members variable value as usable interface.
func GetAttributeInterface(obj reflect.Value, fieldName string) (interface{}, error) {
	val, err := GetAttributeValue(obj, fieldName)
	if err != nil {
		return nil, err
	}
	return ValueToInterface(val), nil
}

// GetAttributeType will return the type of a specific member variable
func GetAttributeType(obj reflect.Value, fieldName string) (reflect.Type, error) {
	if !IsStruct(obj) {
		return nil, fmt.Errorf("GetAttributeType : param is not a struct")
	}
	if !IsValidField(obj, fieldName) {
		return nil, fmt.Errorf("attribute named %s not exist in struct", fieldName)
	}
	structval := obj
	var attrVal reflect.Value
	if structval.Kind() == reflect.Ptr {
		attrVal = structval.Elem().FieldByName(fieldName)
	} else {
		attrVal = structval.FieldByName(fieldName)
	}
	return attrVal.Type(), nil
}

// SetAttributeValue will try to set a member variable value with a new one.
func SetAttributeValue(objVal reflect.Value, fieldName string, value reflect.Value) error {
	if !IsStruct(objVal) {
		return fmt.Errorf("SetAttributeValue : param is not a struct")
	}
	if !IsValidField(objVal, fieldName) {
		return fmt.Errorf("attribute named %s not exist in struct", fieldName)
	}
	var fieldVal reflect.Value
	objType := objVal.Type()
	// If Obj param is a pointer
	if objType.Kind() == reflect.Ptr {
		// And it points to a struct
		if objType.Elem().Kind() == reflect.Struct {
			fieldVal = objVal.Elem().FieldByName(fieldName)
		} else {
			// If its not point to struct ... return error
			return fmt.Errorf("object is pointing a non struct. %s", objType.Elem().Kind().String())
		}
	} else {
		// If Obj param is not a pointer.
		// And its a struct
		if objType.Kind() == reflect.Struct {
			fieldVal = objVal.FieldByName(fieldName)
		} else {
			// If its not a struct ... return error
			return fmt.Errorf("object is not a struct. %s", objType.Kind().String())
		}
	}

	// Check source data type compatibility with the field type
	if GetBaseKind(fieldVal) != GetBaseKind(value) { // pointer check
		if !(IsNumber(fieldVal) && IsNumber(value)) {
			return fmt.Errorf("can not assign type %s to %s", value.Type().String(), fieldVal.Type().String())
		}
	}
	if fieldVal.CanSet() {
		switch fieldVal.Type().Kind() {
		case reflect.String:
			fieldVal.SetString(value.String())
			break
		case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
			if GetBaseKind(value) == reflect.Uint64 {
				fieldVal.SetInt(int64(value.Uint()))
			} else if GetBaseKind(value) == reflect.Float64 {
				fieldVal.SetInt(int64(value.Float()))
			} else {
				fieldVal.SetInt(value.Int())
			}
			break
		case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
			if GetBaseKind(value) == reflect.Uint64 {
				fieldVal.SetUint(value.Uint())
			} else if GetBaseKind(value) == reflect.Float64 {
				fieldVal.SetUint(uint64(value.Float()))
			} else {
				fieldVal.SetUint(uint64(value.Int()))
			}
			break
		case reflect.Float32, reflect.Float64:
			if GetBaseKind(value) == reflect.Uint64 {
				fieldVal.SetFloat(float64(value.Uint()))
			} else if GetBaseKind(value) == reflect.Float64 {
				fieldVal.SetFloat(value.Float())
			} else {
				fieldVal.SetFloat(float64(value.Int()))
			}
			break
		case reflect.Bool:
			fieldVal.SetBool(value.Bool())
			break
		case reflect.Ptr:
			fieldVal.Set(value)
			break
		case reflect.Slice:
			// todo Add setter for slice type field
			return fmt.Errorf("unsupported operation to set slice")
		case reflect.Array:
			// todo Add setter for array type field
			return fmt.Errorf("unsupported operation to set array")
		case reflect.Map:
			// todo Add setter for map type field
			return fmt.Errorf("unsupported operation to set map")
		case reflect.Struct:
			if value.IsValid() {
				if ValueToInterface(value) == nil {
					return fmt.Errorf("Time set failed 1")
				}
				fieldVal.Set(value)
				//t := ValueToInterface(fieldVal).(time.Time)
				if ValueToInterface(fieldVal) == nil {
					return fmt.Errorf("Time set failed 2")
				}
			} else {
				return fmt.Errorf("Setting with nil is not allowed")
			}
			//// todo Add setter for slice type field
			//return fmt.Errorf("unsupported operation to set struct")
		default:
			return fmt.Errorf("unsupported operation to set %s", fieldVal.Type().String())
		}
	} else {
		return fmt.Errorf("can not set field")
	}
	return nil
}

// SetAttributeInterface will try to set a member variable value with a value from an interface
func SetAttributeInterface(obj reflect.Value, fieldName string, value interface{}) error {
	if !IsStruct(obj) {
		return fmt.Errorf("SetAttributeInterface : param is not a struct")
	}
	if !IsValidField(obj, fieldName) {
		return fmt.Errorf("attribute named %s not exist in struct", fieldName)
	}

	return SetAttributeValue(obj, fieldName, reflect.ValueOf(value))
}

// IsAttributeArray validate if a member variable is an array or a slice.
func IsAttributeArray(objVal reflect.Value, fieldName string) (bool, error) {
	if !IsStruct(objVal) {
		return false, fmt.Errorf("IsAttributeArray : param is not a struct")
	}
	if !IsValidField(objVal, fieldName) {
		return false, fmt.Errorf("attribute named %s not exist in struct", fieldName)
	}
	fieldVal := objVal.Elem().FieldByName(fieldName)
	return fieldVal.Type().Kind() == reflect.Array || fieldVal.Type().Kind() == reflect.Slice, nil
}

// SetMapArrayValue will set a value into map array indicated by a selector
func SetMapArrayValue(mapArray, selector reflect.Value, newValue reflect.Value) (err error) {
	objVal := mapArray
	if objVal.Type().Kind() == reflect.Map {
		objVal.SetMapIndex(selector, newValue)
		return nil
	}
	if objVal.Type().Kind() == reflect.Array || objVal.Type().Kind() == reflect.Slice {
		defer func() {
			if errPanic := recover(); errPanic != nil {
				err = fmt.Errorf("index %d is out of bound", selector.Int())
			}
		}()

		idx := 0
		switch GetBaseKind(selector) {
		case reflect.Int64:
			idx = int(selector.Int())
		case reflect.Uint64:
			idx = int(selector.Uint())
		case reflect.Float32:
			idx = int(selector.Float())
		default:
			return fmt.Errorf("array selector can only be numeric type")
		}

		retVal := objVal.Index(idx)
		retVal.Set(newValue)
		return nil
	}
	return fmt.Errorf("argument is not an array, slice nor map")
}

// GetMapArrayValue get value of map, array atau slice by its selector value
func GetMapArrayValue(mapArray, selector interface{}) (ret interface{}, err error) {
	if mapArray == nil {
		return nil, fmt.Errorf("nil map, array or slice")
	}
	objVal := reflect.ValueOf(mapArray)
	if objVal.Type().Kind() == reflect.Map {
		if objVal.Type().Key() == reflect.TypeOf(selector) {
			defer func() {
				if errPanic := recover(); errPanic != nil {
					ret = nil
					err = fmt.Errorf("map key not exist")
				}
			}()
			retVal := objVal.MapIndex(reflect.ValueOf(selector))
			if retVal.IsZero() {
				return nil, fmt.Errorf("selector not exist in map key")
			}
			return ValueToInterface(retVal), nil
		}
		return nil, fmt.Errorf("map requires key of type %s, found %s", objVal.Type().Key().String(), reflect.TypeOf(selector).String())
	}
	if objVal.Type().Kind() == reflect.Array || objVal.Type().Kind() == reflect.Slice {
		defer func() {
			if errPanic := recover(); errPanic != nil {
				ret = nil
				err = fmt.Errorf("index %d is out of bound", selector.(int))
			}
		}()

		idx := 0
		switch GetBaseKind(reflect.ValueOf(selector)) {
		case reflect.Int64:
			idx = int(reflect.ValueOf(selector).Int())
		case reflect.Uint64:
			idx = int(reflect.ValueOf(selector).Uint())
		case reflect.Float32:
			idx = int(reflect.ValueOf(selector).Float())
		default:
			return nil, fmt.Errorf("array selector can only be numeric type")
		}

		retVal := objVal.Index(idx)
		return ValueToInterface(retVal), nil
	}
	return nil, fmt.Errorf("argument is not an array, slice nor map")
}

// IsAttributeMap validate if a member variable is a map.
func IsAttributeMap(obj reflect.Value, fieldName string) (bool, error) {
	if !IsStruct(obj) {
		return false, fmt.Errorf("IsAttributeMap : param is not a struct")
	}
	if !IsValidField(obj, fieldName) {
		return false, fmt.Errorf("attribute named %s not exist in struct", fieldName)
	}
	var fieldVal reflect.Value
	if obj.Kind() == reflect.Ptr {
		fieldVal = obj.Elem().FieldByName(fieldName)
	} else if obj.Kind() == reflect.Struct {
		fieldVal = obj.FieldByName(fieldName)
	}
	return fieldVal.Type().Kind() == reflect.Map, nil
}

// IsAttributeNilOrZero validate if a member variable is nil or zero.
func IsAttributeNilOrZero(obj reflect.Value, fieldName string) (bool, error) {
	if !IsStruct(obj) {
		return false, fmt.Errorf("IsAttributeNilOrZero : param is not a struct")
	}
	if !IsValidField(obj, fieldName) {
		return false, fmt.Errorf("attribute named %s not exist in struct", fieldName)
	}
	var fieldVal reflect.Value
	if obj.Kind() == reflect.Ptr {
		fieldVal = obj.Elem().FieldByName(fieldName)
	} else if obj.Kind() == reflect.Struct {
		fieldVal = obj.FieldByName(fieldName)
	}
	if fieldVal.Kind() == reflect.Ptr {
		return fieldVal.IsNil(), nil
	}
	if fieldVal.Kind() == reflect.Struct {
		z0 := reflect.Zero(fieldVal.Type())
		return ValueToInterface(z0) == ValueToInterface(fieldVal), nil
	}
	if GetBaseKind(fieldVal) == reflect.Int64 {
		return fieldVal.Int() == 0, nil
	}
	if GetBaseKind(fieldVal) == reflect.Uint64 {
		return fieldVal.Uint() == 0, nil
	}
	if GetBaseKind(fieldVal) == reflect.Float64 {
		return fieldVal.Float() == 0, nil
	}
	if GetBaseKind(fieldVal) == reflect.String {
		return len(fieldVal.String()) == 0, nil
	}
	if GetBaseKind(fieldVal) == reflect.Bool {
		return fieldVal.Bool() == false, nil
	}
	if fieldVal.Type().Kind() == reflect.Map || fieldVal.Type().Kind() == reflect.Array || fieldVal.Type().Kind() == reflect.Slice {
		return fieldVal.IsNil() || reflectIsZero(fieldVal), nil
	}
	return false, nil
}

func reflectIsZero(v reflect.Value) bool {
	switch v.Kind() {
	case reflect.Bool:
		return !v.Bool()
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return v.Int() == 0
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr:
		return v.Uint() == 0
	case reflect.Float32, reflect.Float64:
		return math.Float64bits(v.Float()) == 0
	case reflect.Complex64, reflect.Complex128:
		c := v.Complex()
		return math.Float64bits(real(c)) == 0 && math.Float64bits(imag(c)) == 0
	case reflect.Array:
		for i := 0; i < v.Len(); i++ {
			if !reflectIsZero(v.Index(i)) {
				return false
			}
		}
		return true
	case reflect.Chan, reflect.Func, reflect.Interface, reflect.Map, reflect.Ptr, reflect.Slice, reflect.UnsafePointer:
		return v.IsNil()
	case reflect.String:
		return v.Len() == 0
	case reflect.Struct:
		for i := 0; i < v.NumField(); i++ {
			if !reflectIsZero(v.Field(i)) {
				return false
			}
		}
		return true
	default:
		panic(&reflect.ValueError{Method: "reflect.Value.IsZero", Kind: v.Kind()})
	}
}

// GetBaseKind will try to obtain base obtainable kind of a value, so we know what method to call val.Int(), val.Uint(), etc.
func GetBaseKind(val reflect.Value) reflect.Kind {
	switch val.Kind() {
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return reflect.Int64
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		return reflect.Uint64
	case reflect.Float32, reflect.Float64:
		return reflect.Float64
	default:
		return val.Kind()
	}
}

// IsNumber will check a value if its a number eg, int,uint or float
func IsNumber(val reflect.Value) bool {
	switch GetBaseKind(val) {
	case reflect.Int64, reflect.Uint64, reflect.Float64:
		return true
	}
	return false
}
