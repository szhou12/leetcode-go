package leetcode

import "math"

func minLength(s string, numOps int) int {
	n := len(s)

	// convert "010" to [0, 1, 0]
	// nums[] used for handling special case k = 1
	nums := make([]int, 0)
	for _, char := range s {
		nums = append(nums, int(char-'0'))
	}

	// Two Pointers: count length of identical substrs
	lens := make([]int, 0)
	for i := 0; i < n; {
		j := i
		for j < n && s[i] == s[j] {
			j++
		}

		// identical substr [i:j-1]
		lens = append(lens, j-i)

		i = j
	}

	// Binary Search
	left, right := 1, n
	for left < right {
		// mid := target upper limit of identical length
		mid := left + (right-left)/2
		if valid(lens, nums, mid, numOps) {
			right = mid
		} else {
			left = mid + 1
		}
	}

	return left

}

func valid(lens []int, nums []int, k int, numOps int) bool {
	// handle special case k = 1
	if k == 1 {
		flips := 0

		// 010101... -> even index=0, odd index=1
		for i, num := range nums {
			// flip if even index should give 0 but num=1; odd index should give 1 but num=0
			if num != i%2 {
				flips++
			}
		}
		if flips <= numOps {
			return true
		}

		flips = 0
		// 101010... -> even index=1, odd index=0
		for i, num := range nums {
			// flip if even index shouldn't give 0 but num=0
			if num == i%2 {
				flips++
			}
		}
		if flips <= numOps {
			return true
		}

		return false
	}

	// Greedy
	ops := 0
	for _, x := range lens {
		ops += int(math.Ceil(float64(x-k) / float64(k+1)))
		if ops > numOps {
			return false
		}
	}

	return true
}
