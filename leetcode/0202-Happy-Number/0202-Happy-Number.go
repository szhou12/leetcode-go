package leetcode

func isHappy(n int) bool {
	// use map to check if a same number repeats
	// if repeats, then we encounter a cycle. It's not a happy number.
	record := map[int]int{}

	return testHappyNumber(n, record)
}

func testHappyNumber(n int, record map[int]int) bool {
	if n == 1 {
		return true
	}
	if _, ok := record[n]; ok {
		return false
	} else {
		record[n] = n
	}

	newNum := 0

	for n > 0 {
		digit := n % 10
		newNum += square(digit)
		n /= 10
	}

	return testHappyNumber(newNum, record)

}

func square(n int) int {
	return n * n
}
