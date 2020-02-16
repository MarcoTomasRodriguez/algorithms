package utils

import (
	"math/rand"
	"time"
)

// Swap two elements in a slice of integers.
func Swap(slice []int, first int, second int) {
	slice[first], slice[second] = slice[second], slice[first]
}

// GenerateSlice creates a slice with a defined size filled with random numbers
func GenerateSlice(size int) []int {
	slice := make([]int, size, size)
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < size; i++ {
		slice[i] = rand.Intn(999) - rand.Intn(999)
	}
	return slice
}

// IsSorted check if the array is well sorted.
func IsSorted(slice []int) bool {
	for i := range slice {
		if i == len(slice)-1 {
			break
		}
		if slice[i] > slice[i+1] {
			return false
		}
	}
	return true
}
