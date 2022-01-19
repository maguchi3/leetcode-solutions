/*
 * @lc app=leetcode id=617 lang=typescript
 *
 * [617] Merge Two Binary Trees
 */

import TreeNode from "./types/tree-node";

// @lc code=start
/**
 * Definition for a binary tree node.
 * class TreeNode {
 *     val: number
 *     left: TreeNode | null
 *     right: TreeNode | null
 *     constructor(val?: number, left?: TreeNode | null, right?: TreeNode | null) {
 *         this.val = (val===undefined ? 0 : val)
 *         this.left = (left===undefined ? null : left)
 *         this.right = (right===undefined ? null : right)
 *     }
 * }
 */

function mergeTrees(t1: TreeNode | null, t2: TreeNode | null): TreeNode | null {
  if (!t1 && !t2) return null;

  const root = new TreeNode((t1?.val || 0) + (t2?.val || 0));
  root.left = mergeTrees(t1?.left, t2?.left);
  root.right = mergeTrees(t1?.right, t2?.right);

  return root;
}

// @lc code=end
