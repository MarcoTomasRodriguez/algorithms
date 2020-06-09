package sorts

import (
	"testing"

	"github.com/MarcoTomasRodriguez/algorithms/sorts/utils"
)

const (
	Tiny   = 1e2
	Small  = 1e3
	Medium = 1e4
	Big    = 1e5
)

func benchmarkSortAlgorithm(algorithm func(slice []int), lenght int) func(*testing.B) {
	return func(b *testing.B) {
		slice := utils.GenerateSlice(lenght)
		b.ResetTimer()
		BubbleSort(slice)
	}
}

func runBenchmarkSortAlgorithm(b *testing.B, algorithm func(slice []int)) {
	b.Run("Size=Tiny", benchmarkSortAlgorithm(algorithm, Tiny))
	b.Run("Size=Small", benchmarkSortAlgorithm(algorithm, Small))
	b.Run("Size=Medium", benchmarkSortAlgorithm(algorithm, Medium))
	b.Run("Size=Big", benchmarkSortAlgorithm(algorithm, Big))
}

func runTestSortAlgorithm(t *testing.T, algorithm func(slice []int)) {
	slice := utils.GenerateSlice(Tiny)
	algorithm(slice)
	if !utils.IsSorted(slice) {
		t.Error("The slice should be sorted.")
	}
}

func TestBubbleSort(t *testing.T) { runTestSortAlgorithm(t, BubbleSort) }

func TestInsertionSort(t *testing.T) { runTestSortAlgorithm(t, InsertionSort) }

func TestHeapSort(t *testing.T) { runTestSortAlgorithm(t, HeapSort) }

func TestMergeSort(t *testing.T) { runTestSortAlgorithm(t, MergeSort) }

func TestShellSort(t *testing.T) { runTestSortAlgorithm(t, ShellSort) }

func TestQuickSort(t *testing.T) { runTestSortAlgorithm(t, QuickSort) }

func BenchmarkBubbleSort(b *testing.B) { runBenchmarkSortAlgorithm(b, BubbleSort) }

func BenchmarkInsertionSort(b *testing.B) { runBenchmarkSortAlgorithm(b, InsertionSort) }

func BenchmarkHeapSort(b *testing.B) { runBenchmarkSortAlgorithm(b, HeapSort) }

func BenchmarkShellSort(b *testing.B) { runBenchmarkSortAlgorithm(b, ShellSort) }

func BenchmarkMergeSort(b *testing.B) { runBenchmarkSortAlgorithm(b, MergeSort) }

func BenchmarkQuickSort(b *testing.B) { runBenchmarkSortAlgorithm(b, BubbleSort) }
