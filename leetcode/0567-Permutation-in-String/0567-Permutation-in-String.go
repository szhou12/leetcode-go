package leetcode

func checkInclusion(s1 string, s2 string) bool {
	need := make(map[byte]int)
	for i := 0; i < len(s1); i++ {
		need[s1[i]]++
	}

	window := make(map[byte]int)
	left, right := 0, 0
	count := 0
	for right < len(s2) {
		rightElement := s2[right]
		right++
		if _, ok := need[rightElement]; ok {
			window[rightElement]++
			if window[rightElement] == need[rightElement] {
				count++
			}
		}

		for right-left >= len(s1) {
			// update result
			if count == len(need) {
				return true
			}

			leftElement := s2[left]
			left++
			if _, ok := need[leftElement]; ok {
				if window[leftElement] == need[leftElement] {
					count--
				}
				window[leftElement]--
			}
		}
	}
	return false
}

// 写法二: 写法类似 LC 438
func checkInclusion2(s1 string, s2 string) bool {
	need := make(map[byte]int)
	for i := 0; i < len(s1); i++ {
		need[s1[i]]++
	}

	window := make(map[byte]int)
	for right := 0; right < len(s2); right++ {
		// 长
		window[s2[right]]++
		// 过长，收缩
		if right-len(s1) >= 0 {
			window[s2[right-len(s1)]]--
		}
		// left = right - len(s1) + 1
		if right-len(s1)+1 >= 0 && isEqual(need, window) {
			return true
		}
	}
	return false
}

func isEqual(need map[byte]int, window map[byte]int) bool {
	for k, v := range window {
		if need[k] != v {
			return false
		}
	}
	return true
}
