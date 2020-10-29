package ast

import (
	"bytes"
	"fmt"
	"github.com/hyperjumptech/grule-rule-engine/model"
	"reflect"

	"github.com/google/uuid"
	"github.com/hyperjumptech/grule-rule-engine/pkg"
)

// NewVariable create new instance of Variable
func NewVariable() *Variable {
	return &Variable{
		AstID: uuid.New().String(),
	}
}

// Variable AST graph node
type Variable struct {
	AstID   string
	GrlText string

	Name             string
	Constant         *Constant
	Variable         *Variable
	FunctionCall     *FunctionCall
	ArrayMapSelector *ArrayMapSelector

	ValueNode model.ValueNode
	Value     reflect.Value
}

// Clone will clone this Variable. The new clone will have an identical structure
func (e *Variable) Clone(cloneTable *pkg.CloneTable) *Variable {
	clone := &Variable{
		AstID:   uuid.New().String(),
		GrlText: e.GrlText,
		Name:    e.Name,
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

// VariableReceiver should be implemented by AST graph node to receive Variable AST graph node
type VariableReceiver interface {
	AcceptVariable(exp *Variable) error
}

// AcceptVariable accept a variable AST graph into this Variable graph
func (e *Variable) AcceptVariable(vari *Variable) error {
	e.Variable = vari
	return nil
}

// AcceptArrayMapSelector accept an array map selector into this variable graph
func (e *Variable) AcceptArrayMapSelector(sel *ArrayMapSelector) error {
	e.ArrayMapSelector = sel
	return nil
}

// AcceptFunctionCall accept a function call AST graph into this variable graph
func (e *Variable) AcceptFunctionCall(fu *FunctionCall) error {
	e.FunctionCall = fu
	return nil
}

// AcceptConstant accept a constant AST graph into this variable graph.
func (e *Variable) AcceptConstant(con *Constant) error {
	e.Constant = con
	return nil
}

// GetAstID get the UUID asigned for this AST graph node
func (e *Variable) GetAstID() string {
	return e.AstID
}

// GetGrlText get the expression syntax related to this graph when it wast constructed
func (e *Variable) GetGrlText() string {
	return e.GrlText
}

// GetSnapshot will create a structure signature or AST graph
func (e *Variable) GetSnapshot() string {
	var buff bytes.Buffer
	buff.WriteString(VARIABLE)
	buff.WriteString("(")
	if len(e.Name) > 0 && e.Variable == nil {
		buff.WriteString("N:")
		buff.WriteString(e.Name)
	} else if e.Constant != nil {
		buff.WriteString(e.Constant.GetSnapshot())
	} else if e.Variable != nil && e.FunctionCall != nil {
		buff.WriteString(fmt.Sprintf("O:%s->%s", e.Variable.GetSnapshot(), e.FunctionCall.GetSnapshot()))
	} else if e.Variable != nil && len(e.Name) > 0 {
		buff.WriteString(fmt.Sprintf("O:%s->%s", e.Variable.GetSnapshot(), e.Name))
	} else if e.Variable != nil && e.ArrayMapSelector != nil {
		buff.WriteString(fmt.Sprintf("O:%s->%s", e.Variable.GetSnapshot(), e.ArrayMapSelector.GetSnapshot()))
	}
	buff.WriteString(")")
	return buff.String()
}

// SetGrlText set the expression syntax related to this graph when it was constructed. Only ANTLR4 listener should
// call this function.
func (e *Variable) SetGrlText(grlText string) {
	e.GrlText = grlText
}

// Assign will assign the specified value to the variable
func (e *Variable) Assign(newVal reflect.Value, dataContext IDataContext, memory *WorkingMemory) error {
	if len(e.Name) > 0 && e.Variable == nil {
		err := dataContext.Add(e.Name, pkg.ValueToInterface(newVal))
		if err == nil {
			dataContext.IncrementVariableChangeCount()
		}
		return err
	}
	if e.Constant != nil {
		return fmt.Errorf("can not change constant")
	}
	if e.Variable != nil && e.FunctionCall != nil {
		return fmt.Errorf("can not change function call")
	}
	if e.Variable != nil && len(e.Name) > 0 {
		_, err := e.Variable.Evaluate(dataContext, memory)
		if err != nil {
			return err
		}
		err = e.Variable.ValueNode.SetObjectValueByField(e.Name, newVal)
		if err == nil {
			dataContext.IncrementVariableChangeCount()
			memory.ResetVariable(e)
		}
		return err
	}
	if e.Variable != nil && e.ArrayMapSelector != nil {
		_, err := e.Variable.Evaluate(dataContext, memory)
		if err != nil {
			return err
		}
		_, err = e.ArrayMapSelector.Evaluate(dataContext, memory)
		if err != nil {
			return err
		}
		if e.Variable.ValueNode.IsArray() {
			err := e.Variable.ValueNode.SetArrayValueAt(int(e.ArrayMapSelector.Value.Int()), newVal)
			if err == nil {
				memory.ResetVariable(e)
			}
			return err
		}
		if e.Variable.ValueNode.IsMap() {
			err := e.Variable.ValueNode.SetMapValueAt(e.ArrayMapSelector.Value, newVal)
			if err == nil {
				memory.ResetVariable(e)
			}
			return err
		}
	}
	return fmt.Errorf("this code part should not be reached")
}

// Evaluate will evaluate this AST graph for when scope evaluation
func (e *Variable) Evaluate(dataContext IDataContext, memory *WorkingMemory) (reflect.Value, error) {
	if len(e.Name) > 0 && e.Variable == nil {
		valueNode := dataContext.Get(e.Name)
		if valueNode == nil {
			return reflect.ValueOf(nil), fmt.Errorf("non existent key %s", e.Name)
		}
		e.ValueNode = valueNode
		e.Value = valueNode.Value()
		return e.Value, nil
	}
	if e.Constant != nil {
		val, err := e.Constant.Evaluate(dataContext, memory)
		if err != nil {
			return reflect.ValueOf(nil), err
		}
		valueNode := model.NewGoValueNode(val, fmt.Sprintf("%s->%s", val.Type().String(), val.String()))
		e.ValueNode = valueNode
		e.Value = valueNode.Value()
		return e.Value, nil
	}
	if e.Variable != nil && e.FunctionCall != nil {
		_, err := e.Variable.Evaluate(dataContext, memory)
		if err != nil {
			return reflect.ValueOf(nil), err
		}

		args, err := e.FunctionCall.EvaluateArgumentList(dataContext, memory)
		if err != nil {
			return reflect.ValueOf(nil), err
		}

		retVal, err := e.Variable.ValueNode.CallFunction(e.FunctionCall.FunctionName, args...)
		if err != nil {
			return reflect.ValueOf(nil), err
		}

		if retVal.IsValid() {
			e.Value = retVal
		}
		e.ValueNode = e.Variable.ValueNode.ContinueWithValue(retVal, e.FunctionCall.FunctionName)
		return e.Value, nil
	}
	if e.Variable != nil && len(e.Name) > 0 {
		_, err := e.Variable.Evaluate(dataContext, memory)
		if err != nil {
			return reflect.Value{}, err
		}
		valueNode, err := e.Variable.ValueNode.GetChildNodeByField(e.Name)
		if err != nil {
			return reflect.Value{}, err
		}
		e.ValueNode = valueNode
		e.Value = valueNode.Value()
		return e.Value, nil
	}
	if e.Variable != nil && e.ArrayMapSelector != nil {
		_, err := e.Variable.Evaluate(dataContext, memory)
		if err != nil {
			return reflect.Value{}, err
		}
		selValue, err := e.ArrayMapSelector.Evaluate(dataContext, memory)
		if err != nil {
			return reflect.Value{}, err
		}
		var valueNode model.ValueNode
		if e.Variable.ValueNode.IsArray() {
			valueNode, err = e.Variable.ValueNode.GetChildNodeByIndex(int(selValue.Int()))
			if err != nil {
				return reflect.Value{}, err
			}
		} else if e.Variable.ValueNode.IsMap() {
			valueNode, err = e.Variable.ValueNode.GetChildNodeBySelector(selValue)
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
	return reflect.ValueOf(nil), fmt.Errorf("this code part should not be reached")
}
