package main

import (
	"fmt"

	"github.com/MarcoTomasRodriguez/algorithms/sorting/utils"

	"github.com/MarcoTomasRodriguez/algorithms/sorting"
)

func main() {
	testSlice := utils.GenerateSlice(10)
	fmt.Printf("Initial: %T %v\n", testSlice, testSlice)
	sorting.MergeSort(testSlice)
	fmt.Printf("Final: %T %v\n", testSlice, testSlice)
}
