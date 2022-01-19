/*
 * @lc app=leetcode id=213 lang=typescript
 *
 * [213] House Robber II
 */

// @lc code=start
function rob(nums: number[]): number {
  if (nums.length <= 3) return Math.max(...nums);

  const m1 = robFrom(nums, 0, nums.length - 2);
  const m2 = robFrom(nums, 1, nums.length - 1);

  return Math.max(m1, m2);
}

function robFrom(nums: number[], start: number, end: number) {
  let dp = new Array(start + (end - start + 1));

  for (let i = start; i <= end; i++) {
    let before = dp[i - 1] ? dp[i - 1] : 0;
    let twoBefore = dp[i - 2] ? dp[i - 2] : 0;

    dp[i] = Math.max(before, twoBefore + nums[i]);
  }

  return dp[end];
}
// @lc code=end
