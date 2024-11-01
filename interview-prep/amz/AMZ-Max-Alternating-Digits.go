package amz

/*
Given a binary string that reprensents a sequence of "0" and "1".
Given k opertaions that in one operation you can switch "0" -> "1" or "1" -> "0".
Find the longest substring of alternating binary digits that can be obtained by performing at most k operations.
Constraints: 1 <= k <= n

Example:
s = "1001", k=1
"1001" -> "1011" -> "[101]1"

Output: 3
*/

func MaxSwitchingDigits(s string, k int) int {
	pattern01 := maxAlternateSubstr(s, k, '0') // "010101"
	pattern10 := maxAlternateSubstr(s, k, '1') // "101010"
	return max(pattern01, pattern10)
}


// similar leetcode 2024 solution
func maxAlternateSubstr(s string, k int, start byte) int {
	n := len(s)
	res := 0
	right := 0
	flips := 0

	for left := 0; left < n; left++ {

		// 吃
		for right < n && (s[right] == expectedCharAt(right, left, start) || flips < k) {
			if s[right] != expectedCharAt(right, left, start) {
				flips++
			}
			right++
		}

		// update
		res = max(res, right-left)

		// 吐: 如果左边界指向的字符不是期望pattern的字符，说明我们在该字符上使用了一次operation，现在要吐出来
		if s[left] != expectedCharAt(left, left, start) {
			flips--
		}

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