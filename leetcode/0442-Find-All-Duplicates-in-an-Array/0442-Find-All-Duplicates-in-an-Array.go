package leetcode

// Solution: 思路和0448完全一致 - 关联元素与index
func findDuplicates(nums []int) []int {

	var res []int
	for _, num := range nums {
		index := abs(num) - 1 // 注意索引，元素大小从 1 开始，有一位索引偏移
		if nums[index] < 0 {
			// 之前已经把对应索引的元素变成负数了，
			// 这说明 num 重复出现了两次
			res = append(res, abs(num)) // 注意这里要取绝对值，因为可能被变成了负数
		} else {
			// 把索引 num - 1 置为负数
			nums[index] *= -1
		}
	}

	return res
}

func abs(num int) int {
	if num < 0 {
		return -num
	}

	return num
}
