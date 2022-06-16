package leetcode

// Greedy solution
func canCompleteCircuit(gas []int, cost []int) int {
	n := len(gas)

	// pre-process: if total gas < 0, definitely can't complete circulation
	sum := 0
	for i := 0; i < n; i++ {
		sum += gas[i] - cost[i]
	}
	if sum < 0 {
		return sum
	}

	// Greedy part
	tank := 0
	start := 0
	for i := 0; i < n; i++ {
		tank += gas[i] - cost[i]
		if tank < 0 {
			// it means unable to reach i from start, hence try start = i + 1
			tank = 0
			start = i + 1
		}
	}

	if start == n {
		return 0
	}
	return start
}
