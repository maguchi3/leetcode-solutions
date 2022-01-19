/*
 * @lc app=leetcode id=1011 lang=typescript
 *
 * [1011] Capacity To Ship Packages Within D Days
 */

// @lc code=start
function shipWithinDays(weights: number[], days: number): number {
  const getDays = (cap: number) => {
    let days = 1,
      total = 0;

    for (let w of weights) {
      total += w;

      if (total > cap) {
        total = w;
        days++;
      }
    }

    return days;
  };

  let min = Math.max(...weights);
  let max = weights.reduce((a, b) => a + b);

  while (min < max) {
    const med = ~~((min + max) / 2);
    const _days = getDays(med);

    if (_days > days) min = med + 1;
    else max = med;
  }

  return max;
}
// @lc code=end
