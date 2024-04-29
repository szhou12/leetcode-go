package leetcode

func minEnd(n int, x int) int64 {
	m := n - 1             // last element
	last := make([]int, 0) // low -> high
	for m > 0 {
		last = append(last, m%2)
		m /= 2
	}

	bits := make([]int, 0) // low -> high
	for x > 0 {
		bits = append(bits, x%2)
		x /= 2
	}

	// fill bits[] with last[] on positions where bits[i]=0
	j := 0
	for i := 0; i < len(bits); i++ {
		if bits[i] == 1 {
			continue
		}
		if j < len(last) {
			bits[i] = last[j]
			j++
		}
	}
	// it's possible last[] is longer than bits[], continue appending rest of bits of last[]
	for j < len(last) {
		bits = append(bits, last[j])
		j++
	}

	res := 0
	for i := len(bits) - 1; i >= 0; i-- {
		res = res*2 + bits[i]
	}
	return int64(res)
}
