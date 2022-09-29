package leetcode

type Node struct {
	Val   int
	Left  *Node
	Right *Node
	Next  *Node
}

func connect(root *Node) *Node {
	// edge case
	if root == nil {
		return root
	}

	queue := make([]*Node, 0)
	queue = append(queue, root)

	for len(queue) > 0 {
		size := len(queue)
		for i := 0; i < size-1; i++ {
			cur := queue[0]
			queue = queue[1:]
			cur.Next = queue[0]
			if cur.Left != nil {
				queue = append(queue, cur.Left)
			}
			if cur.Right != nil {
				queue = append(queue, cur.Right)
			}

		}
		last := queue[0]
		queue = queue[1:]
		last.Next = nil
		if last.Left != nil {
			queue = append(queue, last.Left)
		}
		if last.Right != nil {
			queue = append(queue, last.Right)
		}
	}

	return root
}
