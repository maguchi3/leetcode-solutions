import TreeNode from "./types/tree-node";

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

/* with recursion */

function splitBST(
  root: TreeNode | null,
  target: number
): Array<TreeNode | null> {
  if (!root) return new Array(2).fill(null);

  if (root.val > target) {
    const [sm, lg] = splitBST(root.left, target);
    root.left = lg;
    return [sm, root];
  } else {
    const [sm, lg] = splitBST(root.right, target);
    root.right = sm;
    return [root, lg];
  }
}

/* with iteration
function splitBST(
  root: TreeNode | null,
  target: number
): Array<TreeNode | null> {
  const ans: TreeNode[] = new Array(2).fill(new TreeNode(null));
  if (!root) return ans;

  const q: TreeNode[] = [];

  while (root) {
    q.push(root);
    if (root.val > target) root = root.left;
    else root = root.right;
  }

  while (q.length) {
    const cur = q.pop();
    if (cur.val > target) {
      cur.left = ans[1];
      ans[1] = cur;
    } else {
      cur.right = ans[0];
      ans[0] = cur;
    }
  }

  return ans;
}
*/
