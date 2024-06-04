package leetcode

import "sort"

func countDays_sweepline(days int, meetings [][]int) int {
	netMeetings := make(map[int]int) // {key: date, value: delta meetings}
	dates := make([]int, 0)
	for _, meeting := range meetings {
		start, end := meeting[0], meeting[1]
		netMeetings[start]++
		netMeetings[end+1]-- // end is the last day of the meeting => end+1 is the first day without this meeting
	}
	// dummy date: because last meeting is likely not to end at the last day, 
	// we set a positive delta at the days+1 so that we can include "no-meeting" duration between the last meeting and the end of the days
	netMeetings[days+1]++

	for date, _ := range netMeetings {
		dates = append(dates, date)
	}
	sort.Ints(dates)

	integral := 0 // cumulative meetings as of (date - 1). 上轮loop累积的meetings数量
	curDate := 1 // last seen frist date when integral becomes 0. curDate的更新永远跟着integral走，每次integral由正数更新成0时，更新curDate
	res := 0
	for _, date := range dates {
		delta := netMeetings[date] // 本轮loop的meetings"增量"
		if integral == 0 && integral+delta > 0 { // "好日子(integral==0)"到头了, 计算“好日子”持续的时长
			res += date - curDate
		} else if integral > 0 && integral+delta == 0 { // "好日子"又开始了, 标记开始日期, 储存到curDate
			curDate = date
		}
		integral += delta
	}

	return res
}
