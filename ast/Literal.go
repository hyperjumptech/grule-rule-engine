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
