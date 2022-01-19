/*
 * @lc app=leetcode id=16 lang=typescript
 *
 * [16] 3Sum Closest
 */

// @lc code=start
function threeSumClosest(nums: number[], target: number): number {
  nums = nums.sort((a, b) => a - b);
  let diff = Number.MAX_VALUE;

  for (let i = 0; i < nums.length; i++) {
    let lo = i + 1;
    let hi = nums.length - 1;

    while (lo < hi) {
      let sum = nums[i] + nums[lo] + nums[hi];

      if (Math.abs(target - sum) < Math.abs(diff)) diff = target - sum;

      if (sum < target) {
        lo++;
      } else if (sum > target) {
        hi--;
      } else {
        return sum;
      }
    }
  }

  return target - diff;
}
// @lc code=end
