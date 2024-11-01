package amz

import "fmt"

/*
Given a binary string that reprensents a sequence of "0" and "1".
Given k opertaions that in one operation you can switch "0" -> "1" or "1" -> "0".
Find the longest substring of alternating binary digits that can be obtained by performing at most k operations.

Example:
s = "1001", k=1
"1001" -> "1011" -> "[101]1"

Output: 3
*/

func MaxSwitchingDigits(s string, k int) int {
	pattern1 := maxAlternateSubstr(s, k, '0') // "010101"

	fmt.Println("+++++++++++++++++")

	pattern2 := maxAlternateSubstr(s, k, '1') // "101010"
	return max(pattern1, pattern2)
}

func maxAlternateSubstr(s string, k int, start byte) int {
	n := len(s)
	res := 0
	right := 0
	flips := 0

	for left := 0; left < n; left++ {

		for right < n && (s[right] == expectedCharAt(right, left, start) || flips < k) {
			if s[right] != expectedCharAt(right, left, start) {
				flips++
			}
			right++
		}

		fmt.Printf("update - l: %v, r: %v, flips: %v\n", left, right, flips)

		res = max(res, right-left)

		if left > right {
			if s[left] != expectedCharAt(left, left, start) {
				flips--
			}
		} else {
			right = left + 1 // bring right back will cause O(n^2)???
		}

		fmt.Printf("postpo - l: %v, r: %v, flips: %v\n\n", left, right, flips)

	}
	return res
}

func expectedCharAt(pos int, left int, startChar byte) byte {
	if (pos-left)%2 == 0 {
		return startChar
	}
	return flipChar(startChar)
}

func flipChar(c byte) byte {
	if c == '0' {
		return '1'
	}
	return '0'
}
