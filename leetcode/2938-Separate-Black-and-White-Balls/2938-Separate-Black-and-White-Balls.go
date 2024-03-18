package leetcode

func minimumSteps(s string) int64 {
	count := 0 // # of 0 encountered so far
	res := 0

	for i := len(s) - 1; i >=0; i-- {
		if s[i] == '0' {
			count++
		} else {
			res += count
		}
	}

	return int64(res)
}