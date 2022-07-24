package leetcode

func convertToTitle(columnNumber int) string {
	var res string
	for columnNumber > 0 {
		temp := columnNumber % 26
		if temp == 0 {
			res = "Z" + res
		} else {
			res = string(rune(temp)) + res
		}
		columnNumber /= 26
	}

	return res
}

// num = 701
// 701 % 26 = 25 -> Y
// 701 / 26 = 26 -> Z

// 701 / 26 = 2
// 52
// 181 / 26 = 6
// 26*6 = 156
// 181-156 = 25
// 26 * 26 = 156 + 520 = 676
// 701-676 = 25 -> Y
