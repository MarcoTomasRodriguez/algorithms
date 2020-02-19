package structures

import (
	"testing"
)

const (
	testVar1 = "Test1"
	testVar2 = "Test2"
)

func TestQueue(t *testing.T) {
	t.Run("EnQueue", func(t *testing.T) {
		var queue Queue = Queue{testVar1}
		queue.EnQueue(testVar2)
		if !(len(queue) == 2 && queue[1] == testVar2) {
			t.Error("EnQueue should put on queue the element")
		}
	})
	t.Run("DeQueue", func(t *testing.T) {
		var queue Queue = Queue{testVar1, testVar2}
		queue.DeQueue()
		if !(len(queue) == 1 && queue[0] == testVar2) {
			t.Error("DeQueue should dequeue the element")
		}
	})
}

func TestStack(t *testing.T) {
	t.Run("Push", func(t *testing.T) {
		var stack Stack = Stack{testVar1}
		stack.Push(testVar2)
		if !(len(stack) == 2 && stack[1] == testVar2) {
			t.Error("Push should stack the element")
		}
	})
	t.Run("Pop", func(t *testing.T) {
		var stack Stack = Stack{testVar1, testVar2}
		stack.Pop()
		if !(len(stack) == 1 && stack[0] == testVar1) {
			t.Error("Pop should destack the element")
		}
	})
}

func TestList(t *testing.T) {
	// Implement
}
