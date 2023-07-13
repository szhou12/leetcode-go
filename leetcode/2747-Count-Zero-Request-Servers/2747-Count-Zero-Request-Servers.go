package leetcode

import "sort"

func countServers(n int, logs [][]int, x int, queries []int) []int {
	// sort logs in ascending order
	sort.Slice(logs, func(i, j int) bool {
		return logs[i][1] < logs[j][1]
	})
	// add idx to queries and sort in ascending order
	queriesId := make([][]int, len(queries))
	for i, q := range queries {
		queriesId[i] = []int{i, q}
	}
	sort.Slice(queriesId, func(i, j int) bool {
		return queriesId[i][1] < queriesId[j][1]
	})

	// sliding window
	res := make([]int, len(queries))
	activeServers := make(map[int]int) // {serverId: count}

	left, right := 0, 0
	for _, q := range queriesId {
		qid := q[0]
		lower, upper := q[1]-x, q[1]

		for right < len(logs) && logs[right][1] <= upper {
			activeServers[logs[right][0]]++
			right++
		}

		for left < len(logs) && logs[left][1] < lower {
			activeServers[logs[left][0]]--
			if activeServers[logs[left][0]] == 0 {
				delete(activeServers, logs[left][0])
			}
			left++
		}

		res[qid] = n - len(activeServers)

	}

	return res

}
