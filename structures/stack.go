package structures

// Stack implementation
type Stack []interface{}

// Push puts an element to the end of the stack
func (stack *Stack) Push(element interface{}) {
	*stack = append(*stack, element)
}

// Pop removes the last element of the stack
func (stack *Stack) Pop(element interface{}) {
	if len(*stack) > 0 {
		*stack = (*stack)[:len(*stack)-1]
	}
}
