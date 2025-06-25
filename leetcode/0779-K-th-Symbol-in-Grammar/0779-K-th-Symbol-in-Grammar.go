package leetcode

func kthGrammar(n int, k int) int {
	// base case
	if n == 1 {
		return 0
	}

	// recursion
	// if k%2 == 0, even, right child, its parent is at k/2
	// if k%2 == 1, odd, left child, its parent is at (k+1)/2
	if k%2 == 0 {
		if kthGrammar(n-1, k/2) == 0 {
			return 1
		} else {
			return 0
		}
	} else {
		if kthGrammar(n-1, (k+1)/2) == 0 {
			return 0
		} else {
			return 1
		}
	}
}