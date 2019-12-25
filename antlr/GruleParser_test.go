package antlr

import (
	"fmt"
	"github.com/antlr/antlr4/runtime/Go/antlr"
	"github.com/hyperjumptech/grule-rule-engine/antlr/parser"
	"github.com/hyperjumptech/grule-rule-engine/model"
	"io/ioutil"
	"reflect"
	"testing"
	"time"
)

func TestLexer(t *testing.T) {
	data, err := ioutil.ReadFile("./sample2.grl")
	if err != nil {
		t.Fatal(err)
	} else {
		is := antlr.NewInputStream(string(data))

		// Create the Lexer
		lexer := parser.NewgruleLexer(is)
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

func TestParser(t *testing.T) {
	data, err := ioutil.ReadFile("./sample2.grl")
	if err != nil {
		t.Fatal(err)
	} else {
		sdata := string(data)

		is := antlr.NewInputStream(sdata)
		lexer := parser.NewgruleLexer(is)
		stream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)

		var parseError error
		listener := NewGruleParserListener(model.NewKnowledgeBase(), func(e error) {
			parseError = e
		})

		psr := parser.NewgruleParser(stream)
		psr.BuildParseTrees = true
		antlr.ParseTreeWalkerDefault.Walk(listener, psr.Root())

		if parseError != nil {
			t.Log(parseError)
			t.FailNow()
		}
	}

}

func TestTimeKind(t *testing.T) {
	n := time.Now()
	nt := reflect.TypeOf(n)
	fmt.Println(nt.String())
}
