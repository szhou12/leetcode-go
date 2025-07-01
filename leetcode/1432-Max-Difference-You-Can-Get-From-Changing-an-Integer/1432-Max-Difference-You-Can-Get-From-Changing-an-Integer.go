package leetcode

import (
	"strconv"
	"strings"
)

func maxDiff(num int) int {
	// replace x with y in a number
	var change func(x, y int) string
	change = func(x, y int) string {
		numStr := strconv.Itoa(num)
		var res strings.Builder
		for _, digit := range numStr {
			if int(digit-'0') == x {
				res.WriteByte(byte('0' + y))
			} else {
				res.WriteRune(digit)
			}
		}
		return res.String()
	}

	minNum := num
	maxNum := num
	// try out all combination: x has 10 options, y has 10 options (0-9)
	for x := 0; x < 10; x++ {
		for y := 0; y < 10; y++ {
			cur := change(x, y)
			// if leading zero, invalid
			if cur[0] == '0' {
				continue
			}
			curNum, _ := strconv.Atoi(cur)
			minNum = min(minNum, curNum)
			maxNum = max(maxNum, curNum)
		}
	}

	return maxNum - minNum
}
