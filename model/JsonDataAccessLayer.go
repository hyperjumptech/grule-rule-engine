package model

import (
	"encoding/json"
	"fmt"
	"github.com/hyperjumptech/grule-rule-engine/pkg"
	"reflect"
	"time"
)

var (
	DateTimeLayout = time.RFC3339
)

func NewJsonValueNode(JSONString, identifiedAs string) (ValueNode, error) {
	var d interface{}
	err := json.Unmarshal([]byte(JSONString), &d)
	if err != nil {
		return nil, err
	}
	return &JsonValueNode{
		parent:       nil,
		identifiedAs: identifiedAs,
		data:         reflect.ValueOf(d),
	}, nil
}

type JsonValueNode struct {
	parent       ValueNode
	identifiedAs string
	data         reflect.Value
}

func (vn *JsonValueNode) IdentifiedAs() string {
	return vn.identifiedAs
}

func (vn *JsonValueNode) Value() reflect.Value {
	return vn.data
}

func (vn *JsonValueNode) HasParent() bool {
	return vn.parent != nil
}

func (vn *JsonValueNode) Parent() ValueNode {
	return vn.parent
}

func (vn *JsonValueNode) ContinueWithValue(value reflect.Value, identifiedAs string) ValueNode {
	return &JsonValueNode{
		parent:       vn,
		identifiedAs: identifiedAs,
		data:         value,
	}
}

func (vn *JsonValueNode) GetValue() (reflect.Value, error) {
	return vn.data, nil
}

func (vn *JsonValueNode) GetType() (reflect.Type, error) {
	return vn.data.Type(), nil
}

func (vn *JsonValueNode) IsArray() bool {
	return vn.data.Kind() == reflect.Array
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
	val := reflect.ValueOf(vn.data)
	fmt.Println(reflect.TypeOf(vn.data).String())
	return val.Kind() == reflect.Map
}

func (vn *JsonValueNode) GetMapValueAt(index reflect.Value) (reflect.Value, error) {
	if !vn.IsMap() {
		return reflect.ValueOf(nil), fmt.Errorf("not a map")
	}
	if index.Kind() != reflect.String {
		return reflect.ValueOf(nil), fmt.Errorf("JSON map selector must be a string")
	}
	tmap := vn.data.(map[string]interface{})
	if itv, ok := tmap[index.String()]; ok {
		return reflect.ValueOf(itv), nil
	}
	return reflect.ValueOf(nil), nil
}

func (vn *JsonValueNode) SetMapValueAt(index, newValue reflect.Value) error {
	if !vn.IsMap() {
		return fmt.Errorf("not a map")
	}
	if index.Kind() != reflect.String {
		return fmt.Errorf("JSON map selector must be a string")
	}
	tmap := vn.data.(map[string]interface{})
	tmap[index.String()] = pkg.ValueToInterface(newValue)
	return nil
}

func (vn *JsonValueNode) GetChildNodeBySelector(index reflect.Value) (ValueNode, error) {
	return nil, nil
}

func (vn *JsonValueNode) IsObject() bool {
	val := reflect.ValueOf(vn.data)
	return val.Kind() == reflect.Map
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
	if !vn.IsObject() {
		return nil, fmt.Errorf("not an object")
	}
	vn.data.
	tmap := vn.data.(map[string]interface{})
	if itv, ok := tmap[field]; ok {
		return vn.ContinueWithValue(reflect.ValueOf(itv), field), nil
	}
	return nil, nil
}

func (vn *JsonValueNode) IsTime() bool {
	if vn.data.Kind() == reflect.String {
		return IsDateFormatValid(DateTimeLayout, vn.data.String())
	}
	return false
}

func (vn *JsonValueNode) IsInteger() bool {
	if vn.data.Kind() == reflect.Float64 {
		f := vn.data.Float()
		i := int64(f)
		f2 := float64(i)
		return f == f2
	}
	return false
}

func (vn *JsonValueNode) IsReal() bool {
	return vn.data.Kind() == reflect.Float64
}

func (vn *JsonValueNode) IsBool() bool {
	return vn.data.Kind() == reflect.Bool
}

func (vn *JsonValueNode) IsString() bool {
	return vn.data.Kind() == reflect.String
}
