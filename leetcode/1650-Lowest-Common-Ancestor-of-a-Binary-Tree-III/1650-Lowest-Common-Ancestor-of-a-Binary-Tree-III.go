package leetcode

type TreeNode struct {
	Val    int
	Left   *TreeNode
	Right  *TreeNode
	Parent *TreeNode
}

// Method 1: HashMap to keep track of all ancestors of p
func lowestCommonAncestor(p, q *TreeNode) *TreeNode {
	ancestors := make(map[*TreeNode]bool)

	for p != nil {
		ancestors[p] = true
		p = p.Parent
	}

	for q != nil {
		if ancestors[q] {
			return q
		}
		q = q.Parent
	}

	return q
}
