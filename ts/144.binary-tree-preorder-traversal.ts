/*
 * @lc app=leetcode id=144 lang=typescript
 *
 * [144] Binary Tree Preorder Traversal
 */

import TreeNode from "./types/tree-node";

// @lc code=start
function preorderTraversal(root: TreeNode | null): number[] {
  if (root === null) return [];
  const stack: TreeNode[] = [root];
  const result = [];

  while (stack.length) {
    root = stack.pop();
    result.push(root.val);
    if (root.right !== null) {
      stack.push(root.right);
    }
    if (root.left !== null) {
      stack.push(root.left);
    }
  }

  return result;
}

// @lc code=end
