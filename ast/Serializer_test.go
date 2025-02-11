//  Copyright DataWiseHQ/grule-rule-engine Authors
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
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestReadWriteString(t *testing.T) {
	str := uuid.New().String()
	buff := &bytes.Buffer{}
	err := WriteStringToWriter(buff, str)
	assert.Nil(t, err)

	buff2 := bytes.NewBuffer(buff.Bytes())
	str2, err := ReadStringFromReader(buff2)
	assert.Nil(t, err)

	assert.Equal(t, str, str2)
}

func TestAssigmentMetaReadWrite(t *testing.T) {
	assigment := &AssigmentMeta{
		NodeMeta: NodeMeta{
			AstID:    uuid.New().String(),
			GrlText:  uuid.New().String(),
			Snapshot: uuid.New().String(),
		},
		VariableID:    uuid.New().String(),
		ExpressionID:  uuid.New().String(),
		IsAssign:      true,
		IsPlusAssign:  false,
		IsMinusAssign: true,
		IsDivAssign:   false,
		IsMulAssign:   true,
	}
	buff := &bytes.Buffer{}
	err := assigment.WriteMetaTo(buff)
	assert.Nil(t, err)

	buff2 := bytes.NewBuffer(buff.Bytes())

	assigment2 := &AssigmentMeta{}
	err = assigment2.ReadMetaFrom(buff2)
	assert.Nil(t, err)

	assert.True(t, assigment.Equals(assigment2))
}

func TestSerialization(t *testing.T) {
	cat := &Catalog{
		KnowledgeBaseName:    uuid.New().String(),
		KnowledgeBaseVersion: uuid.New().String(),
		Data: map[string]Meta{
			uuid.New().String(): &AssigmentMeta{
				NodeMeta: NodeMeta{
					AstID:    uuid.New().String(),
					GrlText:  uuid.New().String(),
					Snapshot: uuid.New().String(),
				},
				VariableID:    uuid.New().String(),
				ExpressionID:  uuid.New().String(),
				IsAssign:      true,
				IsPlusAssign:  false,
				IsMinusAssign: true,
				IsDivAssign:   false,
				IsMulAssign:   true,
			},
			uuid.New().String(): &ConstantMeta{
				NodeMeta: NodeMeta{
					AstID:    uuid.New().String(),
					GrlText:  uuid.New().String(),
					Snapshot: uuid.New().String(),
				},
				ValueType:  TypeString,
				ValueBytes: []byte(uuid.New().String()),
				IsNil:      true,
			},
			uuid.New().String(): &ExpressionMeta{
				NodeMeta: NodeMeta{
					AstID:    uuid.New().String(),
					GrlText:  uuid.New().String(),
					Snapshot: uuid.New().String(),
				},
				LeftExpressionID:   uuid.New().String(),
				RightExpressionID:  uuid.New().String(),
				SingleExpressionID: uuid.New().String(),
				ExpressionAtomID:   uuid.New().String(),
				Operator:           3,
				Negated:            false,
			},
			uuid.New().String(): &ExpressionAtomMeta{
				NodeMeta: NodeMeta{
					AstID:    uuid.New().String(),
					GrlText:  uuid.New().String(),
					Snapshot: uuid.New().String(),
				},
				VariableName:       uuid.New().String(),
				ConstantID:         uuid.New().String(),
				FunctionCallID:     uuid.New().String(),
				VariableID:         uuid.New().String(),
				Negated:            true,
				ExpressionAtomID:   uuid.New().String(),
				ArrayMapSelectorID: uuid.New().String(),
			},
			uuid.New().String(): &FunctionCallMeta{
				NodeMeta: NodeMeta{
					AstID:    uuid.New().String(),
					GrlText:  uuid.New().String(),
					Snapshot: uuid.New().String(),
				},
				FunctionName:   uuid.New().String(),
				ArgumentListID: uuid.New().String(),
			},
			uuid.New().String(): &RuleEntryMeta{
				NodeMeta: NodeMeta{
					AstID:    uuid.New().String(),
					GrlText:  uuid.New().String(),
					Snapshot: uuid.New().String(),
				},
				RuleName:        uuid.New().String(),
				RuleDescription: uuid.New().String(),
				Salience:        234,
				WhenScopeID:     uuid.New().String(),
				ThenScopeID:     uuid.New().String(),
			},
			uuid.New().String(): &ThenExpressionMeta{
				NodeMeta: NodeMeta{
					AstID:    uuid.New().String(),
					GrlText:  uuid.New().String(),
					Snapshot: uuid.New().String(),
				},
				AssignmentID:     "",
				ExpressionAtomID: "",
			},
			uuid.New().String(): &ThenExpressionListMeta{
				NodeMeta: NodeMeta{
					AstID:    uuid.New().String(),
					GrlText:  uuid.New().String(),
					Snapshot: uuid.New().String(),
				},
				ThenExpressionIDs: nil,
			},
			uuid.New().String(): &ThenScopeMeta{
				NodeMeta: NodeMeta{
					AstID:    uuid.New().String(),
					GrlText:  uuid.New().String(),
					Snapshot: uuid.New().String(),
				},
				ThenExpressionListID: "",
			},
			uuid.New().String(): &VariableMeta{
				NodeMeta: NodeMeta{
					AstID:    uuid.New().String(),
					GrlText:  uuid.New().String(),
					Snapshot: uuid.New().String(),
				},
				Name:               uuid.New().String(),
				VariableID:         uuid.New().String(),
				ArrayMapSelectorID: uuid.New().String(),
			},
			uuid.New().String(): &WhenScopeMeta{
				NodeMeta: NodeMeta{
					AstID:    uuid.New().String(),
					GrlText:  uuid.New().String(),
					Snapshot: uuid.New().String(),
				},
				ExpressionID: uuid.New().String(),
			},
		},
		MemoryName:    "MemoryName",
		MemoryVersion: "MemoryVersion",
		MemoryVariableSnapshotMap: map[string]string{
			uuid.New().String(): uuid.New().String(),
			uuid.New().String(): uuid.New().String(),
			uuid.New().String(): uuid.New().String(),
		},
		MemoryExpressionAtomSnapshotMap: map[string]string{
			uuid.New().String(): uuid.New().String(),
			uuid.New().String(): uuid.New().String(),
			uuid.New().String(): uuid.New().String(),
			uuid.New().String(): uuid.New().String(),
		},
		MemoryExpressionVariableMap: map[string][]string{
			uuid.New().String(): {uuid.New().String(), uuid.New().String(), uuid.New().String()},
			uuid.New().String(): {uuid.New().String(), uuid.New().String(), uuid.New().String()},
			uuid.New().String(): {uuid.New().String(), uuid.New().String(), uuid.New().String()},
			uuid.New().String(): {uuid.New().String(), uuid.New().String(), uuid.New().String()},
			uuid.New().String(): {uuid.New().String(), uuid.New().String(), uuid.New().String()},
		},
		MemoryExpressionAtomVariableMap: map[string][]string{
			uuid.New().String(): {uuid.New().String(), uuid.New().String(), uuid.New().String()},
			uuid.New().String(): {uuid.New().String(), uuid.New().String(), uuid.New().String()},
		},
	}

	buffer := &bytes.Buffer{}

	err := cat.WriteCatalogToWriter(buffer)
	assert.Nil(t, err)

	catBytes := buffer.Bytes()

	buffer2 := bytes.NewBuffer(catBytes)
	cat2 := &Catalog{}
	err = cat2.ReadCatalogFromReader(buffer2)
	assert.Nil(t, err)

	assert.True(t, cat.Equals(cat2))
}
