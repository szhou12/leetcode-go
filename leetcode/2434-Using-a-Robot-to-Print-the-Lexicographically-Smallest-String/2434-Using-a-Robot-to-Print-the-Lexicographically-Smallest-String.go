package leetcode

import "strings"

func robotWithString(s string) string {
	n := len(s)
	// nextSmallest[i] := the smallest element in s[i : n-1] (inclusive)
	nextSmallest := make([]int, n)
	curSmallest := 26
	for i := n-1; i >= 0; i-- {
		curSmallest = min(curSmallest, int(s[i] - 'a'))
		nextSmallest[i] = curSmallest
	}

	var res strings.Builder
	stack := make([]int, 0) // this is t
	// We want to maintain the property that stack.top is always the smallest element not yet appended to res

	i := 0
	for i < n {
		// We push s[i] into t/stack if stack is empty or stack.top is not yet the smallest, meaning the there is a smaller element in nextSmallest[i]
		if len(stack) == 0 || nextSmallest[i] < stack[len(stack) - 1] {
			stack = append(stack, int(s[i] - 'a'))
			i++
		} else {
			// else, stack.top is already the smallest element not yet appended to res
			// pop it out and append it to res
			// no increment i
			cur := stack[len(stack) - 1]
			stack = stack[:len(stack) - 1]
			res.WriteByte(byte(cur + 'a'))
		}
	}

	// append the rest of elements still in stack
	for len(stack) > 0 {
		cur := stack[len(stack) - 1]
		stack = stack[:len(stack) - 1]
		res.WriteByte(byte(cur + 'a'))
	}

	return res.String()
}

func robotWithString_lc_soln(s string) string {
	stack := make([]int, 0) // this is t
	freq := make([]int, 26) // count of each charater remained in s

	for _, c := range s {
		freq[int(c-'a')]++
	}

	var smallestRemained func(freq []int) int
	smallestRemained = func(freq []int) int {
		for i := 0; i < 26; i++ {
			if freq[i] > 0 {
				return i
			}
		}
		return -1

	}

	var res strings.Builder

	for _, c := range s {
		curCharIdx := int(c - 'a')
		stack = append(stack, curCharIdx)

		freq[curCharIdx]--

		for len(stack) > 0 && stack[len(stack)-1] <= smallestRemained(freq) {
			cur := stack[len(stack)-1]
			stack = stack[:len(stack) - 1]
			res.WriteByte(byte(cur + 'a'))

		}

	}

	for len(stack) > 0 {
		cur := stack[len(stack) - 1]
		stack = stack[:len(stack) - 1]
		res.WriteByte(byte(cur + 'a'))
	}

	return res.String()
}