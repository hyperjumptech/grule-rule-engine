package ast

import (
	"github.com/google/uuid"
	"strings"
)

// NewWorkingMemory create new instance of WorkingMemory
func NewWorkingMemory() *WorkingMemory {
	return &WorkingMemory{
		ExpressionSnapshotMap: make(map[string]*Expression),
		ExpressionVariableMap: make(map[string][]*Expression),
		ID:                    uuid.New().String(),
	}
}

// WorkingMemory handles states of expression evaluation status
type WorkingMemory struct {
	ExpressionSnapshotMap map[string]*Expression
	ExpressionVariableMap map[string][]*Expression
	ID                    string
}

// IndexVar will index all expression that contains a speciffic variable name
func (wm *WorkingMemory) IndexVar(varName string) {
	if _, ok := wm.ExpressionVariableMap[varName]; ok == false {
		wm.ExpressionVariableMap[varName] = make([]*Expression, 0)
		for snapshot, expr := range wm.ExpressionSnapshotMap {
			if strings.Contains(snapshot, varName) {
				wm.ExpressionVariableMap[varName] = append(wm.ExpressionVariableMap[varName], expr)
			}
		}
	}
}

// Add will add expression into its map if the expression signature is unique
// if the expression is already in its map, it will return one from the map.
func (wm *WorkingMemory) Add(exp *Expression) *Expression {
	if expr, ok := wm.ExpressionSnapshotMap[exp.GetSnapshot()]; ok {
		AstLog.Tracef("%s : Ignored Expression Snapshot : %s", wm.ID, exp.GetSnapshot())
		return expr
	}
	AstLog.Tracef("%s : Added Expression Snapshot : %s", wm.ID, exp.GetSnapshot())
	wm.ExpressionSnapshotMap[exp.GetSnapshot()] = exp
	return exp
}

// Reset will reset the evaluated status of a speciffic expression if its contains a variable name in its signature.
func (wm *WorkingMemory) Reset(variableName string) {
	AstLog.Tracef("%s : Resetting %s", wm.ID, variableName)

	if arr, ok := wm.ExpressionVariableMap[variableName]; ok {
		for _, expr := range arr {
			expr.Evaluated = false
		}
	}
}

// ResetAll sets all expression evaluated status to false.
func (wm *WorkingMemory) ResetAll() {
	for _, expr := range wm.ExpressionSnapshotMap {
		expr.Evaluated = false
	}
}
