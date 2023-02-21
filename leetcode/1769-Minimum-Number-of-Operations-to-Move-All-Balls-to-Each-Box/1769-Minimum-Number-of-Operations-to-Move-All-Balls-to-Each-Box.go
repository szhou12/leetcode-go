package leetcode

// leftMoves[i] := # of moves to move all balls in boxes[0:i-1] to cell i
// rightMoves[i] := # of moves to move all balls in boxes[i+1:n-1] to cell i
// leftMoves[0] = rightMoves[n-1] = 0
// left[i] := # of balls in boxes[0:i-1]
// right[i] := # of balls in boxes[i+1:n-1]
func minOperations(boxes string) []int {
	n := len(boxes)

	leftMoves := make([]int, n)
	left := 0
	for i := 1; i < n; i++ {
		if boxes[i-1] == '1' {
			left += 1
		}
		leftMoves[i] = leftMoves[i-1] + left*1
	}

	rightMoves := make([]int, n)
	right := 0
	for i := n - 2; i >= 0; i-- {
		if boxes[i+1] == '1' {
			right += 1
		}
		rightMoves[i] = rightMoves[i+1] + right*1
	}

	res := make([]int, n)
	for i := 0; i < n; i++ {
		res[i] = leftMoves[i] + rightMoves[i]
	}
	return res

}
