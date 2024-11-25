package leetcode

func uniqueLetterString(s string) int {
	n := len(s)

	positions := make([][]int, 26)
	for ch := 0; ch < 26; ch++ {
		positions[ch] = make([]int, 0)
	}

	// -1 at head
	for ch := 0; ch < 26; ch++ {
		positions[ch] = append(positions[ch], -1)
	}

	// append corresponding char's index
	for i := 0; i < n; i++ {
		positions[int(s[i]-'A')] = append(positions[int(s[i]-'A')], i)
	}

	// n at tail
	for ch := 0; ch < 26; ch++ {
		positions[ch] = append(positions[ch], n)
	}

	res := 0
	for ch := 0; ch < 26; ch++ {
		for i := 1; i < len(positions[ch]) - 1; i++ {
			res += (positions[ch][i] - positions[ch][i-1]) * (positions[ch][i+1] - positions[ch][i])
		}
	}

	return res

}

// [[-1, 0, 3, 6, n], 
//  [-1, 1, 4, 7, n], 
//  [-1, 2, 5, 8, n]]