package ast

import (
	"github.com/hyperjumptech/grule-rule-engine/pkg"
	"github.com/sirupsen/logrus"
	"reflect"
	"strings"
	"time"
)

var (
	// GrlLogger is the logger that be used from within the rule engine GRL
	GrlLogger = logrus.WithFields(logrus.Fields{
		"lib":     "grule",
		"package": "AST",
		"source":  "GRL",
	})
)

// BuildInFunctions strucr hosts the built-in functions ready to invoke from the rule engine execution.
type BuildInFunctions struct {
	Knowledge     *KnowledgeBase
	WorkingMemory *WorkingMemory
}

// MakeTime will create a Time struct according to the argument values.
func (gf *BuildInFunctions) MakeTime(year, month, day, hour, minute, second int64) time.Time {
	return time.Date(int(year), time.Month(month), int(day), int(hour), int(minute), int(second), 0, time.Local)
}

// Changed will enable Grule's working memory to forget about a variable, so in the next cycle
// grue will re-valuate that variable instead of just use the value from its working memory.
// If you change the variable from within grule DRL (using assignment expression, you dont need to call this
// function on that variable since grule will automaticaly see the change. So only call this
// function if the variable got changed from your internal struct logic.
func (gf *BuildInFunctions) Changed(variableName string) {
	gf.WorkingMemory.Reset(variableName)
	gf.Knowledge.DataContext.VariableChangeCount++
}

// Now is an extension tn time.Now().
func (gf *BuildInFunctions) Now() time.Time {
	return time.Now()
}

// Log extension to log.Print
func (gf *BuildInFunctions) Log(text string) {
	GrlLogger.Println(text)
}

// StringContains extension to strings.Contains
func (gf *BuildInFunctions) StringContains(str, substr string) bool {
	return strings.Contains(str, substr)
}

// LogFormat extension to log.Printf
func (gf *BuildInFunctions) LogFormat(format string, i interface{}) {
	GrlLogger.Printf(format, i)
}

// IsNil Enables nill checking on variables.
func (gf *BuildInFunctions) IsNil(i interface{}) bool {
	val := reflect.ValueOf(i)
	if val.Kind() == reflect.Struct {
		return false
	}
	return !val.IsValid() || val.IsNil()
}

// IsZero Enable zero checking
func (gf *BuildInFunctions) IsZero(i interface{}) bool {
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
func (gf *BuildInFunctions) Retract(ruleName string) {
	gf.Knowledge.RetractRule(ruleName)
}

// GetTimeYear will get the year value of time
func (gf *BuildInFunctions) GetTimeYear(time time.Time) int {
	return time.Year()
}

// GetTimeMonth will get the month value of time
func (gf *BuildInFunctions) GetTimeMonth(time time.Time) int {
	return int(time.Month())
}

// GetTimeDay will get the day value of time
func (gf *BuildInFunctions) GetTimeDay(time time.Time) int {
	return time.Day()
}

// GetTimeHour will get the hour value of time
func (gf *BuildInFunctions) GetTimeHour(time time.Time) int {
	return time.Hour()
}

// GetTimeMinute will get the minute value of time
func (gf *BuildInFunctions) GetTimeMinute(time time.Time) int {
	return time.Minute()
}

// GetTimeSecond will get the second value of time
func (gf *BuildInFunctions) GetTimeSecond(time time.Time) int {
	return time.Second()
}

// IsTimeBefore will check if the 1st argument is before the 2nd argument.
func (gf *BuildInFunctions) IsTimeBefore(time, before time.Time) bool {
	return time.Before(before)
}

// IsTimeAfter will check if the 1st argument is after the 2nd argument.
func (gf *BuildInFunctions) IsTimeAfter(time, after time.Time) bool {
	return time.After(after)
}

// TimeFormat will format a time according to format layout.
func (gf *BuildInFunctions) TimeFormat(time time.Time, layout string) string {
	return time.Format(layout)
}
