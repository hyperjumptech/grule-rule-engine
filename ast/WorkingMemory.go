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
	"github.com/hyperjumptech/grule-rule-engine/logger"
	"github.com/hyperjumptech/grule-rule-engine/pkg"
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

// MakeCatalog create a catalog entry of this working memory
func (workingMem *WorkingMemory) MakeCatalog(cat *Catalog) {
	cat.MemoryName = workingMem.Name
	cat.MemoryVersion = workingMem.Version
	cat.MemoryExpressionSnapshotMap = make(map[string]string)
	for key, value := range workingMem.expressionSnapshotMap {
		cat.MemoryExpressionSnapshotMap[key] = value.AstID
	}
	cat.MemoryExpressionAtomSnapshotMap = make(map[string]string)
	for key, value := range workingMem.expressionAtomSnapshotMap {
		cat.MemoryExpressionAtomSnapshotMap[key] = value.AstID
	}
	cat.MemoryVariableSnapshotMap = make(map[string]string)
	for key, value := range workingMem.variableSnapshotMap {
		cat.MemoryVariableSnapshotMap[key] = value.AstID
	}
	cat.MemoryExpressionVariableMap = make(map[string][]string)
	for key, value := range workingMem.expressionVariableMap {
		cat.MemoryExpressionVariableMap[key.AstID] = make([]string, len(value))
		for i, j := range value {
			cat.MemoryExpressionVariableMap[key.AstID][i] = j.AstID
		}
	}
	cat.MemoryExpressionAtomVariableMap = make(map[string][]string)
	for key, value := range workingMem.expressionAtomVariableMap {
		cat.MemoryExpressionAtomVariableMap[key.AstID] = make([]string, len(value))
		for i, j := range value {
			cat.MemoryExpressionAtomVariableMap[key.AstID][i] = j.AstID
		}
	}
}

// DebugContent will shows the working memory mapping content
func (workingMem *WorkingMemory) DebugContent() {
	if AstLog.Level <= logger.DebugLevel {
		for varName, vari := range workingMem.variableSnapshotMap {
			AstLog.Debugf("Variable %s : %s : %s", varName, vari.GrlText, vari.AstID)

			if exprs, ok := workingMem.expressionVariableMap[vari]; ok {
				AstLog.Debugf("  %d expressions", len(exprs))
				for i, ex := range exprs {
					AstLog.Debugf("   expr %d: %s", i, ex.GrlText)
				}
			} else {
				AstLog.Debugf("  no expressions mapped for variable %s", vari.GrlText)
			}
			if expratms, ok := workingMem.expressionAtomVariableMap[vari]; ok {
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
func (workingMem *WorkingMemory) Equals(that *WorkingMemory) bool {
	if workingMem.Name != that.Name {

		return false
	}
	if workingMem.Version != that.Version {

		return false
	}
	if len(workingMem.expressionSnapshotMap) != len(that.expressionSnapshotMap) {

		return false
	}
	if len(workingMem.expressionAtomSnapshotMap) != len(that.expressionAtomSnapshotMap) {

		return false
	}
	if len(workingMem.variableSnapshotMap) != len(that.variableSnapshotMap) {

		return false
	}
	if len(workingMem.expressionVariableMap) != len(that.expressionVariableMap) {

		return false
	}
	if len(workingMem.expressionAtomVariableMap) != len(that.expressionAtomVariableMap) {

		return false
	}

	return true
}

// Clone will clone this WorkingMemory. The new clone will have an identical structure
func (workingMem *WorkingMemory) Clone(cloneTable *pkg.CloneTable) (*WorkingMemory, error) {
	AstLog.Debugf("Cloning working memory %s:%s", workingMem.Name, workingMem.Version)
	clone := NewWorkingMemory(workingMem.Name, workingMem.Version)

	if workingMem.expressionSnapshotMap != nil {
		AstLog.Debugf("Cloning %d expressionSnapshotMap entries", len(workingMem.expressionSnapshotMap))
		for k, expr := range workingMem.expressionSnapshotMap {
			if cloneTable.IsCloned(expr.AstID) {
				clone.expressionSnapshotMap[k] = cloneTable.Records[expr.AstID].CloneInstance.(*Expression)
			} else {

				return nil, fmt.Errorf("expression  %s is not on the clone table - %s", expr.GrlText, expr.GetSnapshot())
			}
		}
	}

	if workingMem.expressionAtomSnapshotMap != nil {
		AstLog.Debugf("Cloning %d expressionAtomSnapshotMap entries", len(workingMem.expressionAtomSnapshotMap))
		for k, exprAtm := range workingMem.expressionAtomSnapshotMap {
			if cloneTable.IsCloned(exprAtm.AstID) {
				clone.expressionAtomSnapshotMap[k] = cloneTable.Records[exprAtm.AstID].CloneInstance.(*ExpressionAtom)
			} else {

				return nil, fmt.Errorf("expression atom %s is not on the clone table. ASTID %s", exprAtm.GrlText, exprAtm.AstID)
			}
		}
	}

	if workingMem.variableSnapshotMap != nil {
		AstLog.Debugf("Cloning %d variableSnapshotMap entries", len(workingMem.variableSnapshotMap))
		for key, variable := range workingMem.variableSnapshotMap {
			if cloneTable.IsCloned(variable.AstID) {
				clone.variableSnapshotMap[key] = cloneTable.Records[variable.AstID].CloneInstance.(*Variable)
			} else {

				panic(fmt.Sprintf("variable %s is not on the clone table", variable.GrlText))
			}
		}
	}

	if workingMem.expressionVariableMap != nil {
		AstLog.Debugf("Cloning %d expressionVariableMap entries", len(workingMem.expressionVariableMap))
		for key, exprArr := range workingMem.expressionVariableMap {
			if cloneTable.IsCloned(key.AstID) {
				clonedVari := cloneTable.Records[key.AstID].CloneInstance.(*Variable)
				clone.expressionVariableMap[clonedVari] = make([]*Expression, len(exprArr))
				for k2, expr := range exprArr {
					if cloneTable.IsCloned(expr.AstID) {
						clone.expressionVariableMap[clonedVari][k2] = cloneTable.Records[expr.AstID].CloneInstance.(*Expression)
					} else {

						panic(fmt.Sprintf("expression %s is not on the clone table", expr.GrlText))
					}
				}
			} else {

				panic(fmt.Sprintf("variable %s is not on the clone table", key.GrlText))
			}
		}
	}

	if workingMem.expressionAtomVariableMap != nil {
		AstLog.Debugf("Cloning %d expressionAtomVariableMap entries", len(workingMem.expressionAtomVariableMap))
		for key, exprAtmArr := range workingMem.expressionAtomVariableMap {
			if cloneTable.IsCloned(key.AstID) {
				clonedVari := cloneTable.Records[key.AstID].CloneInstance.(*Variable)
				clone.expressionAtomVariableMap[clonedVari] = make([]*ExpressionAtom, len(exprAtmArr))
				for k2, expr := range exprAtmArr {
					if cloneTable.IsCloned(expr.AstID) {
						clone.expressionAtomVariableMap[clonedVari][k2] = cloneTable.Records[expr.AstID].CloneInstance.(*ExpressionAtom)
					} else {

						panic(fmt.Sprintf("expression atom %s is not on the clone table", expr.GrlText))
					}
				}
			} else {

				return nil, fmt.Errorf("variable %s is not on the clone table", key.GrlText)
			}
		}
	}

	if workingMem.Equals(clone) {
		clone.DebugContent()

		return clone, nil
	}

	return nil, fmt.Errorf("clone not equals the origin")
}

// IndexVariables will index all expression and expression atoms that contains a speciffic variable name
func (workingMem *WorkingMemory) IndexVariables() {
	if AstLog.Level <= logger.DebugLevel {
		AstLog.Debugf("Indexing %d expressions, %d expression atoms and %d variables.", len(workingMem.expressionSnapshotMap), len(workingMem.expressionAtomSnapshotMap), len(workingMem.variableSnapshotMap))
	}
	start := time.Now()
	defer func() {
		dur := time.Since(start)
		AstLog.Tracef("Working memory indexing takes %d ms", dur/time.Millisecond)
	}()
	workingMem.expressionVariableMap = make(map[*Variable][]*Expression)
	workingMem.expressionAtomVariableMap = make(map[*Variable][]*ExpressionAtom)

	for varSnapshot, variable := range workingMem.variableSnapshotMap {
		if _, ok := workingMem.expressionVariableMap[variable]; ok == false {
			workingMem.expressionVariableMap[variable] = make([]*Expression, 0)
		}
		if _, ok := workingMem.expressionAtomVariableMap[variable]; ok == false {
			workingMem.expressionAtomVariableMap[variable] = make([]*ExpressionAtom, 0)
		}

		for exprSnapshot, expr := range workingMem.expressionSnapshotMap {
			if strings.Contains(exprSnapshot, varSnapshot) {
				workingMem.expressionVariableMap[variable] = append(workingMem.expressionVariableMap[variable], expr)
			}
		}
		for exprAtmSnapshot, exprAtm := range workingMem.expressionAtomSnapshotMap {
			if strings.Contains(exprAtmSnapshot, varSnapshot) {
				workingMem.expressionAtomVariableMap[variable] = append(workingMem.expressionAtomVariableMap[variable], exprAtm)
			}
		}
	}

	workingMem.DebugContent()

}

// AddExpression will add expression into its map if the expression signature is unique
// if the expression is already in its map, it will return one from the map.
func (workingMem *WorkingMemory) AddExpression(exp *Expression) *Expression {
	snapshot := exp.GetSnapshot()
	if expr, ok := workingMem.expressionSnapshotMap[snapshot]; ok {
		AstLog.Tracef("%s : Ignored Expression Snapshot : %s", workingMem.ID, snapshot)

		return expr
	}
	AstLog.Tracef("%s : Added Expression Snapshot : %s", workingMem.ID, snapshot)
	workingMem.expressionSnapshotMap[snapshot] = exp

	return exp
}

// AddExpressionAtom will add expression atom into its map if the expression signature is unique
// if the expression is already in its map, it will return one from the map.
func (workingMem *WorkingMemory) AddExpressionAtom(exp *ExpressionAtom) *ExpressionAtom {
	snapshot := exp.GetSnapshot()
	if expr, ok := workingMem.expressionAtomSnapshotMap[snapshot]; ok {
		AstLog.Tracef("%s : Ignored ExpressionAtom Snapshot : %s", workingMem.ID, snapshot)

		return expr
	}
	AstLog.Tracef("%s : Added ExpressionAtom Snapshot : %s", workingMem.ID, snapshot)
	workingMem.expressionAtomSnapshotMap[snapshot] = exp

	return exp
}

// AddVariable will add variable into its map if the expression signature is unique
// if the expression is already in its map, it will return one from the map.
func (workingMem *WorkingMemory) AddVariable(vari *Variable) *Variable {
	snapshot := vari.GetSnapshot()
	if v, ok := workingMem.variableSnapshotMap[snapshot]; ok {
		AstLog.Tracef("%s : Ignored Variable Snapshot : %s", workingMem.ID, snapshot)

		return v
	}
	AstLog.Tracef("%s : Added Variable Snapshot : %s", workingMem.ID, snapshot)
	workingMem.variableSnapshotMap[snapshot] = vari

	return vari
}

// Reset will reset the evaluated status of a specific variable if its contains a variable name in its signature.
// Returns true if any expression was reset, false if otherwise
func (workingMem *WorkingMemory) Reset(name string) bool {
	AstLog.Tracef("------- resetting  %s", name)
	for _, vari := range workingMem.variableSnapshotMap {
		if vari.GrlText == name {

			return workingMem.ResetVariable(vari)
		}
	}
	for snap, expr := range workingMem.expressionSnapshotMap {
		if strings.Contains(snap, name) || strings.Contains(expr.GrlText, name) {
			expr.Evaluated = false
		}
	}
	for snap, expr := range workingMem.expressionAtomSnapshotMap {
		if strings.Contains(snap, name) || strings.Contains(expr.GrlText, name) {
			expr.Evaluated = false
		}
	}

	return false
}

// ResetVariable will reset the evaluated status of a specific expression if its contains a variable name in its signature.
// Returns true if any expression was reset, false if otherwise
func (workingMem *WorkingMemory) ResetVariable(variable *Variable) bool {
	AstLog.Tracef("------- resetting variable %s : %s", variable.GrlText, variable.AstID)
	if AstLog.Level == logger.TraceLevel {
		AstLog.Tracef("%s : Resetting %s", workingMem.ID, variable.GetSnapshot())
	}
	reseted := false
	if arr, ok := workingMem.expressionVariableMap[variable]; ok {
		for _, expr := range arr {
			AstLog.Tracef("------ reset expr : %s", expr.GrlText)
			expr.Evaluated = false
			reseted = true
		}
	} else {
		AstLog.Warnf("No expression to reset for variable %s", variable.GrlText)
	}
	if arr, ok := workingMem.expressionAtomVariableMap[variable]; ok {
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
func (workingMem *WorkingMemory) ResetAll() bool {
	reseted := false
	for _, expr := range workingMem.expressionSnapshotMap {
		expr.Evaluated = false
		reseted = true
	}
	for _, expr := range workingMem.expressionAtomSnapshotMap {
		expr.Evaluated = false
		reseted = true
	}

	return reseted
}
