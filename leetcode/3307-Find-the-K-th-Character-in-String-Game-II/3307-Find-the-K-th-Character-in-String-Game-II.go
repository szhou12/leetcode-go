package leetcode

func kthCharacter(k int64, operations []int) byte {
	var n int64 = 1 // length of the word
	times := 0 // # of times to recursively double the size of word until its length >= k
	for n < k {
		n *= 2
		times++
	}

	// recursion
	shift := 0 // number of times to shift letter by operation 1
	for i := times - 1; i >= 0; i-- {
		// k is in the first half or the second half of n?
		if k > n / 2 { // if k in the second half
			if operations[i] == 0 {
				k = k - n/2
			} else {
				k = k - n/2
				shift++
			}
		}
		// if k in the first half, no need to change k

		n /= 2
	}

	return byte('a' + (shift % 26))
}