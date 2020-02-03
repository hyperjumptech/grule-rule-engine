package ast

import (
	"fmt"
	"github.com/hyperjumptech/grule-rule-engine/pkg"
	"github.com/juju/errors"
	"reflect"
	"strings"
)

// NewDataContext will create a new DataContext instance
func NewDataContext() *DataContext {
	return &DataContext{
		ObjectStore: make(map[string]interface{}),
		Retracted:   make([]string, 0),
	}
}

// DataContext holds all structs instance to be used in rule execution environment.
type DataContext struct {
	ObjectStore         map[string]interface{}
	Retracted           []string
	VariableChangeCount uint64
}

// Retract temporary retract a fact from data context, making it unavailable for evaluation or modification.
func (ctx *DataContext) Retract(key string) {
	ctx.Retracted = append(ctx.Retracted, key)
}

// Add will add struct instance into rule execution context
func (ctx *DataContext) Add(key string, obj interface{}) error {
	objVal := reflect.ValueOf(obj)
	if objVal.Kind() != reflect.Ptr && objVal.Elem().Kind() != reflect.Struct {
		return errors.New(fmt.Sprintf("you can only insert a pointer to struct as fact. objVal = %s", objVal.Kind().String()))
	}
	ctx.ObjectStore[key] = obj
	return nil
}

// IsRestracted checks if a key fact is currently retracted.
func (ctx *DataContext) IsRestracted(key string) bool {
	for _, v := range ctx.Retracted {
		if v == key {
			return true
		}
	}
	return false
}

// Reset will un-retract all fact, making them available for evaluation and modification.
func (ctx *DataContext) Reset() {
	ctx.Retracted = make([]string, 0)
}

// ExecMethod will execute instance member variable using the supplied arguments.
func (ctx *DataContext) ExecMethod(methodName string, args []reflect.Value) (reflect.Value, error) {
	varArray := strings.Split(methodName, ".")
	if val, ok := ctx.ObjectStore[varArray[0]]; ok {
		if !ctx.IsRestracted(varArray[0]) {
			return traceMethod(val, varArray[1:], args)
		}
		return reflect.ValueOf(nil), errors.New("fact is retracted")
	}
	return reflect.ValueOf(nil), fmt.Errorf("fact [%s] not found while execute method", varArray[0])
}

// GetType will extract type information of data in this context.
func (ctx *DataContext) GetType(variable string) (reflect.Type, error) {
	varArray := strings.Split(variable, ".")
	if val, ok := ctx.ObjectStore[varArray[0]]; ok {
		if !ctx.IsRestracted(varArray[0]) {
			return traceType(val, varArray[1:])
		}
		return nil, errors.New("fact is retracted")
	}
	return nil, fmt.Errorf("fact [%s] not found while obtaining type", variable)
}

// GetValue will get member variables Value information.
// Used by the rule execution to obtain variable value.
func (ctx *DataContext) GetValue(variable string) (reflect.Value, error) {
	varArray := strings.Split(variable, ".")
	if val, ok := ctx.ObjectStore[varArray[0]]; ok {
		if !ctx.IsRestracted(varArray[0]) {
			vval, err := traceValue(val, varArray[1:])
			if err != nil {
				fmt.Printf("blah %s = %v\n", variable, vval)
			}
			return vval, err
		}
		return reflect.ValueOf(nil), errors.New("fact is retracted")
	}
	return reflect.ValueOf(nil), fmt.Errorf("fact [%s] not found while retrieving value", varArray[0])
}

// SetValue will set variable value of an object instance in this data context, Used by rule script to set values.
func (ctx *DataContext) SetValue(variable string, newValue reflect.Value) error {
	varArray := strings.Split(variable, ".")
	if val, ok := ctx.ObjectStore[varArray[0]]; ok {
		if !ctx.IsRestracted(varArray[0]) {
			err := traceSetValue(val, varArray[1:], newValue)
			if err == nil {
				ctx.VariableChangeCount++
			}
			return err
		}
		return errors.New("fact is retracted")
	}
	return fmt.Errorf("fact [%s] not found while setting value", varArray[0])
}

func traceType(obj interface{}, path []string) (reflect.Type, error) {
	switch length := len(path); {
	case length == 1:
		return pkg.GetAttributeType(obj, path[0])
	case length > 1:
		objVal, err := pkg.GetAttributeValue(obj, path[0])
		if err != nil {
			return nil, errors.Trace(err)
		}
		return traceType(pkg.ValueToInterface(objVal), path[1:])
	default:
		return reflect.TypeOf(obj), nil
	}
}

func traceValue(obj interface{}, path []string) (reflect.Value, error) {
	switch length := len(path); {
	case length == 1:
		return pkg.GetAttributeValue(obj, path[0])
	case length > 1:
		objVal, err := pkg.GetAttributeValue(obj, path[0])
		if err != nil {
			return objVal, errors.Trace(err)
		}
		return traceValue(pkg.ValueToInterface(objVal), path[1:])
	default:
		return reflect.ValueOf(obj), nil
	}
}

func traceSetValue(obj interface{}, path []string, newValue reflect.Value) error {
	switch length := len(path); {
	case length == 1:
		return pkg.SetAttributeValue(obj, path[0], newValue)
	case length > 1:
		objVal, err := pkg.GetAttributeValue(obj, path[0])
		if err != nil {
			return errors.Trace(err)
		}
		return traceSetValue(objVal, path[1:], newValue)
	default:
		return errors.Errorf("no attribute path specified")
	}
}

func traceMethod(obj interface{}, path []string, args []reflect.Value) (reflect.Value, error) {

	switch length := len(path); {
	case length == 1:
		// this obj is reflect.Value... it should not.
		types, err := pkg.GetFunctionParameterTypes(obj, path[0])
		if err != nil {
			return reflect.ValueOf(nil),
				errors.Errorf("error while fetching function %s() parameter types. Got %v", path[0], err)
		}
		if len(types) != len(args) {
			return reflect.ValueOf(nil),
				errors.Errorf("invalid argument count for function %s(). need %d argument while there are %d", path[0], len(types), len(args))
		}
		iargs := make([]interface{}, 0)
		for i, t := range types {
			if t.Kind() != args[i].Kind() {
				if t.Kind() == reflect.Interface {
					iargs = append(iargs, pkg.ValueToInterface(args[i]))
				} else {
					return reflect.ValueOf(nil),
						errors.Errorf("invalid argument types for function %s(). argument #%d, require %s but %s", path[0], i, t.Kind().String(), args[i].Kind().String())
				}
			} else {
				iargs = append(iargs, pkg.ValueToInterface(args[i]))
			}
		}
		rets, err := pkg.InvokeFunction(obj, path[0], iargs)
		if err != nil {
			return reflect.ValueOf(nil), err
		}
		switch retLen := len(rets); {
		case retLen > 1:
			return reflect.ValueOf(rets[0]), errors.Errorf("multiple return value for function %s(). ", path[0])
		case retLen == 1:
			return reflect.ValueOf(rets[0]), nil
		default:
			return reflect.ValueOf(nil), nil
		}
	case length > 1:
		objVal, err := pkg.GetAttributeValue(obj, path[0])
		if err != nil {
			return reflect.ValueOf(nil), errors.Trace(err)
		}
		return traceMethod(objVal, path[1:], args)
	default:
		return reflect.ValueOf(nil), errors.Errorf("no function path specified")
	}
}
