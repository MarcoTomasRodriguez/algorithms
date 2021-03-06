package sort

// ShellSort Best: O(n log n), Average: depends, Worst: O(n log 2 2n)
func ShellSort(slice []int) {
	for i := len(slice) / 2; i > 0; i = (i + 1) * 5 / 11 {
		for j := i; j < len(slice); j++ {
			k, temp := j, slice[j]
			for ; k >= i && slice[k-i] > temp; k -= i {
				slice[k] = slice[k-i]
			}
			slice[k] = temp
		}
	}
}
