package leetcode

import "strconv"

func monotoneIncreasingDigits(n int) int {
	s := strconv.Itoa(n)
	bytesArr := []byte(s)

	if len(bytesArr) <= 1 {
		return n
	}

	for i := len(bytesArr) - 1; i > 0; i-- {
		if bytesArr[i-1] > bytesArr[i] {
			bytesArr[i-1] -= 1
			for j := i; j < len(bytesArr); j++ {
				bytesArr[j] = '9'
			}
		}
	}
	res, _ := strconv.Atoi(string(bytesArr))
	return res
}
