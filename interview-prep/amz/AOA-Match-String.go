package amz

/*
Given two arrays of length n: text and pattern
Each string in pattern contains exactly * character
a * char is a wild card that can match any sequence (zero or more) of lowercase English letters
A pattern matches a string if we can replace the * with some expression
For a string text[i], from 0 to n-1, find out if it matches patterns[i]

Example:
n = 2, text = ['code', 'coder'], pattern = ['co*d', 'co*er']
Output: [false, true]
*/


/*
Key Observation:
Only 3 places to place '*': first place, somewhere in the middle, last place
	e.g. abcde -> *cde, a*de, abc*
'*' char splits pattern[i] into 3 parts: [prefix of size p], '*', [suffix of size q]
Solution: compare first p chars between text[i] and pattern[i] && compare last q chars between text[i] and pattern[i]
*/
func matchString(text []string, pattern []string) []bool {
	n := len(text)
	res := make([]bool, 0)

	for i := 0; i < n; i++ {
		res = append(res, isMatch(text[i], pattern[i]))
	}
	return res
}

func isMatch(text, pattern string) bool {
	n, m := len(text), len(pattern)

	// early return: if text has less chars than pattern excluding '*'
	if n < m-1 {
		return false
	}

	// find the index of '*
	index := 0
	for i := 0; i < m; i++ {
		if pattern[i] == '*' {
			index = i
			break
		}
	}

	// compare prefix
	if index > 0 && !(text[:index] == pattern[:index]) {
		return false
	}

	// compare suffix
	// length of suffix in pattern: (m-1) - (index+1) + 1 = m - index - 1
	// start index in text s.t. [start, n-1] gives (n-1) - start + 1 = m - index - 1
	// start = n - m + index + 1
	if index < n-1 && !(text[n-m+index+1:] == pattern[index+1:]) {
		return false
	}

	return true
}