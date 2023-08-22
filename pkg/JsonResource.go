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

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"strconv"
	"strings"
)

// GruleJSON represents a rule in JSON format
type GruleJSON struct {
	Name        string        `json:"name"`
	Description string        `json:"desc"`
	Salience    int           `json:"salience"`
	When        interface{}   `json:"when"`
	Then        []interface{} `json:"then"`
}

// JSONResource will parse rules in JSON fromat from underlying resource provider.
type JSONResource struct {
	subRes Resource
}

// JSONResourceBundle will parse a set of rules in JSON format from an underlying bundle resource provider.
type JSONResourceBundle struct {
	subRes ResourceBundle
}

// NewJSONResourceFromResource innstantiates a new JSON resource parser from an underlying Resource.
func NewJSONResourceFromResource(res Resource) (Resource, error) {
	if _, ok := res.(*JSONResource); ok {

		return nil, fmt.Errorf("not a JSONResource")
	}

	return &JSONResource{
		subRes: res,
	}, nil
}

// Load will load the underlying Resource and parse the JSON rules into standard GRule syntax.
func (jr *JSONResource) Load() ([]byte, error) {
	data, err := jr.subRes.Load()
	if err != nil {

		return nil, err
	}
	firstRune := string(bytes.TrimSpace(data)[0])

	var ruleSet string

	if firstRune == "[" {
		ruleSet, err = ParseJSONRuleset(data)
	} else if firstRune == "{" {
		ruleSet, err = ParseJSONRule(data)
	} else {
		err = errors.New("invalid JSON input")
	}
	if err != nil {

		return nil, err
	}

	return []byte(ruleSet), nil
}

// String will state the resource source.
func (jr *JSONResource) String() string {

	return "JSON Resource, underlying resource: " + jr.subRes.String()
}

// NewJSONResourceBundleFromBundle innstantiates a new bundled JSON resource parser from an underlying ResourceBundle.
func NewJSONResourceBundleFromBundle(bundle ResourceBundle) (ResourceBundle, error) {
	if _, ok := bundle.(*JSONResourceBundle); ok {

		return nil, fmt.Errorf("bundle is not JSONResourceBundle")
	}

	return &JSONResourceBundle{
		subRes: bundle,
	}, nil
}

// Load will load the underlying ResourceBundle and parse the JSON rules into standard GRule syntax.
func (jrb *JSONResourceBundle) Load() ([]Resource, error) {
	ress, err := jrb.subRes.Load()
	if err != nil {

		return nil, err
	}
	nress := make([]Resource, len(ress))
	for i := 0; i < len(ress); i++ {
		nress[i], err = NewJSONResourceFromResource(ress[i])
		if err != nil {

			return nil, err
		}
	}

	return nress, nil
}

// MustLoad operates the same as load except it will panic in the event of an error.
func (jrb *JSONResourceBundle) MustLoad() []Resource {
	ress := jrb.subRes.MustLoad()
	nress := make([]Resource, len(ress))
	for i := 0; i < len(ress); i++ {
		resour, err := NewJSONResourceFromResource(ress[i])
		if err != nil {

			panic(err)
		}
		nress[i] = resour
	}

	return nress
}

// ParseJSONRuleset accepts a byte array containing an array of rules in JSON format to be parsed into GRule syntax.
func ParseJSONRuleset(data []byte) (rs string, err error) {
	defer func() {
		if x := recover(); x != nil {
			err = fmt.Errorf("%v", x)
		}
	}()
	var rules []GruleJSON
	err = json.Unmarshal(data, &rules)
	if err != nil {

		return
	}
	var sb strings.Builder
	for i := 0; i < len(rules); i++ {
		rName, err := parseRule(&rules[i])
		if err != nil {

			return rName, err
		}
		sb.WriteString(rName)
	}
	rs = sb.String()

	return
}

// ParseJSONRule accepts a byte array containing an rule in JSON format to be parsed into GRule syntax.
func ParseJSONRule(data []byte) (rs string, err error) {
	defer func() {
		if x := recover(); x != nil {
			err = fmt.Errorf("%v", x)
		}
	}()
	var rule GruleJSON
	err = json.Unmarshal(data, &rule)
	if err != nil {

		return
	}

	return parseRule(&rule)
}

// ParseRule Accepts a struct of GruleJSON rule and returns the parsed string of GRule.
func ParseRule(rule *GruleJSON) (r string, err error) {
	defer func() {
		if x := recover(); x != nil {
			err = fmt.Errorf("%v", x)
		}
	}()

	return parseRule(rule)
}

func parseRule(rule *GruleJSON) (string, error) {
	if len(rule.Name) == 0 {
		return "", fmt.Errorf("encountered a rule without name")
	}
	if rule.When == nil {

		panic("rule when condition cannot be nil")
	}
	if rule.Then == nil {

		panic("rule thenn condition cannot be nil")
	}
	var stringBuilder strings.Builder
	stringBuilder.WriteString("rule ")
	stringBuilder.WriteString(rule.Name)
	stringBuilder.WriteString(" ")
	stringBuilder.WriteString(strconv.Quote(rule.Description))
	stringBuilder.WriteString(" salience ")
	stringBuilder.WriteString(strconv.Itoa(rule.Salience))
	stringBuilder.WriteString(" {\n    when\n        ")
	stringBuilder.WriteString(parseWhen(rule.When))
	stringBuilder.WriteString("\n    then\n")
	thens := parseThen(rule.Then)
	for i := 0; i < len(thens); i++ {
		stringBuilder.WriteString("        ")
		stringBuilder.WriteString(thens[i])
		stringBuilder.WriteString("\n")
	}
	stringBuilder.WriteString("}\n")

	return stringBuilder.String(), nil
}

func parseThen(ts []interface{}) []string {
	thens := make([]string, len(ts))
	for thenItem := 0; thenItem < len(ts); thenItem++ {
		switch thenType := ts[thenItem].(type) {
		case string:
			thens[thenItem] = thenType
			if !strings.HasSuffix(thens[thenItem], ";") {
				thens[thenItem] += ";"
			}
		case map[string]interface{}:
			thens[thenItem] = buildExpression(thenType, 0) + ";"
		default:

			panic("invalid then type, must be a string or an array of action objects")
		}
	}

	return thens
}

func parseWhen(w interface{}) string {
	switch whenType := w.(type) {
	case string:

		return whenType
	case map[string]interface{}:

		return buildExpression(whenType, 0)
	default:

		panic("invalid when type, must be a string or an array of condition objects")
	}
}

func buildExpression(input map[string]interface{}, depth int) string {
	exp, _ := buildExpressionEx(input, depth)

	return exp
}

func buildExpressionEx(input map[string]interface{}, depth int) (string, bool) {
	if depth > 1024 {

		panic("JSON nesting exceeded 1024 levels, aborting")
	}
	if len(input) > 1 {

		panic("expression objects can only contain a single operation type")
	}
	for key, value := range input {
		switch key {
		case "and":

			return buildCompoundOperator(value, depth, " && ")
		case "or":

			return buildCompoundOperator(value, depth, " || ")
		case "eq":

			return joinOperator(value, " == "), false
		case "not":

			return joinOperator(value, " != "), false
		case "gt":

			return joinOperator(value, " > "), false
		case "gte":

			return joinOperator(value, " >= "), false
		case "lt":

			return joinOperator(value, " < "), false
		case "lte":

			return joinOperator(value, " <= "), false
		case "bor":

			return joinOperator(value, " | "), false
		case "band":

			return joinOperator(value, " & "), false
		case "plus":

			return joinOperator(value, " + "), false
		case "minus":

			return joinOperator(value, " - "), false
		case "div":

			return joinOperator(value, " / "), false
		case "mul":

			return joinOperator(value, " * "), false
		case "mod":

			return joinOperator(value, " % "), false
		case "set":

			return joinSet(value, " = "), true
		case "call":

			return joinCall(value), true
		case "obj":
			if s, ok := value.(string); ok {

				return s, true
			}

			panic("object must be a string")
		case "const":
			switch valueType := value.(type) {
			case string:

				return strconv.Quote(valueType), true
			case float64:

				return strconv.FormatFloat(valueType, 'f', -1, 64), true
			case bool:
				if valueType {

					return "true", true
				}

				return "false", true
			}

			panic("constant must be a string or a numeric value")
		default:

			panic("unknown operator type: " + key)
		}
	}

	panic("boolean expression cannot be empty")
}

func buildCompoundOperator(o interface{}, depth int, operator string) (string, bool) {
	if andarr, ok := o.([]interface{}); ok {
		var ands []string
		if len(andarr) < 2 {

			panic("and operator must have at least 2 operands")
		}
		for i := 0; i < len(andarr); i++ {
			if subVal, ok := andarr[i].(map[string]interface{}); ok {
				ands = append(ands, buildExpression(subVal, depth+1))
			} else {

				panic("and operands must be an array of objects")
			}
		}
		if depth > 0 {

			return "(" + strings.Join(ands, operator) + ")", false
		}

		return strings.Join(ands, operator), false
	}

	panic("compound operator must be an array")
}

func joinCall(v interface{}) string {
	if arr, ok := v.([]interface{}); ok {
		if len(arr) == 0 {

			panic("call operator must have at least one operand")
		}
		var firstCallOperand string
		var ok bool
		if firstCallOperand, ok = arr[0].(string); !ok {

			panic("first call operand must be a string")
		}
		if len(arr) > 1 {
			sars := make([]string, len(arr)-1)
			for i := 1; i < len(arr); i++ {
				sars[i-1] = parseCallOperand(arr[i])
			}

			return firstCallOperand + "(" + strings.Join(sars, ", ") + ")"
		}

		return firstCallOperand + "()"
	}

	panic("operator has an unexpected type")
}

func parseCallOperand(o interface{}) string {
	switch operandType := o.(type) {
	case string:
		if len(operandType) == 0 {
			panic("operand cannnot be empty")
		}

		return operandType
	case float64:

		return fmt.Sprint(operandType)
	case bool:
		if operandType {

			return "true"
		}

		return "false"
	case map[string]interface{}:

		return buildExpression(operandType, 0)
	default:

		panic("operand has an invalid type")
	}
}

func joinOperator(v interface{}, operator string) string {
	if arr, ok := v.([]interface{}); ok {
		if len(arr) == 0 {

			panic("operator cannot have 0 operands")
		}
		ops := make([]string, len(arr))
		for i := 0; i < len(arr); i++ {
			ops[i] = parseOperand(arr[i], false)
		}

		return strings.Join(ops, operator)
	}

	panic("operator has an unexpected type")
}

func joinSet(v interface{}, operator string) string {
	if arr, ok := v.([]interface{}); ok {
		if len(arr) != 2 {

			panic("set operand count must be 2")
		}

		return parseOperand(arr[0], true) + operator + parseOperand(arr[1], true)
	}

	panic("operator has an unexpected type")
}

func parseOperand(o interface{}, noWrap bool) string {
	switch operandType := o.(type) {
	case string:

		return operandType
	case float64:

		return fmt.Sprint(operandType)
	case bool:

		if operandType {

			return "true"
		}

		return "false"
	case map[string]interface{}:
		expr, expNoWrap := buildExpressionEx(operandType, 0)
		if expNoWrap || noWrap {

			return expr
		}

		return "(" + expr + ")"
	default:

		panic("operand has an invalid type")
	}
}
