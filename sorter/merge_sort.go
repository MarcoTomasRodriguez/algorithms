package sorter

import "github.com/MarcoTomasRodriguez/sort/sorter/utils"

// MergeSort Best: O(n log n), Average: O(n log n), Worst: O(n log n)
func MergeSort(slice []int) []int {
	if len(slice) < 2 {
		return slice
	}
	halfIndex := len(slice) / 2
	firstHalf := MergeSort(slice[0:halfIndex])
	secondHalf := MergeSort(slice[halfIndex+1:])
	return merge(firstHalf, secondHalf)
}

func merge(firstHalf []int, secondHalf []int) []int {
	slice := make([]int, len(firstHalf)+len(secondHalf))
	i := 0

	for len(firstHalf) > 0 && len(secondHalf) > 0 {
		if firstHalf[0] < secondHalf[0] {
			utils.MoveFirstElement(firstHalf, slice, i)
		} else {
			utils.MoveFirstElement(secondHalf, slice, i)
		}
		i++
	}

	for j := 0; j < len(firstHalf); j++ {
		slice[i] = firstHalf[j]
		i++
	}
	for j := 0; j < len(secondHalf); j++ {
		slice[i] = secondHalf[j]
		i++
	}

	return slice
}
