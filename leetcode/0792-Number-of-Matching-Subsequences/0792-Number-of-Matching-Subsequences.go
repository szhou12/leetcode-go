package leetcode

// Optimal Solution - Binary Search
// Use 0392 - binary search as helper function
func numMatchingSubseq(s string, words []string) int {
	// Step 1: pre-process s
	dict := make([][]int, 26)
	for index, char := range s {
		dict[int(char-'a')] = append(dict[int(char-'a')], index)
	}

	res := 0
	// Step 2: binary search for each word in words
	for _, word := range words {
		sPointer := 0 // Pointer for s, the target string

		// Pointer for current word: it's used to tell whether we can
		// successfully iter all chars in word. If so, we increment res.
		// If break at any point in between, word is NOT a subseq.
		wordPointer := 0
		for ; wordPointer < len(word); wordPointer++ {
			char := word[wordPointer]

			// Case 1: this char doesn't even exist in dict for s
			if len(dict[int(char-'a')]) == 0 {
				break
			}

			// Case 2: binary search
			pos := leftBound(dict[int(char-'a')], sPointer)
			if pos == -1 {
				break
			}
			sPointer = dict[int(char-'a')][pos] + 1
		}

		// word is a subseq ONLY IF we successfully iter all chars in word
		if wordPointer == len(word) {
			res++
		}
	}

	return res
}

func leftBound(array []int, target int) int {
	left := 0
	right := len(array) - 1

	for left < right {
		mid := left + (right-left)/2
		if array[mid] >= target {
			right = mid
		} else {
			left = mid + 1
		}
	}

	if array[left] < target {
		return -1
	}

	return left
}

// Solution - Brute force
// Use 0392 - greedy algo as helper function
func numMatchingSubseq_BF(s string, words []string) int {
	count := 0
	for _, t := range words {
		if isSubseq_Greedy(s, t) {
			count++
		}
	}

	return count
}

// check if t is subseq of s
func isSubseq_Greedy(s string, t string) bool {
	// edge cases
	// if s empty, t empty => true
	// if s empty, t not => false
	// if s not, t empty => true
	if len(t) == 0 {
		return true
	}
	if len(s) == 0 {
		return false
	}

	// pointer for t
	left := 0
	for right := 0; right < len(s); right++ {
		if len(t) == left {
			return true
		}
		if t[left] == s[right] {
			left++
		}
	}

	return len(t) == left
}
