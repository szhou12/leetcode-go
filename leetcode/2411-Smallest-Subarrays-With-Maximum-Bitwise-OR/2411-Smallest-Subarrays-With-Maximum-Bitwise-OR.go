package leetcode

func smallestSubarrays(nums []int) []int {
	n := len(nums)
	right := n - 1
	count := make([]int, 32) // int has 32 bits
	res := make([]int, n)

	for left := n - 1; left >= 0; left-- {
		// 吃
		// if i-th bit position of nums[left] is 1, then count[i]++
		for k := 0; k < 32; k++ {
			count[k] += (nums[left] >> k) & 1
		}

		// 吐
		for left < right && canRemove(nums[right], count) {
			for k := 0; k < 32; k++ {
				count[k] -= (nums[right] >> k) & 1
			}
			right--
		}

		// update result
		res[left] = right - left + 1
	}

	return res
}

func canRemove(num int, count []int) bool {
	// we can't remove num if any bit position of count[] will become 0 after removal
	// it means num is contributing 1 at this bit position
	for k := 0; k < 32; k++ {
		if count[k] > 0 && (count[k]-(num>>k&1) <= 0) {
			return false
		}
	}

	return true
}
