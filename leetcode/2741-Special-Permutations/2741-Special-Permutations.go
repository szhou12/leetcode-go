package leetcode

var M = int(1e9+7)


// dfs(i, p, state) := # of permutations (counts) if we have set i-th position
// 					with p-th number and we have selected numbers from state being 1
func specialPerm(nums []int) int {
	n := len(nums)

	memo := make([][]int, n)
	for p := 0; p < n; p++ {
		memo[p] = make([]int, 1<<n)
	}
	for p := 0; p < n; p++ {
		for s := 0; s < (1<<n); s++ {
			memo[p][s] = -1
		}
	}

	res := 0
	for p := 0; p < n; p++ {
		res += dfs(0, p, 1<<p, nums, memo)
		res %= M
	}

	return res
}

func dfs(i int, p int, state int, nums []int, memo [][]int) int {
	n := len(nums)

	// base case
	if i == n -1 {
		return 1
	}

	// pruning
	if memo[p][state] != -1 {
		return memo[p][state]
	}

	count := 0
	// set i+1-th position with q-th number
	for q := 0; q < n; q++ {
		if (state>>q) & 1 == 1 {
			continue
		}
		if nums[p] % nums[q] != 0 && nums[q] % nums[p] != 0 {
			continue
		}
		count += dfs(i+1, q, state+(1<<q), nums, memo)
		count %= M
	}

	memo[p][state] = count

	return count
}