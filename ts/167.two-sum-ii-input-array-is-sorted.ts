/*
 * @lc app=leetcode id=167 lang=typescript
 *
 * [167] Two Sum II - Input array is sorted
 */

// @lc code=start
function twoSum(numbers: number[], target: number): number[] {
  let low = 0;
  let high = numbers.length - 1;

  while (low < high) {
    const sum = numbers[low] + numbers[high];
    if (sum === target) {
      return [low + 1, high + 1];
    }
    if (sum < target) {
      low++;
    } else {
      high--;
    }
  }

  return [-1, -1];
}
// @lc code=end
