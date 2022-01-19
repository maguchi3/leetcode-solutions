/*
 * @lc app=leetcode id=139 lang=typescript
 *
 * [139] Word Break
 */

// @lc code=start
function wordBreak(s: string, wordDict: string[]): boolean {
  if (wordDict.length === 0) return false;
  const wordSet = new Set(wordDict);

  const visited = new Set<number>();
  const q = [0];

  while (q.length) {
    const start = q.shift();

    if (!visited.has(start)) {
      for (let end = start + 1; end <= s.length; end++) {
        if (wordSet.has(s.slice(start, end))) {
          if (end === s.length) return true;
          q.push(end);
        }
      }

      visited.add(start);
    }
  }

  return false;
}
// @lc code=end
