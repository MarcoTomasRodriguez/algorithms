package structures

import (
	"testing"
)

func TestQueue(t *testing.T) {
	t.Run("EnQueue", func(t *testing.T) {
		var queue Queue = Queue{0}
		queue.EnQueue(1)
		if !(len(queue) == 2 && queue[1] == 1) {
			t.Error("EnQueue should put on queue the element")
		}
	})
	t.Run("DeQueue", func(t *testing.T) {
		var queue Queue = Queue{0, 1}
		queue.DeQueue()
		if !(len(queue) == 1 && queue[0] == 1) {
			t.Error("DeQueue should dequeue the element")
		}
	})
}

func TestStack(t *testing.T) {
	t.Run("Push", func(t *testing.T) {
		var stack Stack = Stack{0}
		stack.Push(1)
		if !(len(stack) == 2 && stack[1] == 1) {
			t.Error("Push should stack the element")
		}
	})
	t.Run("Pop", func(t *testing.T) {
		var stack Stack = Stack{0, 1}
		stack.Pop()
		if !(len(stack) == 1 && stack[0] == 0) {
			t.Error("Pop should destack the element")
		}
	})
}

func TestList(t *testing.T) {
	t.Run("IndexOutOfRange", func(t *testing.T) {
		var list List = List{0, 0, 0, 0, 0}
		var truthyTestCases []int = []int{-50, -10, -1, 50, 70}
		var falsyTestCases []int = []int{0, 1, 2, 3, 4}
		for _, testCase := range truthyTestCases {
			if !list.IndexOutOfRange(testCase) {
				t.Error("Index should be out of range")
			}
		}
		for _, testCase := range falsyTestCases {
			if list.IndexOutOfRange(testCase) {
				t.Error("Index should not be out of range")
			}
		}
	})
	t.Run("Get", func(t *testing.T) {
		var list List = List{0, 0, 0, 0, 0}
		var truthyTestCases []int = []int{0, 1, 2, 3, 4}
		var falsyTestCases []int = []int{-50, -10, -1, 50, 70}
		for _, testCase := range truthyTestCases {
			if list.Get(testCase) == nil {
				t.Error("Should get the item on place")
			}
		}
		for _, testCase := range falsyTestCases {
			if list.Get(testCase) != nil {
				t.Error("Should not get the item")
			}
		}
	})
	t.Run("Set", func(t *testing.T) {
		var list List = List{0, 0, 0, 0}
		list.Set(1, 10)
		if list.Get(10) != nil {
			t.Error("Should not put the item on the index")
		}
		list.Set(1, 0)
		if list.Get(0) == nil {
			t.Error("Should put the item on the index")
		}
	})
	t.Run("Swap", func(t *testing.T) {
		var list List = List{0, 1, 2, 3}
		list.Swap(0, 20)
		if !(list.Get(0) == 0) {
			t.Error("Should not swap the elements")
		}
		list.Swap(0, 1)
		if !(list.Get(0) == 1) {
			t.Error("Should swap the elements")
		}
	})
	t.Run("Push", func(t *testing.T) {
		var list List = List{}
		list.Push(1)
		if !(len(list) == 1 && list[0] == 1) {
			t.Error("Should push an item")
		}
		list.Push(2)
		if !(len(list) == 2 && list[1] == 2) {
			t.Error("Should push an item")
		}
	})
	t.Run("Pop", func(t *testing.T) {
		var list List = List{1, 2}
		list.Pop()
		if !(len(list) == 1 && list[0] == 1) {
			t.Error("Should pop an item")
		}
		list.Pop()
		if !(len(list) == 0) {
			t.Error("Should pop an item")
		}
	})
}
