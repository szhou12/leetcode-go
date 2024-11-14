package main

// Terminal: > go run main.go

import (
	"fmt"
	cur "github.com/szhou12/leetcode-go/interview-prep/amz"
)

// NOTE: in order to call any func from imported package, this func must be exported!!!
// (i.e. func name MUST start with a capital letter)
func main() {
	nums := []int{2, 2, 1, 5, 3}

	fmt.Println(cur.MaxSumKGroupsMedian(nums, 2))


}

