package leetcode

func judgeCircle(moves string) bool {
	status := []int{0, 0}
	for _, move := range moves {
		if move == 'U' {
			status[0]++
		} else if move == 'D' {
			status[0]--
		} else if move == 'L' {
			status[1]++
		} else { // 'R'
			status[1]--
		}
	}

	return status[0] == 0 && status[1] == 0
}
