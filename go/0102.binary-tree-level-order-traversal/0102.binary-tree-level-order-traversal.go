/*
 * @lc app=leetcode id=102 lang=golang
 *
 * [102] Binary Tree Level Order Traversal
 */

package leetcode

import "github.com/maguroid/leetcode/go/structures"

type TreeNode = structures.TreeNode

// @lc code=start
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func levelOrder(root *TreeNode) [][]int {
	var res [][]int

	if root == nil {
		return res
	}

	q := make([]*TreeNode, 1)
	q[0] = root

	lv := 0

	for len(q) > 0 {
		limit := len(q)
		for i := 0; i < limit; i++ {
			if i == 0 {
				res = append(res, []int{})
			}
			node := q[0]
			q = q[1:]

			res[lv] = append(res[lv], node.Val)

			if node.Left != nil {
				q = append(q, node.Left)
			}

			if node.Right != nil {
				q = append(q, node.Right)
			}
		}

		lv++
	}

	return res
}

// @lc code=end
