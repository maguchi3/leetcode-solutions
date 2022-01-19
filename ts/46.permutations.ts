/*
 * @lc app=leetcode id=46 lang=typescript
 *
 * [46] Permutations
 */

// @lc code=start
function permute(nums: number[]): number[][] {
  const N = nums.length;
  const res = [];
  const backTrack = (start: number) => {
    if (start === N) res.push([...nums]);

    for (let i = start; i < N; i++) {
      swap(start, i);
      backTrack(start + 1);
      swap(start, i);
    }
  };

  const swap = (i: number, j: number) => {
    const temp = nums[i];
    nums[i] = nums[j];
    nums[j] = temp;
  };

  backTrack(0);

  return res;
}
// @lc code=end
