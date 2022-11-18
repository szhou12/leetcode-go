package leetcode

// Demo
// p = ab
// s = e  i  d  b  a  o  o  o
//     lr
//     l  r
//     l->l  r
//     l->l->l  r
//     l->l->l->l  r
func findAnagrams(s string, p string) []int {
	need := make(map[byte]int)
	for i := 0; i < len(p); i++ {
		need[p[i]]++
	}

	window := make(map[byte]int)
	left, right := 0, 0
	res := make([]int, 0)
	count := 0
	for right < len(s) {
		rightElement := s[right] // 吃
		right++
		if _, ok := need[rightElement]; ok {
			window[rightElement]++
			if window[rightElement] == need[rightElement] {
				count++
			}
		}

		// 开始收缩左边界的条件: [left, right) 定长 len(p) (ie. 上限不能超过len(p))
		// 这里看似是 for-loop，其实只会有 len(p) (进loop) 和 len(p)-1 (出loop) 这两种情况
		for right-left >= len(p) {
			// update result
			if count == len(need) {
				res = append(res, left)
			}

			leftElement := s[left] // 吐
			left++
			if _, ok := window[leftElement]; ok {
				if window[leftElement] == need[leftElement] {
					count--
				}
				window[leftElement]--
			}
		}
	}
	return res
}

// 写法二: 写法类似 LC 1052
func findAnagrams2(s string, p string) []int {
	need := make(map[byte]int)
	for i := 0; i < len(p); i++ {
		need[p[i]]++
	}

	window := make(map[byte]int)
	res := make([]int, 0)

	for right := 0; right < len(s); right++ {
		window[s[right]]++
		if right-len(p) >= 0 {
			window[s[right-len(p)]]--
		}
		if right-len(p)+1 >= 0 && isEqual(window, need) {
			res = append(res, right-len(p)+1)
		}
	}

	return res
}

func isEqual(a map[byte]int, b map[byte]int) bool {
	for k, v := range a {
		if b[k] != v {
			return false
		}
	}
	return true
}
