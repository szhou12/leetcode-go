package mathematics

import (
	"fmt"
	"strconv"
)

// Generate all combinations (states) to select k objects out of n objects = C(n, k)
// 生成n元集合所有k元子集的算法
func GospersHack(n int, k int) {

	// init state with the lowest k bits set to 1
	// e.g. n=4, k=2 --> 1<<k = 0100, state = 0011
	state := (1 << k) - 1

	for state < (1 << n) { // state < 2^n. Note: 2^n-1 is the combination when all n objects selected 111...111

		// do something with current state
		fmt.Println(strconv.FormatInt(int64(state), 2))

		// -state : invert bits, then add 1. e.g. -(0011) = 1100+1 = 1101 
		c := state & -state
		r := state + c
		state = (((r ^ state) >> 2) / c) | r
	}
}
