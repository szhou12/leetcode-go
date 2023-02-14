package leetcode

/**
把nums分成 L, M, R三段:
{X X X X X }  X X X  {X X X X X }
		 i-1  i      i+k

物理意义:
	leftMax[i] := the largest k-length subarray sum from [0:i]
	rightMax[i] := the largest k-length subarray sum from [i:n-1]

题目要找的是:
	find index i s.t. leftMax[i-1] + midSum[i:i+k-1] + rightMax[i+k] is MAX

注意:
	leftMax[i]中 i 代表L段的右端点
	rightMax[i]中 i 代表R段的左端点

注意:
	L段 从左往右滑动从而更新 leftMax
	R段 从右往左滑动从而更新 rightMax

注意:
	本题涉及到 subarray sum, 所以事先准备一个 prefix-sum array方便计算： subarray-sum[l:r] = prefix-sum[r] - prefix-sum[l-1]
*/
func maxSumOfThreeSubarrays(nums []int, k int) []int {
	n := len(nums)

	presum := make([]int, n)
	for i := 0; i < n; i++ {
		if i == 0 {
			presum[i] = 0 + nums[i]
		} else {
			presum[i] = presum[i-1] + nums[i]
		}
	}

	// pass L段: i代表右端点; 滑动窗口从左到右滑动
	leftMax := make([]int, n)
	leftIdx := make([]int, n)
	maxSum := 0
	maxIdx := 0
	for i := k - 1; i < n; i++ {
		// [i-k+1 : i]
		var sum int
		// check if out-of-bound: i won't, but i-k will when i = k-1
		if i-k < 0 {
			sum = presum[i] - 0
		} else {
			sum = presum[i] - presum[i-k]
		}

		// update: index尽量靠左, 因为滑窗从左到右滑，所以只有 > 才更新
		if sum > maxSum {
			maxSum = sum
			maxIdx = i - k + 1 // 记录左端点
		}
		leftMax[i] = maxSum
		leftIdx[i] = maxIdx
	}

	// pass R段: i代表左端点; 滑动窗口从右到左滑动
	rightMax := make([]int, n)
	rightIdx := make([]int, n)
	maxSum = 0
	maxIdx = 0
	for i := n - k; i >= 0; i-- {
		// [i : i+k-1]
		var sum int
		// check if out-of-bound: i+k-1 won't, but i-1 will when i = 0
		if i-1 < 0 {
			sum = presum[i+k-1] - 0
		} else {
			sum = presum[i+k-1] - presum[i-1]
		}

		// update: index尽量靠左, 因为滑窗从右到左滑，所以 >= 就更新
		if sum >= maxSum {
			maxSum = sum
			maxIdx = i
		}
		rightMax[i] = maxSum
		rightIdx[i] = maxIdx
	}

	// pass M段: i代表左端点; 滑动窗口从右到左滑动
	maxSum = 0
	indices := make([]int, 3)
	for i := k; (i+k)+k <= n; i++ { // 左右两边各自预留k长度
		// [i : i+k-1]
		midSum := presum[i+k-1] - presum[i-1]

		// update: index尽量靠左, 因为滑窗从左到右滑，所以只有 > 才更新
		// L, M, R段 三段之和
		if leftMax[i-1]+midSum+rightMax[i+k] > maxSum {
			maxSum = leftMax[i-1] + midSum + rightMax[i+k]
			indices = []int{leftIdx[i-1], i, rightIdx[i+k]}
		}
	}

	return indices
}
