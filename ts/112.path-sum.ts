/*
 * @lc app=leetcode id=112 lang=typescript
 *
 * [112] Path Sum
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

function hasPathSum(root: TreeNode | null, targetSum: number): boolean {
  if (root === null) return false;

  targetSum -= root.val;

  if (root.left === null && root.right === null) return targetSum === 0;

  return hasPathSum(root.left, targetSum) || hasPathSum(root.right, targetSum);
}

// @lc code=end
