package examples

import (
	"testing"

	"github.com/hyperjumptech/grule-rule-engine/ast"
	"github.com/hyperjumptech/grule-rule-engine/builder"
	"github.com/hyperjumptech/grule-rule-engine/engine"
	"github.com/hyperjumptech/grule-rule-engine/pkg"
	"github.com/stretchr/testify/assert"
)

const (
	trainRule = `
	rule TrainSpeedAdjust "Adjust Train speed based on passenger count" salience 10 {
        when 
            Train.PassengerCount > 50
        then
            Train.Speed = 75.45;
			Train.Stops = Train.Stops - 1;
			Retract("TrainSpeedAdjust");
    }
	`
)

type Train struct {
	Speed          *float32
	PassengerCount *uint16
	Stops          *int64
}

func TestSetNumberToPointer(t *testing.T) {

	speed := float32(204.31)
	passengerCount := uint16(100)
	stops := int64(3)
	train := &Train{
		Speed:          &speed,
		PassengerCount: &passengerCount,
		Stops:          &stops,
	}

	dataContext := ast.NewDataContext()
	err := dataContext.Add("Train", train)
	assert.NoError(t, err)

	lib := ast.NewKnowledgeLibrary()
	rb := builder.NewRuleBuilder(lib)
	err = rb.BuildRuleFromResource("TestSetNumberToPointer", "0.1.1", pkg.NewBytesResource([]byte(trainRule)))
	assert.NoError(t, err)
	eng := &engine.GruleEngine{MaxCycle: 5}
	kb, err := lib.NewKnowledgeBaseInstance("TestSetNumberToPointer", "0.1.1")
	assert.NoError(t, err)
	err = eng.Execute(dataContext, kb)
	assert.NoError(t, err)

	assert.Equal(t, int64(2), *train.Stops)
	assert.Equal(t, uint16(100), *train.PassengerCount)
	assert.Equal(t, float32(75.45), *train.Speed)
}
