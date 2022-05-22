package leetcode

import (
	"fmt"
	"testing"
)

type question47 struct {
	para47
	ans47
}

// para 是参数
// s 代表第一个参数
type para47 struct {
	s []int
}

// ans 是答案
// one 代表第一个答案
type ans47 struct {
	one [][]int
}

func Test_Problem47(t *testing.T) {
	qs := []question47{

		{
			para47{[]int{1, 1, 2}},
			ans47{[][]int{{1, 1, 2}, {1, 2, 1}, {2, 1, 1}}},
		},

		{
			para47{[]int{1, 2, 3}},
			ans47{[][]int{{1, 2, 3}, {1, 3, 2}, {2, 1, 3}, {2, 3, 1}, {3, 1, 2}, {3, 2, 1}}},
		},
	}

	fmt.Printf("------------------------Leetcode Problem 47------------------------\n")

	for _, q := range qs {
		expect, p := q.ans47, q.para47

		testname := fmt.Sprintf("【input】: %v", p.s)
		t.Run(testname, func(t *testing.T) {
			ans := permuteUnique(p.s)
			if !test2DSlice(ans, expect.one) {
				t.Errorf("got %v, want %v", ans, expect.one)
			}
			fmt.Printf("【expect】:%v       【output】:%v\n", expect.one, ans)
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
