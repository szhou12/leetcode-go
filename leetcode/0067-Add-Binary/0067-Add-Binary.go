package leetcode

import "strconv"

func addBinary(a string, b string) string {
	var res string

	l := len(a) - 1
	r := len(b) - 1
	extra := 0

	var curSum int
	for l >= 0 || r >= 0 {
		if l >= 0 && r >= 0 {
			curSum = int((a[l] - '0')) + int((b[r] - '0')) + extra
			l--
			r--
		} else if l >= 0 {
			curSum = int((a[l] - '0')) + extra
			l--
		} else {
			curSum = int((b[r] - '0')) + extra
			r--
		}

		if curSum > 1 {
			extra = 1
		} else {
			extra = 0
		}

		curBin := curSum % 2
		res = strconv.Itoa(curBin) + res
	}

	if extra == 1 {
		return "1" + res
	} else {
		return res
	}
}
