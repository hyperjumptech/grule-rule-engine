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
	AstID   string
	GrlText string

	FunctionName string
	ArgumentList *ArgumentList
	Value        reflect.Value
}

// Clone will clone this FunctionCall. The new clone will have an identical structure
func (e *FunctionCall) Clone(cloneTable *pkg.CloneTable) *FunctionCall {
	clone := &FunctionCall{
		AstID:        uuid.New().String(),
		GrlText:      e.GrlText,
		FunctionName: e.FunctionName,
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
	buff.WriteString(FUNCTIONCALL)
	buff.WriteString(fmt.Sprintf("(n:%s", e.FunctionName))
	if e.ArgumentList == nil {
		log.Errorf("Argument is nil")
	} else {
		buff.WriteString(",")
		buff.WriteString(e.ArgumentList.GetSnapshot())
	}
	buff.WriteString(")")
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
func (e *FunctionCall) Evaluate(receiver reflect.Value, dataContext IDataContext, memory *WorkingMemory) (reflect.Value, error) {
	args, err := e.ArgumentList.Evaluate(dataContext, memory)
	if err != nil {
		return reflect.ValueOf(nil), err
	}
	if dataContext == nil {
		AstLog.Errorf("Datacontext for function call %s (%s) is nil", e.FunctionName, e.AstID)
	}
	return dataContext.ExecMethod(receiver, e.FunctionName, args)
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
