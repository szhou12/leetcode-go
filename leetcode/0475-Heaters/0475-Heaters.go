package leetcode

import "sort"

func findRadius(houses []int, heaters []int) int {
	// for each house, find the closest heater to its left & to its right, 
    // take the min of two = the min radius to cover this house
    // loop all houses, find the min radius of each and take the max
	sort.Ints(heaters)

	res := 0

	var rad int

	for _, h := range houses {
		nearest := lowerBound(heaters, h)

		// case 1:  heater_0, heater_1, ..., heater_n-1 | house
        // case 2:  house | heater_0, heater_1, ..., heater_n-1
        // case 3: heater_0, heater_1, ... | house | ..., heater_n-1

		if nearest == len(heaters) {
			rad = h - heaters[nearest-1]
		} else if nearest == 0 {
			rad = heaters[nearest] - h
		} else {
			rad = min(heaters[nearest]-h, h-heaters[nearest-1])
		}
		res = max(res, rad)
	}

	return res
}

func lowerBound(nums []int, target int) int {
	res := sort.Search(
		len(nums),
		func(i int) bool {
			return nums[i] >= target
		},
	)

	return res
}
