package leetcode

func minimumOperationsToMakeKPeriodic(word string, k int) int {
	n := len(word)
	count := make(map[string]int)
	for i := 0; i < n; i+=k {
		count[word[i:i+k]]++
	}

	maxFreq := -1
	for _, freq := range count {
		maxFreq = max(maxFreq, freq)
	}

	return n/k-maxFreq
}