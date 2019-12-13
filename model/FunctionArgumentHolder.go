package model

// FunctionArgumentHolder define a graph that should store function argument.
type FunctionArgumentHolder interface {
	AcceptFunctionArgument(funcArg *FunctionArgument) error
}
