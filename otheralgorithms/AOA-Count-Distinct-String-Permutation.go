package otheralgorithms

/*
Given a string, we can select a substring so that when we reverse the substring, we can form a new string.
Find the number of distinct strings that can be formed.

Example:
Input: "abb"
"abb", "bab", "bba"
Output: 3

Input: "abbab"
"abbab": 1
len = 2: 3
len = 3: 2 ([bab] CAN'T)
len = 4: 0
len = 5: 1
Output: 7
*/


/*
Key Observation:
寻找不变量！What is the invariant?
 - if a subtring is a palindrome, then no use to reverse it to form a new string

isPalind[l][r] := true if s[l:r] (inclusive) is a palindrome
*/
func countDistinctStrings(s string) int {
	n := len(s)

	isPalind := make([][]bool, n)
	for i := 0; i < n; i++ {
		isPalind[i] = make([]bool, n)
	}

	totalCnts := 1 // the original string istelf is a distinct string

	// Base case 1: substr length = 1 - single character is a palindrome
	for i := 0; i < n; i++ {
		isPalind[i][i] = true
	}

	// Base case 2: substr length = 2 - two characters are palindrome if they are the same
	for i := 1; i < n; i++ {
		if s[i-1] == s[i] {
			isPalind[i-1][i] = true
		} else {
			totalCnts++ // two characters are distinct so that we can reverse them to form a new string
		}
	}

	// For substring length = 3 ... n
	for length := 3; length <= n; length++ {
		for l := 0; l < n; l++ {
			// length = r - l + 1
			r := length + l - 1

			if r < n {
				// check if substring s[l:r] (inclusive) is a palindrome
				// l [l+1, ..., r-1] r
				if s[l] == s[r] && isPalind[l+1][r-1] {
					isPalind[l][r] = true
				} else {
					// case 1: s[l] != s[r] && isPalind[l+1][r-1] == TRUE
					// e [aba] f => we can reverse
					// case 2: s[l] == s[r] && isPalind[l+1][r-1] == FALSE
					// e [ab] e => we can NOT reverse as [ab] already counted!!!
					// case 3: s[l] != s[r] && isPalind[l+1][r-1] == FALSE
					// e [ac] f => we can reverse
					if s[l] != s[r] {
						totalCnts++
					}
				}
			}
		}
	}

	return totalCnts

}