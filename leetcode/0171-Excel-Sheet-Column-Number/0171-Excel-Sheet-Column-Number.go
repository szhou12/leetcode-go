package leetcode

// Optimal Solution: 最高位 -> 最低位
func titleToNumber(columnTitle string) int {
	res := 0
	val := 0

	for i := 0; i < len(columnTitle); i++ {
		val = int(columnTitle[i] - 'A' + 1)
		res = 26*res + val
	}
	return res
}

// My solution: calculate from the last letter 最低位 -> 最高位 (较为麻烦)
func titleToNumber_MY(columnTitle string) int {
	res := 0

	for i := len(columnTitle) - 1; i >= 0; i-- {
		letterByte := columnTitle[i] - 'A' + 1
		res += pow(26, len(columnTitle)-1-i) * int(letterByte)
	}
	return res
}

func pow(x int, n int) int {
	if n == 0 {
		return 1
	}
	if n == 1 {
		return x
	}

	temp := pow(x, n/2)
	if n%2 == 0 {
		return temp * temp
	}

	return temp * temp * x
}
