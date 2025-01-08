package leetcode

import "sort"

func maximumCoins(coins [][]int, k int) int64 {
	res := 0

	// sort intervals by left-ends in ascending order
	sort.Slice(coins, func(i, j int) bool {
		return coins[i][0] < coins[j][0]
	})

	res = max(res, slideWindow(coins, k))

	// negate coordinates
	for i, coin := range coins {
		// 取负数前一定要先把值取出来，暂存在a,b中，否则直接在coin[0], coin[1]上取负数会出错
		a, b := coin[0], coin[1]
		coins[i][0] = -b
		coins[i][1] = -a
	}

	// sort intervals by right-ends in ascending order
	sort.Slice(coins, func(i, j int) bool {
		return coins[i][0] < coins[j][0]
	})

	res = max(res, slideWindow(coins, k))

	return int64(res)
}

func slideWindow(coins [][]int, k int) int {
	n := len(coins)

	r := 0 // r-th inerval can be reached. 代表第几个区间，而不是某一个区间的右端点

	completeIntervalSum := 0 // 只累计完整覆盖的区间的sum
	res := 0

	for l := 0; l < n; l++ {
		end := coins[l][0] + k - 1
		// 吃
		for r < n && coins[r][1] <= end {
			completeIntervalSum += (coins[r][1] - coins[r][0] + 1) * coins[r][2]
			r++
		}
		// leftover partial interval to be covered
		partial := 0 // 部分覆盖的区间的sum另算，不能直接累加进completeIntervalSum，会出错
		if r < n && coins[r][0] <= end {
			partial = (end - coins[r][0] + 1) * coins[r][2]
		}

		// update result
		res = max(res, completeIntervalSum + partial)

		// 吐
		// 注意：这里的写法是正确的
		// 虽然在interval长度超过k时，当前轮不会吃掉一个完整覆盖的interval，而吐的时候吐掉了一个完整覆盖的interval，导致completeIntervalSum为负数
		// 但是这个负数会在下一轮的吃操作中马上得到修复。因为r在当前轮也相应地没有累加，所以在下一轮中会首先完整吃掉这个当前轮没有完整覆盖的interval，使completeIntervalSum重新归0
		// example: coins := [][]int{{1,4,4}, {31,39,20}, {41,42,100}}, k := 3
		completeIntervalSum -= (coins[l][1] - coins[l][0] + 1) * coins[l][2]
	}

	return res
}
