/*
 * @lc app=leetcode id=105 lang=typescript
 *
 * [105] Construct Binary Tree from Preorder and Inorder Traversal
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

function buildTree(preorder: number[], inorder: number[]): TreeNode | null {
  let pIdx = 0;
  const iIdxMap = new Map<number, number>();

  inorder.forEach((num, i) => iIdxMap.set(num, i));

  const arrayToTree = (left: number, right: number) => {
    if (left > right) return null;

    const rootVal = preorder[pIdx];
    const root = new TreeNode(rootVal);

    pIdx++;

    root.left = arrayToTree(left, iIdxMap.get(rootVal) - 1);
    root.right = arrayToTree(iIdxMap.get(rootVal) + 1, right);

    return root;
  };

  return arrayToTree(0, preorder.length - 1);
}

// @lc code=end
