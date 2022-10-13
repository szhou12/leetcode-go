package leetcode

func canConstruct(ransomNote string, magazine string) bool {
	dict := make(map[rune]int)
	for _, char := range magazine {
		dict[char]++
	}

	for _, char := range ransomNote {
		if count, ok := dict[char]; ok && count > 0 {
			dict[char]--
		} else {
			return false
		}
	}

	return true
}
