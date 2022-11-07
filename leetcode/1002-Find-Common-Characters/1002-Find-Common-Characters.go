package leetcode

import "math"

// My Solution
// 用一个map[key=char, value = [counts0, counts1, ...]] 记录每个出现在第一个word里的字母分别在每个word中出现的次数
// 每个字母对应的count数组取最小值，就是所有单词的common个数
func commonChars(words []string) []string {
	dict := make(map[rune][]int)
	for _, char := range words[0] { // O(m)
		if _, ok := dict[char]; !ok {
			dict[char] = make([]int, len(words))
			dict[char][0] = 1
		} else {
			dict[char][0]++
		}
	}

	for i := 1; i < len(words); i++ { // O(n)
		for _, char := range words[i] { // O(m)
			if _, ok := dict[char]; ok {
				dict[char][i]++
			}
		}
	}

	res := make([]string, 0)
	for char, counts := range dict { // O(m)
		minCount := findMin(counts) // O(n)
		for minCount > 0 {          // O(m)
			res = append(res, string(char))
			minCount--
		}
	}
	return res
}

func findMin(nums []int) int {
	res := math.MaxInt
	for _, val := range nums {
		res = min(res, val)
	}

	return res
}

func min(a int, b int) int {
	if a < b {
		return a
	}
	return b
}
