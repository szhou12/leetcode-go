package mathematics

// 求 最大公约数
// greatest common divisor (GCD) via Euclidean algorithm (辗转相除法)

// recursive way
// 注意！用的if
func gcd(a, b int) int {
	if b == 0 {
		return a
	}
	return gcd(b, a%b)
}

// iterative way
// 注意！用的for
func GCD(a, b int) int {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}
