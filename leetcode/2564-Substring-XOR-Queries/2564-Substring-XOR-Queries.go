package leetcode

func substringXorQueries(s string, queries [][]int) [][]int {
	// Step 1: map XOR value to query index
	val2idx := make(map[int][]int) // {XOR value: [query index]}
	for i, q := range queries {
		first, second := q[0], q[1]
		val := first ^ second
		val2idx[val] = append(val2idx[val], i)
	}

	// Step 2: Sliding window - fix size
	res := make([][]int, len(queries))
	for i := 0; i < len(queries); i++ {
		res[i] = []int{-1, -1}
	}
	// loop window size: 1, 2, 3, ..., 31
	for size := 1; size <= 31; size++ {
		val := 0
		for right := 0; right < len(s); right++ {
			// 吃
			val = val*2 + int(s[right]-'0')
			// 吐
			if right >= size {
				val -= int(s[right-size]-'0') * (1 << size)
			}

			// update result if val exists
			if _, ok := val2idx[val]; ok {
				for _, idx := range val2idx[val] {
					// no longer update if already updated
					if res[idx][0] != -1 {
						break
					}
					res[idx] = []int{right - size + 1, right}
				}
			}
		}
	}
	return res
}
