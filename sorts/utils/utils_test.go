package utils

import "testing"

func TestSwap(t *testing.T) {
	testCases := [][]int{
		[]int{2, 1},
		[]int{2, 1, 3},
		[]int{2, 1, 3, 4},
	}
	for _, testCase := range testCases {
		Swap(testCase, 0, 1)
		if testCase[0] != 1 || testCase[1] != 2 {
			t.Error(testCase, "Should swap the first two numbers.")
		}
	}
}

func TestGenerateSlice(t *testing.T) {
	testCases := []int{
		1,
		1e3,
		1e4,
	}
	for _, testCase := range testCases {
		slice := GenerateSlice(testCase)
		if len(slice) != testCase {
			t.Error(testCase, "Should generate a slice of", testCase, "elements.")
		}
	}
}

func TestIsSorted(t *testing.T) {
	truthyTestCases := [][]int{
		[]int{1, 2},
		[]int{1, 2, 3},
		[]int{1, 2, 3, 4},
	}
	falsyTestCases := [][]int{
		[]int{2, 1},
		[]int{2, 3, 1},
		[]int{2, 3, 1, 4},
	}
	for _, testCase := range truthyTestCases {
		if !IsSorted(testCase) {
			t.Error(testCase, "Should be truthy.")
		}
	}
	for _, testCase := range falsyTestCases {
		if IsSorted(testCase) {
			t.Error(testCase, "Should be falsy.")
		}
	}
}
