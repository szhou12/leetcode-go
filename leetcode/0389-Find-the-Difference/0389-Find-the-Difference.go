package leetcode

// My solution - HashMap
func findTheDifference_map(s string, t string) byte {
	dict := make(map[byte]int)

	for i := 0; i < len(s); i++ {
		dict[s[i]]++
	}

	for j := 0; j < len(t); j++ {
		if _, ok := dict[t[j]]; !ok {
			return t[j]
		}
		if dict[t[j]] == 0 {
			return t[j]
		}

		dict[t[j]]--
	}

	return byte(' ')
}

// Use array of length 26 to replace map
func findTheDifference_array(s string, t string) byte {
	count := make([]int, 26)

	for i := 0; i < len(s); i++ {
		count[int(s[i]-'a')]++
	}

	for j := 0; j < len(t); j++ {
		count[int(t[j]-'a')]--
		if count[int(t[j]-'a')] < 0 {
			return t[j]
		}
	}

	return byte(' ')
}

// Use XOR
func findTheDifference(s string, t string) byte {
	n := len(t)
	char := t[n-1] // t的最后一个作为起始点，因为t比s多一个char

	for i := 0; i < n-1; i++ {
		char ^= s[i]
		char ^= t[i]
	}
	return char
}
