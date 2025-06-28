package leetcode

func findNumbers(nums []int) int {
    res := 0

    for _, v := range nums {
        if hasEvenDigits(v) {
            res++
        }
    }

    return res
}

func hasEvenDigits(num int) bool {
    count := 0
    for num > 0 {
        num /= 10
        count++
    }
    return count % 2 == 0
}