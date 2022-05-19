package leetcode

import (
	"fmt"
	"testing"
)

type question322 struct {
	para322
	ans322
}

// para 是参数
// s 代表第一个参数
type para322 struct {
	coins  []int
	amount int
}

// ans 是答案
// one 代表第一个答案
type ans322 struct {
	one int
}

func Test_Problem322(t *testing.T) {
	qs := []question322{

		{
			para322{[]int{1, 2, 5}, 11},
			ans322{3},
		},

		{
			para322{[]int{2}, 3},
			ans322{-1},
		},

		{
			para322{[]int{1}, 0},
			ans322{0},
		},
	}

	fmt.Printf("------------------------Leetcode Problem 322------------------------\n")

	for _, q := range qs {
		expect, p := q.ans322, q.para322

		testname := fmt.Sprintf("%v,%v", p.coins, p.amount)
		t.Run(testname, func(t *testing.T) {
			ans := coinChange(p.coins, p.amount)
			if ans != expect.one {
				t.Errorf("got %v, want %v", ans, expect.one)
			}
			fmt.Printf("【input】:%v       【output】:%v\n", p, ans)
		})

	}
	fmt.Printf("\n\n\n")
}
