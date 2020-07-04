package search

import (
	"math/rand"
	"testing"

	"github.com/MarcoTomasRodriguez/algorithms/sort"
	"github.com/MarcoTomasRodriguez/algorithms/sort/utils"
)

const (
	Tiny   = 1e2
	Small  = 1e3
	Medium = 1e4
	Big    = 1e5
	Large  = 1e6
)

func benchmarkSearchAlgorithm(algorithm func(slice []int, search int) (int, bool), lenght int) func(*testing.B) {
	return func(b *testing.B) {
		slice := utils.GenerateSlice(lenght)
		sort.QuickSort(slice)
		b.ResetTimer()
		algorithm(slice, rand.Intn(lenght))
	}
}

func runBenchmarkSortAlgorithm(b *testing.B, algorithm func(slice []int, search int) (int, bool)) {
	b.Run("Size=Tiny", benchmarkSearchAlgorithm(algorithm, Tiny))
	b.Run("Size=Small", benchmarkSearchAlgorithm(algorithm, Small))
	b.Run("Size=Medium", benchmarkSearchAlgorithm(algorithm, Medium))
	b.Run("Size=Big", benchmarkSearchAlgorithm(algorithm, Big))
	b.Run("Size=Large", benchmarkSearchAlgorithm(algorithm, Large))
}

// To test a search algorithm this uses the LinearSearch function.
func runTestSearchAlgorithm(t *testing.T, algorithm func(slice []int, search int) (int, bool)) {
	slice := utils.GenerateSlice(Tiny)
	random := rand.Intn(Tiny)
	sort.QuickSort(slice)
	linearIndex, linearFound := LinearSearch(slice, random)
	algorithmIndex, algorithmFound := algorithm(slice, random)
	if linearFound != algorithmFound || linearIndex != algorithmIndex {
		t.Error()
	}
}

func TestLinearSearch(t *testing.T) { runTestSearchAlgorithm(t, LinearSearch) }

func TestBinarySearch(t *testing.T) { runTestSearchAlgorithm(t, BinarySearch) }

func BenchmarkLinearSearch(b *testing.B) { runBenchmarkSortAlgorithm(b, LinearSearch) }

func BenchmarkBinarySearch(b *testing.B) { runBenchmarkSortAlgorithm(b, BinarySearch) }
