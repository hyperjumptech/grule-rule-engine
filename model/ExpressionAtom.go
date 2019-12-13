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
	ruleCtx             *context.RuleContext
	dataCtx             *context.DataContext
}

// Evaluate the object graph against underlined context or execute evaluation in the sub graph.
func (exprAtm *ExpressionAtom) Evaluate() (reflect.Value, error) {
	//logrus.Tracef("ExpressionAtom : %s", exprAtm.Text)
	if len(exprAtm.Variable) > 0 {
		logrus.Tracef("ExpressionAtom Variable : %s", exprAtm.Text)
		return exprAtm.dataCtx.GetValue(exprAtm.Variable)
	} else if exprAtm.Constant != nil {
		logrus.Tracef("ExpressionAtom Constant : %s", exprAtm.Text)
		return exprAtm.Constant.Evaluate()
	} else if exprAtm.FunctionCall != nil {
		logrus.Tracef("ExpressionAtom Function : %s", exprAtm.Text)
		return exprAtm.FunctionCall.Evaluate()
	} else if exprAtm.MethodCall != nil {
		logrus.Tracef("MethodCall Function : %s", exprAtm.Text)
		return exprAtm.MethodCall.Evaluate()
	} else {
		logrus.Tracef("ExpressionAtom MathOps : %s", exprAtm.Text)
		lv, err := exprAtm.ExpressionAtomLeft.Evaluate()
		if err != nil {
			return reflect.ValueOf(nil), errors.Trace(err)
		}
		rv, err := exprAtm.ExpressionAtomRight.Evaluate()
		if err != nil {
			return reflect.ValueOf(nil), errors.Trace(err)
		}
		switch exprAtm.MathOperator {
		case MathOperatorPlus:
			return pkg.ValueAdd(lv, rv)
		case MathOperatorMinus:
			return pkg.ValueSub(lv, rv)
		case MathOperatorMul:
			return pkg.ValueMul(lv, rv)
		case MathOperatorDiv:
			return pkg.ValueDiv(lv, rv)
		}
		return reflect.ValueOf(nil), errors.Errorf("math operation can only be applied to numerical data (eg. int, uit or float) or string")
	}
}

// Initialize will prepare this graph with  contexts
func (exprAtm *ExpressionAtom) Initialize(knowledgeContext *context.KnowledgeContext, ruleCtx *context.RuleContext, dataCtx *context.DataContext) {
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
