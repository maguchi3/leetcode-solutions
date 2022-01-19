/*
 * @lc app=leetcode id=33 lang=typescript
 *
 * [33] Search in Rotated Sorted Array
 */

// @lc code=start
function search(nums: number[], target: number): number {
  const findRotateIndex = (left: number, right: number) => {
    if (nums[left] < nums[right]) return 0;

    while (left <= right) {
      let pivot = Math.floor((left + right) / 2);
      if (nums[pivot] > nums[pivot + 1]) {
        return pivot + 1;
      } else {
        if (nums[pivot] < nums[left]) {
          right = pivot - 1;
        } else {
          left = pivot + 1;
        }
      }
    }
    return 0;
  };

  const binarySearch = (left: number, right: number) => {
    while (left <= right) {
      let pivot = Math.floor((left + right) / 2);
      if (nums[pivot] === target) {
        return pivot;
      } else {
        if (target < nums[pivot]) {
          right = pivot - 1;
        } else {
          left = pivot + 1;
        }
      }
    }

    return -1;
  };

  let n = nums.length;

  if (n === 1) return nums[0] === target ? 0 : -1;

  let rotateIndex = findRotateIndex(0, n - 1);

  if (nums[rotateIndex] === target) return rotateIndex;

  if (rotateIndex === 0) return binarySearch(0, n - 1);

  if (target < nums[0]) return binarySearch(rotateIndex, n - 1);

  return binarySearch(0, rotateIndex);
}
// @lc code=end
