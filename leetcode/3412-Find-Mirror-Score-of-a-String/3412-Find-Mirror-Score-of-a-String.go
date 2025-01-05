package leetcode

func calculateScore(s string) int64 {
	dict := make(map[int][]int)

	res := 0

	for i, letter := range s {
		cur := int(letter - 'a')
		mirror := 25 - cur

		// if mirror exists in dict, use the last element and remove it after use
		// otherwise, put cur into dict as a candidate
		if _, ok := dict[mirror]; ok && len(dict[mirror]) > 0 {
			n := len(dict[mirror])
			res += i - dict[mirror][n-1]
			dict[mirror] = dict[mirror][:n-1]
		} else {
			if _, ok := dict[cur]; !ok {
				dict[cur] = make([]int, 0)
			}
			dict[cur] = append(dict[cur], i)
		}
	}

	return int64(res)
}