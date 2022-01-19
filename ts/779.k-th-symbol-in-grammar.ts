/*
 * @lc app=leetcode id=779 lang=typescript
 *
 * [779] K-th Symbol in Grammar
 */

// @lc code=start
function kthGrammar(n: number, k: number): number {
  if (n === 1) return 0;

  let parent = kthGrammar(n - 1, Math.ceil(k / 2));

  const isEven = k % 2 === 0;

  if (parent === 0) return isEven ? 1 : 0;
  else return isEven ? 0 : 1;
}
// @lc code=end
