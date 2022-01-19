/*
 * @lc app=leetcode id=122 lang=typescript
 *
 * [122] Best Time to Buy and Sell Stock II
 */

// @lc code=start
function maxProfit(prices: number[]): number {
  let profit = 0;

  for (let i = 1; i < prices.length; i++) {
    const cur = prices[i];
    const prev = prices[i - 1];

    if (prev < cur) profit += cur - prev;
  }

  return profit;
}
// @lc code=end
