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

package pkg

// CloneRecord contains information about all AST versions, instance, their cloned version and cloned instance.
type CloneRecord struct {
	OriginAstID    string
	CloneAstID     string
	OriginInstance interface{}
	CloneInstance  interface{}
}

// NewCloneTable create new instance of CloneTable
func NewCloneTable() *CloneTable {
	return &CloneTable{Records: make(map[string]*CloneRecord)}
}

// CloneTable will stores all meta information about AST object being cloned under one KnowledgeBase.
type CloneTable struct {
	Records map[string]*CloneRecord
}

// IsCloned will check if an AST object with identified astId has a clone.
func (tab *CloneTable) IsCloned(astID string) bool {
	_, ok := tab.Records[astID]
	return ok
}

// MarkCloned will record that an Ast object are now been cloned, so all other cloned object should reference to the newly cloned Ast object
func (tab *CloneTable) MarkCloned(originAst, cloneAst string, origin, clone interface{}) {
	tab.Records[originAst] = &CloneRecord{
		OriginAstID:    originAst,
		CloneAstID:     cloneAst,
		OriginInstance: origin,
		CloneInstance:  clone,
	}
}
