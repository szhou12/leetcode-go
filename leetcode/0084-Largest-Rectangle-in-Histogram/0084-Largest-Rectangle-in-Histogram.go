package leetcode

// Stack - prevSmaller, nextSmaller
func largestRectangleArea(heights []int) int {
    n := len(heights)

    stack := make([]int, 0)
    // find nextSmaller[i] = j where heights[j] < heights[i] and j > i
    nextSmaller := make([]int, n)
    for i := 0; i < n; i++ {
        nextSmaller[i] = n
    }

    for i := 0; i < n; i++ {
        for len(stack) > 0 && heights[stack[len(stack)-1]] > heights[i] {
            nextSmaller[stack[len(stack) - 1]] = i
            stack = stack[:len(stack)-1]
        }
        stack = append(stack, i)
    }

    // clear stack
    stack = []int{}

    prevSmaller := make([]int, n)
    for i := 0; i < n; i++ {
        prevSmaller[i] = -1
    }

    for i := n-1; i>=0; i-- {
        for len(stack) > 0 && heights[stack[len(stack)-1]] > heights[i] {
            prevSmaller[stack[len(stack)-1]] = i
            stack = stack[:len(stack) - 1]
        }
        stack = append(stack, i)
    }

    res := -1
    for i := 0; i < n; i++ {
        cur := heights[i] * (nextSmaller[i] - prevSmaller[i] - 1)
        res = max(res, cur)
    }
    return res

}


// DP
func largestRectangleArea_old(heights []int) int {
	n := len(heights)
	lmin := make([]int, n) // lmin[i]: index of first bar [i-1...0] shorter than bar i
	rmin := make([]int, n) // rmin[i]: index of firt bar [i+1...n-1] shorter than bar i

	lmin[0] = -1
	for i := 1; i < n; i++ {
		tempIdx := i - 1
		for tempIdx >= 0 && heights[tempIdx] >= heights[i] {
			tempIdx = lmin[tempIdx]
		}
		lmin[i] = tempIdx
	}

	rmin[n-1] = n
	for i := n - 1; i >= 0; i-- {
		tempIdx := i + 1
		for tempIdx < n && heights[tempIdx] >= heights[i] {
			tempIdx = rmin[tempIdx]
		}
		rmin[i] = tempIdx
	}

	res := 0
	for i := 0; i < n; i++ {
		curArea := heights[i] * (rmin[i] - lmin[i] - 1) // 因为宽要去掉lmin和rmin
		res = max(res, curArea)
	}
	return res
}

func max(a int, b int) int {
	if a > b {
		return a
	}
	return b
}
