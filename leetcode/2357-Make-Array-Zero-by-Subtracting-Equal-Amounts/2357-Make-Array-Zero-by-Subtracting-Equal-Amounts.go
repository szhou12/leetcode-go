package leetcode

func minimumOperations(nums []int) int {
    dict := make(map[int]bool)
    for _, num := range nums {
        if _, ok := dict[num]; !ok && num != 0 {
            dict[num] = true
        }
    }

    return len(dict)
}