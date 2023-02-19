package leetcode

import "math"

/**
left[i] := Rightmost index r for the shortest prefix of s (i.e. s[0:r]) containing t[0:i]
right[j] := Leftmost index l for the shortest suffix of s (ie. s[l:m-1]) containing t[j:n-1]
*/
func minimumScore(s string, t string) int {
	m := len(s)
	n := len(t)

	// prefix of s
	left := make([]int, n)
	sidx := 0 // pointer to s
	for i := 0; i < n; i++ {
		left[i] = math.MaxInt             // 预设无限大表示找不到包含关系的情况, 之后如果可以包含再更新
		for sidx < m && s[sidx] != t[i] { // 没有越界 并且 没找到match的时候 继续走s的pointer
			sidx++
		}
		if sidx < m { // 出界前找到包含关系
			left[i] = sidx
			sidx++
		}
	}

	// suffix of s
	right := make([]int, n)
	sidx = m - 1
	for j := n - 1; j >= 0; j-- {
		right[j] = math.MinInt
		for sidx >= 0 && s[sidx] != t[j] {
			sidx--
		}
		if sidx >= 0 { // 出界前找到包含关系
			right[j] = sidx
			sidx--
		}
	}

	// middle part to be deleted - Binary Search
	// 挡板物理意义: 猜删除的长度, 所以low=删除长度最小值=0=什么都不删, high=删除长度最大值=n=全删
	low, high := 0, n
	for low < high {
		mid := low + (high-low)/2
		if isScoreOK(s, t, left, right, mid) { // mid长度目前是一个可行解, 题目求最短删除长度(minimum score), 再试试短一点呢
			high = mid
		} else {
			low = mid + 1
		}
	}
	return low
}

/*
t:   x x x  {y   y   y}        z z z
             i------i+score-1
以score为长度的sliding window滑一遍t, 看对应的left[i-1], right[i+score-1]有没有在s上overlap
*/
func isScoreOK(s string, t string, left []int, right []int, score int) bool {
	n := len(t)

	// i的物理意义: 以score为长度的sliding window的左端点 (ie. 要删除的subarray的左端点)
	// 注意: 下面这样写的for-loop，里面的if会发生越界 (当sliding window分别顶在t最左边和最右边时)
	// for i := 0; i+score-1 < n; i++ {
	// 	if left[i-1] < right[i+score] {
	// 		return true
	// 	}
	// }

	// sliding window顶在t最左边, 不看left[], 只看right[]
	if right[0+score] != math.MinInt {
		return true
	}

	// sliding window顶在t最右边, 不看right[], 只看left[]
	// n-score是刨除掉sliding window, (n-score)-1是sliding window左端点的前一个index
	if left[n-score-1] != math.MaxInt {
		return true
	}

	// sliding window顶边的时候, 同时看left[]和right[]不重叠
	for i := 1; i+score < n; i++ {
		if left[i-1] < right[i+score] {
			return true
		}
	}

	return false

}
