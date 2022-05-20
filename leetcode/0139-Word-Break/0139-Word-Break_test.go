package leetcode

import (
	"fmt"
	"testing"
)

type question139 struct {
	para139
	ans139
}

// para 是参数
// s 代表第一个参数
type para139 struct {
	s    string
	dict []string
}

// ans 是答案
// one 代表第一个答案
type ans139 struct {
	one bool
}

func Test_Problem139(t *testing.T) {
	qs := []question139{

		{
			para139{"leetcode", []string{"leet", "code"}},
			ans139{true},
		},

		{
			para139{"applepenapple", []string{"apple", "pen"}},
			ans139{true},
		},

		{
			para139{"catsandog", []string{"cats", "dog", "sand", "and", "cat"}},
			ans139{false},
		},
	}

	fmt.Printf("------------------------Leetcode Problem 322------------------------\n")

	for _, q := range qs {
		expect, p := q.ans139, q.para139

		testname := fmt.Sprintf("【input】: %v,%v", p.s, p.dict)
		t.Run(testname, func(t *testing.T) {
			ans := wordBreak(p.s, p.dict)
			if ans != expect.one {
				t.Errorf("got %v, want %v", ans, expect.one)
			}
			fmt.Printf("【expect】:%v       【output】:%v\n", expect.one, ans)
		})

	}
	fmt.Printf("\n\n\n")
}
