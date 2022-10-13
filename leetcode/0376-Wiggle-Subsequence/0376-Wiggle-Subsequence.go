package leetcode

// DP solution - 容易想, 但是时间复杂度高
// dp[i][0] = longest wiggle subseq ending at i and diff < 0
// dp[i][1] = longest wiggle subseq ending at i and diff > 0
// Base Case
// dp[0][0/1] = 0
// dp[1...n][0/1] = 1
// Recurrence
// dp[i][0] = max(dp[j][1]+1) where j < i and nums[i] - nums[j] < 0
// dp[i][1] = max(dp[j][0]+1) where j < i and nums[i] - nums[j] > 0
func wiggleMaxLength_DP(nums []int) int {
	n := len(nums)
	nums = append([]int{-1}, nums...)

	dp := make([][2]int, n+1)

	// Base Cases
	dp[0][0] = 0
	dp[0][1] = 0
	for i := 1; i <= n; i++ {
		dp[i][0] = 1
		dp[i][1] = 1
	}

	// Recurrence
	res := 0
	for i := 2; i <= n; i++ {
		for j := 1; j < i; j++ {
			if nums[i]-nums[j] < 0 {
				dp[i][0] = max(dp[i][0], dp[j][1]+1)
			} else if nums[i]-nums[j] > 0 {
				dp[i][1] = max(dp[i][1], dp[j][0]+1)
			}

		}
		temp := max(dp[i][0], dp[i][1])
		res = max(res, temp)

	}
	return res

}

func max(a int, b int) int {
	if a > b {
		return a
	}
	return b
}

// Optimal Solution - Greedy
// 让局部峰值尽可能的保持局部峰值，然后删除单一坡度上的节点，统计局部峰值个数
func wiggleMaxLength_Greedy(nums []int) int {
	n := len(nums)
	if n < 2 {
		return n
	}
	res := 1 // 记录峰值个数，长度>1的数组，峰值个数至少为1个 e.g [2,2,2,2,2]
	// 判断只有头两个元素的情况下，是否有峰值
	prevDiff := nums[1] - nums[0]
	if prevDiff != 0 {
		res++
	}

	for i := 2; i < n; i++ {
		curDiff := nums[i] - nums[i-1]

		//如果当前差值和上一个差值为一正一负
		//等于0的情况表示初始时的prevDiff, 即，头两个元素相等的情况
		if (curDiff > 0 && prevDiff <= 0) || (curDiff < 0 && prevDiff >= 0) {
			res++
			prevDiff = curDiff
		}
	}
	return res
}
