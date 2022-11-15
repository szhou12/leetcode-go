package leetcode

// My Solution: 每个edge都有的 vertex 就是要找的
func findCenter(edges [][]int) int {
	counter := make(map[int]int)

	for _, edge := range edges {
		left, right := edge[0], edge[1]
		if _, ok := counter[left]; ok {
			counter[left]++
		} else {
			counter[left] = 1
		}
		if _, ok := counter[right]; ok {
			counter[right]++
		} else {
			counter[right] = 1
		}
	}

	n := len(edges)
	for k, v := range counter {
		if v == n {
			return k
		}
	}

	return -1
}

// Solution: 计算每个vertex的入度; 找到vertex的度数==边总数
func findCenter_degree(edges [][]int) int {
	deg := make(map[int]int)

	for _, edge := range edges {
		deg[edge[0]]++
		deg[edge[1]]++
	}

	for k, v := range deg {
		if v == len(edges) {
			return k
		}
	}
	return -1
}
