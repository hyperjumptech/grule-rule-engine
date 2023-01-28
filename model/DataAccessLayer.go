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
	"github.com/hyperjumptech/grule-rule-engine/pkg"
	"reflect"
	"regexp"
	"strings"
)

// ValueNode is an abstraction layer to access underlying dom style data.
// the node have tree kind of structure which each node are tied to an underlying data node.
type ValueNode interface {
	IdentifiedAs() string
	Value() reflect.Value
	HasParent() bool
	Parent() ValueNode

	ContinueWithValue(value reflect.Value, identifiedAs string) ValueNode
	GetValue() (reflect.Value, error)
	GetType() (reflect.Type, error)

	IsArray() bool
	GetArrayType() (reflect.Type, error)
	GetArrayValueAt(index int) (reflect.Value, error)
	GetChildNodeByIndex(index int) (ValueNode, error)
	SetArrayValueAt(index int, value reflect.Value) error
	AppendValue(value []reflect.Value) error
	Length() (int, error)

	IsMap() bool
	GetMapValueAt(index reflect.Value) (reflect.Value, error)
	SetMapValueAt(index, newValue reflect.Value) error
	GetChildNodeBySelector(index reflect.Value) (ValueNode, error)

	IsInterface() bool
	IsObject() bool
	GetObjectValueByField(field string) (reflect.Value, error)
	GetObjectTypeByField(field string) (reflect.Type, error)
	SetObjectValueByField(field string, newValue reflect.Value) error
	CallFunction(funcName string, args ...reflect.Value) (reflect.Value, error)
	GetChildNodeByField(field string) (ValueNode, error)

	IsTime() bool
	IsInteger() bool
	IsReal() bool
	IsBool() bool
	IsString() bool
}

// StrLen is return the string length value
func StrLen(str string, arg []reflect.Value) (reflect.Value, error) {
	if arg != nil && len(arg) != 0 {
		return reflect.ValueOf(nil), fmt.Errorf("function Len requires no argument")
	}
	i := len(str)
	return reflect.ValueOf(i), nil
}

// StrCompare is like strings.compare() function, to be called by the ValueNode function call if the underlying data is string.
func StrCompare(str string, arg []reflect.Value) (reflect.Value, error) {
	if arg == nil || len(arg) != 1 || arg[0].Kind() != reflect.String {
		return reflect.ValueOf(nil), fmt.Errorf("function Compare requires 1 string argument")
	}
	i := strings.Compare(str, arg[0].String())
	return reflect.ValueOf(i), nil
}

// StrContains is like strings.Contains() function, to be called by the ValueNode function call if the underlying data is string. is like strings.compare() function, to be called by the ValueNode functioncall if the underlying data is string.
func StrContains(str string, arg []reflect.Value) (reflect.Value, error) {
	if arg == nil || len(arg) != 1 || arg[0].Kind() != reflect.String {
		return reflect.ValueOf(nil), fmt.Errorf("function Contains requires 1 string argument")
	}

	i := strings.Contains(str, arg[0].String())
	return reflect.ValueOf(i), nil
}

// StrCount is like strings.Count() function, to be called by the ValueNode function call if the underlying data is string.
func StrCount(str string, arg []reflect.Value) (reflect.Value, error) {
	if arg == nil || len(arg) != 1 || arg[0].Kind() != reflect.String {
		return reflect.ValueOf(nil), fmt.Errorf("function Count requires 1 string argument")
	}

	i := strings.Count(str, arg[0].String())
	return reflect.ValueOf(i), nil
}

// StrHasPrefix is like strings.HasPrefix() function, to be called by the ValueNode functioncall if the underlying data is string.
func StrHasPrefix(str string, arg []reflect.Value) (reflect.Value, error) {
	if arg == nil || len(arg) != 1 || arg[0].Kind() != reflect.String {
		return reflect.ValueOf(nil), fmt.Errorf("function HasPrefix requires 1 string argument")
	}

	b := strings.HasPrefix(str, arg[0].String())
	return reflect.ValueOf(b), nil
}

// StrHasSuffix is like strings.HasSuffix() function, to be called by the ValueNode functioncall if the underlying data is string.
func StrHasSuffix(str string, arg []reflect.Value) (reflect.Value, error) {
	if arg == nil || len(arg) != 1 || arg[0].Kind() != reflect.String {
		return reflect.ValueOf(nil), fmt.Errorf("function HasSuffix requires 1 string argument")
	}

	b := strings.HasSuffix(str, arg[0].String())
	return reflect.ValueOf(b), nil
}

// StrIndex is like strings.Index() function, to be called by the ValueNode functioncall if the underlying data is string.
func StrIndex(str string, arg []reflect.Value) (reflect.Value, error) {
	if arg == nil || len(arg) != 1 || arg[0].Kind() != reflect.String {
		return reflect.ValueOf(nil), fmt.Errorf("function Index requires 1 string argument")
	}

	b := strings.Index(str, arg[0].String())
	return reflect.ValueOf(b), nil
}

// StrLastIndex is like strings.LastIndex() function, to be called by the ValueNode functioncall if the underlying data is string.
func StrLastIndex(str string, arg []reflect.Value) (reflect.Value, error) {
	if arg == nil || len(arg) != 1 || arg[0].Kind() != reflect.String {
		return reflect.ValueOf(nil), fmt.Errorf("function LastIndex requires 1 string argument")
	}

	b := strings.LastIndex(str, arg[0].String())
	return reflect.ValueOf(b), nil
}

// StrRepeat is like strings.Repeat() function, to be called by the ValueNode functioncall if the underlying data is string.
func StrRepeat(str string, arg []reflect.Value) (reflect.Value, error) {
	if arg == nil || len(arg) != 1 {
		return reflect.ValueOf(nil), fmt.Errorf("function Repeat requires 1 numeric argument")
	}
	repeat := 0
	switch pkg.GetBaseKind(arg[0]) {
	case reflect.Int64:
		repeat = int(arg[0].Int())
	case reflect.Uint64:
		repeat = int(arg[0].Uint())
	case reflect.Float64:
		repeat = int(arg[0].Float())
	default:
		return reflect.ValueOf(nil), fmt.Errorf("function Repeat requires 1 numeric argument")
	}

	b := strings.Repeat(str, repeat)
	return reflect.ValueOf(b), nil
}

// StrReplace is like strings.Replace() function, to be called by the ValueNode functioncall if the underlying data is string.
func StrReplace(str string, arg []reflect.Value) (reflect.Value, error) {
	if arg == nil || len(arg) != 2 || arg[0].Kind() != reflect.String || arg[1].Kind() != reflect.String {
		return reflect.ValueOf(nil), fmt.Errorf("function Cmpare requires 2 string argument")
	}

	b := strings.ReplaceAll(str, arg[0].String(), arg[1].String())
	return reflect.ValueOf(b), nil
}

// StrSplit is like strings.Split() function, to be called by the ValueNode functioncall if the underlying data is string.
func StrSplit(str string, arg []reflect.Value) (reflect.Value, error) {
	if arg == nil || len(arg) != 1 || arg[0].Kind() != reflect.String {
		return reflect.ValueOf(nil), fmt.Errorf("function Split requires 1 string argument")
	}

	b := strings.Split(str, arg[0].String())
	return reflect.ValueOf(b), nil
}

// StrToLower is like strings.ToLower() function, to be called by the ValueNode functioncall if the underlying data is string.
func StrToLower(str string, arg []reflect.Value) (reflect.Value, error) {
	if arg != nil && len(arg) != 0 {
		return reflect.ValueOf(nil), fmt.Errorf("function ToLower requires no argument")
	}
	b := strings.ToLower(str)
	return reflect.ValueOf(b), nil
}

// StrToUpper is like strings.ToUpper() function, to be called by the ValueNode functioncall if the underlying data is string.
func StrToUpper(str string, arg []reflect.Value) (reflect.Value, error) {
	if arg != nil && len(arg) != 0 {
		return reflect.ValueOf(nil), fmt.Errorf("function ToUpper requires no argument")
	}
	b := strings.ToUpper(str)
	return reflect.ValueOf(b), nil
}

// StrTrim is like strings.Trim() function, to be called by the ValueNode functioncall if the underlying data is string.
func StrTrim(str string, arg []reflect.Value) (reflect.Value, error) {
	if arg != nil && len(arg) != 0 {
		return reflect.ValueOf(nil), fmt.Errorf("function Trim requires no argument")
	}
	b := strings.TrimSpace(str)
	return reflect.ValueOf(b), nil
}

// StrIn will check the string instance if its equals one of the arguments, if no argument specified it will return false
func StrIn(str string, arg []reflect.Value) (reflect.Value, error) {
	for _, a := range arg {
		if !a.IsValid() || a.Kind() != reflect.String {
			return reflect.ValueOf(nil), fmt.Errorf("function StrIn requires string arguments")
		}
		if a.String() == str {
			return reflect.ValueOf(true), nil
		}
	}
	return reflect.ValueOf(false), nil
}

// StrMatchRegexPattern reports whether the string s contains any match of the regular expression pattern.
func StrMatchRegexPattern(str string, arg []reflect.Value) (reflect.Value, error) {
	if arg == nil || len(arg) != 1 || arg[0].Kind() != reflect.String {
		return reflect.ValueOf(nil), fmt.Errorf("function StrMatchRegexPattern requires 1 string argument")
	}
	m, err := regexp.MatchString(arg[0].String(), str)
	if err != nil {
		return reflect.ValueOf(nil), fmt.Errorf("function StrMatchRegexPattern requires valid regex pattern")
	}
	return reflect.ValueOf(m), nil
}

// ArrMapLen will return the size of underlying map, array or slice
func ArrMapLen(arr reflect.Value, arg []reflect.Value) (reflect.Value, error) {
	if arg != nil && len(arg) != 0 {
		return reflect.ValueOf(nil), fmt.Errorf("function Len requires no argument")
	}
	return reflect.ValueOf(arr.Len()), nil
}

// MapEqualValues will check all values are equal
func MapEqualValues(val reflect.Value, arg []reflect.Value) (reflect.Value, error) {
	if arg != nil && len(arg) != 0 {
		return reflect.ValueOf(false), fmt.Errorf("function EqualValues requires no argument")
	}
	m, ok := val.Interface().(map[string]string)
	if !ok || len(m) == 0 {
		return reflect.ValueOf(false), fmt.Errorf("function EqualValues requires map[string]string")
	}

	last := ""
	for _, v := range m {
		if last == "" {
			last = v
		} else if last != v {
			return reflect.ValueOf(false), nil
		}
	}

	return reflect.ValueOf(true), nil
}

// MapCountValue will return number of value which is equal to arg[0]
func MapCountValue(val reflect.Value, arg []reflect.Value) (reflect.Value, error) {
	if arg == nil || len(arg) != 1 || arg[0].Kind() != reflect.String {
		return reflect.ValueOf(0), fmt.Errorf("function MapCountValue requires no argument")
	}
	m, ok := val.Interface().(map[string]string)
	if !ok || len(m) == 0 {
		return reflect.ValueOf(0), fmt.Errorf("function MapCountValue requires map[string]string")
	}

	to := arg[0].String()
	cnt := int(0)
	for _, v := range m {
		if v == to {
			cnt++
		}
	}

	return reflect.ValueOf(cnt), nil
}
