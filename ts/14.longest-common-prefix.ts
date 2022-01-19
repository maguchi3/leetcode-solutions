/*
 * @lc app=leetcode id=14 lang=typescript
 *
 * [14] Longest Common Prefix
 */

// @lc code=start
function longestCommonPrefix(strs: string[]): string {
  if (strs.length === 1) return strs[0];
  let prefix = strs[0];

  strs.slice(1).forEach((str) => {
    while (str.indexOf(prefix) !== 0) {
      if (prefix === "") break;
      prefix = prefix.slice(0, -1);
    }
  });

  return prefix;
}
// @lc code=end
