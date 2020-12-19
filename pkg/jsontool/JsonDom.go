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

package jsontool

import (
	"encoding/json"
	"fmt"
	"reflect"
	"strconv"
	"strings"
)

// NewJSONData will create a new instance of JSONData
func NewJSONData(jsonData []byte) (*JSONData, error) {
	var tm interface{}

	err := json.Unmarshal(jsonData, &tm)
	if err != nil {
		return nil, err
	}
	return &JSONData{jsonRoot: tm}, nil
}

// JSONNode represent a node in JSON Tree
type JSONNode struct {
	interf interface{}
}

// IsArray will check if this node represent an array
func (n *JSONNode) IsArray() bool {
	return reflect.TypeOf(n.interf).String() == "[]interface {}"
}

// IsMap will check if this node represent a Map
func (n *JSONNode) IsMap() bool {
	return reflect.TypeOf(n.interf).String() == "map[string]interface {}"
}

// IsString check if this node represent a string
func (n *JSONNode) IsString() bool {
	return reflect.TypeOf(n.interf).String() == "string"
}

// IsBool check if this node represent a boolean
func (n *JSONNode) IsBool() bool {
	return reflect.TypeOf(n.interf).String() == "bool"
}

// IsFloat check if this node represent a float
func (n *JSONNode) IsFloat() bool {
	return reflect.TypeOf(n.interf).String() == "float64"
}

// IsInt checks if this node represent an int
func (n *JSONNode) IsInt() bool {
	if reflect.TypeOf(n.interf).String() == "float64" {
		v := reflect.ValueOf(n.interf)
		f := v.Float()
		i := int64(f)
		f2 := float64(i)
		return f == f2
	}
	return true
}

// Len return length of element in this array. Will panic if this node is not an array
func (n *JSONNode) Len() int {
	if !n.IsArray() {
		panic("Not array")
	}
	arr := n.interf.([]interface{})
	return len(arr)
}

// GetNodeAt will get the child not on specific index. Will panic if this not is not an array
func (n *JSONNode) GetNodeAt(index int) *JSONNode {
	if !n.IsArray() {
		panic("Not array")
	}
	arr := n.interf.([]interface{})
	return &JSONNode{interf: arr[index]}
}

// HaveKey will check if the map contains specified key. Will panic if this node is not a map
func (n *JSONNode) HaveKey(key string) bool {
	if !n.IsMap() {
		panic("Not map")
	}
	amap := n.interf.(map[string]interface{})
	if _, ok := amap[key]; ok {
		return ok
	}
	return false
}

// Get will fetch the child not designated with specified key. Will panic if this node is not a map
func (n *JSONNode) Get(key string) *JSONNode {
	if !n.IsMap() {
		panic("Not map")
	}
	amap := n.interf.(map[string]interface{})
	return &JSONNode{interf: amap[key]}
}

// Set will set the value of a map designated with specified key. Will panic if this node is not a map
func (n *JSONNode) Set(key string, node *JSONNode) {
	if !n.IsMap() {
		panic("Not map")
	}
	amap := n.interf.(map[string]interface{})
	amap[key] = node.interf
}

// GetString will get the string value of this node. Will panic if this node is not a string
func (n *JSONNode) GetString() string {
	if !n.IsString() {
		panic("Not string")
	}
	return n.interf.(string)
}

// SetString will set this node value with a string value. Will panic if this node is not a string
func (n *JSONNode) SetString(val string) {
	if !n.IsString() {
		panic("Not string")
	}
	n.interf = val
}

// GetBool will get the bool value of this node. Will panic if this node is not a boolean
func (n *JSONNode) GetBool() bool {
	if !n.IsBool() {
		panic("Not boolean")
	}
	return n.interf.(bool)
}

// SetBool will set this node value with boolean value, will panic if this node is not a bool
func (n *JSONNode) SetBool(val bool) {
	if !n.IsBool() {
		panic("Not boolean")
	}
	n.interf = val
}

// GetFloat will get the float value of this node. Will panic if this node is not a float.
func (n *JSONNode) GetFloat() float64 {
	if !n.IsFloat() {
		panic("Not float")
	}
	return n.interf.(float64)
}

// SetFloat will set this node value with float value. Will panic if this node is not a float
func (n *JSONNode) SetFloat(val float64) {
	if !n.IsFloat() {
		panic("Not float")
	}
	n.interf = val
}

// GetInt will get the int value of this node. Will panic if this node is not an int
func (n *JSONNode) GetInt() int {
	if !n.IsInt() {
		panic("Not int")
	}
	fl := n.interf.(float64)
	return int(fl)
}

// SetInt will set this node value with int value. Will panic if this node is not an int
func (n *JSONNode) SetInt(val int) {
	if !n.IsInt() {
		panic("Not int")
	}
	n.interf = float64(val)
}

// JSONData represent a whole Json construct.
type JSONData struct {
	jsonRoot interface{}
}

// GetRootNode will return the root node of this JSONData
func (jo *JSONData) GetRootNode() *JSONNode {
	if jo.jsonRoot == nil {
		panic(fmt.Sprintf("root node is nil"))
	}
	return &JSONNode{interf: jo.jsonRoot}
}

// IsValidPath will check if the provided path is valid
func (jo *JSONData) IsValidPath(path string) bool {
	if len(path) == 0 {
		return true
	}
	pathArr := strings.Split(path, ".")
	node := jo.GetRootNode()
	return jo.validPathCheck(pathArr, node)
}

// validPathCheck is recursion function to traverse the json tree for checking valid path
func (jo *JSONData) validPathCheck(pathArr []string, node *JSONNode) bool {
	if len(pathArr) == 0 && (node.IsString() || node.IsInt() || node.IsFloat() || node.IsBool()) {
		return true
	}
	p := pathArr[0]
	if len(p) == 0 {
		return false
	}
	if p[:1] == "[" && p[len(p)-1:] == "]" {
		if node.IsArray() {
			pn := p[1 : len(p)-1]
			if len(pn) == 0 {
				return false
			}
			n, err := strconv.Atoi(pn)
			if err != nil {
				return false
			}
			if n < 0 || n >= node.Len() {
				return false
			}
			nNode := node.GetNodeAt(n)
			nPathArr := pathArr[1:]
			return jo.validPathCheck(nPathArr, nNode)
		}
		return false
	}
	if node.IsMap() {
		if strings.Contains(p, "[") {
			k := p[:strings.Index(p, "[")]
			if !node.HaveKey(k) {
				return false
			}
			nNode := node.Get(k)
			nPathArr := []string{p[strings.Index(p, "["):]}
			nPathArr = append(nPathArr, pathArr[1:]...)
			return jo.validPathCheck(nPathArr, nNode)
		}
		if node.HaveKey(p) {
			nNode := node.Get(p)
			nPathArr := pathArr[1:]
			return jo.validPathCheck(nPathArr, nNode)
		}
		return false
	}
	return false
}

// Get will retrieve the json node indicated by a path
func (jo *JSONData) Get(path string) *JSONNode {
	if len(path) == 0 {
		return jo.GetRootNode()
	}
	pathArr := strings.Split(path, ".")
	return jo.getByPath(pathArr, jo.GetRootNode())
}

// getByPath is recursion function to traverse the json tree for retrieving node at specified path
func (jo *JSONData) getByPath(pathArr []string, node *JSONNode) *JSONNode {
	if len(pathArr) == 0 && (node.IsString() || node.IsInt() || node.IsFloat() || node.IsBool()) {
		return node
	}
	p := pathArr[0]
	if len(p) == 0 {
		panic("Not a valid path")
	}
	if p[:1] == "[" && p[len(p)-1:] == "]" {
		if node.IsArray() {
			pn := p[1 : len(p)-1]
			if len(pn) == 0 {
				panic("Not a valid path - array do not contain offset number")
			}
			n, err := strconv.Atoi(pn)
			if err != nil {
				panic("Not a valid path - array offset not number")
			}
			if n < 0 || n >= node.Len() {
				panic("Not a valid path - array offset < 0 or >= length")
			}
			nNode := node.GetNodeAt(n)
			nPathArr := pathArr[1:]
			return jo.getByPath(nPathArr, nNode)
		}
		panic("Not a valid path - not an array")
	}
	if node.IsMap() {
		if strings.Contains(p, "[") {
			k := p[:strings.Index(p, "[")]
			if !node.HaveKey(k) {
				panic("Not a valid path - key not exist")
			}
			nNode := node.Get(k)
			nPathArr := []string{p[strings.Index(p, "["):]}
			nPathArr = append(nPathArr, pathArr[1:]...)
			return jo.getByPath(nPathArr, nNode)
		}
		if node.HaveKey(p) {
			nNode := node.Get(p)
			nPathArr := pathArr[1:]
			return jo.getByPath(nPathArr, nNode)
		}
		panic("Not a valid path - key not exist")
	}
	panic("Not a valid path")
}

// GetString will get the string value from a json indicated by specified path. Error returned if path is not valid.
func (jo *JSONData) GetString(path string) (string, error) {
	b, err := jo.IsString(path)
	if err != nil {
		return "", err
	}
	if !b {
		return "", fmt.Errorf("%s is not a string", path)
	}
	node := jo.Get(path)
	return node.GetString(), nil
}

// SetString will set the node at specified path with provided string value
func (jo *JSONData) SetString(path, value string) error {
	// Todo Implement this
	return fmt.Errorf("not yet implemented")
}

// GetBool will get the bool value from a json indicated by specified path. Error returned if path is not valid.
func (jo *JSONData) GetBool(path string) (bool, error) {
	b, err := jo.IsBool(path)
	if err != nil {
		return false, err
	}
	if !b {
		return false, fmt.Errorf("%s is not a boolean", path)
	}
	node := jo.Get(path)
	return node.GetBool(), nil
}

// SetBool will set the node at specified path with provided bool value
func (jo *JSONData) SetBool(path string, value bool) error {
	// Todo Implement this
	return fmt.Errorf("not yet implemented")
}

// GetFloat will get the float value from a json indicated by specified path. Error returned if path is not valid.
func (jo *JSONData) GetFloat(path string) (float64, error) {
	b, err := jo.IsFloat(path)
	if err != nil {
		return 0, err
	}
	if !b {
		return 0, fmt.Errorf("%s is not a float", path)
	}
	node := jo.Get(path)
	return node.GetFloat(), nil
}

// SetFloat will set the node at specified path with provided float value
func (jo *JSONData) SetFloat(path string, value float64) error {
	// Todo Implement this
	return fmt.Errorf("not yet implemented")
}

// GetInt will get the int value from a json indicated by specified path. Error returned if path is not valid.
func (jo *JSONData) GetInt(path string) (int, error) {
	b, err := jo.IsInt(path)
	if err != nil {
		return 0, err
	}
	if !b {
		return 0, fmt.Errorf("%s is not an int", path)
	}
	node := jo.Get(path)
	return node.GetInt(), nil
}

// SetInt will set the node at specified path with provided int value
func (jo *JSONData) SetInt(path string, value int) error {
	// Todo Implement this
	return fmt.Errorf("not yet implemented")
}

// IsArray will check if the node indicated by specified path is an Array node
func (jo *JSONData) IsArray(path string) (bool, error) {
	if !jo.IsValidPath(path) {
		return false, fmt.Errorf("%s is not a valid path", path)
	}
	return jo.Get(path).IsArray(), nil
}

// IsMap will check if the node indicated by specified path is a map node
func (jo *JSONData) IsMap(path string) (bool, error) {
	if !jo.IsValidPath(path) {
		return false, fmt.Errorf("%s is not a valid path", path)
	}
	return jo.Get(path).IsMap(), nil
}

// IsString will check if the node indicated by specified path is a string node
func (jo *JSONData) IsString(path string) (bool, error) {
	if !jo.IsValidPath(path) {
		return false, fmt.Errorf("%s is not a valid path", path)
	}
	return jo.Get(path).IsString(), nil
}

// IsBool will check if the node indicated by specified path is a bool node
func (jo *JSONData) IsBool(path string) (bool, error) {
	if !jo.IsValidPath(path) {
		return false, fmt.Errorf("%s is not a valid path", path)
	}
	return jo.Get(path).IsBool(), nil
}

// IsFloat will check if the node indicated by specified path is a float node
func (jo *JSONData) IsFloat(path string) (bool, error) {
	if !jo.IsValidPath(path) {
		return false, fmt.Errorf("%s is not a valid path", path)
	}
	return jo.Get(path).IsFloat(), nil
}

// IsInt will check if the node indicated by specified path is an int node
func (jo *JSONData) IsInt(path string) (bool, error) {
	if !jo.IsValidPath(path) {
		return false, fmt.Errorf("%s is not a valid path", path)
	}
	return jo.Get(path).IsInt(), nil
}
