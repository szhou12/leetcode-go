package leetcode

func digitCount(num string) bool {
	dict := make([]int, len(num))

	// char is 'rune' type
	for _, char := range num {
		temp := int(char - 48) // ASCII for 0 = 48
		if temp < len(num) {
			dict[temp]++
		}
	}
	for i, char := range num {
		if dict[i] != int(char-48) {
			return false
		}
	}
	return true
}
