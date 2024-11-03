package amz

/*
Given two string arrays of length n: new[] and old[].
new[i] is similar to old[i] if by applying the following operation to some chars in new[i], old[i] becomes a subseqnce of new[i].
Operation: take any char and transform it to its next char in the alphabet. e.g. 'a' -> 'b', 'b' -> c, ..., 'z' -> 'a'.
Return an array of length n: res[] where res[i] = true if new[i] is similar to old[i], otherwise false.

Example:
Input:  new = ["aaccbbee", "aab"], old = ["bdbf", "aee"]

a[ac]cbb[e]e -> a[bd]cbb[f]e
"abdcbbfe" has "bdbf" as a subsequence

Output: [true, false]
*/

/*
Solution: Two Pointers
newPtr: pointer to new[i]
oldPtr: pointer to old[i]

Keep moving newPtr forward.
Only move oldPtr forward iff 1. new[newPtr] == old[oldPtr] OR 2. transform(new[newPtr]) == old[oldPtr].
Return true if when newPtr finishes, oldPtr also finishes. Otherwise, return false.

Time Complexity = O(mn) where m is the length of new[i] and n is the length of new[].
*/

func CheckSimilarStrTransform(new []string, old []string) []bool {
	n := len(new)
	res := make([]bool, 0)

	isSimilar := func(newStr, oldStr string) bool {
		newPointer := 0
		oldPointer := 0

		for newPointer < len(newStr) && oldPointer < len(oldStr) {
			if newStr[newPointer] == oldStr[oldPointer] || transform(newStr[newPointer]) == oldStr[oldPointer] {
				newPointer++
				oldPointer++
			} else {
				newPointer++
			}
		}
		return oldPointer >= len(oldStr)
	}

	for i := 0; i < n; i++ {
		res = append(res, isSimilar(new[i], old[i]))
	}
	return res
}

func transform(char byte) byte {
	if char == 'z' {
		return 'a'
	}
	return char + 1
}
