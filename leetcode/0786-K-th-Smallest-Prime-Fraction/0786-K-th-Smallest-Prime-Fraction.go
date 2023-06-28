package leetcode

func kthSmallestPrimeFraction(arr []int, k int) []int {
	n := len(arr)

	left := 0.0
	right := 1.0

	var mid float64
	for left < right {
		mid = (left + right) / 2

		count := 0
		// for each arr[i], count # of j s.t. arr[i]/arr[j] <= mid
		for i := 0; i < n; i++ {
			// arr[i]/arr[j] <= mid -> arr[j] >= arr[i]/mid
			j := lowerBound(arr, float64(arr[i])/mid)
			count += (n-1) - j + 1
		}

		if count > k {
			right = mid
		} else if count < k {
			left = mid
		} else {
			break
		}
	}

	var res []int
	ans := 0.0
	for i := 0; i < n;i++ {
		j := lowerBound(arr, float64(arr[i])/mid)
		if j < n && float64(arr[i])/float64(arr[j]) > ans {
			ans = float64(arr[i])/float64(arr[j])
			res = []int{arr[i], arr[j]}
		}
	}

	return res
}


// first index >= target
func lowerBound(nums []int, target float64) int {
	left, right := 0, len(nums)-1

	for left < right {
		mid := left + (right-left)/2

		if float64(nums[mid]) < target {
			left = mid + 1
		} else {
			right = mid
		}
	}

	if float64(nums[left]) < target {
		return left + 1
	}
	return left
}