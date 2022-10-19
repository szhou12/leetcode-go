package leetcode

import "strconv"

func compress(chars []byte) int {
	// Edge case
	if len(chars) == 1 {
		return len(chars)
	}

	slow := 0 // 物理意义：slow左边不包含slow都是已经compress好的
	curChar := chars[slow]
	count := 1

	// 注意：
	// fast 从 1 开始, 这是因为count初始值=1, 如果从0开始, 第一个if-condition会导致第一个group多算一个
	// fast 一直循环到 len(chars), 这是为了算到最后一个group, <len(chars)最后一个group无法在这个loop里算到
	for fast := 1; fast <= len(chars); fast++ {
		if fast < len(chars) && chars[fast] == curChar {
			count++ // 数 相同的char遇到了多少个了
		} else { // fast == len(chars) || chars[fast] != curChar 结算时间
			// Step 1: slow继承 curChar, 并推进slow一格
			chars[slow] = curChar
			slow++
			// Step 2: 如果count > 1, 把count加在后面.
			// 注意: 12 占两格 "1", "2" --> Go 中, loop string 来实现
			if count > 1 {
				countStr := strconv.Itoa(count)
				for j := 0; j < len(countStr); j++ {
					chars[slow+j] = countStr[j]
				}
				// Step 3: slow跳过append数字的格子, slow跳到此时fast的位置
				slow += len(countStr)
			}
			// Step 4: 如果fast还没越界，curChar换新，换成现在fast指向的
			if fast < len(chars) {
				curChar = chars[fast]
			}
			// Step 5: 重置count=1
			count = 1

		}
	}
	// 根据slow物理意义, 此时slow就代表了compressed string的length
	return slow
}
