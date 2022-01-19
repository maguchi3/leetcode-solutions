/*
 * @lc app=leetcode id=136 lang=typescript
 *
 * [136] Single Number
 */

// @lc code=start
function singleNumber(nums: number[]): number {
  let answer = 0;
  for (const num of nums) {
    answer ^= num;
  }

  return answer;
}

// @lc code=end
