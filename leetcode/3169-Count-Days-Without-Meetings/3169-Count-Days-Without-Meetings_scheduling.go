package leetcode

import "sort"

func countDays_scheduling(days int, meetings [][]int) int {
	// sort by start date, prioritize larger end date (其实不需要考虑end date顺序，这个算法也正确)
	sort.Slice(meetings, func(i, j int) bool {
		if meetings[i][0] == meetings[j][0] {
			return meetings[i][1] > meetings[j][1]
		}
		return meetings[i][0] < meetings[j][0]
	})

	res := 0
	curDate := 1
	for _, meeting := range meetings {
		start, end := meeting[0], meeting[1]
		if curDate < start {
			res += start - curDate
		}
		if curDate <= end { // 跳到下一次potentially"好日子"的开始
			curDate = end + 1
		}
	}

	if curDate <= days {
		res += days - curDate + 1
	}
	return curDate
}
