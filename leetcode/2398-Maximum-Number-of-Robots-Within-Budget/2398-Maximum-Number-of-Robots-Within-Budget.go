package leetcode

func maximumRobots(chargeTimes []int, runningCosts []int, budget int64) int {
	n := len(chargeTimes)

	left, right := 0, n
	for left < right {
		mid := right - (right-left) / 2
		if isValid(mid, chargeTimes, runningCosts, budget) {
			left = mid
		} else {
			right = mid - 1
		}
	}

	return left
}

// Sliding Window Max
func isValid(k int, chargeTimes []int, runningCosts []int, budget int64) bool {
	n := len(chargeTimes)
	sum := 0

	dq := make([]int, 0) // store chargeTimes[] index [i-k+1 : i] (inclusive)

	for i := 0; i < n; i++ {
		sum += runningCosts[i]

		// retire tail's drawfs
		for len(dq) > 0 && chargeTimes[dq[len(dq)-1]] < chargeTimes[i] {
			dq = dq[:len(dq) - 1]
		}
		// enqueue new element
		dq = append(dq, i)
		// retire head's oldies
		for len(dq) > 0 && dq[0] <= i - k {
			dq = dq[1:]
		}

		// i-k+1 >= 0
		if i >= k-1 {
			cost := chargeTimes[dq[0]] + k * sum
			if int64(cost) <= budget {
				return true
			}
			sum -= runningCosts[i-k+1]

		}
	}

	return false
}