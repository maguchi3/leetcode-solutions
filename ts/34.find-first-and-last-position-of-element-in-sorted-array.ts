/*
 * @lc app=leetcode id=34 lang=typescript
 *
 * [34] Find First and Last Position of Element in Sorted Array
 */

// @lc code=start
function searchRange(nums: number[], target: number): number[] {
  let left = 0;
  let right = nums.length - 1;

  while (left <= right) {
    let pivot = Math.floor((left + right) / 2);

    if (nums[pivot] === target) {
      let prev = pivot;
      let next = pivot;

      while (nums[prev] === target) {
        prev--;
      }

      while (nums[next] === target) {
        next++;
      }

      return [prev + 1, next - 1];
    } else if (nums[pivot] > target) {
      right = pivot - 1;
    } else {
      left = pivot + 1;
    }
  }

  return [-1, -1];
}
// @lc code=end
