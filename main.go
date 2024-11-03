package main

// Terminal: > go run main.go

import (
	"fmt"
	cur "github.com/szhou12/leetcode-go/interview-prep/amz"
)

// NOTE: in order to call any func from imported package, this func must be exported!!!
// (i.e. func name MUST start with a capital letter)
func main() {
	volumes := []int{9, 2, 4, 6}
	k := 3

	fmt.Println(cur.FindMinCapacity(k, volumes))
}

