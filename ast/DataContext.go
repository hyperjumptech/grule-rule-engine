package ast

//go:generate mockgen -destination=../mocks/ast/DataContext.go -package=mocksAst . IDataContext

import (
	"fmt"
	"github.com/hyperjumptech/grule-rule-engine/pkg"
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
	Get(key string) interface{}

	Retract(key string)
	IsRetracted(key string) bool
	Complete()
	IsComplete() bool
	Retracted() []string
	Reset()

	ExecMethod(receiver reflect.Value, methodName string, args []reflect.Value) (reflect.Value, error)
	GetType(receiver reflect.Value, variable string) (reflect.Type, error)
	GetValue(receiver reflect.Value, variable string) (reflect.Value, error)
	SetValue(receiver reflect.Value, variable string, newValue reflect.Value) error
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
func (ctx *DataContext) Get(key string) interface{} {
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

// ExecMethod will execute instance member variable using the supplied arguments.
func (ctx *DataContext) ExecMethod(receiver reflect.Value, methodName string, args []reflect.Value) (reflect.Value, error) {
	switch pkg.GetBaseKind(receiver) {
	case reflect.Int64, reflect.Uint64, reflect.Float64, reflect.Bool:
		return reflect.ValueOf(nil), fmt.Errorf("function %s is not supported for type %s", methodName, receiver.Type().String())
	case reflect.String:
		var strfunc func(string, []reflect.Value) (reflect.Value, error)
		switch methodName {
		case "Compare":
			strfunc = StrCompare
		case "Contains":
			strfunc = StrContains
		case "Count":
			strfunc = StrCount
		case "HasPrefix":
			strfunc = StrHasPrefix
		case "HasSuffix":
			strfunc = StrHasSuffix
		case "Index":
			strfunc = StrIndex
		case "LastIndex":
			strfunc = StrLastIndex
		case "Repeat":
			strfunc = StrRepeat
		case "Replace":
			strfunc = StrReplace
		case "Split":
			strfunc = StrSplit
		case "ToLower":
			strfunc = StrToLower
		case "ToUpper":
			strfunc = StrToUpper
		case "Trim":
			strfunc = StrTrim
		}
		if strfunc != nil {
			val, err := strfunc(receiver.String(), args)
			if err != nil {
				return reflect.ValueOf(nil), err
			}
			return val, nil
		}
		return reflect.ValueOf(nil), fmt.Errorf("function %s is not supported for string", methodName)
	}

	// this obj is reflect.Value... it should not.
	types, variad, err := pkg.GetFunctionParameterTypes(receiver, methodName)
	if err != nil {
		return reflect.ValueOf(nil),
			fmt.Errorf("error while fetching function %s() parameter types. Got %v", methodName, err)
	}

	if len(types) != len(args) && !variad {
		return reflect.ValueOf(nil),
			fmt.Errorf("invalid argument count for function %s(). need %d argument while there are %d", methodName, len(types), len(args))
	}

	retVals, err := pkg.InvokeFunction(receiver, methodName, args)
	if err == nil {
		if len(retVals) == 0 {
			return reflect.ValueOf(nil), nil
		}
		return retVals[0], nil
	}
	return reflect.ValueOf(nil), err
}

// GetType will extract type information of data in this context.
func (ctx *DataContext) GetType(receiver reflect.Value, variable string) (reflect.Type, error) {
	return pkg.GetAttributeType(receiver, variable)
}

// GetValue will get member variables Value information.
// Used by the rule execution to obtain variable value.
func (ctx *DataContext) GetValue(receiver reflect.Value, variable string) (reflect.Value, error) {
	return pkg.GetAttributeValue(receiver, variable)
}

// SetValue will set variable value of an object instance in this data context, Used by rule script to set values.
func (ctx *DataContext) SetValue(receiver reflect.Value, variable string, newValue reflect.Value) error {
	return pkg.SetAttributeValue(receiver, variable, newValue)
}
