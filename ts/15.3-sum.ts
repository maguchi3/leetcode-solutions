/*
 * @lc app=leetcode id=15 lang=typescript
 *
 * [15] 3Sum
 */

// @lc code=start
function threeSum(nums: number[]): number[][] {
  const res: number[][] = [];
  nums = nums.sort((a, b) => a - b);

  for (let i = 0; i < nums.length; i++) {
    if (i === 0 || nums[i - 1] !== nums[i]) {
      movePointer(nums, i, res);
    }
  }

  return res;
}

function movePointer(nums: number[], i: number, res: number[][]): void {
  let low = i + 1;
  let high = nums.length - 1;

  while (low < high) {
    let sum = nums[i] + nums[low] + nums[high];

    if (sum < 0) {
      low++;
    } else if (sum > 0) {
      high--;
    } else {
      res.push([nums[i], nums[low++], nums[high--]]);
      while (low < high && nums[low] === nums[low - 1]) low++;
    }
  }
}
// @lc code=end
