package leetcode

/*
nums1[i] * nums2[j] <= mid
if nums1[i] > 0:
find largest j s.t. nums2[j] <= floor(mid/nums1[i]) --> [0, j] // bc nums2 sorted
count += (j-0+1)
if nums1[i] == 0:
    if mid >= 0: count += len(nums2) // all nums2 elements satisfy the requirement
    if mid < 0: count += 0 // no nums2 element satisfy the requirement
if nums1[i] < 0:
find smallest j s.t. nums2[j] >= ceil(mid/nums1[i]) --> [j, n-1]
count += (n-1-j+1)
*/
func kthSmallestProduct(nums1 []int, nums2 []int, k int64) int64 {
	left := int(-1e5 * 1e5)
	right := int(1e5 * 1e5)

	for left < right {
		mid := left + (right-left)/2

		count := int64(countLessOrEqual(nums1, nums2, mid))
		if count < k {
			left = mid + 1
		} else {
			right = mid
		}
	}

	return int64(left)
}

func countLessOrEqual(nums1 []int, nums2 []int, mid int) int {
	var res int
	for i := 0; i < len(nums1); i++ {
		if nums1[i] > 0 {
			y := mid / nums1[i]
			j := upperBound(nums2, y)
			res += (j + 1)
		} else if nums1[i] == 0 {
			if mid < 0 {
				res += 0
			} else {
				res += len(nums2)
			}
		} else { // nums1[i] < 0
			y := mid/nums1[i] + 1 // +1 for ceiling
			j := lowerBound(nums2, y)
			res += (len(nums2) - 1 - j + 1)
		}
	}

	return res
}

// find smallest index j s.t. nums[j] >= target
func lowerBound(nums []int, target int) int {
	left := 0
	right := len(nums) - 1

	for left < right {
		mid := left + (right-left)/2
		if nums[mid] >= target {
			right = mid
		} else {
			left = mid + 1
		}
	}

	return left

}

// find largest index j s.t. nums[j] <= target
// func upperBound(nums []int, target int) int {
// 	left := 0
// 	right := len(nums) - 1

// 	for left < right {
// 		mid := left + (right-left)/2
// 		if nums[mid] <= target {
// 			left = mid
// 		} else {
// 			right = mid - 1
// 		}
// 	}
// 	return left
// }

func upperBound(nums []int, target int) int {
	left := 0
	right := len(nums) - 1

	for left < right {
		mid := left + (right-left)/2
		if nums[mid] > target {
			right = mid
		} else {
			left = mid + 1
		}
	}
	return left
}
