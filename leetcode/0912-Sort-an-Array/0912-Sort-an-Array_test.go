package leetcode

import (
	"fmt"
	"testing"
)

type question912 struct {
	para912
	ans912
}

// para 是参数
type para912 struct {
	nums []int
}

// ans 是答案
// one 代表第一个答案
type ans912 struct {
	one []int
}

func Test_Problem912(t *testing.T) {
	qs := []question912{

		{
			para912{[]int{5, 2, 3, 1}},
			ans912{[]int{1, 2, 3, 5}},
		},

		{
			para912{[]int{5, 1, 1, 2, 0, 0}},
			ans912{[]int{0, 0, 1, 1, 2, 5}},
		},
	}

	fmt.Printf("------------------------Leetcode Problem 912------------------------\n")

	for _, q := range qs {
		expect, p := q.ans912, q.para912

		testname := fmt.Sprintf("%v", p.nums)
		t.Run(testname, func(t *testing.T) {
			ans := sortArray(p.nums)
			if !testEq(ans, expect.one) {
				t.Errorf("got %v, want %v", ans, expect.one)
			}
			fmt.Printf("【input】:%v       【output】:%v\n", p, ans)
		})

	}
	fmt.Printf("\n\n\n")
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
