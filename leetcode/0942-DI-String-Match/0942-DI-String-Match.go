package leetcode

func diStringMatch(s string) []int {
	s = "I" + s // prepend I
	n := len(s)
	max := -1
	var res []int

	for l := 0; l < n; l++ {
		// Step 1: move right bound
		r := l + 1
		for r < n && s[r] == 'D' { // stop at l == n || s[r] = next 'I'
			r++
		}
		// Step 2: append 'IDDD..D'
		count := r - l
		for k := max + count; k >= max+1; k-- {
			res = append(res, k) // append in decreasing order
		}
		// Step 3: prepare next round: update left bound and max
		max = max + count
		l = r - 1 // r-1 bc next round will l++, making it l = r
	}
	return res
}
