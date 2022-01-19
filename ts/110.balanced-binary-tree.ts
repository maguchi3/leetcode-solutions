/*
 * @lc app=leetcode id=110 lang=typescript
 *
 * [110] Balanced Binary Tree
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

function isBalanced(root: TreeNode | null): boolean {
  if (root === null) return true;

  const leftHeight = calcHeight(root.left);
  const rightHeight = calcHeight(root.right);

  if (Math.abs(leftHeight - rightHeight) > 1) {
    return false;
  } else {
    return isBalanced(root.left) && isBalanced(root.right);
  }
}

function calcHeight(node: TreeNode | null): number {
  if (node === null || node.val === null) return 0;

  const left = calcHeight(node.left);
  const right = calcHeight(node.right);

  return Math.max(left, right) + 1;
}
// @lc code=end
