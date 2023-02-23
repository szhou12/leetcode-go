package leetcode

// leftFlips01[i] := # of flips to make s[0:i] alternate with 01 pattern (leftmost element starts with 0, next with 1)
// leftFlips10[i] := # of flips to make s[0:i] alternate with 10 pattern (leftmost element starts with 1, next with 0)
// rightFlips01[i] := # of flips to make s[i:n-1] alternate with 01 pattern (rightmost element starts with 0, next with 1)
// rightFlips10[i] := # of flips to make s[i:n-1] alternate with 10 pattern (rightmost element starts with 1, next with 0)
func minFlips(s string) int {
	n := len(s)

	leftFlips01 := make([]int, n)
	leftFlips10 := make([]int, n)
	countFlips01 := 0
	countFlips10 := 0
	for i := 0; i < n; i++ {
		if (i%2 == 0 && s[i] == '1') || (i%2 == 1 && s[i] == '0') {
			countFlips01 += 1
		}
		if (i%2 == 0 && s[i] == '0') || (i%2 == 1 && s[i] == '1') {
			countFlips10 += 1
		}
		leftFlips01[i] = countFlips01
		leftFlips10[i] = countFlips10
	}

	rightFlips01 := make([]int, n)
	rightFlips10 := make([]int, n)
	countFlips01 = 0
	countFlips10 = 0
	for i := n - 1; i >= 0; i-- {
		j := n - 1 - i // j = i从右往左数是第几个index, 用来判断奇偶性
		if (j%2 == 0 && s[i] == '1') || (j%2 == 1 && s[i] == '0') {
			countFlips01 += 1
		}
		if (j%2 == 0 && s[i] == '0') || (j%2 == 1 && s[i] == '1') {
			countFlips10 += 1
		}
		rightFlips01[i] = countFlips01
		rightFlips10[i] = countFlips10
	}

	res := min(leftFlips01[n-1], leftFlips10[n-1])
	for i := 0; i+1 < n; i++ {
		res = min(res, leftFlips01[i]+rightFlips10[i+1])
		res = min(res, leftFlips10[i]+rightFlips01[i+1])
	}
	return res
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
