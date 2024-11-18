package leetcode

/*
Greedy + Simulation
O(n^2)
*/
func minMovesToMakePalindrome_sim(s string) int {
	n := len(s)
	str := []rune(s)
	count := 0 // # of rightmost places have been swapped already so that that they cannot be used

	res := 0

	// for any left-hand side i, find its pair from the right-hand side
	for i := 0; i < n/2; i++ {
		j := n - 1 - count
		for str[i] != str[j] {
			j--
		}
		if i != j {
			// swap j by k times to the 'rightmost' place
			k := n - 1 - count - j
			res += k
			for ; k > 0; k-- {
				str[j], str[j+1] = str[j+1], str[i] // swap
				j++
			}
			count++ // one more 'rightmost' place has been occupied
		} else {
			// if i == j, meaning the i-th element is alone and cannot be paired up
			/*
			这里的写法如何达到“先忽略奇数元素”的效果？
			“先忽略”的意思是剩下的元素排列成回文串。
			也就是说下一个元素i+1是固定在i+1位子去找配对的右边元素j
			      i+1    j
			X X w [a a b a] X X
			*/
			res += n/2 - j
		}
	}

	return res
}
