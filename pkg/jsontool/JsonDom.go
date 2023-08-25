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
	var instance interface{}

	err := json.Unmarshal(jsonData, &instance)
	if err != nil {
		return nil, err
	}

	return &JSONData{jsonRoot: instance}, nil
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

		return -1
	}
	arr := n.interf.([]interface{})

	return len(arr)
}

// GetNodeAt will get the child not on specific index. Will panic if this not is not an array
func (n *JSONNode) GetNodeAt(index int) (*JSONNode, error) {
	if !n.IsArray() {

		return nil, fmt.Errorf("not array")
	}
	arr := n.interf.([]interface{})

	return &JSONNode{interf: arr[index]}, nil
}

// HaveKey will check if the map contains specified key. Will panic if this node is not a map
func (n *JSONNode) HaveKey(key string) (bool, error) {
	if !n.IsMap() {

		return false, fmt.Errorf("not map")
	}
	amap := n.interf.(map[string]interface{})
	if _, ok := amap[key]; ok {

		return ok, nil
	}

	return false, nil
}

// Get will fetch the child not designated with specified key. Will panic if this node is not a map
func (n *JSONNode) Get(key string) (*JSONNode, error) {
	if !n.IsMap() {

		return nil, fmt.Errorf("not map")
	}
	amap := n.interf.(map[string]interface{})

	return &JSONNode{interf: amap[key]}, nil
}

// Set will set the value of a map designated with specified key. Will panic if this node is not a map
func (n *JSONNode) Set(key string, node *JSONNode) error {
	if !n.IsMap() {

		return fmt.Errorf("not map")
	}
	amap := n.interf.(map[string]interface{})
	amap[key] = node.interf

	return nil
}

// GetString will get the string value of this node. Will panic if this node is not a string
func (n *JSONNode) GetString() (string, error) {
	if !n.IsString() {

		return "", fmt.Errorf("not string")
	}

	return n.interf.(string), nil
}

// SetString will set this node value with a string value. Will panic if this node is not a string
func (n *JSONNode) SetString(val string) error {
	if !n.IsString() {

		return fmt.Errorf("not string")
	}

	n.interf = val
	return nil
}

// GetBool will get the bool value of this node. Will panic if this node is not a boolean
func (n *JSONNode) GetBool() (bool, error) {
	if !n.IsBool() {

		return false, fmt.Errorf("not boolean")
	}

	return n.interf.(bool), nil
}

// SetBool will set this node value with boolean value, will panic if this node is not a bool
func (n *JSONNode) SetBool(val bool) error {
	if !n.IsBool() {

		return fmt.Errorf("not boolean")
	}
	n.interf = val

	return nil
}

// GetFloat will get the float value of this node. Will panic if this node is not a float.
func (n *JSONNode) GetFloat() (float64, error) {
	if !n.IsFloat() {

		return 0, fmt.Errorf("not float")
	}

	return n.interf.(float64), nil
}

// SetFloat will set this node value with float value. Will panic if this node is not a float
func (n *JSONNode) SetFloat(val float64) error {
	if !n.IsFloat() {

		return fmt.Errorf("not float")
	}
	n.interf = val

	return nil
}

// GetInt will get the int value of this node. Will panic if this node is not an int
func (n *JSONNode) GetInt() (int, error) {
	if !n.IsInt() {

		return 0, fmt.Errorf("not int")
	}
	fl := n.interf.(float64)

	return int(fl), nil
}

// SetInt will set this node value with int value. Will panic if this node is not an int
func (n *JSONNode) SetInt(val int) error {
	if !n.IsInt() {

		return fmt.Errorf("not int")
	}
	n.interf = float64(val)

	return nil
}

// JSONData represent a whole Json construct.
type JSONData struct {
	jsonRoot interface{}
}

// GetRootNode will return the root node of this JSONData
func (jo *JSONData) GetRootNode() (*JSONNode, error) {
	if jo.jsonRoot == nil {

		return nil, fmt.Errorf("root node is nil")
	}

	return &JSONNode{interf: jo.jsonRoot}, nil
}

// IsValidPath will check if the provided path is valid
func (jo *JSONData) IsValidPath(path string) (bool, error) {
	if len(path) == 0 {

		return true, nil
	}
	pathArr := strings.Split(path, ".")
	node, err := jo.GetRootNode()
	if err != nil {

		return false, err
	}

	boolres, err := jo.validPathCheck(pathArr, node)
	if err != nil {

		return false, err
	}

	return boolres, nil
}

// validPathCheck is recursion function to traverse the json tree for checking valid path
func (jo *JSONData) validPathCheck(pathArr []string, node *JSONNode) (bool, error) {
	if len(pathArr) == 0 && (node.IsString() || node.IsInt() || node.IsFloat() || node.IsBool()) {

		return true, nil
	}
	path := pathArr[0]
	if len(path) == 0 {

		return false, nil
	}
	if path[:1] == "[" && path[len(path)-1:] == "]" {
		if node.IsArray() {
			pn := path[1 : len(path)-1]
			if len(pn) == 0 {

				return false, nil
			}
			theInt, err := strconv.Atoi(pn)
			if err != nil {

				return false, nil
			}
			if theInt < 0 || theInt >= node.Len() {

				return false, nil
			}
			nNode, err := node.GetNodeAt(theInt)
			if err != nil {
				return false, err
			}
			nPathArr := pathArr[1:]

			return jo.validPathCheck(nPathArr, nNode)
		}

		return false, nil
	}
	if node.IsMap() {
		if strings.Contains(path, "[") {
			k := path[:strings.Index(path, "[")]
			haveKey, err := node.HaveKey(k)
			if err != nil {

				return false, err
			}
			if !haveKey {

				return false, nil
			}
			nNode, err := node.Get(k)
			if err != nil {

				return false, err
			}
			nPathArr := []string{path[strings.Index(path, "["):]}
			nPathArr = append(nPathArr, pathArr[1:]...)

			boolret, err := jo.validPathCheck(nPathArr, nNode)
			if err != nil {

				return false, err
			}

			return boolret, nil
		}
		hKey, err := node.HaveKey(path)
		if err != nil {

			return false, err
		}
		if hKey {
			nNode, err := node.Get(path)
			if err != nil {
				return false, err
			}
			nPathArr := pathArr[1:]

			return jo.validPathCheck(nPathArr, nNode)
		}

		return false, nil
	}

	return false, nil
}

// Get will retrieve the json node indicated by a path
func (jo *JSONData) Get(path string) (*JSONNode, error) {
	if len(path) == 0 {

		return jo.GetRootNode()
	}
	pathArr := strings.Split(path, ".")

	rNode, err := jo.GetRootNode()
	if err != nil {
		return nil, err
	}

	return jo.getByPath(pathArr, rNode)
}

// getByPath is recursion function to traverse the json tree for retrieving node at specified path
func (jo *JSONData) getByPath(pathArr []string, node *JSONNode) (*JSONNode, error) {
	if len(pathArr) == 0 && (node.IsString() || node.IsInt() || node.IsFloat() || node.IsBool()) {

		return node, nil
	}
	path := pathArr[0]
	if len(path) == 0 {

		return nil, fmt.Errorf("%s not a valid path", strings.Join(pathArr, "."))
	}
	if path[:1] == "[" && path[len(path)-1:] == "]" {
		if node.IsArray() {
			pn := path[1 : len(path)-1]
			if len(pn) == 0 {

				return nil, fmt.Errorf("not a valid path - array do not contain offset number")
			}
			theInt, err := strconv.Atoi(pn)
			if err != nil {

				return nil, fmt.Errorf("not a valid path - array offset not number")
			}
			if theInt < 0 || theInt >= node.Len() {

				return nil, fmt.Errorf("not a valid path - array offset < 0 or >= length")
			}
			nNode, err := node.GetNodeAt(theInt)
			if err != nil {

				return nil, err
			}
			nPathArr := pathArr[1:]

			return jo.getByPath(nPathArr, nNode)
		}

		return nil, fmt.Errorf("not a valid path - not an array")
	}
	if node.IsMap() {
		if strings.Contains(path, "[") {
			k := path[:strings.Index(path, "[")]
			haveKe, err := node.HaveKey(k)
			if err != nil {

				return nil, err
			}
			if !haveKe {

				return nil, fmt.Errorf("not a valid path - key not exist")
			}
			nNode, err := node.Get(k)
			if err != nil {

				return nil, err
			}
			nPathArr := []string{path[strings.Index(path, "["):]}
			nPathArr = append(nPathArr, pathArr[1:]...)

			return jo.getByPath(nPathArr, nNode)
		}
		haveKy, err := node.HaveKey(path)
		if err != nil {
			return nil, err
		}
		if haveKy {
			nNode, err := node.Get(path)
			if err != nil {

				return nil, err
			}
			nPathArr := pathArr[1:]

			return jo.getByPath(nPathArr, nNode)
		}

		return nil, fmt.Errorf("not a valid path - key not exist")
	}

	return nil, fmt.Errorf("not a valid path")
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
	node, err := jo.Get(path)
	if err != nil {
		return "", err
	}
	return node.GetString()
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
	node, err := jo.Get(path)
	if err != nil {
		return false, err
	}
	return node.GetBool()
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
	node, err := jo.Get(path)
	if err != nil {
		return 0, err
	}
	return node.GetFloat()
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
	node, err := jo.Get(path)
	if err != nil {

		return 0, err
	}

	return node.GetInt()
}

// SetInt will set the node at specified path with provided int value
func (jo *JSONData) SetInt(path string, value int) error {
	// Todo Implement this

	return fmt.Errorf("not yet implemented")
}

// IsArray will check if the node indicated by specified path is an Array node
func (jo *JSONData) IsArray(path string) (bool, error) {
	vp, err := jo.IsValidPath(path)
	if err != nil {

		return false, err
	}
	if !vp {

		return false, fmt.Errorf("%s is not a valid path", path)
	}

	jsonN, err := jo.Get(path)
	if err != nil {

		return false, err
	}
	boolRet := jsonN.IsArray()

	return boolRet, nil
}

// IsMap will check if the node indicated by specified path is a map node
func (jo *JSONData) IsMap(path string) (bool, error) {
	isValPath, err := jo.IsValidPath(path)
	if err != nil {

		return false, err
	}
	if !isValPath {

		return false, fmt.Errorf("%s is not a valid path", path)
	}
	node, err := jo.Get(path)
	if err != nil {

		return false, err
	}

	return node.IsMap(), nil
}

// IsString will check if the node indicated by specified path is a string node
func (jo *JSONData) IsString(path string) (bool, error) {
	isValPath, err := jo.IsValidPath(path)
	if err != nil {

		return false, err
	}
	if !isValPath {

		return false, fmt.Errorf("%s is not a valid path", path)
	}

	node, err := jo.Get(path)
	if err != nil {

		return false, err
	}

	return node.IsString(), nil
}

// IsBool will check if the node indicated by specified path is a bool node
func (jo *JSONData) IsBool(path string) (bool, error) {
	isValPath, err := jo.IsValidPath(path)
	if err != nil {

		return false, err
	}
	if !isValPath {

		return false, fmt.Errorf("%s is not a valid path", path)
	}

	node, err := jo.Get(path)
	if err != nil {

		return false, err
	}

	return node.IsBool(), nil
}

// IsFloat will check if the node indicated by specified path is a float node
func (jo *JSONData) IsFloat(path string) (bool, error) {
	isValPath, err := jo.IsValidPath(path)
	if err != nil {

		return false, err
	}
	if !isValPath {

		return false, fmt.Errorf("%s is not a valid path", path)
	}

	node, err := jo.Get(path)
	if err != nil {

		return false, err
	}

	return node.IsFloat(), nil
}

// IsInt will check if the node indicated by specified path is an int node
func (jo *JSONData) IsInt(path string) (bool, error) {
	isValPath, err := jo.IsValidPath(path)
	if err != nil {

		return false, err
	}
	if !isValPath {

		return false, fmt.Errorf("%s is not a valid path", path)
	}

	node, err := jo.Get(path)
	if err != nil {

		return false, err
	}

	return node.IsInt(), nil
}
