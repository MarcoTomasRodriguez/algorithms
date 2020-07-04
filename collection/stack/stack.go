package stack

type (
	Stack struct {
		head *node
		length int
	}
	node struct {
		value interface{}
		prev *node
	}
)

func New() *Stack {
	return &Stack{nil, 0}
}

func (stack *Stack) Len() int {
	return stack.length
}

func (stack *Stack) Peek() interface{} {
	if stack.length >= 1 {
		return stack.head.value
	}
	return nil
}

func (stack *Stack) Push(value interface{})  {
	stack.head = &node{value, stack.head}
	stack.length++
}

func (stack *Stack) Pop() interface{} {
	if stack.length == 0 {
		return nil
	}
	if stack.length == 1 {
		oldHead := stack.head
		stack.head = nil
		stack.length--
		return oldHead.value
	}
	newHead := stack.head.prev
	stack.head = newHead
	stack.length--
	return newHead.value
}
