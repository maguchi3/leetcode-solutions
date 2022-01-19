/*
 * @lc app=leetcode id=28 lang=typescript
 *
 * [28] Implement strStr()
 */

// @lc code=start
function strStr(haystack: string, needle: string): number {
  if (!needle) return 0;
  for (let i = 0; i < haystack.length; i++) {
    const target = haystack.substring(i, needle.length + i);
    if (target === needle) return i;
  }
  return -1;
}
// @lc code=end
