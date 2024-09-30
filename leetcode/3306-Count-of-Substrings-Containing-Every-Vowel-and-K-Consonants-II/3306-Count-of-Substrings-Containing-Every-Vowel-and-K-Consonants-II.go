package leetcode

func countOfSubstrings(word string, k int) int64 {
	n := len(word)

	vowels := map[byte]int {'a': 0, 'e': 0, 'i': 0, 'o': 0, 'u': 0}
	countV := 0 // number of vowel types in the window
	countC := 0 // number of consonants in the window

	A := consecutiveConsonants(word, vowels)

	res := 0
	right := 0
	for left := 0; left < n; left++ {
		// 吃
		for right < n && (countV < 5 || countC < k) {
			if _, ok := vowels[word[right]]; ok {
				vowels[word[right]]++
				// increment vowel type ONLY if 0 -> 1
				if vowels[word[right]] == 1 {
					countV++
				}
			} else {
				countC++
			}

			right++
		}

		// update
		// word[left : right-1] (inclusive) minimally satisfied window
		if countV == 5 && countC == k {
			// 判斷right循环是因为right出界还是因为满足了滑窗条件
			if right == n {
				res += 1
			} else if right < n {
				res += 1 + A[right]
			}
		}

		// 吐
		if _, ok := vowels[word[left]]; ok {
			vowels[word[left]]--
			// decrement vowel type ONLY if 1 -> 0
			if vowels[word[left]] == 0 {
				countV--
			}
		} else {
			countC--
		}
	}

	return int64(res)

}

// res[i] := number of consecutive consonants starting at word[i]
func consecutiveConsonants(word string, vowels map[byte]int) []int {
	n := len(word)
	res := make([]int, n)

	c := 0 // cumulative # of consecutive consonants
	for i := n-1; i >= 0; i-- {
		if _, ok := vowels[word[i]]; ok {
			c++
		} else {
			c = 0
		}

		res[i] = c
	}

	return res
}