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
	return nil
	// TODO resolve this
}
