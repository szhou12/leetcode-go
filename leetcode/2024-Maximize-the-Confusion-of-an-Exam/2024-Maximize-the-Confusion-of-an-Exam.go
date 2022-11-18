package leetcode

func maxConsecutiveAnswers(answerKey string, k int) int {
	convertTs := maxConversion(answerKey, k, 'T')
	convertFs := maxConversion(answerKey, k, 'F')
	return max(convertFs, convertTs)
}

// Same as LC 1004
func maxConversion(answerKey string, k int, target byte) int {
	counter := k
	right := 0
	res := 0

	for left := 0; left < len(answerKey); left++ {
		for right < len(answerKey) && (answerKey[right] != target || counter > 0) {
			if answerKey[right] == target {
				counter--
			}
			right++
		}
		// udpate result
		res = max(res, right-left)

		// 吐
		if answerKey[left] == target {
			counter++
		}
	}
	return res
}

func max(a int, b int) int {
	if a > b {
		return a
	}
	return b
}
