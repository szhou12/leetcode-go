package mathematics

// 最小公倍数

// find LCM of a, b
func lcm(a, b int) int {
	return a * b / gcd(a, b)
}

// find Least Common Multiple (LCM) via GCD
func LCM(a, b int, integers ...int) int {
	result := a * b / GCD(a, b)

	for i := 0; i < len(integers); i++ {
		result = LCM(result, integers[i])
	}

	return result
}
