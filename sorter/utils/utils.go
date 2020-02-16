package utils

import (
	"math/rand"
	"time"
)

// Swap two elements in a slice of integers.
func Swap(slice []int, first int, second int) {
	slice[first], slice[second] = slice[second], slice[first]
}

// MoveFirstElement from a slice to another
func MoveFirstElement(from []int, to []int, toPosition int) {
	to[toPosition] = from[0]
	from = from[1:]
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
