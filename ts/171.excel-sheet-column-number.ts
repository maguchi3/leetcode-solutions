/*
 * @lc app=leetcode id=171 lang=typescript
 *
 * [171] Excel Sheet Column Number
 */

// @lc code=start
function titleToNumber(columnTitle: string): number {
  let result = 0;
  for (let i = 0; i < columnTitle.length; i++) {
    const num = columnTitle.charCodeAt(columnTitle.length - (i + 1)) - 64;
    if (i === 0) {
      result += num;
    } else {
      result += num * 26 ** i;
    }
  }
  return result;
}
// @lc code=end
