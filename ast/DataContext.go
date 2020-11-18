package ast

//go:generate mockgen -destination=../mocks/ast/DataContext.go -package=mocksAst . IDataContext

import (
	"github.com/hyperjumptech/grule-rule-engine/model"
	"reflect"
)

// NewDataContext will create a new DataContext instance
func NewDataContext() IDataContext {
	return &DataContext{
		ObjectStore: make(map[string]interface{}),

		retracted:           make([]string, 0),
		variableChangeCount: 0,
	}
}

// DataContext holds all structs instance to be used in rule execution environment.
type DataContext struct {
	ObjectStore map[string]interface{}

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
	// LETS experiment by disabling this. We can add non struct pointer as fact.
	// Because now we can extract value directly into graph.
	//
	//objVal := reflect.ValueOf(obj)
	//if objVal.Kind() != reflect.Ptr || objVal.Elem().Kind() != reflect.Struct {
	//	return fmt.Errorf("you can only insert a pointer to struct as fact. objVal = %s", objVal.Kind().String())
	//}
	ctx.ObjectStore[key] = obj
	return nil
}

// Get will extract the struct instance
func (ctx *DataContext) Get(key string) model.ValueNode {
	if v, ok := ctx.ObjectStore[key]; ok {
		return model.NewGoValueNode(reflect.ValueOf(v), key)
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
