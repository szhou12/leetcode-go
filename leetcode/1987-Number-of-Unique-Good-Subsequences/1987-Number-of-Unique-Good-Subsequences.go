package leetcode

var M = int(1e9 + 7)

func numberOfUniqueGoodSubsequences(binary string) int {
	// DP def
	one := 0  // # of unique good subseq ending with 1 for s[0:i]
	zero := 0 // # of unique good subseq ending with 0 for s[0:i]

	hasZero := 0 // bool value: whether s has any 0, used for corner case subseq: '0'

	// Recurrence
	for _, char := range binary {
		if char == '0' {
			zero = (one + zero) % M
			hasZero = 1
		} else {
			one = (one + zero + 1) % M
		}
	}

	return (one + zero + hasZero) % M
}
