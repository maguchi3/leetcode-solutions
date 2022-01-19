/*
 * @lc app=leetcode id=145 lang=typescript
 *
 * [145] Binary Tree Postorder Traversal
 */

import TreeNode from "./types/tree-node";

// @lc code=start

function postorderTraversal(root: TreeNode | null): number[] {
  const result: number[] = [];
  function postOrder(root: TreeNode | null) {
    if (root === null) return;

    postOrder(root.left);
    postOrder(root.right);

    result.push(root.val);
  }

  postOrder(root);

  return result;
}
// @lc code=end
