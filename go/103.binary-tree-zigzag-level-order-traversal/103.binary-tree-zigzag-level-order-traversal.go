/*
 * @lc app=leetcode id=103 lang=golang
 *
 * [103] Binary Tree Zigzag Level Order Traversal
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

func zigzagLevelOrder(root *TreeNode) [][]int {
	res := [][]int{}

	if root == nil {
		return res
	}

	curNodes := []*TreeNode{root}
	lv := 0

	for len(curNodes) > 0 {
		curLevelsVal, nextNodes := []int{}, []*TreeNode{}

		for i := 0; i < len(curNodes); i++ {
			pos := i

			// if current level = odd, move right to left
			if lv%2 == 1 {
				pos = len(curNodes) - (i + 1)
			}

			curLevelsVal = append(curLevelsVal, curNodes[pos].Val)

			if curNodes[i].Left != nil {
				nextNodes = append(nextNodes, curNodes[i].Left)
			}

			if curNodes[i].Right != nil {
				nextNodes = append(nextNodes, curNodes[i].Right)
			}
		}
		curNodes = nextNodes
		res = append(res, curLevelsVal)
		lv++

	}

	return res
}

// @lc code=end
