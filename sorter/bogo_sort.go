package sorter

import (
	"math/rand"
	"time"

	"github.com/MarcoTomasRodriguez/sort/sorter/utils"
)

// BogoSort is a non-efficient permutation sort algorithm
// For educational purposes only
func BogoSort(slice []int) {
	for !utils.IsSorted(slice) {
		rand.Seed(time.Now().UnixNano())
		rand.Shuffle(len(slice), func(i, j int) { utils.Swap(slice, i, j) })
	}
}
