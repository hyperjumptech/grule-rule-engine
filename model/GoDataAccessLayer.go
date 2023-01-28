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

package model

import (
	"fmt"
	"reflect"

	"github.com/hyperjumptech/grule-rule-engine/pkg"
)

// NewGoValueNode creates new instance of ValueNode backed by golang reflection
func NewGoValueNode(value reflect.Value, identifiedAs string) ValueNode {
	return &GoValueNode{
		parentNode:   nil,
		identifiedAs: identifiedAs,
		thisValue:    value,
	}
}

// GoValueNode is an implementation of ValueNode that used to traverse native golang primitives through reflect package
type GoValueNode struct {
	parentNode   ValueNode
	identifiedAs string
	thisValue    reflect.Value
}

// Value returns the underlying reflect.Value
func (node *GoValueNode) Value() reflect.Value {
	return node.thisValue
}

// HasParent returns `true` if the current value is a field, function, map, array, slice of another value
func (node *GoValueNode) HasParent() bool {
	return node.parentNode != nil
}

// Parent returns the value node of the parent value, if this node is a field, function, map, array, slice of another value
func (node *GoValueNode) Parent() ValueNode {
	return node.parentNode
}

// IdentifiedAs return the current representation of this Value Node
func (node *GoValueNode) IdentifiedAs() string {
	if node.HasParent() {
		if node.parentNode.IsArray() || node.parentNode.IsMap() {
			return fmt.Sprintf("%s%s", node.parentNode.IdentifiedAs(), node.identifiedAs)
		}
		return fmt.Sprintf("%s.%s", node.parentNode.IdentifiedAs(), node.identifiedAs)
	}
	return node.identifiedAs
}

// ContinueWithValue will return a nother ValueNode to wrap the specified value and treated as child of current node.
// The main purpose of this is for easier debugging.
func (node *GoValueNode) ContinueWithValue(value reflect.Value, identifiedAs string) ValueNode {
	return &GoValueNode{
		parentNode:   node,
		identifiedAs: identifiedAs,
		thisValue:    value,
	}
}

// GetValue will return the underlying reflect.Value
func (node *GoValueNode) GetValue() (reflect.Value, error) {
	return node.thisValue, nil
}

// GetType will return the underlying value's type
func (node *GoValueNode) GetType() (reflect.Type, error) {
	return node.thisValue.Type(), nil
}

// IsArray to check if the underlying value is an array or not
func (node *GoValueNode) IsArray() bool {
	return node.thisValue.Kind() == reflect.Array || node.thisValue.Kind() == reflect.Slice
}

func (node *GoValueNode) IsInterface() bool {
	return node.thisValue.Kind() == reflect.Interface
}

// GetArrayType to get the type of underlying value array element types.
func (node *GoValueNode) GetArrayType() (reflect.Type, error) {
	if node.IsArray() {
		return node.thisValue.Type().Elem(), nil
	}
	return nil, fmt.Errorf("this node identified as \"%s\" is not referring to an array or slice", node.IdentifiedAs())
}

// GetArrayValueAt to get the value of an array element if the current underlying value is an array
func (node *GoValueNode) GetArrayValueAt(index int) (val reflect.Value, err error) {
	if node.IsArray() {
		defer func() {
			if r := recover(); r != nil {
				err = fmt.Errorf("recovered : %v", r)
			}
		}()
		return node.thisValue.Index(index), err
	}
	return reflect.Value{}, fmt.Errorf("this node identified as \"%s\" is not referring to an array or slice", node.IdentifiedAs())
}

// GetChildNodeByIndex is similar to `GetArrayValueAt`, where this will return a ValueNode that wrap the value.
func (node *GoValueNode) GetChildNodeByIndex(index int) (ValueNode, error) {
	if node.IsArray() {
		v, err := node.GetArrayValueAt(index)
		if err != nil {
			return nil, err
		}
		gv := node.ContinueWithValue(v, fmt.Sprintf("[%d]", index))
		return gv, nil
	}
	return nil, fmt.Errorf("this node identified as \"%s\" is not an array. its %s", node.IdentifiedAs(), node.thisValue.Type().String())
}

// SetArrayValueAt will set the value of specified array index on the current underlying array value.
func (node *GoValueNode) SetArrayValueAt(index int, value reflect.Value) (err error) {
	if node.IsArray() {
		defer func() {
			if r := recover(); r != nil {
				err = fmt.Errorf("recovered : %v", r)
			}
		}()
		val := node.thisValue.Index(index)
		if val.CanAddr() && val.CanSet() {
			if pkg.IsNumber(val) && pkg.IsNumber(value) {
				return SetNumberValue(val, value)
			}
			val.Set(value)
			return nil
		}
		return fmt.Errorf("this node identified as \"%s\" can not set value on array index %d", node.IdentifiedAs(), index)
	}
	return fmt.Errorf("this node identified as \"%s\" is not referencing an array or slice", node.IdentifiedAs())
}

// AppendValue will append the new values into the current underlying array.
// will return error if argument list are not compatible with the array element type.
func (node *GoValueNode) AppendValue(value []reflect.Value) (err error) {
	if node.IsArray() {
		arrVal := node.thisValue
		if arrVal.CanSet() {
			defer func() {
				if r := recover(); r != nil {
					err = fmt.Errorf("recovered : %v", r)
				}
			}()
			arrVal.Set(reflect.Append(arrVal, value...))
			return nil
		}
	}
	return fmt.Errorf("this node identified as \"%s\" is not referencing an array or slice", node.IdentifiedAs())
}

// Length will return the length of underlying value if its an array, slice, map or string
func (node *GoValueNode) Length() (int, error) {
	if node.IsArray() || node.IsMap() || node.IsString() {
		return node.thisValue.Len(), nil
	}
	return 0, fmt.Errorf("this node identified as \"%s\" is not referencing an array, slice, map or string", node.IdentifiedAs())
}

// IsMap will validate if the underlying value is a map.
func (node *GoValueNode) IsMap() bool {
	return node.thisValue.Kind() == reflect.Map
}

// GetMapValueAt will retrieve a map value by the specified key argument.
func (node *GoValueNode) GetMapValueAt(index reflect.Value) (reflect.Value, error) {
	if node.IsMap() {
		retVal := node.thisValue.MapIndex(index)
		if retVal.IsValid() {
			return retVal, nil
		}
		return reflect.Value{}, fmt.Errorf("this node identified as \"%s\" have no selector with specified key", node.IdentifiedAs())
	}
	return reflect.Value{}, fmt.Errorf("this node identified as \"%s\" is not referencing a map", node.IdentifiedAs())
}

// SetMapValueAt will set the map value for the specified key, value argument
func (node *GoValueNode) SetMapValueAt(index, newValue reflect.Value) (err error) {
	if node.IsMap() {
		defer func() {
			if r := recover(); r != nil {
				err = fmt.Errorf("recovered : %v", r)
			}
		}()
		node.thisValue.SetMapIndex(index, newValue)
		return nil
	}
	return fmt.Errorf("this node identified as \"%s\" is not referencing a map", node.IdentifiedAs())
}

// GetChildNodeBySelector is similar to GetMapValueAt, it retrieve a value of map element identified by a value index as ValueNode.
func (node *GoValueNode) GetChildNodeBySelector(index reflect.Value) (ValueNode, error) {
	val, err := node.GetMapValueAt(index)
	if err != nil {
		return nil, err
	}
	return node.ContinueWithValue(val, fmt.Sprintf("[%s->%s]", index.Type().String(), index.String())), nil
}

// IsObject will check if the underlying value is a struct or pointer to a struct
func (node *GoValueNode) IsObject() bool {
	if node.thisValue.IsValid() {
		typ := node.thisValue.Type()
		if typ.Kind() == reflect.Ptr {
			return typ.Elem().Kind() == reflect.Struct
		}
		return typ.Kind() == reflect.Struct
	}
	return false
}

// GetObjectValueByField will return underlying value's field
func (node *GoValueNode) GetObjectValueByField(field string) (reflect.Value, error) {
	if node.IsObject() {
		var val reflect.Value
		if node.thisValue.Kind() == reflect.Ptr {
			val = node.thisValue.Elem().FieldByName(field)
		}
		if node.thisValue.Kind() == reflect.Struct {
			val = node.thisValue.FieldByName(field)
		}
		if val.IsValid() {
			return val, nil
		}
		return reflect.Value{}, fmt.Errorf("this node have no field named %s", field)
	}
	return reflect.Value{}, fmt.Errorf("this node identified as \"%s\" is not referencing to an object", node.IdentifiedAs())
}

// GetObjectTypeByField will return underlying type of the value's field
func (node *GoValueNode) GetObjectTypeByField(field string) (typ reflect.Type, err error) {
	if node.IsObject() {
		defer func() {
			if r := recover(); r != nil {
				err = fmt.Errorf("recovered : %v", r)
				typ = nil
			}
		}()
		if node.thisValue.Kind() == reflect.Ptr {
			return node.thisValue.Elem().FieldByName(field).Type(), nil
		}
		if node.thisValue.Kind() == reflect.Struct {
			return node.thisValue.FieldByName(field).Type(), nil
		}
	}
	return nil, fmt.Errorf("this node identified as \"%s\" is not referring to an object", node.IdentifiedAs())
}

// SetNumberValue will assign a numeric value to a numeric target value
// this helper function is to ensure assignment between numerical types is happening regardless of types, int, uint or float.
// The rule designer should be careful as conversion of types in automatic way like this will cause lost of precision
// during conversion. This will be removed in the future version.
func SetNumberValue(target, newvalue reflect.Value) error {
	if pkg.IsNumber(target) && pkg.IsNumber(newvalue) {
		switch target.Type().Kind() {
		case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
			if pkg.GetBaseKind(newvalue) == reflect.Uint64 {
				target.SetInt(int64(newvalue.Uint()))
			} else if pkg.GetBaseKind(newvalue) == reflect.Float64 {
				target.SetInt(int64(newvalue.Float()))
			} else {
				target.SetInt(newvalue.Int())
			}
			return nil
		case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
			if pkg.GetBaseKind(newvalue) == reflect.Uint64 {
				target.SetUint(newvalue.Uint())
			} else if pkg.GetBaseKind(newvalue) == reflect.Float64 {
				target.SetUint(uint64(newvalue.Float()))
			} else {
				target.SetUint(uint64(newvalue.Int()))
			}
			return nil
		case reflect.Float32, reflect.Float64:
			if pkg.GetBaseKind(newvalue) == reflect.Uint64 {
				target.SetFloat(float64(newvalue.Uint()))
			} else if pkg.GetBaseKind(newvalue) == reflect.Float64 {
				target.SetFloat(newvalue.Float())
			} else {
				target.SetFloat(float64(newvalue.Int()))
			}
			return nil
		}
		return fmt.Errorf("this line should not be reached")
	}
	return fmt.Errorf("this function only used for assigning number data to number variable")
}

// SetObjectValueByField will set the underlying value's field with new value.
func (node *GoValueNode) SetObjectValueByField(field string, newValue reflect.Value) (err error) {
	fieldVal := node.thisValue.Elem().FieldByName(field)
	if fieldVal.IsValid() && fieldVal.CanAddr() && fieldVal.CanSet() {
		defer func() {
			if r := recover(); r != nil {
				err = fmt.Errorf("recovered : %v", r)
			}
		}()
		if pkg.IsNumber(fieldVal) && pkg.IsNumber(newValue) {
			return SetNumberValue(fieldVal, newValue)
		}
		fieldVal.Set(newValue)
		return nil
	}
	return fmt.Errorf("this node identified as \"%s\" have field \"%s\" that is not valid nor addressable", node.IdentifiedAs(), field)
}

// CallFunction will call a function owned by the underlying value receiver.
// this function will artificially create a built-in functions for constants, array and map.
func (node *GoValueNode) CallFunction(funcName string, args ...reflect.Value) (retval reflect.Value, err error) {
	switch pkg.GetBaseKind(node.thisValue) {
	case reflect.Int64, reflect.Uint64, reflect.Float64, reflect.Bool:
		return reflect.ValueOf(nil), fmt.Errorf("this node identified as \"%s\" try to call function %s which is not supported for type %s", node.IdentifiedAs(), funcName, node.thisValue.Type().String())
	case reflect.String:
		var strfunc func(string, []reflect.Value) (reflect.Value, error)
		switch funcName {
		case "In":
			strfunc = StrIn
		case "Compare":
			strfunc = StrCompare
		case "Contains":
			strfunc = StrContains
		case "Count":
			strfunc = StrCount
		case "HasPrefix":
			strfunc = StrHasPrefix
		case "HasSuffix":
			strfunc = StrHasSuffix
		case "Index":
			strfunc = StrIndex
		case "LastIndex":
			strfunc = StrLastIndex
		case "Repeat":
			strfunc = StrRepeat
		case "Replace":
			strfunc = StrReplace
		case "Split":
			strfunc = StrSplit
		case "ToLower":
			strfunc = StrToLower
		case "ToUpper":
			strfunc = StrToUpper
		case "Trim":
			strfunc = StrTrim
		case "Len":
			strfunc = StrLen
		case "MatchString":
			strfunc = StrMatchRegexPattern
		}
		if strfunc != nil {
			val, err := strfunc(node.thisValue.String(), args)
			if err != nil {
				return reflect.Value{}, err
			}
			return val, nil
		}
		return reflect.Value{}, fmt.Errorf("this node identified as \"%s\" call function %s is not supported for string", node.IdentifiedAs(), funcName)
	}
	if node.IsArray() {
		var arrFunc func(reflect.Value, []reflect.Value) (reflect.Value, error)
		switch funcName {
		case "Len":
			arrFunc = ArrMapLen
		case "Append":
			node.AppendValue(args)
			return reflect.Value{}, nil
		}
		if arrFunc != nil {
			if funcName == "Clear" {
				val, err := arrFunc(node.thisValue, args)
				if err != nil {
					return reflect.Value{}, err
				}
				return val, nil
			}
			val, err := arrFunc(node.thisValue, args)
			if err != nil {
				return reflect.Value{}, err
			}
			return val, nil
		}
		return reflect.Value{}, fmt.Errorf("this node identified as \"%s\" call function %s is not supported for array", node.IdentifiedAs(), funcName)
	}
	if node.IsMap() {
		var mapFunc func(reflect.Value, []reflect.Value) (reflect.Value, error)
		switch funcName {
		case "Len":
			mapFunc = ArrMapLen
		case "EqualValues":
			mapFunc = MapEqualValues
		case "CountValue":
			mapFunc = MapCountValue
		}
		if mapFunc != nil {
			val, err := mapFunc(node.thisValue, args)
			if err != nil {
				return reflect.Value{}, err
			}
			return val, nil
		}
		return reflect.Value{}, fmt.Errorf("this node identified as \"%s\" call function %s is not supported for map", node.IdentifiedAs(), funcName)
	}

	if node.IsObject() || node.IsInterface() {
		funcValue := node.thisValue.MethodByName(funcName)
		if funcValue.IsValid() {
			rets := funcValue.Call(args)
			if len(rets) > 1 {
				return reflect.Value{}, fmt.Errorf("this node identified as \"%s\" calling function %s which returns multiple values, multiple value returns are not supported", node.IdentifiedAs(), funcName)
			}
			if len(rets) == 1 {
				return rets[0], nil
			}
			return reflect.Value{}, nil
		}
		return reflect.Value{}, fmt.Errorf("this node identified as \"%s\" have no function named %s", node.IdentifiedAs(), funcName)
	}
	return reflect.ValueOf(nil), fmt.Errorf("this node identified as \"%s\" is not referencing an object thus function %s call is not supported. Kind %s", node.IdentifiedAs(), funcName, node.thisValue.Kind().String())
}

// GetChildNodeByField will retrieve the underlying struct's field and return the ValueNode wraper.
func (node *GoValueNode) GetChildNodeByField(field string) (ValueNode, error) {
	val, err := node.GetObjectValueByField(field)
	if err != nil {
		return nil, err
	}
	return node.ContinueWithValue(val, field), nil
}

// IsTime will check if the underlying value is a time.Time
func (node *GoValueNode) IsTime() bool {
	return node.thisValue.Type().String() == "time.Time"
}

// IsInteger will check if the underlying value is a type of int, or uint
func (node *GoValueNode) IsInteger() bool {
	kind := pkg.GetBaseKind(node.thisValue)
	return kind == reflect.Int64 || kind == reflect.Uint64
}

// IsReal will check if the underlying value is a type of real number, float.
func (node *GoValueNode) IsReal() bool {
	kind := pkg.GetBaseKind(node.thisValue)
	return kind == reflect.Float64
}

// IsBool will check if the underlying value is a type of boolean.
func (node *GoValueNode) IsBool() bool {
	return node.thisValue.Kind() == reflect.Bool
}

// IsString will check if the underlying value is a type of string
func (node *GoValueNode) IsString() bool {
	return node.thisValue.Kind() == reflect.String
}
