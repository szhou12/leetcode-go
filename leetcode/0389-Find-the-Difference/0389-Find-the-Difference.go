package leetcode

// My solution - HashMap
func findTheDifference_map(s string, t string) byte {
	dict := make(map[byte]int)

	for i := 0; i < len(s); i++ {
		dict[s[i]]++
	}

	var res byte
	for j := 0; j < len(t); j++ {
		if _, ok := dict[t[j]]; !ok {
			return t[j]
		}
		if dict[t[j]] == 0 {
			return t[j]
		}

		dict[t[j]]--
	}

	return res
}
