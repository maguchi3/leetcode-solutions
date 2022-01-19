/*
 * @lc app=leetcode id=6 lang=typescript
 *
 * [6] Zigzag Conversion
 */

// @lc code=start
function convert(s: string, numRows: number): string {
  if (numRows === 1) return s;
  const zigzag = new Array(numRows).fill("");
  let isIncreasing = true;
  let row = 0;

  for (let i = 0; i < s.length; i++) {
    zigzag[row] += s[i];
    if (row === 0) isIncreasing = true;
    if (row === numRows - 1) isIncreasing = false;

    row += isIncreasing ? 1 : -1;
  }

  return zigzag.join("");
}
// @lc code=end
