package solvers

import (
	"fmt"
	"math/rand"
	"time"
)

// MagicSquare is the representation of the square in integers
type MagicSquare struct {
	size   int
	sum    int
	square [][]int
}

func (magicSquare MagicSquare) set(index int, value int) {
	size := magicSquare.size
	highOrder, lowOrder := index/size, index%size
	magicSquare.square[highOrder][lowOrder] = value
}

// generateStartingPoint ...
func (magicSquare MagicSquare) generateStartingPoint() {
	rand.Seed(time.Now().UnixNano())
	fullLenght := magicSquare.size*magicSquare.size - 1
	randomInt := rand.Intn(fullLenght)
	magicSquare.set(randomInt, 1)
}

// NewMagicSquare creates a new MagicSquare filled with zeros
func NewMagicSquare(size int) *MagicSquare {
	magicSquare := MagicSquare{
		square: make([][]int, size, size),
		size:   size,
		sum:    0,
	}
	for index := range magicSquare.square {
		magicSquare.square[index] = make([]int, size, size)
	}
	return &magicSquare
}

// InitializeMagicSquare generates all the things needed to start solving the MagicSquare
func (magicSquare MagicSquare) InitializeMagicSquare() {
	magicSquare.generateStartingPoint()
}

// SolveMagicSquare solves the MagicSquare
func (magicSquare MagicSquare) SolveMagicSquare() {

}

// ExecMagicSquare creates, initializes and solves the MagicSquare printing the result
func ExecMagicSquare() {
	ms := NewMagicSquare(4)
	ms.InitializeMagicSquare()
	ms.SolveMagicSquare()
	fmt.Println("Magic Square:", ms.square)
}
