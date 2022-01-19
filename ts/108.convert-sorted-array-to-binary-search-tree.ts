/*
 * @lc app=leetcode id=108 lang=typescript
 *
 * [108] Convert Sorted Array to Binary Search Tree
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

function sortedArrayToBST(nums: number[]): TreeNode | null {
  const node = createNode(nums, 0, nums.length - 1);

  return node;
}

function createNode(nums: number[], left: number, right: number) {
  if (left > right) return null;
  const middle = Math.floor((left + right) / 2);
  const root = new TreeNode(nums[middle]);
  root.left = createNode(nums, left, middle - 1);
  root.right = createNode(nums, middle + 1, right);

  return root;
}

// @lc code=end
export default sortedArrayToBST;
