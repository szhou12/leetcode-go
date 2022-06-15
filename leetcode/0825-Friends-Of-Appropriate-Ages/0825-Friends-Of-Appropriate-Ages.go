package leetcode

func numFriendRequests(ages []int) int {
	result := 0
	countMap := make(map[int]int)
	for _, age := range ages {
		countMap[age]++
	}

	for x, xcount := range countMap {
		for y, ycount := range countMap {
			if makeRequest(x, y) {
				if x != y {
					result += xcount * ycount
				} else {
					result += xcount * (ycount - 1)
				}
			}
		}
	}
	return result
}

func makeRequest(x int, y int) bool {
	if 2*(y-7) <= x {
		return false
	}
	if y > x {
		return false
	}
	if y > 100 && x < 100 {
		return false
	}
	return true
}
