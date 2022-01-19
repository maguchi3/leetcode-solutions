/*
 * @lc app=leetcode id=70 lang=typescript
 *
 * [70] Climbing Stairs
 */

// @lc code=start

function climbStairs(n: number): number {
  if (memo.has(n)) return memo.get(n);
  const result = climbStairs(n - 1) + climbStairs(n - 2);

  memo.set(n, result);

  return result;
}

const memo = new Map([
  [1, 1],
  [2, 2],
]);

// @lc code=end
