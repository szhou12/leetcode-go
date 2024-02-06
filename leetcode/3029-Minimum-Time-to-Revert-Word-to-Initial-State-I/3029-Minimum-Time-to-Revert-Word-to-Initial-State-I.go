package leetcode

func minimumTimeToInitialState(word string, k int) int {
	n := len(word)
	t := 1 // 切几刀k = multiple of k
	for t*k < n && word[:n-t*k] != word[t*k:] {
		t++
	}
	return t
}
