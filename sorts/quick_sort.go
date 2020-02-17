package sorts

import (
	"github.com/MarcoTomasRodriguez/algorithms/sorts/utils"
)

// QuickSort Best: O(n log n), Average: O(n log n), Worst: O(n^2)
func QuickSort(slice []int) {
	quickSort(slice, 0, len(slice)-1)
}

func quickSort(slice []int, low int, high int) {
	if low < high {
		pivot := partition(slice, low, high)
		quickSort(slice, low, pivot-1)
		quickSort(slice, pivot+1, high)
	}
}

func partition(slice []int, low, high int) int {
	pivot := slice[high]
	i := low
	for j := low; j < high; j++ {
		if slice[j] <= pivot {
			utils.Swap(slice, i, j)
			i++
		}
	}
	utils.Swap(slice, i, high)
	return i
}
