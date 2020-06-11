package examples

import (
	"fmt"
	"github.com/hyperjumptech/grule-rule-engine/ast"
	"github.com/hyperjumptech/grule-rule-engine/builder"
	engine2 "github.com/hyperjumptech/grule-rule-engine/engine"
	"github.com/hyperjumptech/grule-rule-engine/pkg"
	"time"
)

const (
	// ItemTypeLuxury is a labels as Luxury item
	ItemTypeLuxury = "LUXURY"
	// ItemTypeNormal is a labels as Non Luxury item
	ItemTypeNormal = "NORMAL"
)

// Purchase stores a purchasing example
type Purchase struct {
	PurchaseDate time.Time
	ItemType     string
	Price        float64

	Tax             float64
	PriceAfterTax   float64
	IgnoredPurchase bool
}

// CashFlow stores simulated cash flow
type CashFlow struct {
	TotalPurchases    float64
	PurchaseCount     int
	TotalTax          float64
	PurchasesAfterTax float64
}

// String shows a cash flow.
func (cf *CashFlow) String() string {
	return fmt.Sprintf("Purchase count %d total amount %f. Total tax are %f thus total cash in %f", cf.PurchaseCount, cf.TotalPurchases, cf.TotalTax, cf.PurchasesAfterTax)
}

var (
	// Purchases contains list of purchases to be evaluated.
	Purchases = []*Purchase{
		&Purchase{
			PurchaseDate: time.Date(2019, time.January, 4, 13, 0, 0, 0, time.Local),
			ItemType:     ItemTypeLuxury,
			Price:        100000,
		},
		&Purchase{
			PurchaseDate: time.Date(2019, time.January, 17, 15, 0, 0, 0, time.Local),
			ItemType:     ItemTypeLuxury,
			Price:        100000,
		},
		&Purchase{
			PurchaseDate: time.Date(2019, time.February, 12, 7, 0, 0, 0, time.Local),
			ItemType:     ItemTypeLuxury,
			Price:        100000,
		},
		&Purchase{
			PurchaseDate: time.Date(2019, time.February, 24, 3, 0, 0, 0, time.Local),
			ItemType:     ItemTypeLuxury,
			Price:        100000,
		},
		&Purchase{
			PurchaseDate: time.Date(2019, time.March, 22, 22, 0, 0, 0, time.Local),
			ItemType:     ItemTypeLuxury,
			Price:        100000,
		},
		&Purchase{
			PurchaseDate: time.Date(2019, time.March, 24, 17, 0, 0, 0, time.Local),
			ItemType:     ItemTypeLuxury,
			Price:        100000,
		},
		&Purchase{
			PurchaseDate: time.Date(2019, time.March, 15, 14, 0, 0, 0, time.Local),
			ItemType:     ItemTypeLuxury,
			Price:        100000,
		},
		&Purchase{
			PurchaseDate: time.Date(2019, time.March, 25, 10, 0, 0, 0, time.Local),
			ItemType:     ItemTypeLuxury,
			Price:        100000,
		},
		&Purchase{
			PurchaseDate: time.Date(2019, time.March, 19, 13, 0, 0, 0, time.Local),
			ItemType:     ItemTypeLuxury,
			Price:        100000,
		},
		&Purchase{
			PurchaseDate: time.Date(2019, time.June, 6, 21, 0, 0, 0, time.Local),
			ItemType:     ItemTypeLuxury,
			Price:        100000,
		},
		&Purchase{
			PurchaseDate: time.Date(2019, time.June, 19, 22, 0, 0, 0, time.Local),
			ItemType:     ItemTypeLuxury,
			Price:        100000,
		},
	}
)

// CashFlowCalculator to simulate a calculator
type CashFlowCalculator struct {
}

// CalculatePurchases will calculate a speciffic purchase by rule engine.
func (cf *CashFlowCalculator) CalculatePurchases() {
	cashFlow := &CashFlow{}

	lib := ast.NewKnowledgeLibrary()
	rb := builder.NewRuleBuilder(lib)
	err := rb.BuildRuleFromResource("Purchase Calculator", "0.0.1", pkg.NewFileResource("CashFlowRule.grl"))

	engine := engine2.NewGruleEngine()

	kb := lib.NewKnowledgeBaseInstance("Purchase Calculator", "0.0.1")

	for _, purchase := range Purchases {
		dctx := ast.NewDataContext()
		dctx.Add("CashFlow", cashFlow)
		dctx.Add("Purchase", purchase)
		err = engine.Execute(dctx, kb)
		if err != nil {
			panic(err)
		}
	}

	fmt.Println(cashFlow.String())
}
