package ast

import (
	"github.com/google/uuid"
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
		ID:                        uuid.New().String(),
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

// Clone will clone this WorkingMemory. The new clone will have an identical structure
func (e WorkingMemory) Clone(cloneTable *pkg.CloneTable) *WorkingMemory {
	clone := NewWorkingMemory(e.Name, e.Version)

	if e.expressionSnapshotMap != nil {
		for k, expr := range e.expressionSnapshotMap {
			if cloneTable.IsCloned(expr.AstID) {
				clone.expressionSnapshotMap[k] = cloneTable.Records[expr.AstID].CloneInstance.(*Expression)
			} else {
				cloned := expr.Clone(cloneTable)
				clone.expressionSnapshotMap[k] = cloned
				cloneTable.MarkCloned(expr.AstID, cloned.AstID, expr, cloned)
			}
		}
	}

	if e.expressionVariableMap != nil {
		for k, exprArr := range e.expressionVariableMap {
			clone.expressionVariableMap[k] = make([]*Expression, len(exprArr))
			for k2, expr := range exprArr {
				if cloneTable.IsCloned(expr.AstID) {
					clone.expressionVariableMap[k][k2] = cloneTable.Records[expr.AstID].CloneInstance.(*Expression)
				} else {
					cloned := expr.Clone(cloneTable)
					clone.expressionVariableMap[k][k2] = cloned
					cloneTable.MarkCloned(expr.AstID, cloned.AstID, expr, cloned)
				}
			}
		}
	}

	return clone
}

// IndexVariables will index all expression and expression atoms that contains a speciffic variable name
func (e *WorkingMemory) IndexVariables() {
	if AstLog.Level <= logrus.DebugLevel {
		AstLog.Debugf("Indexing %d expressions, %d expression atoms and %d variables.", len(e.expressionSnapshotMap), len(e.expressionAtomSnapshotMap), len(e.variableSnapshotMap))
	}
	start := time.Now()
	defer func() {
		dur := time.Since(start)
		AstLog.Infof("Working memory indexing takes %d ms", dur/time.Millisecond)
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

// AddExpression will add expression into its map if the expression signature is unique
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
func (e *WorkingMemory) Reset(varName string) bool {
	for _, vari := range e.variableSnapshotMap {
		if vari.GrlText == varName {
			return e.ResetVariable(vari)
		}
	}
	return false
}

// ResetVariable will reset the evaluated status of a specific expression if its contains a variable name in its signature.
// Returns true if any expression was reset, false if otherwise
func (e *WorkingMemory) ResetVariable(variable *Variable) bool {
	if AstLog.Level == logrus.TraceLevel {
		AstLog.Tracef("%s : Resetting %s", e.ID, variable.GetSnapshot())
	}
	reseted := false
	if arr, ok := e.expressionVariableMap[variable]; ok {
		for _, expr := range arr {
			expr.Evaluated = false
			reseted = true
		}
	}
	if arr, ok := e.expressionAtomVariableMap[variable]; ok {
		for _, expr := range arr {
			expr.Evaluated = false
			reseted = true
		}
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
