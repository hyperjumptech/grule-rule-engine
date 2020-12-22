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

// NewThenExpressionList creates new instance of ThenExpressionList
func NewThenExpressionList() *ThenExpressionList {
	return &ThenExpressionList{
		AstID:           unique.NewID(),
		ThenExpressions: make([]*ThenExpression, 0),
	}
}

// ThenExpressionList AST graph node
type ThenExpressionList struct {
	AstID   string
	GrlText string

	ThenExpressions []*ThenExpression
}

// ThenExpressionListReceiver must be implemented by any AST object that hold a ThenExpression list AST object
type ThenExpressionListReceiver interface {
	AcceptThenExpressionList(list *ThenExpressionList) error
}

// AcceptThenExpression will accept ThenExpression AST graph into this ExpressionList
func (e *ThenExpressionList) AcceptThenExpression(expr *ThenExpression) error {
	if e.ThenExpressions == nil {
		e.ThenExpressions = make([]*ThenExpression, 0)
	}
	e.ThenExpressions = append(e.ThenExpressions, expr)
	return nil
}

// Clone will clone this ThenExpressionList. The new clone will have an identical structure
func (e *ThenExpressionList) Clone(cloneTable *pkg.CloneTable) *ThenExpressionList {
	clone := &ThenExpressionList{
		AstID:   unique.NewID(),
		GrlText: e.GrlText,
	}

	if e.ThenExpressions != nil {
		clone.ThenExpressions = make([]*ThenExpression, len(e.ThenExpressions))
		for k, expr := range e.ThenExpressions {
			if cloneTable.IsCloned(expr.AstID) {
				clone.ThenExpressions[k] = cloneTable.Records[expr.AstID].CloneInstance.(*ThenExpression)
			} else {
				cloned := expr.Clone(cloneTable)
				clone.ThenExpressions[k] = cloned
				cloneTable.MarkCloned(expr.AstID, cloned.AstID, expr, cloned)
			}
		}
	}

	return clone
}

// GetAstID get the UUID asigned for this AST graph node
func (e *ThenExpressionList) GetAstID() string {
	return e.AstID
}

// GetGrlText get the expression syntax related to this graph when it wast constructed
func (e *ThenExpressionList) GetGrlText() string {
	return e.GrlText
}

// GetSnapshot will create a structure signature or AST graph
func (e *ThenExpressionList) GetSnapshot() string {
	var buff bytes.Buffer
	buff.WriteString(THENEXPRESSIONLIST)
	buff.WriteString("(")
	if e.ThenExpressions != nil {
		for idx, es := range e.ThenExpressions {
			if idx > 0 {
				buff.WriteString(",")
			}
			buff.WriteString(es.GetSnapshot())
		}
	}
	buff.WriteString(")")
	return buff.String()
}

// SetGrlText set the expression syntax related to this graph when it was constructed. Only ANTLR4 listener should
// call this function.
func (e *ThenExpressionList) SetGrlText(grlText string) {
	e.GrlText = grlText
}

// Execute will execute this graph in the Then scope
func (e *ThenExpressionList) Execute(dataContext IDataContext, memory *WorkingMemory) error {
	for _, es := range e.ThenExpressions {
		err := es.Execute(dataContext, memory)
		if err != nil {
			return err
		}
	}
	return nil
}
