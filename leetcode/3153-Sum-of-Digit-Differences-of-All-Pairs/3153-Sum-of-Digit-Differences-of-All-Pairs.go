package leetcode

import "strconv"

func sumDigitDifferences(nums []int) int64 {
	n := len(nums) // population
	totalDigits := len(strconv.Itoa(nums[0]))

	// count[i][0...9] := frequency of 0...9 at i-th digit
	count := make([][10]int, totalDigits)
	for _, num := range nums {
		str := strconv.Itoa(num)
		for i := 0; i < totalDigits; i++ {
			cur := int(str[i] - '0')
			count[i][cur]++
		}
	}

	res := 0
	// for i-th digit position, # of matching pairs with number k = freq(k)*(n-freq(k))
	for i := 0; i < totalDigits; i++ {
		for digit := 0; digit < 10; digit++ {
			freq := count[i][digit]
			res += freq * (n - freq)
		}
	}

	// divides 2 because (a,b) and (b,a) are the same pair
	return int64(res/2)

}
