package antlr

import (
	"fmt"
	"github.com/antlr/antlr4/runtime/Go/antlr"
	parser "github.com/hyperjumptech/grule-rule-engine/antlr/parser/grulev3"
	"github.com/hyperjumptech/grule-rule-engine/ast/v2"
	v3 "github.com/hyperjumptech/grule-rule-engine/ast/v3"
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"reflect"
	"testing"
)

func TestV3Lexer(t *testing.T) {
	data, err := ioutil.ReadFile("./sample4.grl")
	if err != nil {
		t.Fatal(err)
	} else {
		is := antlr.NewInputStream(string(data))

		// Create the Lexer
		lexer := parser.Newgrulev3Lexer(is)
		//lexer := parser.NewLdifParserLexer(is)

		// Read all tokens
		for {
			nt := lexer.NextToken()
			fmt.Println(nt.GetText())
			if nt.GetTokenType() == antlr.TokenEOF {
				break
			}
		}
	}

}

func TestV3Parser(t *testing.T) {
	// logrus.SetLevel(logrus.TraceLevel)
	data, err := ioutil.ReadFile("./sample4.grl")
	if err != nil {
		t.Fatal(err)
	} else {
		sdata := string(data)

		is := antlr.NewInputStream(sdata)
		lexer := parser.Newgrulev3Lexer(is)
		stream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)

		var parseError error
		kb := v3.NewKnowledgeLibrary().GetKnowledgeBase("T", "1")

		listener := NewGruleV3ParserListener(kb, func(e error) {
			parseError = e
		})

		psr := parser.Newgrulev3Parser(stream)
		psr.BuildParseTrees = true
		antlr.ParseTreeWalkerDefault.Walk(listener, psr.Grl())

		if parseError != nil {
			t.Log(parseError)
			t.FailNow()
		}
	}
}

func TestV3Parser2(t *testing.T) {
	// logrus.SetLevel(logrus.InfoLevel)

	sdata := rules

	is := antlr.NewInputStream(sdata)
	lexer := parser.Newgrulev3Lexer(is)
	stream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)

	var parseError error
	kb := v3.NewKnowledgeLibrary().GetKnowledgeBase("T", "1")

	listener := NewGruleV3ParserListener(kb, func(e error) {
		parseError = e
		panic(e)
	})

	psr := parser.Newgrulev3Parser(stream)
	psr.BuildParseTrees = true
	antlr.ParseTreeWalkerDefault.Walk(listener, psr.Grl())

	if parseError != nil {
		t.Log(parseError)
		t.FailNow()
	}
}

func TestV3ParserEscapedStringInvalid(t *testing.T) {
	// logrus.SetLevel(logrus.DebugLevel)

	is := antlr.NewInputStream(invalidEscapeRule)
	lexer := parser.Newgrulev3Lexer(is)
	stream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)

	var parseError error
	kb := v3.NewKnowledgeLibrary().GetKnowledgeBase("T", "1")

	listener := NewGruleV3ParserListener(kb, func(e error) {
		parseError = e
	})

	psr := parser.Newgrulev3Parser(stream)
	psr.BuildParseTrees = true
	antlr.ParseTreeWalkerDefault.Walk(listener, psr.Grl())

	if parseError == nil {
		t.Fatal("Successfully parsed invalid string literal, should have gotten an error")
	}
}

func TestV3ParserEscapedStringValid(t *testing.T) {
	// logrus.SetLevel(logrus.DebugLevel)

	is := antlr.NewInputStream(validEscapeRule)
	lexer := parser.Newgrulev3Lexer(is)
	stream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)

	var parseError error
	kb := v3.NewKnowledgeLibrary().GetKnowledgeBase("T", "1")

	listener := NewGruleV3ParserListener(kb, func(e error) {
		parseError = e
	})

	psr := parser.Newgrulev3Parser(stream)
	psr.BuildParseTrees = true
	antlr.ParseTreeWalkerDefault.Walk(listener, psr.Grl())

	if parseError != nil {
		t.Fatal("Failed to parse rule with escaped string constant")
	}
}

func TestV3ParserSnapshotEyeBalling(t *testing.T) {
	// logrus.SetLevel(logrus.TraceLevel)

	data := `
rule SpeedUp "When testcar is speeding up we keep increase the speed." salience 10 {
    when
        TestCar.SpeedUp == true && TestCar.Speed < TestCar.MaxSpeed
    then
        TestCar.Speed = TestCar.Speed + TestCar.SpeedIncrement;
		DistanceRecord.TotalDistance = DistanceRecord.TotalDistance + TestCar.Speed;
}
`

	is := antlr.NewInputStream(data)
	lexer := parser.Newgrulev3Lexer(is)
	stream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)

	var parseError error
	kb := v3.NewKnowledgeLibrary().GetKnowledgeBase("T", "1")

	listener := NewGruleV3ParserListener(kb, func(e error) {
		parseError = e
	})

	psr := parser.Newgrulev3Parser(stream)
	psr.BuildParseTrees = true
	antlr.ParseTreeWalkerDefault.Walk(listener, psr.Grl())
	if parseError != nil {
		t.Log(parseError)
		t.FailNow()
	}

	listener.KnowledgeBase.WorkingMemory.IndexVariables()
	if listener.Grl.RuleEntries["SpeedUp"] == nil {
		t.Fatalf("Rule entry not exist")
	}
	if listener.Grl.RuleEntries["SpeedUp"].WhenScope == nil {
		t.Fatalf("When scope not exist")
	}
	if listener.Grl.RuleEntries["SpeedUp"].ThenScope == nil {
		t.Fatalf("Then scope not exist")
	}
	t.Log(listener.Grl.RuleEntries["SpeedUp"].GetSnapshot())
	if listener.Grl.RuleEntries["SpeedUp"].ThenScope.ThenExpressionList == nil {
		t.Fatalf("Then expression list is not exist")
	}
	if listener.Grl.RuleEntries["SpeedUp"].ThenScope.ThenExpressionList.ThenExpressions == nil {
		t.Fatalf("Then expression list array is not exist")
	}
	if len(listener.Grl.RuleEntries["SpeedUp"].ThenScope.ThenExpressionList.ThenExpressions) != 2 {
		t.Fatalf("Then expression list array %s contains not 2 but %d", listener.Grl.RuleEntries["SpeedUp"].ThenScope.ThenExpressionList.GetAstID(), len(listener.Grl.RuleEntries["SpeedUp"].ThenScope.ThenExpressionList.ThenExpressions))
	}
}

func prepareV3TestKnowledgeBase(t *testing.T, grl string) (*v3.KnowledgeBase, *v3.WorkingMemory) {
	is := antlr.NewInputStream(grl)
	lexer := parser.Newgrulev3Lexer(is)
	stream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)

	var parseError error
	kb := v3.NewKnowledgeLibrary().GetKnowledgeBase("T", "1")

	listener := NewGruleV3ParserListener(kb, func(e error) {
		parseError = e
	})

	psr := parser.Newgrulev3Parser(stream)
	psr.BuildParseTrees = true
	antlr.ParseTreeWalkerDefault.Walk(listener, psr.Grl())
	assert.NoError(t, parseError)
	listener.KnowledgeBase.WorkingMemory.IndexVariables()
	return kb, kb.WorkingMemory
}

func TestV3ConstantFunctionAndConstantFunctionChain(t *testing.T) {
	// logrus.SetLevel(logrus.InfoLevel)
	dctx := v3.NewDataContext()

	data := `
rule RuleOne "RuleOneDesc" salience 123 {
    when
        "    true    ".Trim() != "true"
    then
        Log("      Success    ".Trim().ToUpper());
}
`
	kb, wm := prepareV3TestKnowledgeBase(t, data)
	err := dctx.Add("DEFUNC", &v3.BuiltInFunctions{
		Knowledge:     kb,
		WorkingMemory: wm,
		DataContext:   dctx,
	})

	assert.NoError(t, err)

	whenVal, err := kb.RuleEntries["RuleOne"].WhenScope.Evaluate(dctx, wm)
	assert.NoError(t, err)
	assert.True(t, whenVal.IsValid())
	assert.Equal(t, reflect.Bool, whenVal.Kind())
	assert.False(t, whenVal.Bool())
	assert.NotNil(t, kb.RuleEntries["RuleOne"].ThenScope)
	assert.NotNil(t, kb.RuleEntries["RuleOne"].ThenScope.ThenExpressionList)
	assert.NotNil(t, kb.RuleEntries["RuleOne"].ThenScope.ThenExpressionList.ThenExpressions)
	assert.Equal(t, 1, len(kb.RuleEntries["RuleOne"].ThenScope.ThenExpressionList.ThenExpressions))
	assert.NotNil(t, kb.RuleEntries["RuleOne"].ThenScope.ThenExpressionList.ThenExpressions[0])
	err = kb.RuleEntries["RuleOne"].ThenScope.ThenExpressionList.ThenExpressions[0].Execute(dctx, wm)
	assert.NoError(t, err)
}

func TestV3RuleRetract(t *testing.T) {
	// logrus.SetLevel(logrus.InfoLevel)
	dctx := v3.NewDataContext()

	data := `
rule RuleOne "RuleOneDesc" salience 123 {
    when
        "    true    ".Trim() != "true"
    then
        Retract("RuleOne");
}
`

	kb, wm := prepareV3TestKnowledgeBase(t, data)
	err := dctx.Add("DEFUNC", &v3.BuiltInFunctions{
		Knowledge:     kb,
		WorkingMemory: wm,
		DataContext:   dctx,
	})

	assert.False(t, kb.RuleEntries["RuleOne"].Retracted)
	assert.NotNil(t, kb.RuleEntries["RuleOne"].ThenScope)
	assert.NotNil(t, kb.RuleEntries["RuleOne"].ThenScope.ThenExpressionList)
	assert.NotNil(t, kb.RuleEntries["RuleOne"].ThenScope.ThenExpressionList.ThenExpressions)
	assert.Equal(t, 1, len(kb.RuleEntries["RuleOne"].ThenScope.ThenExpressionList.ThenExpressions))
	assert.NotNil(t, kb.RuleEntries["RuleOne"].ThenScope.ThenExpressionList.ThenExpressions[0])

	err = kb.RuleEntries["RuleOne"].ThenScope.ThenExpressionList.ThenExpressions[0].Execute(dctx, wm)
	if err != nil {
		t.Log(err)
		t.FailNow()
	}
	assert.True(t, kb.RuleEntries["RuleOne"].Retracted)
	kb.Reset()
	assert.False(t, kb.RuleEntries["RuleOne"].Retracted)
}

func TestV3RuleAssignment(t *testing.T) {
	// logrus.SetLevel(logrus.TraceLevel)
	dctx := v2.NewDataContext()

	data := `
rule RuleOne "RuleOneDesc" salience 123 {
    when
        Person.Name == "Rudolf"
    then
        Person.Name="Pearson".ToUpper();
}
`

	kb, wm := prepareV3TestKnowledgeBase(t, data)
	err := dctx.Add("DEFUNC", &v3.BuiltInFunctions{
		Knowledge:     kb,
		WorkingMemory: wm,
		DataContext:   dctx,
	})

	p := &Person{
		Name: "Rudolf",
	}
	err = dctx.Add("Person", p)

	assert.NoError(t, err)
	assert.False(t, kb.RuleEntries["RuleOne"].Retracted)
	assert.NotNil(t, kb.RuleEntries["RuleOne"].ThenScope)
	assert.NotNil(t, kb.RuleEntries["RuleOne"].ThenScope.ThenExpressionList)
	assert.NotNil(t, kb.RuleEntries["RuleOne"].ThenScope.ThenExpressionList.ThenExpressions)

	ret, err := kb.RuleEntries["RuleOne"].WhenScope.Evaluate(dctx, wm)
	assert.NoError(t, err)
	assert.True(t, ret.IsValid())
	assert.Equal(t, reflect.Bool, ret.Kind())
	assert.True(t, ret.Bool())

	assert.Equal(t, 1, len(kb.RuleEntries["RuleOne"].ThenScope.ThenExpressionList.ThenExpressions))
	assert.NotNil(t, kb.RuleEntries["RuleOne"].ThenScope.ThenExpressionList.ThenExpressions[0])
	err = kb.RuleEntries["RuleOne"].ThenScope.ThenExpressionList.ThenExpressions[0].Execute(dctx, wm)
	assert.NoError(t, err)
	assert.Equal(t, "PEARSON", p.Name)

	ret, err = kb.RuleEntries["RuleOne"].WhenScope.Evaluate(dctx, wm)
	assert.NoError(t, err)
	assert.True(t, ret.IsValid())
	assert.Equal(t, reflect.Bool, ret.Kind())
	assert.False(t, ret.Bool())

}