package leetcode

import (
	"fmt"
	"testing"
)

type question97 struct {
	para97
	ans97
}

// para 是参数
type para97 struct {
	s1 string
	s2 string
	s3 string
}

// ans 是答案
// one 代表第一个答案
type ans97 struct {
	one bool
}

func Test_Problem97(t *testing.T) {

	qs := []question97{

		{
			para97{"aabcc", "dbbca", "aadbbcbcac"},
			ans97{one: true},
		},
		{
			para97{"aabcc", "dbbca", "aadbbbaccc"},
			ans97{false},
		},
		{
			para97{"", "", ""},
			ans97{true},
		},
		{
			para97{"ab", "c", "abc"},
			ans97{true},
		},
		{
			para97{"broadcast", "parade", "brparoaadcasdet"},
			ans97{true},
		},
	}

	fmt.Printf("------------------------Leetcode Problem 97------------------------\n")

	for _, q := range qs {
		expect, input := q.ans97, q.para97
		testname := fmt.Sprintf("%v,%v,%v", input.s1, input.s2, input.s3)
		t.Run(testname, func(t *testing.T) {
			ans := isInterleave(input.s1, input.s2, input.s3)
			if ans != expect.one {
				t.Errorf("got %v, want %v", ans, expect.one)
			}
			fmt.Printf("【input】:%v       【output】:%v\n", input, ans)
		})
	}
	fmt.Printf("\n\n\n")
}
