/*
 * @lc app=leetcode id=209 lang=typescript
 *
 * [209] Minimum Size Subarray Sum
 */

// @lc code=start
function minSubArrayLen(target: number, nums: number[]): number {
  let l = 0;
  let d = Number.MAX_SAFE_INTEGER;
  let sum = 0;

  for (let r = 0; r < nums.length; r++) {
    sum += nums[r];

    while (sum >= target) {
      d = Math.min(d, r - l + 1);
      sum -= nums[l];
      l++;
    }
  }

  return d === Number.MAX_SAFE_INTEGER ? 0 : d;
}
// @lc code=end
