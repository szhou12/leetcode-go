package leetcode

func findKDistantIndices(nums []int, key int, k int) []int {
	n := len(nums)
	res := make([]int, 0)

	r := 0

	for j := 0; j < n; j++ {
		if nums[j] == key {
			l := max(j-k, r) // [0: r] are elements checked already
			r = min(j+k, n-1) + 1

			for i := l; i < r; i++ {
				res = append(res, i)
			}
		}
		
	}

	return res
}