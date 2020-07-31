package pkg

import (
	"testing"
)

const jsonData = `[{
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

const jsonDataExpanded = `[{
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

const jsonDataVerbose = `[{
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

const jsonDataEscaped = `[{
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

func TestParseJSONRuleset(t *testing.T) {
	rs, err := ParseJSONRuleset([]byte(jsonData))
	if err != nil {
		t.Fatal("Failed to parse flat ruleset: " + err.Error())
	}
	t.Log("Flat rule output:")
	t.Log(rs)
	if rs != expectedRule {
		t.Fatal("Parsed rule does not match expected result")
	}
	rs, err = ParseJSONRuleset([]byte(jsonDataExpanded))
	if err != nil {
		t.Fatal("Failed to parse expanded ruleset: " + err.Error())
	}
	t.Log("Expanded rule output:")
	t.Log(rs)
	if rs != expectedRule {
		t.Fatal("Parsed rule does not match expected result")
	}
	rs, err = ParseJSONRuleset([]byte(jsonDataVerbose))
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
		t.Fatal("Failed to load JSON ruleset: " + err.Error())
	}
	t.Log(string(loaded))
	if string(loaded) != expectedRule {
		t.Fatal("Loaded rule does not match expected result")
	}
}

func TestJSONStringEscaping(t *testing.T) {
	rs, err := ParseJSONRuleset([]byte(jsonDataEscaped))
	if err != nil {
		t.Fatal("Failed to parse flat ruleset: " + err.Error())
	}
	t.Log(rs)
	if rs != expectedEscaped {
		t.Fatal("Rule output doe not match expected value")
	}
}
