package ast

import (
	"github.com/google/uuid"
	"github.com/hyperjumptech/grule-rule-engine/pkg"
	"strings"
)

// NewWorkingMemory create new instance of WorkingMemory
func NewWorkingMemory(name, version string) *WorkingMemory {
	return &WorkingMemory{
		Name:                  name,
		Version:               version,
		ExpressionSnapshotMap: make(map[string]*Expression),
		ExpressionVariableMap: make(map[string][]*Expression),
		ID:                    uuid.New().String(),
	}
}

// WorkingMemory handles states of expression evaluation status
type WorkingMemory struct {
	Name                  string
	Version               string
	ExpressionSnapshotMap map[string]*Expression
	ExpressionVariableMap map[string][]*Expression
	ID                    string
}

// Clone will clone this WorkingMemory. The new clone will have an identical structure
func (e WorkingMemory) Clone(cloneTable *pkg.CloneTable) *WorkingMemory {
	clone := NewWorkingMemory(e.Name, e.Version)

	if e.ExpressionSnapshotMap != nil {
		for k, expr := range e.ExpressionSnapshotMap {
			if cloneTable.IsCloned(expr.AstID) {
				clone.ExpressionSnapshotMap[k] = cloneTable.Records[expr.AstID].CloneInstance.(*Expression)
			} else {
				cloned := expr.Clone(cloneTable)
				clone.ExpressionSnapshotMap[k] = cloned
				cloneTable.MarkCloned(expr.AstID, cloned.AstID, expr, cloned)
			}
		}
	}

	if e.ExpressionVariableMap != nil {
		for k, exprArr := range e.ExpressionVariableMap {
			clone.ExpressionVariableMap[k] = make([]*Expression, len(exprArr))
			for k2, expr := range exprArr {
				if cloneTable.IsCloned(expr.AstID) {
					clone.ExpressionVariableMap[k][k2] = cloneTable.Records[expr.AstID].CloneInstance.(*Expression)
				} else {
					cloned := expr.Clone(cloneTable)
					clone.ExpressionVariableMap[k][k2] = cloned
					cloneTable.MarkCloned(expr.AstID, cloned.AstID, expr, cloned)
				}
			}
		}
	}

	return clone
}

// IndexVar will index all expression that contains a speciffic variable name
func (e *WorkingMemory) IndexVar(varName string) bool {
	indexed := false
	if _, ok := e.ExpressionVariableMap[varName]; ok == false {
		e.ExpressionVariableMap[varName] = make([]*Expression, 0)
		for snapshot, expr := range e.ExpressionSnapshotMap {
			if strings.Contains(snapshot, varName) {
				e.ExpressionVariableMap[varName] = append(e.ExpressionVariableMap[varName], expr)
				indexed = true
			}
		}
	}
	return indexed
}

// Add will add expression into its map if the expression signature is unique
// if the expression is already in its map, it will return one from the map.
func (e *WorkingMemory) Add(exp *Expression) (*Expression, bool) {
	if expr, ok := e.ExpressionSnapshotMap[exp.GetSnapshot()]; ok {
		AstLog.Tracef("%s : Ignored Expression Snapshot : %s", e.ID, exp.GetSnapshot())
		return expr, false
	}
	AstLog.Tracef("%s : Added Expression Snapshot : %s", e.ID, exp.GetSnapshot())
	e.ExpressionSnapshotMap[exp.GetSnapshot()] = exp
	return exp, true
}

// Reset will reset the evaluated status of a speciffic expression if its contains a variable name in its signature.
// Returns true if any expression was reset, false if otherwise
func (e *WorkingMemory) Reset(variableName string) bool {
	AstLog.Tracef("%s : Resetting %s", e.ID, variableName)
	reseted := false
	if arr, ok := e.ExpressionVariableMap[variableName]; ok {
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
	for _, expr := range e.ExpressionSnapshotMap {
		expr.Evaluated = false
		reseted = true
	}
	return reseted
}
