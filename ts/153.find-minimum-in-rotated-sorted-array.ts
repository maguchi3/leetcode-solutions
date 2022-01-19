/*
 * @lc app=leetcode id=153 lang=typescript
 *
 * [153] Find Minimum in Rotated Sorted Array
 */

// @lc code=start
function findMin(nums: number[]): number {
  let l = 0;
  let r = nums.length - 1;

  while (l < r) {
    let m = ~~((l + r) / 2);

    if (nums[m] > nums[r]) l = m + 1;
    else r = m;
  }

  return nums[l];
}
// @lc code=end
