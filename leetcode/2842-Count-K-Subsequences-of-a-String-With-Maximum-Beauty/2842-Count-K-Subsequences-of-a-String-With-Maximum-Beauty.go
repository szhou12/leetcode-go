package leetcode

import "sort"

var M = int(1e9 + 7)

func countKSubsequencesWithMaxBeauty(s string, k int) int {
	f := make(map[byte]int)
	for i := 0; i < len(s); i++ { // O(n)
		f[s[i]]++
	}
	counter := make([]int, 0)
	for _, v := range f { // O(26)
		counter = append(counter, v)
	}
	// early return: if total number of unique characters < k, return 0
	if len(counter) < k {
		return 0
	}

	// sort in decreasing order to get Top K
	sort.Slice(counter, func(i, j int) bool { // O(26log26)
		return counter[i] > counter[j]
	})
	// target beauty = sum(top k)
	targetBeauty := 0
	for i := 0; i < k; i++ { // O(k)
		targetBeauty += counter[i]
	}

	// DFS
	// total number of subsequences that satisfy the condition
	total := 0

	// curChoices: current numeber of subsequences that satisfy the condition
	// at start, curChoinces = 1 cause choosing nothing is one option
	dfs(k, 0, 0, 0, targetBeauty, counter, 1, &total) // O(2^k) where k <= 26

	return total

}

func dfs(k int, curIndex int, picked int, curBeauty int, targetBeauty int, counter []int, curChoices int, total *int) {
	// base cases
	if picked == k && curBeauty == targetBeauty {
		*total += curChoices
		*total %= M
		return
	}

	if curBeauty > targetBeauty {
		return
	}

	if picked > k {
		return
	}

	// pruning
	if curBeauty + sum(counter[curIndex:]) < targetBeauty {
		return
	}

	// recursion
	for i := curIndex; i < len(counter); i++ {
		// NOTE: (curChoices*counter[i]) MUST mod M !
		dfs(k, i+1, picked+1, curBeauty+counter[i], targetBeauty, counter, curChoices*counter[i]%M, total)
	}
}

func sum(arr []int) int {
	res := 0
	for _, count := range arr {
		res += count
	}
	return res
}