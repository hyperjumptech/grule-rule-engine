package examples

import (
	"fmt"
	"github.com/hyperjumptech/grule-rule-engine/builder"
	"github.com/hyperjumptech/grule-rule-engine/context"
	"github.com/hyperjumptech/grule-rule-engine/engine"
	"github.com/hyperjumptech/grule-rule-engine/model"
	"github.com/hyperjumptech/grule-rule-engine/pkg"
)

var (
	// PriceCheckRule is a rule definition for test.
	PriceCheckRule = `
rule ApplyDiscountForItemAbove100 "If an item prices is above 100 we give them discount"  {
    when
        Item.Price > 100 && Item.Discount == 0
    then
        Log("Rule applying discount to " + Item.Name);
		Item.Discount = 10;
}

rule ApplyDiscountForCart "If a cart contain item prices is above 100 we give them discount"  {
    when
        Cart.CountItemWithPriceAboveWithNoDiscount(100) > 0
    then
        Cart.GiveDiscountForItemPriceAbove(100,10);
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
func (cf *ItemPriceChecker) CheckPrices() {
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

	// Prepare knowledgebase and load it with our rule.
	kb := model.NewKnowledgeBase()
	rb := builder.NewRuleBuilder(kb)
	err := rb.BuildRuleFromResource(pkg.NewBytesResource([]byte(PriceCheckRule)))
	if err != nil {
		panic(err)
	}

	// Prepare the engine
	eng := engine.NewGruleEngine()

	// Execute every item in to the engine.
	// Handling of the array is this program's job.
	// Let the rule decide for every item.
	for _, v := range items {
		dctx := context.NewDataContext()
		err = dctx.Add("Item", v)
		if err != nil {
			panic(err)
		}
		err = eng.Execute(dctx, kb)
		if err != nil {
			panic(err)
		}

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
func (cf *ItemPriceChecker) CheckCart() {
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

	// Prepare knowledgebase and load it with our rule.
	kb := model.NewKnowledgeBase()
	rb := builder.NewRuleBuilder(kb)
	err := rb.BuildRuleFromResource(pkg.NewBytesResource([]byte(PriceCheckRule)))
	if err != nil {
		panic(err)
	}

	// Prepare the engine
	eng := engine.NewGruleEngine()

	dctx := context.NewDataContext()
	err = dctx.Add("Cart", cart)
	if err != nil {
		panic(err)
	}
	err = eng.Execute(dctx, kb)
	if err != nil {
		panic(err)
	}
	cart.ShowDiscount()
}
