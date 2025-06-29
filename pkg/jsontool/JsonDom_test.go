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
	"reflect"
	"testing"
)

var (
	JSONText = [...]string{
		`{
	"stringAttr":"stringVal", 
	"intAttr":123, 
	"floatAttr":123.456, 
	"boolAttr": true, 
	"arrayAttr": [1,2,3,4], 
	"objAttr": {
		"attr1":1234, 
		"attr2": "string"
	}}`,
		`[1,2,3,4]`,
		`["1","2","3","4"]`,
		`["1", 2, 3.4, true]`,
		`{}`,
		`[]`,
		`12345`,
		`123.45`,
		`"123"`,
		`true`,
		`false`,
	}

	JSONType = [...]string{
		"obj",
		"arr",
		"arr",
		"arr",
		"obj",
		"arr",
		"int",
		"float",
		"string",
		"bool",
		"bool",
	}

	bigJSON = `
{
	"fullname" : "Bruce Wayne",
	"address" : {
		"street1" : "Super Mansion",
		"street2" : "DC Commic Rd",
		"zip" : 84893,
		"city" : "Metro City"
	},
	"age" : 35,
	"height" : 5.23,
	"gender" : "male",
	"active" : true,
	"friends" : [
		{
			"fullname" : "Peter Parker",
			"address" : {
				"street1" : "Other Super Mansion",
				"street2" : "DC Commic Rd",
				"zip" : 84893,
				"city" : "Metro City"
			},
			"age" : 32,
			"height" : 5.42,
			"gender" : "male",
			"active" : true
		}, {
			"fullname" : "Lara Croft",
			"address" : {
				"street1" : "Lilly Bulevard",
				"street2" : "Clear Hill",
				"zip" : 65223,
				"city" : "Fiction Mount"
			},
			"age" : 28,
			"height" : 5.5,
			"gender" : "female",
			"active" : false
		}
	]
}
`
)

func TestNewJsonObject(t *testing.T) {
	for i, v := range JSONText {
		jobj, err := NewJSONData([]byte(v))
		if err != nil {
			t.Logf("Error : %d - %s. Got %s", i, v, err.Error())
			t.Fail()
		} else {
			typ := reflect.TypeOf(jobj.jsonRoot)
			t.Logf("Data : %d - %s. Type is : %s ", i, v, typ.String())
			node, err := jobj.GetRootNode()
			if err != nil {
				t.Fail()
			}
			switch JSONType[i] {
			case "obj":
				if !node.IsMap() {
					t.Logf("Error : %d - %s. Is not map", i, v)
					t.Fail()
				}
			case "arr":
				if !node.IsArray() {
					t.Logf("Error : %d - %s. Is not array", i, v)
					t.Fail()
				}
			case "int":
				if !node.IsInt() {
					t.Logf("Error : %d - %s. Is not int", i, v)
					t.Fail()
				}
			case "float":
				if !node.IsFloat() {
					t.Logf("Error : %d - %s. Is not float", i, v)
					t.Fail()
				}
				if node.IsInt() {
					t.Logf("Error : %d - %s. Is supposed not to be int", i, v)
					t.Fail()
				}
			case "string":
				if !node.IsString() {
					t.Logf("Error : %d - %s. Is not string", i, v)
					t.Fail()
				}
			case "bool":
				if !node.IsBool() {
					t.Logf("Error : %d - %s. Is not bool", i, v)
					t.Fail()
				}
			default:
				t.Logf("Error : %d - %s. Invalid test data type", i, v)
				t.Fail()
			}
		}
	}
}

func TestJsonNodeOperations(t *testing.T) {
	jdata, err := NewJSONData([]byte(bigJSON))
	if err != nil {
		t.Logf("Got error %s", err.Error())
		t.FailNow()
	}

	rNode, err := jdata.GetRootNode()
	if err != nil {
		t.Logf("Got error %s", err.Error())
		t.FailNow()
	}

	jnode, err := rNode.Get("fullname")
	if err != nil {
		t.Logf("Got error %s", err.Error())
		t.FailNow()
	}

	sn, err := jnode.GetString()
	if err != nil {
		t.Logf("Got error %s", err.Error())
		t.FailNow()
	}

	if sn != "Bruce Wayne" {
		t.Logf("fail validate full name")
		t.Fail()
	}
	s, err := jdata.GetString("fullname")
	if err != nil {
		t.Errorf("%v", err.Error())
	}
	if s != "Bruce Wayne" {
		t.Logf("fail validate full name")
		t.Fail()
	}

	rnode, err := jdata.GetRootNode()
	if err != nil {
		t.Fail()
	}
	ageNode, err := rnode.Get("age")
	if err != nil {
		t.Fail()
	}
	theInt, err := ageNode.GetInt()
	if err != nil {
		t.Fail()
	}

	if theInt != 35 {
		t.Logf("fail validate age")
		t.Fail()
	}
	i, err := jdata.GetInt("age")
	if err != nil {
		t.Errorf("%v", err.Error())
	}
	if i != 35 {
		t.Logf("fail validate age")
		t.Fail()
	}

	rnode, err = jdata.GetRootNode()
	if err != nil {
		t.Logf("Got error %s", err.Error())
		t.FailNow()
	}

	addrNode, err := rnode.Get("address")
	if err != nil {
		t.Logf("Got error %s", err.Error())
		t.FailNow()
	}

	strtNode, err := addrNode.Get("street1")
	if err != nil {
		t.Logf("Got error %s", err.Error())
		t.FailNow()
	}

	strValue, err := strtNode.GetString()
	if err != nil {
		t.Logf("Got error %s", err.Error())
		t.FailNow()
	}

	if strValue != "Super Mansion" {
		t.Logf("fail validate address.street1")
		t.Fail()
	}

	m, err := jdata.IsMap("address")
	if err != nil {
		t.Errorf("%v", err.Error())
	}
	if !m {
		t.Logf("address is a map")
		t.Fail()
	}

	s, err = jdata.GetString("address.street1")
	if err != nil {
		t.Errorf("%v", err.Error())
	}
	if s != "Super Mansion" {
		t.Logf("fail validate address.street1")
		t.Fail()
	}

	rNode, err = jdata.GetRootNode()
	if err != nil {
		t.FailNow()
	}
	nNode, err := rNode.Get("friends")
	if err != nil {
		t.FailNow()
	}

	if !nNode.IsArray() {
		t.Logf("fail validate friends as array")
		t.Fail()
	}
	b, err := jdata.IsArray("friends")
	if err != nil {
		t.Errorf("%v", err.Error())
	}
	if !b {
		t.Logf("fail validate friends as array")
		t.Fail()
	}

	rNode, err = jdata.GetRootNode()
	if err != nil {
		t.FailNow()
	}
	nNode, err = rNode.Get("friends")
	if err != nil {
		t.FailNow()
	}
	nNodeAt, err := nNode.GetNodeAt(1)
	if err != nil {
		t.FailNow()
	}
	nNode2, err := nNodeAt.Get("fullname")
	if err != nil {
		t.FailNow()
	}
	if !nNode2.IsString() {
		t.Logf("fail validate friends[1].fullname type")
		t.Fail()
	}
	b, err = jdata.IsString("friends[1].fullname")
	if err != nil {
		t.Errorf("%v", err.Error())
	}
	if !b {
		t.Logf("fail validate friends[1].fullname type")
		t.Fail()
	}

	rNode, err = jdata.GetRootNode()
	if err != nil {
		t.FailNow()
	}
	nNode, err = rNode.Get("friends")
	if err != nil {
		t.FailNow()
	}
	nNodeAt, err = nNode.GetNodeAt(1)
	if err != nil {
		t.FailNow()
	}
	nNode2, err = nNodeAt.Get("fullname")
	if err != nil {
		t.FailNow()
	}
	str, err := nNode2.GetString()
	if err != nil {
		t.FailNow()
	}
	if str != "Lara Croft" {
		t.Logf("fail validate friends[1].fullname value")
		t.Fail()
	}
	str, err = jdata.GetString("friends[1].fullname")
	if err != nil {
		t.Errorf("%v", err.Error())
	}
	if str != "Lara Croft" {
		t.Logf("fail validate friends[1].fullname value")
		t.Fail()
	}
}

type PathTest struct {
	Path  string
	Valid bool
}

func TestJsonData_IsValidPath(t *testing.T) {

	pTests := make([]PathTest, 0)

	pTests = append(pTests, PathTest{Path: "fullname", Valid: true})
	pTests = append(pTests, PathTest{Path: "fullname.", Valid: false})
	pTests = append(pTests, PathTest{Path: "fullname.abc", Valid: false})
	pTests = append(pTests, PathTest{Path: "abc", Valid: false})
	pTests = append(pTests, PathTest{Path: "", Valid: true})
	pTests = append(pTests, PathTest{Path: "address.street1", Valid: true})
	pTests = append(pTests, PathTest{Path: "address.street5", Valid: false})
	pTests = append(pTests, PathTest{Path: "friends", Valid: true})
	pTests = append(pTests, PathTest{Path: "friends[1]", Valid: true})
	pTests = append(pTests, PathTest{Path: "friends[]", Valid: false})
	pTests = append(pTests, PathTest{Path: "friends[10]", Valid: false})
	pTests = append(pTests, PathTest{Path: "friends[1].address.street1", Valid: true})
	pTests = append(pTests, PathTest{Path: "friends[1].abc.street1", Valid: false})

	jdata, err := NewJSONData([]byte(bigJSON))
	if err != nil {
		t.Logf("Got error %s", err.Error())
		t.FailNow()
	}

	for _, ptest := range pTests {
		if val, err := jdata.IsValidPath(ptest.Path); err != nil {
			t.Logf("got error %s", err.Error())
			t.FailNow()
		} else if val != ptest.Valid {
			t.Logf("'%s' valid path expect '%v' but '%v'", ptest.Path, ptest.Valid, val)
			t.FailNow()
		}
	}

}
