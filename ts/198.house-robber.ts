/*
 * @lc app=leetcode id=198 lang=typescript
 *
 * [198] House Robber
 */

// @lc code=start
function rob(nums: number[]): number {
  const memo = new Array<number>(100).fill(-1);

  const robFrom = (i: number) => {
    if (i >= nums.length) return 0;
    if (memo[i] > -1) return memo[i];

    let total = Math.max(nums[i] + robFrom(i + 2), robFrom(i + 1));

    memo[i] = total;

    return total;
  };

  return robFrom(0);
}
// @lc code=end
