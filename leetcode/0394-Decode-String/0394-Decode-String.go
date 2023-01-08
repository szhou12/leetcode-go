package leetcode

import (
	"bytes"
	"strconv"
)

func decodeString(s string) string {
	dict := make([]byte, 0)
	nums := make([]int, 0)

	for i := 0; i < len(s); i++ {
		char := s[i]

		if isNumber(char) {
			numArr := []byte{char}
			for i < len(s) && isNumber(s[i+1]) {
				numArr = append(numArr, s[i+1])
				i++
			}
			num, _ := strconv.Atoi(string(numArr))
			nums = append(nums, num)
		} else if isLetter(char) {
			dict = append(dict, char)
		} else if char == '[' {
			dict = append(dict, char)
		} else if char == ']' {
			// pop all letters
			letterArr := make([]byte, 0)
			for dict[len(dict)-1] != '[' { // Note: PREPEND elements!!!
				letterArr = append([]byte{dict[len(dict)-1]}, letterArr...)
				dict = dict[:len(dict)-1]
			}
			// dict pop '['
			dict = dict[:len(dict)-1]
			// nums pop top num
			repeats := nums[len(nums)-1]
			nums = nums[:len(nums)-1]
			// copy letterArr n times
			decodeLetterArr := bytes.Repeat(letterArr, repeats)
			// push decodeLetterArr back to dict
			dict = append(dict, decodeLetterArr...)

		}
	}

	return string(dict)
}

func isLetter(char byte) bool {
	return 0 <= char-'a' && char-'a' <= 25
}

func isNumber(char byte) bool {
	return 0 <= char-'0' && char-'0' <= 9
}
