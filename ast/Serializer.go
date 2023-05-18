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
	"encoding/binary"
	"fmt"
	"io"
	"math"
	"reflect"

	"github.com/hyperjumptech/grule-rule-engine/ast/unique"
	"github.com/sirupsen/logrus"
)

// NodeType is to label a Meta information within catalog
type NodeType int

// ValueType will label the datatype when a constant its saved as binary
type ValueType int

const (
	// TypeArgumentList meta type of ArgumentList
	TypeArgumentList NodeType = iota
	// TypeArrayMapSelector meta type of ArrayMapSelector
	TypeArrayMapSelector
	// TypeAssignment meta type of Assigment
	TypeAssignment
	// TypeExpression meta type of Expression
	TypeExpression
	// TypeConstant meta type of Constant
	TypeConstant
	// TypeExpressionAtom meta type of ExpressionAtom
	TypeExpressionAtom
	// TypeFunctionCall meta type of FunctionCall
	TypeFunctionCall
	// TypeRuleEntry meta type of RuleEntry
	TypeRuleEntry
	// TypeThenExpression meta type of ThenExpression
	TypeThenExpression
	// TypeThenExpressionList meta type of ThenExpressionList
	TypeThenExpressionList
	// TypeThenScope meta type of ThenScope
	TypeThenScope
	// TypeVariable meta type of Variable
	TypeVariable
	// TypeWhenScope meta type of WhenScope
	TypeWhenScope

	// TypeString variable type string label
	TypeString ValueType = iota
	// TypeInteger variable type integer label
	TypeInteger
	// TypeFloat variable type float label
	TypeFloat
	// TypeBoolean variable type boolean label
	TypeBoolean

	// Version will be written to the stream and used for compatibility check
	Version = "1.8"
)

// Catalog used to catalog all AST nodes in a KnowledgeBase.
// All nodes will be saved as their Meta information.
// which includes relations between AST Node.
// As RETE algorithm is a prominent aspect of the KnowledgeBase,
// retaining RETE network is very important. The catalog
// provides simple recording of Expression, ExpressionAtoms, Variables
// that capable of supporting the network, enabling the network
// to be saved and reloaded to/from a stream.
//
// This capability alone supposed to store and load huge ruleset
// fast without having to read the rule set from their origin GRL which
// some how a bit expensive due to string parsing and pattern operation
// by ANTLR4
type Catalog struct {
	KnowledgeBaseName               string
	KnowledgeBaseVersion            string
	Data                            map[string]Meta
	MemoryName                      string
	MemoryVersion                   string
	MemoryVariableSnapshotMap       map[string]string
	MemoryExpressionSnapshotMap     map[string]string
	MemoryExpressionAtomSnapshotMap map[string]string
	MemoryExpressionVariableMap     map[string][]string
	MemoryExpressionAtomVariableMap map[string][]string
}

// BuildKnowledgeBase will rebuild a knowledgebase from this Catalog.
// the rebuilt KnowledgeBase is identical to the original KnowledgeBase from
// which this Catalog was built.
func (cat *Catalog) BuildKnowledgeBase() *KnowledgeBase {
	wm := &WorkingMemory{
		Name:                      cat.MemoryName,
		Version:                   cat.MemoryVersion,
		expressionSnapshotMap:     make(map[string]*Expression),
		expressionAtomSnapshotMap: make(map[string]*ExpressionAtom),
		variableSnapshotMap:       make(map[string]*Variable),
		expressionVariableMap:     make(map[*Variable][]*Expression),
		expressionAtomVariableMap: make(map[*Variable][]*ExpressionAtom),
		ID:                        unique.NewID(),
	}
	kb := &KnowledgeBase{
		Name:          cat.KnowledgeBaseName,
		Version:       cat.KnowledgeBaseVersion,
		DataContext:   nil,
		WorkingMemory: wm,
		RuleEntries:   make(map[string]*RuleEntry),
	}
	importTable := make(map[string]Node)

	// creating instances
	for _, meta := range cat.Data {
		switch meta.GetASTType() {
		case TypeArgumentList:
			amet := meta.(*ArgumentListMeta)
			n := &ArgumentList{
				AstID:     amet.AstID,
				GrlText:   amet.GrlText,
				Arguments: nil,
			}
			importTable[meta.GetAstID()] = n
		case TypeArrayMapSelector:
			n := &ArrayMapSelector{
				AstID:      meta.GetAstID(),
				GrlText:    meta.GetGrlText(),
				Expression: nil,
			}
			importTable[meta.GetAstID()] = n
		case TypeAssignment:
			amet := meta.(*AssigmentMeta)
			n := &Assignment{
				AstID:         amet.AstID,
				GrlText:       amet.GrlText,
				Variable:      nil,
				Expression:    nil,
				IsAssign:      amet.IsAssign,
				IsPlusAssign:  amet.IsPlusAssign,
				IsMinusAssign: amet.IsMinusAssign,
				IsDivAssign:   amet.IsDivAssign,
				IsMulAssign:   amet.IsMulAssign,
			}
			importTable[amet.AstID] = n
		case TypeExpression:
			amet := meta.(*ExpressionMeta)
			n := &Expression{
				AstID:            amet.AstID,
				GrlText:          amet.GrlText,
				LeftExpression:   nil,
				RightExpression:  nil,
				SingleExpression: nil,
				ExpressionAtom:   nil,
				Operator:         amet.Operator,
				Negated:          amet.Negated,
			}
			importTable[amet.AstID] = n
		case TypeConstant:
			amet := meta.(*ConstantMeta)
			n := &Constant{
				AstID:    amet.AstID,
				GrlText:  amet.GrlText,
				Snapshot: amet.Snapshot,
				Value:    reflect.Value{},
				IsNil:    amet.IsNil,
			}
			buffer := bytes.NewBuffer(amet.ValueBytes)
			switch amet.ValueType {
			case TypeString:
				length := make([]byte, 8)
				buffer.Read(length)
				dLen := binary.LittleEndian.Uint64(length)
				byteArr := make([]byte, dLen)
				buffer.Read(byteArr)
				n.Value = reflect.ValueOf(string(byteArr))
			case TypeBoolean:
				arr := make([]byte, 1)
				buffer.Read(arr)
				n.Value = reflect.ValueOf(arr[0] == 1)
			case TypeInteger:
				arr := make([]byte, 8)
				buffer.Read(arr)
				n.Value = reflect.ValueOf(int64(binary.LittleEndian.Uint64(arr)))
			case TypeFloat:
				arr := make([]byte, 8)
				buffer.Read(arr)
				bits := binary.LittleEndian.Uint64(arr)
				float := math.Float64frombits(bits)
				n.Value = reflect.ValueOf(float)
			}
			importTable[amet.AstID] = n
		case TypeExpressionAtom:
			amet := meta.(*ExpressionAtomMeta)
			n := &ExpressionAtom{
				AstID:            amet.AstID,
				GrlText:          amet.GrlText,
				VariableName:     amet.VariableName,
				Constant:         nil,
				FunctionCall:     nil,
				Variable:         nil,
				Negated:          amet.Negated,
				ExpressionAtom:   nil,
				ArrayMapSelector: nil,
			}
			importTable[amet.AstID] = n
		case TypeFunctionCall:
			amet := meta.(*FunctionCallMeta)
			n := &FunctionCall{
				AstID:        amet.AstID,
				GrlText:      amet.GrlText,
				FunctionName: amet.FunctionName,
				ArgumentList: nil,
			}
			importTable[amet.AstID] = n
		case TypeRuleEntry:
			amet := meta.(*RuleEntryMeta)
			n := &RuleEntry{
				AstID:           amet.AstID,
				GrlText:         amet.GrlText,
				RuleName:        amet.RuleName,
				RuleDescription: amet.RuleDescription,
				Salience:        amet.Salience,
				WhenScope:       nil,
				ThenScope:       nil,
			}
			importTable[amet.AstID] = n
			kb.RuleEntries[n.RuleName] = n
		case TypeThenExpression:
			amet := meta.(*ThenExpressionMeta)
			n := &ThenExpression{
				AstID:          amet.AstID,
				GrlText:        amet.GrlText,
				Assignment:     nil,
				ExpressionAtom: nil,
			}
			importTable[amet.AstID] = n
		case TypeThenExpressionList:
			amet := meta.(*ThenExpressionListMeta)
			n := &ThenExpressionList{
				AstID:           amet.AstID,
				GrlText:         amet.GrlText,
				ThenExpressions: nil,
			}
			importTable[amet.AstID] = n
		case TypeThenScope:
			amet := meta.(*ThenScopeMeta)
			n := &ThenScope{
				AstID:              amet.AstID,
				GrlText:            amet.GrlText,
				ThenExpressionList: nil,
			}
			importTable[amet.AstID] = n
		case TypeVariable:
			amet := meta.(*VariableMeta)
			n := &Variable{
				AstID:            amet.AstID,
				GrlText:          amet.GrlText,
				Name:             amet.Name,
				Variable:         nil,
				ArrayMapSelector: nil,
			}
			importTable[amet.AstID] = n
		case TypeWhenScope:
			amet := meta.(*WhenScopeMeta)
			n := &WhenScope{
				AstID:      amet.AstID,
				GrlText:    amet.GrlText,
				Expression: nil,
			}
			importTable[amet.AstID] = n
		default:
			panic("Unrecognized meta type")
		}
	}

	// Cross referencing
	for astID, meta := range cat.Data {
		node := importTable[astID]
		switch meta.GetASTType() {
		case TypeArgumentList:
			n := node.(*ArgumentList)
			amet := meta.(*ArgumentListMeta)
			if amet.ArgumentASTIDs != nil && len(amet.ArgumentASTIDs) > 0 {
				n.Arguments = make([]*Expression, len(amet.ArgumentASTIDs))
				for k, v := range amet.ArgumentASTIDs {
					n.Arguments[k] = importTable[v].(*Expression)
				}
			}
		case TypeArrayMapSelector:
			n := node.(*ArrayMapSelector)
			amet := meta.(*ArrayMapSelectorMeta)
			if len(amet.ExpressionID) > 0 {
				n.Expression = importTable[amet.ExpressionID].(*Expression)
			}
		case TypeAssignment:
			n := node.(*Assignment)
			amet := meta.(*AssigmentMeta)
			if len(amet.ExpressionID) > 0 {
				n.Expression = importTable[amet.ExpressionID].(*Expression)
			}
			if len(amet.VariableID) > 0 {
				n.Variable = importTable[amet.VariableID].(*Variable)
			}
		case TypeExpression:
			n := node.(*Expression)
			amet := meta.(*ExpressionMeta)
			if len(amet.LeftExpressionID) > 0 {
				n.LeftExpression = importTable[amet.LeftExpressionID].(*Expression)
			}
			if len(amet.RightExpressionID) > 0 {
				n.RightExpression = importTable[amet.RightExpressionID].(*Expression)
			}
			if len(amet.SingleExpressionID) > 0 {
				n.SingleExpression = importTable[amet.SingleExpressionID].(*Expression)
			}
			if len(amet.ExpressionAtomID) > 0 {
				n.ExpressionAtom = importTable[amet.ExpressionAtomID].(*ExpressionAtom)
			}
		case TypeConstant:
			// nothing todo

		case TypeExpressionAtom:
			n := node.(*ExpressionAtom)
			amet := meta.(*ExpressionAtomMeta)
			if len(amet.ConstantID) > 0 {
				n.Constant = importTable[amet.ConstantID].(*Constant)
			}
			if len(amet.ExpressionAtomID) > 0 {
				n.ExpressionAtom = importTable[amet.ExpressionAtomID].(*ExpressionAtom)
			}
			if len(amet.VariableID) > 0 {
				n.Variable = importTable[amet.VariableID].(*Variable)
			}
			if len(amet.FunctionCallID) > 0 {
				n.FunctionCall = importTable[amet.FunctionCallID].(*FunctionCall)
			}
			if len(amet.ArrayMapSelectorID) > 0 {
				n.ArrayMapSelector = importTable[amet.ArrayMapSelectorID].(*ArrayMapSelector)
			}
		case TypeFunctionCall:
			n := node.(*FunctionCall)
			amet := meta.(*FunctionCallMeta)
			if len(amet.ArgumentListID) > 0 {
				n.ArgumentList = importTable[amet.ArgumentListID].(*ArgumentList)
			}
		case TypeRuleEntry:
			n := node.(*RuleEntry)
			amet := meta.(*RuleEntryMeta)
			if len(amet.WhenScopeID) > 0 {
				n.WhenScope = importTable[amet.WhenScopeID].(*WhenScope)
			}
			if len(amet.ThenScopeID) > 0 {
				n.ThenScope = importTable[amet.ThenScopeID].(*ThenScope)
			}
		case TypeThenExpression:
			n := node.(*ThenExpression)
			amet := meta.(*ThenExpressionMeta)
			if len(amet.AssignmentID) > 0 {
				n.Assignment = importTable[amet.AssignmentID].(*Assignment)
			}
			if len(amet.ExpressionAtomID) > 0 {
				n.ExpressionAtom = importTable[amet.ExpressionAtomID].(*ExpressionAtom)
			}
		case TypeThenExpressionList:
			n := node.(*ThenExpressionList)
			amet := meta.(*ThenExpressionListMeta)
			if amet.ThenExpressionIDs != nil && len(amet.ThenExpressionIDs) > 0 {
				n.ThenExpressions = make([]*ThenExpression, len(amet.ThenExpressionIDs))
				for k, v := range amet.ThenExpressionIDs {
					if node, ok := importTable[v]; ok {
						n.ThenExpressions[k] = node.(*ThenExpression)
					} else {
						logrus.Errorf("then expression with ast id %s not catalogued", v)
					}
				}
			}
		case TypeThenScope:
			n := node.(*ThenScope)
			amet := meta.(*ThenScopeMeta)
			if len(amet.ThenExpressionListID) > 0 {
				n.ThenExpressionList = importTable[amet.ThenExpressionListID].(*ThenExpressionList)
			}
		case TypeVariable:
			n := node.(*Variable)
			amet := meta.(*VariableMeta)
			if len(amet.VariableID) > 0 {
				n.Variable = importTable[amet.VariableID].(*Variable)
			}
			if len(amet.ArrayMapSelectorID) > 0 {
				n.ArrayMapSelector = importTable[amet.ArrayMapSelectorID].(*ArrayMapSelector)
			}
		case TypeWhenScope:
			n := node.(*WhenScope)
			amet := meta.(*WhenScopeMeta)
			if len(amet.ExpressionID) > 0 {
				n.Expression = importTable[amet.ExpressionID].(*Expression)
			}
		default:
			panic("Unrecognized meta type")
		}
	}

	// Rebuilding Working Memory
	if cat.MemoryVariableSnapshotMap != nil && len(cat.MemoryVariableSnapshotMap) > 0 {
		for k, v := range cat.MemoryVariableSnapshotMap {
			if n, ok := importTable[v]; ok {
				wm.variableSnapshotMap[k] = n.(*Variable)
			} else {
				logrus.Warnf("snapshot %s in working memory have no referenced variable with ASTID %s", k, v)
			}
		}
	}
	if cat.MemoryExpressionSnapshotMap != nil && len(cat.MemoryExpressionSnapshotMap) > 0 {
		for k, v := range cat.MemoryExpressionSnapshotMap {
			wm.expressionSnapshotMap[k] = importTable[v].(*Expression)
		}
	}
	if cat.MemoryExpressionAtomSnapshotMap != nil && len(cat.MemoryExpressionAtomSnapshotMap) > 0 {
		for k, v := range cat.MemoryExpressionAtomSnapshotMap {
			wm.expressionAtomSnapshotMap[k] = importTable[v].(*ExpressionAtom)
		}
	}
	if cat.MemoryExpressionVariableMap != nil && len(cat.MemoryExpressionVariableMap) > 0 {
		for k, v := range cat.MemoryExpressionVariableMap {
			variable := importTable[k].(*Variable)
			wm.expressionVariableMap[variable] = make([]*Expression, len(v))
			for i, j := range v {
				wm.expressionVariableMap[variable][i] = importTable[j].(*Expression)
			}
		}
	}
	if cat.MemoryExpressionAtomVariableMap != nil && len(cat.MemoryExpressionAtomVariableMap) > 0 {
		for k, v := range cat.MemoryExpressionAtomVariableMap {
			variable := importTable[k].(*Variable)
			wm.expressionAtomVariableMap[variable] = make([]*ExpressionAtom, len(v))
			for i, j := range v {
				wm.expressionAtomVariableMap[variable][i] = importTable[j].(*ExpressionAtom)
			}
		}
	}

	return kb
}

// Equals used for testing purpose, to ensure that two catalog
// can be compared straight away.
// The comparison is Deep comparison.
func (cat *Catalog) Equals(that *Catalog) bool {
	if cat.KnowledgeBaseName != that.KnowledgeBaseName {
		return false
	}
	if cat.KnowledgeBaseVersion != that.KnowledgeBaseVersion {
		return false
	}
	if cat.MemoryName != that.MemoryName {
		return false
	}
	if cat.MemoryVersion != that.MemoryVersion {
		return false
	}
	if len(cat.Data) != len(that.Data) {
		return false
	}
	for k, v := range cat.Data {
		if j, ok := that.Data[k]; ok {
			if !j.Equals(v) {
				return false
			}
		} else {
			return false
		}
	}
	for k, v := range cat.MemoryVariableSnapshotMap {
		if j, ok := that.MemoryVariableSnapshotMap[k]; ok {
			if j != v {
				return false
			}
		} else {
			return false
		}
	}
	for k, v := range cat.MemoryExpressionSnapshotMap {
		if j, ok := that.MemoryExpressionSnapshotMap[k]; ok {
			if j != v {
				return false
			}
		} else {
			return false
		}
	}
	for k, v := range cat.MemoryExpressionAtomSnapshotMap {
		if j, ok := that.MemoryExpressionAtomSnapshotMap[k]; ok {
			if j != v {
				return false
			}
		} else {
			return false
		}
	}
	for k, v := range cat.MemoryExpressionVariableMap {
		if j, ok := that.MemoryExpressionVariableMap[k]; ok {
			if len(j) != len(v) {
				return false
			}
			for in, st := range v {
				if j[in] != st {
					return false
				}
			}
		} else {
			return false
		}
	}
	for k, v := range cat.MemoryExpressionAtomVariableMap {
		if j, ok := that.MemoryExpressionAtomVariableMap[k]; ok {
			if len(j) != len(v) {
				return false
			}
			for in, st := range v {
				if j[in] != st {
					return false
				}
			}
		} else {
			return false
		}
	}
	return true

}

// ReadCatalogFromReader would read a byte stream from reader
// It will replace all values already sets in a catalog.
// You are responsible for closing the reader stream once its done.
func (cat *Catalog) ReadCatalogFromReader(r io.Reader) error {
	// Read the catalog file version.
	str, err := ReadStringFromReader(r) // V
	if err != nil {
		return err
	}
	if str != Version {
		return fmt.Errorf("invalid version %s", str)
	}

	// Read the knowledgebase name.
	str, err = ReadStringFromReader(r) // V
	if err != nil {
		return err
	}
	cat.KnowledgeBaseName = str

	// Read the knowledgebase version.
	str, err = ReadStringFromReader(r) // V
	if err != nil {
		return err
	}
	cat.KnowledgeBaseVersion = str

	// Writedown meta counts.
	count, err := ReadIntFromReader(r) // V
	if err != nil {
		return err
	}

	cat.Data = make(map[string]Meta)

	for i := uint64(0); i < count; i++ {
		key, err := ReadStringFromReader(r) // V
		if err != nil {
			return err
		}
		metaType, err := ReadIntFromReader(r) // V
		if err != nil {
			return err
		}
		var meta Meta
		switch NodeType(metaType) {
		case TypeArgumentList:
			meta = &ArgumentListMeta{}
		case TypeArrayMapSelector:
			meta = &ArrayMapSelectorMeta{}
		case TypeAssignment:
			meta = &AssigmentMeta{}
		case TypeConstant:
			meta = &ConstantMeta{}
		case TypeExpression:
			meta = &ExpressionMeta{}
		case TypeExpressionAtom:
			meta = &ExpressionAtomMeta{}
		case TypeFunctionCall:
			meta = &FunctionCallMeta{}
		case TypeRuleEntry:
			meta = &RuleEntryMeta{}
		case TypeThenExpression:
			meta = &ThenExpressionMeta{}
		case TypeThenExpressionList:
			meta = &ThenExpressionListMeta{}
		case TypeThenScope:
			meta = &ThenScopeMeta{}
		case TypeVariable:
			meta = &VariableMeta{}
		case TypeWhenScope:
			meta = &WhenScopeMeta{}
		default:
			return fmt.Errorf("unknown meta number %d", metaType)
		}
		err = meta.ReadMetaFrom(r) // V
		if err != nil {
			return err
		}
		cat.Data[key] = meta
	}

	str, err = ReadStringFromReader(r)
	if err != nil {
		return err
	}
	cat.MemoryName = str

	str, err = ReadStringFromReader(r)
	if err != nil {
		return err
	}
	cat.MemoryVersion = str

	// Writedown meta counts.
	count, err = ReadIntFromReader(r)
	if err != nil {
		return err
	}

	cat.MemoryVariableSnapshotMap = make(map[string]string)
	for i := uint64(0); i < count; i++ {
		key, err := ReadStringFromReader(r)
		if err != nil {
			return err
		}
		val, err := ReadStringFromReader(r)
		if err != nil {
			return err
		}
		cat.MemoryVariableSnapshotMap[key] = val
	}

	// MemoryExpressionSnapshotMap meta counts.
	count, err = ReadIntFromReader(r)
	if err != nil {
		return err
	}

	cat.MemoryExpressionSnapshotMap = make(map[string]string)
	for i := uint64(0); i < count; i++ {
		key, err := ReadStringFromReader(r)
		if err != nil {
			return err
		}
		val, err := ReadStringFromReader(r)
		if err != nil {
			return err
		}
		cat.MemoryExpressionSnapshotMap[key] = val
	}

	// MemoryExpressionAtomSnapshotMap meta counts.
	count, err = ReadIntFromReader(r)
	if err != nil {
		return err
	}

	cat.MemoryExpressionAtomSnapshotMap = make(map[string]string)
	for i := uint64(0); i < count; i++ {
		key, err := ReadStringFromReader(r)
		if err != nil {
			return err
		}
		val, err := ReadStringFromReader(r)
		if err != nil {
			return err
		}
		cat.MemoryExpressionAtomSnapshotMap[key] = val
	}

	// MemoryExpressionVariableMap meta counts.
	count, err = ReadIntFromReader(r)
	if err != nil {
		return err
	}

	cat.MemoryExpressionVariableMap = make(map[string][]string)
	for i := uint64(0); i < count; i++ {
		key, err := ReadStringFromReader(r)
		if err != nil {
			return err
		}
		incount, err := ReadIntFromReader(r)
		if err != nil {
			return err
		}
		content := make([]string, incount)
		for j := uint64(0); j < incount; j++ {
			str, err := ReadStringFromReader(r)
			if err != nil {
				return err
			}
			content[j] = str
		}
		cat.MemoryExpressionVariableMap[key] = content
	}

	// MemoryExpressionAtomVariableMap meta counts.
	count, err = ReadIntFromReader(r)
	if err != nil {
		return err
	}

	cat.MemoryExpressionAtomVariableMap = make(map[string][]string)
	for i := uint64(0); i < count; i++ {
		key, err := ReadStringFromReader(r)
		if err != nil {
			return err
		}
		incount, err := ReadIntFromReader(r)
		if err != nil {
			return err
		}
		content := make([]string, incount)
		for j := uint64(0); j < incount; j++ {
			str, err := ReadStringFromReader(r)
			if err != nil {
				return err
			}
			content[j] = str
		}
		cat.MemoryExpressionAtomVariableMap[key] = content
	}

	return nil
}

// WriteCatalogToWriter will store the content of this Catalog
// into a byte stream using provided writer.
// You are responsible for closing the writing stream once its done.
func (cat *Catalog) WriteCatalogToWriter(w io.Writer) error {
	// Write the catalog file version.
	err := WriteStringToWriter(w, Version)
	if err != nil {
		return err
	}

	// Write the knowledgebase name.
	err = WriteStringToWriter(w, cat.KnowledgeBaseName)
	if err != nil {
		return err
	}

	// Write the knowledgebase version.
	err = WriteStringToWriter(w, cat.KnowledgeBaseVersion)
	if err != nil {
		return err
	}

	// Writedown meta counts.
	err = WriteIntToWriter(w, uint64(len(cat.Data)))
	if err != nil {
		return err
	}

	// For each meta.. write them down
	for k, v := range cat.Data {

		// Write the AST ID
		err = WriteStringToWriter(w, k)
		if err != nil {
			return err
		}

		err := WriteIntToWriter(w, uint64(v.GetASTType()))
		if err != nil {
			return err
		}

		// Write the meta
		err = v.WriteMetaTo(w)
		if err != nil {
			return err
		}
	}

	// Write the MemoryName version.
	err = WriteStringToWriter(w, cat.MemoryName)
	if err != nil {
		return err
	}

	// Write the MemoryVersion version.
	err = WriteStringToWriter(w, cat.MemoryVersion)
	if err != nil {
		return err
	}

	// MemoryVariableSnapshotMap meta counts.
	err = WriteIntToWriter(w, uint64(len(cat.MemoryVariableSnapshotMap)))
	if err != nil {
		return err
	}
	for k, v := range cat.MemoryVariableSnapshotMap {
		err = WriteStringToWriter(w, k)
		if err != nil {
			return err
		}
		err = WriteStringToWriter(w, v)
		if err != nil {
			return err
		}
	}

	// MemoryExpressionSnapshotMap meta counts.
	err = WriteIntToWriter(w, uint64(len(cat.MemoryExpressionSnapshotMap)))
	if err != nil {
		return err
	}
	for k, v := range cat.MemoryExpressionSnapshotMap {
		err = WriteStringToWriter(w, k)
		if err != nil {
			return err
		}
		err = WriteStringToWriter(w, v)
		if err != nil {
			return err
		}
	}

	// MemoryExpressionAtomSnapshotMap meta counts.
	err = WriteIntToWriter(w, uint64(len(cat.MemoryExpressionAtomSnapshotMap)))
	if err != nil {
		return err
	}
	for k, v := range cat.MemoryExpressionAtomSnapshotMap {
		err = WriteStringToWriter(w, k)
		if err != nil {
			return err
		}
		err = WriteStringToWriter(w, v)
		if err != nil {
			return err
		}
	}

	// MemoryExpressionVariableMap meta counts.
	err = WriteIntToWriter(w, uint64(len(cat.MemoryExpressionVariableMap)))
	if err != nil {
		return err
	}
	for k, v := range cat.MemoryExpressionVariableMap {
		err = WriteStringToWriter(w, k)
		if err != nil {
			return err
		}
		err = WriteIntToWriter(w, uint64(len(v)))
		if err != nil {
			return err
		}
		for _, j := range v {
			err = WriteStringToWriter(w, j)
			if err != nil {
				return err
			}
		}
	}

	// MemoryExpressionAtomVariableMap meta counts.
	err = WriteIntToWriter(w, uint64(len(cat.MemoryExpressionAtomVariableMap)))
	if err != nil {
		return err
	}
	for k, v := range cat.MemoryExpressionAtomVariableMap {
		err = WriteStringToWriter(w, k)
		if err != nil {
			return err
		}
		err = WriteIntToWriter(w, uint64(len(v)))
		if err != nil {
			return err
		}
		for _, j := range v {
			err = WriteStringToWriter(w, j)
			if err != nil {
				return err
			}
		}
	}

	return nil
}

// AddMeta will add AST Node meta information.
// it will reject duplicated AST ID
func (cat *Catalog) AddMeta(astID string, meta Meta) bool {
	if cat.Data == nil {
		cat.Data = make(map[string]Meta)
	}
	if _, ok := cat.Data[astID]; !ok {
		cat.Data[astID] = meta
		return true
	}
	return false
}

// Meta interface as contract of all AST Node meta information.
type Meta interface {
	GetASTType() NodeType
	GetAstID() string
	GetGrlText() string
	GetSnapshot() string
	WriteMetaTo(writer io.Writer) error
	ReadMetaFrom(reader io.Reader) error
	Equals(that Meta) bool
}

// NodeMeta is a base struct for all ASTNode meta
type NodeMeta struct {
	AstID    string
	GrlText  string
	Snapshot string
}

// GetAstID return the node AST ID
func (meta *NodeMeta) GetAstID() string {
	return meta.AstID
}

// GetGrlText return the node original GRLText, this might not be needed
// but useful for debuging future GRL issue
func (meta *NodeMeta) GetGrlText() string {
	return meta.GrlText
}

// GetSnapshot return the NodeSnapshot, this might not needed but
// could be useful for consistency.
func (meta *NodeMeta) GetSnapshot() string {
	return meta.Snapshot
}

// WriteMetaTo write basic AST Node information meta data into writer.
// One should not use this function directly, unless for testing
// serialization of single ASTNode.
func (meta *NodeMeta) WriteMetaTo(w io.Writer) error {
	// First write the AST ID. this may be redundant.
	err := WriteStringToWriter(w, meta.AstID)
	if err != nil {
		return err
	}
	// Second write the GRL Text.
	err = WriteStringToWriter(w, meta.GrlText)
	if err != nil {
		return err
	}
	// Third write the snapshot. This might be un-necessary.
	err = WriteStringToWriter(w, meta.Snapshot)
	if err != nil {
		return err
	}

	return nil
}

// ReadMetaFrom write basic AST Node information meta data from reader.
// One should not use this function directly, unless for testing
// serialization of single ASTNode.
func (meta *NodeMeta) ReadMetaFrom(reader io.Reader) error {
	str, err := ReadStringFromReader(reader)
	if err != nil {
		return err
	}
	meta.AstID = str

	str, err = ReadStringFromReader(reader)
	if err != nil {
		return err
	}
	meta.GrlText = str

	str, err = ReadStringFromReader(reader)
	if err != nil {
		return err
	}
	meta.Snapshot = str

	return nil
}

// Equals basic function to test equality of two MetaNode
func (meta *NodeMeta) Equals(that Meta) bool {
	if meta.GetAstID() != that.GetAstID() {
		return false
	}
	if meta.GetGrlText() != that.GetGrlText() {
		return false
	}
	if meta.GetSnapshot() != that.GetSnapshot() {
		return false
	}
	return true
}

// ArgumentListMeta meta data for an ArgumentList node
type ArgumentListMeta struct {
	NodeMeta
	ArgumentASTIDs []string
}

// Equals basic function to test equality of two MetaNode
func (meta *ArgumentListMeta) Equals(that Meta) bool {
	if ins, ok := that.(*ArgumentListMeta); ok {
		if !meta.NodeMeta.Equals(that) {
			return false
		}
		if len(meta.ArgumentASTIDs) != len(ins.ArgumentASTIDs) {
			return false
		}
		for k, v := range meta.ArgumentASTIDs {
			if ins.ArgumentASTIDs[k] != v {
				return false
			}
		}
		return true
	}
	return false
}

// GetASTType returns the meta type of this AST Node
func (meta *ArgumentListMeta) GetASTType() NodeType {
	return TypeArgumentList
}

// WriteMetaTo write basic AST Node information meta data into writer.
// One should not use this function directly, unless for testing
// serialization of single ASTNode.
func (meta *ArgumentListMeta) WriteMetaTo(w io.Writer) error {
	err := meta.NodeMeta.WriteMetaTo(w)
	if err != nil {
		return err
	}

	// Write the number of arguments
	err = WriteIntToWriter(w, uint64(len(meta.ArgumentASTIDs)))
	if err != nil {
		return err
	}

	// Write the array content
	for _, v := range meta.ArgumentASTIDs {
		err = WriteStringToWriter(w, v)
		if err != nil {
			return err
		}
	}

	return nil
}

// ReadMetaFrom write basic AST Node information meta data from reader.
// One should not use this function directly, unless for testing
// serialization of single ASTNode.
func (meta *ArgumentListMeta) ReadMetaFrom(r io.Reader) error {
	err := meta.NodeMeta.ReadMetaFrom(r)
	if err != nil {
		return err
	}

	in, err := ReadIntFromReader(r)
	if err != nil {
		return err
	}

	meta.ArgumentASTIDs = make([]string, in)
	for i := uint64(0); i < in; i++ {
		s, err := ReadStringFromReader(r)
		if err != nil {
			return err
		}
		meta.ArgumentASTIDs[i] = s
	}

	return nil
}

// ArrayMapSelectorMeta meta data for an ArrayMapSelector node
type ArrayMapSelectorMeta struct {
	NodeMeta
	ExpressionID string
}

// Equals basic function to test equality of two MetaNode
func (meta *ArrayMapSelectorMeta) Equals(that Meta) bool {
	if ins, ok := that.(*ArrayMapSelectorMeta); ok {
		if !meta.NodeMeta.Equals(that) {
			return false
		}
		if meta.ExpressionID != ins.ExpressionID {
			return false
		}
		return true
	}
	return false
}

// GetASTType returns the meta type of this AST Node
func (meta *ArrayMapSelectorMeta) GetASTType() NodeType {
	return TypeArrayMapSelector
}

// WriteMetaTo write basic AST Node information meta data into writer.
// One should not use this function directly, unless for testing
// serialization of single ASTNode.
func (meta *ArrayMapSelectorMeta) WriteMetaTo(w io.Writer) error {
	err := meta.NodeMeta.WriteMetaTo(w)
	if err != nil {
		return err
	}
	err = WriteStringToWriter(w, meta.ExpressionID)
	if err != nil {
		return err
	}
	return nil
}

// ReadMetaFrom write basic AST Node information meta data from reader.
// One should not use this function directly, unless for testing
// serialization of single ASTNode.
func (meta *ArrayMapSelectorMeta) ReadMetaFrom(r io.Reader) error {
	err := meta.NodeMeta.ReadMetaFrom(r)
	if err != nil {
		return err
	}
	s, err := ReadStringFromReader(r)
	if err != nil {
		return err
	}
	meta.ExpressionID = s
	return nil
}

// AssigmentMeta meta data for an Assigment node
type AssigmentMeta struct {
	NodeMeta
	VariableID    string
	ExpressionID  string
	IsAssign      bool
	IsPlusAssign  bool
	IsMinusAssign bool
	IsDivAssign   bool
	IsMulAssign   bool
}

// Equals basic function to test equality of two MetaNode
func (meta *AssigmentMeta) Equals(that Meta) bool {
	if ins, ok := that.(*AssigmentMeta); ok {
		if !meta.NodeMeta.Equals(that) {
			return false
		}
		if meta.VariableID != ins.VariableID {
			return false
		}
		if meta.ExpressionID != ins.ExpressionID {
			return false
		}
		if meta.IsAssign != ins.IsAssign {
			return false
		}
		if meta.IsPlusAssign != ins.IsPlusAssign {
			return false
		}
		if meta.IsMinusAssign != ins.IsMinusAssign {
			return false
		}
		if meta.IsDivAssign != ins.IsDivAssign {
			return false
		}
		if meta.IsMulAssign != ins.IsMulAssign {
			return false
		}
		return true
	}
	return false
}

// GetASTType returns the meta type of this AST Node
func (meta *AssigmentMeta) GetASTType() NodeType {
	return TypeAssignment
}

// WriteMetaTo write basic AST Node information meta data into writer.
// One should not use this function directly, unless for testing
// serialization of single ASTNode.
func (meta *AssigmentMeta) WriteMetaTo(w io.Writer) error {
	err := meta.NodeMeta.WriteMetaTo(w)
	if err != nil {
		return err
	}

	err = WriteStringToWriter(w, meta.VariableID)
	if err != nil {
		return err
	}
	err = WriteStringToWriter(w, meta.ExpressionID)
	if err != nil {
		return err
	}

	err = WriteBoolToWriter(w, meta.IsAssign)
	if err != nil {
		return err
	}
	err = WriteBoolToWriter(w, meta.IsPlusAssign)
	if err != nil {
		return err
	}
	err = WriteBoolToWriter(w, meta.IsMinusAssign)
	if err != nil {
		return err
	}
	err = WriteBoolToWriter(w, meta.IsDivAssign)
	if err != nil {
		return err
	}
	err = WriteBoolToWriter(w, meta.IsMulAssign)
	if err != nil {
		return err
	}

	return nil
}

// ReadMetaFrom write basic AST Node information meta data from reader.
// One should not use this function directly, unless for testing
// serialization of single ASTNode.
func (meta *AssigmentMeta) ReadMetaFrom(r io.Reader) error {
	err := meta.NodeMeta.ReadMetaFrom(r)
	if err != nil {
		return err
	}

	s, err := ReadStringFromReader(r)
	if err != nil {
		return err
	}
	meta.VariableID = s
	s, err = ReadStringFromReader(r)
	if err != nil {
		return err
	}
	meta.ExpressionID = s

	b, err := ReadBoolFromReader(r)
	if err != nil {
		return err
	}
	meta.IsAssign = b

	b, err = ReadBoolFromReader(r)
	if err != nil {
		return err
	}
	meta.IsPlusAssign = b

	b, err = ReadBoolFromReader(r)
	if err != nil {
		return err
	}
	meta.IsMinusAssign = b

	b, err = ReadBoolFromReader(r)
	if err != nil {
		return err
	}
	meta.IsDivAssign = b

	b, err = ReadBoolFromReader(r)
	if err != nil {
		return err
	}
	meta.IsMulAssign = b

	return nil
}

// ConstantMeta meta data for an Constant node
type ConstantMeta struct {
	NodeMeta

	ValueType  ValueType
	ValueBytes []byte
	IsNil      bool
}

// Equals basic function to test equality of two MetaNode
func (meta *ConstantMeta) Equals(that Meta) bool {
	if ins, ok := that.(*ConstantMeta); ok {
		if !meta.NodeMeta.Equals(that) {
			return false
		}
		if meta.ValueType != ins.ValueType {
			return false
		}
		if meta.IsNil != ins.IsNil {
			return false
		}
		if len(meta.ValueBytes) != len(ins.ValueBytes) {
			return false
		}
		for k, v := range meta.ValueBytes {
			if ins.ValueBytes[k] != v {
				return false
			}
		}
		return true
	}
	return false
}

// GetASTType returns the meta type of this AST Node
func (meta *ConstantMeta) GetASTType() NodeType {
	return TypeConstant
}

// WriteMetaTo write basic AST Node information meta data into writer.
// One should not use this function directly, unless for testing
// serialization of single ASTNode.
func (meta *ConstantMeta) WriteMetaTo(w io.Writer) error {
	err := meta.NodeMeta.WriteMetaTo(w)
	if err != nil {
		return err
	}
	err = WriteIntToWriter(w, uint64(meta.ValueType))
	if err != nil {
		return err
	}
	err = WriteIntToWriter(w, uint64(len(meta.ValueBytes)))
	if err != nil {
		return err
	}
	_, err = w.Write(meta.ValueBytes)
	if err != nil {
		return err
	}
	err = WriteBoolToWriter(w, meta.IsNil)
	if err != nil {
		return err
	}

	return nil
}

// ReadMetaFrom write basic AST Node information meta data from reader.
// One should not use this function directly, unless for testing
// serialization of single ASTNode.
func (meta *ConstantMeta) ReadMetaFrom(r io.Reader) error {
	err := meta.NodeMeta.ReadMetaFrom(r)
	if err != nil {
		return err
	}
	i, err := ReadIntFromReader(r)
	if err != nil {
		return err
	}
	meta.ValueType = ValueType(i)

	length, err := ReadIntFromReader(r)
	if err != nil {
		return err
	}
	byteArr := make([]byte, length)
	n, err := r.Read(byteArr)
	if err != nil {
		return err
	}
	if uint64(n) != length {
		return io.ErrShortBuffer
	}
	meta.ValueBytes = byteArr

	b, err := ReadBoolFromReader(r)
	if err != nil {
		return err
	}
	meta.IsNil = b

	return nil
}

// ExpressionMeta meta data for an Expression node
type ExpressionMeta struct {
	NodeMeta
	LeftExpressionID   string
	RightExpressionID  string
	SingleExpressionID string
	ExpressionAtomID   string
	Operator           int
	Negated            bool
}

// Equals basic function to test equality of two MetaNode
func (meta *ExpressionMeta) Equals(that Meta) bool {
	if ins, ok := that.(*ExpressionMeta); ok {
		if !meta.NodeMeta.Equals(that) {
			return false
		}
		if meta.LeftExpressionID != ins.LeftExpressionID {
			return false
		}
		if meta.RightExpressionID != ins.RightExpressionID {
			return false
		}
		if meta.SingleExpressionID != ins.SingleExpressionID {
			return false
		}
		if meta.ExpressionAtomID != ins.ExpressionAtomID {
			return false
		}
		if meta.Operator != ins.Operator {
			return false
		}
		if meta.Negated != ins.Negated {
			return false
		}
		return true
	}
	return false
}

// GetASTType returns the meta type of this AST Node
func (meta *ExpressionMeta) GetASTType() NodeType {
	return TypeExpression
}

// WriteMetaTo write basic AST Node information meta data into writer.
// One should not use this function directly, unless for testing
// serialization of single ASTNode.
func (meta *ExpressionMeta) WriteMetaTo(w io.Writer) error {
	err := meta.NodeMeta.WriteMetaTo(w)
	if err != nil {
		return err
	}
	err = WriteStringToWriter(w, meta.LeftExpressionID)
	if err != nil {
		return err
	}
	err = WriteStringToWriter(w, meta.RightExpressionID)
	if err != nil {
		return err
	}
	err = WriteStringToWriter(w, meta.SingleExpressionID)
	if err != nil {
		return err
	}
	err = WriteStringToWriter(w, meta.ExpressionAtomID)
	if err != nil {
		return err
	}
	err = WriteIntToWriter(w, uint64(meta.Operator))
	if err != nil {
		return err
	}
	err = WriteBoolToWriter(w, meta.Negated)
	if err != nil {
		return err
	}

	return nil
}

// ReadMetaFrom write basic AST Node information meta data from reader.
// One should not use this function directly, unless for testing
// serialization of single ASTNode.
func (meta *ExpressionMeta) ReadMetaFrom(r io.Reader) error {
	err := meta.NodeMeta.ReadMetaFrom(r)
	if err != nil {
		return err
	}
	s, err := ReadStringFromReader(r)
	if err != nil {
		return err
	}
	meta.LeftExpressionID = s
	s, err = ReadStringFromReader(r)
	if err != nil {
		return err
	}
	meta.RightExpressionID = s
	s, err = ReadStringFromReader(r)
	if err != nil {
		return err
	}
	meta.SingleExpressionID = s
	s, err = ReadStringFromReader(r)
	if err != nil {
		return err
	}
	meta.ExpressionAtomID = s
	i, err := ReadIntFromReader(r)
	if err != nil {
		return err
	}
	meta.Operator = int(i)
	b, err := ReadBoolFromReader(r)
	if err != nil {
		return err
	}
	meta.Negated = b
	return nil
}

// ExpressionAtomMeta meta data for an ExpressionAtom node
type ExpressionAtomMeta struct {
	NodeMeta
	VariableName       string
	ConstantID         string
	FunctionCallID     string
	VariableID         string
	Negated            bool
	ExpressionAtomID   string
	ArrayMapSelectorID string
}

// Equals basic function to test equality of two MetaNode
func (meta *ExpressionAtomMeta) Equals(that Meta) bool {
	if ins, ok := that.(*ExpressionAtomMeta); ok {
		if !meta.NodeMeta.Equals(that) {
			return false
		}
		if meta.VariableName != ins.VariableName {
			return false
		}
		if meta.ConstantID != ins.ConstantID {
			return false
		}
		if meta.FunctionCallID != ins.FunctionCallID {
			return false
		}
		if meta.VariableID != ins.VariableID {
			return false
		}
		if meta.Negated != ins.Negated {
			return false
		}
		if meta.ExpressionAtomID != ins.ExpressionAtomID {
			return false
		}
		if meta.ArrayMapSelectorID != ins.ArrayMapSelectorID {
			return false
		}
		return true
	}
	return false
}

// GetASTType returns the meta type of this AST Node
func (meta *ExpressionAtomMeta) GetASTType() NodeType {
	return TypeExpressionAtom
}

// WriteMetaTo write basic AST Node information meta data into writer.
// One should not use this function directly, unless for testing
// serialization of single ASTNode.
func (meta *ExpressionAtomMeta) WriteMetaTo(w io.Writer) error {
	err := meta.NodeMeta.WriteMetaTo(w)
	if err != nil {
		return err
	}
	err = WriteStringToWriter(w, meta.VariableName)
	if err != nil {
		return err
	}
	err = WriteStringToWriter(w, meta.ConstantID)
	if err != nil {
		return err
	}
	err = WriteStringToWriter(w, meta.FunctionCallID)
	if err != nil {
		return err
	}
	err = WriteStringToWriter(w, meta.VariableID)
	if err != nil {
		return err
	}
	err = WriteBoolToWriter(w, meta.Negated)
	if err != nil {
		return err
	}
	err = WriteStringToWriter(w, meta.ExpressionAtomID)
	if err != nil {
		return err
	}
	err = WriteStringToWriter(w, meta.ArrayMapSelectorID)
	if err != nil {
		return err
	}

	return nil
}

// ReadMetaFrom write basic AST Node information meta data from reader.
// One should not use this function directly, unless for testing
// serialization of single ASTNode.
func (meta *ExpressionAtomMeta) ReadMetaFrom(r io.Reader) error {
	err := meta.NodeMeta.ReadMetaFrom(r)
	if err != nil {
		return err
	}
	s, err := ReadStringFromReader(r)
	if err != nil {
		return err
	}
	meta.VariableName = s
	s, err = ReadStringFromReader(r)
	if err != nil {
		return err
	}
	meta.ConstantID = s
	s, err = ReadStringFromReader(r)
	if err != nil {
		return err
	}
	meta.FunctionCallID = s
	s, err = ReadStringFromReader(r)
	if err != nil {
		return err
	}
	meta.VariableID = s
	b, err := ReadBoolFromReader(r)
	if err != nil {
		return err
	}
	meta.Negated = b
	s, err = ReadStringFromReader(r)
	if err != nil {
		return err
	}
	meta.ExpressionAtomID = s
	s, err = ReadStringFromReader(r)
	if err != nil {
		return err
	}
	meta.ArrayMapSelectorID = s
	return nil
}

// FunctionCallMeta meta data for an FunctionCall node
type FunctionCallMeta struct {
	NodeMeta
	FunctionName   string
	ArgumentListID string
}

// Equals basic function to test equality of two MetaNode
func (meta *FunctionCallMeta) Equals(that Meta) bool {
	if ins, ok := that.(*FunctionCallMeta); ok {
		if !meta.NodeMeta.Equals(that) {
			return false
		}
		if meta.FunctionName != ins.FunctionName {
			return false
		}
		if meta.ArgumentListID != ins.ArgumentListID {
			return false
		}
		return true
	}
	return false
}

// GetASTType returns the meta type of this AST Node
func (meta *FunctionCallMeta) GetASTType() NodeType {
	return TypeFunctionCall
}

// WriteMetaTo write basic AST Node information meta data into writer.
// One should not use this function directly, unless for testing
// serialization of single ASTNode.
func (meta *FunctionCallMeta) WriteMetaTo(w io.Writer) error {
	err := meta.NodeMeta.WriteMetaTo(w)
	if err != nil {
		return err
	}
	err = WriteStringToWriter(w, meta.FunctionName)
	if err != nil {
		return err
	}
	err = WriteStringToWriter(w, meta.ArgumentListID)
	if err != nil {
		return err
	}

	return nil
}

// ReadMetaFrom write basic AST Node information meta data from reader.
// One should not use this function directly, unless for testing
// serialization of single ASTNode.
func (meta *FunctionCallMeta) ReadMetaFrom(r io.Reader) error {
	err := meta.NodeMeta.ReadMetaFrom(r)
	if err != nil {
		return err
	}
	s, err := ReadStringFromReader(r)
	if err != nil {
		return err
	}
	meta.FunctionName = s
	s, err = ReadStringFromReader(r)
	if err != nil {
		return err
	}
	meta.ArgumentListID = s
	return nil
}

// RuleEntryMeta meta data for an RuleEntry node
type RuleEntryMeta struct {
	NodeMeta

	RuleName        string
	RuleDescription string
	Salience        int
	WhenScopeID     string
	ThenScopeID     string
}

// Equals basic function to test equality of two MetaNode
func (meta *RuleEntryMeta) Equals(that Meta) bool {
	if ins, ok := that.(*RuleEntryMeta); ok {
		if !meta.NodeMeta.Equals(that) {
			return false
		}
		if meta.RuleName != ins.RuleName {
			return false
		}
		if meta.RuleDescription != ins.RuleDescription {
			return false
		}
		if meta.Salience != ins.Salience {
			return false
		}
		if meta.WhenScopeID != ins.WhenScopeID {
			return false
		}
		if meta.ThenScopeID != ins.ThenScopeID {
			return false
		}
		return true
	}
	return false
}

// GetASTType returns the meta type of this AST Node
func (meta *RuleEntryMeta) GetASTType() NodeType {
	return TypeRuleEntry
}

// WriteMetaTo write basic AST Node information meta data into writer.
// One should not use this function directly, unless for testing
// serialization of single ASTNode.
func (meta *RuleEntryMeta) WriteMetaTo(w io.Writer) error {
	err := meta.NodeMeta.WriteMetaTo(w)
	if err != nil {
		return err
	}
	err = WriteStringToWriter(w, meta.RuleName)
	if err != nil {
		return err
	}
	err = WriteStringToWriter(w, meta.RuleDescription)
	if err != nil {
		return err
	}
	err = WriteIntToWriter(w, uint64(meta.Salience))
	if err != nil {
		return err
	}
	err = WriteStringToWriter(w, meta.WhenScopeID)
	if err != nil {
		return err
	}
	err = WriteStringToWriter(w, meta.ThenScopeID)
	if err != nil {
		return err
	}

	return nil
}

// ReadMetaFrom write basic AST Node information meta data from reader.
// One should not use this function directly, unless for testing
// serialization of single ASTNode.
func (meta *RuleEntryMeta) ReadMetaFrom(r io.Reader) error {
	err := meta.NodeMeta.ReadMetaFrom(r)
	if err != nil {
		return err
	}
	s, err := ReadStringFromReader(r)
	if err != nil {
		return err
	}
	meta.RuleName = s
	s, err = ReadStringFromReader(r)
	if err != nil {
		return err
	}
	meta.RuleDescription = s
	i, err := ReadIntFromReader(r)
	if err != nil {
		return err
	}
	meta.Salience = int(i)
	s, err = ReadStringFromReader(r)
	if err != nil {
		return err
	}
	meta.WhenScopeID = s
	s, err = ReadStringFromReader(r)
	if err != nil {
		return err
	}
	meta.ThenScopeID = s
	return nil
}

// ThenExpressionMeta meta data for an ThenExpression node
type ThenExpressionMeta struct {
	NodeMeta

	AssignmentID     string
	ExpressionAtomID string
}

// Equals basic function to test equality of two MetaNode
func (meta *ThenExpressionMeta) Equals(that Meta) bool {
	if ins, ok := that.(*ThenExpressionMeta); ok {
		if !meta.NodeMeta.Equals(that) {
			return false
		}
		if meta.AssignmentID != ins.AssignmentID {
			return false
		}
		if meta.ExpressionAtomID != ins.ExpressionAtomID {
			return false
		}
		return true
	}
	return false
}

// GetASTType returns the meta type of this AST Node
func (meta *ThenExpressionMeta) GetASTType() NodeType {
	return TypeThenExpression
}

// WriteMetaTo write basic AST Node information meta data into writer.
// One should not use this function directly, unless for testing
// serialization of single ASTNode.
func (meta *ThenExpressionMeta) WriteMetaTo(w io.Writer) error {
	err := meta.NodeMeta.WriteMetaTo(w)
	if err != nil {
		return err
	}
	err = WriteStringToWriter(w, meta.AssignmentID)
	if err != nil {
		return err
	}
	err = WriteStringToWriter(w, meta.ExpressionAtomID)
	if err != nil {
		return err
	}

	return nil
}

// ReadMetaFrom write basic AST Node information meta data from reader.
// One should not use this function directly, unless for testing
// serialization of single ASTNode.
func (meta *ThenExpressionMeta) ReadMetaFrom(r io.Reader) error {
	err := meta.NodeMeta.ReadMetaFrom(r)
	if err != nil {
		return err
	}
	s, err := ReadStringFromReader(r)
	if err != nil {
		return err
	}
	meta.AssignmentID = s
	s, err = ReadStringFromReader(r)
	if err != nil {
		return err
	}
	meta.ExpressionAtomID = s
	return nil
}

// ThenExpressionListMeta meta data for an ThenExpressionList node
type ThenExpressionListMeta struct {
	NodeMeta

	ThenExpressionIDs []string
}

// Equals basic function to test equality of two MetaNode
func (meta *ThenExpressionListMeta) Equals(that Meta) bool {
	if ins, ok := that.(*ThenExpressionListMeta); ok {
		if !meta.NodeMeta.Equals(that) {
			return false
		}
		if len(meta.ThenExpressionIDs) != len(ins.ThenExpressionIDs) {
			return false
		}
		for k, v := range meta.ThenExpressionIDs {
			if ins.ThenExpressionIDs[k] != v {
				return false
			}
		}
		return true
	}
	return false
}

// GetASTType returns the meta type of this AST Node
func (meta *ThenExpressionListMeta) GetASTType() NodeType {
	return TypeThenExpressionList
}

// WriteMetaTo write basic AST Node information meta data into writer.
// One should not use this function directly, unless for testing
// serialization of single ASTNode.
func (meta *ThenExpressionListMeta) WriteMetaTo(w io.Writer) error {
	err := meta.NodeMeta.WriteMetaTo(w)
	if err != nil {
		return err
	}

	err = WriteIntToWriter(w, uint64(len(meta.ThenExpressionIDs)))
	if err != nil {
		return err
	}

	for _, v := range meta.ThenExpressionIDs {
		err = WriteStringToWriter(w, v)
		if err != nil {
			return err
		}
	}

	return nil
}

// ReadMetaFrom write basic AST Node information meta data from reader.
// One should not use this function directly, unless for testing
// serialization of single ASTNode.
func (meta *ThenExpressionListMeta) ReadMetaFrom(r io.Reader) error {
	err := meta.NodeMeta.ReadMetaFrom(r)
	if err != nil {
		return err
	}

	count, err := ReadIntFromReader(r)
	if err != nil {
		return err
	}

	meta.ThenExpressionIDs = make([]string, count)
	for i := uint64(0); i < count; i++ {
		s, err := ReadStringFromReader(r)
		if err != nil {
			return err
		}
		meta.ThenExpressionIDs[i] = s
	}

	return nil
}

// ThenScopeMeta meta data for an ThenScope node
type ThenScopeMeta struct {
	NodeMeta
	ThenExpressionListID string
}

// Equals basic function to test equality of two MetaNode
func (meta *ThenScopeMeta) Equals(that Meta) bool {
	if ins, ok := that.(*ThenScopeMeta); ok {
		if !meta.NodeMeta.Equals(that) {
			return false
		}
		if meta.ThenExpressionListID != ins.ThenExpressionListID {
			return false
		}
		return true
	}
	return false
}

// GetASTType returns the meta type of this AST Node
func (meta *ThenScopeMeta) GetASTType() NodeType {
	return TypeThenScope
}

// WriteMetaTo write basic AST Node information meta data into writer.
// One should not use this function directly, unless for testing
// serialization of single ASTNode.
func (meta *ThenScopeMeta) WriteMetaTo(w io.Writer) error {
	err := meta.NodeMeta.WriteMetaTo(w)
	if err != nil {
		return err
	}
	err = WriteStringToWriter(w, meta.ThenExpressionListID)
	if err != nil {
		return err
	}

	return nil
}

// ReadMetaFrom write basic AST Node information meta data from reader.
// One should not use this function directly, unless for testing
// serialization of single ASTNode.
func (meta *ThenScopeMeta) ReadMetaFrom(r io.Reader) error {
	err := meta.NodeMeta.ReadMetaFrom(r)
	if err != nil {
		return err
	}
	s, err := ReadStringFromReader(r)
	if err != nil {
		return err
	}
	meta.ThenExpressionListID = s
	return nil
}

// VariableMeta meta data for an Variable node
type VariableMeta struct {
	NodeMeta

	Name               string
	VariableID         string
	ArrayMapSelectorID string
}

// Equals basic function to test equality of two MetaNode
func (meta *VariableMeta) Equals(that Meta) bool {
	if ins, ok := that.(*VariableMeta); ok {
		if !meta.NodeMeta.Equals(that) {
			return false
		}
		if meta.Name != ins.Name {
			return false
		}
		if meta.VariableID != ins.VariableID {
			return false
		}
		if meta.ArrayMapSelectorID != ins.ArrayMapSelectorID {
			return false
		}
		return true
	}
	return false
}

// GetASTType returns the meta type of this AST Node
func (meta *VariableMeta) GetASTType() NodeType {
	return TypeVariable
}

// WriteMetaTo write basic AST Node information meta data into writer.
// One should not use this function directly, unless for testing
// serialization of single ASTNode.
func (meta *VariableMeta) WriteMetaTo(w io.Writer) error {
	err := meta.NodeMeta.WriteMetaTo(w)
	if err != nil {
		return err
	}
	err = WriteStringToWriter(w, meta.Name)
	if err != nil {
		return err
	}
	err = WriteStringToWriter(w, meta.VariableID)
	if err != nil {
		return err
	}
	err = WriteStringToWriter(w, meta.ArrayMapSelectorID)
	if err != nil {
		return err
	}

	return nil
}

// ReadMetaFrom write basic AST Node information meta data from reader.
// One should not use this function directly, unless for testing
// serialization of single ASTNode.
func (meta *VariableMeta) ReadMetaFrom(r io.Reader) error {
	err := meta.NodeMeta.ReadMetaFrom(r)
	if err != nil {
		return err
	}
	s, err := ReadStringFromReader(r)
	if err != nil {
		return err
	}
	meta.Name = s
	s, err = ReadStringFromReader(r)
	if err != nil {
		return err
	}
	meta.VariableID = s
	s, err = ReadStringFromReader(r)
	if err != nil {
		return err
	}
	meta.ArrayMapSelectorID = s
	return nil
}

// WhenScopeMeta meta data for an WhenScope node
type WhenScopeMeta struct {
	NodeMeta
	ExpressionID string
}

// Equals basic function to test equality of two MetaNode
func (meta *WhenScopeMeta) Equals(that Meta) bool {
	if ins, ok := that.(*WhenScopeMeta); ok {
		if !meta.NodeMeta.Equals(that) {
			return false
		}
		if meta.ExpressionID != ins.ExpressionID {
			return false
		}
		return true
	}
	return false
}

// GetASTType returns the meta type of this AST Node
func (meta *WhenScopeMeta) GetASTType() NodeType {
	return TypeWhenScope
}

// WriteMetaTo write basic AST Node information meta data into writer.
// One should not use this function directly, unless for testing
// serialization of single ASTNode.
func (meta *WhenScopeMeta) WriteMetaTo(w io.Writer) error {
	err := meta.NodeMeta.WriteMetaTo(w)
	if err != nil {
		return err
	}
	err = WriteStringToWriter(w, meta.ExpressionID)
	if err != nil {
		return err
	}

	return nil
}

// ReadMetaFrom write basic AST Node information meta data from reader.
// One should not use this function directly, unless for testing
// serialization of single ASTNode.
func (meta *WhenScopeMeta) ReadMetaFrom(r io.Reader) error {
	err := meta.NodeMeta.ReadMetaFrom(r)
	if err != nil {
		return err
	}
	s, err := ReadStringFromReader(r)
	if err != nil {
		return err
	}
	meta.ExpressionID = s
	return nil
}

var (
	// TotalRead counter to track total byte read
	TotalRead = uint64(0)
	// TotalWrite counter to track total bytes written
	TotalWrite = uint64(0)
	// ReadCount read counter
	ReadCount = 0
	// WriteCount write counter
	WriteCount = 0
)

// WriteFull will ensure that a byte array is fully written into writer
func WriteFull(w io.Writer, bytes []byte) (int, error) {
	toWrite := len(bytes)
	written := 0
	for written < toWrite {
		outCount, err := w.Write(bytes[written:])
		if err != nil {
			return written, err
		}
		written += outCount
	}
	return written, nil
}

// WriteStringToWriter write a string into writer.
// the structure is that there's length value written
// prior writing the actual string.
func WriteStringToWriter(w io.Writer, s string) error {
	length := make([]byte, 8)
	data := []byte(s)
	binary.LittleEndian.PutUint64(length, uint64(len(data)))
	c, err := WriteFull(w, length)

	TotalWrite += uint64(c)
	if err != nil {
		return err
	}
	c, err = WriteFull(w, data)
	TotalWrite += uint64(c)
	WriteCount++
	return err
}

// ReadStringFromReader read a string from reader.
func ReadStringFromReader(r io.Reader) (string, error) {
	length := make([]byte, 8)

	c, err := io.ReadFull(r, length)
	TotalRead += uint64(c)
	if err != nil {
		return "", err
	}
	strLen := binary.LittleEndian.Uint64(length)
	strByte := make([]byte, int(strLen))
	c, err = io.ReadFull(r, strByte)
	TotalRead += uint64(c)
	if err != nil {
		return "", err
	}
	ReadCount++
	return string(strByte), nil
}

// WriteIntToWriter write a 64 bit integer into writer.
func WriteIntToWriter(w io.Writer, i uint64) error {
	data := make([]byte, 8)
	binary.LittleEndian.PutUint64(data, i)
	c, err := WriteFull(w, data)
	TotalWrite += uint64(c)
	WriteCount++
	return err
}

// ReadIntFromReader read a 64 bit integer from reader.
func ReadIntFromReader(r io.Reader) (uint64, error) {
	byteArray := make([]byte, 8)
	c, err := io.ReadFull(r, byteArray)
	TotalRead += uint64(c)
	if err != nil {
		return 0, err
	}
	i := binary.LittleEndian.Uint64(byteArray)
	ReadCount++
	return i, nil
}

// WriteBoolToWriter writes a simple boolean into writer
func WriteBoolToWriter(w io.Writer, b bool) error {
	data := make([]byte, 1)
	if b {
		data[0] = 1
	} else {
		data[0] = 0
	}
	c, err := WriteFull(w, data)
	TotalWrite += uint64(c)
	return err
}

// ReadBoolFromReader reads a simple boolean from writer
func ReadBoolFromReader(r io.Reader) (bool, error) {
	byteArray := make([]byte, 1)
	c, err := io.ReadFull(r, byteArray)
	TotalRead += uint64(c)
	if err != nil {
		return false, err
	}
	return byteArray[0] == 1, nil
}

// WriteFloatToWriter write a 64bit float into writer
func WriteFloatToWriter(w io.Writer, f float64) error {
	data := make([]byte, 8)
	binary.LittleEndian.PutUint64(data, math.Float64bits(f))
	c, err := WriteFull(w, data)
	TotalWrite += uint64(c)
	WriteCount++
	return err
}

// ReadFloatFromReader reads a 64bit float from reader
func ReadFloatFromReader(r io.Reader) (float64, error) {
	byteArray := make([]byte, 8)
	c, err := io.ReadFull(r, byteArray)
	TotalRead += uint64(c)
	if err != nil {
		return 0, err
	}
	bits := binary.LittleEndian.Uint64(byteArray)
	float := math.Float64frombits(bits)
	return float, nil
}
