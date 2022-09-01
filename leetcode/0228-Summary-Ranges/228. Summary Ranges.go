package leetcode

import "strconv"

// My Solution: 快慢指针，[slow, fast-1]为连续数组
func summaryRanges(nums []int) []string {
	var res []string
	// Edge Cases
	if len(nums) == 0 {
		return res
	}
	if len(nums) == 1 {
		res = append(res, strconv.Itoa(nums[0]))
		return res
	}

	slow := 0
	count := 1
	for fast := 1; fast < len(nums); fast++ {
		if nums[fast]-nums[slow] == count {
			count++
		} else {
			var s string
			if slow == fast-1 {
				s = strconv.Itoa(nums[slow])
			} else {
				s = strconv.Itoa(nums[slow]) + "->" + strconv.Itoa(nums[fast-1])
			}

			res = append(res, s)

			// reset slow & count
			slow = fast
			count = 1
		}
	}

	// post-processing
	var endS string
	if slow == len(nums)-1 {
		endS = strconv.Itoa(nums[slow])
	} else {
		endS = strconv.Itoa(nums[slow]) + "->" + strconv.Itoa(nums[len(nums)-1])
	}

	res = append(res, endS)

	return res
}
