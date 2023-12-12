package leetcode

var M = int(1e9+7)

func maxSum(nums []int, k int) int {
	counter := make([]int, 32)
	for _, num := range nums {
		for i := 0; i < 32; i++ {
			if (num >> i) & 1 == 1 {
				counter[i]++
			}
		}
	}

	res := 0
	for t := 0; t < k; t++ {
		x := 0
		for i := 31; i>=0; i-- {
			if counter[i] > 0 {
				x += 1<<i
				counter[i]--
			}
		}
		res += x * x
		res %= M
	}

	return res
}