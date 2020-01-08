package model

import (
	"fmt"
	"github.com/hyperjumptech/grule-rule-engine/context"
	"reflect"
)

// ArgumentHolder is a struct part of the rule object graph.
// It holds child graph such as Variable name, Constant data, Function, Expressions, etc.
type ArgumentHolder struct {
	Constant         *Constant
	Variable         string
	FunctionCall     *FunctionCall
	MethodCall       *MethodCall
	Expression       *Expression
	knowledgeContext *context.KnowledgeContext
	ruleCtx          *RuleContext
	dataCtx          *context.DataContext
}

// Initialize this ArgumentHolder instance graph before rule execution start.
func (ah *ArgumentHolder) Initialize(knowledgeContext *context.KnowledgeContext, ruleCtx *RuleContext, dataCtx *context.DataContext) {
	ah.knowledgeContext = knowledgeContext
	ah.ruleCtx = ruleCtx
	ah.dataCtx = dataCtx

	if ah.Constant != nil {
		ah.Constant.Initialize(knowledgeContext, ruleCtx, dataCtx)
	}
	if ah.FunctionCall != nil {
		ah.FunctionCall.Initialize(knowledgeContext, ruleCtx, dataCtx)
	}
	if ah.MethodCall != nil {
		ah.MethodCall.Initialize(knowledgeContext, ruleCtx, dataCtx)
	}
	if ah.Expression != nil {
		ah.Expression.Initialize(knowledgeContext, ruleCtx, dataCtx)
	}
}

// Evaluate the object graph against underlined context or execute evaluation in the sub graph.
func (ah *ArgumentHolder) Evaluate() (reflect.Value, error) {
	if len(ah.Variable) > 0 {
		return ah.dataCtx.GetValue(ah.Variable)
	}
	if ah.Constant != nil {
		return ah.Constant.Evaluate()
	}
	if ah.FunctionCall != nil {
		return ah.FunctionCall.Evaluate()
	}
	if ah.MethodCall != nil {
		return ah.MethodCall.Evaluate()
	}
	if ah.Expression != nil {
		return ah.Expression.Evaluate()
	}
	return reflect.ValueOf(nil), fmt.Errorf("argument holder stores no value")
}

// EqualsTo will compare two literal constants, be it string, int, uint, floats bools and nils
func (ah *ArgumentHolder) EqualsTo(that AlphaNode) bool {
	typ := reflect.TypeOf(that)
	if that == nil {
		return false
	}
	if typ.Kind() == reflect.Ptr {
		if typ.Elem().Name() == "ArgumentHolder" {
			thatAh := that.(*ArgumentHolder)
			if ah.Constant != nil && thatAh.Constant != nil {
				return ah.Constant.EqualsTo(thatAh.Constant)
			}
			if ah.FunctionCall != nil && thatAh.FunctionCall != nil {
				return ah.FunctionCall.EqualsTo(thatAh.FunctionCall)
			}
			if ah.MethodCall != nil && thatAh.MethodCall != nil {
				return ah.MethodCall.EqualsTo(thatAh.MethodCall)
			}
			if ah.Expression != nil && thatAh.Expression != nil {
				return ah.Expression.EqualsTo(thatAh.Expression)
			}
			if ah.Variable == thatAh.Variable {
				return true
			}
		}
	}
	return false
}

// IsContainVariable should check for this ArgumentHolder whether it contains a variable related argument
func (ah *ArgumentHolder) IsContainVariable(varName string) bool {
	if ah.Variable == varName {
		return true
	}
	if ah.MethodCall != nil {
		return ah.MethodCall.IsContainVariable(varName)
	}
	if ah.FunctionCall != nil {
		return ah.FunctionCall.IsContainVariable(varName)
	}
	return false
}
