/*
 * @lc app=leetcode id=101 lang=typescript
 *
 * [101] Symmetric Tree
 */

export class TreeNode {
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

function isSymmetric(root: TreeNode | null): boolean {
  return helper(root.left, root.right);
}

function helper(left: TreeNode | null, right: TreeNode | null) {
  if (left === null && right === null) return true;
  if (left === null || right === null) return false;
  if (left.val !== right.val) return false;

  return helper(left.left, right.right) && helper(left.right, right.left);
}
// @lc code=end
export default isSymmetric;
