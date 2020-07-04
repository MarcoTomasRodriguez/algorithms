package queue

type (
	Queue struct {
		head *node
		tail *node
		length int
	}
	node struct {
		value interface{}
		next *node
	}
)

func New() *Queue {
	return &Queue{nil, nil, 0}
}

func (queue *Queue) Len() int {
	return queue.length
}

func (queue *Queue) Peek() interface{} {
	if queue.length >= 1 {
		return queue.head.value
	}
	return nil
}

func (queue *Queue) Enqueue(value interface{}) {
	if queue.length == 0 {
		queue.head = &node{value, nil}
	} else if queue.length == 1 {
		queue.tail = &node{value, nil}
		queue.head = &node{queue.head.value, queue.tail}
	} else {
		queue.tail = &node{value, queue.tail}
	}
	queue.length++
}

func (queue *Queue) Dequeue() interface{} {
	if queue.length == 0 {
		return nil
	}
	oldHead := queue.head
	if queue.length == 1 {
		queue.head = nil
	} else {
		queue.head = queue.head.next
	}
	queue.length--
	return oldHead.value
}
