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
	"github.com/hyperjumptech/grule-rule-engine/ast/unique"
	"reflect"

	"github.com/hyperjumptech/grule-rule-engine/pkg"
)

// NewWhenScope creates new instance of WhenScope
func NewWhenScope() *WhenScope {
	return &WhenScope{
		AstID: unique.NewID(),
	}
}

// WhenScope AST graph node
type WhenScope struct {
	AstID   string
	GrlText string

	Expression *Expression
}

// WhenScopeReceiver must be implemented by AST object that stores WhenScope
type WhenScopeReceiver interface {
	AcceptWhenScope(whenScope *WhenScope) error
}

// Clone will clone this Clone. The new clone will have an identical structure
func (e *WhenScope) Clone(cloneTable *pkg.CloneTable) *WhenScope {
	clone := &WhenScope{
		AstID:   unique.NewID(),
		GrlText: e.GrlText,
	}

	if e.Expression != nil {
		if cloneTable.IsCloned(e.Expression.AstID) {
			clone.Expression = cloneTable.Records[e.Expression.AstID].CloneInstance.(*Expression)
		} else {
			cloned := e.Expression.Clone(cloneTable)
			clone.Expression = cloned
			cloneTable.MarkCloned(e.Expression.AstID, cloned.AstID, e.Expression, cloned)
		}
	}

	return clone
}

// AcceptExpression will accept Expression AST graph node into this node
func (e *WhenScope) AcceptExpression(exp *Expression) error {
	if e.Expression == nil {
		e.Expression = exp
		return nil
	}
	return errors.New("expression for when scope already assigned")
}

// GetAstID get the UUID asigned for this AST graph node
func (e *WhenScope) GetAstID() string {
	return e.AstID
}

// GetGrlText get the expression syntax related to this graph when it wast constructed
func (e *WhenScope) GetGrlText() string {
	return e.GrlText
}

// GetSnapshot will create a structure signature or AST graph
func (e *WhenScope) GetSnapshot() string {
	var buff bytes.Buffer
	buff.WriteString(WHENSCOPE)
	buff.WriteString("(")
	if e.Expression != nil {
		buff.WriteString(e.Expression.GetSnapshot())
	}
	buff.WriteString(")")
	return buff.String()
}

// SetGrlText set the expression syntax related to this graph when it was constructed. Only ANTLR4 listener should
// call this function.
func (e *WhenScope) SetGrlText(grlText string) {
	e.GrlText = grlText
}

// Evaluate will evaluate this AST graph for when scope evaluation
func (e *WhenScope) Evaluate(dataContext IDataContext, memory *WorkingMemory) (reflect.Value, error) {
	return e.Expression.Evaluate(dataContext, memory)
}
