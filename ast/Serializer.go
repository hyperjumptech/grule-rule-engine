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
	"github.com/hyperjumptech/grule-rule-engine/ast/unique"
	"github.com/sirupsen/logrus"
	"io"
	"math"
	"reflect"
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
func (cat *Catalog) BuildKnowledgeBase() (*KnowledgeBase, error) {
	workingMem := &WorkingMemory{
		Name:                      cat.MemoryName,
		Version:                   cat.MemoryVersion,
		expressionSnapshotMap:     make(map[string]*Expression),
		expressionAtomSnapshotMap: make(map[string]*ExpressionAtom),
		variableSnapshotMap:       make(map[string]*Variable),
		expressionVariableMap:     make(map[*Variable][]*Expression),
		expressionAtomVariableMap: make(map[*Variable][]*ExpressionAtom),
		ID:                        unique.NewID(),
	}
	knowledgeBase := &KnowledgeBase{
		Name:          cat.KnowledgeBaseName,
		Version:       cat.KnowledgeBaseVersion,
		DataContext:   nil,
		WorkingMemory: workingMem,
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
			arrayMapSelect := &ArrayMapSelector{
				AstID:      meta.GetAstID(),
				GrlText:    meta.GetGrlText(),
				Expression: nil,
			}
			importTable[meta.GetAstID()] = arrayMapSelect
		case TypeAssignment:
			amet := meta.(*AssigmentMeta)
			assignment := &Assignment{
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
			importTable[amet.AstID] = assignment
		case TypeExpression:
			amet := meta.(*ExpressionMeta)
			expression := &Expression{
				AstID:            amet.AstID,
				GrlText:          amet.GrlText,
				LeftExpression:   nil,
				RightExpression:  nil,
				SingleExpression: nil,
				ExpressionAtom:   nil,
				Operator:         amet.Operator,
				Negated:          amet.Negated,
			}
			importTable[amet.AstID] = expression
		case TypeConstant:
			amet := meta.(*ConstantMeta)
			newConst := &Constant{
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
				newConst.Value = reflect.ValueOf(string(byteArr))
			case TypeBoolean:
				arr := make([]byte, 1)
				buffer.Read(arr)
				newConst.Value = reflect.ValueOf(arr[0] == 1)
			case TypeInteger:
				arr := make([]byte, 8)
				buffer.Read(arr)
				newConst.Value = reflect.ValueOf(int64(binary.LittleEndian.Uint64(arr)))
			case TypeFloat:
				arr := make([]byte, 8)
				buffer.Read(arr)
				bits := binary.LittleEndian.Uint64(arr)
				float := math.Float64frombits(bits)
				newConst.Value = reflect.ValueOf(float)
			}
			importTable[amet.AstID] = newConst
		case TypeExpressionAtom:
			amet := meta.(*ExpressionAtomMeta)
			expressionAtm := &ExpressionAtom{
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
			importTable[amet.AstID] = expressionAtm
		case TypeFunctionCall:
			amet := meta.(*FunctionCallMeta)
			funcCall := &FunctionCall{
				AstID:        amet.AstID,
				GrlText:      amet.GrlText,
				FunctionName: amet.FunctionName,
				ArgumentList: nil,
			}
			importTable[amet.AstID] = funcCall
		case TypeRuleEntry:
			amet := meta.(*RuleEntryMeta)
			ruleEntry := &RuleEntry{
				AstID:           amet.AstID,
				GrlText:         amet.GrlText,
				RuleName:        amet.RuleName,
				RuleDescription: amet.RuleDescription,
				Salience:        amet.Salience,
				WhenScope:       nil,
				ThenScope:       nil,
			}
			importTable[amet.AstID] = ruleEntry
			knowledgeBase.RuleEntries[ruleEntry.RuleName] = ruleEntry
		case TypeThenExpression:
			amet := meta.(*ThenExpressionMeta)
			thenExp := &ThenExpression{
				AstID:          amet.AstID,
				GrlText:        amet.GrlText,
				Assignment:     nil,
				ExpressionAtom: nil,
			}
			importTable[amet.AstID] = thenExp
		case TypeThenExpressionList:
			amet := meta.(*ThenExpressionListMeta)
			thenExprList := &ThenExpressionList{
				AstID:           amet.AstID,
				GrlText:         amet.GrlText,
				ThenExpressions: nil,
			}
			importTable[amet.AstID] = thenExprList
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
			variable := &Variable{
				AstID:            amet.AstID,
				GrlText:          amet.GrlText,
				Name:             amet.Name,
				Variable:         nil,
				ArrayMapSelector: nil,
			}
			importTable[amet.AstID] = variable
		case TypeWhenScope:
			amet := meta.(*WhenScopeMeta)
			n := &WhenScope{
				AstID:      amet.AstID,
				GrlText:    amet.GrlText,
				Expression: nil,
			}
			importTable[amet.AstID] = n
		default:
			return nil, fmt.Errorf("unrecognized meta type")
		}
	}

	// Cross referencing
	for astID, meta := range cat.Data {
		node := importTable[astID]
		switch meta.GetASTType() {
		case TypeArgumentList:
			argList := node.(*ArgumentList)
			amet := meta.(*ArgumentListMeta)
			if amet.ArgumentASTIDs != nil && len(amet.ArgumentASTIDs) > 0 {
				argList.Arguments = make([]*Expression, len(amet.ArgumentASTIDs))
				for k, v := range amet.ArgumentASTIDs {
					argList.Arguments[k] = importTable[v].(*Expression)
				}
			}
		case TypeArrayMapSelector:
			arrayMapSel := node.(*ArrayMapSelector)
			amet := meta.(*ArrayMapSelectorMeta)
			if len(amet.ExpressionID) > 0 {
				arrayMapSel.Expression = importTable[amet.ExpressionID].(*Expression)
			}
		case TypeAssignment:
			assignment := node.(*Assignment)
			amet := meta.(*AssigmentMeta)
			if len(amet.ExpressionID) > 0 {
				assignment.Expression = importTable[amet.ExpressionID].(*Expression)
			}
			if len(amet.VariableID) > 0 {
				assignment.Variable = importTable[amet.VariableID].(*Variable)
			}
		case TypeExpression:
			expr := node.(*Expression)
			amet := meta.(*ExpressionMeta)
			if len(amet.LeftExpressionID) > 0 {
				expr.LeftExpression = importTable[amet.LeftExpressionID].(*Expression)
			}
			if len(amet.RightExpressionID) > 0 {
				expr.RightExpression = importTable[amet.RightExpressionID].(*Expression)
			}
			if len(amet.SingleExpressionID) > 0 {
				expr.SingleExpression = importTable[amet.SingleExpressionID].(*Expression)
			}
			if len(amet.ExpressionAtomID) > 0 {
				expr.ExpressionAtom = importTable[amet.ExpressionAtomID].(*ExpressionAtom)
			}
		case TypeConstant:
			// nothing todo

		case TypeExpressionAtom:
			expressAtm := node.(*ExpressionAtom)
			amet := meta.(*ExpressionAtomMeta)
			if len(amet.ConstantID) > 0 {
				expressAtm.Constant = importTable[amet.ConstantID].(*Constant)
			}
			if len(amet.ExpressionAtomID) > 0 {
				expressAtm.ExpressionAtom = importTable[amet.ExpressionAtomID].(*ExpressionAtom)
			}
			if len(amet.VariableID) > 0 {
				expressAtm.Variable = importTable[amet.VariableID].(*Variable)
			}
			if len(amet.FunctionCallID) > 0 {
				expressAtm.FunctionCall = importTable[amet.FunctionCallID].(*FunctionCall)
			}
			if len(amet.ArrayMapSelectorID) > 0 {
				expressAtm.ArrayMapSelector = importTable[amet.ArrayMapSelectorID].(*ArrayMapSelector)
			}
		case TypeFunctionCall:
			funcCall := node.(*FunctionCall)
			amet := meta.(*FunctionCallMeta)
			if len(amet.ArgumentListID) > 0 {
				funcCall.ArgumentList = importTable[amet.ArgumentListID].(*ArgumentList)
			}
		case TypeRuleEntry:
			ruleEntry := node.(*RuleEntry)
			amet := meta.(*RuleEntryMeta)
			if len(amet.WhenScopeID) > 0 {
				ruleEntry.WhenScope = importTable[amet.WhenScopeID].(*WhenScope)
			}
			if len(amet.ThenScopeID) > 0 {
				ruleEntry.ThenScope = importTable[amet.ThenScopeID].(*ThenScope)
			}
		case TypeThenExpression:
			thenExpr := node.(*ThenExpression)
			amet := meta.(*ThenExpressionMeta)
			if len(amet.AssignmentID) > 0 {
				thenExpr.Assignment = importTable[amet.AssignmentID].(*Assignment)
			}
			if len(amet.ExpressionAtomID) > 0 {
				thenExpr.ExpressionAtom = importTable[amet.ExpressionAtomID].(*ExpressionAtom)
			}
		case TypeThenExpressionList:
			ThenExprList := node.(*ThenExpressionList)
			amet := meta.(*ThenExpressionListMeta)
			if amet.ThenExpressionIDs != nil && len(amet.ThenExpressionIDs) > 0 {
				ThenExprList.ThenExpressions = make([]*ThenExpression, len(amet.ThenExpressionIDs))
				for k, v := range amet.ThenExpressionIDs {
					if node, ok := importTable[v]; ok {
						ThenExprList.ThenExpressions[k] = node.(*ThenExpression)
					} else {
						logrus.Errorf("then expression with ast id %s not catalogued", v)
					}
				}
			}
		case TypeThenScope:
			thenScope := node.(*ThenScope)
			amet := meta.(*ThenScopeMeta)
			if len(amet.ThenExpressionListID) > 0 {
				thenScope.ThenExpressionList = importTable[amet.ThenExpressionListID].(*ThenExpressionList)
			}
		case TypeVariable:
			variable := node.(*Variable)
			amet := meta.(*VariableMeta)
			if len(amet.VariableID) > 0 {
				variable.Variable = importTable[amet.VariableID].(*Variable)
			}
			if len(amet.ArrayMapSelectorID) > 0 {
				variable.ArrayMapSelector = importTable[amet.ArrayMapSelectorID].(*ArrayMapSelector)
			}
		case TypeWhenScope:
			whenScope := node.(*WhenScope)
			amet := meta.(*WhenScopeMeta)
			if len(amet.ExpressionID) > 0 {
				whenScope.Expression = importTable[amet.ExpressionID].(*Expression)
			}
		default:
			panic("Unrecognized meta type")
		}
	}

	// Rebuilding Working Memory
	if cat.MemoryVariableSnapshotMap != nil && len(cat.MemoryVariableSnapshotMap) > 0 {
		for key, value := range cat.MemoryVariableSnapshotMap {
			if n, ok := importTable[value]; ok {
				workingMem.variableSnapshotMap[key] = n.(*Variable)
			} else {
				logrus.Warnf("snapshot %s in working memory have no referenced variable with ASTID %s", key, value)
			}
		}
	}
	if cat.MemoryExpressionSnapshotMap != nil && len(cat.MemoryExpressionSnapshotMap) > 0 {
		for key, value := range cat.MemoryExpressionSnapshotMap {
			workingMem.expressionSnapshotMap[key] = importTable[value].(*Expression)
		}
	}
	if cat.MemoryExpressionAtomSnapshotMap != nil && len(cat.MemoryExpressionAtomSnapshotMap) > 0 {
		for key, value := range cat.MemoryExpressionAtomSnapshotMap {
			workingMem.expressionAtomSnapshotMap[key] = importTable[value].(*ExpressionAtom)
		}
	}
	if cat.MemoryExpressionVariableMap != nil && len(cat.MemoryExpressionVariableMap) > 0 {
		for key, value := range cat.MemoryExpressionVariableMap {
			variable := importTable[key].(*Variable)
			workingMem.expressionVariableMap[variable] = make([]*Expression, len(value))
			for i, j := range value {
				workingMem.expressionVariableMap[variable][i] = importTable[j].(*Expression)
			}
		}
	}
	if cat.MemoryExpressionAtomVariableMap != nil && len(cat.MemoryExpressionAtomVariableMap) > 0 {
		for key, value := range cat.MemoryExpressionAtomVariableMap {
			variable := importTable[key].(*Variable)
			workingMem.expressionAtomVariableMap[variable] = make([]*ExpressionAtom, len(value))
			for i, j := range value {
				workingMem.expressionAtomVariableMap[variable][i] = importTable[j].(*ExpressionAtom)
			}
		}
	}

	return knowledgeBase, nil
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
	for key, value := range cat.Data {
		if j, ok := that.Data[key]; ok {
			if !j.Equals(value) {

				return false
			}
		} else {

			return false
		}
	}
	for key, value := range cat.MemoryVariableSnapshotMap {
		if j, ok := that.MemoryVariableSnapshotMap[key]; ok {
			if j != value {

				return false
			}
		} else {

			return false
		}
	}
	for key, value := range cat.MemoryExpressionSnapshotMap {
		if j, ok := that.MemoryExpressionSnapshotMap[key]; ok {
			if j != value {

				return false
			}
		} else {

			return false
		}
	}
	for key, value := range cat.MemoryExpressionAtomSnapshotMap {
		if j, ok := that.MemoryExpressionAtomSnapshotMap[key]; ok {
			if j != value {

				return false
			}
		} else {

			return false
		}
	}
	for key, value := range cat.MemoryExpressionVariableMap {
		if mapValue, ok := that.MemoryExpressionVariableMap[key]; ok {
			if len(mapValue) != len(value) {

				return false
			}
			for in, st := range value {
				if mapValue[in] != st {

					return false
				}
			}
		} else {

			return false
		}
	}
	for key, value := range cat.MemoryExpressionAtomVariableMap {
		if vari, ok := that.MemoryExpressionAtomVariableMap[key]; ok {
			if len(vari) != len(value) {

				return false
			}
			for in, st := range value {
				if vari[in] != st {

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
func (cat *Catalog) ReadCatalogFromReader(reader io.Reader) error {
	// Read the catalog file version.
	str, err := ReadStringFromReader(reader) // V
	if err != nil {

		return err
	}
	if str != Version {

		return fmt.Errorf("invalid version %s", str)
	}

	// Read the knowledgebase name.
	str, err = ReadStringFromReader(reader) // V
	if err != nil {

		return err
	}
	cat.KnowledgeBaseName = str

	// Read the knowledgebase version.
	str, err = ReadStringFromReader(reader) // V
	if err != nil {

		return err
	}
	cat.KnowledgeBaseVersion = str

	// Writedown meta counts.
	count, err := ReadIntFromReader(reader) // V
	if err != nil {

		return err
	}

	cat.Data = make(map[string]Meta)

	for i := uint64(0); i < count; i++ {
		key, err := ReadStringFromReader(reader) // V
		if err != nil {

			return err
		}
		metaType, err := ReadIntFromReader(reader) // V
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
		err = meta.ReadMetaFrom(reader) // V
		if err != nil {

			return err
		}
		cat.Data[key] = meta
	}

	str, err = ReadStringFromReader(reader)
	if err != nil {

		return err
	}
	cat.MemoryName = str

	str, err = ReadStringFromReader(reader)
	if err != nil {

		return err
	}
	cat.MemoryVersion = str

	// Writedown meta counts.
	count, err = ReadIntFromReader(reader)
	if err != nil {

		return err
	}

	cat.MemoryVariableSnapshotMap = make(map[string]string)
	for index := uint64(0); index < count; index++ {
		key, err := ReadStringFromReader(reader)
		if err != nil {

			return err
		}
		val, err := ReadStringFromReader(reader)
		if err != nil {

			return err
		}
		cat.MemoryVariableSnapshotMap[key] = val
	}

	// MemoryExpressionSnapshotMap meta counts.
	count, err = ReadIntFromReader(reader)
	if err != nil {

		return err
	}

	cat.MemoryExpressionSnapshotMap = make(map[string]string)
	for index := uint64(0); index < count; index++ {
		key, err := ReadStringFromReader(reader)
		if err != nil {

			return err
		}
		val, err := ReadStringFromReader(reader)
		if err != nil {

			return err
		}
		cat.MemoryExpressionSnapshotMap[key] = val
	}

	// MemoryExpressionAtomSnapshotMap meta counts.
	count, err = ReadIntFromReader(reader)
	if err != nil {

		return err
	}

	cat.MemoryExpressionAtomSnapshotMap = make(map[string]string)
	for index := uint64(0); index < count; index++ {
		key, err := ReadStringFromReader(reader)
		if err != nil {

			return err
		}
		val, err := ReadStringFromReader(reader)
		if err != nil {

			return err
		}
		cat.MemoryExpressionAtomSnapshotMap[key] = val
	}

	// MemoryExpressionVariableMap meta counts.
	count, err = ReadIntFromReader(reader)
	if err != nil {

		return err
	}

	cat.MemoryExpressionVariableMap = make(map[string][]string)
	for index := uint64(0); index < count; index++ {
		key, err := ReadStringFromReader(reader)
		if err != nil {

			return err
		}
		incount, err := ReadIntFromReader(reader)
		if err != nil {

			return err
		}
		content := make([]string, incount)
		for subIndex := uint64(0); subIndex < incount; subIndex++ {
			str, err := ReadStringFromReader(reader)
			if err != nil {

				return err
			}
			content[subIndex] = str
		}
		cat.MemoryExpressionVariableMap[key] = content
	}

	// MemoryExpressionAtomVariableMap meta counts.
	count, err = ReadIntFromReader(reader)
	if err != nil {

		return err
	}

	cat.MemoryExpressionAtomVariableMap = make(map[string][]string)
	for index := uint64(0); index < count; index++ {
		key, err := ReadStringFromReader(reader)
		if err != nil {

			return err
		}
		incount, err := ReadIntFromReader(reader)
		if err != nil {

			return err
		}
		content := make([]string, incount)
		for subIndex := uint64(0); subIndex < incount; subIndex++ {
			str, err := ReadStringFromReader(reader)
			if err != nil {

				return err
			}
			content[subIndex] = str
		}
		cat.MemoryExpressionAtomVariableMap[key] = content
	}

	return nil
}

// WriteCatalogToWriter will store the content of this Catalog
// into a byte stream using provided writer.
// You are responsible for closing the writing stream once its done.
func (cat *Catalog) WriteCatalogToWriter(writer io.Writer) error {
	// Write the catalog file version.
	err := WriteStringToWriter(writer, Version)
	if err != nil {

		return err
	}

	// Write the knowledgebase name.
	err = WriteStringToWriter(writer, cat.KnowledgeBaseName)
	if err != nil {

		return err
	}

	// Write the knowledgebase version.
	err = WriteStringToWriter(writer, cat.KnowledgeBaseVersion)
	if err != nil {

		return err
	}

	// Writedown meta counts.
	err = WriteIntToWriter(writer, uint64(len(cat.Data)))
	if err != nil {

		return err
	}

	// For each meta.. write them down
	for key, value := range cat.Data {

		// Write the AST ID
		err = WriteStringToWriter(writer, key)
		if err != nil {

			return err
		}

		err := WriteIntToWriter(writer, uint64(value.GetASTType()))
		if err != nil {

			return err
		}

		// Write the meta
		err = value.WriteMetaTo(writer)
		if err != nil {

			return err
		}
	}

	// Write the MemoryName version.
	err = WriteStringToWriter(writer, cat.MemoryName)
	if err != nil {

		return err
	}

	// Write the MemoryVersion version.
	err = WriteStringToWriter(writer, cat.MemoryVersion)
	if err != nil {

		return err
	}

	// MemoryVariableSnapshotMap meta counts.
	err = WriteIntToWriter(writer, uint64(len(cat.MemoryVariableSnapshotMap)))
	if err != nil {

		return err
	}
	for key, value := range cat.MemoryVariableSnapshotMap {
		err = WriteStringToWriter(writer, key)
		if err != nil {

			return err
		}
		err = WriteStringToWriter(writer, value)
		if err != nil {

			return err
		}
	}

	// MemoryExpressionSnapshotMap meta counts.
	err = WriteIntToWriter(writer, uint64(len(cat.MemoryExpressionSnapshotMap)))
	if err != nil {

		return err
	}
	for key, value := range cat.MemoryExpressionSnapshotMap {
		err = WriteStringToWriter(writer, key)
		if err != nil {

			return err
		}
		err = WriteStringToWriter(writer, value)
		if err != nil {

			return err
		}
	}

	// MemoryExpressionAtomSnapshotMap meta counts.
	err = WriteIntToWriter(writer, uint64(len(cat.MemoryExpressionAtomSnapshotMap)))
	if err != nil {

		return err
	}
	for key, value := range cat.MemoryExpressionAtomSnapshotMap {
		err = WriteStringToWriter(writer, key)
		if err != nil {

			return err
		}
		err = WriteStringToWriter(writer, value)
		if err != nil {

			return err
		}
	}

	// MemoryExpressionVariableMap meta counts.
	err = WriteIntToWriter(writer, uint64(len(cat.MemoryExpressionVariableMap)))
	if err != nil {

		return err
	}
	for key, value := range cat.MemoryExpressionVariableMap {
		err = WriteStringToWriter(writer, key)
		if err != nil {

			return err
		}
		err = WriteIntToWriter(writer, uint64(len(value)))
		if err != nil {

			return err
		}
		for _, j := range value {
			err = WriteStringToWriter(writer, j)
			if err != nil {

				return err
			}
		}
	}

	// MemoryExpressionAtomVariableMap meta counts.
	err = WriteIntToWriter(writer, uint64(len(cat.MemoryExpressionAtomVariableMap)))
	if err != nil {

		return err
	}
	for key, value := range cat.MemoryExpressionAtomVariableMap {
		err = WriteStringToWriter(writer, key)
		if err != nil {

			return err
		}
		err = WriteIntToWriter(writer, uint64(len(value)))
		if err != nil {

			return err
		}
		for _, j := range value {
			err = WriteStringToWriter(writer, j)
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
func (meta *NodeMeta) WriteMetaTo(writer io.Writer) error {
	// First write the AST ID. this may be redundant.
	err := WriteStringToWriter(writer, meta.AstID)
	if err != nil {

		return err
	}
	// Second write the GRL Text.
	err = WriteStringToWriter(writer, meta.GrlText)
	if err != nil {

		return err
	}
	// Third write the snapshot. This might be un-necessary.
	err = WriteStringToWriter(writer, meta.Snapshot)
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
func (meta *ArgumentListMeta) WriteMetaTo(writer io.Writer) error {
	err := meta.NodeMeta.WriteMetaTo(writer)
	if err != nil {

		return err
	}

	// Write the number of arguments
	err = WriteIntToWriter(writer, uint64(len(meta.ArgumentASTIDs)))
	if err != nil {

		return err
	}

	// Write the array content
	for _, v := range meta.ArgumentASTIDs {
		err = WriteStringToWriter(writer, v)
		if err != nil {

			return err
		}
	}

	return nil
}

// ReadMetaFrom write basic AST Node information meta data from reader.
// One should not use this function directly, unless for testing
// serialization of single ASTNode.
func (meta *ArgumentListMeta) ReadMetaFrom(reader io.Reader) error {
	err := meta.NodeMeta.ReadMetaFrom(reader)
	if err != nil {

		return err
	}

	integer, err := ReadIntFromReader(reader)
	if err != nil {

		return err
	}

	meta.ArgumentASTIDs = make([]string, integer)
	for index := uint64(0); index < integer; index++ {
		s, err := ReadStringFromReader(reader)
		if err != nil {

			return err
		}
		meta.ArgumentASTIDs[index] = s
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
func (meta *ArrayMapSelectorMeta) WriteMetaTo(writer io.Writer) error {
	err := meta.NodeMeta.WriteMetaTo(writer)
	if err != nil {

		return err
	}
	err = WriteStringToWriter(writer, meta.ExpressionID)
	if err != nil {

		return err
	}

	return nil
}

// ReadMetaFrom write basic AST Node information meta data from reader.
// One should not use this function directly, unless for testing
// serialization of single ASTNode.
func (meta *ArrayMapSelectorMeta) ReadMetaFrom(reader io.Reader) error {
	err := meta.NodeMeta.ReadMetaFrom(reader)
	if err != nil {

		return err
	}
	s, err := ReadStringFromReader(reader)
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
func (meta *AssigmentMeta) WriteMetaTo(writer io.Writer) error {
	err := meta.NodeMeta.WriteMetaTo(writer)
	if err != nil {

		return err
	}

	err = WriteStringToWriter(writer, meta.VariableID)
	if err != nil {

		return err
	}
	err = WriteStringToWriter(writer, meta.ExpressionID)
	if err != nil {

		return err
	}

	err = WriteBoolToWriter(writer, meta.IsAssign)
	if err != nil {

		return err
	}
	err = WriteBoolToWriter(writer, meta.IsPlusAssign)
	if err != nil {

		return err
	}
	err = WriteBoolToWriter(writer, meta.IsMinusAssign)
	if err != nil {

		return err
	}
	err = WriteBoolToWriter(writer, meta.IsDivAssign)
	if err != nil {

		return err
	}
	err = WriteBoolToWriter(writer, meta.IsMulAssign)
	if err != nil {

		return err
	}

	return nil
}

// ReadMetaFrom write basic AST Node information meta data from reader.
// One should not use this function directly, unless for testing
// serialization of single ASTNode.
func (meta *AssigmentMeta) ReadMetaFrom(reader io.Reader) error {
	err := meta.NodeMeta.ReadMetaFrom(reader)
	if err != nil {

		return err
	}

	stringFromReader, err := ReadStringFromReader(reader)
	if err != nil {

		return err
	}
	meta.VariableID = stringFromReader
	stringFromReader, err = ReadStringFromReader(reader)
	if err != nil {

		return err
	}
	meta.ExpressionID = stringFromReader

	boolReaded, err := ReadBoolFromReader(reader)
	if err != nil {

		return err
	}
	meta.IsAssign = boolReaded

	boolReaded, err = ReadBoolFromReader(reader)
	if err != nil {

		return err
	}
	meta.IsPlusAssign = boolReaded

	boolReaded, err = ReadBoolFromReader(reader)
	if err != nil {

		return err
	}
	meta.IsMinusAssign = boolReaded

	boolReaded, err = ReadBoolFromReader(reader)
	if err != nil {

		return err
	}
	meta.IsDivAssign = boolReaded

	boolReaded, err = ReadBoolFromReader(reader)
	if err != nil {

		return err
	}
	meta.IsMulAssign = boolReaded

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
func (meta *ConstantMeta) WriteMetaTo(writer io.Writer) error {
	err := meta.NodeMeta.WriteMetaTo(writer)
	if err != nil {

		return err
	}
	err = WriteIntToWriter(writer, uint64(meta.ValueType))
	if err != nil {

		return err
	}
	err = WriteIntToWriter(writer, uint64(len(meta.ValueBytes)))
	if err != nil {

		return err
	}
	_, err = writer.Write(meta.ValueBytes)
	if err != nil {

		return err
	}
	err = WriteBoolToWriter(writer, meta.IsNil)
	if err != nil {

		return err
	}

	return nil
}

// ReadMetaFrom write basic AST Node information meta data from reader.
// One should not use this function directly, unless for testing
// serialization of single ASTNode.
func (meta *ConstantMeta) ReadMetaFrom(reader io.Reader) error {
	err := meta.NodeMeta.ReadMetaFrom(reader)
	if err != nil {

		return err
	}
	i, err := ReadIntFromReader(reader)
	if err != nil {

		return err
	}
	meta.ValueType = ValueType(i)

	length, err := ReadIntFromReader(reader)
	if err != nil {

		return err
	}
	byteArr := make([]byte, length)
	_, err = reader.Read(byteArr)
	if err != nil {

		return err
	}
	meta.ValueBytes = byteArr

	b, err := ReadBoolFromReader(reader)
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
func (meta *ExpressionMeta) WriteMetaTo(writer io.Writer) error {
	err := meta.NodeMeta.WriteMetaTo(writer)
	if err != nil {

		return err
	}
	err = WriteStringToWriter(writer, meta.LeftExpressionID)
	if err != nil {

		return err
	}
	err = WriteStringToWriter(writer, meta.RightExpressionID)
	if err != nil {

		return err
	}
	err = WriteStringToWriter(writer, meta.SingleExpressionID)
	if err != nil {

		return err
	}
	err = WriteStringToWriter(writer, meta.ExpressionAtomID)
	if err != nil {

		return err
	}
	err = WriteIntToWriter(writer, uint64(meta.Operator))
	if err != nil {

		return err
	}
	err = WriteBoolToWriter(writer, meta.Negated)
	if err != nil {

		return err
	}

	return nil
}

// ReadMetaFrom write basic AST Node information meta data from reader.
// One should not use this function directly, unless for testing
// serialization of single ASTNode.
func (meta *ExpressionMeta) ReadMetaFrom(reader io.Reader) error {
	err := meta.NodeMeta.ReadMetaFrom(reader)
	if err != nil {

		return err
	}
	theString, err := ReadStringFromReader(reader)
	if err != nil {

		return err
	}
	meta.LeftExpressionID = theString
	theString, err = ReadStringFromReader(reader)
	if err != nil {

		return err
	}
	meta.RightExpressionID = theString
	theString, err = ReadStringFromReader(reader)
	if err != nil {

		return err
	}
	meta.SingleExpressionID = theString
	theString, err = ReadStringFromReader(reader)
	if err != nil {

		return err
	}
	meta.ExpressionAtomID = theString
	i, err := ReadIntFromReader(reader)
	if err != nil {

		return err
	}
	meta.Operator = int(i)
	b, err := ReadBoolFromReader(reader)
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
func (meta *ExpressionAtomMeta) WriteMetaTo(writer io.Writer) error {
	err := meta.NodeMeta.WriteMetaTo(writer)
	if err != nil {

		return err
	}
	err = WriteStringToWriter(writer, meta.VariableName)
	if err != nil {

		return err
	}
	err = WriteStringToWriter(writer, meta.ConstantID)
	if err != nil {

		return err
	}
	err = WriteStringToWriter(writer, meta.FunctionCallID)
	if err != nil {

		return err
	}
	err = WriteStringToWriter(writer, meta.VariableID)
	if err != nil {

		return err
	}
	err = WriteBoolToWriter(writer, meta.Negated)
	if err != nil {

		return err
	}
	err = WriteStringToWriter(writer, meta.ExpressionAtomID)
	if err != nil {

		return err
	}
	err = WriteStringToWriter(writer, meta.ArrayMapSelectorID)
	if err != nil {

		return err
	}

	return nil
}

// ReadMetaFrom write basic AST Node information meta data from reader.
// One should not use this function directly, unless for testing
// serialization of single ASTNode.
func (meta *ExpressionAtomMeta) ReadMetaFrom(reader io.Reader) error {
	err := meta.NodeMeta.ReadMetaFrom(reader)
	if err != nil {

		return err
	}
	stringFromReader, err := ReadStringFromReader(reader)
	if err != nil {

		return err
	}
	meta.VariableName = stringFromReader
	stringFromReader, err = ReadStringFromReader(reader)
	if err != nil {

		return err
	}
	meta.ConstantID = stringFromReader
	stringFromReader, err = ReadStringFromReader(reader)
	if err != nil {

		return err
	}
	meta.FunctionCallID = stringFromReader
	stringFromReader, err = ReadStringFromReader(reader)
	if err != nil {

		return err
	}
	meta.VariableID = stringFromReader
	b, err := ReadBoolFromReader(reader)
	if err != nil {

		return err
	}
	meta.Negated = b
	stringFromReader, err = ReadStringFromReader(reader)
	if err != nil {

		return err
	}
	meta.ExpressionAtomID = stringFromReader
	stringFromReader, err = ReadStringFromReader(reader)
	if err != nil {

		return err
	}
	meta.ArrayMapSelectorID = stringFromReader

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
func (meta *FunctionCallMeta) WriteMetaTo(writer io.Writer) error {
	err := meta.NodeMeta.WriteMetaTo(writer)
	if err != nil {

		return err
	}
	err = WriteStringToWriter(writer, meta.FunctionName)
	if err != nil {

		return err
	}
	err = WriteStringToWriter(writer, meta.ArgumentListID)
	if err != nil {

		return err
	}

	return nil
}

// ReadMetaFrom write basic AST Node information meta data from reader.
// One should not use this function directly, unless for testing
// serialization of single ASTNode.
func (meta *FunctionCallMeta) ReadMetaFrom(reader io.Reader) error {
	err := meta.NodeMeta.ReadMetaFrom(reader)
	if err != nil {

		return err
	}
	stringFromReader, err := ReadStringFromReader(reader)
	if err != nil {

		return err
	}
	meta.FunctionName = stringFromReader
	stringFromReader, err = ReadStringFromReader(reader)
	if err != nil {

		return err
	}
	meta.ArgumentListID = stringFromReader

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
func (meta *RuleEntryMeta) WriteMetaTo(writer io.Writer) error {
	err := meta.NodeMeta.WriteMetaTo(writer)
	if err != nil {

		return err
	}
	err = WriteStringToWriter(writer, meta.RuleName)
	if err != nil {

		return err
	}
	err = WriteStringToWriter(writer, meta.RuleDescription)
	if err != nil {

		return err
	}
	err = WriteIntToWriter(writer, uint64(meta.Salience))
	if err != nil {

		return err
	}
	err = WriteStringToWriter(writer, meta.WhenScopeID)
	if err != nil {

		return err
	}
	err = WriteStringToWriter(writer, meta.ThenScopeID)
	if err != nil {

		return err
	}

	return nil
}

// ReadMetaFrom write basic AST Node information meta data from reader.
// One should not use this function directly, unless for testing
// serialization of single ASTNode.
func (meta *RuleEntryMeta) ReadMetaFrom(reader io.Reader) error {
	err := meta.NodeMeta.ReadMetaFrom(reader)
	if err != nil {

		return err
	}
	stringFromReader, err := ReadStringFromReader(reader)
	if err != nil {

		return err
	}
	meta.RuleName = stringFromReader
	stringFromReader, err = ReadStringFromReader(reader)
	if err != nil {

		return err
	}
	meta.RuleDescription = stringFromReader
	i, err := ReadIntFromReader(reader)
	if err != nil {

		return err
	}
	meta.Salience = int(i)
	stringFromReader, err = ReadStringFromReader(reader)
	if err != nil {

		return err
	}
	meta.WhenScopeID = stringFromReader
	stringFromReader, err = ReadStringFromReader(reader)
	if err != nil {

		return err
	}
	meta.ThenScopeID = stringFromReader

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
func (meta *ThenExpressionMeta) WriteMetaTo(writer io.Writer) error {
	err := meta.NodeMeta.WriteMetaTo(writer)
	if err != nil {

		return err
	}
	err = WriteStringToWriter(writer, meta.AssignmentID)
	if err != nil {

		return err
	}
	err = WriteStringToWriter(writer, meta.ExpressionAtomID)
	if err != nil {

		return err
	}

	return nil
}

// ReadMetaFrom write basic AST Node information meta data from reader.
// One should not use this function directly, unless for testing
// serialization of single ASTNode.
func (meta *ThenExpressionMeta) ReadMetaFrom(reader io.Reader) error {
	err := meta.NodeMeta.ReadMetaFrom(reader)
	if err != nil {

		return err
	}
	theString, err := ReadStringFromReader(reader)
	if err != nil {

		return err
	}
	meta.AssignmentID = theString
	theString, err = ReadStringFromReader(reader)
	if err != nil {

		return err
	}
	meta.ExpressionAtomID = theString

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
func (meta *ThenExpressionListMeta) WriteMetaTo(writer io.Writer) error {
	err := meta.NodeMeta.WriteMetaTo(writer)
	if err != nil {

		return err
	}

	err = WriteIntToWriter(writer, uint64(len(meta.ThenExpressionIDs)))
	if err != nil {

		return err
	}

	for _, v := range meta.ThenExpressionIDs {
		err = WriteStringToWriter(writer, v)
		if err != nil {

			return err
		}
	}

	return nil
}

// ReadMetaFrom write basic AST Node information meta data from reader.
// One should not use this function directly, unless for testing
// serialization of single ASTNode.
func (meta *ThenExpressionListMeta) ReadMetaFrom(reader io.Reader) error {
	err := meta.NodeMeta.ReadMetaFrom(reader)
	if err != nil {

		return err
	}

	count, err := ReadIntFromReader(reader)
	if err != nil {

		return err
	}

	meta.ThenExpressionIDs = make([]string, count)
	for index := uint64(0); index < count; index++ {
		s, err := ReadStringFromReader(reader)
		if err != nil {

			return err
		}
		meta.ThenExpressionIDs[index] = s
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
func (meta *ThenScopeMeta) WriteMetaTo(writer io.Writer) error {
	err := meta.NodeMeta.WriteMetaTo(writer)
	if err != nil {

		return err
	}
	err = WriteStringToWriter(writer, meta.ThenExpressionListID)
	if err != nil {

		return err
	}

	return nil
}

// ReadMetaFrom write basic AST Node information meta data from reader.
// One should not use this function directly, unless for testing
// serialization of single ASTNode.
func (meta *ThenScopeMeta) ReadMetaFrom(reader io.Reader) error {
	err := meta.NodeMeta.ReadMetaFrom(reader)
	if err != nil {

		return err
	}
	s, err := ReadStringFromReader(reader)
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
func (meta *VariableMeta) WriteMetaTo(writer io.Writer) error {
	err := meta.NodeMeta.WriteMetaTo(writer)
	if err != nil {

		return err
	}
	err = WriteStringToWriter(writer, meta.Name)
	if err != nil {

		return err
	}
	err = WriteStringToWriter(writer, meta.VariableID)
	if err != nil {

		return err
	}
	err = WriteStringToWriter(writer, meta.ArrayMapSelectorID)
	if err != nil {

		return err
	}

	return nil
}

// ReadMetaFrom write basic AST Node information meta data from reader.
// One should not use this function directly, unless for testing
// serialization of single ASTNode.
func (meta *VariableMeta) ReadMetaFrom(reader io.Reader) error {
	err := meta.NodeMeta.ReadMetaFrom(reader)
	if err != nil {

		return err
	}
	stringFromReader, err := ReadStringFromReader(reader)
	if err != nil {

		return err
	}
	meta.Name = stringFromReader
	stringFromReader, err = ReadStringFromReader(reader)
	if err != nil {

		return err
	}
	meta.VariableID = stringFromReader
	stringFromReader, err = ReadStringFromReader(reader)
	if err != nil {

		return err
	}
	meta.ArrayMapSelectorID = stringFromReader

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
func (meta *WhenScopeMeta) WriteMetaTo(writer io.Writer) error {
	err := meta.NodeMeta.WriteMetaTo(writer)
	if err != nil {

		return err
	}
	err = WriteStringToWriter(writer, meta.ExpressionID)
	if err != nil {

		return err
	}

	return nil
}

// ReadMetaFrom write basic AST Node information meta data from reader.
// One should not use this function directly, unless for testing
// serialization of single ASTNode.
func (meta *WhenScopeMeta) ReadMetaFrom(reader io.Reader) error {
	err := meta.NodeMeta.ReadMetaFrom(reader)
	if err != nil {

		return err
	}
	s, err := ReadStringFromReader(reader)
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
func WriteStringToWriter(writer io.Writer, s string) error {
	length := make([]byte, 8)
	data := []byte(s)
	binary.LittleEndian.PutUint64(length, uint64(len(data)))
	writeCount, err := WriteFull(writer, length)

	TotalWrite += uint64(writeCount)
	if err != nil {

		return err
	}
	writeCount, err = WriteFull(writer, data)
	TotalWrite += uint64(writeCount)
	WriteCount++

	return err
}

// ReadStringFromReader read a string from reader.
func ReadStringFromReader(reader io.Reader) (string, error) {
	length := make([]byte, 8)

	counter, err := io.ReadFull(reader, length)
	TotalRead += uint64(counter)
	if err != nil {

		return "", err
	}
	strLen := binary.LittleEndian.Uint64(length)
	strByte := make([]byte, int(strLen))
	counter, err = io.ReadFull(reader, strByte)
	TotalRead += uint64(counter)
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
func WriteBoolToWriter(writer io.Writer, aBoolean bool) error {
	data := make([]byte, 1)
	if aBoolean {
		data[0] = 1
	} else {
		data[0] = 0
	}
	c, err := WriteFull(writer, data)
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
