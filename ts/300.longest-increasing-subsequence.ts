/*
 * @lc app=leetcode id=300 lang=typescript
 *
 * [300] Longest Increasing Subsequence
 */

// @lc code=start
function lengthOfLIS(nums: number[]): number {
  const dp = new Array(nums.length).fill(1);
  for (let r = 1; r < nums.length; r++) {
    for (let l = 0; l < r; l++) {
      if (nums[r] > nums[l]) dp[r] = Math.max(dp[r], dp[l] + 1);
    }
  }

  let longest = 0;

  for (let len of dp) longest = Math.max(longest, len);

  return longest;
}

// @lc code=end
