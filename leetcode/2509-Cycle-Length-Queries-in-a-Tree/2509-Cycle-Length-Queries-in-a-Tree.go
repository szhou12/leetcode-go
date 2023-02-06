package leetcode

func cycleLengthQueries(n int, queries [][]int) []int {
	var res []int
	for _, query := range queries {
		count := 0
		a, b := query[0], query[1]
		for a != b {
			if a > b {
				a = a / 2
			} else {
				b = b / 2
			}
			count++
		}

		res = append(res, count+1)
	}

	return res
}
