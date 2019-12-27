package model

import (
	"github.com/hyperjumptech/grule-rule-engine/context"
	"reflect"
)

// Constant holds a constants, it holds a simple golang value.
type Constant struct {
	ConstantValue    reflect.Value
	knowledgeContext *context.KnowledgeContext
	ruleCtx          *RuleContext
	dataCtx          *context.DataContext
}

// Initialize will initialize this graph with context
func (cons *Constant) Initialize(knowledgeContext *context.KnowledgeContext, ruleCtx *RuleContext, dataCtx *context.DataContext) {
	cons.knowledgeContext = knowledgeContext
	cons.ruleCtx = ruleCtx
	cons.dataCtx = dataCtx
}

// Evaluate the object graph against underlined context or execute evaluation in the sub graph.
func (cons *Constant) Evaluate() (reflect.Value, error) {
	return cons.ConstantValue, nil
}

// AcceptDecimal prepare this graph with a decimal value.
func (cons *Constant) AcceptDecimal(val int64) error {
	cons.ConstantValue = reflect.ValueOf(val)
	return nil
}

// EqualsTo will compare two literal constants, be it string, int, uint, floats bools and nils
func (cons *Constant) EqualsTo(that AlphaNode) bool {
	typ := reflect.TypeOf(that)
	if that == nil {
		return false
	}
	if typ.Kind() == reflect.Ptr {
		if typ.Elem().Name() == "Constant" {
			thatConstant := that.(*Constant)
			if cons.ConstantValue.Kind() == thatConstant.ConstantValue.Kind() {
				switch cons.ConstantValue.Kind() {
				case reflect.String:
					return cons.ConstantValue.String() == thatConstant.ConstantValue.String()
				case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
					return cons.ConstantValue.Int() == thatConstant.ConstantValue.Int()
				case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
					return cons.ConstantValue.Uint() == thatConstant.ConstantValue.Uint()
				case reflect.Float32, reflect.Float64:
					return cons.ConstantValue.Float() == thatConstant.ConstantValue.Float()
				case reflect.Bool:
					return cons.ConstantValue.Bool() == thatConstant.ConstantValue.Bool()
				case reflect.Invalid:
					return thatConstant.ConstantValue.Kind() == reflect.Invalid
				default:
					println(cons.ConstantValue.Kind().String())
				}
			}
		}
	}
	return false
}
