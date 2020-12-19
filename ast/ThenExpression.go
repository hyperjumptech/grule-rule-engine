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

// NewThenExpression create new instance of ThenExpression
func NewThenExpression() *ThenExpression {
	return &ThenExpression{
		AstID: unique.NewID(),
	}
}

// ThenExpression AST graph node
type ThenExpression struct {
	AstID   string
	GrlText string

	Assignment     *Assignment
	ExpressionAtom *ExpressionAtom
}

// ThenExpressionReceiver must be implemented by any AST object that will store a Then expression
type ThenExpressionReceiver interface {
	AcceptThenExpression(expr *ThenExpression) error
}

// Clone will clone this ThenExpression. The new clone will have an identical structure
func (e *ThenExpression) Clone(cloneTable *pkg.CloneTable) *ThenExpression {
	clone := &ThenExpression{
		AstID:   unique.NewID(),
		GrlText: e.GrlText,
	}

	if e.Assignment != nil {
		if cloneTable.IsCloned(e.Assignment.AstID) {
			clone.Assignment = cloneTable.Records[e.Assignment.AstID].CloneInstance.(*Assignment)
		} else {
			cloned := e.Assignment.Clone(cloneTable)
			clone.Assignment = cloned
			cloneTable.MarkCloned(e.Assignment.AstID, cloned.AstID, e.Assignment, cloned)
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

	return clone
}

// AcceptAssignment will accept Assignment AST graph into this Then ast graph
func (e *ThenExpression) AcceptAssignment(assignment *Assignment) error {
	e.Assignment = assignment
	return nil
}

// AcceptExpressionAtom will accept an AcceptExpressionAtom AST graph into this ast graph
func (e *ThenExpression) AcceptExpressionAtom(exp *ExpressionAtom) error {
	e.ExpressionAtom = exp
	return nil
}

// GetAstID get the UUID asigned for this AST graph node
func (e *ThenExpression) GetAstID() string {
	return e.AstID
}

// GetGrlText get the expression syntax related to this graph when it wast constructed
func (e *ThenExpression) GetGrlText() string {
	return e.GrlText
}

// GetSnapshot will create a structure signature or AST graph
func (e *ThenExpression) GetSnapshot() string {
	var buff bytes.Buffer
	buff.WriteString(THENEXPRESSION)
	buff.WriteString("(")
	if e.Assignment != nil {
		buff.WriteString(e.Assignment.GetSnapshot())
	}
	if e.ExpressionAtom != nil {
		buff.WriteString(e.ExpressionAtom.GetSnapshot())
	}
	buff.WriteString(")")
	return buff.String()
}

// SetGrlText set the expression syntax related to this graph when it was constructed. Only ANTLR4 listener should
// call this function.
func (e *ThenExpression) SetGrlText(grlText string) {
	e.GrlText = grlText
}

// Execute will execute this graph in the Then scope
func (e *ThenExpression) Execute(dataContext IDataContext, memory *WorkingMemory) error {
	if e.Assignment != nil {
		err := e.Assignment.Execute(dataContext, memory)
		if err != nil {
			AstLog.Errorf("error while executing assignment %s. got %s", e.Assignment.GrlText, err.Error())
		} else {
			AstLog.Debugf("success executing assignment %s", e.Assignment.GrlText)
		}
		return err
	}
	if e.ExpressionAtom != nil {
		_, err := e.ExpressionAtom.Evaluate(dataContext, memory)
		if err != nil {
			AstLog.Errorf("error while executing expression %s. got %s", e.ExpressionAtom.GrlText, err.Error())
			return err
		}
		AstLog.Debugf("success executing ExpressionAtom %s", e.ExpressionAtom.GrlText)
		return nil
	}
	return nil
}
