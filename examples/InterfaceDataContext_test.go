package examples

import (
	"github.com/hyperjumptech/grule-rule-engine/ast"
	"github.com/hyperjumptech/grule-rule-engine/builder"
	"github.com/hyperjumptech/grule-rule-engine/engine"
	"github.com/hyperjumptech/grule-rule-engine/pkg"
	"github.com/stretchr/testify/assert"
	"testing"
)

type Vehicle interface {
	GetPrice() int64
	GetName() string
}

type Car struct {
	price int64
	name  string
}

func (c *Car) GetPrice() int64 {
	return c.price
}

func (c *Car) GetName() string {
	return c.name
}

type Motorcycle struct {
	price int64
	name  string
}

func (m *Motorcycle) GetPrice() int64 {
	return m.price
}

func (m *Motorcycle) GetName() string {
	return m.name
}

func TestInterfaceDataContext(t *testing.T) {
	rule := `rule VehiclePriceCheck "Checking vehicle price" salience 100 {
when
	v[1].GetPrice() > 60000
then
	Log("Too expensive");
	Retract("VehiclePriceCheck");
}`
	vehicles := []Vehicle{
		&Car{price: 50000, name: "Car A"},
		&Car{price: 90000, name: "Car B"},
		&Motorcycle{price: 50000, name: "Motor A"},
		&Motorcycle{price: 90000, name: "Motor A"},
	}

	lib := ast.NewKnowledgeLibrary()
	rb := builder.NewRuleBuilder(lib)
	err := rb.BuildRuleFromResource("CarPriceTest", "0.1.1", pkg.NewBytesResource([]byte(rule)))
	assert.NoError(t, err)
	kb, err := lib.NewKnowledgeBaseInstance("CarPriceTest", "0.1.1")
	assert.NoError(t, err)
	eng := &engine.GruleEngine{MaxCycle: 3}

	dataContext := ast.NewDataContext()
	err = dataContext.Add("v", vehicles)
	assert.NoError(t, err)

	err = eng.Execute(dataContext, kb)
	assert.NoError(t, err)
}
