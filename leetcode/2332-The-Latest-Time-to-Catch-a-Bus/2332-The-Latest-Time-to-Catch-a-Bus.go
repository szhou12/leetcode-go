package leetcode

import "sort"

func latestTimeCatchTheBus(buses []int, passengers []int, capacity int) int {
	sort.Slice(buses, func(i, j int) bool {
		return buses[i] < buses[j]
	})
	sort.Slice(passengers, func(i, j int) bool {
		return passengers[i] < passengers[j]
	})

	res := -1
	n := len(passengers)
	j := 0 // pointer to passenger
	for i := 0; i < len(buses); i++ {
		cap := capacity

		// case 1: 看哪些乘客可以赶上 i-th bus, 挤掉乘客j, 自己上位
		// 注意: 会有乘客 本可以赶上 i-th bus 但没上去的 (bc cap == 0), 他们会被遗留到 i+1-th bus
		for j < n && passengers[j] <= buses[i] && cap > 0 {
			if j == 0 || (j >= 1 && passengers[j]-1 != passengers[j-1]) {
				// 挤掉 j-th passenger, 自己上位
				res = max(res, passengers[j]-1)
			}
			j++
			cap--
		}

		// case 2: 如果 i-th bus 还可以容纳更多人, 紧着bus出发时间来安排我自己的时间
		if cap > 0 {
			if j == 0 || (j >= 1 && passengers[j-1] != buses[i]) {
				res = max(res, buses[i])
			}
		}
	}

	return res
}

// max 可以省略，因为一开始就sort过了, 所以 res一定是单调递增的
func max(a int, b int) int {
	if a > b {
		return a
	}
	return b
}
