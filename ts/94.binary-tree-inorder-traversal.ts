/*
 * @lc app=leetcode id=94 lang=typescript
 *
 * [94] Binary Tree Inorder Traversal
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

function inorderTraversal(root: TreeNode | null): number[] {
  const result = [];
  helper(root, result);

  return result;
}

function helper(root: TreeNode, res: Array<number>): void {
  if (root !== null) {
    helper(root.left, res);
    res.push(root.val);
    helper(root.right, res);
  }
}
// @lc code=end
export default inorderTraversal;
