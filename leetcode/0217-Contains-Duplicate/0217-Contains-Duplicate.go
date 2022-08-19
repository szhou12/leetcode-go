package leetcode

func containsDuplicate(nums []int) bool {
	dict := make(map[int]bool)

	for _, num := range nums {
		if _, ok := dict[num]; !ok {
			dict[num] = true
		} else {
			return true
		}
	}

	return false
}
