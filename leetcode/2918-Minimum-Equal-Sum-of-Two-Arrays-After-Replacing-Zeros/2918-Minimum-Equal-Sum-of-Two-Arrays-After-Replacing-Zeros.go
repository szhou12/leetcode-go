package leetcode

func minSum(nums1 []int, nums2 []int) int64 {
	sum1, sum2 := 0, 0
	miss1, miss2 := 0, 0 // # of missing places

	for _, e := range nums1 {
		if e == 0 {
			miss1++
		}
		sum1 += e
	}
	for _, e := range nums2 {
		if e == 0 {
			miss2++
		}
		sum2 += e
	}

	// 分类讨论
	// Case 1: both arrays have NO missing places
	if miss1 == 0 && miss2 == 0 {
		if sum1 == sum2 {
			return int64(sum1)
		}
		return -1
	}
	// Case 2 & 3: one array has NO missing places while the other does
	if miss1 == 0 {
		if sum1 < sum2+miss2 {
			return -1
		}
		return int64(sum1)
	}
	if miss2 == 0 {
		if sum2 < sum1+miss1 {
			return -1
		}
		return int64(sum2)
	}
	// Case 4: both arrays have missing places, then when filling missing places with 1, whoever gets smaller sum is the answer
	return int64(max(sum1+miss1, sum2+miss2))

}
