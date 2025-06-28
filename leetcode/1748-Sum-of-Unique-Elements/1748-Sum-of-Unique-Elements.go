package leetcode

func sumOfUnique(nums []int) int {
    freq := make(map[int]int)
    for _, num := range nums {
        freq[num]++
    }

    res := 0
    for k, cnt := range freq {
        if cnt == 1 {
            res += k
        }
    }

    return res
}