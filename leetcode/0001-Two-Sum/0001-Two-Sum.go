package leetcode

func twoSum(nums []int, target int) []int {
	dict := make(map[int]int) // key = num, value = num's index

	for idx, val := range nums {
		if matchIdx, ok := dict[target-val]; ok {
			return []int{idx, matchIdx}
		}
		dict[val] = idx
	}
	return nil
}
