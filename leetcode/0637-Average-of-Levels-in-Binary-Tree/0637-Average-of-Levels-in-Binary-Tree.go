package leetcode

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

func averageOfLevels(root *TreeNode) []float64 {

	queue := make([]*TreeNode, 0)
	counts := make([]int, 3) // [curLevelCnt, nextLevelCnt]

	res := make([]float64, 0)

	// start node
	queue = append(queue, root)
	counts[0] = 1

	// loop
	for len(queue) > 0 {
		sum := 0

		for i := 0; i < counts[0]; i++ {
			// pop the node
			cur := queue[0]
			queue = queue[1:]

			sum += cur.Val
			
			// move to the next level
			if cur.Left != nil {
				queue = append(queue, cur.Left)
				counts[1]++
			}
			if cur.Right != nil {
				queue = append(queue, cur.Right)
				counts[1]++
			}
		}

		// calculate the average
		res = append(res, float64(sum) / float64(counts[0]))
		counts[0] = counts[1]
		counts[1] = 0

	}

	return res

}

