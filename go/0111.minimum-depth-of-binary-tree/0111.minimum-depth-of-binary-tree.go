/*
 * @lc app=leetcode id=111 lang=golang
 *
 * [111] Minimum Depth of Binary Tree
 */

package leetcode

import (
	"math"

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

// iteration with Que
func minDepth(root *TreeNode) int {
	depth := 0
	if root == nil {
		return depth
	}

	curQue := []*TreeNode{root}

	for len(curQue) > 0 {
		depth++
		nextQue := []*TreeNode{}

		for _, node := range curQue {
			if node.Left == nil && node.Right == nil {
				return depth
			}

			if node.Left != nil {
				nextQue = append(nextQue, node.Left)
			}

			if node.Right != nil {
				nextQue = append(nextQue, node.Right)
			}
		}

		curQue = nextQue
	}

	return depth
}

// @lc code=end

// recursion
func _(root *TreeNode) int {
	if root == nil {
		return 0
	}

	if root.Left == nil && root.Right == nil {
		return 1
	}

	min := math.MaxInt

	if root.Left != nil {
		left := minDepth(root.Left)
		if left < min {
			min = left
		}
	}

	if root.Right != nil {
		right := minDepth(root.Right)
		if right < min {
			min = right
		}
	}

	return min + 1
}
