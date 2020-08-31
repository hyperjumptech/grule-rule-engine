package model

import (
	"fmt"
	"github.com/hyperjumptech/grule-rule-engine/pkg"
	"reflect"
)

func NewGoValueNode(value reflect.Value, identifiedAs string) ValueNode {
	return &GoValueNode{
		parentNode:   nil,
		identifiedAs: identifiedAs,
		thisValue:    value,
	}
}

type GoValueNode struct {
	parentNode   ValueNode
	identifiedAs string
	thisValue    reflect.Value
}

func (node *GoValueNode) Value() reflect.Value {
	return node.thisValue
}

func (node *GoValueNode) HasParent() bool {
	return node.parentNode != nil
}

func (node *GoValueNode) Parent() ValueNode {
	return node.parentNode
}

func (node *GoValueNode) IdentifiedAs() string {
	if node.HasParent() {
		if node.parentNode.IsArray() || node.parentNode.IsMap() {
			return fmt.Sprintf("%s%s", node.parentNode.IdentifiedAs(), node.identifiedAs)
		}
		return fmt.Sprintf("%s.%s", node.parentNode.IdentifiedAs(), node.identifiedAs)
	}
	return node.identifiedAs
}
func (node *GoValueNode) ContinueWithValue(value reflect.Value, identifiedAs string) ValueNode {
	return &GoValueNode{
		parentNode:   node,
		identifiedAs: identifiedAs,
		thisValue:    value,
	}
}
func (node *GoValueNode) GetValue() (reflect.Value, error) {
	return node.thisValue, nil
}
func (node *GoValueNode) GetType() (reflect.Type, error) {
	return node.thisValue.Type(), nil
}
func (node *GoValueNode) IsArray() bool {
	return node.thisValue.Kind() == reflect.Array || node.thisValue.Kind() == reflect.Slice
}
func (node *GoValueNode) GetArrayType() (reflect.Type, error) {
	if node.IsArray() {
		return node.thisValue.Type().Elem(), nil
	}
	return nil, fmt.Errorf("this node identified as \"%s\" is not referring to an array or slice", node.IdentifiedAs())
}
func (node *GoValueNode) GetArrayValueAt(index int) (reflect.Value, error) {
	if node.IsArray() {
		return node.thisValue.Index(index), nil
	}
	return reflect.Value{}, fmt.Errorf("this node identified as \"%s\" is not referring to an array or slice", node.IdentifiedAs())
}
func (node *GoValueNode) GetChildNodeByIndex(index int) (ValueNode, error) {
	if node.IsArray() {
		v, err := node.GetArrayValueAt(index)
		if err != nil {
			return nil, err
		}
		gv := node.ContinueWithValue(v, fmt.Sprintf("[%d]", index))
		return gv, nil
	} else {
		return nil, fmt.Errorf("this node identified as \"%s\" is not an array. its %s", node.IdentifiedAs(), node.thisValue.Type().String())
	}
}
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
func (node *GoValueNode) Length() (int, error) {
	if node.IsArray() || node.IsMap() || node.IsString() {
		return node.thisValue.Len(), nil
	}
	return 0, fmt.Errorf("this node identified as \"%s\" is not referencing an array, slice, map or string", node.IdentifiedAs())
}

func (node *GoValueNode) IsMap() bool {
	return node.thisValue.Kind() == reflect.Map
}
func (node *GoValueNode) GetMapValueAt(index reflect.Value) (reflect.Value, error) {
	if node.IsMap() {
		retVal := node.thisValue.MapIndex(index)
		if retVal.IsValid() {
			return retVal, nil
		} else {
			return reflect.Value{}, fmt.Errorf("this node identified as \"%s\" have no selector with specified key", node.IdentifiedAs())
		}
	}
	return reflect.Value{}, fmt.Errorf("this node identified as \"%s\" is not referencing a map", node.IdentifiedAs())
}
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
func (node *GoValueNode) GetChildNodeBySelector(index reflect.Value) (ValueNode, error) {
	val, err := node.GetMapValueAt(index)
	if err != nil {
		return nil, err
	}
	return node.ContinueWithValue(val, fmt.Sprintf("[%s->%s]", index.Type().String(), index.String())), nil
}

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
	} else {
		return fmt.Errorf("this function only used for assigning number data to number variable")
	}
}

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

func (node *GoValueNode) CallFunction(funcName string, args ...reflect.Value) (retval reflect.Value, err error) {
	switch pkg.GetBaseKind(node.thisValue) {
	case reflect.Int64, reflect.Uint64, reflect.Float64, reflect.Bool:
		return reflect.ValueOf(nil), fmt.Errorf("this node identified as \"%s\" try to call function %s which is not supported for type %s", node.IdentifiedAs(), funcName, node.thisValue.Type().String())
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
		case "Clear":
			arrFunc = ArrClear
		}
		if arrFunc != nil {
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
		case "Clear":
			mapFunc = MapClear
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

	if node.IsObject() {
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
	return reflect.ValueOf(nil), fmt.Errorf("this node identified as \"%s\" is not referencing an object thus function %s call is not supported", node.IdentifiedAs(), funcName)
}

func (node *GoValueNode) GetChildNodeByField(field string) (ValueNode, error) {
	val, err := node.GetObjectValueByField(field)
	if err != nil {
		return nil, err
	}
	return node.ContinueWithValue(val, field), nil
}

func (node *GoValueNode) IsTime() bool {
	return node.thisValue.Type().String() == "time.Time"
}
func (node *GoValueNode) IsInteger() bool {
	kind := pkg.GetBaseKind(node.thisValue)
	return kind == reflect.Int64 || kind == reflect.Uint64
}
func (node *GoValueNode) IsReal() bool {
	kind := pkg.GetBaseKind(node.thisValue)
	return kind == reflect.Float64
}
func (node *GoValueNode) IsBool() bool {
	return node.thisValue.Kind() == reflect.Bool
}
func (node *GoValueNode) IsString() bool {
	return node.thisValue.Kind() == reflect.String
}
