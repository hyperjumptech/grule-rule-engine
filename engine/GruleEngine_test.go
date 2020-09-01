package engine

import (
	"context"
	"fmt"
	"github.com/hyperjumptech/grule-rule-engine/ast"
	"github.com/hyperjumptech/grule-rule-engine/builder"
	"github.com/hyperjumptech/grule-rule-engine/pkg"
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
	assert.NoError(t, err)
	err = dctx.Add("DistanceRecord", dr)
	assert.NoError(t, err)

	lib := ast.NewKnowledgeLibrary()
	rb := builder.NewRuleBuilder(lib)
	err = rb.BuildRuleFromResource("Test", "0.1.1", pkg.NewBytesResource([]byte(rules)))
	assert.NoError(t, err)
	engine := NewGruleEngine()
	kb := lib.NewKnowledgeBaseInstance("Test", "0.1.1")
	start := time.Now()
	err = engine.Execute(dctx, kb)
	assert.NoError(t, err)
	dur := time.Since(start)
	t.Log(dr.TotalDistance)
	t.Logf("Duration %d ms", dur.Milliseconds())
}

func getTypeOf(i interface{}) string {
	t := reflect.TypeOf(i)
	if t.Kind() == reflect.Ptr {
		return fmt.Sprintf("*%s", t.Elem().Name())
	}
	return t.Name()
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
	assert.NoError(t, err)

	lib := ast.NewKnowledgeLibrary()
	rb := builder.NewRuleBuilder(lib)
	err = rb.BuildRuleFromResource("Test", "0.1.1", pkg.NewBytesResource([]byte(complexRule1)))
	assert.NoError(t, err)
	kb := lib.NewKnowledgeBaseInstance("Test", "0.1.1")

	engine := NewGruleEngine()
	err = engine.Execute(dctx, kb)
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
	assert.NoError(t, err)

	lib := ast.NewKnowledgeLibrary()
	rb := builder.NewRuleBuilder(lib)
	err = rb.BuildRuleFromResource("Test", "0.1.1", pkg.NewBytesResource([]byte(complexRule2)))
	assert.NoError(t, err)
	kb := lib.NewKnowledgeBaseInstance("Test", "0.1.1")

	engine := NewGruleEngine()
	err = engine.Execute(dctx, kb)
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
	assert.NoError(t, err)

	lib := ast.NewKnowledgeLibrary()
	rb := builder.NewRuleBuilder(lib)
	err = rb.BuildRuleFromResource("Test", "0.1.1", pkg.NewBytesResource([]byte(complexRule3)))
	assert.NoError(t, err)
	kb := lib.NewKnowledgeBaseInstance("Test", "0.1.1")

	engine := NewGruleEngine()
	err = engine.Execute(dctx, kb)
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
	assert.NoError(t, err)

	lib := ast.NewKnowledgeLibrary()
	rb := builder.NewRuleBuilder(lib)
	err = rb.BuildRuleFromResource("Test", "0.1.1", pkg.NewBytesResource([]byte(complexRule4)))
	assert.NoError(t, err)
	kb := lib.NewKnowledgeBaseInstance("Test", "0.1.1")

	engine := NewGruleEngine()
	err = engine.Execute(dctx, kb)
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
	assert.NoError(t, err)

	lib := ast.NewKnowledgeLibrary()
	rb := builder.NewRuleBuilder(lib)
	err = rb.BuildRuleFromResource("Test", "0.1.1", pkg.NewBytesResource([]byte(OpPresedenceRule)))
	assert.NoError(t, err)
	kb := lib.NewKnowledgeBaseInstance("Test", "0.1.1")

	engine := NewGruleEngine()
	err = engine.Execute(dctx, kb)
	assert.NoError(t, err)

	assert.Equal(t, int64(3), ts.Result)
}

type ESTest struct {
	Var1    string
	Var2    string
	Result1 bool
	Result2 bool
}

const escapedRules = `rule EscapedRuleA "test string escaping" salience 10 {
    when
        ESTest.Var1 == "C:\\Windows\\System32\\ntdll.dll"
    then
        ESTest.Result1 = true;
		Retract("EscapedRuleA");
}

rule EscapedRuleB "test string escaping" salience 10 {
    when
        ESTest.Var2 == "some \"escaped\" string value\nAnd another"
    then
        ESTest.Result2 = true;
		Retract("EscapedRuleB");
}`

func TestEngine_EscapedStrings(t *testing.T) {
	es := &ESTest{}
	es.Var1 = `C:\Windows\System32\ntdll.dll`
	es.Var2 = `some "escaped" string value
And another`

	dctx := ast.NewDataContext()
	err := dctx.Add("ESTest", es)
	assert.NoError(t, err)

	lib := ast.NewKnowledgeLibrary()
	rb := builder.NewRuleBuilder(lib)
	err = rb.BuildRuleFromResource("Test", "0.1.1", pkg.NewBytesResource([]byte(escapedRules)))
	assert.NoError(t, err)
	kb := lib.NewKnowledgeBaseInstance("Test", "0.1.1")

	assert.False(t, es.Result1)
	assert.False(t, es.Result2)

	engine := NewGruleEngine()
	err = engine.Execute(dctx, kb)
	assert.NoError(t, err)

	assert.True(t, es.Result1)
	assert.True(t, es.Result2)
}

type Sleeper struct {
	Count int
}

func (s *Sleeper) SleepMore() {
	time.Sleep(1 * time.Second)
}

func TestGruleEngine_ExecuteWithContext(t *testing.T) {
	ts := &Sleeper{}

	dctx := ast.NewDataContext()
	err := dctx.Add("TS", ts)
	assert.NoError(t, err)

	lib := ast.NewKnowledgeLibrary()
	rb := builder.NewRuleBuilder(lib)
	err = rb.BuildRuleFromResource("TestTimer", "0.1.1", pkg.NewBytesResource([]byte(`
rule KeepSleep "test string escaping" salience 10 {
    when
        TS.Count < 4
    then
        TS.Count = TS.Count + 1;
		TS.SleepMore();
}
`)))
	assert.NoError(t, err)
	kb := lib.NewKnowledgeBaseInstance("TestTimer", "0.1.1")
	engine := NewGruleEngine()
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	err = engine.ExecuteWithContext(ctx, dctx, kb)
	if err == nil {
		t.Logf("Should have failed since its was timeout")
		t.Fail()
	} else {
		t.Logf("got %s", err.Error())
	}
}

type Fact struct {
	NetAmount float32
	Distance  int32
	Duration  int32
	Result    bool
}

const duplicateRules = `rule  DuplicateRule1  "Duplicate Rule 1"  salience 10 {
when
(Fact.Distance > 5000  &&   Fact.Duration > 120) && (Fact.Result == false)
Then
   Fact.NetAmount=143.320007;
   Fact.Result=true;
}
rule  DuplicateRule2  "Duplicate Rule 2"  salience 10 {
when
(Fact.Distance > 5000  &&   Fact.Duration > 120) && (Fact.Result == false)
Then
   Fact.NetAmount=143.320007;
   Fact.Result=true;
}


rule  DuplicateRule3  "Duplicate Rule 3"  salience 10 {
when
(Fact.Distance > 5000  &&   Fact.Duration > 120) && (Fact.Result == false)
Then
   Fact.NetAmount=143.320007;
   Fact.Result=true;
}


rule  DuplicateRule4  "Duplicate Rule 4"  salience 10 {
when
(Fact.Distance > 5000  &&   Fact.Duration > 120) && (Fact.Result == false)
Then
   Fact.NetAmount=143.320007;
   Fact.Result=true;
}


rule  DuplicateRule5  "Duplicate Rule 5"  salience 10 {
when
(Fact.Distance > 5000  &&   Fact.Duration > 120) && (Fact.Result == false)
Then
   Output.NetAmount=143.320007;
   Fact.Result=true;
}`

func TestGruleEngine_FetchMatchingRules_Having_Same_Salience(t *testing.T) {
	//Given
	fact := &Fact{
		Distance: 6000,
		Duration: 123,
	}
	dctx := ast.NewDataContext()
	err := dctx.Add("Fact", fact)
	assert.NoError(t, err)
	lib := ast.NewKnowledgeLibrary()
	rb := builder.NewRuleBuilder(lib)
	err = rb.BuildRuleFromResource("conflict_rules_test", "0.1.1", pkg.NewBytesResource([]byte(duplicateRules)))
	assert.NoError(t, err)
	kb := lib.NewKnowledgeBaseInstance("conflict_rules_test", "0.1.1")

	//When
	engine := NewGruleEngine()
	ruleEntries, err := engine.FetchMatchingRules(dctx, kb)

	//Then
	assert.NoError(t, err)
	assert.Equal(t, 5, len(ruleEntries))
}

const duplicateRulesWithDiffSalience = `rule  DuplicateRule1  "Duplicate Rule 1"  salience 5 {
when
(Fact.Distance > 5000  &&   Fact.Duration > 120) && (Fact.Result == false)
Then
   Fact.NetAmount=143.320007;
   Fact.Result=true;
}
rule  DuplicateRule2  "Duplicate Rule 2"  salience 6 {
when
(Fact.Distance > 5000  &&   Fact.Duration > 120) && (Fact.Result == false)
Then
   Fact.NetAmount=143.320007;
   Fact.Result=true;
}


rule  DuplicateRule3  "Duplicate Rule 3"  salience 7 {
when
(Fact.Distance > 5000  &&   Fact.Duration > 120) && (Fact.Result == false)
Then
   Fact.NetAmount=143.320007;
   Fact.Result=true;
}


rule  DuplicateRule4  "Duplicate Rule 4"  salience 8 {
when
(Fact.Distance > 5000  &&   Fact.Duration > 120) && (Fact.Result == false)
Then
   Fact.NetAmount=143.320007;
   Fact.Result=true;
}


rule  DuplicateRule5  "Duplicate Rule 5"  salience 9 {
when
(Fact.Distance > 5000  &&   Fact.Duration == 120) && (Fact.Result == false)
Then
   Output.NetAmount=143.320007;
   Fact.Result=true;
}`

func TestGruleEngine_FetchMatchingRules_Having_Diff_Salience(t *testing.T) {
	//Given
	fact := &Fact{
		Distance: 6000,
		Duration: 121,
	}
	dctx := ast.NewDataContext()
	err := dctx.Add("Fact", fact)
	assert.NoError(t, err)
	lib := ast.NewKnowledgeLibrary()
	rb := builder.NewRuleBuilder(lib)
	err = rb.BuildRuleFromResource("conflict_rules_test", "0.1.1", pkg.NewBytesResource([]byte(duplicateRulesWithDiffSalience)))
	assert.NoError(t, err)
	kb := lib.NewKnowledgeBaseInstance("conflict_rules_test", "0.1.1")

	//When
	engine := NewGruleEngine()
	ruleEntries, err := engine.FetchMatchingRules(dctx, kb)

	//Then
	assert.NoError(t, err)
	assert.Equal(t, 4, len(ruleEntries))
	assert.Equal(t, 8, ruleEntries[0].Salience.SalienceValue)
	assert.Equal(t, 7, ruleEntries[1].Salience.SalienceValue)
	assert.Equal(t, 6, ruleEntries[2].Salience.SalienceValue)
	assert.Equal(t, 5, ruleEntries[3].Salience.SalienceValue)
}

//This TestCase is to test whether grule-rule-engine follows logical operator precedence
// ! - Highest Priority
// && - Medium Priority
// || - Lowest Priority
// Credits: https://chortle.ccsu.edu/java5/Notes/chap40/ch40_16.html
const logicalOperatorPrecedenceRules = `
rule  ComplicatedLogicalOperatorRule  "Complicated logical operator rule" {
when
Fact.Distance > 5000  ||   Fact.Duration > 120 || Fact.RideType == "On-Demand" && Fact.IsFrequentCustomer == true
Then
   Fact.NetAmount=143.320007;
   Fact.Result=true;
   Complete();
}`

/**
Evaluation must be done below way if you follow logical operator precedence (identify parentheses arrangement)
(Fact.Distance > 5000  ||   Fact.Duration > 120 || (Fact.RideType == "On-Demand" && Fact.IsFrequentCustomer == true))
Result:
Logical Operator Precedence: true
No precedence: false
**/
type LogicalOperatorRuleFact struct {
	Distance           int32
	Duration           int32
	RideType           string
	IsFrequentCustomer bool
	Result             bool
	NetAmount          float32
}

func TestGruleEngine_Follows_logical_operator_precedence(t *testing.T) {
	//Given
	fact := &LogicalOperatorRuleFact{
		Distance:           2000,
		Duration:           121,
		RideType:           "Pre-Booked",
		IsFrequentCustomer: true,
	}
	dctx := ast.NewDataContext()
	err := dctx.Add("Fact", fact)
	assert.NoError(t, err)
	lib := ast.NewKnowledgeLibrary()
	rb := builder.NewRuleBuilder(lib)
	err = rb.BuildRuleFromResource("logical_operator_rules_test", "0.1.1", pkg.NewBytesResource([]byte(logicalOperatorPrecedenceRules)))
	assert.NoError(t, err)
	kb := lib.NewKnowledgeBaseInstance("logical_operator_rules_test", "0.1.1")

	//When
	engine := NewGruleEngine()
	err = engine.Execute(dctx, kb)

	//Then
	assert.NoError(t, err)
	assert.Equal(t, fact.Result, true)
	assert.Equal(t, fact.NetAmount, float32(143.32))
}
