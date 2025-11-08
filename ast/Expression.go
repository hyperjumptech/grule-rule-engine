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

package ast

import (
	"errors"
	"fmt"
	"reflect"
	"strings"

	"github.com/hyperjumptech/grule-rule-engine/ast/unique"

	"github.com/hyperjumptech/grule-rule-engine/pkg"
)

const (
	// OpMul Multiplication operator
	OpMul int = iota
	// OpDiv Divisioon operator
	OpDiv
	// OpMod Modulus operator
	OpMod
	// OpAdd Addition operator
	OpAdd
	// OpSub Substraction operator
	OpSub
	// OpBitAnd Bitwise And operator
	OpBitAnd
	// OpBitOr Bitwise Or operator
	OpBitOr
	// OpGT Greater Than operator
	OpGT
	// OpLT Lesser Than operator
	OpLT
	// OpGTE Greater Than or Equal operator
	OpGTE
	// OpLTE Lesser Than or Equal operator
	OpLTE
	// OpEq Equals operator
	OpEq
	// OpNEq Not Equals operator
	OpNEq
	// OpAnd Logical And operator
	OpAnd
	// OpOr Logical Or operator
	OpOr
)

// NewExpression creates new Expression instance
func NewExpression() *Expression {

	return &Expression{
		AstID:    unique.NewID(),
		Operator: 0,
	}
}

// Expression AST Graph node
type Expression struct {
	AstID   string
	GrlText string

	LeftExpression   *Expression
	RightExpression  *Expression
	SingleExpression *Expression
	ExpressionAtom   *ExpressionAtom
	Operator         int
	Negated          bool
	Value            reflect.Value

	Evaluated        bool
	CompareNilValues bool
}

// MakeCatalog will create a catalog entry from Expression node.
func (e *Expression) MakeCatalog(cat *Catalog) {
	meta := &ExpressionMeta{
		NodeMeta: NodeMeta{
			AstID:    e.AstID,
			GrlText:  e.GrlText,
			Snapshot: e.GetSnapshot(),
		},
	}
	if cat.AddMeta(e.AstID, meta) {
		if e.LeftExpression != nil {
			meta.LeftExpressionID = e.LeftExpression.AstID
			e.LeftExpression.MakeCatalog(cat)
		}
		if e.RightExpression != nil {
			meta.RightExpressionID = e.RightExpression.AstID
			e.RightExpression.MakeCatalog(cat)
		}
		if e.SingleExpression != nil {
			meta.SingleExpressionID = e.SingleExpression.AstID
			e.SingleExpression.MakeCatalog(cat)
		}
		if e.ExpressionAtom != nil {
			meta.ExpressionAtomID = e.ExpressionAtom.AstID
			e.ExpressionAtom.MakeCatalog(cat)
		}
		meta.Operator = e.Operator
		meta.Negated = e.Negated
	}
}

// Clone will clone this Expression. The new clone will have an identical structure
func (e *Expression) Clone(cloneTable *pkg.CloneTable) *Expression {
	clone := &Expression{
		AstID:    unique.NewID(),
		GrlText:  e.GrlText,
		Operator: e.Operator,
		Negated:  e.Negated,
	}

	if e.LeftExpression != nil {
		if cloneTable.IsCloned(e.LeftExpression.AstID) {
			clone.LeftExpression = cloneTable.Records[e.LeftExpression.AstID].CloneInstance.(*Expression)
		} else {
			cloned := e.LeftExpression.Clone(cloneTable)
			clone.LeftExpression = cloned
			cloneTable.MarkCloned(e.LeftExpression.AstID, cloned.AstID, e.LeftExpression, cloned)
		}
	}

	if e.RightExpression != nil {
		if cloneTable.IsCloned(e.RightExpression.AstID) {
			clone.RightExpression = cloneTable.Records[e.RightExpression.AstID].CloneInstance.(*Expression)
		} else {
			cloned := e.RightExpression.Clone(cloneTable)
			clone.RightExpression = cloned
			cloneTable.MarkCloned(e.RightExpression.AstID, cloned.AstID, e.RightExpression, cloned)
		}
	}

	if e.SingleExpression != nil {
		if cloneTable.IsCloned(e.SingleExpression.AstID) {
			clone.SingleExpression = cloneTable.Records[e.SingleExpression.AstID].CloneInstance.(*Expression)
		} else {
			cloned := e.SingleExpression.Clone(cloneTable)
			clone.SingleExpression = cloned
			clone.Negated = e.Negated
			cloneTable.MarkCloned(e.SingleExpression.AstID, cloned.AstID, e.SingleExpression, cloned)
		}
	}

	if e.ExpressionAtom != nil {
		if cloneTable.IsCloned(e.ExpressionAtom.AstID) {
			clone.ExpressionAtom = cloneTable.Records[e.ExpressionAtom.AstID].CloneInstance.(*ExpressionAtom)
		} else {
			cloned := e.ExpressionAtom.Clone(cloneTable)
			clone.ExpressionAtom = cloned
			cloneTable.MarkCloned(e.ExpressionAtom.AstID, cloned.AstID, e.ExpressionAtom, cloned)
		}
	}

	return clone
}

// AcceptExpression will accept an Expression AST graph into this ast graph
func (e *Expression) AcceptExpression(exp *Expression) error {
	if e.SingleExpression == nil && e.LeftExpression == nil {
		e.SingleExpression = exp
	} else if e.SingleExpression != nil && e.LeftExpression == nil {
		e.LeftExpression = e.SingleExpression
		e.RightExpression = exp
		e.SingleExpression = nil
	} else {

		return errors.New("left or right side expression already assigned")
	}

	return nil
}

// ExpressionReceiver contains function to be implemented by other AST graph to receive an Expression AST graph
type ExpressionReceiver interface {
	AcceptExpression(exp *Expression) error
}

// AcceptExpressionAtom will accept ExpressionAtom into this Expression
func (e *Expression) AcceptExpressionAtom(atom *ExpressionAtom) error {
	e.ExpressionAtom = atom

	return nil
}

// GetAstID get the UUID asigned for this AST graph node
func (e *Expression) GetAstID() string {

	return e.AstID
}

// GetGrlText get the expression syntax related to this graph when it wast constructed
func (e *Expression) GetGrlText() string {

	return e.GrlText
}

// GetSnapshot will create a structure signature or AST graph
func (e *Expression) GetSnapshot() string {
	var buff strings.Builder
	buff.WriteString(EXPRESSION)
	buff.WriteString("(")
	if e.SingleExpression != nil {
		buff.WriteString("SE(")
		if e.Negated {
			buff.WriteString("!")
		}
		buff.WriteString(e.SingleExpression.GetSnapshot())
		buff.WriteString(")")
	}
	if e.LeftExpression != nil && e.RightExpression != nil {
		buff.WriteString("EL(")
		buff.WriteString(e.LeftExpression.GetSnapshot())
		buff.WriteString(")")

		switch e.Operator {
		case OpMul:
			buff.WriteString("*")
		case OpDiv:
			buff.WriteString("/")
		case OpMod:
			buff.WriteString("%")
		case OpAdd:
			buff.WriteString("+")
		case OpSub:
			buff.WriteString("-")
		case OpBitAnd:
			buff.WriteString("&")
		case OpBitOr:
			buff.WriteString("|")
		case OpGT:
			buff.WriteString(">")
		case OpLT:
			buff.WriteString("<")
		case OpGTE:
			buff.WriteString(">=")
		case OpLTE:
			buff.WriteString("<=")
		case OpEq:
			buff.WriteString("==")
		case OpNEq:
			buff.WriteString("!=")
		case OpAnd:
			buff.WriteString("&&")
		case OpOr:
			buff.WriteString("||")
		}

		buff.WriteString("ER(")
		buff.WriteString(e.RightExpression.GetSnapshot())
		buff.WriteString(")")
	}
	if e.ExpressionAtom != nil {
		buff.WriteString("EA(")
		buff.WriteString(e.ExpressionAtom.GetSnapshot())
		buff.WriteString(")")
	}
	buff.WriteString(")")

	return buff.String()
}

// SetGrlText set the expression syntax related to this graph when it was constructed. Only ANTLR4 listener should
// call this function.
func (e *Expression) SetGrlText(grlText string) {
	e.GrlText = grlText
}

// Evaluate will evaluate this AST graph for when scope evaluation
func (e *Expression) Evaluate(dataContext IDataContext, memory *WorkingMemory) (reflect.Value, error) {
	compareNode := dataContext.Get("COMPARE_NILS")
	if compareNode != nil {
		if rv := compareNode.Value(); rv.IsValid() && rv.Kind() == reflect.Bool {
			e.CompareNilValues = rv.Bool()
		} else {
			e.CompareNilValues = false
		}
	} else {
		e.CompareNilValues = false
	}

	if e.Evaluated {

		return e.Value, nil
	}
	if e.ExpressionAtom != nil {
		val, err := e.ExpressionAtom.Evaluate(dataContext, memory)
		if err == nil {
			e.Value = val
			e.Evaluated = true
		}

		return val, err
	}
	if e.SingleExpression != nil {
		val, err := e.SingleExpression.Evaluate(dataContext, memory)
		if err == nil {
			e.Value = val
			if e.Negated {
				if e.Value.Kind() == reflect.Bool {
					e.Value = reflect.ValueOf(!e.Value.Bool())
				} else {
					AstLog.Warnf("Expression \"%s\" is a negation to non boolean value, negation is ignored.", e.SingleExpression.GrlText)
				}
			}
			e.Evaluated = true
		}

		return e.Value, err
	}
	if e.LeftExpression != nil && e.RightExpression != nil {
		var val reflect.Value
		var opErr error

		lval, lerr := e.LeftExpression.Evaluate(dataContext, memory)
		if e.Operator == OpAnd {
			if lerr != nil {

				return reflect.Value{}, fmt.Errorf("left hand expression error. got %v", lerr)
			}
			val, opErr = pkg.EvaluateLogicSingle(lval)
			if opErr == nil && !val.Bool() {
				e.Value = val
				e.Evaluated = true

				return val, opErr
			}
		}
		if e.Operator == OpOr {
			if lerr != nil {

				return reflect.Value{}, fmt.Errorf("left hand expression error. got %v", lerr)
			}
			val, opErr = pkg.EvaluateLogicSingle(lval)
			if opErr == nil && val.Bool() {
				e.Value = val
				e.Evaluated = true

				return val, opErr
			}
		}

		rval, rerr := e.RightExpression.Evaluate(dataContext, memory)
		if lerr != nil {

			return reflect.Value{}, fmt.Errorf("left hand expression error. got %v", lerr)
		}
		if rerr != nil {

			return reflect.Value{}, fmt.Errorf("right hand expression error.  got %v", rerr)
		}

		if e.CompareNilValues && (!lval.IsValid() || !rval.IsValid()) {
			if e.CompareNilValues {
				AstLog.Debugf("Values have invalid value (%v and %v) but continuing with null handling", lval, rval)
				switch e.Operator {
				case OpMul, OpDiv, OpBitAnd, OpBitOr, OpMod:
					// Can be left as nil, as these operators with Nil are not defined
					e.Evaluated = true
				case OpAdd:
					if lval.IsValid() {
						val = lval
					} else if rval.IsValid() {
						val = rval
					}
					e.Evaluated = true
				case OpSub:
					if lval.IsValid() {
						val = lval
					} else if rval.IsValid() {
						val, _ = pkg.EvaluateSubtraction(reflect.ValueOf(0), rval)
					}
					e.Evaluated = true
				case OpOr:
					lvale := pkg.GetValueElem(lval)
					rvale := pkg.GetValueElem(rval)
					if (lvale.IsValid() && lvale.Kind() == reflect.Bool && lvale.Bool()) || (rvale.IsValid() && rvale.Kind() == reflect.Bool && rvale.Bool()) {
						val = reflect.ValueOf(true)
					} else {
						val = reflect.ValueOf(false)
					}
					e.Value = val
					e.Evaluated = true
				case OpEq, OpGTE, OpLTE, OpGT, OpLT, OpAnd:
					val = reflect.ValueOf(false)
					e.Value = val
					e.Evaluated = true
				case OpNEq:
					val = reflect.ValueOf(true)
					e.Value = val
					e.Evaluated = true
				}
			}
		} else {
			switch e.Operator {
			case OpMul:
				val, opErr = pkg.EvaluateMultiplication(lval, rval)
			case OpDiv:
				val, opErr = pkg.EvaluateDivision(lval, rval)
			case OpMod:
				val, opErr = pkg.EvaluateModulo(lval, rval)
			case OpAdd:
				val, opErr = pkg.EvaluateAddition(lval, rval)
			case OpSub:
				val, opErr = pkg.EvaluateSubtraction(lval, rval)
			case OpBitAnd:
				val, opErr = pkg.EvaluateBitAnd(lval, rval)
			case OpBitOr:
				val, opErr = pkg.EvaluateBitOr(lval, rval)
			case OpGT:
				val, opErr = pkg.EvaluateGreaterThan(lval, rval)
			case OpLT:
				val, opErr = pkg.EvaluateLesserThan(lval, rval)
			case OpGTE:
				val, opErr = pkg.EvaluateGreaterThanEqual(lval, rval)
			case OpLTE:
				val, opErr = pkg.EvaluateLesserThanEqual(lval, rval)
			case OpEq:
				val, opErr = pkg.EvaluateEqual(lval, rval)
			case OpNEq:
				val, opErr = pkg.EvaluateNotEqual(lval, rval)
			case OpAnd:
				val, opErr = pkg.EvaluateLogicAnd(lval, rval)
			case OpOr:
				val, opErr = pkg.EvaluateLogicOr(lval, rval)
			}
			if opErr == nil {
				e.Value = val
				e.Evaluated = true
			}
		}

		return val, opErr
	}

	return reflect.Value{}, nil
}
