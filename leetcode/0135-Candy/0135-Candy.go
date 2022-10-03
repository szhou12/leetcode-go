package leetcode

func candy(ratings []int) int {
	n := len(ratings)
	candies := make([]int, n)

	// init: 每个人至少一个糖果
	for i := 0; i < n; i++ {
		candies[i] = 1
	}

	// Step 1: From left to right: 当前人比前一个人高，则要比前一个人多一个糖
	for i := 1; i < n; i++ {
		if ratings[i] > ratings[i-1] {
			candies[i] = candies[i-1] + 1
		}
	}

	// Step 2: From right to left: 当前人比后一个人高，则要比后一个人多一个糖
	// Step 3: max(从左遍历来的糖，从右遍历来的糖) - 求最大值才能保证当前糖数 同时 高于左右邻居
	for i := n - 2; i >= 0; i++ {
		if ratings[i] > ratings[i+1] {
			candies[i] = max(candies[i], candies[i+1]+1)
		}
	}

	// Step 4: sum all candies
	res := 0
	for i := 0; i < n; i++ {
		res += candies[i]
	}
	return res

}

func max(a int, b int) int {
	if a > b {
		return a
	}
	return b
}
