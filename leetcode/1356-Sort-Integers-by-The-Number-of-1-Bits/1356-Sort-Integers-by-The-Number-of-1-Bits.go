package leetcode

import (
	"sort"
	"strconv"
)

// My solution - convert number to binary representation
func sortByBits(arr []int) []int {
	sort.Slice(arr, func(i, j int) bool {
		if countOnes(arr[i]) == countOnes(arr[j]) {
			return arr[i] < arr[j]
		}
		return countOnes(arr[i]) < countOnes(arr[j])
	})
	return arr

}

func convertBits(x int) []byte {
	bitStr := strconv.FormatInt(int64(x), 2)
	return []byte(bitStr)
}

func countOnes(x int) int {
	xBytes := convertBits(x)
	res := 0
	for _, bit := range xBytes {
		res += int(bit - '0')
	}
	return res
}
