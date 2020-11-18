package ast

// IntegerLiteral will hold IntegerLiteral constant AST data
type IntegerLiteral struct {
	Integer int64
}

// StringLiteral will hold StringLiteral constant AST data
type StringLiteral struct {
	String string
}

// FloatLiteral will hold FloatLiteral constant AST data
type FloatLiteral struct {
	Float float64
}

// BooleanLiteral will hold BooleanLiteral constant AST data
type BooleanLiteral struct {
	Boolean bool
}

// IntegerLiteralReceiver should be implemented by AST graph node to receive a IntegerLiteral AST graph node
type IntegerLiteralReceiver interface {
	AcceptIntegerLiteral(fun *IntegerLiteral)
}

// StringLiteralReceiver should be implemented by AST graph node to receive a StringLiteral AST graph node
type StringLiteralReceiver interface {
	AcceptStringLiteral(fun *StringLiteral)
}

// FloatLiteralReceiver should be implemented by AST graph node to receive a FloatLiteral AST graph node
type FloatLiteralReceiver interface {
	AcceptFloatLiteral(fun *FloatLiteral)
}

// BooleanLiteralReceiver should be implemented by AST graph node to receive a BooleanLiteral AST graph node
type BooleanLiteralReceiver interface {
	AcceptBooleanLiteral(fun *BooleanLiteral)
}
