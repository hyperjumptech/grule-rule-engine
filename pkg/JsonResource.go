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
	subRes ResouceBundle
}

// NewJSONResourceFromResource innstantiates a new JSON resource parser from an underlying Resource.
func NewJSONResourceFromResource(res Resource) Resource {
	if _, ok := res.(*JSONResource); ok {
		panic("cannot create JSON resource from JSON resource")
	}
	return &JSONResource{
		subRes: res,
	}
}

// Load will load the underlying Resource and parse the JSON rules into standard GRule syntax.
func (jr *JSONResource) Load() ([]byte, error) {
	data, err := jr.subRes.Load()
	if err != nil {
		return nil, err
	}
	firstRune := string(bytes.TrimSpace(data)[0])

	var rs string

	if firstRune == "[" {
		rs, err = ParseJSONRuleset(data)
	} else if firstRune == "{" {
		rs, err = ParseJSONRule(data)
	} else {
		err = errors.New("invalid JSON input")
	}
	if err != nil {
		return nil, err
	}
	return []byte(rs), nil
}

// String will state the resource source.
func (jr *JSONResource) String() string {
	return "JSON Resource, underlying resource: " + jr.subRes.String()
}

// NewJSONResourceBundleFromBundle innstantiates a new bundled JSON resource parser from an underlying ResourceBundle.
func NewJSONResourceBundleFromBundle(bundle ResouceBundle) ResouceBundle {
	if _, ok := bundle.(*JSONResourceBundle); ok {
		panic("cannot create JSON resource bundle from JSON resource bundle")
	}
	return &JSONResourceBundle{
		subRes: bundle,
	}
}

// Load will load the underlying ResourceBundle and parse the JSON rules into standard GRule syntax.
func (jrb *JSONResourceBundle) Load() ([]Resource, error) {
	ress, err := jrb.subRes.Load()
	if err != nil {
		return nil, err
	}
	nress := make([]Resource, len(ress))
	for i := 0; i < len(ress); i++ {
		nress[i] = NewJSONResourceFromResource(ress[i])
	}
	return nress, nil
}

// MustLoad operates the same as load except it will panic in the event of an error.
func (jrb *JSONResourceBundle) MustLoad() []Resource {
	ress := jrb.subRes.MustLoad()
	nress := make([]Resource, len(ress))
	for i := 0; i < len(ress); i++ {
		nress[i] = NewJSONResourceFromResource(ress[i])
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
		sb.WriteString(parseRule(&rules[i]))
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

	rs = parseRule(&rule)
	return
}

// ParseRule Accepts a struct of GruleJSON rule and returns the parsed string of GRule.
func ParseRule(rule *GruleJSON) (r string, err error) {
	defer func() {
		if x := recover(); x != nil {
			err = fmt.Errorf("%v", x)
		}
	}()
	r = parseRule(rule)
	return
}

func parseRule(rule *GruleJSON) string {
	if len(rule.Name) == 0 {
		panic("rule name cannot be blank")
	}
	if rule.When == nil {
		panic("rule when condition cannot be nil")
	}
	if rule.Then == nil {
		panic("rule thenn condition cannot be nil")
	}
	var sb strings.Builder
	sb.WriteString("rule ")
	sb.WriteString(rule.Name)
	sb.WriteString(" ")
	sb.WriteString(strconv.Quote(rule.Description))
	sb.WriteString(" salience ")
	sb.WriteString(strconv.Itoa(rule.Salience))
	sb.WriteString(" {\n    when\n        ")
	sb.WriteString(parseWhen(rule.When))
	sb.WriteString("\n    then\n")
	thens := parseThen(rule.Then)
	for i := 0; i < len(thens); i++ {
		sb.WriteString("        ")
		sb.WriteString(thens[i])
		sb.WriteString("\n")
	}
	sb.WriteString("}\n")
	return sb.String()
}

func parseThen(ts []interface{}) []string {
	thens := make([]string, len(ts))
	for i := 0; i < len(ts); i++ {
		switch x := ts[i].(type) {
		case string:
			thens[i] = x
			if !strings.HasSuffix(thens[i], ";") {
				thens[i] += ";"
			}
		case map[string]interface{}:
			thens[i] = buildExpression(x, 0) + ";"
		default:
			panic("invalid then type, must be a string or an array of action objects")
		}
	}
	return thens
}

func parseWhen(w interface{}) string {
	switch x := w.(type) {
	case string:
		return x
	case map[string]interface{}:
		return buildExpression(x, 0)
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
	for k, v := range input {
		switch k {
		case "and":
			return buildCompoundOperator(v, depth, " && ")
		case "or":
			return buildCompoundOperator(v, depth, " || ")
		case "eq":
			return joinOperator(v, " == "), false
		case "not":
			return joinOperator(v, " != "), false
		case "gt":
			return joinOperator(v, " > "), false
		case "gte":
			return joinOperator(v, " >= "), false
		case "lt":
			return joinOperator(v, " < "), false
		case "lte":
			return joinOperator(v, " <= "), false
		case "bor":
			return joinOperator(v, " | "), false
		case "band":
			return joinOperator(v, " & "), false
		case "plus":
			return joinOperator(v, " + "), false
		case "minus":
			return joinOperator(v, " - "), false
		case "div":
			return joinOperator(v, " / "), false
		case "mul":
			return joinOperator(v, " * "), false
		case "mod":
			return joinOperator(v, " % "), false
		case "set":
			return joinSet(v, " = "), true
		case "call":
			return joinCall(v), true
		case "obj":
			if s, ok := v.(string); ok {
				return s, true
			}
			panic("object must be a string")
		case "const":
			switch x := v.(type) {
			case string:
				return strconv.Quote(x), true
			case float64:
				return strconv.FormatFloat(x, 'f', -1, 64), true
			case bool:
				if x {
					return "true", true
				}
				return "false", true
			}
			panic("constant must be a string or a numeric value")
		default:
			panic("unknown operator type: " + k)
		}
	}
	panic("boolean expression cannot be empty")
}

func buildCompoundOperator(o interface{}, depth int, op string) (string, bool) {
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
			return "(" + strings.Join(ands, op) + ")", false
		}
		return strings.Join(ands, op), false
	}
	panic("compound operator must be an array")
}

func joinCall(v interface{}) string {
	if arr, ok := v.([]interface{}); ok {
		if len(arr) == 0 {
			panic("call operator must have at least one operand")
		}
		var f string
		var ok bool
		if f, ok = arr[0].(string); !ok {
			panic("first call operand must be a string")
		}
		if len(arr) > 1 {
			sars := make([]string, len(arr)-1)
			for i := 1; i < len(arr); i++ {
				sars[i-1] = parseCallOperand(arr[i])
			}
			return f + "(" + strings.Join(sars, ", ") + ")"
		}
		return f + "()"
	}
	panic("operator has an unexpected type")
}

func parseCallOperand(o interface{}) string {
	switch x := o.(type) {
	case string:
		if len(x) == 0 {
			panic("operand cannnot be empty")
		}
		return x
	case float64:
		return fmt.Sprint(x)
	case bool:
		if x {
			return "true"
		}
		return "false"
	case map[string]interface{}:
		return buildExpression(x, 0)
	default:
		panic("operand has an invalid type")
	}
}

func joinOperator(v interface{}, op string) string {
	if arr, ok := v.([]interface{}); ok {
		if len(arr) == 0 {
			panic("operator cannot have 0 operands")
		}
		ops := make([]string, len(arr))
		for i := 0; i < len(arr); i++ {
			ops[i] = parseOperand(arr[i], false)
		}
		return strings.Join(ops, op)
	}
	panic("operator has an unexpected type")
}

func joinSet(v interface{}, op string) string {
	if arr, ok := v.([]interface{}); ok {
		if len(arr) != 2 {
			panic("set operand count must be 2")
		}
		return parseOperand(arr[0], true) + op + parseOperand(arr[1], true)
	}
	panic("operator has an unexpected type")
}

func parseOperand(o interface{}, noWrap bool) string {
	switch x := o.(type) {
	case string:
		return x
	case float64:
		return fmt.Sprint(x)
	case bool:
		if x {
			return "true"
		}
		return "false"
	case map[string]interface{}:
		expr, expNoWrap := buildExpressionEx(x, 0)
		if expNoWrap || noWrap {
			return expr
		}
		return "(" + expr + ")"
	default:
		panic("operand has an invalid type")
	}
}
