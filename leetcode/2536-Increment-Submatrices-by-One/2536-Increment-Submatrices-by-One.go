package leetcode

func rangeAddQueries(n int, queries [][]int) [][]int {
	sol := newDiff2D(n, n)
	for _, query := range queries { // O(q)
		sol.set(query[0], query[1], query[2], query[3], 1)
	}
	sol.compute() // O(n*n)

	res := make([][]int, n)
	for i := 0; i < n; i++ {
		res[i] = make([]int, n)
	}

	// sol.f 最后一行和最后一列是冗余, 只是为了set()计算差分, 输出时要去掉
	// O(n*n)
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			res[i][j] = sol.f[i][j]
		}
	}
	return res
}

type Diff2D struct {
	f    [][]int
	diff [][]int
	m, n int
}

func newDiff2D(m, n int) *Diff2D {

	f := make([][]int, m+1)
	diff := make([][]int, m+1)

	for i := 0; i <= m; i++ {
		f[i] = make([]int, n+1)
		diff[i] = make([]int, n+1)
	}

	diff2D := Diff2D{
		f:    f,
		diff: diff,
		m:    m,
		n:    n,
	}
	return &diff2D
}

func (diff2D *Diff2D) set(x0 int, y0 int, x1 int, y1 int, val int) {
	diff2D.diff[x0][y0] += val
	diff2D.diff[x0][y1+1] -= val
	diff2D.diff[x1+1][y0] -= val
	diff2D.diff[x1+1][y1+1] += val
}

func (diff2D *Diff2D) compute() {
	diff2D.f[0][0] = diff2D.diff[0][0]
	for i := 0; i < diff2D.m; i++ {
		for j := 0; j < diff2D.n; j++ {
			var a, b, c int
			if i == 0 {
				a = 0
			} else {
				a = diff2D.f[i-1][j]
			}
			if j == 0 {
				b = 0
			} else {
				b = diff2D.f[i][j-1]
			}
			if i == 0 || j == 0 {
				c = 0
			} else {
				c = diff2D.f[i-1][j-1]
			}
			diff2D.f[i][j] = a + b - c + diff2D.diff[i][j]
		}
	}
}
