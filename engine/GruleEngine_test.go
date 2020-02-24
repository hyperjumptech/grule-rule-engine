package engine

import (
	"fmt"
	"github.com/hyperjumptech/grule-rule-engine/ast"
	"github.com/hyperjumptech/grule-rule-engine/builder"
	"github.com/hyperjumptech/grule-rule-engine/events"
	"github.com/hyperjumptech/grule-rule-engine/pkg"
	"github.com/hyperjumptech/grule-rule-engine/pkg/eventbus"
	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
	"reflect"
	"sort"
	"testing"
	"time"
)

type Sorting struct {
	Val int
}

func TestGruleSorting(t *testing.T) {
	arr := make([]*Sorting, 0)
	arr = append(arr, &Sorting{Val: 4})
	arr = append(arr, &Sorting{Val: 7})
	arr = append(arr, &Sorting{Val: 3})
	arr = append(arr, &Sorting{Val: 6})
	arr = append(arr, &Sorting{Val: 9})
	arr = append(arr, &Sorting{Val: 8})
	arr = append(arr, &Sorting{Val: 1})
	arr = append(arr, &Sorting{Val: 2})

	sort.Slice(arr, func(i, j int) bool {
		return arr[i].Val > arr[j].Val
	})

	if arr[0].Val != 9 {
		t.FailNow()
	}
}

type TestCar struct {
	SpeedUp        bool
	Speed          int
	MaxSpeed       int
	SpeedIncrement int
}

type DistanceRecorder struct {
	TotalDistance int
	TestTime      time.Time
}

const (
	rules = `
rule SpeedUp "When testcar is speeding up we keep increase the speed." salience 10 {
    when
        TestCar.SpeedUp == true && TestCar.Speed < TestCar.MaxSpeed
    then
        TestCar.Speed = TestCar.Speed + TestCar.SpeedIncrement;
		DistanceRecord.TotalDistance = DistanceRecord.TotalDistance + TestCar.Speed;
}

rule StartSpeedDown "When testcar is speeding up and over max speed we change to speed down." salience 10  {
    when
        TestCar.SpeedUp == true && TestCar.Speed >= TestCar.MaxSpeed
    then
        TestCar.SpeedUp = false;
		Log("Now we slow down");
}

rule SlowDown "When testcar is slowing down we keep decreasing the speed." salience 10  {
    when
        TestCar.SpeedUp == false && TestCar.Speed > 0
    then
        TestCar.Speed = TestCar.Speed - TestCar.SpeedIncrement;
		DistanceRecord.TotalDistance = DistanceRecord.TotalDistance + TestCar.Speed;
}

rule SetTime "When Distance Recorder time not set, set it." {
	when
		IsZero(DistanceRecord.TestTime)
	then
		Log("Set the test time");
		DistanceRecord.TestTime = Now();
		Log(TimeFormat(DistanceRecord.TestTime,"Mon Jan _2 15:04:05 2006"));
}
`
)

func TestGrule_Execute(t *testing.T) {
	logrus.SetLevel(logrus.DebugLevel)
	tc := &TestCar{
		SpeedUp:        true,
		Speed:          0,
		MaxSpeed:       100,
		SpeedIncrement: 2,
	}
	dr := &DistanceRecorder{
		TotalDistance: 0,
	}
	dctx := ast.NewDataContext()
	err := dctx.Add("TestCar", tc)
	if err != nil {
		t.Fatal(err)
	}
	err = dctx.Add("DistanceRecord", dr)
	if err != nil {
		t.Fatal(err)
	}
	memory := ast.NewWorkingMemory()
	kb := ast.NewKnowledgeBase("Test", "0.1.1")
	rb := builder.NewRuleBuilder(kb, memory)
	err = rb.BuildRuleFromResource(pkg.NewBytesResource([]byte(rules)))
	if err != nil {
		t.Errorf("Got error : %v", err)
		t.FailNow()
	} else {
		engine := NewGruleEngine()
		start := time.Now()
		err = engine.Execute(dctx, kb, memory)
		if err != nil {
			t.Errorf("Got error : %v", err)
			t.FailNow()
		} else {
			dur := time.Since(start)
			t.Log(dr.TotalDistance)
			t.Logf("Duration %d ms", dur.Milliseconds())
		}
	}
}

func getTypeOf(i interface{}) string {
	t := reflect.TypeOf(i)
	if t.Kind() == reflect.Ptr {
		return fmt.Sprintf("*%s", t.Elem().Name())
	}
	return t.Name()
}

func TestGrule_ExecuteWithSubscribers(t *testing.T) {
	logrus.SetLevel(logrus.InfoLevel)
	tc := &TestCar{
		SpeedUp:        true,
		Speed:          0,
		MaxSpeed:       100,
		SpeedIncrement: 2,
	}
	dr := &DistanceRecorder{
		TotalDistance: 0,
	}
	dctx := ast.NewDataContext()
	err := dctx.Add("TestCar", tc)
	if err != nil {
		t.Fatal(err)
	}
	err = dctx.Add("DistanceRecord", dr)
	if err != nil {
		t.Fatal(err)
	}

	ruleEntrySubscriber := eventbus.DefaultBrooker.GetSubscriber(events.RuleEntryEventTopic, func(i interface{}) error {
		if i != nil && getTypeOf(i) == "*RuleEntryEvent" {
			event := i.(*events.RuleEntryEvent)
			if event.EventType == events.RuleEntryExecuteStartEvent {
				log.Infof("Rule executed %s", event.RuleName)
			}
		} else if i != nil {
			log.Infof("RuleEntry Subscriber, Receive type is %s ", getTypeOf(i))
		}
		return nil
	})
	ruleEntrySubscriber.Subscribe()

	ruleEngineSubscriber := eventbus.DefaultBrooker.GetSubscriber(events.RuleEngineEventTopic, func(i interface{}) error {
		if i != nil && getTypeOf(i) == "*RuleEngineEvent" {
			event := i.(*events.RuleEngineEvent)
			if event.EventType == events.RuleEngineEndEvent {
				log.Infof("Engine finished in %d cycles", event.Cycle)
			}
		} else if i != nil {
			log.Infof("RuleEngine Subscriber, Receive type is %s ", getTypeOf(i))
		}
		return nil
	})
	ruleEngineSubscriber.Subscribe()

	//f := func(r *ast.RuleEntry) {
	//	log.Debugf("Now executing rule %s", r.Name)
	//}

	memory := ast.NewWorkingMemory()
	kb := ast.NewKnowledgeBase("Test", "0.1.1")
	rb := builder.NewRuleBuilder(kb, memory)
	err = rb.BuildRuleFromResource(pkg.NewBytesResource([]byte(rules)))
	if err != nil {
		t.Errorf("Got error : %v", err)
		t.FailNow()
	} else {
		engine := NewGruleEngine()

		start := time.Now()
		err = engine.Execute(dctx, kb, memory)
		if err != nil {
			t.Errorf("Got error : %v", err)
			t.FailNow()
		} else {
			dur := time.Since(start)
			t.Log(dr.TotalDistance)
			t.Logf("Duration %d ms", dur.Milliseconds())
		}
	}
}

func TestEmptyValueEquality(t *testing.T) {
	t1 := time.Time{}
	tv1 := reflect.ValueOf(t1)
	tv2 := reflect.Zero(tv1.Type())

	if tv1.Type() != tv2.Type() {
		t.Logf("%s vs %s", tv1.Type().String(), tv2.Type().String())
		t.FailNow()
	}

	if pkg.ValueToInterface(tv1) != pkg.ValueToInterface(tv2) {
		t.Logf("%s vs %s", tv1.Kind().String(), tv2.Kind().String())
		t.Logf("%s vs %s", tv1.Type().String(), tv2.Type().String())
		t.Logf("%v vs %v", tv1.IsValid(), tv2.IsValid())

		t.FailNow()
	}
}

type TestStruct struct {
	Param1 bool
	Param2 bool
	Param3 bool
	Param4 bool
	Result int64
}

const complexRule1 = `rule ComplexRule "test complex rule" salience 10 {
    when
        TestStruct.Param1 == true && TestStruct.Param2 == true || 
		TestStruct.Param3 == true && TestStruct.Param4 == true
    then
        TestStruct.Result = 1;
		Retract("ComplexRule");
}`

func TestEngine_ComplexRule1(t *testing.T) {

	ts := &TestStruct{
		Param1: true,
		Param2: true,
		Param3: true,
		Param4: true,
	}

	dctx := ast.NewDataContext()
	err := dctx.Add("TestStruct", ts)
	if err != nil {
		t.Fatal(err)
	}

	memory := ast.NewWorkingMemory()
	kb := ast.NewKnowledgeBase("Test", "0.0.1")
	rb := builder.NewRuleBuilder(kb, memory)
	err = rb.BuildRuleFromResource(pkg.NewBytesResource([]byte(complexRule1)))
	assert.NoError(t, err)

	engine := NewGruleEngine()
	err = engine.Execute(dctx, kb, memory)
	assert.NoError(t, err)

	assert.Equal(t, int64(1), ts.Result)
}

const complexRule2 = `rule ComplexRule "test complex rule" salience 10 {
    when
        TestStruct.Param1 == true && TestStruct.Param2 == true || 
		TestStruct.Param3 == true && TestStruct.Param4 == false
    then
        TestStruct.Result = 1;
		Retract("ComplexRule");
}`

func TestEngine_ComplexRule2(t *testing.T) {

	ts := &TestStruct{
		Param1: false,
		Param2: false,
		Param3: true,
		Param4: false,
	}

	dctx := ast.NewDataContext()
	err := dctx.Add("TestStruct", ts)
	if err != nil {
		t.Fatal(err)
	}

	memory := ast.NewWorkingMemory()
	kb := ast.NewKnowledgeBase("Test", "0.0.1")
	rb := builder.NewRuleBuilder(kb, memory)
	err = rb.BuildRuleFromResource(pkg.NewBytesResource([]byte(complexRule2)))
	assert.NoError(t, err)

	engine := NewGruleEngine()
	err = engine.Execute(dctx, kb, memory)
	assert.NoError(t, err)

	assert.Equal(t, int64(1), ts.Result)
}

const complexRule3 = `rule ComplexRule "test complex rule" salience 10 {
    when
        TestStruct.Param1 == true && TestStruct.Param2 == true  || 
		TestStruct.Param1 == true && TestStruct.Param3 == false ||
		TestStruct.Param4 == true
    then
        TestStruct.Result = 1;
		Retract("ComplexRule");
}`

func TestEngine_ComplexRule3(t *testing.T) {

	ts := &TestStruct{
		Param1: false,
		Param2: false,
		Param3: true,
		Param4: true,
	}

	dctx := ast.NewDataContext()
	err := dctx.Add("TestStruct", ts)
	if err != nil {
		t.Fatal(err)
	}

	memory := ast.NewWorkingMemory()
	kb := ast.NewKnowledgeBase("Test", "0.0.1")
	rb := builder.NewRuleBuilder(kb, memory)
	err = rb.BuildRuleFromResource(pkg.NewBytesResource([]byte(complexRule3)))
	assert.NoError(t, err)

	engine := NewGruleEngine()
	err = engine.Execute(dctx, kb, memory)
	assert.NoError(t, err)

	assert.Equal(t, int64(1), ts.Result)
}

const complexRule4 = `rule ComplexRule "test complex rule" salience 10 {
    when
        TestStruct.Param1 == true 	&& 
		(TestStruct.Param2 == true 	|| 
		TestStruct.Param3 == true	|| 
		TestStruct.Param4 == false) 
    then
        TestStruct.Result = 1;
		Retract("ComplexRule");
}`

func TestEngine_ComplexRule4(t *testing.T) {

	ts := &TestStruct{
		Param1: true,
		Param2: false,
		Param3: true,
		Param4: true,
	}

	dctx := ast.NewDataContext()
	err := dctx.Add("TestStruct", ts)
	if err != nil {
		t.Fatal(err)
	}

	memory := ast.NewWorkingMemory()
	kb := ast.NewKnowledgeBase("Test", "0.0.1")
	rb := builder.NewRuleBuilder(kb, memory)
	err = rb.BuildRuleFromResource(pkg.NewBytesResource([]byte(complexRule4)))
	assert.NoError(t, err)

	engine := NewGruleEngine()
	err = engine.Execute(dctx, kb, memory)
	assert.NoError(t, err)

	assert.Equal(t, int64(1), ts.Result)
}

const OpPresedenceRule = `rule OpPresedenceRule "test operator presedence" salience 10 {
    when
        1 + 2 + 3 * 4 == 15
    then
        TestStruct.Result = 3;
		Retract("OpPresedenceRule");
}`

func TestEngine_OperatorPrecedence(t *testing.T) {

	ts := &TestStruct{}

	dctx := ast.NewDataContext()
	err := dctx.Add("TestStruct", ts)
	if err != nil {
		t.Fatal(err)
	}

	memory := ast.NewWorkingMemory()
	kb := ast.NewKnowledgeBase("Test", "0.0.1")
	rb := builder.NewRuleBuilder(kb, memory)
	err = rb.BuildRuleFromResource(pkg.NewBytesResource([]byte(OpPresedenceRule)))
	assert.NoError(t, err)

	engine := NewGruleEngine()
	err = engine.Execute(dctx, kb, memory)
	assert.NoError(t, err)

	assert.Equal(t, int64(3), ts.Result)
}
