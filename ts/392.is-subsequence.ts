/*
 * @lc app=leetcode id=392 lang=typescript
 *
 * [392] Is Subsequence
 */

// @lc code=start
function isSubsequence(s: string, t: string): boolean {
  if (!s.length) return true;

  let posS = 0;
  let posT = 0;

  while (posS < s.length && posT < t.length) {
    if (s[posS] === t[posT++]) posS++;
  }

  return posS === s.length;
}

// @lc code=end
