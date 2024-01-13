package leetcode

import (
	"bytes"
	"strconv"
	"strings"
)

func nearestPalindromic(n string) string {
	num, _ := strconv.Atoi(n)

	small := nextSmaller(n)
	large := nextGreater(n)

	sml, _ := strconv.Atoi(small)
	lrg, _ := strconv.Atoi(large)

	if num-sml <= lrg-num {
		return small
	} else {
		return large
	}
}

func nextSmaller(s string) string {
	n := len(s)
	arr := []rune(s)

	// mirror first-half to second-half
	for i, j := 0, n-1; i < j; i, j = i+1, j-1 {
		arr[j] = arr[i]
	}

	if string(arr) < s {
		return string(arr)
	}

	// mirror (first-half - 1) to second half + handle 退位
	// e.g. [10]00 -> [09]99 -> 999
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
		arr[n-1-i] = arr[i]
	}

	// chop off leading 0
	// special case to keep leading 0: 1 -> 0
	if arr[0] == '0' && n > 1 {
		return strings.Repeat("9", n-1)
	} else {
		return string(arr)
	}
}

func nextGreater(s string) string {
	n := len(s)
	arr := []rune(s)

	for i, j := 0, n-1; i < j; i, j = i+1, j-1 {
		arr[j] = arr[i]
	}

	if string(arr) > s {
		return string(arr)
	}

	// mirror (first-half + 1) to second half + handle 进位
	// e.g. [89]1 -> [90]9 -> 909
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
		arr[n-1-i] = arr[i]
	}

	if carry == 1 {
		temp := bytes.Repeat([]byte{'0'}, n+1)
		temp[0], temp[n] = '1', '1'
		return string(temp)
	} else {
		return string(arr)
	}
}
