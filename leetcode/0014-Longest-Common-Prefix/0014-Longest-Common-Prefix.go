package leetcode

import "sort"

func longestCommonPrefix(strs []string) string {
	// step 1: sort all words by length in ascending order
	sort.Slice(strs, func(i, j int) bool {
		return len(strs[i]) < len(strs[j])
	})

	// step 2: compare each char
	prefix := strs[0]
	for i := 1; i < len(strs); i++ {
		for j := 0; j < len(prefix); j++ {
			if strs[i][j] != prefix[j] {
				prefix = prefix[0:j]
				break
			}
		}
	}

	return prefix
}

// 优化：去掉sort步骤
func longestCommonPrefix_optimal(strs []string) string {
	prefix := strs[0]
	for i := 1; i < len(strs); i++ {
		for j := 0; j < len(prefix); j++ {
			if len(strs[i]) <= j || strs[i][j] != prefix[j] { // 添加 len(strs[i]) <= j 以达到挑选长度较小word的目的
				prefix = prefix[0:j]
				break
			}
		}
	}

	return prefix
}
