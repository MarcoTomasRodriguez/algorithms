package stack

import (
	"testing"
)

func TestStack(t *testing.T) {
	q := New()

	if q.Len() != 0 {
		t.Errorf("Length should be 0")
	}

	q.Push(1)

	if q.Len() != 1 {
		t.Errorf("Length should be 1")
	}

	if q.Peek().(int) != 1 {
		t.Errorf("Pushed value should be 1")
	}

	if q.Pop().(int) != 1 {
		t.Errorf("Popped value should be 1")
	}

	if q.Peek() != nil || q.Pop() != nil {
		t.Errorf("Empty stack should have no values")
	}

	q.Push(1)
	q.Push(2)

	if q.Len() != 2 {
		t.Errorf("Length should be 2")
	}

	if q.Peek().(int) != 2 {
		t.Errorf("First value should be 2")
	}

	q.Pop()

	if q.Peek().(int) != 1 {
		t.Errorf("Next value should be 1")
	}
}
