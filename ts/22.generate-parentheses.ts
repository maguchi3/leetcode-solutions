/*
 * @lc app=leetcode id=22 lang=typescript
 *
 * [22] Generate Parentheses
 */

// @lc code=start
function generateParenthesis(n: number): string[] {
  const res: string[] = [];
  generateAll(new Array(2 * n), 0, res);

  return res;
}

function generateAll(cur: string[], pos: number, res: string[]) {
  if (pos === cur.length) {
    if (isValid(cur)) res.push(cur.join(""));
  } else {
    cur[pos] = "(";
    generateAll(cur, pos + 1, res);
    cur[pos] = ")";
    generateAll(cur, pos + 1, res);
  }
}

function isValid(cur: string[]): boolean {
  let balance = 0;

  for (let c of cur) {
    if (c === "(") {
      balance++;
    } else {
      balance--;
    }

    if (balance < 0) return false;
  }

  return balance === 0;
}
// @lc code=end
