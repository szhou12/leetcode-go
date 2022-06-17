package leetcode

// Memorization
// Time = O(n), Space = O(n)
func trap(height []int) int {
	n := len(height)
	res := 0

	// memo
	lmax := make([]int, n)
	lmax[0] = height[0]
	for i := 1; i < n; i++ {
		lmax[i] = max(lmax[i-1], height[i])
	}

	rmax := make([]int, n)
	rmax[n-1] = height[n-1]
	for i := n - 2; i >= 0; i-- {
		rmax[i] = max(rmax[i+1], height[i])
	}

	// Note: i in [1...n-2] bc we don't take into account of both ends
	for i := 1; i < n-1; i++ {
		res += min(lmax[i], rmax[i]) - height[i]
	}
	return res
}

// Time = O(n^2)
func trap_subopt(height []int) int {
	n := len(height)
	res := 0
	for i := 1; i < n-1; i++ {
		lmax := 0
		rmax := 0
		for j := 0; j <= i; j++ {
			lmax = max(lmax, height[j])
		}
		for j := i; j < n; j++ {
			rmax = max(rmax, height[j])
		}
		res += min(lmax, rmax) - height[i]
	}
	return res
}

// Two Pointers
// Time = O(n), Space = O(1)
func trap_opt(height []int) int {
	n := len(height)
	res := 0

	left := 0
	right := n - 1

	lmax, rmax := 0, 0

	for left < right {
		lmax = max(lmax, height[left])
		rmax = max(rmax, height[right])

		// Key!!
		if lmax < rmax {
			res += lmax - height[left]
			left++
		} else {
			res += rmax - height[right]
			right--
		}
	}

	return res
}

func max(a int, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a int, b int) int {
	if a > b {
		return b
	}
	return a
}
