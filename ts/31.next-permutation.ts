/*
 * @lc app=leetcode id=31 lang=typescript
 *
 * [31] Next Permutation
 */

// @lc code=start
/**
 Do not return anything, modify nums in-place instead.
 */
function nextPermutation(nums: number[]): void {
  const reverse = (start: number) => {
    let i = start;
    let j = nums.length - 1;

    while (i < j) {
      swap(i, j);
      i++;
      j--;
    }
  };

  const swap = (i: number, j: number) => {
    let temp = nums[i];
    nums[i] = nums[j];
    nums[j] = temp;
  };

  let i = nums.length - 2;
  while (i >= 0 && nums[i + 1] <= nums[i]) i--;

  if (i >= 0) {
    let j = nums.length - 1;
    while (nums[j] <= nums[i]) j--;
    swap(i, j);
  }

  reverse(i + 1);
}
// @lc code=end
