package leetcode

// simulation
func triangularSum(nums []int) int {
	n := len(nums)

	for n > 1 {
		for i := 0; i < n-1; i++ {
			nums[i] = (nums[i] + nums[i+1]) % 10
		}
		n--
	}

	return nums[0]
}

// pascal's triangle
func triangularSum_pt(nums []int) int {
	n := len(nums)
	comb := pascalTriangle(n)

	res := 0
	for i := 0; i < n; i++ {
		res += (comb[n-1][i] * nums[i]) % 10
	}
	return res % 10
}

func pascalTriangle(n int) [][]int {
	// comb[n][k] := C(n, k)
	comb := make([][]int, n)
	for i := 0; i < n; i++ {
		comb[i] = make([]int, n)
	}

	for i := 0; i < n; i++ {
		comb[i][0] = 1
		comb[i][i] = 1
		if i == 0 {
			continue
		}
		for j := 1; j < i; j++ {
			comb[i][j] = comb[i-1][j-1] + comb[i-1][j]
			comb[i][j] %= 10
		}
	}
	return comb
}