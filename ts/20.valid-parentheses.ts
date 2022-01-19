/*
 * @lc app=leetcode id=20 lang=typescript
 *
 * [20] Valid Parentheses
 */

// @lc code=start
function isValid(s: string): boolean {
  const stack = [];

  for (let i = 0; i < s.length; i++) {
    if (OPEN_SET.has(s[i])) {
      stack.push(s[i]);
    } else if (PARENTHESES_MAP.get(s[i]) === stack[stack.length - 1]) {
      stack.pop();
    } else {
      return false;
    }
  }

  return stack.length === 0;
}

const OPEN_SET = new Set(["(", "{", "["]);

const PARENTHESES_MAP = new Map([
  [")", "("],
  ["}", "{"],
  ["]", "["],
]);
// @lc code=end
