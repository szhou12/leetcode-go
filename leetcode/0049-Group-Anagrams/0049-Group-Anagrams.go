package leetcode

import "sort"

// method 1: sort every word, use map to group words that are same after sorted
// {sorted word: [w1, w2, ...]}
func groupAnagrams(strs []string) [][]string {
	dict := make(map[string][]string)

	for _, word := range strs {
		sorted := sortWord(word)
		if _, ok := dict[sorted]; !ok {
			dict[sorted] = []string{word}
		} else {
			dict[sorted] = append(dict[sorted], word)
		}
	}

	res := make([][]string, 0)
	for _, group := range dict {
		res = append(res, group)
	}

	return res
}

func sortWord(word string) string {
	chars := []rune(word)
	sort.Slice(chars, func(i, j int) bool {
		return chars[i] < chars[j]
	})

	return string(chars)
}
