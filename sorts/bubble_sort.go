package sorts

import (
	"github.com/MarcoTomasRodriguez/algorithms/sorts/utils"
)

// BubbleSort Best: O(n), Average: O(n^2), Worst: O(n^2)
func BubbleSort(slice []int) {
	for i := len(slice); i > 0; i-- {
		for j := 1; j < i; j++ {
			if slice[j-1] > slice[j] {
				utils.Swap(slice, j, j-1)
			}
		}
	}
}
