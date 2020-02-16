package main

import (
	"fmt"

	sorter "github.com/MarcoTomasRodriguez/sort/sorter"
	sorterUtils "github.com/MarcoTomasRodriguez/sort/sorter/utils"
)

func main() {
	t := sorterUtils.GenerateSlice(1e5)
	fmt.Println(sorter.BubbleSort(t))
	fmt.Println(sorter.InsertionSort(t))
	fmt.Println(sorter.HeapSort(t))
	fmt.Println(sorter.MergeSort(t))
}
