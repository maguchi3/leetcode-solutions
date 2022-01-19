/*
 * @lc app=leetcode id=202 lang=typescript
 *
 * [202] Happy Number
 */

// @lc code=start
function nextNumber(n: number): number {
  let sum = 0;
  while (n > 0) {
    const digit = n % 10;
    n = Math.floor(n / 10);
    sum += digit ** 2;
  }

  return sum;
}

function isHappy(n: number): boolean {
  const seen = new Set<number>();
  while (!seen.has(n)) {
    seen.add(n);
    n = nextNumber(n);
  }

  return n === 1;
}
// @lc code=end
