package mathematics
/*
Pascal's Triangle
[0]	1          => coeffs in (a + b)^0
[1]	1 1        => coeffs in (a + b)^1
[2]	1 2 1      => coeffs in (a + b)^2
[3]	1 3 3 1    => coeffs in (a + b)^3
[4]	1 4 6 4 1  => coeffs in (a + b)^4
...

利用杨辉三角，计算任意组合数C(n, k)的值，也就是二项式系数。
*/
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
		}
	}
	return comb
}