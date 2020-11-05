package v3

import (
	"bytes"
	"fmt"
	"github.com/hyperjumptech/grule-rule-engine/ast/unique"
	"reflect"

	"github.com/hyperjumptech/grule-rule-engine/pkg"
)

// NewConstant will create new instance of Constant
func NewConstant() *Constant {
	return &Constant{
		AstID: unique.NewId(),
	}
}

// Constant AST node that stores AST graph for Constants
type Constant struct {
	AstID         string
	GrlText       string
	Snapshot      string
	DataContext   IDataContext
	WorkingMemory *WorkingMemory
	Value         reflect.Value
	IsNil         bool
}

// Clone will clone this Constant. The new clone will have an identical structure
func (e *Constant) Clone(cloneTable *pkg.CloneTable) *Constant {
	clone := &Constant{
		AstID:   unique.NewId(),
		GrlText: e.GrlText,
		Value:   e.Value,
	}

	return clone
}

// ConstantReceiver should be implemented by AST Graph node to receive a Constant Graph Node.
type ConstantReceiver interface {
	AcceptConstant(con *Constant) error
}

// GetAstID get the UUID asigned for this AST graph node
func (e *Constant) GetAstID() string {
	return e.AstID
}

// GetGrlText get the expression syntax related to this graph when it wast constructed
func (e *Constant) GetGrlText() string {
	return e.GrlText
}

// GetSnapshot will create a structure signature or AST graph
func (e *Constant) GetSnapshot() string {
	var buff bytes.Buffer
	buff.WriteString(CONSTANT)
	buff.WriteString("(")
	buff.WriteString(e.Value.Kind().String())
	buff.WriteString("->")
	switch e.Value.Kind() {
	case reflect.String:
		buff.WriteString(fmt.Sprintf("\"%s\"", e.Value.String()))
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		buff.WriteString(fmt.Sprintf("%d", e.Value.Int()))
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		buff.WriteString(fmt.Sprintf("%d", e.Value.Uint()))
	case reflect.Float32, reflect.Float64:
		buff.WriteString(fmt.Sprintf("%f", e.Value.Float()))
	case reflect.Bool:
		buff.WriteString(fmt.Sprintf("%v", e.Value.Bool()))
	}
	buff.WriteString(")")
	return buff.String()
}

// SetGrlText set the expression syntax related to this graph when it was constructed. Only ANTLR4 listener should
// call this function.
func (e *Constant) SetGrlText(grlText string) {
	e.GrlText = grlText
}

// AcceptIntegerLiteral will accept integer literal
func (e *Constant) AcceptIntegerLiteral(fun *IntegerLiteral) {
	e.Value = reflect.ValueOf(fun.Integer)
}

// AcceptStringLiteral will accept string literal
func (e *Constant) AcceptStringLiteral(fun *StringLiteral) {
	e.Value = reflect.ValueOf(fun.String)
}

// AcceptFloatLiteral will accept float literal
func (e *Constant) AcceptFloatLiteral(fun *FloatLiteral) {
	e.Value = reflect.ValueOf(fun.Float)
}

// AcceptBooleanLiteral will accept boolean literal
func (e *Constant) AcceptBooleanLiteral(fun *BooleanLiteral) {
	e.Value = reflect.ValueOf(fun.Boolean)
}

// Evaluate will evaluate this AST graph for when scope evaluation
func (e *Constant) Evaluate(dataContext IDataContext, memory *WorkingMemory) (reflect.Value, error) {
	if e.IsNil {
		return reflect.ValueOf(nil), nil
	}
	return e.Value, nil
}