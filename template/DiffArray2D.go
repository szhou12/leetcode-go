package template

type Diff2D struct {
	sum  [][]int
	diff [][]int
	m, n int
}

func newDiff2D(m, n int) *Diff2D {

	sum := make([][]int, m+1)
	diff := make([][]int, m+1)

	for i := 0; i <= m; i++ {
		sum[i] = make([]int, n+1)
		diff[i] = make([]int, n+1)
	}

	diff2D := Diff2D{
		sum:  sum,
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
	diff2D.sum[0][0] = diff2D.diff[0][0]
	for i := 0; i < diff2D.m; i++ {
		for j := 0; j < diff2D.n; j++ {
			var a, b, c int
			if i == 0 {
				a = 0
			} else {
				a = diff2D.sum[i-1][j]
			}
			if j == 0 {
				b = 0
			} else {
				b = diff2D.sum[i][j-1]
			}
			if i == 0 || j == 0 {
				c = 0
			} else {
				c = diff2D.sum[i-1][j-1]
			}
			diff2D.sum[i][j] = a + b - c + diff2D.diff[i][j]
		}
	}
}
