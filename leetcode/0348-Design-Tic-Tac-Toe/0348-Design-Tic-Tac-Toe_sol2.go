package leetcode

/*
Leetcode Solution 1: Optimized Brute Force

The simplest and most intuitive approach is to check on every move if the current player has won. Each player makes a move by marking a cell on the board.

4 ways to win:
1. entire row
2. entire column
3. diagonal
4. anti-diagonal (col = n - row - 1)

Time complexity = O(n) as for every move we are iterating over n cells 4 times to check for each of the column, row, diagonal row, and anti-diagonal. This gives us a total O(4n)

Space complexity = O(n^2) to store the board
*/
type TicTacToe1 struct {
	board [][]int
	n int
}

func NewTicTacToe1(n int) TicTacToe1 {
	board := make([][]int, n)
	for i := 0; i < n; i++ {
		board[i] = make([]int, n)
	}

	return TicTacToe1{
		board: board,
		n: n,
	}
} 

func (this *TicTacToe1) Move(row int, col int, player int) int {
	// make a move
	this.board[row][col] = player

	// check if the current player has won
	if this.winsCol(col, player) || this.winsRow(row, player) || (row == col &&this.winsDiagonal(player)) || (col == this.n - row - 1 && this.winsAntiDiagonal(player)) {
		return player
	}

	// no one wins
	return 0
}

func (this *TicTacToe1) winsCol(col int, player int) bool {
	// given a col, loop through all rows
	for row := 0; row < this.n; row++ {
		if this.board[row][col] != player {
			return false
		}
	}
	return true
}

func (this *TicTacToe1) winsRow(row int, player int) bool {
	// given a row, loop through all cols
	for col := 0; col < this.n; col++ {
		if this.board[row][col] != player {
			return false
		}
	}
	return true
}

func (this *TicTacToe1) winsDiagonal(player int) bool {
	for row := 0; row < this.n; row++ {
		if this.board[row][row] != player {
			return false
		}
	}
	return true
}

func (this *TicTacToe1) winsAntiDiagonal(player int) bool {
	for row := 0; row < this.n; row++ {
		if this.board[row][this.n - row - 1] != player {
			return false
		}
	}
	return true
}