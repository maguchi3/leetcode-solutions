/*
 * @lc app=leetcode id=169 lang=typescript
 *
 * [169] Majority Element
 */

// @lc code=start
function majorityElement(nums: number[]): number {
  const appeared = new Map<number, number>();

  for (let i = 0; i < nums.length; i++) {
    const num = nums[i];
    if (appeared.has(num)) {
      appeared.set(num, appeared.get(num) + 1);
    } else {
      appeared.set(num, 1);
    }
  }

  let majority: [number, number] = null;

  for (let [key, value] of appeared.entries()) {
    if (majority === null || value > majority[1]) majority = [key, value];
  }

  return majority[0];
}
// @lc code=end
