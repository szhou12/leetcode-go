package leetcode

/*
Binary Search

mid = (1 + n) / 2
for each mid, count the # of elemenets <= mid
if count > mid, by Pigeonhole Principle, left half contains the duplicate and mid may be a candidate:
	right = mid
else:
	left = mid + 1

return left
*/
func findDuplicate(nums []int) int {
	n := len(nums)

	left, right := 1, n-1

	for left < right {
		mid := left + (right - left) / 2

		// count the # of elements <= mid
		count := 0
		for _, num := range nums {
			if num <= mid {
				count++
			}
		}

		if count > mid {
			right = mid
		} else {
			left = mid + 1
		}
	}

	return left
}