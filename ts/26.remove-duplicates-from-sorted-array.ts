/*
 * @lc app=leetcode id=26 lang=typescript
 *
 * [26] Remove Duplicates from Sorted Array
 */

// @lc code=start
function removeDuplicates(nums: number[]): number {
  for (let i = 0; i < nums.length; i++) {
    while (nums[i] === nums[i + 1]) {
      nums.splice(i + 1, 1);
    }
  }

  return nums.length;
}
// @lc code=end
