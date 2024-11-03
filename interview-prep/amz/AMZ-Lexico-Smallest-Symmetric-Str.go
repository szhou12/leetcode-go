package amz

import "sort"

/*
Given a symmetrical string s, find the lexicographically smallest symmetric string that can be obtained by re-arranging s.

Example:
s = "yxxy"
Output: "xyyx"
*/


/*
Idea: 切一半，前一半sort，后一半reverse，然后拼接

Time complexity = O(nlogn)
*/
func LexicoSmallestSymmetricStr(s string) string {
	n := len(s)
	mid := n / 2

	firstHalf := s[:mid]
	firstHalfArr := []rune(firstHalf)

	sort.Slice(firstHalfArr, func(i, j int) bool {
		return firstHalfArr[i] < firstHalfArr[j]
	})

	reFirstHalf := string(firstHalfArr)

	// reverse
	for i, j := 0, len(firstHalfArr)-1; i < j; i, j = i+1, j-1 {
		firstHalfArr[i], firstHalfArr[j] = firstHalfArr[j], firstHalfArr[i]
	}

	reSeondHalf := string(firstHalfArr)

	if n % 2 == 0 {
		return reFirstHalf + reSeondHalf
	}
	return reFirstHalf + string(s[mid]) + reSeondHalf

}
