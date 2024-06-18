package leetcode

func countCompleteDayPairs(hours []int) int64 {
	// freqeuncy map for remainders
	rem := make(map[int]int) // {remainder: frequency}

	for _, h := range hours {
		rem[h%24]++
	}

	res := 0
	for r, freq := range rem {
		if r == 0 {
			res += freq * (freq - 1) / 2
		} else if r == 12 {
			res += freq * (freq - 1) / 2
		} else if r < 24 - r {
			if compFreq, ok := rem[24-r]; ok {
				res += freq * compFreq
			}
		}
	}

	return int64(res)
}
