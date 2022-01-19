/*
 * @lc app=leetcode id=322 lang=typescript
 *
 * [322] Coin Change
 */

// @lc code=start

// bottom-up solution

function coinChange(coins: number[], amount: number) {
  const dp = new Array(amount + 1).fill(Infinity);
  dp[0] = 0;

  for (let i = 1; i <= amount; i++) {
    for (let coin of coins) {
      if (i - coin >= 0) dp[i] = Math.min(dp[i], dp[i - coin] + 1);
    }
  }

  return dp[amount] === Infinity ? -1 : dp[amount];
}

// top-down solution

// function coinChange(coins: number[], amount: number, dp = {}): number {
//   let min = Number.MAX_SAFE_INTEGER;

//   if (amount < 0) return -1;
//   if (amount === 0) return 0;

//   if (amount in dp) return dp[amount];

//   for (let coin of coins) {
//     const count = coinChange(coins, amount - coin, dp);

//     if (count !== -1) min = Math.min(min, count + 1);
//   }

//   dp[amount] = min === Number.MAX_SAFE_INTEGER ? -1 : min;

//   return dp[amount];
// }

// @lc code=end
