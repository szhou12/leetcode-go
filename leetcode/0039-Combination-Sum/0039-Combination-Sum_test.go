package leetcode

import (
	"fmt"
	"testing"
)

type question39 struct {
	para39
	ans39
}

// para 是参数
type para39 struct {
	candidates []int
	target     int
}

// ans 是答案
// one 代表第一个答案
type ans39 struct {
	one [][]int
}

func Test_Problem39(t *testing.T) {
	qs := []question39{

		{
			para39{[]int{2, 3, 6, 7}, 7},
			ans39{[][]int{{7}, {2, 2, 3}}},
		},

		{
			para39{[]int{2, 3, 5}, 8},
			ans39{[][]int{{2, 2, 2, 2}, {2, 3, 3}, {3, 5}}},
		},
	}

	fmt.Printf("------------------------Leetcode Problem 39------------------------\n")

	for _, q := range qs {
		expect, p := q.ans39, q.para39

		testname := fmt.Sprintf("%v,%v", p.candidates, p.target)
		t.Run(testname, func(t *testing.T) {
			ans := combinationSum(p.candidates, p.target)
			if !test2DSlice(ans, expect.one) {
				t.Errorf("got %v, want %v", ans, expect.one)
			}
			fmt.Printf("【input】:%v       【output】:%v\n", p, ans)
		})

	}
	fmt.Printf("\n\n\n")
}

func test2DSlice(a [][]int, b [][]int) bool {
	if len(a) != len(b) {
		return false
	}

	for _, ai := range a {
		exist := false
		for _, bj := range b {
			if len(ai) == len(bj) && testEq(ai, bj) {
				exist = true
			}
		}
		if !exist {
			return false
		}
	}

	return true

}

func testEq(a []int, b []int) bool {
	if len(a) != len(b) {
		return false
	}

	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}

	return true
}
