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
		AstID: unique.NewID(),
	}
}

// ExpressionAtom AST node graph
type ExpressionAtom struct {
	AstID   string
	GrlText string

	VariableName     string
	Constant         *Constant
	FunctionCall     *FunctionCall
	Variable         *Variable
	Negated          bool
	ExpressionAtom   *ExpressionAtom
	Value            reflect.Value
	ValueNode        model.ValueNode
	ArrayMapSelector *ArrayMapSelector

	Evaluated bool
}

// ExpressionAtomReceiver contains function to be implemented by other AST graph to receive an ExpressionAtom AST graph
type ExpressionAtomReceiver interface {
	AcceptExpressionAtom(exp *ExpressionAtom) error
}

// Clone will clone this ExpressionAtom. The new clone will have an identical structure
func (e *ExpressionAtom) Clone(cloneTable *pkg.CloneTable) *ExpressionAtom {
	clone := &ExpressionAtom{
		AstID:        unique.NewID(),
		GrlText:      e.GrlText,
		VariableName: e.VariableName,
		Negated:      e.Negated,
	}

	if e.Constant != nil {
		if cloneTable.IsCloned(e.Constant.AstID) {
			clone.Constant = cloneTable.Records[e.Constant.AstID].CloneInstance.(*Constant)
		} else {
			cloned := e.Constant.Clone(cloneTable)
			clone.Constant = cloned
			cloneTable.MarkCloned(e.Constant.AstID, cloned.AstID, e.Constant, cloned)
		}
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

	if e.ExpressionAtom != nil {
		if cloneTable.IsCloned(e.ExpressionAtom.AstID) {
			clone.ExpressionAtom = cloneTable.Records[e.ExpressionAtom.AstID].CloneInstance.(*ExpressionAtom)
		} else {
			cloned := e.ExpressionAtom.Clone(cloneTable)
			clone.ExpressionAtom = cloned
			cloneTable.MarkCloned(e.ExpressionAtom.AstID, cloned.AstID, e.ExpressionAtom, cloned)
		}
	}

	if e.ArrayMapSelector != nil {
		if cloneTable.IsCloned(e.ArrayMapSelector.AstID) {
			clone.ArrayMapSelector = cloneTable.Records[e.ArrayMapSelector.AstID].CloneInstance.(*ArrayMapSelector)
		} else {
			cloned := e.ArrayMapSelector.Clone(cloneTable)
			clone.ArrayMapSelector = cloned
			cloneTable.MarkCloned(e.ArrayMapSelector.AstID, cloned.AstID, e.ArrayMapSelector, cloned)
		}
	}

	return clone
}

// AcceptMemberVariable accept a member variable AST graph into this Variable graph
func (e *ExpressionAtom) AcceptMemberVariable(name string) {
	e.VariableName = name
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

// AcceptExpressionAtom will accept an ExpressionAtom AST graph into this ast graph
func (e *ExpressionAtom) AcceptExpressionAtom(ea *ExpressionAtom) error {
	if e.ExpressionAtom != nil {
		return errors.New("expression atom for ExpressionAtom already assigned")
	}
	e.ExpressionAtom = ea
	return nil
}

// AcceptConstant will accept a Constant AST graph into this ast graph
func (e *ExpressionAtom) AcceptConstant(cons *Constant) error {
	if e.Constant != nil {
		return errors.New("constant for ExpressionAtom already assigned")
	}
	e.Constant = cons
	return nil
}

// AcceptArrayMapSelector accept an array map selector into this variable graph
func (e *ExpressionAtom) AcceptArrayMapSelector(sel *ArrayMapSelector) error {
	e.ArrayMapSelector = sel
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
	} else if e.Constant != nil {
		buff.WriteString(e.Constant.GetSnapshot())
	} else if e.FunctionCall != nil && e.ExpressionAtom == nil {
		buff.WriteString(e.FunctionCall.GetSnapshot())
	} else if e.FunctionCall == nil && e.ExpressionAtom != nil && len(e.VariableName) == 0 {
		if e.Negated {
			buff.WriteString("!")
		}
		buff.WriteString(e.ExpressionAtom.GetSnapshot())
	} else if e.FunctionCall != nil && e.ExpressionAtom != nil {
		buff.WriteString(e.ExpressionAtom.GetSnapshot())
		buff.WriteString("->")
		buff.WriteString(e.FunctionCall.GetSnapshot())
	} else if len(e.VariableName) > 0 && e.ExpressionAtom != nil {
		buff.WriteString(e.ExpressionAtom.GetSnapshot())
		buff.WriteString("->MV:")
		buff.WriteString(e.VariableName)
	} else if e.ArrayMapSelector != nil && e.ExpressionAtom != nil {
		buff.WriteString(e.ExpressionAtom.GetSnapshot())
		buff.WriteString("-[]>")
		buff.WriteString(e.ArrayMapSelector.GetSnapshot())
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
func (e *ExpressionAtom) Evaluate(dataContext IDataContext, memory *WorkingMemory) (val reflect.Value, err error) {
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
		//t, _ := e.Variable.ValueNode.GetType()
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
		// e.Evaluated = true
		return ret, err
	}
	if e.ExpressionAtom != nil && e.FunctionCall == nil && len(e.VariableName) == 0 && e.ArrayMapSelector == nil {
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
		return e.Value, err
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
		e.Evaluated = true
		return e.Value, nil
	}
	if e.ExpressionAtom != nil && len(e.VariableName) > 0 {
		_, err := e.ExpressionAtom.Evaluate(dataContext, memory)
		if err != nil {
			return reflect.Value{}, err
		}
		valueNode, err := e.ExpressionAtom.ValueNode.GetChildNodeByField(e.VariableName)
		if err != nil {
			return reflect.Value{}, err
		}
		e.ValueNode = valueNode
		e.Value = valueNode.Value()
		e.Evaluated = true

		return e.Value, nil
	}
	if e.ExpressionAtom != nil && e.ArrayMapSelector != nil && len(e.VariableName) == 0 {

		_, err := e.ExpressionAtom.Evaluate(dataContext, memory)
		if err != nil {
			return reflect.Value{}, err
		}
		selValue, err := e.ArrayMapSelector.Evaluate(dataContext, memory)
		if err != nil {
			return reflect.Value{}, err
		}
		var valueNode model.ValueNode
		if e.ExpressionAtom.ValueNode.IsArray() {
			valueNode, err = e.ExpressionAtom.ValueNode.GetChildNodeByIndex(int(selValue.Int()))
			if err != nil {
				return reflect.Value{}, err
			}
		} else if e.ExpressionAtom.ValueNode.IsMap() {
			valueNode, err = e.ExpressionAtom.ValueNode.GetChildNodeBySelector(selValue)
			if err != nil {
				return reflect.Value{}, err
			}
		} else {
			return reflect.Value{}, fmt.Errorf("%s is not an array nor map", e.Variable.ValueNode.IdentifiedAs())
		}

		e.ValueNode = valueNode
		e.Value = valueNode.Value()

		return e.Value, nil
	}
	panic("should not be reached")
}
