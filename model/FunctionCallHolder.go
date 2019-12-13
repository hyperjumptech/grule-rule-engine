package model

// FunctionCallHolder defines a graph that should be able to store function call.
type FunctionCallHolder interface {
	AcceptFunctionCall(funcCall *FunctionCall) error
}
