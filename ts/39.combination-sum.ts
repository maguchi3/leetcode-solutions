/*
 * @lc app=leetcode id=39 lang=typescript
 *
 * [39] Combination Sum
 */

// @lc code=start
function combinationSum(candidates: number[], target: number): number[][] {
  const comb: number[] = [];
  const res: number[][] = [];

  const backTrack = (remain: number, start: number) => {
    if (remain === 0) {
      res.push([...comb]);
    } else if (remain < 0) {
      return;
    }

    for (let i = start; i < candidates.length; i++) {
      comb.push(candidates[i]);
      backTrack(remain - candidates[i], i);
      comb.pop();
    }
  };

  backTrack(target, 0);

  return res;
}
// @lc code=end
