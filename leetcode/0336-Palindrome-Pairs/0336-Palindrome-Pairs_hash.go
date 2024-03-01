package leetcode

import "sort"

// WARNING: Hashing will TLE!
func palindromePairs_hash(words []string) [][]int {
	n := len(words)
	dict := make(map[string]int)
	for i, word := range words {
		dict[word] = i
	}

	// sort by length WHY?
	// Only compare current word with hashed words shorter than it
	// To avoid duplicating pairs
	sort.Slice(words, func(i, j int) bool {
		return len(words[i]) < len(words[j])
	})

	res := make([][]int, 0)
	set := make(map[string]bool)

	for i := 0; i < n; i++ {
		suffix := words[i]
		for k := 0; k <= len(suffix); k++ {
			s1 := suffix[:k]
			s2 := suffix[k:]
			// s2' | s1s2
			match := reverse(s2)
			if _, ok := set[match]; ok && isPalidrome(s1) {
				res = append(res, []int{dict[match], dict[suffix]})

			}

			// s1s2 | s1'
			match = reverse(s1)
			if _, ok := set[match]; ok && isPalidrome(s2) {
				res = append(res, []int{dict[suffix], dict[match]})
			}
		}
		set[suffix] = true
	}
	return res
}

func reverse(word string) string {
	chars := []rune(word)
	for i, j := 0, len(word)-1; i < j; i, j = i+1, j-1 {
		chars[i], chars[j] = chars[j], chars[i]
	}
	return string(chars)
}

func isPalidrome(word string) bool {
	chars := []rune(word)
	for i, j := 0, len(word)-1; i < j; i, j = i+1, j-1 {
		if chars[i] != chars[j] {
			return false
		}
	}
	return true
}
