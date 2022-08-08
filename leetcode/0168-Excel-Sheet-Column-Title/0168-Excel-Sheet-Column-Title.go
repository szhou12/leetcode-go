package leetcode

func convertToTitle(columnNumber int) string {
	var res string

	for columnNumber > 0 {
		lastLetterByte := (columnNumber - 1) % 26 // Note: 这里 (columnNumber - 1) 为了 'A' + byte - 1 e.g. 'Z' = 'A' + 26 - 1
		letter := string('A' + lastLetterByte)
		res = letter + res                     // 每一轮取模顺序是从最低位 -> 最高位，所以这里是prepend
		columnNumber = (columnNumber - 1) / 26 // Note: 这里 (columnNumber - 1) 因为题意 A-Z 对应 1-26, 这样，Z/26 应该为0 但为1. -1使A-Z 对应 0-25
	}

	return res
}

// num = 701 => ZY
// (701-1) % 26 = 24 -> Y
// (701-1) / 26 = 26
// (26-1) % 26 = 25 -> Z
