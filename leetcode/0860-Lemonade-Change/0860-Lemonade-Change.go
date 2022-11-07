package leetcode

func lemonadeChange(bills []int) bool {
	// edge case: the first bill MUST be $5
	if bills[0] != 5 {
		return false
	}

	// savings[0] = # of $5
	// savings[1] = # of $10
	savings := []int{0, 0}
	for i := 0; i < len(bills); i++ {
		if bills[i] == 5 {
			savings[0]++
		}

		if bills[i] == 10 {
			savings[1]++
		}

		giveback := bills[i] - 5
		// received $10, MUST give back $5
		if giveback == 5 {
			if savings[0] > 0 {
				savings[0]--
			} else {
				return false
			}
		}
		// received $20, give back either 1*$10 + 1*$5 (higher piority) or 3*$5
		if giveback == 15 {
			if savings[0] > 0 && savings[1] > 0 {
				savings[0]--
				savings[1]--
			} else if savings[0] >= 3 {
				savings[0] -= 3
			} else {
				return false
			}
		}
	}

	return true
}
