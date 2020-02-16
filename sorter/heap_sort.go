package sorter

import (
	"github.com/MarcoTomasRodriguez/sort/sorter/utils"
)

// HeapSort Best: O(n log n), Average: O(n log n), Worst: O(n log n)
func HeapSort(slice []int) {
	heapify(slice)
	for i := len(slice) - 1; i > 0; i-- {
		slice[0], slice[i] = slice[i], slice[0]
		siftDown(slice, 0, i)
	}
}

func heapify(slice []int) {
	for i := (len(slice) - 1) / 2; i >= 0; i-- {
		siftDown(slice, i, len(slice))
	}
}

func siftDown(slice []int, lowest int, highest int) {
	root := lowest
	for {
		child := root*2 + 1
		if child >= highest {
			break
		}
		if child+1 < highest && slice[child] < slice[child+1] {
			child++
		}
		if slice[root] < slice[child] {
			utils.Swap(slice, root, child)
			root = child
		} else {
			break
		}

	}
}
