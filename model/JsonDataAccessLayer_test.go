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

package model

import (
	"github.com/stretchr/testify/assert"
	"reflect"
	"testing"
)

var (
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

func TestNewJsonValueNode(t *testing.T) {
	vn, err := NewJSONValueNode(bigJSON, "json")
	if err != nil {
		t.Log(err.Error())
		t.Fail()
	}
	assert.True(t, vn.IsObject())
	assert.True(t, vn.IsMap())

	vnName, err := vn.GetChildNodeByField("fullname")
	if err != nil {
		t.Log(err.Error())
		t.Fail()
	}
	assert.True(t, vnName.IsString())
	assert.Equal(t, "Bruce Wayne", vnName.Value().String())

	vnName, err = vn.GetChildNodeBySelector(reflect.ValueOf("fullname"))
	if err != nil {
		t.Log(err.Error())
		t.Fail()
	}
	assert.True(t, vnName.IsString())
	assert.Equal(t, "Bruce Wayne", vnName.Value().String())

	// todo definitely testing this in the future.
}

func TestIdentifiedAs(t *testing.T) {
	vn, err := NewJSONValueNode("\"Its A String\"", "json")
	if err != nil {
		t.Log(err.Error())
		t.Fail()
	}
	assert.Equal(t, "json", vn.IdentifiedAs())
}

func TestIsString(t *testing.T) {
	vn, err := NewJSONValueNode("\"Its A String\"", "json")
	if err != nil {
		t.Log(err.Error())
		t.Fail()
	}
	assert.True(t, vn.IsString())
}

func TestIsArray(t *testing.T) {
	vn, err := NewJSONValueNode("[1,2,3,4]", "json")
	if err != nil {
		t.Log(err.Error())
		t.Fail()
	}
	assert.True(t, vn.IsArray())
}

func TestGetArrayValueAt(t *testing.T) {
	vn, err := NewJSONValueNode("[1,\"2\",3,4]", "json")
	if err != nil {
		t.Log(err.Error())
		t.Fail()
	}

	v, err := vn.GetArrayValueAt(0)
	if err != nil {
		t.Log(err.Error())
		t.Fail()
	}
	assert.Equal(t, reflect.Float64, v.Kind())
	vn2 := vn.ContinueWithValue(v, "[0]")
	assert.True(t, vn2.IsInteger())
	assert.True(t, vn2.IsReal())
	assert.False(t, vn2.IsString())

	v, err = vn.GetArrayValueAt(1)
	if err != nil {
		t.Log(err.Error())
		t.Fail()
	}
	assert.Equal(t, reflect.String, v.Kind())
	vn2 = vn.ContinueWithValue(v, "[1]")
	assert.True(t, vn2.IsString())
	assert.False(t, vn2.IsReal())
	assert.Equal(t, "2", v.String())

	vn.SetArrayValueAt(1, reflect.ValueOf("2.0"))
	v, err = vn.GetArrayValueAt(1)
	if err != nil {
		t.Log(err.Error())
		t.Fail()
	}
	assert.Equal(t, reflect.String, v.Kind())
	vn2 = vn.ContinueWithValue(v, "[1]")
	assert.True(t, vn2.IsString())
	assert.False(t, vn2.IsReal())
	assert.Equal(t, "2.0", v.String())

	ln, err := vn.Length()
	if err != nil {
		t.Log(err.Error())
		t.Fail()
	}
	assert.Equal(t, 4, ln)

	err = vn.AppendValue([]reflect.Value{reflect.ValueOf("five"), reflect.ValueOf(6)})

	ln, err = vn.Length()
	if err != nil {
		t.Log(err.Error())
		t.Fail()
	}
	assert.Equal(t, 6, ln)
}

func TestIsMap(t *testing.T) {
	vn, err := NewJSONValueNode("{\"one\": \"1\", \"two\": 2}", "json")
	if err != nil {
		t.Log(err.Error())
		t.Fail()
	}
	assert.True(t, vn.IsMap())

	eval, err := vn.GetMapValueAt(reflect.ValueOf("one"))
	if err != nil {
		t.Log(err.Error())
		t.Fail()
	}
	assert.Equal(t, reflect.String, eval.Kind())
	assert.Equal(t, "1", eval.String())

	eval, err = vn.GetMapValueAt(reflect.ValueOf("two"))
	if err != nil {
		t.Log(err.Error())
		t.Fail()
	}
	assert.Equal(t, reflect.Float64, eval.Kind())
	assert.Equal(t, 2.0, eval.Float())

	err = vn.SetMapValueAt(reflect.ValueOf("three"), reflect.ValueOf("3"))
	if err != nil {
		t.Log(err.Error())
		t.Fail()
	}

	eval, err = vn.GetMapValueAt(reflect.ValueOf("three"))
	if err != nil {
		t.Log(err.Error())
		t.Fail()
	}
	assert.Equal(t, reflect.String, eval.Kind())
	assert.Equal(t, "3", eval.String())

	err = vn.SetMapValueAt(reflect.ValueOf("three"), reflect.ValueOf("3.0"))
	if err != nil {
		t.Log(err.Error())
		t.Fail()
	}

	eval, err = vn.GetMapValueAt(reflect.ValueOf("three"))
	if err != nil {
		t.Log(err.Error())
		t.Fail()
	}
	assert.Equal(t, reflect.String, eval.Kind())
	assert.Equal(t, "3.0", eval.String())

	ln, err := vn.Length()
	if err != nil {
		t.Log(err.Error())
		t.Fail()
	}
	assert.Equal(t, 3, ln)
}

func TestIsObject(t *testing.T) {
	vn, err := NewJSONValueNode("{\"one\": \"1\", \"two\": 2}", "json")
	if err != nil {
		t.Log(err.Error())
		t.Fail()
	}
	assert.True(t, vn.IsObject())

	eval, err := vn.GetObjectValueByField("one")
	if err != nil {
		t.Log(err.Error())
		t.Fail()
	}
	assert.Equal(t, reflect.String, eval.Kind())
	assert.Equal(t, "1", eval.String())

	eval, err = vn.GetObjectValueByField("two")
	if err != nil {
		t.Log(err.Error())
		t.Fail()
	}
	assert.Equal(t, reflect.Float64, eval.Kind())
	assert.Equal(t, 2.0, eval.Float())

	err = vn.SetObjectValueByField("three", reflect.ValueOf("3"))
	if err != nil {
		t.Log(err.Error())
		t.Fail()
	}

	eval, err = vn.GetObjectValueByField("three")
	if err != nil {
		t.Log(err.Error())
		t.Fail()
	}
	assert.Equal(t, reflect.String, eval.Kind())
	assert.Equal(t, "3", eval.String())

	err = vn.SetObjectValueByField("three", reflect.ValueOf("3.0"))
	if err != nil {
		t.Log(err.Error())
		t.Fail()
	}

	eval, err = vn.GetObjectValueByField("three")
	if err != nil {
		t.Log(err.Error())
		t.Fail()
	}
	assert.Equal(t, reflect.String, eval.Kind())
	assert.Equal(t, "3.0", eval.String())

	ln, err := vn.Length()
	if err != nil {
		t.Log(err.Error())
		t.Fail()
	}
	assert.Equal(t, 3, ln)
}
