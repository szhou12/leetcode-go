package amz

import (
	"testing"
)

/*
Terminal Commands:

	Run all tests that contains func_name:
	/amz $ go test -run TestMatchString

	Print out result (t.Logf) in console:
	/amz $ go test -v -run TestMatchString

	Run the only test named exactly as func_name:
	/amz $ go test -run ^TestMatchString$

	Run the only test named as TestMatchString from root directory 
		(NOTE: MUST use regex otherwise TestMatchString will be mistakenly treated as a package name):
	/leetcode-go $ go test -run ./interview-prep/amz ^TestMatchString$
*/
func TestMatchString(t *testing.T) {
	text := []string{"code", "coder"}
	pattern := []string{"co*d", "co*er"}
	expected := []bool{false, true}

	result := matchString(text, pattern)

	t.Logf("result: %v", result) // Print out the result

	for i := range result {
		if result[i] != expected[i] {
			t.Errorf("matchString(%v, %v) = %v; want %v", text, pattern, result, expected)
		}
	}
}

func TestCatchSmallerGoods(t *testing.T) {
	items := []int{1, 2, 7, 4, 5}
	start := []int{0, 0, 1}
	end := []int{1, 2, 2}
	query := []int{2, 4}

	expected := []int{2, 5}
	result := CatchSmallerGoods(items, start, end, query)

	t.Logf("result: %v", result)

	for i := range result {
		if result[i] != expected[i] {
			t.Errorf("catchSmallerGoods(%v, %v, %v, %v) = %v; want %v", items, start, end, query, result, expected)
		}
	}
}

func TestMinInputsPerWeek(t *testing.T) {
	costs := []int{1000, 500, 2000, 8000, 1500}
	weeks := 3
	expected := 9500
	result := minInputsPerWeek(costs, weeks)

	t.Logf("result: %v", result)

	if result != expected {
		t.Errorf("minInputsPerWeek(%v, %v) = %v; want %v", costs, weeks, result, expected)
	}
}

func TestEqualizeElements(t *testing.T) {
	nums1 := []int{2, 4, 1}
	nums2 := []int{1, 2, 3}
	expected := 2
	result := equalizeElements(nums1, nums2)

	t.Logf("result: %v", result)

	if result != expected {
		t.Errorf("equalizeElements(%v, %v) = %v; want %v", nums1, nums2, result, expected)
	}
}

func TestMaxFlipKeepPosCumsum(t *testing.T) {
	nums := []int{5, 2, 3, 5, 2, 3}
	expected := 3

	result := maxFlipKeepPosCumsum(nums)
	t.Logf("result: %v", result)

	if result != expected {
		t.Errorf("maxFlipKeepPosCumsum(%v) = %v; want %v", nums, result, expected)
	}
}

func TestGetMaxEnergyBurned(t *testing.T) {
	height := []int{5, 2, 5}
	expected := 43

	result := getMaxEnergyBurned(height)
	t.Logf("result: %v", result)

	if result != expected {
		t.Errorf("getMaxEnergyBurned(%v) = %v; want %v", height, result, expected)
	}
}

func TestGetFirstNonPosPresum(t *testing.T) {
	nums := []int{1, 2}
	expected := -1

	result := getFirstNonPosPresum(nums)
	t.Logf("result: %v", result)

	if result != expected {
		t.Errorf("getFirstNonPosPresum(%v) = %v; want %v", nums, result, expected)
	}

	nums = []int{2, -4, 1}
	expected = 2

	result = getFirstNonPosPresum(nums)
	t.Logf("result: %v", result)

	if result != expected {
		t.Errorf("getFirstNonPosPresum(%v) = %v; want %v", nums, result, expected)
	}

	nums = []int{1, 2, 3, -6}
	expected = 3

	result = getFirstNonPosPresum(nums)
	t.Logf("result: %v", result)

	if result != expected {
		t.Errorf("getFirstNonPosPresum(%v) = %v; want %v", nums, result, expected)
	}
}

func TestMaxMinorCourses(t *testing.T) {
	major := []int{4, 5, 2, 4}
	minor := []int{5, 6, 3, 4}
	limit := 7
	expected := 2

	result := maxMinorCourses(major, minor, limit)
	t.Logf("result: %v", result)

	if result != expected {
		t.Errorf("maxMinorCourses(%v, %v, %v) = %v; want %v", major, minor, limit, result, expected)
	}
}

func TestMaxDaysBeforeBroke(t *testing.T) {
	lend := []int{4, 6, 1, 8}
	debt := []int{7, 10, 3, 9}
	expected := 3

	result := maxDaysBeforeBroke(lend, debt)
	t.Logf("result: %v", result)

	if result != expected {
		t.Errorf("maxDaysBeforeBroke(%v, %v) = %v; want %v", lend, debt, result, expected)
	}

	lend = []int{2, 1, 5}
	debt = []int{2, 2, 5}
	expected = 3

	result = maxDaysBeforeBroke(lend, debt)
	t.Logf("result: %v", result)

	if result != expected {
		t.Errorf("maxDaysBeforeBroke(%v, %v) = %v; want %v", lend, debt, result, expected)
	}

	lend = []int{1, 1, 1, 2}
	debt = []int{2, 2, 2, 3}
	expected = 2

	result = maxDaysBeforeBroke(lend, debt)
	t.Logf("result: %v", result)

	if result != expected {
		t.Errorf("maxDaysBeforeBroke(%v, %v) = %v; want %v", lend, debt, result, expected)
	}

}

func TestMinCostPurchase(t *testing.T) {
	cost := []int{1,2,3}
	joinCost := 2
	k := 1

	expected := 3

	result := minCostPurchase(cost, joinCost, k)
	t.Logf("result: %v", result)

	if result != expected {
		t.Errorf("minCostPurchase(%v, %v, %v) = %v; want %v", cost, joinCost, k, result, expected)
	}
}

func TestGetRequestsInQueue(t *testing.T) {
	wait := []int{2, 2, 3, 1}
	expected := []int{4, 2, 1, 0}

	result := getRequestsInQueue(wait)
	t.Logf("result: %v", result)

	for i := range result {
		if result[i] != expected[i] {
			t.Errorf("getRequestsInQueue(%v) = %v; want %v", wait, result, expected)
		}
	}

	wait = []int{4, 4, 4}
	expected = []int{3, 2, 1, 0}

	result = getRequestsInQueue(wait)
	t.Logf("result: %v", result)

	for i := range result {
		if result[i] != expected[i] {
			t.Errorf("getRequestsInQueue(%v) = %v; want %v", wait, result, expected)
		}
	}

	wait = []int{3, 1, 2, 1}
	expected = []int{4, 1, 0}

	result = getRequestsInQueue(wait)
	t.Logf("result: %v", result)

	for i := range result {
		if result[i] != expected[i] {
			t.Errorf("getRequestsInQueue(%v) = %v; want %v", wait, result, expected)
		}
	}

	wait = []int{3, 1, 3, 1}
	expected = []int{4, 1, 0}

	result = getRequestsInQueue(wait)
	t.Logf("result: %v", result)

	for i := range result {
		if result[i] != expected[i] {
			t.Errorf("getRequestsInQueue(%v) = %v; want %v", wait, result, expected)
		}
	}

	wait = []int{1,1,2,5,8}
	expected = []int{5,3,2,1,0}

	result = getRequestsInQueue(wait)
	t.Logf("result: %v", result)

	for i := range result {
		if result[i] != expected[i] {
			t.Errorf("getRequestsInQueue(%v) = %v; want %v", wait, result, expected)
		}
	}
}

func TestLeastInterval(t *testing.T) {
	requests := "aaabbb"
	gap := 2
	expected := 8

	result := leastInterval(requests, gap)
	t.Logf("result: %v", result)
	
	if result != expected {
		t.Errorf("leastInterval(%v, %v) = %v; want %v", requests, gap, result, expected)
	}

	requests = "abacadaeafag"
	gap = 2
	expected = 16

	result = leastInterval(requests, gap)
	t.Logf("result: %v", result)
	
	if result != expected {
		t.Errorf("leastInterval(%v, %v) = %v; want %v", requests, gap, result, expected)
	}
}

func TestMergeIntervals(t *testing.T) {
	intervals := [][]int{{1,3},{2,6},{8,10},{15,18}}
	expected := [][]int{{1,6},{8,10},{15,18}}

	result := mergeIntervals(intervals)
	t.Logf("result: %v", result)

	for i := range result {
		if result[i][0] != expected[i][0] || result[i][1] != expected[i][1] {
			t.Errorf("mergeIntervals(%v) = %v; want %v", intervals, result, expected)
		}
	}

	intervals = [][]int{{1,4},{4,5}}
	expected = [][]int{{1,5}}

	result = mergeIntervals(intervals)
	t.Logf("result: %v", result)

	for i := range result {
		if result[i][0] != expected[i][0] || result[i][1] != expected[i][1] {
			t.Errorf("mergeIntervals(%v) = %v; want %v", intervals, result, expected)
		}
	}
}

func TestMaxSwitchingDigits(t *testing.T) {
	s := "1001"
	k := 1
	expected := 3

	result := MaxSwitchingDigits(s, k)
	t.Logf("result: %v", result)

	if result != expected {
		t.Errorf("MaxSwitchingDigits(%v, %v) = %v; want %v", s, k, result, expected)
	}

	// s = "1001"
	// k = 0 // k = 0 doesn't fit the constraint
	// expected = 2

	// result = MaxSwitchingDigits(s, k)
	// t.Logf("result: %v", result)

	// if result != expected {
	// 	t.Errorf("MaxSwitchingDigits(%v, %v) = %v; want %v", s, k, result, expected)
	// }

	s = "1001"
	k = 2
	expected = 4

	result = MaxSwitchingDigits(s, k)
	t.Logf("result: %v", result)

	if result != expected {
		t.Errorf("MaxSwitchingDigits(%v, %v) = %v; want %v", s, k, result, expected)
	}
}

func TestLexicoSmallestSymmetricStr(t *testing.T) {
	s := "yxxy"
	expected := "xyyx"

	result := LexicoSmallestSymmetricStr(s)
	t.Logf("result: %v", result)

	if result != expected {
		t.Errorf("LexicoSmallestSymmetricStr(%v) = %v; want %v", s, result, expected)
	}

	s = "ded"
	expected = "ded"

	result = LexicoSmallestSymmetricStr(s)
	t.Logf("result: %v", result)

	if result != expected {
		t.Errorf("LexicoSmallestSymmetricStr(%v) = %v; want %v", s, result, expected)
	}
}

func TestCheckSimilarStrTransform(t *testing.T) {
	new := []string{"aaccbbee", "aab"}
	old := []string{"bdbf", "aee"}
	expected := []bool{true, false}

	result := CheckSimilarStrTransform(new, old)
	t.Logf("result: %v", result)

	for i := range result {
		if result[i] != expected[i] {
			t.Errorf("CheckSimilarStrTransform(%v, %v) = %v; want %v", new, old, result, expected)
		}
	}

	new = []string{"baacbab", "accdb", "baacba"}
	old = []string{"abdbc", "ach", "abb"}
	expected = []bool{true, false, true}

	result = CheckSimilarStrTransform(new, old)
	t.Logf("result: %v", result)
	
	for i := range result {
		if result[i] != expected[i] {
			t.Errorf("CheckSimilarStrTransform(%v, %v) = %v; want %v", new, old, result, expected)
		}
	}

}

func TestFindMinCapacity(t *testing.T) {
	volumes := []int{9, 2, 4, 6}
	k := 3
	// expected := 9
	expected := findMinCapacity_v1(k, volumes)

	result := FindMinCapacity(k, volumes)
	t.Logf("result: %v", result)

	if result != expected {
		t.Errorf("FindMinCapacity(%v, %v) = %v; want %v", k, volumes, result, expected)
	}
}

func TestNumNotWorkingMachinese(t *testing.T) {
	x := []int{0, 0, 0, 0, 0, 1, 1, 1, 2, -1, -1, -2, -1}
	y := []int{-1, 0, 1, 2, -2, 0, 1, -1, 0, 1, -1, 0, 0}
	expected := 5

	result := NumNotWorkingMachinese(x, y)
	t.Logf("result: %v", result)

	if result != expected {
		t.Errorf("NumNotWorkingMachinese(%v, %v) = %v; want %v", x, y, result, expected)
	}

	x = []int{1, 1, 1, 2, 2, 2, 2, 3, 3, 3}
	y = []int{1, 2, 3, 1, 2, 3, 5, 1, 2, 3}
	expected = 2

	result = NumNotWorkingMachinese(x, y)
	t.Logf("result: %v", result)

	if result != expected {
		t.Errorf("NumNotWorkingMachinese(%v, %v) = %v; want %v", x, y, result, expected)
	}

}

func TestFindMaxValueSizeK(t *testing.T) {
	segments := [][]int{
		{1, 4, 2},
		{6, 6, 5},
		{7, 7, 7},
		{7, 9, 1},
		{9, 10, 1},
	}
	k := 5

	expected := 16

	result := FindMaxValueSizeK(segments, k)

	t.Logf("result: %v", result)

	if result != expected {
		t.Errorf("FindMaxValueSizeK(%v, %v) = %v; want %v", segments, k, result, expected)
	}

}

func TestMaxSumKGroupsMedian(t *testing.T) {
	nums := []int{2, 2, 1, 5, 3}
	k := 2
	expected := 7

	result := MaxSumKGroupsMedian(nums, k)

	t.Logf("result: %v", result)

	if result != expected {
		t.Errorf("MaxSumKGroupsMedian(%v, %v) = %v; want %v", nums, k, result, expected)
	}

	nums = []int{1, 2, 3, 4, 5}
	k = 2
	expected = 8

	result = MaxSumKGroupsMedian(nums, k)

	if result != expected {
		t.Errorf("MaxSumKGroupsMedian(%v, %v) = %v; want %v", nums, k, result, expected)
	}

	nums = []int{89, 48, 10}
	k = 3
	expected = 147

	result = MaxSumKGroupsMedian(nums, k)

	if result != expected {
		t.Errorf("MaxSumKGroupsMedian(%v, %v) = %v; want %v", nums, k, result, expected)
	}

}