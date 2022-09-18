package leetcode

func longestPalindrome(s string) int {
	dict := make(map[string]int)

	res := 0
	for i := 0; i < len(s); i++ {
		char := s[i : i+1]
		if _, ok := dict[char]; ok { // if map contains this char, it means we've found a pair
			res += 2
			delete(dict, char) // remove this char from map
		} else { // otherwise, add this char to map
			dict[char] = 1
		}
	}

	if len(dict) > 0 { // if len of map > 0, it means we can add one more char at center of palindrome
		res += 1
	}

	return res
}
