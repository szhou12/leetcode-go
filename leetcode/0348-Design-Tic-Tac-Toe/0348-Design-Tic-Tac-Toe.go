package leetcode

/*
Leetcode Solution 2: Optimised Approach

Intuition:
Each player makes a move by marking a cell with a value (e.g., 1 or -1). There are always n cells on a row/col/diagonal/anti-diagonal. Thus, to win either of these, a player must have marked n times (sum = n or -n). So after a move, we check if i-th row has sum being n or -n OR if i-th col has sum being n or -n OR if the diagonal OR anti-diagonal has sum being n or -n.

Time complexity = O(1) because for every move, we mark a particular row, column, diagonal, and anti-diagonal in constant time.

Space complexity = O(n) because we use arrays rows and cols of size n. The variables diagonal and antiDiagonal use constant extra space.
*/
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
		diagonal: 0,
		antiDiagonal: 0,
		size: n,
	}
}

func (this *TicTacToe) Move(row int, col int, player int) int {
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

	if col == len(this.cols)-row-1 {
		this.antiDiagonal += toAdd
	}

	// Check win
	// Time = O(1)
	if abs(this.rows[row]) == this.size || abs(this.cols[col]) == this.size ||
		abs(this.diagonal) == this.size || abs(this.antiDiagonal) == this.size {
		return player
	}

	// no one wins
	return 0
}

func abs(a int) int {
	if a >= 0 {
		return a
	}
	return -a
}
