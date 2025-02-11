//  Copyright DataWiseHQ/grule-rule-engine Authors
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

package pkg

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

const jsonData = `{
    "name": "SpeedUp",
    "desc": "When testcar is speeding up we keep increase the speed.",
    "salience": 10,
    "when": "TestCar.SpeedUp == true && TestCar.Speed < TestCar.MaxSpeed",
    "then": [
        "TestCar.Speed = TestCar.Speed + TestCar.SpeedIncrement",
        "DistanceRecord.TotalDistance = DistanceRecord.TotalDistance + TestCar.Speed",
        "Log(\"Speed increased\")"
    ]
}`

const arrayJSONData = `[{
    "name": "SpeedUp",
    "desc": "When testcar is speeding up we keep increase the speed.",
    "salience": 10,
    "when": "TestCar.SpeedUp == true && TestCar.Speed < TestCar.MaxSpeed",
    "then": [
        "TestCar.Speed = TestCar.Speed + TestCar.SpeedIncrement",
        "DistanceRecord.TotalDistance = DistanceRecord.TotalDistance + TestCar.Speed",
        "Log(\"Speed increased\")"
    ]
}]`

const jsonDataExpanded = `{
    "name": "SpeedUp",
    "desc": "When testcar is speeding up we keep increase the speed.",
    "salience": 10,
    "when": {
       "and": [
           {"eq": ["TestCar.SpeedUp", true]},
           {"lt": ["TestCar.Speed", "TestCar.MaxSpeed"]}
       ]
    },
    "then": [
        {"set": ["TestCar.Speed", {"plus": ["TestCar.Speed", "TestCar.SpeedIncrement"]}]},
        {"set": ["DistanceRecord.TotalDistance", {"plus": ["DistanceRecord.TotalDistance", "TestCar.Speed"]}]},
        {"call": ["Log", {"const": "Speed increased"}]}
    ]
}`

const arrayJSONDataExpanded = `[{
    "name": "SpeedUp",
    "desc": "When testcar is speeding up we keep increase the speed.",
    "salience": 10,
    "when": {
       "and": [
           {"eq": ["TestCar.SpeedUp", true]},
           {"lt": ["TestCar.Speed", "TestCar.MaxSpeed"]}
       ]
    },
    "then": [
        {"set": ["TestCar.Speed", {"plus": ["TestCar.Speed", "TestCar.SpeedIncrement"]}]},
        {"set": ["DistanceRecord.TotalDistance", {"plus": ["DistanceRecord.TotalDistance", "TestCar.Speed"]}]},
        {"call": ["Log", {"const": "Speed increased"}]}
    ]
}]`

const jsonDataVerbose = `{
    "name": "SpeedUp",
    "desc": "When testcar is speeding up we keep increase the speed.",
    "salience": 10,
    "when": {
       "and": [
           {"eq": [{"obj": "TestCar.SpeedUp"}, {"const": true}]},
           {"lt": [{"obj": "TestCar.Speed"}, {"obj": "TestCar.MaxSpeed"}]}
       ]
    },
    "then": [
        {"set": [{"obj": "TestCar.Speed"}, {"plus": [{"obj": "TestCar.Speed"}, {"obj": "TestCar.SpeedIncrement"}]}]},
        {"set": [{"obj": "DistanceRecord.TotalDistance"}, {"plus": [{"obj": "DistanceRecord.TotalDistance"}, {"obj": "TestCar.Speed"}]}]},
        {"call": ["Log", {"const": "Speed increased"}]}
    ]
}`

const arrayJSONDataVerbose = `[{
    "name": "SpeedUp",
    "desc": "When testcar is speeding up we keep increase the speed.",
    "salience": 10,
    "when": {
       "and": [
           {"eq": [{"obj": "TestCar.SpeedUp"}, {"const": true}]},
           {"lt": [{"obj": "TestCar.Speed"}, {"obj": "TestCar.MaxSpeed"}]}
       ]
    },
    "then": [
        {"set": [{"obj": "TestCar.Speed"}, {"plus": [{"obj": "TestCar.Speed"}, {"obj": "TestCar.SpeedIncrement"}]}]},
        {"set": [{"obj": "DistanceRecord.TotalDistance"}, {"plus": [{"obj": "DistanceRecord.TotalDistance"}, {"obj": "TestCar.Speed"}]}]},
        {"call": ["Log", {"const": "Speed increased"}]}
    ]
}]`

const expectedRule = `rule SpeedUp "When testcar is speeding up we keep increase the speed." salience 10 {
    when
        TestCar.SpeedUp == true && TestCar.Speed < TestCar.MaxSpeed
    then
        TestCar.Speed = TestCar.Speed + TestCar.SpeedIncrement;
        DistanceRecord.TotalDistance = DistanceRecord.TotalDistance + TestCar.Speed;
        Log("Speed increased");
}
`

const jsonDataEscaped = `{
    "name": "SpeedUp",
    "desc": "When testcar is speeding up we keep increase the \"speed\".",
    "salience": 10,
    "when": {
       "and": [
           {"eq": [{"obj": "TestCar.SpeedUp"}, {"const": true}]},
           {"lt": [{"obj": "TestCar.Speed"}, {"obj": "TestCar.MaxSpeed"}]}
       ]
    },
    "then": [
        {"set": [{"obj": "TestCar.Speed"}, {"plus": [{"obj": "TestCar.Speed"}, {"obj": "TestCar.SpeedIncrement"}]}]},
        {"set": [{"obj": "DistanceRecord.TotalDistance"}, {"plus": [{"obj": "DistanceRecord.TotalDistance"}, {"obj": "TestCar.Speed"}]}]},
        {"call": ["Log", {"const": "\"Speed\" increased\n"}]}
    ]
}`

const arrayJSONDataEscaped = `[{
    "name": "SpeedUp",
    "desc": "When testcar is speeding up we keep increase the \"speed\".",
    "salience": 10,
    "when": {
       "and": [
           {"eq": [{"obj": "TestCar.SpeedUp"}, {"const": true}]},
           {"lt": [{"obj": "TestCar.Speed"}, {"obj": "TestCar.MaxSpeed"}]}
       ]
    },
    "then": [
        {"set": [{"obj": "TestCar.Speed"}, {"plus": [{"obj": "TestCar.Speed"}, {"obj": "TestCar.SpeedIncrement"}]}]},
        {"set": [{"obj": "DistanceRecord.TotalDistance"}, {"plus": [{"obj": "DistanceRecord.TotalDistance"}, {"obj": "TestCar.Speed"}]}]},
        {"call": ["Log", {"const": "\"Speed\" increased\n"}]}
    ]
}]`

const expectedEscaped = `rule SpeedUp "When testcar is speeding up we keep increase the \"speed\"." salience 10 {
    when
        TestCar.SpeedUp == true && TestCar.Speed < TestCar.MaxSpeed
    then
        TestCar.Speed = TestCar.Speed + TestCar.SpeedIncrement;
        DistanceRecord.TotalDistance = DistanceRecord.TotalDistance + TestCar.Speed;
        Log("\"Speed\" increased\n");
}
`

const jsonDataBigIntConversion = `{
    "name": "SpeedUp",
    "desc": "When testcar is speeding up we keep increase the \"speed\".",
    "salience": 10,
    "when": {
       "and": [
           {"eq": [{"obj": "TestCar.SpeedUp"}, {"const": true}]},
           {"lt": [{"obj": "TestCar.Speed"}, {"const": 10000000}]}
       ]
    },
    "then": [
        {"set": [{"obj": "TestCar.Speed"}, {"plus": [{"obj": "TestCar.Speed"}, {"obj": "TestCar.SpeedIncrement"}]}]},
        {"set": [{"obj": "DistanceRecord.TotalDistance"}, {"plus": [{"obj": "DistanceRecord.TotalDistance"}, {"obj": "TestCar.Speed"}]}]},
        {"call": ["Log", {"const": "\"Speed\" increased\n"}]}
    ]
}`

const arrayJSONDataBigIntConversion = `[{
    "name": "SpeedUp",
    "desc": "When testcar is speeding up we keep increase the \"speed\".",
    "salience": 10,
    "when": {
       "and": [
           {"eq": [{"obj": "TestCar.SpeedUp"}, {"const": true}]},
           {"lt": [{"obj": "TestCar.Speed"}, {"const": 10000000}]}
       ]
    },
    "then": [
        {"set": [{"obj": "TestCar.Speed"}, {"plus": [{"obj": "TestCar.Speed"}, {"obj": "TestCar.SpeedIncrement"}]}]},
        {"set": [{"obj": "DistanceRecord.TotalDistance"}, {"plus": [{"obj": "DistanceRecord.TotalDistance"}, {"obj": "TestCar.Speed"}]}]},
        {"call": ["Log", {"const": "\"Speed\" increased\n"}]}
    ]
}]`

const expectedBigIntConversion = `rule SpeedUp "When testcar is speeding up we keep increase the \"speed\"." salience 10 {
    when
        TestCar.SpeedUp == true && TestCar.Speed < 10000000
    then
        TestCar.Speed = TestCar.Speed + TestCar.SpeedIncrement;
        DistanceRecord.TotalDistance = DistanceRecord.TotalDistance + TestCar.Speed;
        Log("\"Speed\" increased\n");
}
`

func TestParseJSONRuleset(t *testing.T) {
	rs, err := ParseJSONRule([]byte(jsonData))
	assert.NoError(t, err)
	assert.Equal(t, expectedRule, rs)
	rs, err = ParseJSONRuleset([]byte(arrayJSONData))
	assert.NoError(t, err)
	assert.Equal(t, expectedRule, rs)
	rs, err = ParseJSONRule([]byte(jsonDataExpanded))
	assert.NoError(t, err)
	assert.Equal(t, expectedRule, rs)
	rs, err = ParseJSONRuleset([]byte(arrayJSONDataExpanded))
	assert.NoError(t, err)
	assert.Equal(t, expectedRule, rs)
	rs, err = ParseJSONRule([]byte(jsonDataVerbose))
	assert.NoError(t, err)
	assert.Equal(t, expectedRule, rs)
	rs, err = ParseJSONRuleset([]byte(arrayJSONDataVerbose))
	assert.NoError(t, err)
	assert.Equal(t, expectedRule, rs)
}

func TestNewJSONResourceFromResource(t *testing.T) {
	underlyingResource := NewBytesResource([]byte(jsonDataExpanded))
	resource, err := NewJSONResourceFromResource(underlyingResource)
	assert.NoError(t, err)
	loaded, err := resource.Load()
	assert.NoError(t, err)
	assert.Equal(t, expectedRule, string(loaded))
	underlyingResource = NewBytesResource([]byte(arrayJSONDataExpanded))
	resource, err = NewJSONResourceFromResource(underlyingResource)
	assert.NoError(t, err)
	loaded, err = resource.Load()
	assert.NoError(t, err)
	assert.Equal(t, expectedRule, string(loaded))
}

func TestJSONStringEscaping(t *testing.T) {
	rs, err := ParseJSONRule([]byte(jsonDataEscaped))
	assert.NoError(t, err)
	assert.Equal(t, expectedEscaped, rs)
	rs, err = ParseJSONRuleset([]byte(arrayJSONDataEscaped))
	assert.NoError(t, err)
	assert.Equal(t, expectedEscaped, rs)
}

func TestJSONBigIntConversion(t *testing.T) {
	rs, err := ParseJSONRule([]byte(jsonDataBigIntConversion))
	assert.NoError(t, err)
	assert.Equal(t, expectedBigIntConversion, rs)
	rs, err = ParseJSONRuleset([]byte(arrayJSONDataBigIntConversion))
	assert.NoError(t, err)
	assert.Equal(t, expectedBigIntConversion, rs)
}
