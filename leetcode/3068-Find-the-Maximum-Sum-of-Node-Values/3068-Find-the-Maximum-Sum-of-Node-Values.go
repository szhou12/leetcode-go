package leetcode

import "sort"

func maximumValueSum(nums []int, k int, edges [][]int) int64 {
	diff := make([]int, 0) // 增量array: 表示每个节点 XOR k 一次之后可以增加多少
	for _, x := range nums {
		diff = append(diff, (x^k)-x)
	}

	// sort diff in descending order
	sort.Slice(diff, func(i, j int) bool {
		return diff[i] > diff[j]
	})

	// sequentially select a pair of nodes at a time
	maxDiffSum := 0
	for i := 0; i+1 < len(nums); i+=2 {
		maxDiffSum = max(maxDiffSum, maxDiffSum+diff[i]+diff[i+1])
	}

	nodesSum := 0
	for _, x := range nums {
		nodesSum += x
	}

	return int64(nodesSum + maxDiffSum)
}
