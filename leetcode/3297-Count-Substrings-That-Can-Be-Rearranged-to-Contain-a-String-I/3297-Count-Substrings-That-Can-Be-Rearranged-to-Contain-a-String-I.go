package leetcode

func validSubstringCount(word1 string, word2 string) int64 {
	target := make([]int, 26)
	for i := 0; i < len(word2); i++ {
		target[word2[i]-'a']++
	}   

	n := len(word1)
	right := 0
	count := make([]int, 26)
	res := 0
	for left := 0; left < n; left++ {
		for !containWord2(count, target) && right < n {
			count[word1[right]-'a']++
			right++
		}

		if containWord2(count, target) {
			res += (n-1) - (right-1) + 1
		}

		count[word1[left]-'a']--
	}

	return int64(res)
}

func containWord2(count []int, target []int) bool {
	for i := 0; i < 26; i++ {
		if count[i] < target[i] {
			return false
		}
	}
	return true
}