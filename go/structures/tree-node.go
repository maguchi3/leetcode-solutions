package structures

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

const NULL = -1 << 63 // using express null node val

func IntsToTree(ints []int) *TreeNode {
	n := len(ints)
	if n == 0 {
		return nil
	}
	root := &TreeNode{Val: ints[0]}
	q := make([]*TreeNode, 1, n)
	q[0] = root

	i := 1

	for i < n {
		node := q[0]
		q = q[1:]

		if ints[i] != NULL {
			node.Left = &TreeNode{Val: ints[i]}
			q = append(q, node.Left)
		}
		i++
		if i < n && ints[i] != NULL {
			node.Right = &TreeNode{Val: ints[i]}
			q = append(q, node.Right)
		}
		i++
	}

	return root
}
