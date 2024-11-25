package leetcode

func appealSum(s string) int64 {
	n := len(s)

	lastPositions := make([]int, 26) // last appearance (index) of a character (a-z)
	for c := 0; c < 26; c++ {
		lastPositions[c] = -1 // init to -1 when no appearance
	}

	res := 0
	/*
	  X  X  X  a [X  a  X  X  a  X] 
	          lI ----i------------- n
	*/      
	for i := 0; i < n; i++ {
		lastIdx := lastPositions[int(s[i]-'a')]
		res += (i - lastIdx) * (n - i)
		// update last appearance
		lastPositions[int(s[i]-'a')] = i
	}

	return int64(res)
}
