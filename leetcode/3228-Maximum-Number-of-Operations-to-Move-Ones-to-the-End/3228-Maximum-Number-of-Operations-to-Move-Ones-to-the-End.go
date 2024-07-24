package leetcode

func maxOperations(s string) int {
	n := len(s)

	cnt := 0 // # of 1 as of s[0:i] (inclusive)
	res  := 0 // max # of moves

	for i := 0; i < n; i++ {
		if s[i] == '0' {
			res += cnt
			for i < n && s[i] == '0' { // skip all consecutive 0
				i++
			}
		}
		cnt++
	}

	return res
}