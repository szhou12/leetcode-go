package leetcode

// My Solution: 用一个额外的array来模拟hashset
func findDisappearedNumbers(nums []int) []int {

	dict := make([]int, len(nums)+1)
	for _, cur := range nums {
		dict[cur] = 1
	}

	var res []int

	for i := 1; i <= len(nums); i++ {
		if dict[i] == 0 {
			res = append(res, i)
		}
	}

	return res
}

// Optimal Solution: 关联元素与index，不需要额外空间
func findDisappearedNumbers_optimal(nums []int) []int {

	for _, num := range nums {
		index := abs(num) - 1 // 注意索引，元素大小从 1 开始，有一位index偏移
		if nums[index] < 0 {
			// 之前已经把对应index位置上的元素nums[index]变成负数了，
			// 这说明 num 重复出现了两次，但我们不用做，让nums[index]继续保持负数
		} else {
			// 把索引 |num| - 1 置为负数
			nums[index] *= -1
		}
	}

	var res []int

	for i := 0; i < len(nums); i++ {
		if nums[i] > 0 {
			// 说明没有元素和这个索引对应，即找到一个缺失元素
			res = append(res, i+1) // 注意索引，要把index偏移回去，即 i=0 意味着该元素是 i+1=1
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
