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

package antlr

type (
	stack struct {
		top    *node
		length int
	}
	node struct {
		value interface{}
		prev  *node
	}
)

// Create a new stack
func newStack() *stack {
	return &stack{nil, 0}
}

// Len return the number of items in the stack
func (s *stack) Len() int {
	return s.length
}

// Peek views the top item on the stack
func (s *stack) Peek() interface{} {
	if s.length == 0 {
		return nil
	}
	return s.top.value
}

// Pop the top item of the stack and return it
func (s *stack) Pop() interface{} {
	if s.length == 0 {
		return nil
	}

	n := s.top
	s.top = n.prev
	s.length--
	return n.value
}

// Push a value onto the top of the stack
func (s *stack) Push(value interface{}) {
	n := &node{value, s.top}
	s.top = n
	s.length++
}
