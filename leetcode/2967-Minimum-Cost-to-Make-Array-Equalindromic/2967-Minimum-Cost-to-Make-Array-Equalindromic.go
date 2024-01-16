package leetcode

import (
	"bytes"
	"math"
	"slices"
	"strconv"
	"strings"
)

func minimumCost(nums []int) int64 {
	n := len(nums)
	slices.Sort(nums)

	var median int
	if n%2 == 1 {
		median = nums[n/2]
	} else {
		median = (nums[n/2] + nums[n/2-1]) / 2
	}

	// find closest palindrome to median
	medianStr := strconv.Itoa(median)
	smaller := nextSmallerOrEqual(medianStr)
	larger := nextGreaterOrEqual(medianStr)
	cand1, _ := strconv.Atoi(smaller)
	cand2, _ := strconv.Atoi(larger)

	res := math.MaxInt
	res = min(res, cost(nums, cand1))
	res = min(res, cost(nums, cand2))

	return int64(res)
}

func cost(nums []int, target int) int {
	totalCost := 0
	for _, num := range nums {
		totalCost += abs(num - target)
	}
	return totalCost
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

func nextSmallerOrEqual(s string) string {
	n := len(s)
	arr := []rune(s)

	// STEP 1: mirror left-half to right-half
	for i, j := 0, n-1; i < j; i, j = i+1, j-1 {
		arr[j] = arr[i]
	}

	// allow equal
	if string(arr) <= s {
		return string(arr)
	}

	// STEP 2: mirror (left-half - 1) to right-half + handle 借位
	carry := 1
	for i := (n - 1) / 2; i >= 0; i-- {
		d := int(arr[i]-'0') - carry
		if d >= 0 {
			arr[i] = rune(d + '0')
			carry = 0
		} else {
			arr[i] = '9'
			carry = 1
		}
		// mirroring
		arr[n-i-1] = arr[i]
	}

	if arr[0] == '0' && n > 1 {
		return strings.Repeat("9", n-1)
	} else {
		return string(arr)
	}
}

func nextGreaterOrEqual(s string) string {
	n := len(s)
	arr := []rune(s)

	// STEP 1: mirror left-half to right-half
	for i, j := 0, n-1; i < j; i, j = i+1, j-1 {
		arr[j] = arr[i]
	}
	// allow equal
	if string(arr) >= s {
		return string(arr)
	}

	// STEP 2: mirror (left-half + 1) to right-half + handle 进位
	carry := 1
	for i := (n - 1) / 2; i >= 0; i-- {
		d := int(arr[i]-'0') + carry
		if d <= 9 {
			arr[i] = rune(d + '0')
			carry = 0
		} else {
			arr[i] = '0'
			carry = 1
		}
		// mirroring
		arr[n-i-1] = arr[i]
	}

	if carry == 1 {
		newS := bytes.Repeat([]byte{'0'}, n+1)
		newS[0], newS[n] = '1', '1'
		return string(newS)
	} else {
		return string(arr)
	}
}
