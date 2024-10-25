package amz

import "sort"

/*
1   2   3
abc def ghi
4   5   6
jkl mno pqr
7   8   9
stu vwx yz

Given a word, we want to re-design the keyboard s.t. the user can type the word with minimum clicks.
e.g. typing "o" takes 3 clicks
*/

// Key Observation: higher frequency characters should be placed before lower frequency characters
// we don't need to exactly what characters are placed, we manipulate by their frequencies
func findMinClicks(word string) int {
	n := len(word)

	freq := make([]int, 26)
	for i := 0; i < n; i++ {
		freq[word[i]-'a']++
	}

	// sort frequencies in descending order
	// e.g. [5, 4, 2, 1, 0, 0, 0 ...]
	sort.Slice(freq, func(i, j int) bool {
		return freq[i] > freq[j]
	})

	clicks := 1 // # of clicks to reach and type a character
	pos := 1 // keyboard number 1-9
	res := 0
	for i := 0; i < 26; i++ {
		if freq[i] == 0 { // no more characters in given word
			break
		}

		res += freq[i] * clicks // it takes 'clicks' to type this character once and we need to repeat 'freq[i]' times
		pos++ // move to next keyboard number

		if pos == 10 {
			// loop keyboard number back to 1
			pos = 1
			clicks++
		}

	}

	return res

}
