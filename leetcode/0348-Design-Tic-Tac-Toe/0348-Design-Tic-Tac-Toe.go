package leetcode

type TicTacToe struct {
	rows         []int
	cols         []int
	diagonal     int
	antiDiagonal int
	size         int
}

func Constructor(n int) TicTacToe {
	return TicTacToe{
		rows: make([]int, n),
		cols: make([]int, n),
		size: n,
	}
}

func (this *TicTacToe) move(row int, col int, player int) int {
	var toAdd int
	if player == 1 {
		toAdd = 1
	} else {
		toAdd = -1
	}

	this.rows[row] += toAdd
	this.cols[col] += toAdd

	if row == col {
		this.diagonal += toAdd
	}

	if col+row == len(this.cols)-1 {
		this.antiDiagonal += toAdd
	}

	// Check win
	// Time = O(1)
	if abs(this.rows[row]) == this.size || abs(this.cols[col]) == this.size ||
		abs(this.diagonal) == this.size || abs(this.antiDiagonal) == this.size {
		return player
	}

	return 0
}

func abs(a int) int {
	if a >= 0 {
		return a
	}
	return -a
}
