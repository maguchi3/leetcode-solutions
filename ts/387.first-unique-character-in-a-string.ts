/*
 * @lc app=leetcode id=387 lang=typescript
 *
 * [387] First Unique Character in a String
 */

// @lc code=start
function firstUniqChar(s: string): number {
  const appeared = new Map<string, number>();

  for (let i = 0; i < s.length; i++) {
    const c = s[i];
    appeared.set(c, (appeared.get(c) || 0) + 1);
  }

  for (let [char, count] of appeared) {
    if (count === 1) return s.indexOf(char);
  }

  return -1;
}
// @lc code=end
