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
	return nil, fmt.Errorf("this node is not referring to an array or slice")
}
func (node *GoValueNode) GetArrayValueAt(index int) (reflect.Value, error) {
	if node.IsArray() {
		return node.thisValue.Index(index), nil
	}
	return reflect.Value{}, fmt.Errorf("this node is not referring to an array or slice")
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
		return nil, fmt.Errorf("this node is not an array. its %s", node.thisValue.Type().String())
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
			val.Set(value)
			return nil
		}
		return fmt.Errorf("can not set value on array index %d", index)
	}
	return fmt.Errorf("this node is not referencing an array or slice")
}
func (node *GoValueNode) AppendValue(value reflect.Value) (err error) {
	if node.IsArray() {
		arrVal := node.thisValue
		if arrVal.CanSet() {
			defer func() {
				if r := recover(); r != nil {
					err = fmt.Errorf("recovered : %v", r)
				}
			}()
			arrVal.Set(reflect.Append(arrVal, value))
			return nil
		}
	}
	return fmt.Errorf("this node is not referencing an array or slice")
}
func (node *GoValueNode) Length() (int, error) {
	if node.IsArray() || node.IsMap() || node.IsString() {
		return node.thisValue.Len(), nil
	}
	return 0, fmt.Errorf("this node is not referencing an array, slice, map or string")
}

func (node *GoValueNode) IsMap() bool {
	return node.thisValue.Kind() == reflect.Map
}
func (node *GoValueNode) GetMapValueAt(index reflect.Value) (reflect.Value, error) {
	panic("not yet implemented")
}
func (node *GoValueNode) SetMapValueAt(index, newValue reflect.Value) error {
	panic("not yet implemented")
}
func (node *GoValueNode) GetChildNodeBySelector(index reflect.Value) (ValueNode, error) {
	panic("not yet implemented")
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
	return reflect.Value{}, fmt.Errorf("this node is not referencing to an object")
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
	return nil, fmt.Errorf("this node is not referring to an object")
}

func (node *GoValueNode) SetObjectValueByField(field string, newValue reflect.Value) (err error) {
	fieldVal := node.thisValue.Elem().FieldByName(field)
	if fieldVal.IsValid() && fieldVal.CanAddr() && fieldVal.CanSet() {
		defer func() {
			if r := recover(); r != nil {
				err = fmt.Errorf("recovered : %v", r)
			}
		}()
		fieldVal.Set(newValue)
		return nil
	}
	return fmt.Errorf("node is not valid nor addressable")
}

func (node *GoValueNode) CallFunction(funcName string, args ...reflect.Value) (reflect.Value, error) {
	panic("not yet implemented")
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
