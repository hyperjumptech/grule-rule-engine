package model

// VariableHolder should be implemented by any object graph that would hold a variable name.
type VariableHolder interface {
	AcceptVariable(name string) error
}
