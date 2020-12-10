package model

import (
	"reflect"
	"time"
)

var (
	DateTimeLayout = time.RFC3339
)

type JsonValueNode struct {
	parent       ValueNode
	identifiedAs string
	data         interface{}
}

func (vn *JsonValueNode) IdentifiedAs() string {
	return vn.identifiedAs
}

func (vn *JsonValueNode) Value() reflect.Value {
	return reflect.ValueOf(vn.data)
}

func (vn *JsonValueNode) HasParent() bool {
	return vn.parent != nil
}

func (vn *JsonValueNode) Parent() ValueNode {
	return vn.parent
}

func (vn *JsonValueNode) ContinueWithValue(value reflect.Value, identifiedAs string) ValueNode {
	return nil
}

func (vn *JsonValueNode) GetValue() (reflect.Value, error) {
	return reflect.ValueOf(nil), nil
}

func (vn *JsonValueNode) GetType() (reflect.Type, error) {
	return reflect.TypeOf(nil), nil
}

func (vn *JsonValueNode) IsArray() bool {
	return false
}

func (vn *JsonValueNode) GetArrayType() (reflect.Type, error) {
	return reflect.TypeOf(nil), nil
}

func (vn *JsonValueNode) GetArrayValueAt(index int) (reflect.Value, error) {
	return reflect.ValueOf(nil), nil
}

func (vn *JsonValueNode) GetChildNodeByIndex(index int) (ValueNode, error) {
	return nil, nil
}

func (vn *JsonValueNode) SetArrayValueAt(index int, value reflect.Value) error {
	return nil
}

func (vn *JsonValueNode) AppendValue(value []reflect.Value) error {
	return nil
}

func (vn *JsonValueNode) Length() (int, error) {
	return 0, nil
}

func (vn *JsonValueNode) IsMap() bool {
	return false
}

func (vn *JsonValueNode) GetMapValueAt(index reflect.Value) (reflect.Value, error) {
	return reflect.ValueOf(nil), nil
}

func (vn *JsonValueNode) SetMapValueAt(index, newValue reflect.Value) error {
	return nil
}

func (vn *JsonValueNode) GetChildNodeBySelector(index reflect.Value) (ValueNode, error) {
	return nil, nil
}

func (vn *JsonValueNode) IsObject() bool {
	return false
}

func (vn *JsonValueNode) GetObjectValueByField(field string) (reflect.Value, error) {
	return reflect.ValueOf(nil), nil
}

func (vn *JsonValueNode) GetObjectTypeByField(field string) (reflect.Type, error) {
	return reflect.TypeOf(nil), nil
}

func (vn *JsonValueNode) SetObjectValueByField(field string, newValue reflect.Value) error {
	return nil
}

func (vn *JsonValueNode) CallFunction(funcName string, args ...reflect.Value) (reflect.Value, error) {
	return reflect.ValueOf(nil), nil
}

func (vn *JsonValueNode) GetChildNodeByField(field string) (ValueNode, error) {
	return nil, nil
}

func (vn *JsonValueNode) IsTime() bool {
	if _, ok := vn.data.(string); ok {
		return IsDateFormatValid(DateTimeLayout, vn.data.(string))
	}
	return false
}

func (vn *JsonValueNode) IsInteger() bool {
	if _, ok := vn.data.(string); ok {
		v := reflect.ValueOf(vn.data)
		f := v.Float()
		i := int64(f)
		f2 := float64(i)
		return f == f2
	}
	return false
}

func (vn *JsonValueNode) IsReal() bool {
	if _, ok := vn.data.(float64); ok {
		return true
	}
	return false
}

func (vn *JsonValueNode) IsBool() bool {
	if _, ok := vn.data.(bool); ok {
		return true
	}
	return false
}

func (vn *JsonValueNode) IsString() bool {
	if _, ok := vn.data.(string); ok {
		return true
	}
	return false
}
