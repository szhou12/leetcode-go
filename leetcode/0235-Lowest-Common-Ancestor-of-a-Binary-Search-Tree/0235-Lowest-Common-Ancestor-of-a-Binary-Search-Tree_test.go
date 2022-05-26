package leetcode

import (
	"fmt"
	"testing"

	"github.com/szhou12/leetcode-go/structures"
)

type question235 struct {
	para235
	ans235
}

// para 是参数
// s 代表第一个参数
type para235 struct {
	root []int
	p    []int
	q    []int
}

// ans 是答案
// one 代表第一个答案
type ans235 struct {
	one []int
}

func Test_Problem235(t *testing.T) {
	qs := []question235{

		{
			para235{[]int{}, []int{}, []int{}},
			ans235{[]int{}},
		},
		{
			para235{[]int{2, 1}, []int{2}, []int{1}},
			ans235{[]int{2}},
		},
		{
			para235{[]int{6, 2, 8, 0, 4, 7, 9, structures.NULL, structures.NULL, 3, 5}, []int{2}, []int{4}},
			ans235{[]int{2}},
		},
	}

	fmt.Printf("------------------------Leetcode Problem 235------------------------\n")

	for _, q := range qs {
		expect, param := q.ans235, q.para235

		testname := fmt.Sprintf("%v", param)
		t.Run(testname, func(t *testing.T) {
			root := structures.Ints2TreeNode(param.root)
			p := structures.Ints2TreeNode(param.p)
			q := structures.Ints2TreeNode(param.q)
			exp := structures.Ints2TreeNode(expect.one)
			ans := lowestCommonAncestor(root, p, q)
			if ans != exp {
				fmt.Println(ans)
				t.Errorf("got %v, want %v", ans, exp)
			}
			fmt.Printf("【input】:%v       【output】:%v\n", param, ans)
			fmt.Printf("【output】:%v      \n", lowestCommonAncestor(root, p, q))
		})

	}
	fmt.Printf("\n\n\n")
}
