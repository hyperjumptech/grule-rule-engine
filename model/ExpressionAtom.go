package model

import (
	"github.com/hyperjumptech/grule-rule-engine/context"
	"github.com/hyperjumptech/grule-rule-engine/pkg"
	"github.com/juju/errors"
	"github.com/sirupsen/logrus"
	"reflect"
)

// ExpressionAtom holds an expression atom graph. it can form a mathematical expression, a simple contants, function  all, method call.
type ExpressionAtom struct {
	Text                string
	ExpressionAtomLeft  *ExpressionAtom
	ExpressionAtomRight *ExpressionAtom
	MathOperator        MathOperator
	Variable            string
	Constant            *Constant
	FunctionCall        *FunctionCall
	MethodCall          *MethodCall
	knowledgeContext    *context.KnowledgeContext
	ruleCtx             *RuleContext
	dataCtx             *context.DataContext

	evaluated       bool
	evalValueResult reflect.Value
	evalErrorResult error

	SerialNumber int
}

// Reset will mark this expression as not evaluated, thus next call to Evaluate will run normally.
func (exprAtm *ExpressionAtom) Reset() {
	exprAtm.evaluated = false
}

// Evaluate the object graph against underlined context or execute evaluation in the sub graph.
func (exprAtm *ExpressionAtom) Evaluate() (reflect.Value, error) {
	//logrus.Trace(exprAtm.Text)
	if exprAtm.evaluated {
		if len(exprAtm.Variable) > 0 {
			logrus.Tracef("Variable %s #%d is NOT FROM working memory", exprAtm.Text, exprAtm.SerialNumber)
		}
		if exprAtm.Constant != nil {
			logrus.Tracef("Constant %s #%d is NOT FROM working memory", exprAtm.Text, exprAtm.SerialNumber)
		}
		if exprAtm.FunctionCall != nil {
			logrus.Tracef("Function %s #%d is NOT FROM working memory", exprAtm.Text, exprAtm.SerialNumber)
		}
		if exprAtm.MethodCall != nil {
			logrus.Tracef("Method %s #%d is NOT FROM working memory", exprAtm.Text, exprAtm.SerialNumber)
		}
		return exprAtm.evalValueResult, exprAtm.evalErrorResult
	}
	exprAtm.evaluated = true
	//logrus.Tracef("ExpressionAtom : %s", exprAtm.Text)
	if len(exprAtm.Variable) > 0 {
		logrus.Tracef("Variable %s #%d is FROM working memory", exprAtm.Text, exprAtm.SerialNumber)
		exprAtm.evalValueResult, exprAtm.evalErrorResult = exprAtm.dataCtx.GetValue(exprAtm.Variable)
		return exprAtm.evalValueResult, exprAtm.evalErrorResult
	}
	if exprAtm.Constant != nil {
		logrus.Tracef("ExpressionAtom Constant FROM working memory: %s", exprAtm.Text)
		exprAtm.evalValueResult, exprAtm.evalErrorResult = exprAtm.Constant.Evaluate()
		return exprAtm.evalValueResult, exprAtm.evalErrorResult
	}
	if exprAtm.FunctionCall != nil {
		logrus.Tracef("ExpressionAtom Function FROM working memory: %s", exprAtm.Text)
		exprAtm.evalValueResult, exprAtm.evalErrorResult = exprAtm.FunctionCall.Evaluate()
		return exprAtm.evalValueResult, exprAtm.evalErrorResult
	}
	if exprAtm.MethodCall != nil {
		logrus.Tracef("MethodCall Function FROM working memory: %s", exprAtm.Text)
		exprAtm.evalValueResult, exprAtm.evalErrorResult = exprAtm.MethodCall.Evaluate()
		return exprAtm.evalValueResult, exprAtm.evalErrorResult
	}

	logrus.Tracef("ExpressionAtom MathOps : %s", exprAtm.Text)
	lv, err := exprAtm.ExpressionAtomLeft.Evaluate()
	if err != nil {
		exprAtm.evalValueResult, exprAtm.evalErrorResult = reflect.ValueOf(nil), errors.Trace(err)
		return exprAtm.evalValueResult, exprAtm.evalErrorResult
	}
	rv, err := exprAtm.ExpressionAtomRight.Evaluate()
	if err != nil {
		exprAtm.evalValueResult, exprAtm.evalErrorResult = reflect.ValueOf(nil), errors.Trace(err)
		return exprAtm.evalValueResult, exprAtm.evalErrorResult
	}
	switch exprAtm.MathOperator {
	case MathOperatorPlus:
		exprAtm.evalValueResult, exprAtm.evalErrorResult = pkg.ValueAdd(lv, rv)
		return exprAtm.evalValueResult, exprAtm.evalErrorResult
	case MathOperatorMinus:
		exprAtm.evalValueResult, exprAtm.evalErrorResult = pkg.ValueSub(lv, rv)
		return exprAtm.evalValueResult, exprAtm.evalErrorResult
	case MathOperatorMul:
		exprAtm.evalValueResult, exprAtm.evalErrorResult = pkg.ValueMul(lv, rv)
		return exprAtm.evalValueResult, exprAtm.evalErrorResult
	case MathOperatorDiv:
		exprAtm.evalValueResult, exprAtm.evalErrorResult = pkg.ValueDiv(lv, rv)
		return exprAtm.evalValueResult, exprAtm.evalErrorResult
	}
	exprAtm.evalValueResult, exprAtm.evalErrorResult = reflect.ValueOf(nil), errors.Errorf("math operation can only be applied to numerical data (eg. int, uit or float) or string")
	return exprAtm.evalValueResult, exprAtm.evalErrorResult
}

// Initialize will prepare this graph with  contexts
func (exprAtm *ExpressionAtom) Initialize(knowledgeContext *context.KnowledgeContext, ruleCtx *RuleContext, dataCtx *context.DataContext) {
	exprAtm.knowledgeContext = knowledgeContext
	exprAtm.ruleCtx = ruleCtx
	exprAtm.dataCtx = dataCtx

	if exprAtm.ExpressionAtomLeft != nil {
		exprAtm.ExpressionAtomLeft.Initialize(knowledgeContext, ruleCtx, dataCtx)
	}

	if exprAtm.ExpressionAtomRight != nil {
		exprAtm.ExpressionAtomRight.Initialize(knowledgeContext, ruleCtx, dataCtx)
	}

	if exprAtm.Constant != nil {
		exprAtm.Constant.Initialize(knowledgeContext, ruleCtx, dataCtx)
	}

	if exprAtm.FunctionCall != nil {
		exprAtm.FunctionCall.Initialize(knowledgeContext, ruleCtx, dataCtx)
	}

	if exprAtm.MethodCall != nil {
		exprAtm.MethodCall.Initialize(knowledgeContext, ruleCtx, dataCtx)
	}
}

// AcceptExpressionAtom will prepare this graph an expression atom. The first invocation to this function will set the
// left hand value, the second will set the right hand to be evaluated with math operator.
func (exprAtm *ExpressionAtom) AcceptExpressionAtom(exprAtom *ExpressionAtom) error {
	if exprAtm.ExpressionAtomLeft == nil {
		exprAtm.ExpressionAtomLeft = exprAtom
	} else if exprAtm.ExpressionAtomRight == nil {
		exprAtm.ExpressionAtomRight = exprAtom
	} else {
		return errors.Errorf("expression alredy set twice")
	}
	return nil
}

// AcceptFunctionCall will prepare this expression atom as a function call
func (exprAtm *ExpressionAtom) AcceptFunctionCall(funcCall *FunctionCall) error {
	if exprAtm.FunctionCall != nil {
		return errors.Errorf("functioncall alredy set")
	}
	exprAtm.FunctionCall = funcCall
	return nil
}

// AcceptMethodCall will prepare this expression atom as a method call.
func (exprAtm *ExpressionAtom) AcceptMethodCall(methodCall *MethodCall) error {
	if exprAtm.MethodCall != nil {
		return errors.Errorf("method call alredy set")
	}
	exprAtm.MethodCall = methodCall
	return nil
}

// AcceptVariable will prepare this expression atom as a variable.
func (exprAtm *ExpressionAtom) AcceptVariable(name string) error {
	if exprAtm.Variable == "" {
		exprAtm.Variable = name
		return nil
	}
	return errors.Errorf("variable already defined")
}

// AcceptConstant will prepare this expression as a constant.
func (exprAtm *ExpressionAtom) AcceptConstant(cons *Constant) error {
	if exprAtm.Constant == nil {
		exprAtm.Constant = cons
		return nil
	}
	return errors.Errorf("constant already defined")
}

// EqualsTo will compare two literal constants, be it string, int, uint, floats bools and nils
func (exprAtm *ExpressionAtom) EqualsTo(that AlphaNode) bool {
	typ := reflect.TypeOf(that)
	if that == nil {
		return false
	}
	if typ.Kind() == reflect.Ptr {
		if typ.Elem().Name() == "ExpressionAtom" {
			thatExprAtm := that.(*ExpressionAtom)
			if len(exprAtm.Variable) > 0 && exprAtm.Variable == thatExprAtm.Variable {
				return true
			}
			if exprAtm.Constant != nil && thatExprAtm.Constant != nil && exprAtm.Constant.EqualsTo(thatExprAtm.Constant) {
				return true
			}
			if exprAtm.FunctionCall != nil && thatExprAtm.FunctionCall != nil && exprAtm.FunctionCall.EqualsTo(thatExprAtm.FunctionCall) {
				return true
			}
			if exprAtm.MethodCall != nil && thatExprAtm.MethodCall != nil && exprAtm.MethodCall.EqualsTo(thatExprAtm.MethodCall) {
				return true
			}
			if exprAtm.MathOperator == thatExprAtm.MathOperator &&
				exprAtm.ExpressionAtomLeft != nil && thatExprAtm.ExpressionAtomLeft != nil &&
				exprAtm.ExpressionAtomRight != nil && thatExprAtm.ExpressionAtomRight != nil {
				switch exprAtm.MathOperator {
				case MathOperatorPlus, MathOperatorMul:
					if (exprAtm.ExpressionAtomLeft.EqualsTo(thatExprAtm.ExpressionAtomLeft) &&
						exprAtm.ExpressionAtomRight.EqualsTo(thatExprAtm.ExpressionAtomRight)) ||
						(exprAtm.ExpressionAtomLeft.EqualsTo(thatExprAtm.ExpressionAtomRight) &&
							exprAtm.ExpressionAtomRight.EqualsTo(thatExprAtm.ExpressionAtomLeft)) {
						return true
					}
				default:
					if exprAtm.ExpressionAtomLeft.EqualsTo(thatExprAtm.ExpressionAtomLeft) &&
						exprAtm.ExpressionAtomRight.EqualsTo(thatExprAtm.ExpressionAtomRight) {
						return true
					}
				}
			}
		}
	}
	return false
}

// IsContainVariable should check for this ExpressionAtom whether it contains a variable
func (exprAtm *ExpressionAtom) IsContainVariable(atm *ExpressionAtom, varName string) bool {
	if atm.Variable == varName {
		return true
	}
	if atm.FunctionCall != nil && atm.FunctionCall.FunctionArguments != nil {
		return atm.FunctionCall.IsContainVariable(varName)
	}
	if atm.MethodCall != nil && atm.MethodCall.MethodArguments != nil {
		return atm.MethodCall.IsContainVariable(varName)
	}
	if atm.ExpressionAtomLeft != nil {
		return atm.ExpressionAtomLeft.IsContainVariable(atm.ExpressionAtomLeft, varName)
	}
	if atm.ExpressionAtomRight != nil {
		return atm.ExpressionAtomRight.IsContainVariable(atm.ExpressionAtomRight, varName)
	}
	return false
}
