package leetcode

// Method 1: Divide and Conquer
func longestSubstring(s string, k int) int {
	return findLongest(s, 0, len(s), k)
}

func findLongest(s string, start int, end int, k int) int {
	// base case: if length of (end - start) < k, already cannot satisfy condition, return 0
	if end-start < k {
		return 0
	}

	// create freq map for each character between s[start : end]
	count := make([]int, 26)
	for i := start; i < end; i++ { // O(n)
		idx := int(s[i] - 'a')
		count[idx]++
	}

	// find spliter char in s[start : end]
	for i := 0; i < 26; i++ { // O(26)
		if 0 < count[i] && count[i] < k {
			for j := start; j < end; j++ { // O(n)
				if int(s[j]) == i+'a' {
					left := findLongest(s, start, j, k)  // O(n)
					right := findLongest(s, j+1, end, k) // O(n)
					return max(left, right)
				}
			}
		}
	}

	// if jump out of loop, it means all chars >= k
	return end - start
}

func max(a int, b int) int {
	if a > b {
		return a
	}
	return b
}

// Method 2: Sliding Window
func longestSubstring_sliding_window(s string, k int) int {
	res := 0

	for m := 0; m <= 26; m++ {
		res = max(res, fingLongestforM(s, k, m))
	}

	return res
}

func fingLongestforM(s string, k int, m int) int {
	right := 0
	res := 0

	letterFreq := make(map[byte]int) // count frequency of each char within the window
	count := 0                       // number of distinct chars who repeat >= k times within the window

	for left := 0; left < len(s); left++ {
		for right < len(s) && len(letterFreq) <= m { // 注意这里是 <= m 而不是 < m 因为下一个有可能遇到之前遇到的元素, 此时, 我们应该把它包括进来
			rightElement := s[right]
			letterFreq[rightElement]++
			if letterFreq[rightElement] == k {
				count++
			}

			// update result
			if count == m && len(letterFreq) == m {
				res = max(res, right-left+1)
			}

			right++
		}

		leftElement := s[left]
		letterFreq[leftElement]--
		if letterFreq[leftElement] == k-1 {
			/*
			* 注意:
			*  这里的条件一定一定一定不能写成 letterFreq[leftElement] < k !!!
			*  这是因为, 如果写成这样，那么在吐相同元素时，上一轮freq=k-1, 这一轮freq=k-2
			*  就都满足 < k, 都会诱发 count--.
			*  但, count记录的是不同元素的个数, 同一个元素只能诱发一次 count-- (第一次<k时, ie. ==k-1)
			 */
			count--
		}
		if letterFreq[leftElement] == 0 {
			delete(letterFreq, leftElement)
		}
	}
	return res
}
