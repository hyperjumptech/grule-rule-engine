package antlr

import (
	"github.com/antlr/antlr4/runtime/Go/antlr"
	parser "github.com/hyperjumptech/grule-rule-engine/antlr/parser/grulev2"
	"github.com/hyperjumptech/grule-rule-engine/ast"
	"github.com/hyperjumptech/grule-rule-engine/builder"
	"github.com/hyperjumptech/grule-rule-engine/pkg"
	"github.com/sirupsen/logrus"
	"io/ioutil"
	"testing"
)

func TestV2Lexer(t *testing.T) {
	data, err := ioutil.ReadFile("./sample2.grl")
	if err != nil {
		t.Fatal(err)
	} else {
		is := antlr.NewInputStream(string(data))

		// Create the Lexer
		lexer := parser.Newgrulev2Lexer(is)
		//lexer := parser.NewLdifParserLexer(is)

		// Read all tokens
		for {
			nt := lexer.NextToken()
			if nt.GetTokenType() == antlr.TokenEOF {
				break
			}
		}
	}

}

func TestV2Parser(t *testing.T) {
	logrus.SetLevel(logrus.TraceLevel)
	data, err := ioutil.ReadFile("./sample3.grl")
	if err != nil {
		t.Fatal(err)
	} else {
		sdata := string(data)

		is := antlr.NewInputStream(sdata)
		lexer := parser.Newgrulev2Lexer(is)
		stream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)

		var parseError error
		kb := ast.NewKnowledgeLibrary().GetKnowledgeBase("T", "1")

		listener := NewGruleV2ParserListener(kb, func(e error) {
			parseError = e
		})

		psr := parser.Newgrulev2Parser(stream)
		psr.BuildParseTrees = true
		antlr.ParseTreeWalkerDefault.Walk(listener, psr.Grl())

		if parseError != nil {
			t.Log(parseError)
			t.FailNow()
		}
	}
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
		ABC.abc = "cde";
		Log("Set the test time");
		DistanceRecord.TestTime = Now();
		Log(TimeFormat(DistanceRecord.TestTime,"Mon Jan _2 15:04:05 2006"));
		ABC.abc = "cde";

}
`
	invalidEscapeRule = `rule SetTime "When Distance Recorder time not set, set it." {
	when
		IsZero(DistanceRecord.TestTime)
	then
		ABC.abc = "abc\cde";
		Log("Set the test time");
		DistanceRecord.TestTime = Now();
		Log(TimeFormat(DistanceRecord.TestTime,"Mon Jan _2 15:04:05 2006"));

}`
	validEscapeRule = `rule SetTime "When Distance Recorder time not set, set it." {
	when
		IsZero(DistanceRecord.TestTime)
	then
		ABC.abc = "abc\\cde";
		Log("Set the test time");
		DistanceRecord.TestTime = Now();
		Log(TimeFormat(DistanceRecord.TestTime,"Mon Jan _2 15:04:05 2006"));

}`
)

func TestV2Parser2(t *testing.T) {
	logrus.SetLevel(logrus.InfoLevel)

	sdata := rules

	is := antlr.NewInputStream(sdata)
	lexer := parser.Newgrulev2Lexer(is)
	stream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)

	var parseError error
	kb := ast.NewKnowledgeLibrary().GetKnowledgeBase("T", "1")

	listener := NewGruleV2ParserListener(kb, func(e error) {
		parseError = e
		panic(e)
	})

	psr := parser.Newgrulev2Parser(stream)
	psr.BuildParseTrees = true
	antlr.ParseTreeWalkerDefault.Walk(listener, psr.Grl())

	if parseError != nil {
		t.Log(parseError)
		t.FailNow()
	}
}

func TestV2ParserEscapedStringInvalid(t *testing.T) {
	logrus.SetLevel(logrus.DebugLevel)

	is := antlr.NewInputStream(invalidEscapeRule)
	lexer := parser.Newgrulev2Lexer(is)
	stream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)

	var parseError error
	kb := ast.NewKnowledgeLibrary().GetKnowledgeBase("T", "1")

	listener := NewGruleV2ParserListener(kb, func(e error) {
		parseError = e
	})

	psr := parser.Newgrulev2Parser(stream)
	psr.BuildParseTrees = true
	antlr.ParseTreeWalkerDefault.Walk(listener, psr.Grl())

	if parseError == nil {
		t.Fatal("Successfully parsed invalid string literal, should have gotten an error")
	}
}

func TestV2ParserEscapedStringValid(t *testing.T) {
	logrus.SetLevel(logrus.DebugLevel)

	is := antlr.NewInputStream(validEscapeRule)
	lexer := parser.Newgrulev2Lexer(is)
	stream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)

	var parseError error
	kb := ast.NewKnowledgeLibrary().GetKnowledgeBase("T", "1")

	listener := NewGruleV2ParserListener(kb, func(e error) {
		parseError = e
	})

	psr := parser.Newgrulev2Parser(stream)
	psr.BuildParseTrees = true
	antlr.ParseTreeWalkerDefault.Walk(listener, psr.Grl())

	if parseError != nil {
		t.Fatal("Failed to parse rule with escaped string constant")
	}
}

func TestV2ParserSnapshotEyeBalling(t *testing.T) {
	logrus.SetLevel(logrus.TraceLevel)

	data := `
rule SpeedUp "When testcar is speeding up we keep increase the speed." salience 10 {
    when
        TestCar.SpeedUp == true && TestCar.Speed < TestCar.MaxSpeed
    then
        TestCar.Speed = TestCar.Speed + TestCar.SpeedIncrement;
		DistanceRecord.TotalDistance = DistanceRecord.TotalDistance + TestCar.Speed;
}
`

	lib := ast.NewKnowledgeLibrary()
	rb := builder.NewRuleBuilder(lib)
	err := rb.BuildRuleFromResource("Test", "0.1.1", pkg.NewBytesResource([]byte(data)))

	if err != nil {
		t.Fatalf(err.Error())
	}

	kb := lib.GetKnowledgeBase("Test", "0.1.1")

	if len(kb.RuleEntries) != 1 {
		t.Fatalf("Expect 1 but %d", len(kb.RuleEntries))
	}
	t.Logf("WhenScope snapshot : %s", kb.RuleEntries["RuleName"].WhenScope.GetSnapshot())
}
