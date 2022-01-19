/*
 * @lc app=leetcode id=53 lang=typescript
 *
 * [53] Maximum Subarray
 */

// @lc code=start
function maxSubArray(nums: number[]): number {
  let currentSubarray = nums[0];
  let maxSubArray = currentSubarray;
  for (let i = 1; i < nums.length; i++) {
    let num = nums[i];
    currentSubarray = Math.max(num, currentSubarray + num);
    maxSubArray = Math.max(currentSubarray, maxSubArray);
  }

  return maxSubArray;
}
// @lc code=end
