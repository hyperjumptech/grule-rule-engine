package examples

import (
	"fmt"
	"github.com/antlr/antlr4/runtime/Go/antlr"
	antlr2 "github.com/hyperjumptech/grule-rule-engine/antlr"
	parser2 "github.com/hyperjumptech/grule-rule-engine/antlr/parser/grulev2.g4"
	"github.com/hyperjumptech/grule-rule-engine/ast"
	"testing"
)

func TestItemPriceChecker_TestLexer(t *testing.T) {
	is := antlr.NewInputStream(PriceCheckRule1)

	// Create the Lexer
	lexer := parser2.Newgrulev2Lexer(is)
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

func TestItemPriceChecker_TestParser(t *testing.T) {
	nis := antlr.NewInputStream(PriceCheckRule1)

	lexer := parser2.Newgrulev2Lexer(nis)
	stream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)

	mem := ast.NewWorkingMemory()

	var parseError error
	listener := antlr2.NewGruleV2ParserListener(ast.NewKnowledgeBase("Test", "0.1.1"), mem, func(e error) {
		parseError = e
	})

	psr := parser2.Newgrulev2Parser(stream)
	psr.BuildParseTrees = true
	antlr.ParseTreeWalkerDefault.Walk(listener, psr.Root())

	if parseError != nil {
		t.Log(parseError)
		t.FailNow()
	}

}

func TestItemPriceChecker_CheckPrices(t *testing.T) {
	c := &ItemPriceChecker{}
	c.CheckPrices()
}

func TestItemPriceChecker_CheckCart(t *testing.T) {
	c := &ItemPriceChecker{}
	c.CheckCart()
}
