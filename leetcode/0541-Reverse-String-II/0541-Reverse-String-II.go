package leetcode

// My Solution
// 难点：主要是 2*k 一跳，每一个 2*k 区间 反转前k个, 要找的也就是每 2*k 区间的起点
func reverseStr(s string, k int) string {
	chars := []byte(s)
	n := len(chars)
	right := n - 1
	left := 0

	for left < n {
		numChars := right - left + 1
		if numChars < k {
			reverse(&chars, left, right)
		} else if numChars >= k {
			reverse(&chars, left, left+k-1)
		}
		left += 2 * k
	}

	return string(chars)
}

func reverse(chars *[]byte, start int, end int) {
	l := start
	r := end

	for l <= r {
		(*chars)[l], (*chars)[r] = (*chars)[r], (*chars)[l]
		l++
		r--
	}
}
