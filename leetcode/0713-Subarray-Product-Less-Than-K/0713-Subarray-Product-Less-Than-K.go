package leetcode

// Two Pointers
// 固定左边界，右边界不停探索至极限
// 到达极限后，收缩左边界(右移一位)，右边界再探索至极限
func numSubarrayProductLessThanK(nums []int, k int) int {
	n := len(nums)
	fast := 0
	product := 1

	count := 0

	for slow := 0; slow < n; slow++ {

		// 注意！！！要防止slow超过fast的情况
		// e.g. [10, 1, 2]  k = 5
		if slow > fast {
			fast = slow // 推进fast与slow同步
			product = 1 // 并且重置product重新计算
		}

		for fast < n && product*nums[fast] < k { // 右边界不停探索至极限
			product *= nums[fast]
			fast++
		}
		// 跳出时说明到达极限 (要不是fast到达终点，要不是product >= k)
		// 注意！！！有效区间是左闭右开 [slow, fast), 其长度就是满足条件的 固定在当前左边界情况下subarray的个数
		// 此时，收缩左边界 (product 砍去左边界的值)
		count += fast - slow
		product /= nums[slow]

	}

	return count

}
