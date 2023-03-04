package leetcode

// nextSmaller[i] := index of next element > ith element
// prevSmaller[i] := index of previous element >= ith element (必须用 >= 去重)
// 2 8 5 6 5 [7 5 5 5 5 6] 2
//         a    i          b
func sumSubarrayMins(arr []int) int {
	n := len(arr)

	nextSmaller := make([]int, n)
	for i := 0; i < n; i++ {
		nextSmaller[i] = n // 找不到的情况, 定为最右边的下一个, 也就是 index=n
	}
	stack := make([]int, 0)
	for i := 0; i < n; i++ { // i 跑得快, 作为nextSmaller candidate
		for len(stack) > 0 && arr[stack[len(stack)-1]] > arr[i] { // 什么时候开始标记和出栈: 栈顶元素严格大于的时候
			nextSmaller[stack[len(stack)-1]] = i
			stack = stack[:len(stack)-1]
		}
		stack = append(stack, i)
	}

	prevSmaller := make([]int, n)
	for i := 0; i < n; i++ {
		prevSmaller[i] = -1 // 找不到的情况, 定为最左边的前一个, 也就是 index=-1
	}
	stack = make([]int, 0)
	for i := n - 1; i >= 0; i-- { // i 跑得快, 跑在前面, 作为prevSmaller candidate
		for len(stack) > 0 && arr[stack[len(stack)-1]] >= arr[i] { // 什么时候开始标记和出栈: 栈顶元素大于等于的时候
			prevSmaller[stack[len(stack)-1]] = i
			stack = stack[:len(stack)-1]
		}
		stack = append(stack, i)
	}

	M := int(1e9 + 7)
	res := 0
	for i := 0; i < n; i++ {
		a := prevSmaller[i]
		b := nextSmaller[i]
		count := (i - a) * (b - i)
		res += arr[i] * count % M
	}
	return res % M
}
