package ast

import (
	"bytes"
	"errors"
	"fmt"
	"github.com/google/uuid"
	"reflect"
	"time"
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
		if lerr != nil || rerr != nil {
			return reflect.ValueOf(nil), fmt.Errorf("left hand or right hand expression error. left got %v and/or right got %v", lerr, rerr)
		}

		var val reflect.Value
		var opErr error

		switch e.Operator {
		case OpMul:
			val, opErr = EvaluateMultiplication(lval, rval)
		case OpDiv:
			val, opErr = EvaluateDivision(lval, rval)
		case OpMod:
			val, opErr = EvaluateModulo(lval, rval)
		case OpAdd:
			val, opErr = EvaluateAddition(lval, rval)
		case OpSub:
			val, opErr = EvaluateSubstraction(lval, rval)
		case OpBitAnd:
			val, opErr = EvaluateBitAnd(lval, rval)
		case OpBitOr:
			val, opErr = EvaluateBitOr(lval, rval)
		case OpGT:
			val, opErr = EvaluateGreaterThan(lval, rval)
		case OpLT:
			val, opErr = EvaluateLesserThan(lval, rval)
		case OpGTE:
			val, opErr = EvaluateGreaterThanEqual(lval, rval)
		case OpLTE:
			val, opErr = EvaluateLesserThanEqual(lval, rval)
		case OpEq:
			val, opErr = EvaluateEqual(lval, rval)
		case OpNEq:
			val, opErr = EvaluateNotEqual(lval, rval)
		case OpAnd:
			val, opErr = EvaluateLogicAnd(lval, rval)
		case OpOr:
			val, opErr = EvaluateLogicOr(lval, rval)
		}
		if opErr == nil {
			e.Value = val
			e.Evaluated = true
		}
		return val, opErr
	}
	return reflect.ValueOf(nil), nil
}

// EvaluateMultiplication will evaluate multiplication operation over two value
func EvaluateMultiplication(left, right reflect.Value) (reflect.Value, error) {
	switch left.Kind() {
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		lv := left.Int()
		switch right.Kind() {
		case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
			rv := right.Int()
			return reflect.ValueOf(lv * rv), nil
		case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
			rv := right.Uint()
			return reflect.ValueOf(lv * int64(rv)), nil
		case reflect.Float32, reflect.Float64:
			rv := right.Float()
			return reflect.ValueOf(float64(lv) * rv), nil
		default:
			return reflect.ValueOf(nil), fmt.Errorf("can not multipy data type of %s", right.Kind().String())
		}
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		lv := left.Uint()
		switch right.Kind() {
		case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
			rv := right.Int()
			return reflect.ValueOf(int64(lv) * rv), nil
		case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
			rv := right.Uint()
			return reflect.ValueOf(lv * rv), nil
		case reflect.Float32, reflect.Float64:
			rv := right.Float()
			return reflect.ValueOf(float64(lv) * rv), nil
		default:
			return reflect.ValueOf(nil), fmt.Errorf("can not multipy data type of %s", right.Kind().String())
		}
	case reflect.Float32, reflect.Float64:
		lv := left.Float()
		switch right.Kind() {
		case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
			rv := right.Int()
			return reflect.ValueOf(lv * float64(rv)), nil
		case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
			rv := right.Uint()
			return reflect.ValueOf(lv * float64(rv)), nil
		case reflect.Float32, reflect.Float64:
			rv := right.Float()
			return reflect.ValueOf(lv * rv), nil
		default:
			return reflect.ValueOf(nil), fmt.Errorf("can not multipy data type of %s", right.Kind().String())
		}
	default:
		return reflect.ValueOf(nil), fmt.Errorf("can not multipy data type of %s", left.Kind().String())
	}
}

// EvaluateDivision will evaluate division operation over two value
func EvaluateDivision(left, right reflect.Value) (reflect.Value, error) {
	switch left.Kind() {
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		lv := left.Int()
		switch right.Kind() {
		case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
			rv := right.Int()
			return reflect.ValueOf(float64(lv) / float64(rv)), nil
		case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
			rv := right.Uint()
			return reflect.ValueOf(float64(lv) / float64(rv)), nil
		case reflect.Float32, reflect.Float64:
			rv := right.Float()
			return reflect.ValueOf(float64(lv) / rv), nil
		default:
			return reflect.ValueOf(nil), fmt.Errorf("can not use data type of %s in division", right.Kind().String())
		}
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		lv := left.Uint()
		switch right.Kind() {
		case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
			rv := right.Int()
			return reflect.ValueOf(float64(lv) / float64(rv)), nil
		case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
			rv := right.Uint()
			return reflect.ValueOf(float64(lv) / float64(rv)), nil
		case reflect.Float32, reflect.Float64:
			rv := right.Float()
			return reflect.ValueOf(float64(lv) / rv), nil
		default:
			return reflect.ValueOf(nil), fmt.Errorf("can not use data type of %s in division", right.Kind().String())
		}
	case reflect.Float32, reflect.Float64:
		lv := left.Float()
		switch right.Kind() {
		case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
			rv := right.Int()
			return reflect.ValueOf(lv / float64(rv)), nil
		case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
			rv := right.Uint()
			return reflect.ValueOf(lv / float64(rv)), nil
		case reflect.Float32, reflect.Float64:
			rv := right.Float()
			return reflect.ValueOf(lv / rv), nil
		default:
			return reflect.ValueOf(nil), fmt.Errorf("can not use data type of %s in division", right.Kind().String())
		}
	default:
		return reflect.ValueOf(nil), fmt.Errorf("can not use data type of %s in division", left.Kind().String())
	}
}

// EvaluateModulo will evaluate modulo operation over two value
func EvaluateModulo(left, right reflect.Value) (reflect.Value, error) {
	switch left.Kind() {
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		lv := left.Int()
		switch right.Kind() {
		case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
			rv := right.Int()
			return reflect.ValueOf(lv % rv), nil
		case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
			rv := right.Uint()
			return reflect.ValueOf(lv % int64(rv)), nil
		default:
			return reflect.ValueOf(nil), fmt.Errorf("can not use data type of %s in modulo", right.Kind().String())
		}
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		lv := left.Uint()
		switch right.Kind() {
		case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
			rv := right.Int()
			return reflect.ValueOf(int64(lv) % rv), nil
		case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
			rv := right.Uint()
			return reflect.ValueOf(int64(lv) % int64(rv)), nil
		default:
			return reflect.ValueOf(nil), fmt.Errorf("can not use data type of %s in modulo", right.Kind().String())
		}
	default:
		return reflect.ValueOf(nil), fmt.Errorf("can not use data type of %s in modulo", left.Kind().String())
	}
}

// EvaluateAddition will evaluate addition operation over two value
func EvaluateAddition(left, right reflect.Value) (reflect.Value, error) {
	switch left.Kind() {
	case reflect.String:
		lv := left.String()
		switch right.Kind() {
		case reflect.String:
			rv := right.String()
			return reflect.ValueOf(fmt.Sprintf("%s%s", lv, rv)), nil
		case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
			rv := right.Int()
			return reflect.ValueOf(fmt.Sprintf("%s%d", lv, rv)), nil
		case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
			rv := right.Uint()
			return reflect.ValueOf(fmt.Sprintf("%s%d", lv, rv)), nil
		case reflect.Float32, reflect.Float64:
			rv := right.Float()
			return reflect.ValueOf(fmt.Sprintf("%s%f", lv, rv)), nil
		case reflect.Bool:
			rv := right.Bool()
			return reflect.ValueOf(fmt.Sprintf("%s%v", lv, rv)), nil
		default:
			if right.Type().String() == "time.Time" {
				rv := right.Interface().(time.Time)
				return reflect.ValueOf(fmt.Sprintf("%s%s", lv, rv.Format(time.RFC3339))), nil
			}
			return reflect.ValueOf(nil), fmt.Errorf("can not use data type of %s in addition", right.Kind().String())
		}
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		lv := left.Int()
		switch right.Kind() {
		case reflect.String:
			rv := right.String()
			return reflect.ValueOf(fmt.Sprintf("%d%s", lv, rv)), nil
		case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
			rv := right.Int()
			return reflect.ValueOf(lv + rv), nil
		case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
			rv := right.Uint()
			return reflect.ValueOf(lv + int64(rv)), nil
		case reflect.Float32, reflect.Float64:
			rv := right.Float()
			return reflect.ValueOf(float64(lv) + rv), nil
		default:
			return reflect.ValueOf(nil), fmt.Errorf("can not use data type of %s in addition", right.Kind().String())
		}
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		lv := left.Uint()
		switch right.Kind() {
		case reflect.String:
			rv := right.String()
			return reflect.ValueOf(fmt.Sprintf("%d%s", lv, rv)), nil
		case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
			rv := right.Int()
			return reflect.ValueOf(int64(lv) + rv), nil
		case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
			rv := right.Uint()
			return reflect.ValueOf(lv + rv), nil
		case reflect.Float32, reflect.Float64:
			rv := right.Float()
			return reflect.ValueOf(float64(lv) + rv), nil
		default:
			return reflect.ValueOf(nil), fmt.Errorf("can not use data type of %s in division", right.Kind().String())
		}
	case reflect.Float32, reflect.Float64:
		lv := left.Float()
		switch right.Kind() {
		case reflect.String:
			rv := right.String()
			return reflect.ValueOf(fmt.Sprintf("%f%s", lv, rv)), nil
		case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
			rv := right.Int()
			return reflect.ValueOf(lv + float64(rv)), nil
		case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
			rv := right.Uint()
			return reflect.ValueOf(lv + float64(rv)), nil
		case reflect.Float32, reflect.Float64:
			rv := right.Float()
			return reflect.ValueOf(lv + rv), nil
		default:
			return reflect.ValueOf(nil), fmt.Errorf("can not use data type of %s in division", right.Kind().String())
		}
	default:
		return reflect.ValueOf(nil), fmt.Errorf("can not use data type of %s in division", left.Kind().String())
	}
}

// EvaluateSubstraction will evaluate substraction operation over two value
func EvaluateSubstraction(left, right reflect.Value) (reflect.Value, error) {
	switch left.Kind() {
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		lv := left.Int()
		switch right.Kind() {
		case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
			rv := right.Int()
			return reflect.ValueOf(lv - rv), nil
		case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
			rv := right.Uint()
			return reflect.ValueOf(lv - int64(rv)), nil
		case reflect.Float32, reflect.Float64:
			rv := right.Float()
			return reflect.ValueOf(float64(lv) - rv), nil
		default:
			return reflect.ValueOf(nil), fmt.Errorf("can not use data type of %s in substraction", right.Kind().String())
		}
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		lv := left.Uint()
		switch right.Kind() {
		case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
			rv := right.Int()
			return reflect.ValueOf(int64(lv) - rv), nil
		case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
			rv := right.Uint()
			return reflect.ValueOf(lv - rv), nil
		case reflect.Float32, reflect.Float64:
			rv := right.Float()
			return reflect.ValueOf(float64(lv) - rv), nil
		default:
			return reflect.ValueOf(nil), fmt.Errorf("can not use data type of %s in substraction", right.Kind().String())
		}
	case reflect.Float32, reflect.Float64:
		lv := left.Float()
		switch right.Kind() {
		case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
			rv := right.Int()
			return reflect.ValueOf(lv - float64(rv)), nil
		case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
			rv := right.Uint()
			return reflect.ValueOf(lv - float64(rv)), nil
		case reflect.Float32, reflect.Float64:
			rv := right.Float()
			return reflect.ValueOf(lv - rv), nil
		default:
			return reflect.ValueOf(nil), fmt.Errorf("can not use data type of %s in substraction", right.Kind().String())
		}
	default:
		return reflect.ValueOf(nil), fmt.Errorf("can not use data type of %s in substraction", left.Kind().String())
	}
}

// EvaluateBitAnd will evaluate Bitwise And operation over two value
func EvaluateBitAnd(left, right reflect.Value) (reflect.Value, error) {
	switch left.Kind() {
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		lv := left.Int()
		switch right.Kind() {
		case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
			rv := right.Int()
			return reflect.ValueOf(lv & rv), nil
		case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
			rv := right.Uint()
			return reflect.ValueOf(lv & int64(rv)), nil
		default:
			return reflect.ValueOf(nil), fmt.Errorf("can not use data type of %s in bitwise AND operation", right.Kind().String())
		}
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		lv := left.Uint()
		switch right.Kind() {
		case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
			rv := right.Int()
			return reflect.ValueOf(int64(lv) & rv), nil
		case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
			rv := right.Uint()
			return reflect.ValueOf(lv & rv), nil
		default:
			return reflect.ValueOf(nil), fmt.Errorf("can not use data type of %s in bitwise AND operation", right.Kind().String())
		}
	default:
		return reflect.ValueOf(nil), fmt.Errorf("can not use data type of %s in bitwise AND operation", left.Kind().String())
	}
}

// EvaluateBitOr will evaluate Bitwise Or operation over two value
func EvaluateBitOr(left, right reflect.Value) (reflect.Value, error) {
	switch left.Kind() {
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		lv := left.Int()
		switch right.Kind() {
		case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
			rv := right.Int()
			return reflect.ValueOf(lv | rv), nil
		case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
			rv := right.Uint()
			return reflect.ValueOf(lv | int64(rv)), nil
		default:
			return reflect.ValueOf(nil), fmt.Errorf("can not use data type of %s in bitwise OR operation", right.Kind().String())
		}
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		lv := left.Uint()
		switch right.Kind() {
		case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
			rv := right.Int()
			return reflect.ValueOf(int64(lv) | rv), nil
		case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
			rv := right.Uint()
			return reflect.ValueOf(lv | rv), nil
		default:
			return reflect.ValueOf(nil), fmt.Errorf("can not use data type of %s in bitwise OR operation", right.Kind().String())
		}
	default:
		return reflect.ValueOf(nil), fmt.Errorf("can not use data type of %s in bitwise OR operation", left.Kind().String())
	}
}

// EvaluateGreaterThan will evaluate GreaterThan operation over two value
func EvaluateGreaterThan(left, right reflect.Value) (reflect.Value, error) {
	switch left.Kind() {
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		lv := left.Int()
		switch right.Kind() {
		case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
			rv := right.Int()
			return reflect.ValueOf(lv > rv), nil
		case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
			rv := right.Uint()
			return reflect.ValueOf(lv > int64(rv)), nil
		case reflect.Float32, reflect.Float64:
			rv := right.Float()
			return reflect.ValueOf(float64(lv) > rv), nil
		default:
			return reflect.ValueOf(nil), fmt.Errorf("can not use data type of %s in GT comparison", right.Kind().String())
		}
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		lv := left.Uint()
		switch right.Kind() {
		case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
			rv := right.Int()
			return reflect.ValueOf(int64(lv) > rv), nil
		case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
			rv := right.Uint()
			return reflect.ValueOf(lv > rv), nil
		case reflect.Float32, reflect.Float64:
			rv := right.Float()
			return reflect.ValueOf(float64(lv) > rv), nil
		default:
			return reflect.ValueOf(nil), fmt.Errorf("can not use data type of %s in GT comparison", right.Kind().String())
		}
	case reflect.Float32, reflect.Float64:
		lv := left.Float()
		switch right.Kind() {
		case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
			rv := right.Int()
			return reflect.ValueOf(lv > float64(rv)), nil
		case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
			rv := right.Uint()
			return reflect.ValueOf(lv > float64(rv)), nil
		case reflect.Float32, reflect.Float64:
			rv := right.Float()
			return reflect.ValueOf(lv > rv), nil
		default:
			return reflect.ValueOf(nil), fmt.Errorf("can not use data type of %s in GT comparison", right.Kind().String())
		}
	default:
		if left.Type().String() == "time.Time" && right.Type().String() == "time.Time" {
			lv := left.Interface().(time.Time)
			rv := right.Interface().(time.Time)
			return reflect.ValueOf(lv.After(rv)), nil
		}
		return reflect.ValueOf(nil), fmt.Errorf("can not use data type of %s in GT comparison", left.Kind().String())
	}
}

// EvaluateLesserThan will evaluate LesserThan operation over two value
func EvaluateLesserThan(left, right reflect.Value) (reflect.Value, error) {
	switch left.Kind() {
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		lv := left.Int()
		switch right.Kind() {
		case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
			rv := right.Int()
			return reflect.ValueOf(lv < rv), nil
		case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
			rv := right.Uint()
			return reflect.ValueOf(lv < int64(rv)), nil
		case reflect.Float32, reflect.Float64:
			rv := right.Float()
			return reflect.ValueOf(float64(lv) < rv), nil
		default:
			return reflect.ValueOf(nil), fmt.Errorf("can not use data type of %s in LT comparison", right.Kind().String())
		}
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		lv := left.Uint()
		switch right.Kind() {
		case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
			rv := right.Int()
			return reflect.ValueOf(int64(lv) < rv), nil
		case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
			rv := right.Uint()
			return reflect.ValueOf(lv < rv), nil
		case reflect.Float32, reflect.Float64:
			rv := right.Float()
			return reflect.ValueOf(float64(lv) < rv), nil
		default:
			return reflect.ValueOf(nil), fmt.Errorf("can not use data type of %s in LT comparison", right.Kind().String())
		}
	case reflect.Float32, reflect.Float64:
		lv := left.Float()
		switch right.Kind() {
		case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
			rv := right.Int()
			return reflect.ValueOf(lv < float64(rv)), nil
		case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
			rv := right.Uint()
			return reflect.ValueOf(lv < float64(rv)), nil
		case reflect.Float32, reflect.Float64:
			rv := right.Float()
			return reflect.ValueOf(lv < rv), nil
		default:
			return reflect.ValueOf(nil), fmt.Errorf("can not use data type of %s in LT comparison", right.Kind().String())
		}
	default:
		if left.Type().String() == "time.Time" && right.Type().String() == "time.Time" {
			lv := left.Interface().(time.Time)
			rv := right.Interface().(time.Time)
			return reflect.ValueOf(lv.Before(rv)), nil
		}
		return reflect.ValueOf(nil), fmt.Errorf("can not use data type of %s in LT comparison", left.Kind().String())
	}
}

// EvaluateGreaterThanEqual will evaluate GreaterThanEqual operation over two value
func EvaluateGreaterThanEqual(left, right reflect.Value) (reflect.Value, error) {
	switch left.Kind() {
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		lv := left.Int()
		switch right.Kind() {
		case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
			rv := right.Int()
			return reflect.ValueOf(lv >= rv), nil
		case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
			rv := right.Uint()
			return reflect.ValueOf(lv >= int64(rv)), nil
		case reflect.Float32, reflect.Float64:
			rv := right.Float()
			return reflect.ValueOf(float64(lv) >= rv), nil
		default:
			return reflect.ValueOf(nil), fmt.Errorf("can not use data type of %s in GTE comparison", right.Kind().String())
		}
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		lv := left.Uint()
		switch right.Kind() {
		case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
			rv := right.Int()
			return reflect.ValueOf(int64(lv) >= rv), nil
		case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
			rv := right.Uint()
			return reflect.ValueOf(lv >= rv), nil
		case reflect.Float32, reflect.Float64:
			rv := right.Float()
			return reflect.ValueOf(float64(lv) >= rv), nil
		default:
			return reflect.ValueOf(nil), fmt.Errorf("can not use data type of %s in GTE comparison", right.Kind().String())
		}
	case reflect.Float32, reflect.Float64:
		lv := left.Float()
		switch right.Kind() {
		case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
			rv := right.Int()
			return reflect.ValueOf(lv >= float64(rv)), nil
		case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
			rv := right.Uint()
			return reflect.ValueOf(lv >= float64(rv)), nil
		case reflect.Float32, reflect.Float64:
			rv := right.Float()
			return reflect.ValueOf(lv >= rv), nil
		default:
			return reflect.ValueOf(nil), fmt.Errorf("can not use data type of %s in GTE comparison", right.Kind().String())
		}
	default:
		if left.Type().String() == "time.Time" && right.Type().String() == "time.Time" {
			lv := left.Interface().(time.Time)
			rv := right.Interface().(time.Time)
			return reflect.ValueOf(lv.After(rv) || lv == rv), nil
		}
		return reflect.ValueOf(nil), fmt.Errorf("can not use data type of %s in GTE comparison", left.Kind().String())
	}
}

// EvaluateLesserThanEqual will evaluate LesserThanEqual operation over two value
func EvaluateLesserThanEqual(left, right reflect.Value) (reflect.Value, error) {
	switch left.Kind() {
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		lv := left.Int()
		switch right.Kind() {
		case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
			rv := right.Int()
			return reflect.ValueOf(lv <= rv), nil
		case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
			rv := right.Uint()
			return reflect.ValueOf(lv <= int64(rv)), nil
		case reflect.Float32, reflect.Float64:
			rv := right.Float()
			return reflect.ValueOf(float64(lv) <= rv), nil
		default:
			return reflect.ValueOf(nil), fmt.Errorf("can not use data type of %s in LTE comparison", right.Kind().String())
		}
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		lv := left.Uint()
		switch right.Kind() {
		case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
			rv := right.Int()
			return reflect.ValueOf(int64(lv) <= rv), nil
		case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
			rv := right.Uint()
			return reflect.ValueOf(lv <= rv), nil
		case reflect.Float32, reflect.Float64:
			rv := right.Float()
			return reflect.ValueOf(float64(lv) <= rv), nil
		default:
			return reflect.ValueOf(nil), fmt.Errorf("can not use data type of %s in LTE comparison", right.Kind().String())
		}
	case reflect.Float32, reflect.Float64:
		lv := left.Float()
		switch right.Kind() {
		case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
			rv := right.Int()
			return reflect.ValueOf(lv <= float64(rv)), nil
		case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
			rv := right.Uint()
			return reflect.ValueOf(lv <= float64(rv)), nil
		case reflect.Float32, reflect.Float64:
			rv := right.Float()
			return reflect.ValueOf(lv <= rv), nil
		default:
			return reflect.ValueOf(nil), fmt.Errorf("can not use data type of %s in LTE comparison", right.Kind().String())
		}
	default:
		if left.Type().String() == "time.Time" && right.Type().String() == "time.Time" {
			lv := left.Interface().(time.Time)
			rv := right.Interface().(time.Time)
			return reflect.ValueOf(lv.Before(rv) || lv == rv), nil
		}
		return reflect.ValueOf(nil), fmt.Errorf("can not use data type of %s in LTE comparison", left.Kind().String())
	}
}

// EvaluateEqual will evaluate Equal operation over two value
func EvaluateEqual(left, right reflect.Value) (reflect.Value, error) {
	switch left.Kind() {
	case reflect.String:
		lv := left.String()
		if right.Kind() == reflect.String {
			rv := right.String()
			return reflect.ValueOf(lv == rv), nil
		}
		return reflect.ValueOf(false), nil
	case reflect.Bool:
		lv := left.Bool()
		if right.Kind() == reflect.Bool {
			rv := right.Bool()
			return reflect.ValueOf(lv == rv), nil
		}
		return reflect.ValueOf(false), nil
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		lv := left.Int()
		switch right.Kind() {
		case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
			rv := right.Int()
			return reflect.ValueOf(lv == rv), nil
		case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
			rv := right.Uint()
			return reflect.ValueOf(lv == int64(rv)), nil
		case reflect.Float32, reflect.Float64:
			rv := right.Float()
			return reflect.ValueOf(float64(lv) == rv), nil
		default:
			return reflect.ValueOf(nil), fmt.Errorf("can not use data type of %s in EQ comparison", right.Kind().String())
		}
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		lv := left.Uint()
		switch right.Kind() {
		case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
			rv := right.Int()
			return reflect.ValueOf(int64(lv) == rv), nil
		case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
			rv := right.Uint()
			return reflect.ValueOf(lv == rv), nil
		case reflect.Float32, reflect.Float64:
			rv := right.Float()
			return reflect.ValueOf(float64(lv) == rv), nil
		default:
			return reflect.ValueOf(nil), fmt.Errorf("can not use data type of %s in EQ comparison", right.Kind().String())
		}
	case reflect.Float32, reflect.Float64:
		lv := left.Float()
		switch right.Kind() {
		case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
			rv := right.Int()
			return reflect.ValueOf(lv == float64(rv)), nil
		case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
			rv := right.Uint()
			return reflect.ValueOf(lv == float64(rv)), nil
		case reflect.Float32, reflect.Float64:
			rv := right.Float()
			return reflect.ValueOf(lv == rv), nil
		default:
			return reflect.ValueOf(nil), fmt.Errorf("can not use data type of %s in EQ comparison", right.Kind().String())
		}
	default:
		if left.Type().String() == "time.Time" && right.Type().String() == "time.Time" {
			lv := left.Interface().(time.Time)
			rv := right.Interface().(time.Time)
			return reflect.ValueOf(lv == rv), nil
		}
		return reflect.ValueOf(nil), fmt.Errorf("can not use data type of %s in EQ comparison", left.Kind().String())
	}
}

// EvaluateNotEqual will evaluate NotEqual operation over two value
func EvaluateNotEqual(left, right reflect.Value) (reflect.Value, error) {
	switch left.Kind() {
	case reflect.String:
		lv := left.String()
		if right.Kind() == reflect.String {
			rv := right.String()
			return reflect.ValueOf(lv != rv), nil
		}
		return reflect.ValueOf(false), nil
	case reflect.Bool:
		lv := left.Bool()
		if right.Kind() == reflect.Bool {
			rv := right.Bool()
			return reflect.ValueOf(lv != rv), nil
		}
		return reflect.ValueOf(false), nil
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		lv := left.Int()
		switch right.Kind() {
		case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
			rv := right.Int()
			return reflect.ValueOf(lv != rv), nil
		case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
			rv := right.Uint()
			return reflect.ValueOf(lv != int64(rv)), nil
		case reflect.Float32, reflect.Float64:
			rv := right.Float()
			return reflect.ValueOf(float64(lv) != rv), nil
		default:
			return reflect.ValueOf(nil), fmt.Errorf("can not use data type of %s in EQ comparison", right.Kind().String())
		}
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		lv := left.Uint()
		switch right.Kind() {
		case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
			rv := right.Int()
			return reflect.ValueOf(int64(lv) != rv), nil
		case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
			rv := right.Uint()
			return reflect.ValueOf(lv != rv), nil
		case reflect.Float32, reflect.Float64:
			rv := right.Float()
			return reflect.ValueOf(float64(lv) != rv), nil
		default:
			return reflect.ValueOf(nil), fmt.Errorf("can not use data type of %s in EQ comparison", right.Kind().String())
		}
	case reflect.Float32, reflect.Float64:
		lv := left.Float()
		switch right.Kind() {
		case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
			rv := right.Int()
			return reflect.ValueOf(lv != float64(rv)), nil
		case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
			rv := right.Uint()
			return reflect.ValueOf(lv != float64(rv)), nil
		case reflect.Float32, reflect.Float64:
			rv := right.Float()
			return reflect.ValueOf(lv != rv), nil
		default:
			return reflect.ValueOf(nil), fmt.Errorf("can not use data type of %s in EQ comparison", right.Kind().String())
		}
	default:
		if left.Type().String() == "time.Time" && right.Type().String() == "time.Time" {
			lv := left.Interface().(time.Time)
			rv := right.Interface().(time.Time)
			return reflect.ValueOf(lv != rv), nil
		}
		return reflect.ValueOf(nil), fmt.Errorf("can not use data type of %s in EQ comparison", left.Kind().String())
	}
}

// EvaluateLogicAnd will evaluate LogicalAnd operation over two value
func EvaluateLogicAnd(left, right reflect.Value) (reflect.Value, error) {
	if left.Kind() == reflect.Bool && right.Kind() == reflect.Bool {
		lv := left.Bool()
		rv := right.Bool()
		return reflect.ValueOf(lv && rv), nil
	}
	return reflect.ValueOf(nil), fmt.Errorf("can not use data type of %s in Logical AND comparison", left.Kind().String())
}

// EvaluateLogicOr will evaluate LogicalOr operation over two value
func EvaluateLogicOr(left, right reflect.Value) (reflect.Value, error) {
	if left.Kind() == reflect.Bool && right.Kind() == reflect.Bool {
		lv := left.Bool()
		rv := right.Bool()
		return reflect.ValueOf(lv || rv), nil
	}
	return reflect.ValueOf(nil), fmt.Errorf("can not use data type of %s in Logical OR comparison", left.Kind().String())
}
