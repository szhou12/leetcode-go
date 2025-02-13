package leetcode

func maxScore(points []int, m int) int64 {
	left, right := 0, int(1e15) // (10^9)*(10^6)

	for left < right {
		mid := right - (right-left)/2
		if isValid(points, m, mid) {
			left = mid
		} else {
			right = mid - 1
		}
	}

	return int64(left)
}

func isValid(points []int, m int, targetScore int) bool {
	n := len(points)

	// initial state: jump in
	count := 1
	curScore := points[0] // current score at i

	for i := 0; i < n; i++ {
		// special case: if arrive at n-1, it means we may need to fill up the gap at n-1
		// we move back and forth between n-2 and n-1 (jump to left) to add points at n-1
		if i == n-1 {
			// MUST check if n-1 is fully pre-filled cuz even tho else-case will only have n-1 not fully filled, but if-case may have n-1 fully pre-filled
			if curScore >= targetScore {
				return true
			}
			d := (targetScore - curScore - 1) / points[i] + 1
			return count+(2*d) <= m
		}

		if curScore >= targetScore { // i already satisfied
			count++ // first count the move that the pointer moves to the right i+1. we will def move to right cuz we have not added any points to i+1 yet
			if count > m {
				return false
			}
			curScore = points[i+1] // if pass the check, prepare initial score at i+1
		} else {
			// we need to fill up the gap by moving back and forth
			d := (targetScore-curScore-1)/points[i] + 1 // ceil(a/b) = (a-1)/b + 1

			// early checking: if filling at n-2 also satisfies the filling at n-1, then no need to move to n-1
			if i+1 == n-1 && points[i+1]*d >= targetScore && count+(2*d) <= m {
				return true
			}

			count += 2*d + 1 // 2*d for back and forth, 1 for moving to the right
			if count > m {
				return false
			}
			curScore = points[i+1] * (d + 1) // prepare initial score at i+1
		}

	}
	return true
}
