/*
 * @lc app=leetcode id=102 lang=typescript
 *
 * [102] Binary Tree Level Order Traversal
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

function levelOrder(root: TreeNode | null): number[][] {
  if (root === null) return [];

  const levels: number[][] = [];

  const order = (node: TreeNode | null, lv: number) => {
    if (node === null) return;

    if (levels.length === lv) levels[lv] = new Array<number>();

    levels[lv].push(node.val);

    if (node.left) order(node.left, lv + 1);
    if (node.right) order(node.right, lv + 1);
  };

  order(root, 0);

  return levels;
}
// @lc code=end
