package jsontool

import (
	"encoding/json"
	"fmt"
	"reflect"
	"strconv"
	"strings"
)

func NewJsonData(jsonData []byte) (*JsonData, error) {
	var tm interface{}

	err := json.Unmarshal(jsonData, &tm)
	if err != nil {
		return nil, err
	}
	return &JsonData{jsonRoot: tm}, nil
}

type JsonNode struct {
	interf interface{}
}

func (n *JsonNode) IsArray() bool {
	return reflect.TypeOf(n.interf).String() == "[]interface {}"
}

func (n *JsonNode) IsMap() bool {
	return reflect.TypeOf(n.interf).String() == "map[string]interface {}"
}

func (n *JsonNode) IsString() bool {
	return reflect.TypeOf(n.interf).String() == "string"
}

func (n *JsonNode) IsBool() bool {
	return reflect.TypeOf(n.interf).String() == "bool"
}

func (n *JsonNode) IsFloat() bool {
	return reflect.TypeOf(n.interf).String() == "float64"
}

func (n *JsonNode) IsInt() bool {
	if reflect.TypeOf(n.interf).String() == "float64" {
		v := reflect.ValueOf(n.interf)
		f := v.Float()
		i := int64(f)
		f2 := float64(i)
		return f == f2
	}
	return true
}

func (n *JsonNode) Len() int {
	if !n.IsArray() {
		panic("Not array")
	}
	arr := n.interf.([]interface{})
	return len(arr)
}

func (n *JsonNode) NodeAt(index int) *JsonNode {
	if !n.IsArray() {
		panic("Not array")
	}
	arr := n.interf.([]interface{})
	return &JsonNode{interf: arr[index]}
}

func (n *JsonNode) HaveKey(key string) bool {
	if !n.IsMap() {
		panic("Not map")
	}
	amap := n.interf.(map[string]interface{})
	if _, ok := amap[key]; ok {
		return ok
	}
	return false
}

func (n *JsonNode) Get(key string) *JsonNode {
	if !n.IsMap() {
		panic("Not map")
	}
	amap := n.interf.(map[string]interface{})
	return &JsonNode{interf: amap[key]}
}

func (n *JsonNode) GetString() string {
	if !n.IsString() {
		panic("Not string")
	}
	return n.interf.(string)
}

func (n *JsonNode) GetBool() bool {
	if !n.IsBool() {
		panic("Not boolean")
	}
	return n.interf.(bool)
}

func (n *JsonNode) GetFloat() float64 {
	if !n.IsFloat() {
		panic("Not float")
	}
	return n.interf.(float64)
}

func (n *JsonNode) GetInt() int {
	if !n.IsInt() {
		panic("Not int")
	}
	fl := n.interf.(float64)
	return int(fl)
}

type JsonData struct {
	jsonRoot interface{}
}

func (jo *JsonData) GetRootNode() *JsonNode {
	if jo.jsonRoot == nil {
		panic(fmt.Sprintf("root node is nil"))
	}
	return &JsonNode{interf: jo.jsonRoot}
}

func (jo *JsonData) IsValidPath(path string) bool {
	if len(path) == 0 {
		return true
	}
	pathArr := strings.Split(path, ".")
	node := jo.GetRootNode()
	return jo.validPathCheck(pathArr, node)
}

func (jo *JsonData) validPathCheck(pathArr []string, node *JsonNode) bool {
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
			nNode := node.NodeAt(n)
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

func (jo *JsonData) Get(path string) *JsonNode {
	if len(path) == 0 {
		return jo.GetRootNode()
	}
	pathArr := strings.Split(path, ".")
	return jo.getByPath(pathArr, jo.GetRootNode())
}

func (jo *JsonData) getByPath(pathArr []string, node *JsonNode) *JsonNode {
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
			nNode := node.NodeAt(n)
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

func (n *JsonData) GetString(path string) (string, error) {
	b, err := n.IsString(path)
	if err != nil {
		return "", err
	}
	if !b {
		return "", fmt.Errorf("%s is not a string", path)
	}
	node := n.Get(path)
	return node.GetString(), nil
}

func (n *JsonData) GetBool(path string) (bool, error) {
	b, err := n.IsBool(path)
	if err != nil {
		return false, err
	}
	if !b {
		return false, fmt.Errorf("%s is not a boolean", path)
	}
	node := n.Get(path)
	return node.GetBool(), nil
}

func (n *JsonData) GetFloat(path string) (float64, error) {
	b, err := n.IsFloat(path)
	if err != nil {
		return 0, err
	}
	if !b {
		return 0, fmt.Errorf("%s is not a float", path)
	}
	node := n.Get(path)
	return node.GetFloat(), nil
}

func (n *JsonData) GetInt(path string) (int, error) {
	b, err := n.IsInt(path)
	if err != nil {
		return 0, err
	}
	if !b {
		return 0, fmt.Errorf("%s is not an int", path)
	}
	node := n.Get(path)
	return node.GetInt(), nil
}

func (n *JsonData) IsArray(path string) (bool, error) {
	if !n.IsValidPath(path) {
		return false, fmt.Errorf("%s is not a valid path", path)
	}
	return n.Get(path).IsArray(), nil
}

func (n *JsonData) IsMap(path string) (bool, error) {
	if !n.IsValidPath(path) {
		return false, fmt.Errorf("%s is not a valid path", path)
	}
	return n.Get(path).IsMap(), nil
}

func (n *JsonData) IsString(path string) (bool, error) {
	if !n.IsValidPath(path) {
		return false, fmt.Errorf("%s is not a valid path", path)
	}
	return n.Get(path).IsString(), nil
}

func (n *JsonData) IsBool(path string) (bool, error) {
	if !n.IsValidPath(path) {
		return false, fmt.Errorf("%s is not a valid path", path)
	}
	return n.Get(path).IsBool(), nil
}

func (n *JsonData) IsFloat(path string) (bool, error) {
	if !n.IsValidPath(path) {
		return false, fmt.Errorf("%s is not a valid path", path)
	}
	return n.Get(path).IsFloat(), nil
}

func (n *JsonData) IsInt(path string) (bool, error) {
	if !n.IsValidPath(path) {
		return false, fmt.Errorf("%s is not a valid path", path)
	}
	return n.Get(path).IsInt(), nil
}
