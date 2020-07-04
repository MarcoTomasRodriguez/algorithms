package sort

// MergeSort Best: O(n log n), Average: O(n log n), Worst: O(n log n)
func MergeSort(slice []int) {
	copy(slice, mergeSort(slice))
}

func mergeSort(slice []int) []int {
	if len(slice) <= 1 {
		return slice
	}
	middle := len(slice) / 2
	firstHalf := mergeSort(slice[:middle])
	secondHalf := mergeSort(slice[middle:])
	return merge(firstHalf, secondHalf)
}

func merge(firstHalf []int, secondHalf []int) []int {
	slice := make([]int, len(firstHalf)+len(secondHalf))
	for i := 0; ; i++ {
		if len(firstHalf) > 0 && len(secondHalf) > 0 {
			if firstHalf[0] > secondHalf[0] {
				slice[i] = secondHalf[0]
				secondHalf = secondHalf[1:]
			} else {
				slice[i] = firstHalf[0]
				firstHalf = firstHalf[1:]
			}
		} else if len(firstHalf) > 0 {
			copy(slice[i:], firstHalf)
			break
		} else if len(secondHalf) > 0 {
			copy(slice[i:], secondHalf)
			break
		}
	}
	return slice
}
