package structures

// List implementation
type List []interface{}

// IndexOutOfRange returns true if index is out of range, otherwise returns false
func (list List) IndexOutOfRange(index int) bool {
	return !(index >= 0 && index < len(list))
}

// Get returns the element located in a specific index, if none returns nil
func (list List) Get(index int) interface{} {
	if !list.IndexOutOfRange(index) {
		return list[index]
	}
	return nil
}

// Put sets a value to an index in a list if the index is not out of range
func (list List) Put(index int, element interface{}) {
	if !list.IndexOutOfRange(index) {
		list[index] = element
	}
}

// Swap swaps two items by his index
func (list List) Swap(from int, to int) {
	if !list.IndexOutOfRange(from) && !list.IndexOutOfRange(to) {
		list[from], list[to] = list[to], list[from]
	}
}

// Push adds an element to the end of the list
func (list *List) Push(element interface{}) {
	*list = append(*list, element)
}

// Pop removes the last element of the list
func (list *List) Pop() {
	if len(*list) > 0 {
		*list = (*list)[:len(*list)-1]
	}
}
