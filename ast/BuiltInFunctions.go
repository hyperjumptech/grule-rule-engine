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

import (
	"github.com/hyperjumptech/grule-rule-engine/logger"
	"math"
	"reflect"
	"strings"
	"time"

	"github.com/hyperjumptech/grule-rule-engine/pkg"
	"github.com/sirupsen/logrus"
)

var (
	// GrlLogger is the logger that be used from within the rule engine GRL
	GrlLogger = logger.Log.WithFields(logrus.Fields{
		"package": "AST",
		"source":  "GRL",
	})
)

// BuiltInFunctions struct hosts the built-in functions ready to invoke from the rule engine execution.
type BuiltInFunctions struct {
	Knowledge     *KnowledgeBase
	WorkingMemory *WorkingMemory
	DataContext   IDataContext
}

// Complete will cause the engine to stop processing further rules in the current cycle.
func (gf *BuiltInFunctions) Complete() {
	gf.DataContext.Complete()
}

// MakeTime will create a Time struct according to the argument values.
func (gf *BuiltInFunctions) MakeTime(year, month, day, hour, minute, second int64) time.Time {
	return time.Date(int(year), time.Month(month), int(day), int(hour), int(minute), int(second), 0, time.Local)
}

// Changed is another name for Forget function. This function is retained for backward compatibility reason and will be removed in the future.
func (gf *BuiltInFunctions) Changed(variableName string) {
	gf.WorkingMemory.Reset(variableName)
	gf.Knowledge.DataContext.IncrementVariableChangeCount()
}

// Forget will force Grule's working memory to forget about a variable, or function call, so in the next cycle
// grue will re-valuate that variable/function instead of just use the value from its working memory.
// If you change the variable from within grule GRL (using assignment expression, you dont need to call this
// function on that variable since grule will automaticaly see the change. So only call this
// function if the variable got changed from your internal struct logic.
func (gf *BuiltInFunctions) Forget(snippet string) {
	gf.WorkingMemory.Reset(snippet)
	gf.Knowledge.DataContext.IncrementVariableChangeCount()
}

// Now is an extension tn time.Now().
func (gf *BuiltInFunctions) Now() time.Time {
	return time.Now()
}

// Log extension to log.Print
func (gf *BuiltInFunctions) Log(text string) {
	GrlLogger.Println(text)
}

// StringContains extension to strings.Contains
func (gf *BuiltInFunctions) StringContains(str, substr string) bool {
	return strings.Contains(str, substr)
}

// LogFormat extension to log.Printf
func (gf *BuiltInFunctions) LogFormat(format string, i interface{}) {
	GrlLogger.Printf(format, i)
}

// IsNil Enables nill checking on variables.
func (gf *BuiltInFunctions) IsNil(i interface{}) bool {
	val := reflect.ValueOf(i)
	if val.Kind() == reflect.Struct {
		return false
	}
	return !val.IsValid() || val.IsNil()
}

// IsZero Enable zero checking
func (gf *BuiltInFunctions) IsZero(i interface{}) bool {
	val := reflect.ValueOf(i)
	switch val.Kind() {
	case reflect.Struct:
		if val.Type().String() == "time.Time" {
			return i.(time.Time).IsZero()
		}
		return false
	case reflect.Ptr:
		return val.IsNil()
	default:
		switch pkg.GetBaseKind(val) {
		case reflect.String:
			return len(val.String()) == 0
		case reflect.Int64:
			return val.Int() == 0
		case reflect.Uint64:
			return val.Uint() == 0
		case reflect.Float64:
			return val.Float() == 0
		default:
			return false
		}
	}
}

// Retract will retract a rule from next evaluation cycle.
func (gf *BuiltInFunctions) Retract(ruleName string) {
	gf.Knowledge.RetractRule(ruleName)
}

// GetTimeYear will get the year value of time
func (gf *BuiltInFunctions) GetTimeYear(time time.Time) int {
	return time.Year()
}

// GetTimeMonth will get the month value of time
func (gf *BuiltInFunctions) GetTimeMonth(time time.Time) int {
	return int(time.Month())
}

// GetTimeDay will get the day value of time
func (gf *BuiltInFunctions) GetTimeDay(time time.Time) int {
	return time.Day()
}

// GetTimeHour will get the hour value of time
func (gf *BuiltInFunctions) GetTimeHour(time time.Time) int {
	return time.Hour()
}

// GetTimeMinute will get the minute value of time
func (gf *BuiltInFunctions) GetTimeMinute(time time.Time) int {
	return time.Minute()
}

// GetTimeSecond will get the second value of time
func (gf *BuiltInFunctions) GetTimeSecond(time time.Time) int {
	return time.Second()
}

// IsTimeBefore will check if the 1st argument is before the 2nd argument.
func (gf *BuiltInFunctions) IsTimeBefore(time, before time.Time) bool {
	return time.Before(before)
}

// IsTimeAfter will check if the 1st argument is after the 2nd argument.
func (gf *BuiltInFunctions) IsTimeAfter(time, after time.Time) bool {
	return time.After(after)
}

// TimeFormat will format a time according to format layout.
func (gf *BuiltInFunctions) TimeFormat(time time.Time, layout string) string {
	return time.Format(layout)
}

// Max will pick the biggest of value in the arguments
func (gf *BuiltInFunctions) Max(vals ...float64) float64 {
	val := float64(0)
	for i, v := range vals {
		if i == 0 {
			val = v
		} else {
			if v > val {
				val = v
			}
		}
	}
	return val
}

// Min will pick the smallest of value in the arguments
func (gf *BuiltInFunctions) Min(vals ...float64) float64 {
	val := float64(0)
	for i, v := range vals {
		if i == 0 {
			val = v
		} else {
			if v < val {
				val = v
			}
		}
	}
	return val
}

// Abs is a wrapper function for math.Abs function
func (gf *BuiltInFunctions) Abs(x float64) float64 {
	return math.Abs(x)
}

// Acos is a wrapper function for math.Acos function
func (gf *BuiltInFunctions) Acos(x float64) float64 {
	return math.Acos(x)
}

// Acosh is a wrapper function for math.Acosh function
func (gf *BuiltInFunctions) Acosh(x float64) float64 {
	return math.Acosh(x)
}

// Asin is a wrapper function for math.Asin function
func (gf *BuiltInFunctions) Asin(x float64) float64 {
	return math.Asin(x)
}

// Asinh is a wrapper function for math.Asinh function
func (gf *BuiltInFunctions) Asinh(x float64) float64 {
	return math.Asinh(x)
}

// Atan is a wrapper function for math.Atan function
func (gf *BuiltInFunctions) Atan(x float64) float64 {
	return math.Atan(x)
}

// Atan2 is a wrapper function for math.Atan2 function
func (gf *BuiltInFunctions) Atan2(y, x float64) float64 {
	return math.Atan2(y, x)
}

// Atanh is a wrapper function for math.Atanh function
func (gf *BuiltInFunctions) Atanh(x float64) float64 {
	return math.Atanh(x)
}

// Cbrt is a wrapper function for math.Cbrt function
func (gf *BuiltInFunctions) Cbrt(x float64) float64 {
	return math.Cbrt(x)
}

// Ceil is a wrapper function for math.Ceil function
func (gf *BuiltInFunctions) Ceil(x float64) float64 {
	return math.Ceil(x)
}

// Copysign is a wrapper function for math.Copysign function
func (gf *BuiltInFunctions) Copysign(x, y float64) float64 {
	return math.Copysign(x, y)
}

// Cos is a wrapper function for math.Cos function
func (gf *BuiltInFunctions) Cos(x float64) float64 {
	return math.Cos(x)
}

// Cosh is a wrapper function for math.Cosh function
func (gf *BuiltInFunctions) Cosh(x float64) float64 {
	return math.Cosh(x)
}

// Dim is a wrapper function for math.Dim function
func (gf *BuiltInFunctions) Dim(x, y float64) float64 {
	return math.Dim(x, y)
}

// Erf is a wrapper function for math.Erf function
func (gf *BuiltInFunctions) Erf(x float64) float64 {
	return math.Erf(x)
}

// Erfc is a wrapper function for math.Erfc function
func (gf *BuiltInFunctions) Erfc(x float64) float64 {
	return math.Erfc(x)
}

// Erfcinv is a wrapper function for math.Erfcinv function
func (gf *BuiltInFunctions) Erfcinv(x float64) float64 {
	return math.Erfcinv(x)
}

// Erfinv is a wrapper function for math.Erfinv function
func (gf *BuiltInFunctions) Erfinv(x float64) float64 {
	return math.Erfinv(x)
}

// Exp is a wrapper function for math.Exp function
func (gf *BuiltInFunctions) Exp(x float64) float64 {
	return math.Exp(x)
}

// Exp2 is a wrapper function for math.Exp2 function
func (gf *BuiltInFunctions) Exp2(x float64) float64 {
	return math.Exp2(x)
}

// Expm1 is a wrapper function for math.Expm1 function
func (gf *BuiltInFunctions) Expm1(x float64) float64 {
	return math.Expm1(x)
}

// Float64bits is a wrapper function for math.Float64bits function
func (gf *BuiltInFunctions) Float64bits(f float64) uint64 {
	return math.Float64bits(f)
}

// Float64frombits is a wrapper function for math.Float64frombits function
func (gf *BuiltInFunctions) Float64frombits(b uint64) float64 {
	return math.Float64frombits(b)
}

// Floor is a wrapper function for math.Floor function
func (gf *BuiltInFunctions) Floor(x float64) float64 {
	return math.Floor(x)
}

// Gamma is a wrapper function for math.Gamma function
func (gf *BuiltInFunctions) Gamma(x float64) float64 {
	return math.Gamma(x)
}

// Hypot is a wrapper function for math.Hypot function
func (gf *BuiltInFunctions) Hypot(p, q float64) float64 {
	return math.Hypot(p, q)
}

// Ilogb is a wrapper function for math.Ilogb function
func (gf *BuiltInFunctions) Ilogb(x float64) int {
	return math.Ilogb(x)
}

// IsInf is a wrapper function for math.IsInf function
func (gf *BuiltInFunctions) IsInf(f float64, sign int64) bool {
	return math.IsInf(f, int(sign))
}

// IsNaN is a wrapper function for math.IsNaN function
func (gf *BuiltInFunctions) IsNaN(f float64) (is bool) {
	return math.IsNaN(f)
}

// J0 is a wrapper function for math.J0 function
func (gf *BuiltInFunctions) J0(x float64) float64 {
	return math.J0(x)
}

// J1 is a wrapper function for math.J1 function
func (gf *BuiltInFunctions) J1(x float64) float64 {
	return math.J1(x)
}

// Jn is a wrapper function for math.Jn function
func (gf *BuiltInFunctions) Jn(n int64, x float64) float64 {
	return math.Jn(int(n), x)
}

// Ldexp is a wrapper function for math.Ldexp function
func (gf *BuiltInFunctions) Ldexp(frac float64, exp int64) float64 {
	return math.Ldexp(frac, int(exp))
}

// MathLog is a wrapper function for math.MathLog function
func (gf *BuiltInFunctions) MathLog(x float64) float64 {
	return math.Log(x)
}

// Log10 is a wrapper function for math.Log10 function
func (gf *BuiltInFunctions) Log10(x float64) float64 {
	return math.Log10(x)
}

// Log1p is a wrapper function for math.Log1p function
func (gf *BuiltInFunctions) Log1p(x float64) float64 {
	return math.Log1p(x)
}

// Log2 is a wrapper function for math.Log2 function
func (gf *BuiltInFunctions) Log2(x float64) float64 {
	return math.Log2(x)
}

// Logb is a wrapper function for math.Logb function
func (gf *BuiltInFunctions) Logb(x float64) float64 {
	return math.Logb(x)
}

// Mod is a wrapper function for math.Mod function
func (gf *BuiltInFunctions) Mod(x, y float64) float64 {
	return math.Mod(x, y)
}

// NaN is a wrapper function for math.NaN function
func (gf *BuiltInFunctions) NaN() float64 {
	return math.NaN()
}

// Pow is a wrapper function for math.Pow function
func (gf *BuiltInFunctions) Pow(x, y float64) float64 {
	return math.Pow(x, y)
}

// Pow10 is a wrapper function for math.Pow10 function
func (gf *BuiltInFunctions) Pow10(n int64) float64 {
	return math.Pow10(int(n))
}

// Remainder is a wrapper function for math.Remainder function
func (gf *BuiltInFunctions) Remainder(x, y float64) float64 {
	return math.Remainder(x, y)
}

// Round is a wrapper function for math.Round function
func (gf *BuiltInFunctions) Round(x float64) float64 {
	return math.Round(x)
}

// RoundToEven is a wrapper function for math.RoundToEven function
func (gf *BuiltInFunctions) RoundToEven(x float64) float64 {
	return math.RoundToEven(x)
}

// Signbit is a wrapper function for math.Signbit function
func (gf *BuiltInFunctions) Signbit(x float64) bool {
	return math.Signbit(x)
}

// Sin is a wrapper function for math.Sin function
func (gf *BuiltInFunctions) Sin(x float64) float64 {
	return math.Sin(x)
}

// Sinh is a wrapper function for math.Sinh function
func (gf *BuiltInFunctions) Sinh(x float64) float64 {
	return math.Sinh(x)
}

// Sqrt is a wrapper function for math.Sqrt function
func (gf *BuiltInFunctions) Sqrt(x float64) float64 {
	return math.Sqrt(x)
}

// Tan is a wrapper function for math.Tan function
func (gf *BuiltInFunctions) Tan(x float64) float64 {
	return math.Tan(x)
}

// Tanh is a wrapper function for math.Tanh function
func (gf *BuiltInFunctions) Tanh(x float64) float64 {
	return math.Tanh(x)
}

// Trunc is a wrapper function for math.Trunc function
func (gf *BuiltInFunctions) Trunc(x float64) float64 {
	return math.Trunc(x)
}
