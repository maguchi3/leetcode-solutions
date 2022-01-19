/*
 * @lc app=leetcode id=58 lang=typescript
 *
 * [58] Length of Last Word
 */

// @lc code=start
function lengthOfLastWord(s: string): number {
  const reversedString = s.trim().split("").reverse();

  if (reversedString.every((s) => s !== " ")) return reversedString.length;

  return reversedString.findIndex((s) => s === " ");
}
// @lc code=end
