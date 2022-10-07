# [744. Find Smallest Letter Greater Than Target](https://leetcode.com/problems/find-smallest-letter-greater-than-target/)

## Solution idea

### Binary Search 经典变形题 - 找最小的大于Target的元素

* 注意Post-processing找不到的情况: 数组内所有元素都 $\leq$ Target

Time compelxity = $O(\log n)$


### 拓展：找最大的小于Target的元素

```
func binarySearch(nums []int, target int) int {
	left := 0
	right := len(nums) - 1
	for left < right {
		mid := right - (right-left)/2
		if nums[mid] == target {
			right = mid - 1
		} else if nums[mid] < target {
			left = mid
		} else {
			right = mid - 1
		}
	}

	// post-processing: 找不到小于的元素
	if nums[left] >= target {
		return -1
	}

	return left
}
```