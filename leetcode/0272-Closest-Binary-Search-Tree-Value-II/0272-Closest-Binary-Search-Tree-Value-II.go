package leetcode

import "math"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func closestKValues(root *TreeNode, target float64, k int) []int {
	queue := make([]int, 0)

	dfs(root, target, k, &queue)

	res := make([]int, 0)
	for len(queue) > 0 {
		res = append(res, queue[0])
		queue = queue[1:]
	}

	return res
}

/* 
In-Order Traversal
  Note: queue needs to pass in as a pointer because there is append() operation in recursion
*/
func dfs(root *TreeNode, target float64, k int, queue *[]int) {
	// base case
	if root == nil {
		return
	}

	// In-order traversal
	// left-subtree
	dfs(root.Left, target, k, queue)
	// current node
	if len(*queue) < k {
		*queue = append(*queue, root.Val)
	} else {
		if math.Abs(target-float64((*queue)[0])) > math.Abs(target-float64(root.Val)) {
			*queue = (*queue)[1:]
			*queue = append(*queue, root.Val)
		} else {
			return
		}
	}
	// right-subtree
	dfs(root.Right, target, k, queue)

}

/*
* Input: root = [4,2,5,1,3], target = 3.714286, k = 2
* Output: [3,4]
*/