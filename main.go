package main

// Terminal: > go run main.go

import (
	"fmt"
	cur "github.com/szhou12/leetcode-go/interview-prep/amz"
)

// NOTE: in order to call any func from imported package, this func must be exported!!!
// (i.e. func name MUST start with a capital letter)
func main() {
	s := "1001"
	k := 0
	fmt.Println(cur.MaxSwitchingDigits(s, k))
}

