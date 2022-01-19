/*
 * @lc app=leetcode id=8 lang=typescript
 *
 * [8] String to Integer (atoi)
 */

// @lc code=start
function myAtoi(s: string): number {
  let sign = 1;
  let result = 0;
  let index = 0;

  let max = Math.pow(2, 31) - 1;
  let min = -Math.pow(2, 31);

  while (s[index] === " ") index++;

  if (s[index] === "+") {
    sign = 1;
    index++;
  } else if (s[index] === "-") {
    sign = -1;
    index++;
  }

  while (s[index] >= "0" && s[index] <= "9") {
    const digit = Number(s[index]);

    const isOverflow =
      result > Math.floor(max / 10) ||
      (result === Math.floor(max / 10) && digit > max % 10);

    if (isOverflow) return sign === 1 ? max : min;

    result = 10 * result + digit;
    index++;
  }

  return result * sign;
}
// @lc code=end
