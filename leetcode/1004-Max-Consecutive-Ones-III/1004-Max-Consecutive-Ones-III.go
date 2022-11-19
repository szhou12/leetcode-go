package leetcode

// Optimized Solution:
// 模版Flex
// 优化: 每一轮挡板l, 挡板r不从挡板l处重新开始, 而是从上一轮r停留处继续往前探
func longestOnes(nums []int, k int) int {
	res := 0

	r := 0
	counter := k
	// 固定左边界l, 延伸右边界r
	for l := 0; l < len(nums); l++ {
		for r < len(nums) && (nums[r] == 1 || counter > 0) { // 注意: 延伸右边界有可能出现出界的情况，一定要检查!!!
			if nums[r] == 0 {
				counter--
			}
			// res = max(res, r-l+1)
			r++
		} // 退出for-loop: 说明在当前左边界l下, 已到达最远右边界; 可以前进左边界l进入下一轮, 尝试新的最长了

		// update result
		res = max(res, r-l)

		// 进入下一轮前，如果左边界l吐出来0，说明在下一轮我们可以多一次flip
		if nums[l] == 0 {
			counter++
		}

	}
	return res
}

// My Initial Soution: 固定每一个挡板l, 挡板r尽可能往右边延伸, 看最长能延伸到哪儿
func longestOnes_firstSoln(nums []int, k int) int {
	counter := k
	gmax := 0
	for l := 0; l < len(nums); l++ {
		r := l // 此处可优化
		for r < len(nums) && (nums[r] == 1 || counter > 0) {
			if nums[r] == 0 {
				counter--
			}
			r++
		}
		gmax = max(gmax, r-l)
		// reset counter
		counter = k
	}
	return gmax
}

func max(a int, b int) int {
	if a > b {
		return a
	}
	return b
}
