/*
 * @lc app=leetcode id=18 lang=typescript
 *
 * [18] 4Sum
 */

// @lc code=start
function fourSum(nums: number[], target: number): number[][] {
  nums = nums.sort((a, b) => a - b);

  return kSum(nums, target, 0, 4);
}

function kSum(nums: number[], target: number, start: number, k: number) {
  const res: number[][] = [];

  if (start === nums.length) {
    return res;
  }

  let avr = target / k;

  if (nums[start] > avr || avr > nums[nums.length - 1]) return res;

  if (k === 2) return twoSum(nums, target, start);

  for (let i = start; i < nums.length; i++) {
    if (i === start || nums[i] !== nums[i - 1]) {
      for (let set of kSum(nums, target - nums[i], i + 1, k - 1)) {
        set.push(nums[i]);
        res.push(set);
      }
    }
  }

  return res;
}

function twoSum(nums: number[], target: number, start: number) {
  const res: number[][] = [];

  let lo = start;
  let hi = nums.length - 1;

  while (lo < hi) {
    let sum = nums[lo] + nums[hi];

    if (sum < target || (lo > start && nums[lo] === nums[lo - 1])) {
      lo++;
    } else if (
      sum > target ||
      (hi < nums.length - 1 && nums[hi] === nums[hi + 1])
    ) {
      hi--;
    } else {
      res.push([nums[lo++], nums[hi--]]);
    }
  }

  return res;
}
// @lc code=end
