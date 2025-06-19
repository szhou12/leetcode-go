package leetcode

func isZeroArray(nums []int, queries [][]int) bool {
	n := len(nums)
	diff := make([]int, n+1)
	for i := 0; i < n; i++ {
		if i == 0 {
			diff[i] = nums[i]
		} else {
            diff[i] = nums[i] - nums[i-1]
        }
		
	}
	// range decrement by 1
	for _, q := range queries {
		l, r := q[0], q[1]
		diff[l] -= 1
		diff[r+1] += 1
	}
	// prefix sum of diff[]
	presum := make([]int, n+1)
	for i := 0; i < len(diff); i++ {
		if i == 0 {
			presum[i] = diff[i]
		} else {
            presum[i] = presum[i-1] + diff[i]
        }
		
	}

	// every element from 0 to n-1 should be <= 0
	// otherwise, we cannot zero out the array
	for i := 0; i < n; i++ {
		if presum[i] > 0 {
			return false
		}
	}
	return true
}