package leetcode

func totalStrength(strength []int) int {
	M := int(1e9 + 7)
	n := len(strength)

	// Step 1: prepare Prefix Sum
	strength = append([]int{0}, strength...)
	presum := make([]int, n+2)
	presum2 := make([]int, n+2)

	for i := 1; i <= n; i++ {
		presum[i] = (presum[i-1] + strength[i]) % M
		presum2[i] = (presum2[i-1] + strength[i]*i) % M
	}

	// Step 2: Stack to find Next Smaller index & Prev Smaller index
	stack := make([]int, 0)
	nextSmaller := make([]int, n+2)
	prevSmaller := make([]int, n+2)
	for i := range nextSmaller {
		nextSmaller[i] = n + 1
	}

	for i := 1; i <= n; i++ {
		for len(stack) > 0 && strength[stack[len(stack)-1]] > strength[i] {
			topIdx := stack[len(stack)-1]
			nextSmaller[topIdx] = i
			stack = stack[:len(stack)-1]
		}
		if len(stack) > 0 {
			prevSmaller[i] = stack[len(stack)-1]
		}

		stack = append(stack, i)
	}

	// Step 3: Subarray Sum
	res := 0
	for i := 1; i <= n; i++ {
		a := prevSmaller[i]
		b := nextSmaller[i]
		x := i - a
		y := b - i

		left := (presum2[i-1] - presum2[a]) - a * (presum[i-1] - presum[a])
		left = (left + M) % M // + M bc substraction may lead to negative

		right := b * (presum[b-1] - presum[i]) - (presum2[b-1] - presum2[i])
		right = (right + M) % M

		mid := strength[i] * x * y % M

		first := (left * y) % M
		second :=(right * x) % M
		total := (first + mid + second) * strength[i] % M

		res = (res + total) % M
	}

	return res

}

/*
除去 modulo, 保留算法逻辑本身
*/
func totalStrength_NoMod(strength []int) int {
	n := len(strength)

	// Step 1: prepare Prefix Sum
	strength = append([]int{0}, strength...) // only [1:n] (inclusive) have meaning
	presum := make([]int, n+2)               // 0...n+1 only [1:n] (inclusive) have meaning
	presum2 := make([]int, n+2)              // 新式前缀和

	for i := 1; i <= n; i++ {
		presum[i] = presum[i-1] + strength[i]
		presum2[i] = presum2[i-1] + strength[i]*i
	}

	// Step 2: Stack to find Next Smaller index & Prev Smaller index
	stack := make([]int, 0)
	nextSmaller := make([]int, n+2) // nextSmaller[i] = n+1 means i-th element has no next smaller
	prevSmaller := make([]int, n+2) // prevSmaller[i] = 0 means i-th element has no prev smaller
	for i := range nextSmaller {
		nextSmaller[i] = n + 1
	}
	// stack 维护递增元素: 遇到更小的元素，就 pop 出来
	for i := 1; i <= n; i++ {
		// 寻找stack顶元素的next smaller
		for len(stack) > 0 && strength[stack[len(stack)-1]] > strength[i] {
			topIdx := stack[len(stack)-1]
			nextSmaller[topIdx] = i
			stack = stack[:len(stack)-1] // pop
		}
		// pop干净后，现在stack顶元素恰好是当前i-th元素的prev smaller or equal
		if len(stack) > 0 {
			prevSmaller[i] = stack[len(stack)-1]
		}

		stack = append(stack, i)
	}

	// Step 3: Subarray Sum
	// a X X X X i X X X b
	//   ---x---   --y--
	res := 0
	// Given i, find all subarrays whose weakest element is strength[i]
	for i := 1; i <= n; i++ {
		a := prevSmaller[i]
		b := nextSmaller[i]
		x := i - a
		y := b - i

		// left := subarray sum value of [a+1:i-1] (inclusive)
		// left will be included y times
		left := (presum2[i-1] - presum2[a]) - (a * (presum[i-1] - presum[a]))

		// right := subarray sum value of [i+1:b-1] (inclusive)
		// right will be included x times
		right := (b * (presum[b-1] - presum[i])) - (presum2[b-1] - presum2[i])

		// mid := i-th element value
		// mid will be included x*y times
		mid := strength[i]

		res += (left * y + mid * x * y + right * x) * strength[i]
	}

	return res

}
