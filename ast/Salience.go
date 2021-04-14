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

// NewSalience create new Salience AST object
func NewSalience(val int) *Salience {
	return &Salience{
		SalienceValue: val,
	}
}

// Salience is a simple AST object that stores salience
type Salience struct {
	SalienceValue int
}

// SalienceReceiver must be implemented by any AST object that stores salience
type SalienceReceiver interface {
	AcceptSalience(salience *Salience) error
}

// AcceptIntegerLiteral accept the assigned integer
func (sal *Salience) AcceptIntegerLiteral(lit *IntegerLiteral) {
	sal.SalienceValue = int(lit.Integer)
}
