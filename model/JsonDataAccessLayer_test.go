package model

import (
	"github.com/stretchr/testify/assert"
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
	vn, err := NewJsonValueNode(bigJSON, "json")
	if err != nil {
		t.Log(err.Error())
		t.Fail()
	}
	assert.True(t, vn.IsObject())
	assert.True(t, vn.IsMap())
}
