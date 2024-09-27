package leetcode

func validSubstringCount(word1 string, word2 string) int64 {
	target := make([]int, 26)
	for i := 0; i < len(word2); i++ {
		target[word2[i] - 'a']++
	}
	T := 0 // number of kinds of distinct characters in word2
	for i := 0; i < 26; i++ {
		if target[i] > 0 {
			T++
		}
	}

	n := len(word1)
	res := 0
	count := make([]int, 26)
	right := 0
	t := 0 // number of kinds of distinct characters in the current sliding window
	for left := 0; left < n; left++ {
		for !containWord2(t, T) && right < n {
			count[word1[right] - 'a']++ 
			// 如果"吃掉"right指向的这个char后，刚好达到target中该char的数量
			if count[word1[right] - 'a'] == target[word1[right] - 'a'] {
				t++
			}
			right++
		}
		// 上一个for循环结束跳出时，right指向合法滑窗的下一位
		if containWord2(t, T) {
			res += (n-1) - (right-1) + 1
		}

		count[word1[left] - 'a']--
		// 如果"吐出"left指向的这个char后，刚好达到target中该char的数量-1
		if count[word1[left] - 'a'] == target[word1[left] - 'a'] - 1 {
			t--
		}
	}

	return int64(res)

}

func containWord2(t, T int) bool {
	return t == T
}