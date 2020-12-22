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
	"fmt"
	"github.com/hyperjumptech/grule-rule-engine/ast/unique"
	"github.com/hyperjumptech/grule-rule-engine/pkg"
	"github.com/sirupsen/logrus"
	"strings"
	"time"
)

// NewWorkingMemory create new instance of WorkingMemory
func NewWorkingMemory(name, version string) *WorkingMemory {
	return &WorkingMemory{
		Name:                      name,
		Version:                   version,
		variableSnapshotMap:       make(map[string]*Variable),
		expressionSnapshotMap:     make(map[string]*Expression),
		expressionAtomSnapshotMap: make(map[string]*ExpressionAtom),
		expressionVariableMap:     make(map[*Variable][]*Expression),
		expressionAtomVariableMap: make(map[*Variable][]*ExpressionAtom),
		ID:                        unique.NewID(),
	}
}

// WorkingMemory handles states of expression evaluation status
type WorkingMemory struct {
	Name                      string
	Version                   string
	expressionSnapshotMap     map[string]*Expression
	expressionAtomSnapshotMap map[string]*ExpressionAtom
	variableSnapshotMap       map[string]*Variable
	expressionVariableMap     map[*Variable][]*Expression
	expressionAtomVariableMap map[*Variable][]*ExpressionAtom
	ID                        string
}

// DebugContent will shows the working memory mapping content
func (e *WorkingMemory) DebugContent() {
	if AstLog.Level <= logrus.DebugLevel {
		for varName, vari := range e.variableSnapshotMap {
			AstLog.Debugf("Variable %s : %s : %s", varName, vari.GrlText, vari.AstID)

			if exprs, ok := e.expressionVariableMap[vari]; ok {
				AstLog.Debugf("  %d expressions", len(exprs))
				for i, ex := range exprs {
					AstLog.Debugf("   expr %d: %s", i, ex.GrlText)
				}
			} else {
				AstLog.Debugf("  no expressions mapped for variable %s", vari.GrlText)
			}
			if expratms, ok := e.expressionAtomVariableMap[vari]; ok {
				AstLog.Debugf("  %d expression atoms", len(expratms))
				for i, ex := range expratms {
					AstLog.Debugf("   expr atm %d: %s", i, ex.GrlText)
				}
			} else {
				AstLog.Debugf("  no expressions atom mapped for variable %s", vari.GrlText)
			}
		}
	}
}

// Equals shallowly equals check this Working Memory against other working memory
func (e *WorkingMemory) Equals(that *WorkingMemory) bool {
	if e.Name != that.Name {
		return false
	}
	if e.Version != that.Version {
		return false
	}
	if len(e.expressionSnapshotMap) != len(that.expressionSnapshotMap) {
		return false
	}
	if len(e.expressionAtomSnapshotMap) != len(that.expressionAtomSnapshotMap) {
		return false
	}
	if len(e.variableSnapshotMap) != len(that.variableSnapshotMap) {
		return false
	}
	if len(e.expressionVariableMap) != len(that.expressionVariableMap) {
		return false
	}
	if len(e.expressionAtomVariableMap) != len(that.expressionAtomVariableMap) {
		return false
	}
	return true
}

// Clone will clone this WorkingMemory. The new clone will have an identical structure
func (e *WorkingMemory) Clone(cloneTable *pkg.CloneTable) *WorkingMemory {
	AstLog.Debugf("Cloning working memory %s:%s", e.Name, e.Version)
	clone := NewWorkingMemory(e.Name, e.Version)

	if e.expressionSnapshotMap != nil {
		AstLog.Debugf("Cloning %d expressionSnapshotMap entries", len(e.expressionSnapshotMap))
		for k, expr := range e.expressionSnapshotMap {
			if cloneTable.IsCloned(expr.AstID) {
				clone.expressionSnapshotMap[k] = cloneTable.Records[expr.AstID].CloneInstance.(*Expression)
			} else {
				panic(fmt.Sprintf("expression  %s is not on the clone table", expr.GrlText))
			}
		}
	}

	if e.expressionAtomSnapshotMap != nil {
		AstLog.Debugf("Cloning %d expressionAtomSnapshotMap entries", len(e.expressionAtomSnapshotMap))
		for k, exprAtm := range e.expressionAtomSnapshotMap {
			if cloneTable.IsCloned(exprAtm.AstID) {
				clone.expressionAtomSnapshotMap[k] = cloneTable.Records[exprAtm.AstID].CloneInstance.(*ExpressionAtom)
			} else {
				panic(fmt.Sprintf("expression atom %s is not on the clone table. ASTID %s", exprAtm.GrlText, exprAtm.AstID))
			}
		}
	}

	if e.variableSnapshotMap != nil {
		AstLog.Debugf("Cloning %d variableSnapshotMap entries", len(e.variableSnapshotMap))
		for k, vari := range e.variableSnapshotMap {
			if cloneTable.IsCloned(vari.AstID) {
				clone.variableSnapshotMap[k] = cloneTable.Records[vari.AstID].CloneInstance.(*Variable)
			} else {
				panic(fmt.Sprintf("variable %s is not on the clone table", vari.GrlText))
			}
		}
	}

	if e.expressionVariableMap != nil {
		AstLog.Debugf("Cloning %d expressionVariableMap entries", len(e.expressionVariableMap))
		for k, exprArr := range e.expressionVariableMap {
			if cloneTable.IsCloned(k.AstID) {
				clonedVari := cloneTable.Records[k.AstID].CloneInstance.(*Variable)
				clone.expressionVariableMap[clonedVari] = make([]*Expression, len(exprArr))
				for k2, expr := range exprArr {
					if cloneTable.IsCloned(expr.AstID) {
						clone.expressionVariableMap[clonedVari][k2] = cloneTable.Records[expr.AstID].CloneInstance.(*Expression)
					} else {
						panic(fmt.Sprintf("expression %s is not on the clone table", expr.GrlText))
					}
				}
			} else {
				panic(fmt.Sprintf("variable %s is not on the clone table", k.GrlText))
			}
		}
	}

	if e.expressionAtomVariableMap != nil {
		AstLog.Debugf("Cloning %d expressionAtomVariableMap entries", len(e.expressionAtomVariableMap))
		for k, exprAtmArr := range e.expressionAtomVariableMap {
			if cloneTable.IsCloned(k.AstID) {
				clonedVari := cloneTable.Records[k.AstID].CloneInstance.(*Variable)
				clone.expressionAtomVariableMap[clonedVari] = make([]*ExpressionAtom, len(exprAtmArr))
				for k2, expr := range exprAtmArr {
					if cloneTable.IsCloned(expr.AstID) {
						clone.expressionAtomVariableMap[clonedVari][k2] = cloneTable.Records[expr.AstID].CloneInstance.(*ExpressionAtom)
					} else {
						panic(fmt.Sprintf("expression atom %s is not on the clone table", expr.GrlText))
					}
				}
			} else {
				panic(fmt.Sprintf("variable %s is not on the clone table", k.GrlText))
			}
		}
	}

	if e.Equals(clone) {
		clone.DebugContent()
		return clone
	}
	panic("Clone not equals the origin.")
}

// IndexVariables will index all expression and expression atoms that contains a speciffic variable name
func (e *WorkingMemory) IndexVariables() {
	if AstLog.Level <= logrus.DebugLevel {
		AstLog.Debugf("Indexing %d expressions, %d expression atoms and %d variables.", len(e.expressionSnapshotMap), len(e.expressionAtomSnapshotMap), len(e.variableSnapshotMap))
	}
	start := time.Now()
	defer func() {
		dur := time.Since(start)
		AstLog.Tracef("Working memory indexing takes %d ms", dur/time.Millisecond)
	}()
	e.expressionVariableMap = make(map[*Variable][]*Expression)
	e.expressionAtomVariableMap = make(map[*Variable][]*ExpressionAtom)

	for varSnapshot, variable := range e.variableSnapshotMap {
		if _, ok := e.expressionVariableMap[variable]; ok == false {
			e.expressionVariableMap[variable] = make([]*Expression, 0)
		}
		if _, ok := e.expressionAtomVariableMap[variable]; ok == false {
			e.expressionAtomVariableMap[variable] = make([]*ExpressionAtom, 0)
		}

		for exprSnapshot, expr := range e.expressionSnapshotMap {
			if strings.Contains(exprSnapshot, varSnapshot) {
				e.expressionVariableMap[variable] = append(e.expressionVariableMap[variable], expr)
			}
		}
		for exprAtmSnapshot, exprAtm := range e.expressionAtomSnapshotMap {
			if strings.Contains(exprAtmSnapshot, varSnapshot) {
				e.expressionAtomVariableMap[variable] = append(e.expressionAtomVariableMap[variable], exprAtm)
			}
		}
	}

	e.DebugContent()

}

// AddExpression will add expression into its map if the expression signature is unique
// if the expression is already in its map, it will return one from the map.
func (e *WorkingMemory) AddExpression(exp *Expression) *Expression {
	snapshot := exp.GetSnapshot()
	if expr, ok := e.expressionSnapshotMap[snapshot]; ok {
		AstLog.Tracef("%s : Ignored Expression Snapshot : %s", e.ID, snapshot)
		return expr
	}
	AstLog.Tracef("%s : Added Expression Snapshot : %s", e.ID, snapshot)
	e.expressionSnapshotMap[snapshot] = exp
	return exp
}

// AddExpressionAtom will add expression atom into its map if the expression signature is unique
// if the expression is already in its map, it will return one from the map.
func (e *WorkingMemory) AddExpressionAtom(exp *ExpressionAtom) *ExpressionAtom {
	snapshot := exp.GetSnapshot()
	if expr, ok := e.expressionAtomSnapshotMap[snapshot]; ok {
		AstLog.Tracef("%s : Ignored ExpressionAtom Snapshot : %s", e.ID, snapshot)
		return expr
	}
	AstLog.Tracef("%s : Added ExpressionAtom Snapshot : %s", e.ID, snapshot)
	e.expressionAtomSnapshotMap[snapshot] = exp
	return exp
}

// AddVariable will add variable into its map if the expression signature is unique
// if the expression is already in its map, it will return one from the map.
func (e *WorkingMemory) AddVariable(vari *Variable) *Variable {
	snapshot := vari.GetSnapshot()
	if v, ok := e.variableSnapshotMap[snapshot]; ok {
		AstLog.Tracef("%s : Ignored Variable Snapshot : %s", e.ID, snapshot)
		return v
	}
	AstLog.Tracef("%s : Added Variable Snapshot : %s", e.ID, snapshot)
	e.variableSnapshotMap[snapshot] = vari
	return vari
}

// Reset will reset the evaluated status of a specific variable if its contains a variable name in its signature.
// Returns true if any expression was reset, false if otherwise
func (e *WorkingMemory) Reset(name string) bool {
	AstLog.Tracef("------- resetting  %s", name)
	for _, vari := range e.variableSnapshotMap {
		if vari.GrlText == name {
			return e.ResetVariable(vari)
		}
	}
	for snap, expr := range e.expressionSnapshotMap {
		if strings.Contains(snap, name) || strings.Contains(expr.GrlText, name) {
			expr.Evaluated = false
		}
	}
	for snap, expr := range e.expressionAtomSnapshotMap {
		if strings.Contains(snap, name) || strings.Contains(expr.GrlText, name) {
			expr.Evaluated = false
		}
	}
	return false
}

// ResetVariable will reset the evaluated status of a specific expression if its contains a variable name in its signature.
// Returns true if any expression was reset, false if otherwise
func (e *WorkingMemory) ResetVariable(variable *Variable) bool {
	AstLog.Tracef("------- resetting variable %s : %s", variable.GrlText, variable.AstID)
	if AstLog.Level == logrus.TraceLevel {
		AstLog.Tracef("%s : Resetting %s", e.ID, variable.GetSnapshot())
	}
	reseted := false
	if arr, ok := e.expressionVariableMap[variable]; ok {
		for _, expr := range arr {
			AstLog.Tracef("------ reset expr : %s", expr.GrlText)
			expr.Evaluated = false
			reseted = true
		}
	} else {
		AstLog.Warnf("No expression to reset for variable %s", variable.GrlText)
	}
	if arr, ok := e.expressionAtomVariableMap[variable]; ok {
		for _, expr := range arr {
			AstLog.Tracef("------ reset expr atm : %s", expr.GrlText)
			expr.Evaluated = false
			reseted = true
		}
	} else {
		AstLog.Warnf("No expression atom to reset for variable %s", variable.GrlText)
	}
	return reseted
}

// ResetAll sets all expression evaluated status to false.
// Returns true if any expression was reset, false if otherwise
func (e *WorkingMemory) ResetAll() bool {
	reseted := false
	for _, expr := range e.expressionSnapshotMap {
		expr.Evaluated = false
		reseted = true
	}
	for _, expr := range e.expressionAtomSnapshotMap {
		expr.Evaluated = false
		reseted = true
	}
	return reseted
}
