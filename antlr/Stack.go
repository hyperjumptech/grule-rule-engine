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
