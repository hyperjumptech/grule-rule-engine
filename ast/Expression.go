package ast

import (
	"bytes"
	"errors"
	"fmt"
	"reflect"

	"github.com/google/uuid"
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
		AstID:    uuid.New().String(),
		Operator: 0,
	}
}

// Expression AST Graph node
type Expression struct {
	AstID         string
	GrlText       string
	DataContext   *DataContext
	WorkingMemory *WorkingMemory

	LeftExpression   *Expression
	RightExpression  *Expression
	SingleExpression *Expression
	ExpressionAtom   *ExpressionAtom
	Operator         int
	Value            reflect.Value

	Evaluated bool
}

// InitializeContext will initialize this AST graph with data context and working memory before running rule on them.
func (e *Expression) InitializeContext(dataCtx *DataContext, memory *WorkingMemory) {
	e.DataContext = dataCtx
	e.WorkingMemory = memory
	if e.LeftExpression != nil {
		e.LeftExpression.InitializeContext(dataCtx, memory)
	}
	if e.RightExpression != nil {
		e.RightExpression.InitializeContext(dataCtx, memory)
	}
	if e.SingleExpression != nil {
		e.SingleExpression.InitializeContext(dataCtx, memory)
	}
	if e.ExpressionAtom != nil {
		e.ExpressionAtom.InitializeContext(dataCtx, memory)
	}
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
	var buff bytes.Buffer
	if e.SingleExpression != nil {
		buff.WriteString(e.SingleExpression.GetSnapshot())
	}
	if e.LeftExpression != nil && e.RightExpression != nil {
		buff.WriteString("(")
		buff.WriteString(e.LeftExpression.GetSnapshot())
		switch e.Operator {
		case OpMul:
			buff.WriteString(" x ")
		case OpDiv:
			buff.WriteString(" / ")
		case OpMod:
			buff.WriteString(" % ")
		case OpAdd:
			buff.WriteString(" + ")
		case OpSub:
			buff.WriteString(" - ")
		case OpBitAnd:
			buff.WriteString(" & ")
		case OpBitOr:
			buff.WriteString(" | ")
		case OpGT:
			buff.WriteString(" > ")
		case OpLT:
			buff.WriteString(" < ")
		case OpGTE:
			buff.WriteString(" >= ")
		case OpLTE:
			buff.WriteString(" <= ")
		case OpEq:
			buff.WriteString(" == ")
		case OpNEq:
			buff.WriteString(" != ")
		case OpAnd:
			buff.WriteString(" && ")
		case OpOr:
			buff.WriteString(" || ")
		}
		buff.WriteString(e.RightExpression.GetSnapshot())
		buff.WriteString(")")
	}
	if e.ExpressionAtom != nil {
		buff.WriteString(e.ExpressionAtom.GetSnapshot())
	}
	return buff.String()
}

// SetGrlText set the expression syntax related to this graph when it was constructed. Only ANTLR4 listener should
// call this function.
func (e *Expression) SetGrlText(grlText string) {
	e.GrlText = grlText
}

// Evaluate will evaluate this AST graph for when scope evaluation
func (e *Expression) Evaluate() (reflect.Value, error) {
	if e.Evaluated == true {
		return e.Value, nil
	}
	if e.ExpressionAtom != nil {
		val, err := e.ExpressionAtom.Evaluate()
		if err == nil {
			e.Value = val
			e.Evaluated = true
		}
		return val, err
	}
	if e.SingleExpression != nil {
		val, err := e.SingleExpression.Evaluate()
		if err == nil {
			e.Value = val
			e.Evaluated = true
		}
		return val, err
	}
	if e.LeftExpression != nil && e.RightExpression != nil {
		lval, lerr := e.LeftExpression.Evaluate()
		rval, rerr := e.RightExpression.Evaluate()
		if lerr != nil {
			return reflect.ValueOf(nil), fmt.Errorf("left hand expression error. got %v", lerr)
		}
		if rerr != nil {
			return reflect.ValueOf(nil), fmt.Errorf("right hand expression error.  got %v", rerr)
		}

		var val reflect.Value
		var opErr error

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
		return val, opErr
	}
	return reflect.ValueOf(nil), nil
}
