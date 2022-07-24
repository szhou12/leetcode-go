package leetcode

func strStr(haystack string, needle string) int {
	// edge case 1: needle is empty string
	if len(needle) == 0 {
		return 0
	}
	// edge case 2: needle is longer than haystack
	if len(haystack) < len(needle) {
		return -1
	}

	r := 0

	for l := 0; l < len(haystack); l++ {
		if haystack[l] != needle[r] {
			r = 0
		} else {
			r++
			if r == len(needle) {
				return l - len(needle) + 1
			}
		}
	}

	return -1
}

//    0 1 2 3 4
// h: h e l l o
//          l
//    0 1
// n: l l
//        r
