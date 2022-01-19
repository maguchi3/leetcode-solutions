/*
 * @lc app=leetcode id=98 lang=typescript
 *
 * [98] Validate Binary Search Tree
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

function isValidBST(root: TreeNode | null): boolean {
  const validate = (node: TreeNode | null, hi: number, lo: number) => {
    if (node === null) return true;

    if ((hi !== null && node.val >= hi) || (lo !== null && node.val <= lo))
      return false;

    return (
      validate(node.left, node.val, lo) && validate(node.right, hi, node.val)
    );
  };

  return validate(root, null, null);
}
// @lc code=end
