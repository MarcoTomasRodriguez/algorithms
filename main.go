package main

import (
	"fmt"
	"math/rand"

	"github.com/MarcoTomasRodriguez/algorithms/searches"
	"github.com/MarcoTomasRodriguez/algorithms/sorts"
	"github.com/MarcoTomasRodriguez/algorithms/sorts/utils"
)

func main() {
	testSlice := utils.GenerateSlice(1000)
	wanted := rand.Intn(1000)
	fmt.Printf("Initial: %T %v\n", testSlice, len(testSlice))
	sorts.QuickSort(testSlice)
	fmt.Printf("Final: %T %v\n", testSlice, len(testSlice))
	index, found := searches.BinarySearch(testSlice, wanted)
	fmt.Printf("Search: %T %v, Index: %T %v, Found: %T %v\n", wanted, wanted, index, index, found, found)
}
