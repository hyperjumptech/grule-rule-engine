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
	"encoding/binary"
	"fmt"
	"io"
)

type NodeType int
type ValueType int

const (
	TypeArgumentList NodeType = iota
	TypeArrayMapSelector
	TypeAssignment
	TypeExpression
	TypeConstant
	TypeExpressionAtom
	TypeFunctionCall
	TypeRuleEntry
	TypeThenExpression
	TypeThenExpressionList
	TypeThenScope
	TypeVariable
	TypeWhenScope

	TypeString ValueType = iota
	TypeInteger
	TypeFloat
	TypeBoolean

	Version = "1.8"
)

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

func (cat *Catalog) AddMeta(astID string, meta Meta) bool {
	if cat.Data == nil {
		cat.Data = make(map[string]Meta)
	}
	if _, ok := cat.Data[astID]; !ok {
		cat.Data[astID] = meta
	}
	return false
}

type Meta interface {
	GetASTType() NodeType
	GetAstID() string
	GetGrlText() string
	GetSnapshot() string
	WriteMetaTo(writer io.Writer) error
	ReadMetaFrom(reader io.Reader) error
	Equals(that Meta) bool
}

type NodeMeta struct {
	AstID    string
	GrlText  string
	Snapshot string
}

func (meta *NodeMeta) GetAstID() string {
	return meta.AstID
}
func (meta *NodeMeta) GetGrlText() string {
	return meta.GrlText
}
func (meta *NodeMeta) GetSnapshot() string {
	return meta.Snapshot
}
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

type ArgumentListMeta struct {
	NodeMeta
	ArgumentASTIDs []string
}

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

func (meta *ArgumentListMeta) GetASTType() NodeType {
	return TypeArgumentList
}

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

type ArrayMapSelectorMeta struct {
	NodeMeta
	ExpressionID string
}

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

func (meta *ArrayMapSelectorMeta) GetASTType() NodeType {
	return TypeArrayMapSelector
}

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

func (meta *AssigmentMeta) GetASTType() NodeType {
	return TypeAssignment
}

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

type ConstantMeta struct {
	NodeMeta

	ValueType  ValueType
	ValueBytes []byte
	IsNil      bool
}

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

func (meta *ConstantMeta) GetASTType() NodeType {
	return TypeConstant
}

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
	_, err = r.Read(byteArr)
	if err != nil {
		return err
	}
	meta.ValueBytes = byteArr

	b, err := ReadBoolFromReader(r)
	if err != nil {
		return err
	}
	meta.IsNil = b

	return nil
}

type ExpressionMeta struct {
	NodeMeta
	LeftExpressionID   string
	RightExpressionID  string
	SingleExpressionID string
	ExpressionAtomID   string
	Operator           int
	Negated            bool
}

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

func (meta *ExpressionMeta) GetASTType() NodeType {
	return TypeExpression
}

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

func (meta *ExpressionAtomMeta) GetASTType() NodeType {
	return TypeExpressionAtom
}

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

type FunctionCallMeta struct {
	NodeMeta
	FunctionName   string
	ArgumentListID string
}

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

func (meta *FunctionCallMeta) GetASTType() NodeType {
	return TypeFunctionCall
}

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

type RuleEntryMeta struct {
	NodeMeta

	RuleName        string
	RuleDescription string
	Salience        int
	WhenScopeID     string
	ThenScopeID     string
}

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

func (meta *RuleEntryMeta) GetASTType() NodeType {
	return TypeRuleEntry
}

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

type ThenExpressionMeta struct {
	NodeMeta

	AssignmentID     string
	ExpressionAtomID string
}

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

func (meta *ThenExpressionMeta) GetASTType() NodeType {
	return TypeThenExpression
}

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

type ThenExpressionListMeta struct {
	NodeMeta

	ThenExpressionIDs []string
}

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

func (meta *ThenExpressionListMeta) GetASTType() NodeType {
	return TypeThenExpressionList
}

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

type ThenScopeMeta struct {
	NodeMeta
	ThenExpressionListID string
}

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

func (meta *ThenScopeMeta) GetASTType() NodeType {
	return TypeThenScope
}

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

type VariableMeta struct {
	NodeMeta

	Name               string
	VariableID         string
	ArrayMapSelectorID string
}

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

func (meta *VariableMeta) GetASTType() NodeType {
	return TypeVariable
}

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

type WhenScopeMeta struct {
	NodeMeta
	ExpressionID string
}

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

func (meta *WhenScopeMeta) GetASTType() NodeType {
	return TypeWhenScope
}

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
	TotalRead  = uint64(0)
	TotalWrite = uint64(0)
	ReadCount  = 0
	WriteCount = 0
)

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

func WriteIntToWriter(w io.Writer, i uint64) error {
	data := make([]byte, 8)
	binary.LittleEndian.PutUint64(data, i)
	c, err := WriteFull(w, data)
	TotalWrite += uint64(c)
	WriteCount++
	return err
}

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

func ReadBoolFromReader(r io.Reader) (bool, error) {
	byteArray := make([]byte, 1)
	c, err := io.ReadFull(r, byteArray)
	TotalRead += uint64(c)
	if err != nil {
		return false, err
	}
	return byteArray[0] == 1, nil
}

func WriteFloatToWriter(w io.Writer, f float64) error {
	return WriteIntToWriter(w, uint64(f))
}

func ReadFloatFromReader(r io.Reader) (float64, error) {
	i, err := ReadIntFromReader(r)
	if err != nil {
		return 0, err
	}
	return float64(i), nil
}
