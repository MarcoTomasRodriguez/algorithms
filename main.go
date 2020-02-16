package main

import (
	"fmt"

	"github.com/MarcoTomasRodriguez/algorithms/sorts/utils"

	"github.com/MarcoTomasRodriguez/algorithms/sorts"
)

func main() {
	testSlice := utils.GenerateSlice(10)
	fmt.Printf("Initial: %T %v\n", testSlice, testSlice)
	sorts.MergeSort(testSlice)
	fmt.Printf("Final: %T %v\n", testSlice, testSlice)
}
