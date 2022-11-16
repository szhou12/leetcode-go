package leetcode

func uniqueOccurrences(arr []int) bool {
	count := make(map[int]int)
	for _, val := range arr {
		if _, ok := count[val]; !ok {
			count[val] = 1
		} else {
			count[val]++
		}
	}
	dups := make(map[int]bool)
	for _, c := range count {
		if _, ok := dups[c]; !ok {
			dups[c] = true
		} else {
			return false
		}
	}
	return true
}
