package leetcode

import "strconv"

func evalRPN(tokens []string) int {
	stack := make([]int, 0)

	for _, t := range tokens {
		v, err := strconv.Atoi(t)
		if err == nil { // current char is a number
			stack = append(stack, v)
		} else {
			num1, num2 := stack[len(stack)-2], stack[len(stack)-1]
			stack = stack[:len(stack)-2]
			switch t {
			case "+":
				stack = append(stack, num1+num2)
			case "-":
				stack = append(stack, num1-num2)
			case "*":
				stack = append(stack, num1*num2)
			case "/":
				stack = append(stack, num1/num2)
			}
		}
	}
	return stack[0]
}

// Demo: 2 1 + 3 *
// | 2 1
// |    2+1
// | 3
// | 3 3
// |    3*3
// | 9

// Demo: 4 13 5 / +
// | 4 13 5
// | 4   13 / 5
// | 4 2
// |    4 + 2
// | 6
