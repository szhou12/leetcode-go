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

// Demo:
// 0 1 2 3
// I D I D
// l   r
//     l   r
// max = -1
// count = 2; i = -1+2=1
// res = [1, 0]
// max = -1+2=1
