package leetcode

import "sort"

func partitionLabels(s string) []int {
	positionList := findPositions(s) // O(n)

	// sort by starting index (left) in increasing order
	sort.Slice(positionList, func(i, j int) bool { // O(mlogm)
		return positionList[i][0] < positionList[j][0]
	})

	res := make([]int, 0)
	left := positionList[0][0]
	right := positionList[0][1]
	for i := 0; i < len(positionList); i++ { // O(m)
		// 一旦下一区间左边界大于当前右边界，即可认为出现分割点
		// find a new partition
		if positionList[i][0] > right {
			// conclude 目前的 partition
			pLength := right - left + 1
			res = append(res, pLength)
			// update left
			left = positionList[i][0]
		}
		right = max(right, positionList[i][1])
	}

	// the last partition hasn't been added yet!!!
	res = append(res, right-left+1)
	return res

}

func findPositions(s string) [][]int {
	res := make([][]int, 0)

	dict := make(map[rune][]int) // map each char -> [start idx, end idx]
	for idx, char := range s {
		if _, ok := dict[char]; !ok {
			// if first time seeing current char, adding current index as BOTH start idx and end idx
			dict[char] = make([]int, 2)
			dict[char][0] = idx
			dict[char][1] = idx // MUST have this!!!
		} else {
			// if not, update end idx
			dict[char][1] = idx
		}
	}

	for _, pos := range dict {
		res = append(res, pos)
	}
	return res
}

func max(a int, b int) int {
	if a > b {
		return a
	}
	return b
}
