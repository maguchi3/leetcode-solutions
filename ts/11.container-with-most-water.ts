/*
 * @lc app=leetcode id=11 lang=typescript
 *
 * [11] Container With Most Water
 */

// @lc code=start
function maxArea(height: number[]): number {
  let max = 0;
  let left = 0;
  let right = height.length - 1;

  while (left < right) {
    const lh = height[left];
    const rh = height[right];
    const area = Math.min(lh, rh) * (right - left);
    max = Math.max(max, area);

    lh < rh ? left++ : right--;
  }

  return max;
}
// @lc code=end
