/*
 * @lc app=leetcode id=168 lang=typescript
 *
 * [168] Excel Sheet Column Title
 */

// @lc code=start
function convertToTitle(columnNumber: number): string {
  let title = "";

  while (columnNumber > 0) {
    const res = columnNumber % 26;
    if (res === 0) {
      title = String.fromCharCode(90) + title;
    } else {
      title = String.fromCharCode(res + 64) + title;
    }
    if (res === 0) {
      columnNumber = Math.floor(columnNumber / 26) - 1;
    } else {
      columnNumber = Math.floor(columnNumber / 26);
    }
  }

  return title;
}
// @lc code=end

701;
