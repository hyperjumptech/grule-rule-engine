package ast

import (
	"bytes"
	"errors"
	"fmt"
	"github.com/google/uuid"
	"reflect"
)

// NewAssignment will create new instance of Assignment AST Node
func NewAssignment() *Assignment {
	return &Assignment{
		AstID: uuid.New().String(),
	}
}

// Assignment ast node to store assigment expression.
type Assignment struct {
	AstID         string
	GrlText       string
	DataContext   *DataContext
	WorkingMemory *WorkingMemory

	Variable   *Variable
	Expression *Expression
}

// InitializeContext will initialize this AST graph with data context and working memory before running rule on them.
func (e *Assignment) InitializeContext(dataCtx *DataContext, workingMemory *WorkingMemory) {
	e.DataContext = dataCtx
	e.WorkingMemory = workingMemory
	e.Variable.InitializeContext(dataCtx, workingMemory)
	e.Expression.InitializeContext(dataCtx, workingMemory)
}

// AcceptExpression will accept an Expression AST graph into this ast graph
func (e *Assignment) AcceptExpression(exp *Expression) error {
	if e.Expression != nil {
		return errors.New("expression for assignment already assigned")
	}
	e.Expression = exp
	return nil
}

// AcceptVariable will accept an Variable AST graph into this ast graph
func (e *Assignment) AcceptVariable(vari *Variable) error {
	if e.Variable != nil {
		return errors.New("variable for assignment already assigned")
	}
	e.Variable = vari
	return nil
}

// GetAstID get the UUID asigned for this AST graph node
func (e *Assignment) GetAstID() string {
	return e.AstID
}

// GetGrlText get the expression syntax related to this graph when it wast constructed
func (e *Assignment) GetGrlText() string {
	return e.GrlText
}

// GetSnapshot will create a structure signature or AST graph
func (e *Assignment) GetSnapshot() string {
	var buff bytes.Buffer
	buff.WriteString(e.Variable.GetSnapshot())
	buff.WriteString("=")
	buff.WriteString(e.Expression.GetSnapshot())
	buff.WriteString(";")
	return buff.String()
}

// SetGrlText set the expression syntax related to this graph when it was constructed. Only ANTLR4 listener should
// call this function.
func (e *Assignment) SetGrlText(grlText string) {
	e.GrlText = grlText
}

// Execute will execute this graph in the Then scope
func (e *Assignment) Execute() error {
	varVal, err := e.Variable.Evaluate()
	if err != nil {
		return err
	}
	exprVal, err := e.Expression.Evaluate()
	if err != nil {
		return err
	}
	e.WorkingMemory.Reset(e.Variable.Name)
	switch varVal.Kind() {
	case reflect.Int:
		switch exprVal.Kind() {
		case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
			return e.DataContext.SetValue(e.Variable.Name, exprVal)
		case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
			v := exprVal.Uint()
			return e.DataContext.SetValue(e.Variable.Name, reflect.ValueOf(int(v)))
		case reflect.Float32, reflect.Float64:
			v := exprVal.Float()
			return e.DataContext.SetValue(e.Variable.Name, reflect.ValueOf(int(v)))
		default:
			return fmt.Errorf("can not assign %s to %s", exprVal.Kind().String(), varVal.Kind().String())
		}
	case reflect.Int8:
		switch exprVal.Kind() {
		case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
			v := exprVal.Int()
			if v >= -128 && v <= 127 {
				return e.DataContext.SetValue(e.Variable.Name, reflect.ValueOf(int8(v)))
			}
			return fmt.Errorf("variable of type %s will be overflowed with value %d", varVal.Kind().String(), v)
		case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
			v := exprVal.Uint()
			if v <= 127 {
				return e.DataContext.SetValue(e.Variable.Name, reflect.ValueOf(int8(v)))
			}
			return fmt.Errorf("variable of type %s will be overflowed with value %d", varVal.Kind().String(), v)
		case reflect.Float32, reflect.Float64:
			v := exprVal.Float()
			if v >= -128 && v <= 127 {
				return e.DataContext.SetValue(e.Variable.Name, reflect.ValueOf(int8(v)))
			}
			return fmt.Errorf("variable of type %s will be overflowed with value %f", varVal.Kind().String(), v)
		default:
			return fmt.Errorf("can not assign %s to %s", exprVal.Kind().String(), varVal.Kind().String())
		}
	case reflect.Int16:
		switch exprVal.Kind() {
		case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
			v := exprVal.Int()
			if v >= -32768 && v <= 32767 {
				return e.DataContext.SetValue(e.Variable.Name, reflect.ValueOf(int16(v)))
			}
			return fmt.Errorf("variable of type %s will be overflowed with value %d", varVal.Kind().String(), v)
		case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
			v := exprVal.Uint()
			if v <= 32767 {
				return e.DataContext.SetValue(e.Variable.Name, reflect.ValueOf(int16(v)))
			}
			return fmt.Errorf("variable of type %s will be overflowed with value %d", varVal.Kind().String(), v)
		case reflect.Float32, reflect.Float64:
			v := exprVal.Float()
			if v >= -32768 && v <= 32767 {
				return e.DataContext.SetValue(e.Variable.Name, reflect.ValueOf(int16(v)))
			}
			return fmt.Errorf("variable of type %s will be overflowed with value %f", varVal.Kind().String(), v)
		default:
			return fmt.Errorf("can not assign %s to %s", exprVal.Kind().String(), varVal.Kind().String())
		}
	case reflect.Int32:
		switch exprVal.Kind() {
		case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
			v := exprVal.Int()
			if v >= -2147483648 && v <= 2147483647 {
				return e.DataContext.SetValue(e.Variable.Name, reflect.ValueOf(int32(v)))
			}
			return fmt.Errorf("variable of type %s will be overflowed with value %d", varVal.Kind().String(), v)
		case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
			v := exprVal.Uint()
			if v <= 2147483647 {
				return e.DataContext.SetValue(e.Variable.Name, reflect.ValueOf(int32(v)))
			}
			return fmt.Errorf("variable of type %s will be overflowed with value %d", varVal.Kind().String(), v)
		case reflect.Float32, reflect.Float64:
			v := exprVal.Float()
			if v >= -2147483648 && v <= 2147483647 {
				return e.DataContext.SetValue(e.Variable.Name, reflect.ValueOf(int32(v)))
			}
			return fmt.Errorf("variable of type %s will be overflowed with value %f", varVal.Kind().String(), v)
		default:
			return fmt.Errorf("can not assign %s to %s", exprVal.Kind().String(), varVal.Kind().String())
		}
	case reflect.Int64:
		switch exprVal.Kind() {
		case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
			v := exprVal.Int()
			return e.DataContext.SetValue(e.Variable.Name, reflect.ValueOf(v))
		case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
			v := exprVal.Uint()
			if v <= 9223372036854775807 {
				return e.DataContext.SetValue(e.Variable.Name, reflect.ValueOf(int64(v)))
			}
			return fmt.Errorf("variable of type %s will be overflowed with value %d", varVal.Kind().String(), v)
		case reflect.Float32, reflect.Float64:
			v := exprVal.Float()
			if v >= -9223372036854775808 && v <= 9223372036854775807 {
				return e.DataContext.SetValue(e.Variable.Name, reflect.ValueOf(int64(v)))
			}
			return fmt.Errorf("variable of type %s will be overflowed with value %f", varVal.Kind().String(), v)
		default:
			return fmt.Errorf("can not assign %s to %s", exprVal.Kind().String(), varVal.Kind().String())
		}
	case reflect.Uint:
		switch exprVal.Kind() {
		case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
			v := exprVal.Int()
			if v >= 0 {
				return e.DataContext.SetValue(e.Variable.Name, reflect.ValueOf(uint(v)))
			}
			return fmt.Errorf("variable of type %s will be overflowed with value %d", varVal.Kind().String(), v)
		case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
			v := exprVal.Uint()
			return e.DataContext.SetValue(e.Variable.Name, reflect.ValueOf(uint(v)))
		case reflect.Float32, reflect.Float64:
			v := exprVal.Float()
			if v >= 0 {
				return e.DataContext.SetValue(e.Variable.Name, reflect.ValueOf(uint(v)))
			}
			return fmt.Errorf("variable of type %s will be overflowed with value %f", varVal.Kind().String(), v)
		default:
			return fmt.Errorf("can not assign %s to %s", exprVal.Kind().String(), varVal.Kind().String())
		}
	case reflect.Uint8:
		switch exprVal.Kind() {
		case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
			v := exprVal.Int()
			if v >= 0 && v <= 255 {
				return e.DataContext.SetValue(e.Variable.Name, reflect.ValueOf(uint8(v)))
			}
			return fmt.Errorf("variable of type %s will be overflowed with value %d", varVal.Kind().String(), v)
		case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
			v := exprVal.Uint()
			if v >= 0 && v <= 255 {
				return e.DataContext.SetValue(e.Variable.Name, reflect.ValueOf(uint8(v)))
			}
			return fmt.Errorf("variable of type %s will be overflowed with value %d", varVal.Kind().String(), v)
		case reflect.Float32, reflect.Float64:
			v := exprVal.Float()
			if v >= 0 && v <= 255 {
				return e.DataContext.SetValue(e.Variable.Name, reflect.ValueOf(uint8(v)))
			}
			return fmt.Errorf("variable of type %s will be overflowed with value %f", varVal.Kind().String(), v)
		default:
			return fmt.Errorf("can not assign %s to %s", exprVal.Kind().String(), varVal.Kind().String())
		}
	case reflect.Uint16:
		switch exprVal.Kind() {
		case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
			v := exprVal.Int()
			if v >= 0 && v <= 65535 {
				return e.DataContext.SetValue(e.Variable.Name, reflect.ValueOf(uint16(v)))
			}
			return fmt.Errorf("variable of type %s will be overflowed with value %d", varVal.Kind().String(), v)
		case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
			v := exprVal.Uint()
			if v >= 0 && v <= 65535 {
				return e.DataContext.SetValue(e.Variable.Name, reflect.ValueOf(uint16(v)))
			}
			return fmt.Errorf("variable of type %s will be overflowed with value %d", varVal.Kind().String(), v)
		case reflect.Float32, reflect.Float64:
			v := exprVal.Float()
			if v >= 0 && v <= 65535 {
				return e.DataContext.SetValue(e.Variable.Name, reflect.ValueOf(uint16(v)))
			}
			return fmt.Errorf("variable of type %s will be overflowed with value %f", varVal.Kind().String(), v)
		default:
			return fmt.Errorf("can not assign %s to %s", exprVal.Kind().String(), varVal.Kind().String())
		}
	case reflect.Uint32:
		switch exprVal.Kind() {
		case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
			v := exprVal.Int()
			if v >= 0 && v <= 4294967295 {
				return e.DataContext.SetValue(e.Variable.Name, reflect.ValueOf(uint32(v)))
			}
			return fmt.Errorf("variable of type %s will be overflowed with value %d", varVal.Kind().String(), v)
		case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
			v := exprVal.Uint()
			if v >= 0 && v <= 4294967295 {
				return e.DataContext.SetValue(e.Variable.Name, reflect.ValueOf(uint32(v)))
			}
			return fmt.Errorf("variable of type %s will be overflowed with value %d", varVal.Kind().String(), v)
		case reflect.Float32, reflect.Float64:
			v := exprVal.Float()
			if v >= 0 && v <= 4294967295 {
				return e.DataContext.SetValue(e.Variable.Name, reflect.ValueOf(uint32(v)))
			}
			return fmt.Errorf("variable of type %s will be overflowed with value %f", varVal.Kind().String(), v)
		default:
			return fmt.Errorf("can not assign %s to %s", exprVal.Kind().String(), varVal.Kind().String())
		}
	case reflect.Uint64:
		switch exprVal.Kind() {
		case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
			v := exprVal.Int()
			if v >= 0 {
				return e.DataContext.SetValue(e.Variable.Name, reflect.ValueOf(uint64(v)))
			}
			return fmt.Errorf("variable of type %s will be overflowed with value %d", varVal.Kind().String(), v)
		case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
			return e.DataContext.SetValue(e.Variable.Name, exprVal)
		case reflect.Float32, reflect.Float64:
			v := exprVal.Float()
			if v >= 0 && v <= 18446744073709551615 {
				return e.DataContext.SetValue(e.Variable.Name, reflect.ValueOf(uint64(v)))
			}
			return fmt.Errorf("variable of type %s will be overflowed with value %f", varVal.Kind().String(), v)
		default:
			return fmt.Errorf("can not assign %s to %s", exprVal.Kind().String(), varVal.Kind().String())
		}
	case reflect.Float32:
		switch exprVal.Kind() {
		case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
			v := exprVal.Int()
			return e.DataContext.SetValue(e.Variable.Name, reflect.ValueOf(float32(v)))
		case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
			v := exprVal.Uint()
			return e.DataContext.SetValue(e.Variable.Name, reflect.ValueOf(float32(v)))
		case reflect.Float32, reflect.Float64:
			v := exprVal.Float()
			return e.DataContext.SetValue(e.Variable.Name, reflect.ValueOf(float32(v)))
		default:
			return fmt.Errorf("can not assign %s to %s", exprVal.Kind().String(), varVal.Kind().String())
		}
	case reflect.Float64:
		switch exprVal.Kind() {
		case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
			v := exprVal.Int()
			return e.DataContext.SetValue(e.Variable.Name, reflect.ValueOf(float64(v)))
		case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
			v := exprVal.Uint()
			return e.DataContext.SetValue(e.Variable.Name, reflect.ValueOf(float64(v)))
		case reflect.Float32, reflect.Float64:
			return e.DataContext.SetValue(e.Variable.Name, exprVal)
		default:
			return fmt.Errorf("can not assign %s to %s", exprVal.Kind().String(), varVal.Kind().String())
		}
	default:
		if varVal.Kind() == exprVal.Kind() {
			return e.DataContext.SetValue(e.Variable.Name, exprVal)
		}
		return fmt.Errorf("can not assign %s to %s", exprVal.Kind().String(), varVal.Kind().String())
	}
}
