package sorting

import (
	"fmt"
	"math/rand"
	"testing"
)

// generateRandomArray creates an array of given size with random integers
func generateRandomArray(size int) []int {
	arr := make([]int, size)
	for i := 0; i < size; i++ {
		arr[i] = rand.Intn(1000000) // Random numbers between 0 and 1000000
	}
	return arr
}

// generateSortedArray creates a sorted array of given size
func generateSortedArray(size int) []int {
	arr := make([]int, size)
	for i := 0; i < size; i++ {
		arr[i] = i
	}
	return arr
}

// generateReverseSortedArray creates a reverse sorted array of given size
func generateReverseSortedArray(size int) []int {
	arr := make([]int, size)
	for i := 0; i < size; i++ {
		arr[i] = size - i
	}
	return arr
}

/*
* Benchmark results:
* Operations per second | Time per operation | Memory allocations | Bytes allocated per operation
*
* Speedup = Sequential Time / Parallel Time
* To compare speedup, run the following command:
* cd sorting && go test -bench=. -benchtime=1s
*/
func BenchmarkQuickSort(b *testing.B) {
	sizes := []int{1000, 10000, 100000}

	for _, size := range sizes {
		// Test random arrays
		b.Run(fmt.Sprintf("Sequential_Random_%d", size), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				b.StopTimer()
				arr := generateRandomArray(size)
				b.StartTimer()
				QuickSort(arr, false)
			}
		})

		b.Run(fmt.Sprintf("Parallel_Random_%d", size), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				b.StopTimer()
				arr := generateRandomArray(size)
				b.StartTimer()
				QuickSort(arr, true)
			}
		})

		// Test sorted arrays
		// b.Run(fmt.Sprintf("Sequential_Sorted_%d", size), func(b *testing.B) {
		// 	for i := 0; i < b.N; i++ {
		// 		b.StopTimer()
		// 		arr := generateSortedArray(size)
		// 		b.StartTimer()
		// 		QuickSort(arr, false)
		// 	}
		// })

		// b.Run(fmt.Sprintf("Parallel_Sorted_%d", size), func(b *testing.B) {
		// 	for i := 0; i < b.N; i++ {
		// 		b.StopTimer()
		// 		arr := generateSortedArray(size)
		// 		b.StartTimer()
		// 		QuickSort(arr, true)
		// 	}
		// })

		// // Test reverse sorted arrays
		// b.Run(fmt.Sprintf("Sequential_ReverseSorted_%d", size), func(b *testing.B) {
		// 	for i := 0; i < b.N; i++ {
		// 		b.StopTimer()
		// 		arr := generateReverseSortedArray(size)
		// 		b.StartTimer()
		// 		QuickSort(arr, false)
		// 	}
		// })

		// b.Run(fmt.Sprintf("Parallel_ReverseSorted_%d", size), func(b *testing.B) {
		// 	for i := 0; i < b.N; i++ {
		// 		b.StopTimer()
		// 		arr := generateReverseSortedArray(size)
		// 		b.StartTimer()
		// 		QuickSort(arr, true)
		// 	}
		// })
	}
}

// TestQuickSortCorrectness tests if both sequential and parallel QuickSort produce correct results
func TestQuickSortCorrectness(t *testing.T) {
	sizes := []int{10, 100, 1000, 10000}

	for _, size := range sizes {
		// Test random arrays
		arr1 := generateRandomArray(size)
		arr2 := make([]int, size)
		copy(arr2, arr1)

		QuickSort(arr1, false) // Sequential
		QuickSort(arr2, true)  // Parallel

		// Verify both arrays are sorted
		for i := 1; i < size; i++ {
			if arr1[i] < arr1[i-1] {
				t.Errorf("Sequential QuickSort failed: array not sorted at index %d", i)
			}
			if arr2[i] < arr2[i-1] {
				t.Errorf("Parallel QuickSort failed: array not sorted at index %d", i)
			}
		}
	}
}
