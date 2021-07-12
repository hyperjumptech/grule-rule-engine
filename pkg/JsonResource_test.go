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

package pkg

import (
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
	if err != nil {
		t.Fatal("Failed to parse flat rule: " + err.Error())
	}
	t.Log("Flat rule output:")
	t.Log(rs)
	if rs != expectedRule {
		t.Fatal("Parsed rule does not match expected result")
	}
	rs, err = ParseJSONRuleset([]byte(arrayJSONData))
	if err != nil {
		t.Fatal("Failed to parse flat ruleset: " + err.Error())
	}
	t.Log("Flat rule output:")
	t.Log(rs)
	if rs != expectedRule {
		t.Fatal("Parsed rule does not match expected result")
	}
	rs, err = ParseJSONRule([]byte(jsonDataExpanded))
	if err != nil {
		t.Fatal("Failed to parse expanded rule: " + err.Error())
	}
	t.Log("Expanded rule output:")
	t.Log(rs)
	if rs != expectedRule {
		t.Fatal("Parsed rule does not match expected result")
	}
	rs, err = ParseJSONRuleset([]byte(arrayJSONDataExpanded))
	if err != nil {
		t.Fatal("Failed to parse expanded ruleset: " + err.Error())
	}
	t.Log("Expanded rule output:")
	t.Log(rs)
	if rs != expectedRule {
		t.Fatal("Parsed rule does not match expected result")
	}
	rs, err = ParseJSONRule([]byte(jsonDataVerbose))
	if err != nil {
		t.Fatal("Failed to parse verbose rule: " + err.Error())
	}
	t.Log("Verbose rule output:")
	t.Log(rs)
	if rs != expectedRule {
		t.Fatal("Parsed rule does not match expected result")
	}
	rs, err = ParseJSONRuleset([]byte(arrayJSONDataVerbose))
	if err != nil {
		t.Fatal("Failed to parse verbose ruleset: " + err.Error())
	}
	t.Log("Verbose rule output:")
	t.Log(rs)
	if rs != expectedRule {
		t.Fatal("Parsed rule does not match expected result")
	}
}

func TestNewJSONResourceFromResource(t *testing.T) {
	underlyingResource := NewBytesResource([]byte(jsonDataExpanded))
	resource := NewJSONResourceFromResource(underlyingResource)
	loaded, err := resource.Load()
	if err != nil {
		t.Fatal("Failed to load JSON rule: " + err.Error())
	}
	t.Log(string(loaded))
	if string(loaded) != expectedRule {
		t.Fatal("Loaded rule does not match expected result")
	}
	underlyingResource = NewBytesResource([]byte(arrayJSONDataExpanded))
	resource = NewJSONResourceFromResource(underlyingResource)
	loaded, err = resource.Load()
	if err != nil {
		t.Fatal("Failed to load JSON ruleset: " + err.Error())
	}
	t.Log(string(loaded))
	if string(loaded) != expectedRule {
		t.Fatal("Loaded rule does not match expected result")
	}
}

func TestJSONStringEscaping(t *testing.T) {
	rs, err := ParseJSONRule([]byte(jsonDataEscaped))
	if err != nil {
		t.Fatal("Failed to parse flat rule: " + err.Error())
	}
	t.Log(rs)
	if rs != expectedEscaped {
		t.Fatal("Rule output doe not match expected value")
	}
	rs, err = ParseJSONRuleset([]byte(arrayJSONDataEscaped))
	if err != nil {
		t.Fatal("Failed to parse flat ruleset: " + err.Error())
	}
	t.Log(rs)
	if rs != expectedEscaped {
		t.Fatal("Rule output doe not match expected value")
	}
}

func TestJSONBigIntConversion(t *testing.T) {
	rs, err := ParseJSONRule([]byte(jsonDataBigIntConversion))
	if err != nil {
		t.Fatal("Failed to parse flat rule: " + err.Error())
	}
	t.Log(rs)
	if rs != expectedBigIntConversion {
		t.Fatal("Rule output doe not match expected value")
	}
	rs, err = ParseJSONRuleset([]byte(arrayJSONDataBigIntConversion))
	if err != nil {
		t.Fatal("Failed to parse flat ruleset: " + err.Error())
	}
	t.Log(rs)
	if rs != expectedBigIntConversion {
		t.Fatal("Rule output doe not match expected value")
	}
}
