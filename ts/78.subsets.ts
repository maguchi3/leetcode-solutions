/*
 * @lc app=leetcode id=78 lang=typescript
 *
 * [78] Subsets
 */

// @lc code=start
function subsets(nums: number[]): number[][] {
  const n = nums.length;
  const res = [];
  const set = [];

  const backTrack = (start: number, size: number) => {
    if (set.length === size) {
      res.push([...set]);
      return;
    }

    for (let i = start; i < n; i++) {
      set.push(nums[i]);
      backTrack(i + 1, size);
      set.pop();
    }
  };

  for (let k = 0; k < n + 1; k++) {
    backTrack(0, k);
  }

  return res;
}
// @lc code=end
