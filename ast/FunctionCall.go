package ast

import (
	"bytes"
	"fmt"
	"reflect"
	"strings"

	"github.com/google/uuid"
	"github.com/hyperjumptech/grule-rule-engine/pkg"
	log "github.com/sirupsen/logrus"
)

// NewFunctionCall creates new instance of FunctionCall
func NewFunctionCall() *FunctionCall {
	return &FunctionCall{
		AstID:        uuid.New().String(),
		ArgumentList: NewArgumentList(),
	}
}

// FunctionCall AST graph node
type FunctionCall struct {
	AstID         string
	GrlText       string
	DataContext   IDataContext
	WorkingMemory *WorkingMemory

	FunctionName string
	ArgumentList *ArgumentList
	Value        reflect.Value
}

// Clone will clone this FunctionCall. The new clone will have an identical structure
func (e FunctionCall) Clone(cloneTable *pkg.CloneTable) *FunctionCall {
	clone := &FunctionCall{
		AstID:         uuid.New().String(),
		GrlText:       e.GrlText,
		DataContext:   nil,
		WorkingMemory: nil,
		FunctionName:  e.FunctionName,
	}

	if e.ArgumentList != nil {
		if cloneTable.IsCloned(e.ArgumentList.AstID) {
			clone.ArgumentList = cloneTable.Records[e.ArgumentList.AstID].CloneInstance.(*ArgumentList)
		} else {
			cloned := e.ArgumentList.Clone(cloneTable)
			clone.ArgumentList = cloned
			cloneTable.MarkCloned(e.ArgumentList.AstID, cloned.AstID, e.ArgumentList, cloned)
		}
	}
	return clone
}

// InitializeContext will initialize this AST graph with data context and working memory before running rule on them.
func (e *FunctionCall) InitializeContext(dataCtx IDataContext, WorkingMemory *WorkingMemory) {
	e.DataContext = dataCtx
	e.WorkingMemory = WorkingMemory
	if e.ArgumentList != nil {
		e.ArgumentList.InitializeContext(dataCtx, WorkingMemory)
	}
}

// FunctionCallReceiver should be implemented bu AST graph node to receive a FunctionCall AST graph mode
type FunctionCallReceiver interface {
	AcceptFunctionCall(fun *FunctionCall) error
}

// GetAstID get the UUID asigned for this AST graph node
func (e *FunctionCall) GetAstID() string {
	return e.AstID
}

// GetGrlText get the expression syntax related to this graph when it wast constructed
func (e *FunctionCall) GetGrlText() string {
	return e.GrlText
}

// GetSnapshot will create a structure signature or AST graph
func (e *FunctionCall) GetSnapshot() string {
	var buff bytes.Buffer
	if e.ArgumentList == nil {
		log.Errorf("Argument is nil")
	} else {
		if e.ArgumentList.Arguments == nil {
			log.Errorf("Argument.Argumeent array is nil")
		}
	}
	buff.WriteString(fmt.Sprintf("func->%s%s", e.FunctionName, e.ArgumentList.GetSnapshot()))
	return buff.String()
}

// SetGrlText set the expression syntax related to this graph when it was constructed. Only ANTLR4 listener should
// call this function.
func (e *FunctionCall) SetGrlText(grlText string) {
	e.GrlText = grlText
}

// AcceptArgumentList will accept an ArgumentList AST graph into this ast graph
func (e *FunctionCall) AcceptArgumentList(argList *ArgumentList) error {
	log.Tracef("Method received argument list")
	e.ArgumentList = argList
	return nil
}

// Evaluate will evaluate this AST graph for when scope evaluation
func (e *FunctionCall) Evaluate(receiver reflect.Value) (reflect.Value, error) {
	args, err := e.ArgumentList.Evaluate()
	if err != nil {
		return reflect.ValueOf(nil), err
	}

	switch pkg.GetBaseKind(receiver) {
	case reflect.Int64, reflect.Uint64, reflect.Float64, reflect.Bool:
		return reflect.ValueOf(nil), fmt.Errorf("function %s is not supported for type %s", e.FunctionName, receiver.Type().String())
	case reflect.String:
		var strfunc func(string, []reflect.Value) (reflect.Value, error)
		switch e.FunctionName {
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
			e.Value = val
			return val, nil
		}
		return reflect.ValueOf(nil), fmt.Errorf("function %s is not supported for string", e.FunctionName)
	}

	// this obj is reflect.Value... it should not.
	types, variad, err := pkg.GetFunctionParameterTypes(receiver, e.FunctionName)
	if err != nil {
		return reflect.ValueOf(nil),
			fmt.Errorf("error while fetching function %s() parameter types. Got %v", e.FunctionName, err)
	}

	if len(types) != len(args) && !variad {
		return reflect.ValueOf(nil),
			fmt.Errorf("invalid argument count for function %s(). need %d argument while there are %d", e.FunctionName, len(types), len(args))
	}

	iargs := make([]interface{}, len(args))
	for i, t := range types {
		if variad && i == len(types)-1 {
			break
		}
		if t.Kind() != args[i].Kind() {
			if t.Kind() == reflect.Interface {
				iargs[i] = pkg.ValueToInterface(args[i])
			} else {
				return reflect.ValueOf(nil),
					fmt.Errorf("invalid argument types for function %s(). argument #%d, require %s but %s", e.FunctionName, i, t.Kind().String(), args[i].Kind().String())
			}
		} else {
			iargs[i] = pkg.ValueToInterface(args[i])
		}
	}

	retVals, err := pkg.InvokeFunction(pkg.ValueToInterface(receiver), e.FunctionName, iargs)
	//retVal, err := e.DataContext.ExecMethod(objName, args)
	if err == nil {
		e.Value = reflect.ValueOf(retVals[0])
		return e.Value, nil
	}
	return reflect.ValueOf(nil), err
}

func StrCompare(str string, arg []reflect.Value) (reflect.Value, error) {
	if arg == nil || len(arg) != 1 || arg[0].Kind() != reflect.String {
		return reflect.ValueOf(nil), fmt.Errorf("function Compare requires 1 string argument")
	}
	i := strings.Compare(str, arg[0].String())
	return reflect.ValueOf(i), nil
}

func StrContains(str string, arg []reflect.Value) (reflect.Value, error) {
	if arg == nil || len(arg) != 1 || arg[0].Kind() != reflect.String {
		return reflect.ValueOf(nil), fmt.Errorf("function Contains requires 1 string argument")
	}

	i := strings.Contains(str, arg[0].String())
	return reflect.ValueOf(i), nil
}

func StrCount(str string, arg []reflect.Value) (reflect.Value, error) {
	if arg == nil || len(arg) != 1 || arg[0].Kind() != reflect.String {
		return reflect.ValueOf(nil), fmt.Errorf("function Count requires 1 string argument")
	}

	i := strings.Count(str, arg[0].String())
	return reflect.ValueOf(i), nil
}

func StrHasPrefix(str string, arg []reflect.Value) (reflect.Value, error) {
	if arg == nil || len(arg) != 1 || arg[0].Kind() != reflect.String {
		return reflect.ValueOf(nil), fmt.Errorf("function HasPrefix requires 1 string argument")
	}

	b := strings.HasPrefix(str, arg[0].String())
	return reflect.ValueOf(b), nil
}

func StrHasSuffix(str string, arg []reflect.Value) (reflect.Value, error) {
	if arg == nil || len(arg) != 1 || arg[0].Kind() != reflect.String {
		return reflect.ValueOf(nil), fmt.Errorf("function HasSuffix requires 1 string argument")
	}

	b := strings.HasSuffix(str, arg[0].String())
	return reflect.ValueOf(b), nil
}

func StrIndex(str string, arg []reflect.Value) (reflect.Value, error) {
	if arg == nil || len(arg) != 1 || arg[0].Kind() != reflect.String {
		return reflect.ValueOf(nil), fmt.Errorf("function Index requires 1 string argument")
	}

	b := strings.Index(str, arg[0].String())
	return reflect.ValueOf(b), nil
}

func StrLastIndex(str string, arg []reflect.Value) (reflect.Value, error) {
	if arg == nil || len(arg) != 1 || arg[0].Kind() != reflect.String {
		return reflect.ValueOf(nil), fmt.Errorf("function LastIndex requires 1 string argument")
	}

	b := strings.LastIndex(str, arg[0].String())
	return reflect.ValueOf(b), nil
}

func StrRepeat(str string, arg []reflect.Value) (reflect.Value, error) {
	if arg == nil || len(arg) != 1 {
		return reflect.ValueOf(nil), fmt.Errorf("function Repeat requires 1 numeric argument")
	}
	repeat := 0
	switch pkg.GetBaseKind(arg[0]) {
	case reflect.Int64:
		repeat = int(arg[0].Int())
	case reflect.Uint64:
		repeat = int(arg[0].Uint())
	case reflect.Float64:
		repeat = int(arg[0].Float())
	default:
		return reflect.ValueOf(nil), fmt.Errorf("function Repeat requires 1 numeric argument")
	}

	b := strings.Repeat(str, repeat)
	return reflect.ValueOf(b), nil
}

func StrReplace(str string, arg []reflect.Value) (reflect.Value, error) {
	if arg == nil || len(arg) != 2 || arg[0].Kind() != reflect.String || arg[1].Kind() != reflect.String {
		return reflect.ValueOf(nil), fmt.Errorf("function Cmpare requires 2 string argument")
	}

	b := strings.ReplaceAll(str, arg[0].String(), arg[1].String())
	return reflect.ValueOf(b), nil
}

func StrSplit(str string, arg []reflect.Value) (reflect.Value, error) {
	if arg == nil || len(arg) != 1 || arg[0].Kind() != reflect.String {
		return reflect.ValueOf(nil), fmt.Errorf("function Split requires 1 string argument")
	}

	b := strings.Split(str, arg[0].String())
	return reflect.ValueOf(b), nil
}

func StrToLower(str string, arg []reflect.Value) (reflect.Value, error) {
	if arg == nil || len(arg) != 0 {
		return reflect.ValueOf(nil), fmt.Errorf("function ToLower requires no argument")
	}
	b := strings.ToLower(str)
	return reflect.ValueOf(b), nil
}

func StrToUpper(str string, arg []reflect.Value) (reflect.Value, error) {
	if arg == nil || len(arg) != 0 {
		return reflect.ValueOf(nil), fmt.Errorf("function ToUpper requires no argument")
	}
	b := strings.ToUpper(str)
	return reflect.ValueOf(b), nil

}

func StrTrim(str string, arg []reflect.Value) (reflect.Value, error) {
	if arg == nil || len(arg) != 0 {
		return reflect.ValueOf(nil), fmt.Errorf("function Trim requires no argument")
	}
	b := strings.TrimSpace(str)
	return reflect.ValueOf(b), nil
}
