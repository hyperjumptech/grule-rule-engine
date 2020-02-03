package antlr

import (
	"fmt"
	"github.com/antlr/antlr4/runtime/Go/antlr"
	parser "github.com/hyperjumptech/grule-rule-engine/antlr/parser/grulev2.g4"
	"github.com/hyperjumptech/grule-rule-engine/ast"
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
			fmt.Println(nt.GetText())
		}
	}

}

func TestV2Parser(t *testing.T) {
	logrus.SetLevel(logrus.DebugLevel)
	data, err := ioutil.ReadFile("./sample2.grl")
	if err != nil {
		t.Fatal(err)
	} else {
		sdata := string(data)

		is := antlr.NewInputStream(sdata)
		lexer := parser.Newgrulev2Lexer(is)
		stream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)

		var parseError error

		memory := ast.NewWorkingMemory()
		kb := ast.NewKnowledgeBase("KB", "1.0.0")

		listener := NewGruleV2ParserListener(kb, memory, func(e error) {
			parseError = e
		})

		psr := parser.Newgrulev2Parser(stream)
		psr.BuildParseTrees = true
		antlr.ParseTreeWalkerDefault.Walk(listener, psr.Root())

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
)

func TestV2Parser2(t *testing.T) {
	logrus.SetLevel(logrus.DebugLevel)

	sdata := rules

	is := antlr.NewInputStream(sdata)
	lexer := parser.Newgrulev2Lexer(is)
	stream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)

	var parseError error

	memory := ast.NewWorkingMemory()
	kb := ast.NewKnowledgeBase("KB", "1.0.0")

	listener := NewGruleV2ParserListener(kb, memory, func(e error) {
		parseError = e
		panic(e)
	})

	psr := parser.Newgrulev2Parser(stream)
	psr.BuildParseTrees = true
	antlr.ParseTreeWalkerDefault.Walk(listener, psr.Root())

	if parseError != nil {
		t.Log(parseError)
		t.FailNow()
	}
}
