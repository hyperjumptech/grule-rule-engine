package v3

type IntegerLiteral struct {
	Integer int64
}

type StringLiteral struct {
	String string
}

type FloatLiteral struct {
	Float float64
}

type BooleanLiteral struct {
	Boolean bool
}

// FunctionCallReceiver should be implemented bu AST graph node to receive a FunctionCall AST graph mode
type IntegerLiteralReceiver interface {
	AcceptIntegerLiteral(fun *IntegerLiteral)
}

// FunctionCallReceiver should be implemented bu AST graph node to receive a FunctionCall AST graph mode
type StringLiteralReceiver interface {
	AcceptStringLiteral(fun *StringLiteral)
}

// FunctionCallReceiver should be implemented bu AST graph node to receive a FunctionCall AST graph mode
type FloatLiteralReceiver interface {
	AcceptFloatLiteral(fun *FloatLiteral)
}

// FunctionCallReceiver should be implemented bu AST graph node to receive a FunctionCall AST graph mode
type BooleanLiteralReceiver interface {
	AcceptBooleanLiteral(fun *BooleanLiteral)
}
