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
