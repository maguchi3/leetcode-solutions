/*
 * @lc app=leetcode id=13 lang=typescript
 *
 * [13] Roman to Integer
 */

// @lc code=start
function romanToInt(s: string): number {
  let result = 0;

  for (let i = 0; i < s.length; i++) {
    let number = ROMAN_INT_MAP.get(s[i]);
    let nextNumber = ROMAN_INT_MAP.get(s[i + 1]);

    if ([1, 10, 100].includes(number) && number < nextNumber) {
      result -= number;
      continue;
    }
    result += number;
  }

  return result;
}

const ROMAN_INT_MAP = new Map([
  ["I", 1],
  ["V", 5],
  ["X", 10],
  ["L", 50],
  ["C", 100],
  ["D", 500],
  ["M", 1000],
]);
// @lc code=end
