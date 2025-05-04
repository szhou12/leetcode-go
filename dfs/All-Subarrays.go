package dfs

// Get all subarrays of a given array
// Time = O(n^2)
func allSubarrays(arr []int) [][]int {
	var res [][]int

	n := len(arr)
	for i := 0; i < n; i++ {
		for j := i; j < n; j++ {
			subarray := make([]int, len(arr[i:j+1]))
			copy(subarray, arr[i:j+1])
			res = append(res, subarray)
		}
	}

	return res
}
