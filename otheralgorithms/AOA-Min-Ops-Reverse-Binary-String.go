
package otheralgorithms


/*
Given a binary string: 001011
Operation: pick a character from the string, take it out, and append it to the end of the string
Find the minimum operations to reverse the string

Example: 
001011 --reverse--> 110100
pick 1st '0':       010110
pick 2nd '1':	    001101
pick 1st '0':	    011010
pick 1nd '1':	    110100

Output: 4
*/


/* Key Observation:
1. 寻求不变量: What is the invariant before and after the transformation? 
  - To minimize the # operations is to find the longest matching in the original string that can be the prefix of the reversed string !!!
2. Since we can pick any char from the original string and throw it to the end in any order, the longest matching in the original string doesn't have to be contiguous (i.e. no need to be substring). In other words, we need to find the longest subsequence in the original string that can match the prefix of the reversed string!
Thus, if the longest subsequence is of length k, the min # operations = n - k
*/

func findMinOps(s string) int {
	n := len(s)

	reversed := reverse(s)

	index := 0
	for i := 0; i < n; i++ {
		if s[i] == reversed[index] {
			index++
		}
	}
	// after the for-loop, index is pointing to the next position of the matching prefix
	// therefore, index is the length of the longest matching prefix
	return n-index
}

func reverse(s string) string {
	n := len(s)
	runes := []rune(s)
	for i, j := 0, n-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}