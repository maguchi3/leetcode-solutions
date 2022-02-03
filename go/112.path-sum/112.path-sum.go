/*
 * @lc app=leetcode id=112 lang=golang
 *
 * [112] Path Sum
 */

package leetcode

import "github.com/maguroid/leetcode/go/structures"

type TreeNode = structures.TreeNode

// @lc code=start

func hasPathSum(root *TreeNode, targetSum int) bool {
	if root == nil {
		return false
	}

	if root.Left == nil && root.Right == nil {
		return root.Val == targetSum
	}

	return hasPathSum(root.Left, targetSum-root.Val) || hasPathSum(root.Right, targetSum-root.Val)
}

// @lc code=end
