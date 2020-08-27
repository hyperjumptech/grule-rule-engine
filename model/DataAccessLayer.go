package model

import (
	"fmt"
	"github.com/hyperjumptech/grule-rule-engine/pkg"
	"reflect"
	"strings"
)

type ValueNode interface {
	IdentifiedAs() string
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
	AppendValue(value reflect.Value) error
	Length() (int, error)

	IsMap() bool
	GetMapValueAt(index reflect.Value) (reflect.Value, error)
	SetMapValueAt(index, newValue reflect.Value) error
	GetChildNodeBySelector(index reflect.Value) (ValueNode, error)

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

func StrCompare(str string, arg []reflect.Value) (reflect.Value, error) {
	if arg == nil || len(arg) != 1 || arg[0].Kind() != reflect.String {
		return reflect.ValueOf(nil), fmt.Errorf("function Compare requires 1 string argument")
	}
	i := strings.Compare(str, arg[0].String())
	return reflect.ValueOf(i), nil
}

func StrContains(str string, arg []reflect.Value) (reflect.Value, error) {
	if arg == nil || len(arg) != 1 || arg[0].Kind() != reflect.String {
		return reflect.ValueOf(nil), fmt.Errorf("function Contains requires 1 string argument")
	}

	i := strings.Contains(str, arg[0].String())
	return reflect.ValueOf(i), nil
}

func StrCount(str string, arg []reflect.Value) (reflect.Value, error) {
	if arg == nil || len(arg) != 1 || arg[0].Kind() != reflect.String {
		return reflect.ValueOf(nil), fmt.Errorf("function Count requires 1 string argument")
	}

	i := strings.Count(str, arg[0].String())
	return reflect.ValueOf(i), nil
}

func StrHasPrefix(str string, arg []reflect.Value) (reflect.Value, error) {
	if arg == nil || len(arg) != 1 || arg[0].Kind() != reflect.String {
		return reflect.ValueOf(nil), fmt.Errorf("function HasPrefix requires 1 string argument")
	}

	b := strings.HasPrefix(str, arg[0].String())
	return reflect.ValueOf(b), nil
}

func StrHasSuffix(str string, arg []reflect.Value) (reflect.Value, error) {
	if arg == nil || len(arg) != 1 || arg[0].Kind() != reflect.String {
		return reflect.ValueOf(nil), fmt.Errorf("function HasSuffix requires 1 string argument")
	}

	b := strings.HasSuffix(str, arg[0].String())
	return reflect.ValueOf(b), nil
}

func StrIndex(str string, arg []reflect.Value) (reflect.Value, error) {
	if arg == nil || len(arg) != 1 || arg[0].Kind() != reflect.String {
		return reflect.ValueOf(nil), fmt.Errorf("function Index requires 1 string argument")
	}

	b := strings.Index(str, arg[0].String())
	return reflect.ValueOf(b), nil
}

func StrLastIndex(str string, arg []reflect.Value) (reflect.Value, error) {
	if arg == nil || len(arg) != 1 || arg[0].Kind() != reflect.String {
		return reflect.ValueOf(nil), fmt.Errorf("function LastIndex requires 1 string argument")
	}

	b := strings.LastIndex(str, arg[0].String())
	return reflect.ValueOf(b), nil
}

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

func StrReplace(str string, arg []reflect.Value) (reflect.Value, error) {
	if arg == nil || len(arg) != 2 || arg[0].Kind() != reflect.String || arg[1].Kind() != reflect.String {
		return reflect.ValueOf(nil), fmt.Errorf("function Cmpare requires 2 string argument")
	}

	b := strings.ReplaceAll(str, arg[0].String(), arg[1].String())
	return reflect.ValueOf(b), nil
}

func StrSplit(str string, arg []reflect.Value) (reflect.Value, error) {
	if arg == nil || len(arg) != 1 || arg[0].Kind() != reflect.String {
		return reflect.ValueOf(nil), fmt.Errorf("function Split requires 1 string argument")
	}

	b := strings.Split(str, arg[0].String())
	return reflect.ValueOf(b), nil
}

func StrToLower(str string, arg []reflect.Value) (reflect.Value, error) {
	if arg != nil && len(arg) != 0 {
		return reflect.ValueOf(nil), fmt.Errorf("function ToLower requires no argument")
	}
	b := strings.ToLower(str)
	return reflect.ValueOf(b), nil
}

func StrToUpper(str string, arg []reflect.Value) (reflect.Value, error) {
	if arg != nil && len(arg) != 0 {
		return reflect.ValueOf(nil), fmt.Errorf("function ToUpper requires no argument")
	}
	b := strings.ToUpper(str)
	return reflect.ValueOf(b), nil

}

func StrTrim(str string, arg []reflect.Value) (reflect.Value, error) {
	if arg != nil && len(arg) != 0 {
		return reflect.ValueOf(nil), fmt.Errorf("function Trim requires no argument")
	}
	b := strings.TrimSpace(str)
	return reflect.ValueOf(b), nil
}
