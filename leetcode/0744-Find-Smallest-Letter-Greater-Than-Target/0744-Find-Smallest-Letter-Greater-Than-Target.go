package leetcode

func nextGreatestLetter(letters []byte, target byte) byte {
	left := 0
	right := len(letters) - 1

	for left < right {
		mid := left + (right-left)/2
		if letters[mid] == target {
			left = mid + 1
		} else if letters[mid] > target {
			right = mid
		} else {
			left = mid + 1
		}
	}

	// post-processing: all elements <= target
	if letters[left] <= target {
		return letters[0]
	}

	return letters[left]
}
