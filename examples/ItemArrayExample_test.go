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

package examples

import (
	"fmt"
	"github.com/antlr/antlr4/runtime/Go/antlr"
	antlr2 "github.com/hyperjumptech/grule-rule-engine/antlr"
	parser3 "github.com/hyperjumptech/grule-rule-engine/antlr/parser/grulev3"
	"github.com/hyperjumptech/grule-rule-engine/ast"
	"github.com/hyperjumptech/grule-rule-engine/builder"
	"github.com/hyperjumptech/grule-rule-engine/engine"
	"github.com/hyperjumptech/grule-rule-engine/pkg"
	"github.com/stretchr/testify/assert"
	"testing"
)

var (
	// PriceCheckRule1 is a rule definition for test.
	PriceCheckRule1 = `
rule ApplyDiscountForItemAbove100 "If an item prices is above 100 we give them discount"  {
    when
        Item.Price > 100 && Item.Discount == 0
    then
        Log("Rule applying discount to " + Item.Name);
		Item.Discount = 10;
}`

	// PriceCheckRule2 is a rule definition for test.
	PriceCheckRule2 = `
rule ApplyDiscountForCart "If a cart contain item prices is above 100 we give them discount"  {
    when
        Cart.CountItemWithPriceAboveWithNoDiscount(100) > 0
    then
        Cart.GiveDiscountForItemPriceAbove(100,10);
		Changed("CountItemWithPriceAboveWithNoDiscount");
		Log("Applying discount to cart item");
}
`
)

// Item store item definition.
type Item struct {
	Name     string
	Price    int64
	Discount int64
}

// ItemPriceChecker serve a checker object
type ItemPriceChecker struct {
}

// CheckPrices will test the rule to check item prices.
func (cf *ItemPriceChecker) CheckPrices(t *testing.T) {
	// Our array of items
	items := make([]*Item, 0)
	items = append(items, &Item{
		Name:  "Honda",
		Price: 80,
	}, &Item{
		Name:  "Toyota",
		Price: 90,
	}, &Item{
		Name:  "Bugatti",
		Price: 200,
	}, &Item{
		Name:  "Mazda",
		Price: 110,
	})

	lib := ast.NewKnowledgeLibrary()
	rb := builder.NewRuleBuilder(lib)

	// Prepare knowledgebase and load it with our rule.
	err := rb.BuildRuleFromResource("PriceCheck", "0.0.1", pkg.NewBytesResource([]byte(PriceCheckRule1)))
	assert.NoError(t, err)

	kb := lib.NewKnowledgeBaseInstance("PriceCheck", "0.0.1")

	// Prepare the engine
	eng := engine.NewGruleEngine()

	// Execute every item in to the engine.
	// Handling of the array is this program's job.
	// Let the rule decide for every item.
	for _, v := range items {
		dctx := ast.NewDataContext()
		err = dctx.Add("Item", v)
		assert.NoError(t, err)
		err = eng.Execute(dctx, kb)
		assert.NoError(t, err)

		if v.Discount > 0 {
			fmt.Printf("%s got %d discount\n", v.Name, v.Discount)
		}
		fmt.Println("---")

	}
}

// ItemCart simulates a shopping cart.
type ItemCart struct {
	Items []*Item
}

// CountItemWithPriceAboveWithNoDiscount will count items in a cart that have some minimum price
func (cart *ItemCart) CountItemWithPriceAboveWithNoDiscount(minimumPrice int64) int {
	count := 0
	for _, v := range cart.Items {
		if v.Price > minimumPrice && v.Discount == 0 {
			count++
		}
	}
	return count
}

// ShowDiscount will print out each discounted items.
func (cart *ItemCart) ShowDiscount() {
	for _, v := range cart.Items {
		fmt.Printf("Name %s Price %d Discount %d\n", v.Name, v.Price, v.Discount)
	}
}

// GiveDiscountForItemPriceAbove will apply discount to items that have some minimum price.
func (cart *ItemCart) GiveDiscountForItemPriceAbove(minimumPrice int64, discount int64) {
	fmt.Println("Applying discount for item in cart")
	for _, v := range cart.Items {
		if v.Price > minimumPrice {
			v.Discount = discount
		}
	}
}

// CheckCart will execute cart checking.
func (cf *ItemPriceChecker) CheckCart(t *testing.T) {
	// Our array of items
	items := make([]*Item, 0)
	items = append(items, &Item{
		Name:  "Honda",
		Price: 80,
	}, &Item{
		Name:  "Toyota",
		Price: 90,
	}, &Item{
		Name:  "Bugatti",
		Price: 200,
	}, &Item{
		Name:  "Mazda",
		Price: 110,
	})
	cart := &ItemCart{Items: items}

	// Prepare knowledgebase library and load it with our rule.
	lib := ast.NewKnowledgeLibrary()
	rb := builder.NewRuleBuilder(lib)
	err := rb.BuildRuleFromResource("Cart Check Rules", "0.0.1", pkg.NewBytesResource([]byte(PriceCheckRule2)))
	assert.NoError(t, err)

	kb := lib.NewKnowledgeBaseInstance("Cart Check Rules", "0.0.1")

	// Prepare the engine
	eng := engine.NewGruleEngine()

	dctx := ast.NewDataContext()
	err = dctx.Add("Cart", cart)
	assert.NoError(t, err)
	err = eng.Execute(dctx, kb)
	assert.NoError(t, err)
	cart.ShowDiscount()
}

func TestItemPriceChecker_TestLexer(t *testing.T) {
	is := antlr.NewInputStream(PriceCheckRule1)

	// Create the Lexer
	lexer := parser3.Newgrulev3Lexer(is)
	//lexer := parser.NewLdifParserLexer(is)

	// Read all tokens
	for {
		nt := lexer.NextToken()
		if nt.GetTokenType() == antlr.TokenEOF {
			break
		}
		t.Logf(nt.GetText())
	}
}

func TestItemPriceChecker_TestParser(t *testing.T) {
	nis := antlr.NewInputStream(PriceCheckRule1)

	lexer := parser3.Newgrulev3Lexer(nis)
	stream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)

	lib := ast.NewKnowledgeLibrary()
	kb := lib.GetKnowledgeBase("Test", "0.1.1")

	var parseError error
	listener := antlr2.NewGruleV3ParserListener(kb, func(e error) {
		parseError = e
	})

	psr := parser3.Newgrulev3Parser(stream)
	psr.BuildParseTrees = true
	antlr.ParseTreeWalkerDefault.Walk(listener, psr.Grl())
	assert.NoError(t, parseError)

}

func TestItemPriceChecker_CheckPrices(t *testing.T) {
	c := &ItemPriceChecker{}
	c.CheckPrices(t)
}

func TestItemPriceChecker_CheckCart(t *testing.T) {
	c := &ItemPriceChecker{}
	c.CheckCart(t)
}
