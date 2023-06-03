package leetcode

func maxTotalFruits(fruits [][]int, startPos int, k int) int {
	n := len(fruits)
	fruitPos := make([]int, n)
	prefixSum := make([]int, n)
	for i, fruit := range fruits {
		fruitPos[i] = fruit[0]
		if i == 0 {
			prefixSum[i] = 0 + fruit[1]
		} else {
			prefixSum[i] = prefixSum[i-1] + fruit[1]
		}
	}

	res := 0

	// Turn on left side
	// 固定右边界，从左收缩左边界
	rightStart := lowerBound(fruitPos, startPos)
	left := 0
	for right := rightStart; right < n; right++ {
		// 收缩左边界: 从界外走到界内
		for fruitPos[left] <= startPos && fruitPos[right]-startPos+2*(startPos-fruitPos[left]) > k {
			left++
		}

		if fruitPos[left] <= startPos {
			if left-1 >= 0 {
				res = max(res, prefixSum[right]-prefixSum[left-1])
			} else {
				res = max(res, prefixSum[right]-0)
			}
		} else if fruitPos[right]-startPos <= k {
			if left-1 >= 0 {
				res = max(res, prefixSum[right]-prefixSum[left-1])
			} else {
				res = max(res, prefixSum[right]-0)
			}
		}
	}

	// Turn on right side
	// 固定左边界，从右收缩右边界
	leftStart := upperBound(fruitPos, startPos) - 1
	right := n - 1
	for left := leftStart; left >= 0; left-- {
		// 收缩右边界: 从界外走到界内
		for fruitPos[right] >= startPos && 2*(fruitPos[right]-startPos)+startPos-fruitPos[left] > k {
			right--
		}

		if fruitPos[right] >= startPos {
			if left-1 >= 0 {
				res = max(res, prefixSum[right]-prefixSum[left-1])
			} else {
				res = max(res, prefixSum[right]-0)
			}
		} else if startPos-fruitPos[left] <= k {
			if left-1 >= 0 {
				res = max(res, prefixSum[right]-prefixSum[left-1])
			} else {
				res = max(res, prefixSum[right]-0)
			}
		}
	}

	return res

}

// first index whose value > target
func upperBound(nums []int, target int) int {
	left := 0
	right := len(nums) - 1
	for left < right {
		mid := left + (right-left)/2
		if nums[mid] <= target {
			left = mid + 1
		} else {
			right = mid
		}
	}

	if nums[left] <= target {
		return left + 1
	} else {
		return left
	}
}

// first index whose value >= target
func lowerBound(nums []int, target int) int {
	left := 0
	right := len(nums) - 1
	for left < right {
		mid := left + (right-left)/2
		if nums[mid] < target {
			left = mid + 1
		} else {
			right = mid
		}
	}

	if nums[left] < target {
		return left + 1
	}
	return left
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
