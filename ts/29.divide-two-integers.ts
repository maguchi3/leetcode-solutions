/*
 * @lc app=leetcode id=29 lang=typescript
 *
 * [29] Divide Two Integers
 */

// @lc code=start
function divide(dividend: number, divisor: number): number {
  const MAX_VALUE = 2147483647; // 2**31 - 1
  const MIN_VALUE = -2147483648; // -2**31
  const HALF_INT_MIN = -1073741824;

  if (dividend === MIN_VALUE && divisor === -1) return MAX_VALUE;

  let negatives = 2;

  if (dividend > 0) {
    negatives--;
    dividend = -dividend;
  }

  if (divisor > 0) {
    negatives--;
    divisor = -divisor;
  }

  let quotient = 0;

  while (divisor >= dividend) {
    let powerOfTwo = -1;
    let value = divisor;

    while (value >= HALF_INT_MIN && value + value >= dividend) {
      value += value;
      powerOfTwo += powerOfTwo;
    }

    quotient += powerOfTwo;
    dividend -= value;
  }

  if (negatives !== 1) return -quotient;

  return quotient;
}
// @lc code=end
