package examples

import (
	"fmt"
	"github.com/antlr/antlr4/runtime/Go/antlr"
	antlr2 "github.com/hyperjumptech/grule-rule-engine/antlr"
	"github.com/hyperjumptech/grule-rule-engine/antlr/parser"
	"github.com/hyperjumptech/grule-rule-engine/model"
	"testing"
)

func TestItemPriceChecker_TestLexer(t *testing.T) {
	is := antlr.NewInputStream(PriceCheckRule)

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

func TestItemPriceChecker_TestParser(t *testing.T) {
	nis := antlr.NewInputStream(PriceCheckRule)

	lexer := parser.NewgruleLexer(nis)
	stream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)

	listener := antlr2.NewGruleParserListener(model.NewKnowledgeBase())

	psr := parser.NewgruleParser(stream)
	psr.BuildParseTrees = true
	antlr.ParseTreeWalkerDefault.Walk(listener, psr.Root())

	for _, e := range listener.ParseErrors {
		t.Log(e)
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
