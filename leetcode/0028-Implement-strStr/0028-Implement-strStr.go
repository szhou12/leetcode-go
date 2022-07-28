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
	// edge case 3: haystack is needle (i.e. same length & same chars)
	if haystack == needle {
		return 0
	}

	// sliding window
	l := 0
	r := len(needle)
	for r <= len(haystack) { // Note: r can be len(haystack)
		if haystack[l:r] == needle {
			return l
		}
		l++
		r++
	}

	return -1
}
