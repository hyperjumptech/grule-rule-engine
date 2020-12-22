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

//go:generate mockgen -destination=../mocks/ast/DataContext.go -package=mocksAst . IDataContext

import (
	"github.com/hyperjumptech/grule-rule-engine/model"
	"reflect"
)

// NewDataContext will create a new DataContext instance
func NewDataContext() IDataContext {
	return &DataContext{
		ObjectStore: make(map[string]model.ValueNode),

		retracted:           make([]string, 0),
		variableChangeCount: 0,
	}
}

// DataContext holds all structs instance to be used in rule execution environment.
type DataContext struct {
	ObjectStore map[string]model.ValueNode

	retracted           []string
	variableChangeCount uint64
	complete            bool
}

// Complete marks the DataContext as completed, telling the engine to stop processing rules
func (ctx *DataContext) Complete() {
	ctx.complete = true
}

// IsComplete checks whether the DataContext has been completed
func (ctx *DataContext) IsComplete() bool {
	return ctx.complete
}

// IDataContext is the interface for the DataContext struct.
type IDataContext interface {
	ResetVariableChangeCount()
	IncrementVariableChangeCount()
	HasVariableChange() bool

	Add(key string, obj interface{}) error
	AddJSON(key string, JSON []byte) error
	Get(key string) model.ValueNode

	Retract(key string)
	IsRetracted(key string) bool
	Complete()
	IsComplete() bool
	Retracted() []string
	Reset()
}

// ResetVariableChangeCount will reset the variable change count
func (ctx *DataContext) ResetVariableChangeCount() {
	ctx.variableChangeCount = 0
}

// IncrementVariableChangeCount will increment the variable change count
func (ctx *DataContext) IncrementVariableChangeCount() {
	ctx.variableChangeCount++
}

// HasVariableChange returns true if there are variable changes
func (ctx *DataContext) HasVariableChange() bool {
	return ctx.variableChangeCount > 0
}

// Add will add struct instance into rule execution context
func (ctx *DataContext) Add(key string, obj interface{}) error {
	ctx.ObjectStore[key] = model.NewGoValueNode(reflect.ValueOf(obj), key)
	return nil
}

// AddJSON will add struct instance into rule execution context
func (ctx *DataContext) AddJSON(key string, JSON []byte) error {
	vn, err := model.NewJSONValueNode(string(JSON), key)
	if err != nil {
		return err
	}
	ctx.ObjectStore[key] = vn
	return nil
}

// Get will extract the struct instance
func (ctx *DataContext) Get(key string) model.ValueNode {
	if v, ok := ctx.ObjectStore[key]; ok {
		return v
	}
	return nil
}

// Retract temporary retract a fact from data context, making it unavailable for evaluation or modification.
func (ctx *DataContext) Retract(key string) {
	ctx.retracted = append(ctx.retracted, key)
}

// IsRetracted checks if a key fact is currently retracted.
func (ctx *DataContext) IsRetracted(key string) bool {
	for _, v := range ctx.retracted {
		if v == key {
			return true
		}
	}
	return false
}

// Retracted returns list of retracted key facts.
func (ctx *DataContext) Retracted() []string {
	return ctx.retracted
}

// Reset will un-retract all fact, making them available for evaluation and modification.
func (ctx *DataContext) Reset() {
	ctx.retracted = make([]string, 0)
}
