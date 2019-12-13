package model

import "reflect"

// Evaluator define the interface for all element that able to evaluate rule element against underlying datacontext
// within the rule engine execution.
type Evaluator interface {
	Evaluate() (reflect.Value, error)
}
