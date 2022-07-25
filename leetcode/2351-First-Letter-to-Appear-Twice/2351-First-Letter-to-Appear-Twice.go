package leetcode

func repeatedCharacter(s string) byte {
	dict := make(map[rune]int)
	for _, char := range s {
		if _, ok := dict[char]; ok {
			if dict[char] == 1 {
				return byte(char)
			}

		} else {
			dict[char] = 1
		}
	}

	return 0
}
