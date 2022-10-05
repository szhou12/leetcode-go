package leetcode

func rotate(nums []int, k int) {
	k %= len(nums) // 防止k > len(nums)的情况发生
	reverse(nums)
	reverse(nums[:k])
	reverse(nums[k:])
}

func reverse(array []int) {
	left, right := 0, len(array)-1
	for left <= right {
		array[left], array[right] = array[right], array[left]
		left++
		right--
	}
}
