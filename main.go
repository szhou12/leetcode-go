package main

// Terminal: > go run main.go

import (
	"fmt"
	cur "github.com/szhou12/leetcode-go/interview-prep/amz"
)

// NOTE: in order to call any func from imported package, this func must be exported!!!
// (i.e. func name MUST start with a capital letter)
func main() {
	x := []int{0, 0, 0, 0, 0, 1, 1, 1, 2, -1, -1, -2, -1}
	y := []int{-1, 0, 1, 2, -2, 0, 1, -1, 0, 1, -1, 0, 0}

	fmt.Println(cur.NumNotWorkingMachinese(x, y))
}

