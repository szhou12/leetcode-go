package leetcode

// Optimal Solution - 单调栈模版
// 重点！！！stack存入元素的index 而不是元素本身！！！
// 这样有index方便计算距离

func dailyTemperatures(temperatures []int) []int {
	res := make([]int, len(temperatures))
	// 非常重要！！！stack里放元素index，而不是元素
	stack := make([]int, 0)

	for i := len(temperatures) - 1; i >= 0; i-- {
		// 挤掉不满足要求的栈顶元素
		for len(stack) > 0 && temperatures[i] >= temperatures[stack[len(stack)-1]] {
			stack = stack[:len(stack)-1]
		}
		// 满足要求的栈顶元素index - 当前元素的index = 我们要求的距离/天数
		if len(stack) > 0 {
			res[i] = stack[len(stack)-1] - i
		} else {
			res[i] = 0
		}
		stack = append(stack, i)

	}
	return res
}

// My Solution - suboptimal
// 用一个 pair struct
// temp 表示 temperatures[i]
// count 表示 temperatures[i]是挤掉了多少个stack里的元素才入栈的
// 写的过于复杂, 算间隔天数 == 算距离 == 用index就可以计算
type pair struct {
	temp  int
	count int
}

func dailyTemperatures_mysoln(temperatures []int) []int {
	res := make([]int, len(temperatures))
	stack := make([]pair, 0)

	for i := len(temperatures) - 1; i >= 0; i-- {
		if len(stack) > 0 && temperatures[i] < stack[len(stack)-1].temp { // 栈顶元素就已经满足条件
			res[i] = 1
			stack = append(stack, pair{temp: temperatures[i], count: 0})
			continue
		}

		curCount := 0
		for len(stack) > 0 && temperatures[i] >= stack[len(stack)-1].temp {
			curCount += 1                         // 挤掉栈顶元素
			curCount += stack[len(stack)-1].count // 同时要加上被挤掉元素它自己挤掉了多少元素
			stack = stack[:len(stack)-1]
		}
		if len(stack) > 0 {
			res[i] = curCount + 1 // + 1 因为是挤掉所有元素之后的下一天
		} else {
			res[i] = 0
		}
		stack = append(stack, pair{temp: temperatures[i], count: curCount})
	}

	return res
}
