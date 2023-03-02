package leetcode

// left[i] := # of distinct chars in s[0:i]
// right[i] := # of distinct chars in s[i:n-1]
func numSplits(s string) int {
	n := len(s)

	left := make([]int, n)
	leftMap := make(map[byte]bool)
	count := 0
	for i := 0; i < n; i++ {
		if _, ok := leftMap[s[i]]; !ok {
			count += 1
			leftMap[s[i]] = true
		}
		left[i] = count
	}

	right := make([]int, n)
	rightMap := make(map[byte]bool)
	count = 0
	for i := n - 1; i >= 0; i-- {
		if _, ok := rightMap[s[i]]; !ok {
			count += 1
			rightMap[s[i]] = true
		}
		right[i] = count
	}

	res := 0
	for i := 0; i+1 < n; i++ {
		if left[i] == right[i+1] {
			res += 1
		}
	}
	return res

}
