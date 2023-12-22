package leetcode

var M = int(1e9 + 7)

func numberOfGoodPartitions(nums []int) int {
	n := len(nums)
	// Step 1: find left & right boundaries for each element
	// map{key=element : value=index}
	right := make(map[int]int)
	for i := 0; i < n; i++ {
		right[nums[i]] = i
	}
	left := make(map[int]int)
	for i := n - 1; i >= 0; i-- {
		left[nums[i]] = i
	}

	// Step 2: diff array - annotate left boundary as +1, right boundary as -1
	diff := make([]int, n)
	for element, _ := range left {
		diff[left[element]]++
		diff[right[element]]--
	}
	// Step 3: diff array - Sweep Line to accumulate the sum
	for i := 1; i < n; i++ {
		diff[i] += diff[i-1]
	}

	// Step 4: count the most number of non-overlapped intervals
	count := 0
	for i := 0; i < n; i++ {
		if diff[i] == 0 {
			count++
		}
	}

	// Step 5: calculate # of partitions
	// n non-overlapped intervals have n-1 spaces to put the partition
	// total # of choices to put the partition = 2^(n-1)
	res := 1
	for i := 1; i <= count-1; i++ {
		res *= 2
		res %= M
	}

	return res
}
