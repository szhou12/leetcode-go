package leetcode

// [a, b]
// last[a] = index of nearest candle on the right of a (or itself if a is candle)
// next[b] = index of nearest candle on the left of b (or itself if b is candle)
// [a.....|..........|.......b]
//      next[a]....last[b]
// presum[i] = number of plates from 0 to i
func platesBetweenCandles(s string, queries [][]int) []int {
	n := len(s)

	presum := make([]int, n)
	plates := 0
	for i := 0; i < n; i++ {
		if s[i] == '*' {
			plates++
		}
		presum[i] = plates
	}

	next := make([]int, n)
	closestIndex := n // no candle on the right exists
	for i := n-1; i >=0; i-- {
		if s[i] == '|' { // update closest right candle each time encounter a new candle
			closestIndex = i
		}
		next[i] = closestIndex
	}

	last := make([]int, n)
	closestIndex = -1 // no candle on the left exists
	for i := 0; i < n; i++ {
		if s[i] == '|' {
			closestIndex = i // update closest left candle each time encounter a new candle
		}
		last[i] = closestIndex
	}

	res := make([]int, 0)
	for _, q := range queries {
		a, b := q[0], q[1]
		// Check if a valid substring
		if next[a] <= last[b] {
			res = append(res, presum[last[b]] - presum[next[a]])
		} else {
			// no valid plates between candles
			res = append(res, 0)
		}
	}
	return res
}