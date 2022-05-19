package leetcode

import (
	"fmt"
	"testing"
)

type question516 struct {
	para516
	ans516
}

// para 是参数
type para516 struct {
	s1 string
}

// ans 是答案
// one 代表第一个答案
type ans516 struct {
	one int
}

func Test_Problem516(t *testing.T) {

	qs := []question516{

		{
			para516{"character"},
			ans516{one: 5},
		},
		{
			para516{"bbbab"},
			ans516{one: 4},
		},
		{
			para516{"cbbd"},
			ans516{one: 2},
		},
	}

	fmt.Printf("------------------------Leetcode Problem 516------------------------\n")

	for _, q := range qs {
		expect, input := q.ans516, q.para516
		testname := fmt.Sprintf("%v", input.s1)
		t.Run(testname, func(t *testing.T) {
			ans := longestPalindromeSubseq(input.s1)
			if ans != expect.one {
				t.Errorf("got %v, want %v", ans, expect.one)
			}
			fmt.Printf("【input】:%v       【output】:%v\n", input, ans)
		})
	}
	fmt.Printf("\n\n\n")
}
