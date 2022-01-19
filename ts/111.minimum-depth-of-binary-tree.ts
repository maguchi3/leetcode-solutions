/*
 * @lc app=leetcode id=111 lang=typescript
 *
 * [111] Minimum Depth of Binary Tree
 */

class TreeNode {
  val: number;
  left: TreeNode | null;
  right: TreeNode | null;
  constructor(val?: number, left?: TreeNode | null, right?: TreeNode | null) {
    this.val = val === undefined ? 0 : val;
    this.left = left === undefined ? null : left;
    this.right = right === undefined ? null : right;
  }
}
// @lc code=start

function minDepth(root: TreeNode | null): number {
  if (root === null) return 0;
  if (root.left === null && root.right === null) return 1;

  let min = Number.MAX_VALUE;

  if (root.left !== null) {
    min = Math.min(min, minDepth(root.left));
  }

  if (root.right !== null) {
    min = Math.min(min, minDepth(root.right));
  }

  return min + 1;
}
// @lc code=end
