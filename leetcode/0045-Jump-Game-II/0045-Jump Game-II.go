package leetcode

import "math"


// Greedy: 想象成BFS, 每一层=可以跳到的一段连续区间
func jump(nums []int) int {
    start, end := 0, 0

    if len(nums) == 1 {
        return 0
    }

    step := 0

    for start <= end {
        newEnd := end
        for i := start; i <= end; i++ {
            newEnd = max(newEnd, i+nums[i])
            if newEnd >= len(nums)-1 {
                return step + 1
            }
        }
        start = end + 1
        end = newEnd
        step++
    }

    return -1
}

func jump_DP(nums []int) int {
	n := len(nums)
	dp := make([]int, n)
	// base case: at starting position, need 0 jumps
	dp[0] = 0

	for i := 1; i < n; i++ {
		dp[i] = math.MaxInt
	}

	for i := 1; i < n; i++ {
		for j := 0; j < i; j++ {
			if nums[j] >= i-j {
				dp[i] = min(dp[i], dp[j]+1)
			}
		}
	}

	return dp[n-1]
}


// testing: [2,3,1,1,4]
// dp: [0, max, max, max, max]
// i: 1 j: 0
// dp: [0, 1, max, max, max]
// i: 2 j: 0, 1
// dp: [0, 1, min(2,1)=1, max, max]
// i: 3 j: 0, 1, 2
// dp: [0, 1, 1, min(2,2)=2, max]
// i: 4 j: 0, 1, 2, 3
// dp: [0, 1, 1, 2, min(2, 3)=2]
