/*
 * @lc app=leetcode id=104 lang=golang
 *
 * [104] Maximum Depth of Binary Tree
 */

package leetcode

import (
	"container/list"

	"github.com/maguroid/leetcode/go/structures"
)

type TreeNode = structures.TreeNode

// @lc code=start

// recursion
func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}

	if root.Left == nil && root.Right == nil {
		return 1
	}

	leftDepth := maxDepth(root.Left) + 1
	rightDepth := maxDepth(root.Right) + 1

	if leftDepth > rightDepth {
		return leftDepth
	}

	return rightDepth
}

// @lc code=end

// iteration with queue(container/list)
func maxDepth2(root *TreeNode) int {
	if root == nil {
		return 0
	}

	q := list.New()
	q.PushBack(root)
	lv := 0

	for q.Len() > 0 {
		n := q.Len()

		for i := 0; i < n; i++ {
			f := q.Front()
			q.Remove(f)

			node := f.Value.(*TreeNode)
			if node.Left != nil {
				q.PushBack(node.Left)
			}
			if node.Right != nil {
				q.PushBack(node.Right)
			}
		}
		lv++
	}

	return lv
}
