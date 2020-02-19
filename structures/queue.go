package structures

// Queue implementation
type Queue []interface{}

// EnQueue adds an item to the end of the queue
func (queue *Queue) EnQueue(element interface{}) {
	*queue = append(*queue, element)
}

// DeQueue removes the first element of the queue if any
func (queue *Queue) DeQueue() {
	if len(*queue) > 0 {
		*queue = (*queue)[1:]
	}
}
