/*
 * @lc app=leetcode id=190 lang=typescript
 *
 * [190] Reverse Bits
 */

// @lc code=start
function reverseBits(n: number): number {
  let result = 0;
  let power = 31;

  while (n !== 0) {
    result += (n & 1) << power;
    n = n >> 1;
    power--;
  }

  return result;
}
// @lc code=end
