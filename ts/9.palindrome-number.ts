/*
 * @lc app=leetcode id=9 lang=typescript
 *
 * [9] Palindrome Number
 */

// @lc code=start
function isPalindrome(x: number): boolean {
  if (x < 0) return false;
  if (x < 10) return true;

  let digits: number[] = [];
  let currentNumber = x;

  while (currentNumber >= 1) {
    digits.push(currentNumber % 10);
    currentNumber = Math.floor(currentNumber / 10);
  }

  const lastIndex = digits.length - 1;
  for (let i = 0; i <= lastIndex; i++) {
    if (digits[i] !== digits[lastIndex - i]) return false;
  }

  return true;
}
// @lc code=end
