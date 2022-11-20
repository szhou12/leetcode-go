package leetcode

// Two Pointers
// 固定左边界，右边界不停探索至极限
// 到达极限后，收缩左边界(右移一位)，右边界再探索至极限
func numSubarrayProductLessThanK(nums []int, k int) int {
	n := len(nums)
	right := 0
	product := 1

	res := 0

	for left := 0; left < n; left++ {

		// 注意！！！要防止left超过right的情况
		// e.g. [10, 1, 2]  k = 5, 窗口一开始是[10]
		if left > right {
			right = left // 推进right与left同步
			product = 1  // 并且重置product重新计算
		}

		for right < n && product*nums[right] < k { // 右边界不停探索至极限
			product *= nums[right]
			right++
		}
		// 跳出时说明到达极限 (要不是right到达终点，要不是product >= k)
		// 注意！！！有效区间是左闭右开 [left, right), 其长度就是满足条件的 固定在当前左边界情况下subarray的个数
		// 此时，收缩左边界 (product 砍去左边界的值)

		// update result
		res += right - left
		// 吐
		product /= nums[left]

	}

	return res

}

// 第二种写法
func numSubarrayProductLessThanK_2(nums []int, k int) int {
	right := 0
	product := 1

	res := 0
	for left := 0; left < len(nums); left++ {
		for right < len(nums) && product*nums[right] < k {
			product *= nums[right]
			right++
		}

		// right的loop没有进入有一种特殊情况: 窗口内只有一个元素, 而这个元素 >= k
		// 这种情况下, right没有推进, 需要人为推进
		// e.g. 窗口=[100], k=100
		// 			  lr
		// right 指向这个元素, 但我们要规避掉, 所以 不更新结果, 直接 right++ 跳过该元素, 使 right 和 left 在下一轮共同进入下一格 (ie. 推进right与left同步)
		if left == right {
			right++
		} else {
			// update result
			res = right - left
			// 吐
			product /= nums[left]
		}
	}
	return res
}
