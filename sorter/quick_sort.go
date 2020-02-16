package sorter

import (
	"github.com/MarcoTomasRodriguez/sort/sorter/utils"
)

// QuickSort Best: O(n log n), Average: O(n log n), Worst: O(n^2)
func QuickSort(slice []int) {
	quickSort(slice, 0, len(slice)-1)
}

func quickSort(slice []int, lowest int, highest int) {
	if lowest < highest {
		pivot := partition(slice, lowest, highest)
		quickSort(slice, lowest, pivot-1)
		quickSort(slice, pivot+1, highest)
	}
}

func partition(slice []int, lowest, highest int) int {
	pivot := slice[highest]
	i := lowest
	for j := lowest; j < highest; j++ {
		if slice[j] <= pivot {
			utils.Swap(slice, i, j)
			i++
		}
	}
	utils.Swap(slice, i, highest)
	return i
}
