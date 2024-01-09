package leetcode

func myPow(x float64, n int) float64 {
	if n < 0 {
		return 1 / pow(x, -n)
	}

	return pow(x, n)
}

func pow(x float64, n int) float64 {
	// base cases
	if n == 0 {
		return 1
	}
	if n == 1 {
		return x
	}

	// recursion
	temp := pow(x, n/2)

	if n%2 == 0 {
		return temp * temp
	}

	return temp * temp * x
}
