package leetcode

import "strconv"

func letterCombinations(digits string) []string {
	// edge case
	if len(digits) == 0 {
		return []string{}
	}

	dict := []string{"abc", "def", "ghi", "jkl", "mno", "pqrs", "tuv", "wxyz"}

	res := make([]string, 0)
	cur := ""

	dfs(digits, cur, 0, &res, dict)
	return res

}

func dfs(digits string, cur string, index int, res *[]string, dict []string) {
	// base case
	if index == len(digits) {
		*res = append(*res, cur)
		return
	}

	curDigit, _ := strconv.Atoi(digits[index : index+1])
	letters := dict[curDigit-2]
	for i := 0; i < len(letters); i++ {
		char := letters[i : i+1]
		cur += char
		dfs(digits, cur, index+1, res, dict)
		cur = cur[:len(cur)-1]
	}
}

// # levels = # of digits
// # branches/level = chars each digit mapped to
