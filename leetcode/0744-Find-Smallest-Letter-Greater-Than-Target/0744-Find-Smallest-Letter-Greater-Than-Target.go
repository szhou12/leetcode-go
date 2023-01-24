package leetcode

func nextGreatestLetter(letters []byte, target byte) byte {
	left := 0
	right := len(letters) - 1

	for left < right {
		mid := left + (right-left)/2
		if letters[mid] > target {
			right = mid // mid指向的元素有可能是答案
		} else { // letters[mid] <= target mid指向的元素一定不是答案
			left = mid + 1
		}
	}

	// post-processing: all elements <= target
	if letters[left] <= target {
		return letters[0]
	}

	return letters[left]
}
