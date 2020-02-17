package searches

// BinarySearch Best: O(1) Average: O(log n) Worst: O(log n)
func BinarySearch(slice []int, search int) (int, bool) {
	return binarySearch(slice, search, 0, len(slice)-1)
}

func binarySearch(slice []int, search int, low int, high int) (int, bool) {
	if high >= low {
		middle := low + (high-low)/2
		if slice[middle] == search {
			return middle, true
		} else if slice[middle] > search {
			return binarySearch(slice, search, low, middle-1)
		} else {
			return binarySearch(slice, search, middle+1, high)
		}
	} else {
		return -1, false
	}
}
