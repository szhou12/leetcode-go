package leetcode

// In-place:
// 每个偶数位置检查，如果是奇数元素，则从奇数位置从小到大找存有偶数元素的，进行交换
func sortArrayByParityII(nums []int) []int {
	oddIndex := 1
	for evenIndex := 0; evenIndex < len(nums); evenIndex += 2 {
		if nums[evenIndex]%2 != 0 {
			for nums[oddIndex]%2 == 1 {
				oddIndex += 2
			}
			nums[oddIndex], nums[evenIndex] = nums[evenIndex], nums[oddIndex]
		}
	}
	return nums
}

// naive solution
// 两个数组分别存下偶数元素，奇数元素
// 再合并起来
func sortArrayByParityII_naive(nums []int) []int {
	evens := make([]int, 0)
	odds := make([]int, 0)

	for _, val := range nums {
		if val%2 == 0 {
			evens = append(evens, val)
		} else {
			odds = append(odds, val)
		}
	}

	res := make([]int, len(nums))
	for i := 0; i < len(res); i++ {
		if i%2 == 0 {
			res[i] = evens[0]
			evens = evens[1:]
		} else {
			res[i] = odds[0]
			odds = odds[1:]
		}
	}
	return res

}
