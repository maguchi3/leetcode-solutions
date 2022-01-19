/*
 * @lc app=leetcode id=5 lang=typescript
 *
 * [5] Longest Palindromic Substring
 */

// @lc code=start
function longestPalindrome(s: string): string {
  if (s === null || s.length < 1) return "";
  let start = 0;
  let end = 0;

  for (let i = 0; i < s.length; i++) {
    let oddLen = expandAroundCenter(s, i, i);
    let evenLen = expandAroundCenter(s, i, i + 1);
    let largerLen = Math.max(oddLen, evenLen);

    if (largerLen > end - start + 1) {
      start = i - Math.floor((largerLen - 1) / 2);
      end = i + Math.floor(largerLen / 2);
    }
  }

  return s.substring(start, end + 1);
}

function expandAroundCenter(s: string, left: number, right: number): number {
  let l = left;
  let r = right;

  while (l >= 0 && r < s.length && s[l] === s[r]) {
    l--;
    r++;
  }
  return r - l - 1;
}

// @lc code=end
