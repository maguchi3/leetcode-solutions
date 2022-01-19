/*
 * @lc app=leetcode id=125 lang=typescript
 *
 * [125] Valid Palindrome
 */

// @lc code=start
function isPalindrome(s: string): boolean {
  const sanitized: string[] = [];

  for (let i = 0; i < s.length; i++) {
    const charcode = s.charCodeAt(i);

    if (
      (charcode >= 48 && charcode <= 57) ||
      (charcode >= 97 && charcode <= 122)
    ) {
      sanitized.push(s[i]);
    }

    if (charcode >= 65 && charcode <= 90) {
      sanitized.push(String.fromCharCode(charcode + 32));
    }
  }

  for (let j = 0; j < sanitized.length; j++) {
    if (sanitized[j] !== sanitized[sanitized.length - (j + 1)]) return false;
  }

  return true;
}
// @lc code=end
