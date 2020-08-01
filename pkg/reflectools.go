package pkg

import (
	"fmt"
	"math"
	"reflect"
	"time"

	"github.com/juju/errors"
	"github.com/sirupsen/logrus"
)

// GetFunctionList get list of functions in a struct instance
func GetFunctionList(obj interface{}) ([]string, error) {
	if !IsStruct(obj) {
		return nil, errors.Errorf("param is not a struct")
	}
	ret := make([]string, 0)

	objType := reflect.TypeOf(obj)
	for i := 0; i < objType.NumMethod(); i++ {
		ret = append(ret, objType.Method(i).Name)
	}
	return ret, nil
}

// GetFunctionParameterTypes get list of parameter types of specific function in a struct instance
func GetFunctionParameterTypes(obj interface{}, methodName string) ([]reflect.Type, bool, error) {
	if !IsStruct(obj) {
		return nil, false, errors.Errorf("param is not a struct")
	}
	ret := make([]reflect.Type, 0)
	objType := reflect.TypeOf(obj)

	var meth reflect.Method
	var found bool

	if objType.String() == "reflect.Value" {
		val := obj.(reflect.Value)
		meth, found = val.Type().MethodByName(methodName)
	} else {
		meth, found = objType.MethodByName(methodName)
	}
	if found {
		x := meth.Type
		for i := 1; i < x.NumIn(); i++ {
			ret = append(ret, x.In(i))
		}
		return ret, meth.Type.IsVariadic(), nil
	}
	return nil, false, errors.Errorf("function %s not found", methodName)
}

// GetFunctionReturnTypes get list of return types of specific function in a struct instance
func GetFunctionReturnTypes(obj interface{}, methodName string) ([]reflect.Type, error) {
	if !IsStruct(obj) {
		return nil, errors.Errorf("param is not a struct")
	}
	ret := make([]reflect.Type, 0)
	objType := reflect.TypeOf(obj)

	meth, found := objType.MethodByName(methodName)
	if found {
		x := meth.Type
		for i := 0; i < x.NumOut(); i++ {
			ret = append(ret, x.Out(i))
		}
	} else {
		return nil, errors.Errorf("function %s not found", methodName)
	}
	return ret, nil
}

// InvokeFunction invokes a specific function in a struct instance, using parameters array
func InvokeFunction(obj interface{}, methodName string, param []interface{}) ([]interface{}, error) {
	if !IsStruct(obj) {
		return nil, errors.Errorf("param is not a struct")
	}
	var objVal reflect.Value
	if reflect.TypeOf(obj).Name() == "Value" {
		objVal = obj.(reflect.Value)
	} else {
		objVal = reflect.ValueOf(obj)
	}
	funcVal := objVal.MethodByName(methodName)

	if !funcVal.IsValid() {
		return nil, errors.New(fmt.Sprintf("invalid function %s", methodName))
	}

	argVals := make([]reflect.Value, len(param))
	for idx, val := range param {
		argVals[idx] = reflect.ValueOf(val)
	}

	retVals := funcVal.Call(argVals)
	ret := make([]interface{}, len(retVals))
	for idx, r := range retVals {
		ret[idx] = ValueToInterface(r)
	}
	return ret, nil
}

// IsValidField validates if an instance struct have a field with such name
func IsValidField(obj interface{}, fieldName string) bool {
	if !IsStruct(obj) {
		return false
	}
	objType := reflect.TypeOf(obj)
	objVal := reflect.ValueOf(obj)
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
func IsStruct(obj interface{}) bool {
	if !reflect.ValueOf(obj).IsValid() {
		return false
	}
	objType := reflect.TypeOf(obj)
	if objType.Kind() != reflect.Ptr {
		return objType.Kind() == reflect.Struct
	}
	return objType.Elem().Kind() == reflect.Struct
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
		logrus.Errorf("Can't interface value of struct %v", v)
		return nil
	default:
		return nil
	}
}

// GetAttributeList will populate list of struct's public member variable.
func GetAttributeList(obj interface{}) ([]string, error) {
	if !IsStruct(obj) {
		return nil, errors.Errorf("param is not a struct")
	}
	strRet := make([]string, 0)
	v := reflect.ValueOf(obj)
	e := v.Elem()
	for i := 0; i < e.Type().NumField(); i++ {
		strRet = append(strRet, e.Type().Field(i).Name)
	}
	return strRet, nil
}

// GetAttributeValue will retrieve a members variable value.
func GetAttributeValue(obj interface{}, fieldName string) (reflect.Value, error) {
	if !IsStruct(obj) {
		return reflect.ValueOf(nil), errors.Errorf("param is not a struct")
	}
	if !IsValidField(obj, fieldName) {
		return reflect.ValueOf(nil), errors.Errorf("attribute named %s not exist in struct", fieldName)
	}
	structval := reflect.ValueOf(obj)
	var attrVal reflect.Value
	if structval.Kind() == reflect.Ptr {
		attrVal = structval.Elem().FieldByName(fieldName)
	} else {
		attrVal = structval.FieldByName(fieldName)
	}
	return attrVal, nil
}

// GetAttributeInterface will retrieve a members variable value as usable interface.
func GetAttributeInterface(obj interface{}, fieldName string) (interface{}, error) {
	val, err := GetAttributeValue(obj, fieldName)
	if err != nil {
		return nil, err
	}
	return ValueToInterface(val), nil
}

// GetAttributeType will return the type of a specific member variable
func GetAttributeType(obj interface{}, fieldName string) (reflect.Type, error) {
	if !IsStruct(obj) {
		return nil, errors.Errorf("param is not a struct")
	}
	if !IsValidField(obj, fieldName) {
		return nil, errors.Errorf("attribute named %s not exist in struct", fieldName)
	}
	structval := reflect.ValueOf(obj)
	var attrVal reflect.Value
	if structval.Kind() == reflect.Ptr {
		attrVal = structval.Elem().FieldByName(fieldName)
	} else {
		attrVal = structval.FieldByName(fieldName)
	}
	return attrVal.Type(), nil
}

// SetAttributeValue will try to set a member variable value with a new one.
func SetAttributeValue(obj interface{}, fieldName string, value reflect.Value) error {
	if !IsStruct(obj) {
		return errors.Errorf("param is not a struct")
	}
	if !IsValidField(obj, fieldName) {
		return errors.Errorf("attribute named %s not exist in struct", fieldName)
	}
	var fieldVal reflect.Value
	objType := reflect.TypeOf(obj)
	objVal := reflect.ValueOf(obj)
	// If Obj param is a pointer
	if objType.Kind() == reflect.Ptr {
		// And it points to a struct
		if objType.Elem().Kind() == reflect.Struct {
			fieldVal = objVal.Elem().FieldByName(fieldName)
		} else {
			// If its not point to struct ... return error
			return errors.Errorf("object is pointing a non struct. %s", objType.Elem().Kind().String())
		}
	} else {
		// If Obj param is not a pointer.
		// And its a struct
		if objType.Kind() == reflect.Struct {
			fieldVal = objVal.FieldByName(fieldName)
		} else {
			// If its not a struct ... return error
			return errors.Errorf("object is not a struct. %s", objType.Kind().String())
		}
	}

	// Check source data type compatibility with the field type
	if GetBaseKind(fieldVal) != GetBaseKind(value) { // pointer check
		return errors.Errorf("can not assign type %s to %s", value.Type().String(), fieldVal.Type().String())
	}
	if fieldVal.CanSet() {
		switch fieldVal.Type().Kind() {
		case reflect.String:
			fieldVal.SetString(value.String())
			break
		case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
			fieldVal.SetInt(value.Int())
			break
		case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
			fieldVal.SetUint(value.Uint())
			break
		case reflect.Float32, reflect.Float64:
			fieldVal.SetFloat(value.Float())
			break
		case reflect.Bool:
			fieldVal.SetBool(value.Bool())
			break
		case reflect.Ptr:
			fieldVal.Set(value)
			break
		case reflect.Slice:
			// todo Add setter for slice type field
			return errors.Errorf("unsupported operation to set slice")
		case reflect.Array:
			// todo Add setter for array type field
			return errors.Errorf("unsupported operation to set array")
		case reflect.Map:
			// todo Add setter for map type field
			return errors.Errorf("unsupported operation to set map")
		case reflect.Struct:
			if value.IsValid() {
				if ValueToInterface(value) == nil {
					return errors.Errorf("Time set failed 1")
				}
				fieldVal.Set(value)
				//t := ValueToInterface(fieldVal).(time.Time)
				if ValueToInterface(fieldVal) == nil {
					return errors.Errorf("Time set failed 2")
				}
			} else {
				return errors.Errorf("Setting with nil is not allowed")
			}
			//// todo Add setter for slice type field
			//return errors.Errorf("unsupported operation to set struct")
		default:
			return nil
		}
	} else {
		return errors.Errorf("can not set field")
	}
	return nil
}

// SetAttributeInterface will try to set a member variable value with a value from an interface
func SetAttributeInterface(obj interface{}, fieldName string, value interface{}) error {
	if !IsStruct(obj) {
		return errors.Errorf("param is not a struct")
	}
	if !IsValidField(obj, fieldName) {
		return errors.Errorf("attribute named %s not exist in struct", fieldName)
	}

	return SetAttributeValue(obj, fieldName, reflect.ValueOf(value))
}

// IsAttributeArray validate if a member variable is an array or a slice.
func IsAttributeArray(obj interface{}, fieldName string) (bool, error) {
	if !IsStruct(obj) {
		return false, errors.Errorf("param is not a struct")
	}
	if !IsValidField(obj, fieldName) {
		return false, errors.Errorf("attribute named %s not exist in struct", fieldName)
	}
	objVal := reflect.ValueOf(obj)
	fieldVal := objVal.Elem().FieldByName(fieldName)
	return fieldVal.Type().Kind() == reflect.Array || fieldVal.Type().Kind() == reflect.Slice, nil
}

// IsAttributeMap validate if a member variable is a map.
func IsAttributeMap(obj interface{}, fieldName string) (bool, error) {
	if !IsStruct(obj) {
		return false, errors.Errorf("param is not a struct")
	}
	if !IsValidField(obj, fieldName) {
		return false, errors.Errorf("attribute named %s not exist in struct", fieldName)
	}
	objVal := reflect.ValueOf(obj)
	fieldVal := objVal.Elem().FieldByName(fieldName)
	return fieldVal.Type().Kind() == reflect.Map, nil
}

// IsAttributeNilOrZero validate if a member variable is nil or zero.
func IsAttributeNilOrZero(obj interface{}, fieldName string) (bool, error) {
	if !IsStruct(obj) {
		return false, errors.Errorf("param is not a struct")
	}
	if !IsValidField(obj, fieldName) {
		return false, errors.Errorf("attribute named %s not exist in struct", fieldName)
	}
	objVal := reflect.ValueOf(obj)
	fieldVal := objVal.Elem().FieldByName(fieldName)
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

// GetAttributeStringValue will try to obtain member variable's string value
func GetAttributeStringValue(obj interface{}, fieldName string) (string, error) {
	val, err := GetAttributeInterface(obj, fieldName)
	if err != nil {
		return "", err
	}
	return val.(string), err
}

// SetAttributeStringValue will try to set member variable's string value
func SetAttributeStringValue(obj interface{}, fieldName string, newValue string) error {
	return SetAttributeInterface(obj, fieldName, newValue)
}

// GetAttributeIntValue will try to obtain member variable's int value
func GetAttributeIntValue(obj interface{}, fieldName string) (int, error) {
	val, err := GetAttributeInterface(obj, fieldName)
	if err != nil {
		return 0, err
	}
	return val.(int), err

}

// SetAttributeIntValue will try to set member variable's int value
func SetAttributeIntValue(obj interface{}, fieldName string, newValue int) error {
	return SetAttributeInterface(obj, fieldName, newValue)
}

// GetAttributeInt8Value will try to obtain member variable's int8 value
func GetAttributeInt8Value(obj interface{}, fieldName string) (int8, error) {
	val, err := GetAttributeInterface(obj, fieldName)
	if err != nil {
		return 0, err
	}
	return val.(int8), err
}

// SetAttributeInt8Value will try to set member variable's int8 value
func SetAttributeInt8Value(obj interface{}, fieldName string, newValue int8) error {
	return SetAttributeInterface(obj, fieldName, newValue)
}

// GetAttributeInt16Value will try to obtain member variable's int16 value
func GetAttributeInt16Value(obj interface{}, fieldName string) (int16, error) {
	val, err := GetAttributeInterface(obj, fieldName)
	if err != nil {
		return 0, err
	}
	return val.(int16), err
}

// SetAttributeInt16Value will try to set member variable's int16 value
func SetAttributeInt16Value(obj interface{}, fieldName string, newValue int16) error {
	return SetAttributeInterface(obj, fieldName, newValue)
}

// GetAttributeInt32Value will try to obtain member variable's int32 value
func GetAttributeInt32Value(obj interface{}, fieldName string) (int32, error) {
	val, err := GetAttributeInterface(obj, fieldName)
	if err != nil {
		return 0, err
	}
	return val.(int32), err
}

// SetAttributeInt32Value will try to set member variable's int32 value
func SetAttributeInt32Value(obj interface{}, fieldName string, newValue int32) error {
	return SetAttributeInterface(obj, fieldName, newValue)
}

// GetAttributeInt64Value will try to obtain member variable's int64 value
func GetAttributeInt64Value(obj interface{}, fieldName string) (int64, error) {
	val, err := GetAttributeInterface(obj, fieldName)
	if err != nil {
		return 0, err
	}
	return val.(int64), err
}

// SetAttributeInt64Value will try to set member variable's int64 value
func SetAttributeInt64Value(obj interface{}, fieldName string, newValue int64) error {
	return SetAttributeInterface(obj, fieldName, newValue)
}

// GetAttributeUIntValue will try to obtain member variable's uint value
func GetAttributeUIntValue(obj interface{}, fieldName string) (uint, error) {
	val, err := GetAttributeInterface(obj, fieldName)
	if err != nil {
		return 0, err
	}
	return val.(uint), err
}

// SetAttributeUIntValue will try to set member variable's uint value
func SetAttributeUIntValue(obj interface{}, fieldName string, newValue uint) error {
	return SetAttributeInterface(obj, fieldName, newValue)
}

// GetAttributeUInt8Value will try to obtain member variable's uint8 value
func GetAttributeUInt8Value(obj interface{}, fieldName string) (uint8, error) {
	val, err := GetAttributeInterface(obj, fieldName)
	if err != nil {
		return 0, err
	}
	return val.(uint8), err
}

// SetAttributeUInt8Value will try to set member variable's uint8 value
func SetAttributeUInt8Value(obj interface{}, fieldName string, newValue uint8) error {
	return SetAttributeInterface(obj, fieldName, newValue)
}

// GetAttributeUInt16Value will try to obtain member variable's uint16 value
func GetAttributeUInt16Value(obj interface{}, fieldName string) (uint16, error) {
	val, err := GetAttributeInterface(obj, fieldName)
	if err != nil {
		return 0, err
	}
	return val.(uint16), err
}

// SetAttributeUInt16Value will try to set member variable's uint16 value
func SetAttributeUInt16Value(obj interface{}, fieldName string, newValue uint16) error {
	return SetAttributeInterface(obj, fieldName, newValue)
}

// GetAttributeUInt32Value will try to obtain member variable's uint32 value
func GetAttributeUInt32Value(obj interface{}, fieldName string) (uint32, error) {
	val, err := GetAttributeInterface(obj, fieldName)
	if err != nil {
		return 0, err
	}
	return val.(uint32), err
}

// SetAttributeUInt32Value will try to set member variable's uint32 value
func SetAttributeUInt32Value(obj interface{}, fieldName string, newValue uint32) error {
	return SetAttributeInterface(obj, fieldName, newValue)
}

// GetAttributeUInt64Value will try to obtain member variable's uint64 value
func GetAttributeUInt64Value(obj interface{}, fieldName string) (uint64, error) {
	val, err := GetAttributeInterface(obj, fieldName)
	if err != nil {
		return 0, err
	}
	return val.(uint64), err
}

// SetAttributeUInt64Value will try to set member variable's uint64 value
func SetAttributeUInt64Value(obj interface{}, fieldName string, newValue uint64) error {
	return SetAttributeInterface(obj, fieldName, newValue)
}

// GetAttributeBoolValue will try to obtain member variable's bool value
func GetAttributeBoolValue(obj interface{}, fieldName string) (bool, error) {
	val, err := GetAttributeInterface(obj, fieldName)
	if err != nil {
		return false, err
	}
	return val.(bool), err
}

// SetAttributeBoolValue will try to set member variable's bool value
func SetAttributeBoolValue(obj interface{}, fieldName string, newValue bool) error {
	return SetAttributeInterface(obj, fieldName, newValue)
}

// GetAttributeFloat32Value will try to obtain member variable's float32 value
func GetAttributeFloat32Value(obj interface{}, fieldName string) (float32, error) {
	val, err := GetAttributeInterface(obj, fieldName)
	if err != nil {
		return 0, err
	}
	return val.(float32), err
}

// SetAttributeFloat32Value will try to set member variable's float32 value
func SetAttributeFloat32Value(obj interface{}, fieldName string, newValue float32) error {
	return SetAttributeInterface(obj, fieldName, newValue)
}

// GetAttributeFloat64Value will try to obtain member variable's float64 value
func GetAttributeFloat64Value(obj interface{}, fieldName string) (float64, error) {
	val, err := GetAttributeInterface(obj, fieldName)
	if err != nil {
		return 0, err
	}
	return val.(float64), err
}

// SetAttributeFloat64Value will try to set member variable's float64 value
func SetAttributeFloat64Value(obj interface{}, fieldName string, newValue float64) error {
	return SetAttributeInterface(obj, fieldName, newValue)
}

// GetAttributeTimeValue will try to obtain member variable's time.Time value
func GetAttributeTimeValue(obj interface{}, fieldName string) (time.Time, error) {
	val, err := GetAttributeInterface(obj, fieldName)
	if err != nil {
		return time.Now(), err
	}
	return val.(time.Time), err
}

// SetAttributeTimeValue will try to set member variable's time.Time value
func SetAttributeTimeValue(obj interface{}, fieldName string, newValue time.Time) error {
	return SetAttributeInterface(obj, fieldName, newValue)
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
