package leetcode

// Merge method in MergeSort
func merge(nums1 []int, m int, nums2 []int, n int) {
	// corner case 1
	if n == 0 {
		return
	}
	// corner case 2
	if m == 0 {
		copy(nums1, nums2)
		return
	}

	helper := make([]int, m)
	for k := 0; k < m; k++ {
		helper[k] = nums1[k]
	}

	l, r, left := 0, 0, 0
	for l < m && r < n {
		if helper[l] <= nums2[r] {
			nums1[left] = helper[l]
			left++
			l++
		} else {
			nums1[left] = nums2[r]
			left++
			r++
		}
	}

	for l < m {
		nums1[left] = helper[l]
		left++
		l++
	}

	for r < n {
		nums1[left] = nums2[r]
		left++
		r++
	}
}

// Optimal solution: loop backwards + select the larger element each time
func merge_optimal(nums1 []int, m int, nums2 []int, n int) {
	for p := m + n; m > 0 && n > 0; p-- {
		if nums1[m-1] <= nums2[n-1] {
			nums1[p-1] = nums2[n-1]
			n--
		} else {
			nums1[p-1] = nums1[m-1]
			m--
		}
	}

	//post-processing
	for ; n > 0; n-- {
		nums1[n-1] = nums2[n-1]
	}
}
