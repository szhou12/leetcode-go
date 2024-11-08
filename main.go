package main

// Terminal: > go run main.go

import (
	"fmt"
	cur "github.com/szhou12/leetcode-go/interview-prep/amz"
)

// NOTE: in order to call any func from imported package, this func must be exported!!!
// (i.e. func name MUST start with a capital letter)
func main() {
	segment := [][]int{
		{1, 4, 2},
		{6, 6, 5},
		{7, 7, 7},
		{7, 9, 1},
		{9, 10, 1},
	}

	fmt.Println(cur.FindMaxValueSizeK(segment, 5))


}

