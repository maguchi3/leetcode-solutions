/*
 * @lc app=leetcode id=98 lang=golang
 *
 * [98] Validate Binary Search Tree
 */

package leetcode

import (
	"github.com/maguroid/leetcode/go/structures"
)

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

/* recursive traversal */
func isValidBST(root *TreeNode) bool {
	maxInt, minInt := 1<<63-1, -1<<63
	return validate(root, maxInt, minInt)
}

func validate(node *TreeNode, high int, low int) bool {
	if node == nil {
		return true
	}

	if node.Val >= high || node.Val <= low {
		return false
	}

	return validate(node.Left, node.Val, low) && validate(node.Right, high, node.Val)
}

// @lc code=end

/* iterative in-order traversal */
func _(root *TreeNode) bool {
	stack, prev := []*TreeNode{}, -1<<63

	for root != nil || len(stack) != 0 {
		for root != nil {
			stack = append(stack, root)
			root = root.Left
		}

		if len(stack) != 0 {
			root = stack[len(stack)-1]
			stack = stack[:len(stack)-1]
		}

		if root.Val <= prev {
			return false
		}
		prev = root.Val
		root = root.Right
	}

	return true
}
