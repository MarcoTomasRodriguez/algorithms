package main

import (
	"fmt"

	"github.com/MarcoTomasRodriguez/sort/sorter/utils"

	"github.com/MarcoTomasRodriguez/sort/sorter"
)

func main() {
	testSlice := utils.GenerateSlice(10)
	fmt.Printf("Initial: %T %v\n", testSlice, testSlice)
	sorter.MergeSort(testSlice)
	fmt.Printf("Final: %T %v\n", testSlice, testSlice)
}
