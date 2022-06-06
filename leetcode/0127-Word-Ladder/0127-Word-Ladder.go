package leetcode

func ladderLength(beginWord string, endWord string, wordList []string) int {
	depth := 0
	queue := []string{beginWord}
	wordMap := getWordMap(beginWord, wordList)

	for len(queue) > 0 {
		qlen := len(queue)
		for i := 0; i < qlen; i++ {
			word := queue[0]
			queue = queue[1:]
			candidates := getCandidates(word, wordMap)
			for _, candidate := range candidates {
				if _, ok := wordMap[candidate]; ok {
					if candidate == endWord {
						return depth + 1
					}
					delete(wordMap, candidate) // MUST have this (相当于 mark as visited), otherwise 会超时
					queue = append(queue, candidate)
				}
			}
		}
		depth += 1
	}
	// return 0: if wordList has no endWord
	return 0

}

func getWordMap(beginWord string, wordList []string) map[string]bool {
	wordMap := make(map[string]bool)

	for _, word := range wordList {
		wordMap[word] = true
	}
	return wordMap

}

func getCandidates(word string, wordMap map[string]bool) []string {
	var res []string
	for i := 0; i < len(word); i++ {
		for j := 0; j < 26; j++ {
			if word[i] != byte(int('a')+j) {
				candidate := word[:i] + string(rune(int('a')+j)) + word[i+1:]
				if _, ok := wordMap[candidate]; ok {
					res = append(res, candidate)
				}
			}
		}
	}
	return res
}
