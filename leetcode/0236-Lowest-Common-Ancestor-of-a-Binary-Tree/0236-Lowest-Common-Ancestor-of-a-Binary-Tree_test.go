package leetcode

import (
	"fmt"
	"testing"

	"github.com/szhou12/leetcode-go/structures"
)

type question236 struct {
	para236
	ans236
}

// para 是参数
// s 代表第一个参数
type para236 struct {
	root []int
	p    []int
	q    []int
}

// ans 是答案
// one 代表第一个答案
type ans236 struct {
	one []int
}

func Test_Problem236(t *testing.T) {
	qs := []question236{

		{
			para236{[]int{}, []int{}, []int{}},
			ans236{[]int{}},
		},
		{
			para236{[]int{3, 5, 1, 6, 2, 0, 8, structures.NULL, structures.NULL, 7, 4}, []int{5}, []int{1}},
			ans236{[]int{3}},
		},
		{
			para236{[]int{6, 2, 8, 0, 4, 7, 9, structures.NULL, structures.NULL, 3, 5}, []int{2}, []int{8}},
			ans236{[]int{6}},
		},
		{
			para236{[]int{6, 2, 8, 0, 4, 7, 9, structures.NULL, structures.NULL, 3, 5}, []int{2}, []int{4}},
			ans236{[]int{2}},
		},
	}

	fmt.Printf("------------------------Leetcode Problem 236------------------------\n")

	for _, q := range qs {
		expect, param := q.ans236, q.para236

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
