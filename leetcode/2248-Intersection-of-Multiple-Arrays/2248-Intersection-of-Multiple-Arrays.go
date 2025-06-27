package leetcode

func intersection(nums [][]int) []int {
    n := len(nums)

    freq := make(map[int]int)
    for _, arr := range nums {
        for _, num := range arr {
            freq[num]++
        }
    }

    res := make([]int, 0)
    for _, k := range nums[0] {
        if freq[k] == n {
            res = append(res, k)
        }
    }

    return res
}