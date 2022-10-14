package leetcode

import "sort"

// Greedy Algorithm
func reconstructQueue(people [][]int) [][]int {
	// Step 1: sort by height in decreasing order
	sort.Slice(people, func(i, j int) bool {
		if people[i][0] == people[j][0] { // 如果身高相同，把需要排前面的人数少的优先
			return people[i][1] < people[j][1]
		}
		return people[i][0] > people[j][0]
	})

	// Step 2: greedily insert each person
	var res [][]int
	for _, p := range people {
		res = append(res, p)
		copy(res[p[1]+1:], res[p[1]:]) // 将插入位置之后的元素后移动一位（i.e.腾出空间）
		res[p[1]] = p                  // 将插入元素位置插入元素
	}
	return res
}
