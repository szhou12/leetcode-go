package leetcode

import (
	"math"
	"sort"
)

// 模版Flex
func minOperations(nums []int) int {
	// deduplication
	set := make(map[int]int)
	for _, val := range nums {
		set[val]++
	}
	arr := make([]int, 0)
	for k, _ := range set {
		arr = append(arr, k)
	}
	// sort in ascending order
	sort.Ints(arr)

	N := len(nums)
	right := 0
	res := math.MaxInt
	for left := 0; left < len(arr); left++ {
		for right < len(arr) && arr[right]-arr[left]+1 <= N {
			// update result: 在loop里更新result, 要+1
			// res = min(res, N-(right-left+1))

			right++
		}
		// update result: 在loop外更新result, 不用+1
		// 因为跳出loop时, [left, right) 是在当前左边界下，能到达的最远右边界的下一个
		res = min(res, N-(right-left))
	}

	return res

}

func min(a int, b int) int {
	if a < b {
		return a
	}
	return b
}
