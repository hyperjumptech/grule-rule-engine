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
	"fmt"
	"github.com/hyperjumptech/grule-rule-engine/ast/unique"
	"github.com/hyperjumptech/grule-rule-engine/pkg"
	"reflect"
)

// NewFunctionCall creates new instance of FunctionCall
func NewFunctionCall() *FunctionCall {
	return &FunctionCall{
		AstID:        unique.NewID(),
		ArgumentList: NewArgumentList(),
	}
}

// FunctionCall AST graph node
type FunctionCall struct {
	AstID   string
	GrlText string

	FunctionName string
	ArgumentList *ArgumentList
	Value        reflect.Value
}

// Clone will clone this FunctionCall. The new clone will have an identical structure
func (e *FunctionCall) Clone(cloneTable *pkg.CloneTable) *FunctionCall {
	clone := &FunctionCall{
		AstID:        unique.NewID(),
		GrlText:      e.GrlText,
		FunctionName: e.FunctionName,
	}

	if e.ArgumentList != nil {
		if cloneTable.IsCloned(e.ArgumentList.AstID) {
			clone.ArgumentList = cloneTable.Records[e.ArgumentList.AstID].CloneInstance.(*ArgumentList)
		} else {
			cloned := e.ArgumentList.Clone(cloneTable)
			clone.ArgumentList = cloned
			cloneTable.MarkCloned(e.ArgumentList.AstID, cloned.AstID, e.ArgumentList, cloned)
		}
	}
	return clone
}

// FunctionCallReceiver should be implemented bu AST graph node to receive a FunctionCall AST graph mode
type FunctionCallReceiver interface {
	AcceptFunctionCall(fun *FunctionCall) error
}

// GetAstID get the UUID asigned for this AST graph node
func (e *FunctionCall) GetAstID() string {
	return e.AstID
}

// GetGrlText get the expression syntax related to this graph when it wast constructed
func (e *FunctionCall) GetGrlText() string {
	return e.GrlText
}

// GetSnapshot will create a structure signature or AST graph
func (e *FunctionCall) GetSnapshot() string {
	var buff bytes.Buffer
	buff.WriteString(FUNCTIONCALL)
	buff.WriteString(fmt.Sprintf("(n:%s", e.FunctionName))
	if e.ArgumentList != nil {
		buff.WriteString(",")
		buff.WriteString(e.ArgumentList.GetSnapshot())
	}
	buff.WriteString(")")
	return buff.String()
}

// SetGrlText set the expression syntax related to this graph when it was constructed. Only ANTLR4 listener should
// call this function.
func (e *FunctionCall) SetGrlText(grlText string) {
	e.GrlText = grlText
}

// AcceptArgumentList will accept an ArgumentList AST graph into this ast graph
func (e *FunctionCall) AcceptArgumentList(argList *ArgumentList) error {
	AstLog.Tracef("Method received argument list")
	e.ArgumentList = argList
	return nil
}

// EvaluateArgumentList will evaluate all arguments and ensure it can be passed into function.
func (e *FunctionCall) EvaluateArgumentList(dataContext IDataContext, memory *WorkingMemory) ([]reflect.Value, error) {
	args, err := e.ArgumentList.Evaluate(dataContext, memory)
	if err != nil {
		return nil, err
	}
	if dataContext == nil {
		AstLog.Errorf("Datacontext for function call %s (%s) is nil", e.FunctionName, e.AstID)
	}
	return args, nil
}
