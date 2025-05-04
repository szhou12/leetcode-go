package dfs

/*
 * All Valid Permutations Of Parentheses II
 *
 * Get all valid permutations of l pairs of (), m pairs of [] and n pairs of {}.
 *
 * Assumptions
 * l, m, n >= 0
 * l + m + n > 0
 *
 * Examples
 * l = 1, m = 1, n = 0, all the valid permutations are ["()[]", "([])", "[()]", "[]()"]
 */
func validParenthesesII(l int, m int, n int) []string {
	var result []string
	var subset []rune
	brackets := []rune{'(', ')', '[', ']', '{', '}'}
	numRemained := []int{l, l, m, m, n, n}
	targetLength := 2*l + 2*m + 2*n
	stack := make([]rune, 0)

	findValidParenthesesII(targetLength, subset, brackets, numRemained, stack, &result)
	return result
}

func findValidParenthesesII(targetLength int, subset []rune, brackets []rune, numRemained []int, stack []rune, result *[]string) {
	// base case
	if len(subset) == targetLength {
		subsetCopy := make([]rune, len(subset))
		copy(subsetCopy, subset)
		*result = append(*result, string(subsetCopy))
		return
	}

	// branches = loop through each element of numRemained
	for i := 0; i < len(numRemained); i++ {
		if i%2 == 0 { // left bracket
			if numRemained[i] > 0 {
				subset = append(subset, brackets[i])
				stack = append(stack, brackets[i])
				numRemained[i]--

				findValidParenthesesII(targetLength, subset, brackets, numRemained, stack, result)

				//backtracking
				subset = subset[:len(subset)-1]
				stack = stack[:len(stack)-1]
				numRemained[i]--
			} else { // right bracket
				if len(stack) > 0 && stack[len(stack)-1] == brackets[i-1] {
					subset = append(subset, brackets[i])
					stack = stack[:len(stack)-1]
					numRemained[i]--

					findValidParenthesesII(targetLength, subset, brackets, numRemained, stack, result)

					//backtracking
					subset = subset[:len(subset)-1]
					stack = append(stack, brackets[i-1])
					numRemained[i]++
				}
			}
		}
	}
}
