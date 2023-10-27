package leetcode

var M = int(1e9 + 7)

func numberOfWays(s string, t string, k int64) int {
	// Step 1: find p = # of matches in n shifts
	// cyclic shift
	ss := s + s
	// cyclic shift trick: remove last char
	ss = ss[:len(ss)-1]
	// KMP
	p := strStr(ss, t)

	// Step 2: state transformation
	// n shifts
	n := len(s)
	T := []int{n - p - 1, n - p, p, p - 1}
	// Recursion
	Tk := quickMul(T, k)

	if s == t {
		return Tk[3] // Tk * (0, 1)'
	} else {
		return Tk[2] // Tk * (1, 0)'
	}
}

func quickMul(mat []int, k int64) []int {
	// base case
	if k == 0 {
		return []int{1, 0, 0, 1}
	}

	mat2 := quickMul(mat, k/2)
	if k%2 == 0 {
		return multiply(mat2, mat2)
	} else {
		return multiply(multiply(mat2, mat2), mat)
	}
}

func multiply(mat1 []int, mat2 []int) []int {
	// a b    x y
	// c d    z w
	a, b, c, d := mat1[0], mat1[1], mat1[2], mat1[3]
	x, y, z, w := mat2[0], mat2[1], mat2[2], mat2[3]

	return []int{
		(a*x + b*z) % M,
		(a*y + b*w) % M,
		(c*x + d*z) % M,
		(c*y + d*w) % M,
	}
}

func strStr(target string, pattern string) int {
	count := 0
	n, m := len(target), len(pattern)

	// Step 1: self-match pattern
	// lsp[i] = max length k s.t. pattern[0:k-1] == pattern[i-k+1:i]
	lsp := preprocess(pattern)

	// Step 2: pattern matching
	// dp[i] = max length k s.t. pattern[0:k-1] == target[i-k+1:i]
	dp := make([]int, n)
	// base case
	if pattern[0] == target[0] {
		dp[0] = 1
	}
	if dp[0] == 1 && m == 1 {
		count++
	}
	// recurrence
	for i := 1; i < n; i++ {
		k := dp[i-1]
		// Note: MUST check k >= len(pattern)
		for k > 0 && (k >= len(pattern) || pattern[k] != target[i]) {
			k = lsp[k-1]
		}
		if pattern[k] == target[i] {
			dp[i] = k + 1
		} else {
			dp[i] = k
		}
		if dp[i] == m {
			count++
		}
	}
	return count
}

func preprocess(pattern string) []int {
	n := len(pattern)

	dp := make([]int, n)

	// base case
	dp[0] = 0

	// recurrence
	for i := 1; i < n; i++ {
		k := dp[i-1]
		for k > 0 && pattern[k] != pattern[i] {
			k = dp[k-1]
		}
		if pattern[k] == pattern[i] {
			dp[i] = k + 1
		} else {
			dp[i] = k
		}
	}
	return dp
}


// 1. shift s is like 扑克切牌 => rotate string.
// round 1: +1, +2, +3, ..., +(n-1) // +x = move suffix of x to the front
// round 2: +1, +2, +3, ..., +(n-1)
// round 3: +1, +2, +3, ..., +(n-1)
// ...
/// round k: +1, +2, +3, ..., +(n-1)
// Time = O((n-1)^k)
// 2. assume s' is s after k shifts, s' can be in many forms: 
//   s' = s(0), s(1), s(2), ..., s(n-1) where s(x) means shifting the suffix of s starting at index x
// then how do we know if t is among them? => double the string s to find cyclic shift
// trick: bc s(0) == s(n-1) == no shift at all (will cause double counting), cutoff the last char when double the string s. e.g. s = abababa => 2s-1 = ababab ababa
// find the number of appearances of t (pattern string) in 2s-1 (target string) => KMP
// 3. find all possible k rounds that lead to each form of s' is time-consuming (it follows normal distribution)! We need to think a faster way
// => Assume there are p forms of s' that equal to t. p is calculated by KMP
// f[j] = # of bad strings (s' != t) after j rounds of shift
// g[j] = # of good strings (s' == t) after j round of shift
// how we associate f[j-1], g[j-1] with f[j], g[j]???
// n forms after shift, among which p forms (good shift) == t
// -1 bc we cannot shift nothing in one round
// f[j] = f[j-1] * (n-p-1) + g[j-1] * (n-p)
// meaning: at j round, # of bad strings = j-1 round of bad strings keep being bad (cannot shift nothing) + j-1 round of good strings turn into bad
// why multiply? eg. f[j-1] * (n-p-1)?
// bc each bad string of f[j-1] after one round of shift will have n forms, among which p forms are good
// g[j] = f[j-1] * (p) + g[j-1] * (p-1)
// meaning: at j round, # of good string = j-1 round of bad strings turn into good + j-1 round of good strings keep being good (cannot shift nothing)
// 4. Now can you think of this state transformation as matrix mult? 
// (f[j]) = |n-p-1 n-p| (f[j-1])
// (g[j]) = |  p   p-1| (g[j-1])
// ie. v' = T*v
// apply Associative property
// v_k = T^k * v_0
// 5. Recursion to find T^k = (T^k/2)^2  (k even) = ... = T^0 = (1 0)
//                      T^k = (T^k/2)^2*T (k odd)               (0 1)
// 6. f[0], g[0] = (0 1) if s == t before any shifts
//               = (1 0) if s != t before any shifts