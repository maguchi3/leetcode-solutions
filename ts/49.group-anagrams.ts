/*
 * @lc app=leetcode id=49 lang=typescript
 *
 * [49] Group Anagrams
 */

// @lc code=start
function groupAnagrams(strs: string[]): string[][] {
  if (strs.length === 0) return [];

  const ans = new Map<string, string[]>();

  for (let s of strs) {
    let ca = s.split("");
    ca = ca.sort();
    const key = ca.join("");

    if (!ans.has(key)) ans.set(key, []);
    ans.get(key).push(s);
  }

  return [...ans.values()];
}
// @lc code=end
