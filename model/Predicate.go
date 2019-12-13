package model

import (
	"github.com/hyperjumptech/grule-rule-engine/context"
	"github.com/hyperjumptech/grule-rule-engine/pkg"
	"github.com/juju/errors"
	"reflect"
	"time"
)

const (
	// TimeTypeString store the value of Time type.
	TimeTypeString = "time.Time"
)

// Predicate holds the left and right Expression Atom graph. And apply comparisson operator from both
// expression atom result.
type Predicate struct {
	ExpressionAtomLeft  *ExpressionAtom
	ExpressionAtomRight *ExpressionAtom
	ComparisonOperator  ComparisonOperator
	knowledgeContext    *context.KnowledgeContext
	ruleCtx             *context.RuleContext
	dataCtx             *context.DataContext
}

// Initialize initialize this graph with context
func (prdct *Predicate) Initialize(knowledgeContext *context.KnowledgeContext, ruleCtx *context.RuleContext, dataCtx *context.DataContext) {
	prdct.knowledgeContext = knowledgeContext
	prdct.ruleCtx = ruleCtx
	prdct.dataCtx = dataCtx

	if prdct.ExpressionAtomLeft != nil {
		prdct.ExpressionAtomLeft.Initialize(knowledgeContext, ruleCtx, dataCtx)
	}
	if prdct.ExpressionAtomRight != nil {
		prdct.ExpressionAtomRight.Initialize(knowledgeContext, ruleCtx, dataCtx)
	}
}

// AcceptExpressionAtom configure this graph with left and right side of expression atom. The first call
// to this function will set the left hand side and the second call will set the right.
func (prdct *Predicate) AcceptExpressionAtom(exprAtom *ExpressionAtom) error {
	if prdct.ExpressionAtomLeft == nil {
		prdct.ExpressionAtomLeft = exprAtom
	} else if prdct.ExpressionAtomRight == nil {
		prdct.ExpressionAtomRight = exprAtom
	} else {
		return errors.Errorf("expression alredy set twice")
	}
	return nil
}

// Evaluate the object graph against underlined context or execute evaluation in the sub graph.
func (prdct *Predicate) Evaluate() (reflect.Value, error) {
	if prdct.ExpressionAtomRight == nil {
		return prdct.ExpressionAtomLeft.Evaluate()
	}
	lv, err := prdct.ExpressionAtomLeft.Evaluate()
	if err != nil {
		return reflect.ValueOf(nil), errors.Trace(err)
	}
	rv, err := prdct.ExpressionAtomRight.Evaluate()
	if err != nil {
		return reflect.ValueOf(nil), errors.Trace(err)
	}
	if lv.Kind() == rv.Kind() && (prdct.ComparisonOperator == ComparisonOperatorEQ || prdct.ComparisonOperator == ComparisonOperatorNEQ) {
		if prdct.ComparisonOperator == ComparisonOperatorEQ {
			switch lv.Kind() {
			case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
				return reflect.ValueOf(lv.Int() == rv.Int()), nil
			case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
				return reflect.ValueOf(lv.Uint() == rv.Uint()), nil
			case reflect.Float64, reflect.Float32:
				return reflect.ValueOf(lv.Float() == rv.Float()), nil
			case reflect.String:
				return reflect.ValueOf(lv.String() == rv.String()), nil
			case reflect.Bool:
				return reflect.ValueOf(lv.Bool() == rv.Bool()), nil
			}
			if lv.String() == TimeTypeString {
				tl := pkg.ValueToInterface(lv).(time.Time)
				tr := pkg.ValueToInterface(rv).(time.Time)
				return reflect.ValueOf(tl.Equal(tr)), nil
			}
		} else {
			switch lv.Kind() {
			case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
				return reflect.ValueOf(lv.Int() != rv.Int()), nil
			case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
				return reflect.ValueOf(lv.Uint() != rv.Uint()), nil
			case reflect.Float64, reflect.Float32:
				return reflect.ValueOf(lv.Float() != rv.Float()), nil
			case reflect.String:
				return reflect.ValueOf(lv.String() != rv.String()), nil
			case reflect.Bool:
				return reflect.ValueOf(lv.Bool() != rv.Bool()), nil
			}
			if lv.String() == TimeTypeString {
				tl := pkg.ValueToInterface(lv).(time.Time)
				tr := pkg.ValueToInterface(rv).(time.Time)
				return reflect.ValueOf(!tl.Equal(tr)), nil
			}
		}
	} else if lv.Type().String() == TimeTypeString && rv.Type().String() == TimeTypeString {
		tl := pkg.ValueToInterface(lv).(time.Time)
		tr := pkg.ValueToInterface(rv).(time.Time)
		switch prdct.ComparisonOperator {
		case ComparisonOperatorEQ:
			return reflect.ValueOf(tl.Equal(tr)), nil
		case ComparisonOperatorNEQ:
			return reflect.ValueOf(!tl.Equal(tr)), nil
		case ComparisonOperatorGT:
			return reflect.ValueOf(tl.After(tr)), nil
		case ComparisonOperatorGTE:
			return reflect.ValueOf(tl.After(tr) || tl.Equal(tr)), nil
		case ComparisonOperatorLT:
			return reflect.ValueOf(tl.Before(tr)), nil
		case ComparisonOperatorLTE:
			return reflect.ValueOf(tl.Before(tr) || tl.Equal(tr)), nil
		}
	} else {
		var lf, rf float64
		switch pkg.GetBaseKind(lv) {
		case reflect.Int64:
			lf = float64(lv.Int())
		case reflect.Uint64:
			lf = float64(lv.Uint())
		case reflect.Float64:
			lf = lv.Float()
		default:
			return reflect.ValueOf(nil), errors.Errorf("comparison operator can only between strings, time or numbers")
		}
		switch pkg.GetBaseKind(rv) {
		case reflect.Int64:
			rf = float64(rv.Int())
		case reflect.Uint64:
			rf = float64(rv.Uint())
		case reflect.Float64:
			rf = rv.Float()
		default:
			return reflect.ValueOf(nil), errors.Errorf("comparison operator can only between strings, time or numbers")
		}
		switch prdct.ComparisonOperator {
		case ComparisonOperatorEQ:
			return reflect.ValueOf(lf == rf), nil
		case ComparisonOperatorNEQ:
			return reflect.ValueOf(lf != rf), nil
		case ComparisonOperatorGT:
			return reflect.ValueOf(lf > rf), nil
		case ComparisonOperatorGTE:
			return reflect.ValueOf(lf >= rf), nil
		case ComparisonOperatorLT:
			return reflect.ValueOf(lf < rf), nil
		case ComparisonOperatorLTE:
			return reflect.ValueOf(lf <= rf), nil
		}
	}
	return reflect.ValueOf(nil), nil
}
