package leetcode

//  ------------------------
//             mid
//     ----          ----
//     i   r         l   j
// pre[r] := max prizes that can be collected from k-len segment in p[0:r]
// post[l] := max prizes that can be collected from k-len segment in p[l:n-1]
func maximizeWin(prizePositions []int, k int) int {
	n := len(prizePositions)

	// Edge case: if max position length <= 2k, must overlap, then we can pick up all prizes
	if prizePositions[n-1]-prizePositions[0] <= 2*k {
		return n
	}

	gMax := 0 // record max length for each iter
	pre := make([]int, n)
	i := 0
	for r := 0; r < n; r++ {
		for prizePositions[r]-prizePositions[i] > k {
			i++
		}
		// prizePositions[r]-prizePositions[i] == k
		length := r - i + 1
		gMax = max(gMax, length)
		pre[r] = gMax
	}

	gMax = 0
	post := make([]int, n)
	j := n - 1
	for l := n - 1; l >= 0; l-- {
		for prizePositions[j]-prizePositions[l] > k {
			j--
		}
		// prizePositions[j]-prizePositions[l] == k
		length := j - l + 1
		gMax = max(gMax, length)
		post[l] = gMax
	}

	res := 0
	for mid := 0; mid+1 < n; mid++ { // 不考虑顶到右边的情况, 因为这样只剩1个segment (pre[]), 题目要求必须有2个segment
		res = max(res, pre[mid]+post[mid+1])
	}
	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
