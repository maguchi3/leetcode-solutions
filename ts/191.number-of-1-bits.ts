/*
 * @lc app=leetcode id=191 lang=typescript
 *
 * [191] Number of 1 Bits
 */

// @lc code=start
function hammingWeight(n: number): number {
  let bits = 0;
  let mask = 1;

  for (let i = 0; i < 32; i++) {
    if ((n & mask) !== 0) bits++;
    mask <<= 1;
  }

  return bits;
}
// @lc code=end
