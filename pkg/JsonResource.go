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

		return nil, fmt.Errorf("cannot create JSON resource from JSON resource")
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
		return nil, fmt.Errorf("cannot create JSON resource bundle from JSON resource bundle")
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
	var err error
	for i := 0; i < len(ress); i++ {
		nress[i], err = NewJSONResourceFromResource(ress[i])
		if err != nil {
			panic(err.Error())
		}
	}

	return nress
}

// ParseJSONRuleset accepts a byte array containing an array of rules in JSON format to be parsed into GRule syntax.
func ParseJSONRuleset(data []byte) (rs string, err error) {
	var rules []GruleJSON
	err = json.Unmarshal(data, &rules)
	if err != nil {

		return
	}
	var sb strings.Builder
	for i := 0; i < len(rules); i++ {
		rulestr, err := parseRule(&rules[i])
		if err != nil {
			return "", err
		}
		sb.WriteString(rulestr)
	}
	rs = sb.String()

	return
}

// ParseJSONRule accepts a byte array containing an rule in JSON format to be parsed into GRule syntax.
func ParseJSONRule(data []byte) (rs string, err error) {
	var rule GruleJSON
	err = json.Unmarshal(data, &rule)
	if err != nil {

		return
	}

	rs, err = parseRule(&rule)

	return
}

// ParseRule Accepts a struct of GruleJSON rule and returns the parsed string of GRule.
func ParseRule(rule *GruleJSON) (r string, err error) {
	r, err = parseRule(rule)

	return
}

func parseRule(rule *GruleJSON) (string, error) {
	if len(rule.Name) == 0 {

		return "", fmt.Errorf("rule name cannot be blank")
	}
	if rule.When == nil {

		return "", fmt.Errorf("rule when condition cannot be nil")
	}
	if rule.Then == nil {

		return "", fmt.Errorf("rule thenn condition cannot be nil")
	}
	var stringBuilder strings.Builder
	stringBuilder.WriteString("rule ")
	stringBuilder.WriteString(rule.Name)
	stringBuilder.WriteString(" ")
	stringBuilder.WriteString(strconv.Quote(rule.Description))
	stringBuilder.WriteString(" salience ")
	stringBuilder.WriteString(strconv.Itoa(rule.Salience))
	stringBuilder.WriteString(" {\n    when\n        ")
	when, err := parseWhen(rule.When)
	if err != nil {

		return "", err
	}
	stringBuilder.WriteString(when)
	stringBuilder.WriteString("\n    then\n")
	thens, err := parseThen(rule.Then)
	if err != nil {
		return "", err
	}
	for i := 0; i < len(thens); i++ {
		stringBuilder.WriteString("        ")
		stringBuilder.WriteString(thens[i])
		stringBuilder.WriteString("\n")
	}
	stringBuilder.WriteString("}\n")

	return stringBuilder.String(), nil
}

func parseThen(ts []interface{}) ([]string, error) {
	thens := make([]string, len(ts))
	for thenItem := 0; thenItem < len(ts); thenItem++ {
		switch thenType := ts[thenItem].(type) {
		case string:
			thens[thenItem] = thenType
			if !strings.HasSuffix(thens[thenItem], ";") {
				thens[thenItem] += ";"
			}
		case map[string]interface{}:
			bldExpr, err := buildExpression(thenType, 0)
			if err != nil {
				return nil, err
			}
			thens[thenItem] = bldExpr + ";"
		default:

			return nil, fmt.Errorf("invalid then type, must be a string or an array of action objects")
		}
	}

	return thens, nil
}

func parseWhen(w interface{}) (string, error) {
	switch whenType := w.(type) {
	case string:

		return whenType, nil
	case map[string]interface{}:

		return buildExpression(whenType, 0)
	default:

		return "", fmt.Errorf("invalid when type, must be a string or an array of condition objects")
	}
}

func buildExpression(input map[string]interface{}, depth int) (string, error) {
	exp, _, err := buildExpressionEx(input, depth)

	return exp, err
}

func buildExpressionEx(input map[string]interface{}, depth int) (string, bool, error) {
	if depth > 1024 {

		return "", false, fmt.Errorf("JSON nesting exceeded 1024 levels, aborting")
	}
	if len(input) > 1 {

		return "", false, fmt.Errorf("expression objects can only contain a single operation type")
	}
	for key, value := range input {
		switch key {
		case "and":

			return buildCompoundOperator(value, depth, " && ")
		case "or":

			return buildCompoundOperator(value, depth, " || ")
		case "eq":
			opers, err := joinOperator(value, " == ")

			return opers, false, err
		case "not":
			opers, err := joinOperator(value, " != ")

			return opers, false, err
		case "gt":
			opers, err := joinOperator(value, " > ")

			return opers, false, err
		case "gte":
			opers, err := joinOperator(value, " >= ")

			return opers, false, err
		case "lt":
			opers, err := joinOperator(value, " < ")

			return opers, false, err
		case "lte":
			opers, err := joinOperator(value, " <= ")

			return opers, false, err
		case "bor":
			opers, err := joinOperator(value, " | ")

			return opers, false, err
		case "band":
			opers, err := joinOperator(value, " & ")

			return opers, false, err
		case "plus":
			opers, err := joinOperator(value, " + ")

			return opers, false, err
		case "minus":
			opers, err := joinOperator(value, " - ")

			return opers, false, err
		case "div":
			opers, err := joinOperator(value, " / ")

			return opers, false, err
		case "mul":
			opers, err := joinOperator(value, " * ")

			return opers, false, err
		case "mod":
			opers, err := joinOperator(value, " % ")

			return opers, false, err
		case "set":
			opers, err := joinSet(value, " = ")

			return opers, true, err
		case "call":
			joinStr, err := joinCall(value)

			return joinStr, true, err
		case "obj":
			if s, ok := value.(string); ok {

				return s, true, nil
			}

			return "", false, fmt.Errorf("object must be a string")
		case "const":
			switch valueType := value.(type) {
			case string:

				return strconv.Quote(valueType), true, nil
			case float64:

				return strconv.FormatFloat(valueType, 'f', -1, 64), true, nil
			case bool:
				if valueType {

					return "true", true, nil
				}

				return "false", true, nil
			}

			return "", false, fmt.Errorf("constant must be a string or a numeric value")
		default:

			return "", false, fmt.Errorf("unknown operator type: %s", key)
		}
	}

	return "", false, fmt.Errorf("boolean expression cannot be empty")
}

func buildCompoundOperator(o interface{}, depth int, operator string) (string, bool, error) {
	if andarr, ok := o.([]interface{}); ok {
		var ands []string
		if len(andarr) < 2 {

			return "", false, fmt.Errorf("and operator must have at least 2 operands")
		}
		for i := 0; i < len(andarr); i++ {
			if subVal, ok := andarr[i].(map[string]interface{}); ok {
				bldexpr, err := buildExpression(subVal, depth+1)
				if err != nil {

					return "", false, err
				}

				ands = append(ands, bldexpr)
			} else {

				return "", false, fmt.Errorf("and operands must be an array of objects")
			}
		}
		if depth > 0 {

			return "(" + strings.Join(ands, operator) + ")", false, nil
		}

		return strings.Join(ands, operator), false, nil
	}

	return "", false, fmt.Errorf("compound operator must be an array")
}

func joinCall(v interface{}) (string, error) {
	if arr, ok := v.([]interface{}); ok {
		if len(arr) == 0 {

			return "", fmt.Errorf("call operator must have at least one operand")
		}
		var firstCallOperand string
		var ok bool
		if firstCallOperand, ok = arr[0].(string); !ok {

			return "", fmt.Errorf("first call operand must be a string")
		}
		if len(arr) > 1 {
			sars := make([]string, len(arr)-1)
			for i := 1; i < len(arr); i++ {
				operandStr, err := parseCallOperand(arr[i])
				if err != nil {
					return "", err
				}
				sars[i-1] = operandStr
			}

			return firstCallOperand + "(" + strings.Join(sars, ", ") + ")", nil
		}

		return firstCallOperand + "()", nil
	}

	return "", fmt.Errorf("operator has an unexpected type")
}

func parseCallOperand(o interface{}) (string, error) {
	switch operandType := o.(type) {
	case string:
		if len(operandType) == 0 {
			return "", fmt.Errorf("operand cannnot be empty")
		}

		return operandType, nil
	case float64:

		return fmt.Sprint(operandType), nil
	case bool:
		if operandType {

			return "true", nil
		}

		return "false", nil
	case map[string]interface{}:

		return buildExpression(operandType, 0)
	default:

		return "", fmt.Errorf("operand has an invalid type")
	}
}

func joinOperator(v interface{}, operator string) (string, error) {
	if arr, ok := v.([]interface{}); ok {
		if len(arr) == 0 {

			return "", fmt.Errorf("operator cannot have 0 operands")
		}
		ops := make([]string, len(arr))
		for i := 0; i < len(arr); i++ {
			ope, err := parseOperand(arr[i], false)
			if err != nil {

				return "", err
			}
			ops[i] = ope
		}

		return strings.Join(ops, operator), nil
	}

	return "", fmt.Errorf("operator has an unexpected type")
}

func joinSet(v interface{}, operator string) (string, error) {
	if arr, ok := v.([]interface{}); ok {
		if len(arr) != 2 {

			return "", fmt.Errorf("set operand count must be 2")
		}
		leftOpe, err := parseOperand(arr[0], true)
		if err != nil {

			return "", err
		}
		rightOpe, err := parseOperand(arr[1], true)
		if err != nil {

			return "", err
		}

		return leftOpe + operator + rightOpe, nil
	}

	return "", fmt.Errorf("operator has an unexpected type")
}

func parseOperand(o interface{}, noWrap bool) (string, error) {
	switch operandType := o.(type) {
	case string:

		return operandType, nil
	case float64:

		return fmt.Sprint(operandType), nil
	case bool:

		if operandType {

			return "true", nil
		}

		return "false", nil
	case map[string]interface{}:
		expr, expNoWrap, err := buildExpressionEx(operandType, 0)

		if err != nil {

			return expr, err
		}
		if expNoWrap || noWrap {

			return expr, nil
		}

		return "(" + expr + ")", nil
	default:

		return "", fmt.Errorf("operand has an invalid type")
	}
}
