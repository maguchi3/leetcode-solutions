/*
 * @lc app=leetcode id=50 lang=typescript
 *
 * [50] Pow(x, n)
 */

// @lc code=start
function myPow(x: number, n: number): number {
  const fastPow = (x: number, n: number) => {
    if (n === 0) return 1;

    const half = fastPow(x, Math.floor(n / 2));

    if (n % 2 === 0) {
      return half * half;
    } else {
      return half * half * x;
    }
  };

  if (n < 0) {
    x = 1 / x;
    n = -n;
  }

  return fastPow(x, n);
}
// @lc code=end
