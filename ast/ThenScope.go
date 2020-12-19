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
	"github.com/hyperjumptech/grule-rule-engine/ast/unique"
	"github.com/hyperjumptech/grule-rule-engine/pkg"
)

// NewThenScope will create new instance of ThenScope
func NewThenScope() *ThenScope {
	return &ThenScope{
		AstID: unique.NewID(),
	}
}

// ThenScope AST graph node
type ThenScope struct {
	AstID   string
	GrlText string

	ThenExpressionList *ThenExpressionList
}

// ThenScopeReceiver must be implemented by any AST object that will hold a ThenScope
type ThenScopeReceiver interface {
	AcceptThenScope(thenScope *ThenScope) error
}

// Clone will clone this ThenScope. The new clone will have an identical structure
func (e *ThenScope) Clone(cloneTable *pkg.CloneTable) *ThenScope {
	clone := &ThenScope{
		AstID:   unique.NewID(),
		GrlText: e.GrlText,
	}

	if e.ThenExpressionList != nil {
		if cloneTable.IsCloned(e.ThenExpressionList.AstID) {
			clone.ThenExpressionList = cloneTable.Records[e.ThenExpressionList.AstID].CloneInstance.(*ThenExpressionList)
		} else {
			cloned := e.ThenExpressionList.Clone(cloneTable)
			clone.ThenExpressionList = cloned
			cloneTable.MarkCloned(e.ThenExpressionList.AstID, cloned.AstID, e.ThenExpressionList, cloned)
		}
	}

	return clone
}

// AcceptThenExpressionList will accept ThenExpressionList graph into this ThenScope
func (e *ThenScope) AcceptThenExpressionList(list *ThenExpressionList) error {
	e.ThenExpressionList = list
	return nil
}

// GetAstID get the UUID asigned for this AST graph node
func (e *ThenScope) GetAstID() string {
	return e.AstID
}

// GetGrlText get the expression syntax related to this graph when it wast constructed
func (e *ThenScope) GetGrlText() string {
	return e.GrlText
}

// GetSnapshot will create a structure signature or AST graph
func (e *ThenScope) GetSnapshot() string {
	var buff bytes.Buffer
	buff.WriteString(THENSCOPE)
	buff.WriteString("(")
	if e.ThenExpressionList != nil {
		buff.WriteString(e.ThenExpressionList.GetSnapshot())
	}
	buff.WriteString(")")
	return buff.String()
}

// SetGrlText set the expression syntax related to this graph when it was constructed. Only ANTLR4 listener should
// call this function.
func (e *ThenScope) SetGrlText(grlText string) {
	e.GrlText = grlText
}

// Execute will execute this graph in the Then scope
func (e *ThenScope) Execute(dataContext IDataContext, memory *WorkingMemory) error {
	if e.ThenExpressionList == nil {
		AstLog.Warnf("Can not execute nil expression list")
	}
	return e.ThenExpressionList.Execute(dataContext, memory)
}
