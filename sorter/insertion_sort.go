package sorter

import (
	"github.com/MarcoTomasRodriguez/sort/sorter/utils"
)

// InsertionSort Best: O(n), Average: O(n^2), Worst: O(n^2)
func InsertionSort(slice []int) {
	for i := 1; i < len(slice); i++ {
		for j := i; j > 0 && slice[j-1] > slice[j]; j-- {
			utils.Swap(slice, j, j-1)
		}
	}
}
