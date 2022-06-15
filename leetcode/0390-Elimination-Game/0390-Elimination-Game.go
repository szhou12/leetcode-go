package leetcode

func lastRemaining(n int) int {
	left := true
	remaining := n
	head := 1
	step := 1
	for remaining > 1 {
		if left || remaining%2 == 1 {
			head = head + step
		}
		remaining = remaining / 2
		step = step * 2 // KEY!!
		left = !left
	}
	return head
}
