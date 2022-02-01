/*
 * @lc app=leetcode id=105 lang=golang
 *
 * [105] Construct Binary Tree from Preorder and Inorder Traversal
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
func buildTree(preorder []int, inorder []int) *TreeNode {

	if len(preorder) == 0 {
		return nil
	}

	// first of preorder is absolutely root
	rootVal := preorder[0]
	// find index of root in inorder
	piv := indexOf(inorder, rootVal)

	root := &TreeNode{
		Val: rootVal,
		// leftSubtree's end is equal to inorder's root
		Left: buildTree(preorder[1:piv+1], inorder[:piv]),
		// rightSubtree is next of inorder's root
		Right: buildTree(preorder[piv+1:], inorder[piv+1:]),
	}

	return root
}

func indexOf(s []int, v int) int {
	for i, e := range s {
		if e == v {
			return i
		}
	}

	return -1
}

// @lc code=end
