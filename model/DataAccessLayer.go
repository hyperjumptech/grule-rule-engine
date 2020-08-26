package model

import "reflect"

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
