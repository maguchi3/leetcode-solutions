/*
 * @lc app=leetcode id=35 lang=typescript
 *
 * [35] Search Insert Position
 */

// @lc code=start
function searchInsert(nums: number[], target: number): number {
  let pivot,
    left = 0;
  let right = nums.length - 1;
  while (left <= right) {
    pivot = Math.floor(left + (right - left) / 2);
    if (nums[pivot] === target) return pivot;
    if (nums[pivot] > target) {
      right = pivot - 1;
    } else {
      left = pivot + 1;
    }
  }
  return left;
}
// @lc code=end
