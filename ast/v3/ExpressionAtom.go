package v3

import (
	"bytes"
	"errors"
	"fmt"
	"github.com/hyperjumptech/grule-rule-engine/ast/unique"
	"github.com/hyperjumptech/grule-rule-engine/model"
	"reflect"

	"github.com/hyperjumptech/grule-rule-engine/pkg"
)

// NewExpressionAtom create new instance of ExpressionAtom
func NewExpressionAtom() *ExpressionAtom {
	return &ExpressionAtom{
		AstID: unique.NewId(),
	}
}

// ExpressionAtom AST node graph
type ExpressionAtom struct {
	AstID   string
	GrlText string

	Constant       *Constant
	FunctionCall   *FunctionCall
	Variable       *Variable
	Negated        bool
	ExpressionAtom *ExpressionAtom
	Value          reflect.Value
	ValueNode      model.ValueNode

	Evaluated bool
}

// ExpressionAtomReceiver contains function to be implemented by other AST graph to receive an ExpressionAtom AST graph
type ExpressionAtomReceiver interface {
	AcceptExpressionAtom(exp *ExpressionAtom) error
}

// Clone will clone this ExpressionAtom. The new clone will have an identical structure
func (e *ExpressionAtom) Clone(cloneTable *pkg.CloneTable) *ExpressionAtom {
	clone := &ExpressionAtom{
		AstID:   unique.NewId(),
		GrlText: e.GrlText,
	}

	if e.Variable != nil {
		if cloneTable.IsCloned(e.Variable.AstID) {
			clone.Variable = cloneTable.Records[e.Variable.AstID].CloneInstance.(*Variable)
		} else {
			cloned := e.Variable.Clone(cloneTable)
			clone.Variable = cloned
			cloneTable.MarkCloned(e.Variable.AstID, cloned.AstID, e.Variable, cloned)
		}
	}

	if e.FunctionCall != nil {
		if cloneTable.IsCloned(e.FunctionCall.AstID) {
			clone.FunctionCall = cloneTable.Records[e.FunctionCall.AstID].CloneInstance.(*FunctionCall)
		} else {
			cloned := e.FunctionCall.Clone(cloneTable)
			clone.FunctionCall = cloned
			cloneTable.MarkCloned(e.FunctionCall.AstID, cloned.AstID, e.FunctionCall, cloned)
		}
	}

	return clone
}

// AcceptVariable will accept an Variable AST graph into this ast graph
func (e *ExpressionAtom) AcceptVariable(vari *Variable) error {
	if e.Variable != nil {
		return errors.New("variable for ExpressionAtom already assigned")
	}
	e.Variable = vari
	return nil
}

// AcceptFunctionCall will accept an FunctionCall AST graph into this ast graph
func (e *ExpressionAtom) AcceptFunctionCall(fun *FunctionCall) error {
	if e.FunctionCall != nil {
		return errors.New("function call for ExpressionAtom already assigned")
	}
	e.FunctionCall = fun
	return nil
}

func (e *ExpressionAtom) AcceptExpressionAtom(ea *ExpressionAtom) error {
	if e.ExpressionAtom != nil {
		return errors.New("expression atom for ExpressionAtom already assigned")
	}
	e.ExpressionAtom = ea
	return nil
}

func (e *ExpressionAtom) AcceptConstant(cons *Constant) error {
	if e.Constant != nil {
		return errors.New("constant for ExpressionAtom already assigned")
	}
	e.Constant = cons
	return nil
}

// GetAstID get the UUID asigned for this AST graph node
func (e *ExpressionAtom) GetAstID() string {
	return e.AstID
}

// GetGrlText get the expression syntax related to this graph when it wast constructed
func (e *ExpressionAtom) GetGrlText() string {
	return e.GrlText
}

// GetSnapshot will create a structure signature or AST graph
func (e *ExpressionAtom) GetSnapshot() string {
	var buff bytes.Buffer
	buff.WriteString(EXPRESSIONATOM)
	buff.WriteString("(")
	if e.Variable != nil {
		buff.WriteString(e.Variable.GetSnapshot())
	} else if e.FunctionCall != nil {
		buff.WriteString(e.FunctionCall.GetSnapshot())
	} else if e.Constant != nil {
		buff.WriteString(e.Constant.GetSnapshot())
	} else if e.ExpressionAtom != nil {
		if e.Negated {
			buff.WriteString("!")
		}
		buff.WriteString(e.ExpressionAtom.GetSnapshot())
	}
	buff.WriteString(")")
	return buff.String()
}

// SetGrlText set the expression syntax related to this graph when it was constructed. Only ANTLR4 listener should
// call this function.
func (e *ExpressionAtom) SetGrlText(grlText string) {
	e.GrlText = grlText
}

// Evaluate will evaluate this AST graph for when scope evaluation
func (e *ExpressionAtom) Evaluate(dataContext IDataContext, memory *WorkingMemory) (reflect.Value, error) {
	if e.Evaluated == true {
		return e.Value, nil
	}
	if e.Constant != nil {
		val, err := e.Constant.Evaluate(dataContext, memory)
		if err != nil {
			return reflect.Value{}, err
		}
		e.Value = val
		e.ValueNode = model.NewGoValueNode(val, fmt.Sprintf("%s->%s", val.Type().String(), val.String()))
		e.Evaluated = true
		return val, err
	}
	if e.Variable != nil {
		val, err := e.Variable.Evaluate(dataContext, memory)
		if err != nil {
			return reflect.Value{}, err
		}
		e.Value = val
		e.ValueNode = e.Variable.ValueNode
		e.Evaluated = true
		return val, err
	}
	if e.ExpressionAtom == nil && e.FunctionCall != nil {
		valueNode := dataContext.Get("DEFUNC")
		args, err := e.FunctionCall.EvaluateArgumentList(dataContext, memory)
		if err != nil {
			return reflect.Value{}, err
		}
		ret, err := valueNode.CallFunction(e.FunctionCall.FunctionName, args...)
		if err != nil {
			return reflect.Value{}, err
		}
		e.Value = ret
		e.ValueNode = model.NewGoValueNode(e.Value, fmt.Sprintf("%s()", e.FunctionCall.FunctionName))
		e.Evaluated = true
		return ret, err
	}
	if e.ExpressionAtom != nil && e.FunctionCall == nil {
		val, err := e.ExpressionAtom.Evaluate(dataContext, memory)
		if err != nil {
			return reflect.Value{}, err
		}
		e.Value = val
		e.ValueNode = e.ExpressionAtom.ValueNode
		if e.Negated {
			if e.Value.Kind() == reflect.Bool {
				e.Value = reflect.ValueOf(!e.Value.Bool())
				e.ValueNode = model.NewGoValueNode(e.Value, fmt.Sprintf("!%s", e.GrlText))
			} else {
				AstLog.Warnf("Expression \"%s\" is a negation to non boolean value, negation is ignored.", e.ExpressionAtom.GrlText)
			}
		}
		e.Evaluated = true
		return val, err
	}
	if e.ExpressionAtom != nil && e.FunctionCall != nil {
		_, err := e.ExpressionAtom.Evaluate(dataContext, memory)
		if err != nil {
			return reflect.ValueOf(nil), err
		}

		args, err := e.FunctionCall.EvaluateArgumentList(dataContext, memory)
		if err != nil {
			return reflect.ValueOf(nil), err
		}

		retVal, err := e.ExpressionAtom.ValueNode.CallFunction(e.FunctionCall.FunctionName, args...)
		if err != nil {
			return reflect.ValueOf(nil), err
		}

		if retVal.IsValid() {
			e.Value = retVal
		}
		e.ValueNode = e.ExpressionAtom.ValueNode.ContinueWithValue(retVal, e.FunctionCall.FunctionName)
		return e.Value, nil
	}
	panic("should not be reached")
}
