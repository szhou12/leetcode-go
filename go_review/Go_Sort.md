# [Go Sort Package](https://github.com/szhou12/leetcode-go/blob/main/go_review/README.md)

## Contents
* [SearchInts](#searchints)

## SearchInts 
```go
// arr MUST be sorted in ascending order
sort.SearchInts(arr []int, x int) int
```
Returns: first (smallest) index `i` where `arr[i] >= x` (`len(arr)` if `x` is greater than all elements in `arr`).

Example:
```linux
arr := []int{1, 2, 2, 4, 6, 7, 8}

x := 2
i := sort.SearchInts(arr, x) // Return 1

x := 0
i := sort.SearchInts(arr, x) // Return 0

x = 5
i := sort.SearchInts(arr, x) // Return 4

x = 100
i := sort.SearchInts(arr, x) // Return 7 = len(arr)
```