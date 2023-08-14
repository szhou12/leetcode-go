package leetcode

func stoneGameIX(stones []int) bool {
	count := make([]int, 3)
	for _, stone := range stones {
		count[stone%3]++
	}

	temp := make([]int, len(count))
	copy(temp, count)

	// 1st Round: Alice (player==0)
	// select mod==1
	if temp[1] > 0 {
		temp[1]--
		// Alice wins iff Bob loses
		if !dfs(temp, 1, 1) {
			return true
		}
	}

	// reset count[]
	temp = count
	// select mod==2
	if temp[2] > 0 {
		temp[2]--
		// Alice wins iff Bob loses
		if !dfs(temp, 2, 1) {
			return true
		}
	}

	// Alice loses if there's no certainty to win
	return false
}

// dfs(count, sum, player) := if player can win given the sum of all removed stones so far and
//
//	remaining stones (count[]) left to select
func dfs(count []int, sum int, player int) bool {
	// base case: no stones left => Bob wins
	if count[0]+count[1]+count[2] == 0 {
		if player == 1 { // Bob (player==0)
			return true
		} else {
			return false
		}
	}

	// priority: select count[0]
	if count[0] > 0 {
		count[0]--
		return !dfs(count, sum, 1-player)
	} else if sum%3 == 1 {
		if count[1] > 0 {
			count[1]--
			return !dfs(count, sum+1, 1-player)
		} else {
			return false
		}
	} else { // sum%3 == 2
		if count[2] > 0 {
			count[2]--
			return !dfs(count, sum+2, 1-player)
		} else {
			return false
		}
	}
}
