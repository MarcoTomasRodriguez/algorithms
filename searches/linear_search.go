package searches

// LinearSearch Best: O(1) Average: O(n) Worst: O(n)
func LinearSearch(slice []int, search int) (int, bool) {
	for index, value := range slice {
		if value == search {
			return index, true
		}
	}
	return -1, false
}
