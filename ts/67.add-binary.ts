/*
 * @lc app=leetcode id=67 lang=typescript
 *
 * [67] Add Binary
 */

// @lc code=start
function addBinary(a: string, b: string): string {
  let n = a.length;
  let m = b.length;
  if (m > n) return addBinary(b, a);

  let carry = 0;
  let arr = [];
  let j = m - 1;

  for (let i = n - 1; i > -1; i--) {
    if (a[i] === "1") carry++;
    if (j > -1 && b[j--] === "1") carry++;

    carry % 2 === 0 ? arr.unshift("0") : arr.unshift("1");

    carry = Math.floor(carry / 2);
  }

  if (carry === 1) arr.unshift(1);

  return arr.join("");
}
// @lc code=end
