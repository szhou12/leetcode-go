package tricks

// Split integer into slice of single digits
// 注意：这个方法是从末位往前得到digit
// e.g. 273 -> [3, 7, 2]
func splitInt(n int) []int {
	res := make([]int, 0)
	for n > 0 {
		curDigit := n % 10 // 得到最末位的digit
		res = append(res, curDigit)
		n /= 10 // 砍掉最末位的digit
	}

	return res
}
