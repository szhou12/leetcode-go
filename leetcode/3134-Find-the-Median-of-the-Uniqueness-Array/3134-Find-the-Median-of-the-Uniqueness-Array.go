package leetcode

func medianOfUniquenessArray(nums []int) int {
	n := len(nums)
	total := n * (n + 1) / 2 // total number of subarrays
	// guess k (median distinct numbers of subarray)
	// 猜一个distinct number的个数k，看最多k个不同元素的subarray总数是否刚好是中位数
	left, right := 1, n // left, right := # of distinct numbers in a subarray
	for left < right {
		mid := left + (right-left)/2
		if subarraysMostKDistinct(nums, mid) >= (total+1)/2 {
			// WHY (total+1)/2?
			// 当total为奇数时, 比如total=3, 因为向下取整, total/2没有包含到中位数, 而是中位数的左边(取到了1而中位数是2)
			right = mid
		} else {
			left = mid + 1
		}
	}

	return left
}

// Return: # of subarrays with at most k distinct numbers
func subarraysMostKDistinct(nums []int, k int) int {
	n := len(nums)
	res := 0
	
	upperRight := atMost(nums, k)
	for i := 0; i < n; i++ {
		res += upperRight[i] - i
	}

	return res
}

// Return: upper[i] := for each left index i, the first index j s.t. nums[i...(j-1)] that has <= k distinct numbers
func atMost(nums []int, k int) []int {
	n := len(nums)
	upper := make([]int, n)
	window := make(map[int]int) // record appearances of each distinct number within the window
	count := 0 // count distinct numbers within the window
	right := 0

	for left := 0; left < n; left++ {
		// 吃
		for right < n && count <= k {
			rightElement := nums[right]
			if window[rightElement] == 0 {
				count++
			}
			window[rightElement]++
			right++
		}

		// update upper[]
		if !(right < n) && count <= k {
			upper[left] = right
		} else if (right < n) && !(count <= k) {
			upper[left] = right - 1
		} else {
			upper[left] = right - 1
		}

		// 吐
		leftElement := nums[left]
		window[leftElement]--
		if window[leftElement] == 0 {
			count--
		}
	}

	return upper
}