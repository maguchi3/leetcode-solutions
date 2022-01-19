/*
 * @lc app=leetcode id=69 lang=typescript
 *
 * [69] Sqrt(x)
 */

// @lc code=start
function mySqrt(x: number): number {
  if (x < 2) return x;
  let pivot: number;
  let left = 2;
  let right = x / 2;

  while (left <= right) {
    pivot = Math.floor(left + (right - left) / 2);
    let prod = pivot * pivot;
    if (prod < x) {
      left = pivot + 1;
    } else if (prod > x) {
      right = pivot - 1;
    } else {
      return pivot;
    }
  }

  return Math.floor(right);
}
// @lc code=end
