package leetcode

// 6   1   7   4   1
//    i-1  i  i+1
// leftEven[i] := even sum in nums[0:i]
// leftOdd[i] := odd sum in nums[0:i]
// rightEven[i] := even sum in nums[i:n-1]
// rightOdd[i] := odd sum in nums[i:n-1]
func waysToMakeFair(nums []int) int {
	// Edge Cases
	if len(nums) == 1 {
		return 1
	}
	if len(nums) == 2 {
		return 0
	}

	n := len(nums)

	leftEven := make([]int, n)
	leftOdd := make([]int, n)
	evenSum, oddSum := 0, 0
	for i := 0; i < n; i++ {
		if i%2 == 0 { // even index
			evenSum += nums[i]
		} else { // odd index
			oddSum += nums[i]
		}
		leftEven[i] = evenSum
		leftOdd[i] = oddSum
	}

	rightEven := make([]int, n)
	rightOdd := make([]int, n)
	evenSum, oddSum = 0, 0
	for i := n - 1; i >= 0; i-- {
		// i号元素被移除后, R段整体前移, 所以, 原来i=even, 前移后, i=odd, 所以加在odd sum头上
		if i%2 == 0 { // even index => odd index
			oddSum += nums[i]
		} else { // odd index => even index
			evenSum += nums[i]
		}
		rightEven[i] = evenSum
		rightOdd[i] = oddSum
	}

	res := 0
	// 顶到最右边(i=n-1), 只有L段[0:n-2], 没有R段
	if leftEven[n-2] == leftOdd[n-2] {
		res += 1
	}
	// 顶到最左边(i=0), 只有R段[1:n-1], 没有L段
	if rightEven[1] == rightOdd[1] {
		res += 1
	}
	// 6   1   7   4   1
	//    i-1  i  i+1
	for i := 1; i <= n-2; i++ {
		if leftEven[i-1]+rightEven[i+1] == leftOdd[i-1]+rightOdd[i+1] {
			res += 1
		}
	}

	return res

}
