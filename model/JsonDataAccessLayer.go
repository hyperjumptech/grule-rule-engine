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
	"encoding/json"
	"fmt"
	"github.com/hyperjumptech/grule-rule-engine/pkg"
	"reflect"
	"time"
)

var (
	// DateTimeLayout contains the date time layouting used by this data access layer.
	DateTimeLayout = time.RFC3339
)

// NewJSONValueNode will create a new ValueNode structure backend using data structure as provided by JSON parser.
func NewJSONValueNode(JSONString, identifiedAs string) (ValueNode, error) {
	var d interface{}
	err := json.Unmarshal([]byte(JSONString), &d)
	if err != nil {
		return nil, err
	}
	return &JSONValueNode{
		parent:       nil,
		identifiedAs: identifiedAs,
		data:         reflect.ValueOf(d),
	}, nil
}

// JSONValueNode will hold the json root object as the result of JSON unmarshal
type JSONValueNode struct {
	parent       ValueNode
	identifiedAs string
	data         reflect.Value
}

// IdentifiedAs will return the node label
func (vn *JSONValueNode) IdentifiedAs() string {
	return vn.identifiedAs
}

// Value returns the reflect.Value of this node
func (vn *JSONValueNode) Value() reflect.Value {
	return vn.data
}

// HasParent will return true if this node has parent node, other wise return false
func (vn *JSONValueNode) HasParent() bool {
	return vn.parent != nil
}

// Parent returns the parent node of this node.
func (vn *JSONValueNode) Parent() ValueNode {
	return vn.parent
}

// ContinueWithValue will return a new node contains the specified value and the parent will be this value.
func (vn *JSONValueNode) ContinueWithValue(value reflect.Value, identifiedAs string) ValueNode {
	return &JSONValueNode{
		parent:       vn,
		identifiedAs: identifiedAs,
		data:         value,
	}
}

// GetValue same as Value()
func (vn *JSONValueNode) GetValue() (reflect.Value, error) {
	return vn.data, nil
}

// GetType return the reflect.Type of the value in this node
func (vn *JSONValueNode) GetType() (reflect.Type, error) {
	return vn.data.Type(), nil
}

// IsArray will validate if this node's value is of kind Array or Slice
func (vn *JSONValueNode) IsArray() bool {
	return vn.data.Kind() == reflect.Slice || vn.data.Kind() == reflect.Array
}

// GetArrayType return the content of an array. Since json array can contain any type, it will aways return type of nil.
func (vn *JSONValueNode) GetArrayType() (reflect.Type, error) {
	return reflect.TypeOf(nil), nil
}

// GetArrayValueAt return the value of array element specified by index. It will return error if this node is not array or slice.
func (vn *JSONValueNode) GetArrayValueAt(index int) (reflect.Value, error) {
	if vn.IsArray() {
		return vn.data.Index(index).Elem(), nil
	}
	return reflect.ValueOf(nil), fmt.Errorf("this node identified as \"%s\" is not an array. its %s", vn.IdentifiedAs(), vn.data.Type().String())
}

// GetChildNodeByIndex will return the array node of array element specified by index. It will return error if its not an array nor slice.
func (vn *JSONValueNode) GetChildNodeByIndex(index int) (ValueNode, error) {
	val, err := vn.GetArrayValueAt(index)
	if err != nil {
		return nil, err
	}
	return vn.ContinueWithValue(val, fmt.Sprintf("[%d]", index)), nil
}

// SetArrayValueAt sets this node array element specified at index with new value. User should be careful to not set element with out of bound index.
// It will return error if its not an array nor slice.
func (vn *JSONValueNode) SetArrayValueAt(index int, value reflect.Value) error {
	itv := vn.data.Index(index)
	itv.Set(value)
	return nil
}

// AppendValue will append an array of reflect.Value(s) into the end of this array/slice node.
// It will return error if its not an array nor slice.
func (vn *JSONValueNode) AppendValue(value []reflect.Value) error {
	if !vn.IsArray() {
		return fmt.Errorf("not an array or slice")
	}
	vn.data = reflect.Append(vn.data, value...)
	return nil
}

// Length return the length of this node. It will return error if not type of string, map, array/slice or object.
func (vn *JSONValueNode) Length() (l int, e error) {
	defer func() {
		if r := recover(); r != nil {
			l = 0
			e = fmt.Errorf("can not get the length of value other than string, map, array/slice or object")
		}
	}()
	return vn.data.Len(), nil
}

// IsMap will validate if this node is a map.
func (vn *JSONValueNode) IsMap() bool {
	return vn.data.Kind() == reflect.Map
}

// GetMapValueAt get the value of this map node at specific index/selector value.
// In json, the index selector must be of type of string.
func (vn *JSONValueNode) GetMapValueAt(index reflect.Value) (reflect.Value, error) {
	if !vn.IsMap() {
		return reflect.ValueOf(nil), fmt.Errorf("not a map")
	}
	if index.Kind() != reflect.String {
		return reflect.ValueOf(nil), fmt.Errorf("JSON map selector must be a string")
	}
	tmap := vn.data.MapIndex(index)
	return tmap.Elem(), nil
}

// SetMapValueAt set the value in this map as specific index/selector value.
// In json, the index selector must be of type of string.
func (vn *JSONValueNode) SetMapValueAt(index, newValue reflect.Value) error {
	if !vn.IsMap() {
		return fmt.Errorf("not an object or map")
	}
	vn.data.SetMapIndex(index, newValue)
	return nil
}

// GetChildNodeBySelector get the ValueNode
func (vn *JSONValueNode) GetChildNodeBySelector(index reflect.Value) (ValueNode, error) {
	val, err := vn.GetMapValueAt(index)
	if err != nil {
		return nil, err
	}
	return vn.ContinueWithValue(val, fmt.Sprintf("[%s]", index.String())), nil
}

// IsObject returns true if this node is an object or map.
func (vn *JSONValueNode) IsObject() bool {
	return vn.data.Kind() == reflect.Map
}

// GetObjectValueByField get the value of this node by the specified field.
func (vn *JSONValueNode) GetObjectValueByField(field string) (reflect.Value, error) {
	if !vn.IsObject() {
		return reflect.ValueOf(nil), fmt.Errorf("not an object or map")
	}
	tmap := vn.data.MapIndex(reflect.ValueOf(field))
	return tmap.Elem(), nil
}

// GetObjectTypeByField get the type of the value by specified field. Since in json any field could store any field and
// there are no definition of what type on any field, this function will always return value of nil
func (vn *JSONValueNode) GetObjectTypeByField(field string) (reflect.Type, error) {
	return reflect.TypeOf(nil), nil
}

// SetObjectValueByField set the value in the node by specified field name.
func (vn *JSONValueNode) SetObjectValueByField(field string, newValue reflect.Value) error {
	if !vn.IsObject() {
		return fmt.Errorf("not an object or map")
	}
	vn.data.SetMapIndex(reflect.ValueOf(field), newValue)
	return nil
}

// CallFunction will always return an error, as Json data do not have any function in them.
func (vn *JSONValueNode) CallFunction(funcName string, args ...reflect.Value) (reflect.Value, error) {
	switch pkg.GetBaseKind(vn.data) {
	case reflect.Int64, reflect.Uint64, reflect.Float64, reflect.Bool:
		return reflect.ValueOf(nil), fmt.Errorf("this node identified as \"%s\" try to call function %s which is not supported for type %s", vn.IdentifiedAs(), funcName, vn.data.Type().String())
	case reflect.String:
		var strfunc func(string, []reflect.Value) (reflect.Value, error)
		switch funcName {
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
		}
		if strfunc != nil {
			val, err := strfunc(vn.data.String(), args)
			if err != nil {
				return reflect.Value{}, err
			}
			return val, nil
		}
		return reflect.Value{}, fmt.Errorf("this node identified as \"%s\" call function %s is not supported for string", vn.IdentifiedAs(), funcName)
	}
	if vn.IsArray() {
		var arrFunc func(reflect.Value, []reflect.Value) (reflect.Value, error)
		switch funcName {
		case "Len":
			arrFunc = ArrMapLen
		case "Append":
			vn.AppendValue(args)
			return reflect.Value{}, nil
		}
		if arrFunc != nil {
			if funcName == "Clear" {
				val, err := arrFunc(vn.data, args)
				if err != nil {
					return reflect.Value{}, err
				}
				return val, nil
			}
			val, err := arrFunc(vn.data, args)
			if err != nil {
				return reflect.Value{}, err
			}
			return val, nil
		}
		return reflect.Value{}, fmt.Errorf("this node identified as \"%s\" call function %s is not supported for array", vn.IdentifiedAs(), funcName)
	}
	if vn.IsMap() {
		var mapFunc func(reflect.Value, []reflect.Value) (reflect.Value, error)
		switch funcName {
		case "Len":
			mapFunc = ArrMapLen
		}
		if mapFunc != nil {
			val, err := mapFunc(vn.data, args)
			if err != nil {
				return reflect.Value{}, err
			}
			return val, nil
		}
		return reflect.Value{}, fmt.Errorf("this node identified as \"%s\" call function %s is not supported for map", vn.IdentifiedAs(), funcName)
	}

	if vn.IsObject() {
		funcValue := vn.data.MethodByName(funcName)
		if funcValue.IsValid() {
			rets := funcValue.Call(args)
			if len(rets) > 1 {
				return reflect.Value{}, fmt.Errorf("this node identified as \"%s\" calling function %s which returns multiple values, multiple value returns are not supported", vn.IdentifiedAs(), funcName)
			}
			if len(rets) == 1 {
				return rets[0], nil
			}
			return reflect.Value{}, nil
		}
		return reflect.Value{}, fmt.Errorf("this node identified as \"%s\" have no function named %s", vn.IdentifiedAs(), funcName)
	}
	return reflect.ValueOf(nil), fmt.Errorf("this node identified as \"%s\" is not referencing an object thus function %s call is not supported", vn.IdentifiedAs(), funcName)
}

// GetChildNodeByField will return the field ValueNode
func (vn *JSONValueNode) GetChildNodeByField(field string) (ValueNode, error) {
	val, err := vn.GetObjectValueByField(field)
	if err != nil {
		return nil, err
	}
	return vn.ContinueWithValue(val, field), nil
}

// IsTime return true if the value of this node is of type string with specified DateTimeLayout
func (vn *JSONValueNode) IsTime() bool {
	if vn.data.Kind() == reflect.String {
		return IsDateFormatValid(DateTimeLayout, vn.data.String())
	}
	return false
}

// IsInteger return true if the value of this node is conform to an integer. (no floating point value).
func (vn *JSONValueNode) IsInteger() bool {
	if vn.data.Kind() == reflect.Float64 {
		f := vn.data.Float()
		i := int64(f)
		f2 := float64(i)
		return f == f2
	}
	return false
}

// IsReal return true if the value of this node contains integer or floating point.
func (vn *JSONValueNode) IsReal() bool {
	return vn.data.Kind() == reflect.Float64
}

// IsBool return true if the value of this node contains a boolean.
func (vn *JSONValueNode) IsBool() bool {
	return vn.data.Kind() == reflect.Bool
}

// IsString returns true if the value of this node contains a string.
func (vn *JSONValueNode) IsString() bool {
	return vn.data.Kind() == reflect.String
}
