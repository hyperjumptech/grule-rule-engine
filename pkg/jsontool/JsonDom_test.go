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
			node := jobj.GetRootNode()
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

	if jdata.GetRootNode().Get("fullname").GetString() != "Bruce Wayne" {
		t.Logf("fail validate full name")
		t.Fail()
	}
	s, err := jdata.GetString("fullname")
	if err != nil {
		t.Errorf(err.Error())
	}
	if s != "Bruce Wayne" {
		t.Logf("fail validate full name")
		t.Fail()
	}

	if jdata.GetRootNode().Get("age").GetInt() != 35 {
		t.Logf("fail validate age")
		t.Fail()
	}
	i, err := jdata.GetInt("age")
	if err != nil {
		t.Errorf(err.Error())
	}
	if i != 35 {
		t.Logf("fail validate age")
		t.Fail()
	}

	if jdata.GetRootNode().Get("address").Get("street1").GetString() != "Super Mansion" {
		t.Logf("fail validate address.street1")
		t.Fail()
	}

	m, err := jdata.IsMap("address")
	if err != nil {
		t.Errorf(err.Error())
	}
	if !m {
		t.Logf("address is a map")
		t.Fail()
	}

	s, err = jdata.GetString("address.street1")
	if err != nil {
		t.Errorf(err.Error())
	}
	if s != "Super Mansion" {
		t.Logf("fail validate address.street1")
		t.Fail()
	}

	if !jdata.GetRootNode().Get("friends").IsArray() {
		t.Logf("fail validate friends as array")
		t.Fail()
	}
	b, err := jdata.IsArray("friends")
	if err != nil {
		t.Errorf(err.Error())
	}
	if !b {
		t.Logf("fail validate friends as array")
		t.Fail()
	}

	if !jdata.GetRootNode().Get("friends").GetNodeAt(1).Get("fullname").IsString() {
		t.Logf("fail validate friends[1].fullname type")
		t.Fail()
	}
	b, err = jdata.IsString("friends[1].fullname")
	if err != nil {
		t.Errorf(err.Error())
	}
	if !b {
		t.Logf("fail validate friends[1].fullname type")
		t.Fail()
	}

	if jdata.GetRootNode().Get("friends").GetNodeAt(1).Get("fullname").GetString() != "Lara Croft" {
		t.Logf("fail validate friends[1].fullname value")
		t.Fail()
	}
	str, err := jdata.GetString("friends[1].fullname")
	if err != nil {
		t.Errorf(err.Error())
	}
	if str != "Lara Croft" {
		t.Logf("fail validate friends[1].fullname value")
		t.Fail()
	}
}

func TestJsonData_IsValidPath(t *testing.T) {
	jdata, err := NewJSONData([]byte(bigJSON))
	if err != nil {
		t.Logf("Got error %s", err.Error())
		t.FailNow()
	}
	if !jdata.IsValidPath("fullname") {
		t.Logf("fullname is a valid path")
		t.Fail()
	}
	if jdata.IsValidPath("fullname.") {
		t.Logf("fullname. is not a valid path")
		t.Fail()
	}
	if jdata.IsValidPath("fullname.abc") {
		t.Logf("fullname.abc is not a valid path")
		t.Fail()
	}
	if jdata.IsValidPath("fullname[]") {
		t.Logf("fullname[] is not a valid path")
		t.Fail()
	}
	if jdata.IsValidPath("abc") {
		t.Logf("abs is not a valid path")
		t.Fail()
	}
	if !jdata.IsValidPath("") {
		t.Logf("empty string is a valid path")
		t.Fail()
	}
	if !jdata.IsValidPath("address.street1") {
		t.Logf("\"address.street1\" is a valid path")
		t.Fail()
	}
	if jdata.IsValidPath("address.street5") {
		t.Logf("\"address.street5\" is NOT a valid path")
		t.Fail()
	}
	if !jdata.IsValidPath("friends") {
		t.Logf("\"friends\" is a valid path")
		t.Fail()
	}
	if !jdata.IsValidPath("friends[1]") {
		t.Logf("\"friends[1]\" is a valid path")
		t.Fail()
	}
	if jdata.IsValidPath("friends[]") {
		t.Logf("\"friends[]\" is NOT a valid path")
		t.Fail()
	}
	if jdata.IsValidPath("friends[10]") {
		t.Logf("\"friends[10]\" is NOT a valid path")
		t.Fail()
	}
	if !jdata.IsValidPath("friends[1].address.street1") {
		t.Logf("\"friends[1].address.street1\" is a valid path")
		t.Fail()
	}
	if jdata.IsValidPath("friends[1].abc.street1") {
		t.Logf("\"friends[1].abc.street1\" is a valid path")
		t.Fail()
	}
}
