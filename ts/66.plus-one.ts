/*
 * @lc app=leetcode id=66 lang=typescript
 *
 * [66] Plus One
 */

// @lc code=start
function plusOne(digits: number[]): number[] {
  const reversedDigits = digits.reverse();
  const firstNumber = reversedDigits[0];
  if (firstNumber === 9) {
    let i = 0;
    while (reversedDigits[i] === 9) {
      reversedDigits.splice(i, 1, 0);
      i++;
    }

    reversedDigits[i]
      ? reversedDigits.splice(i, 1, reversedDigits[i] + 1)
      : reversedDigits.push(1);
  } else {
    reversedDigits.splice(0, 1, firstNumber + 1);
  }

  const result = reversedDigits.reverse();

  return result;
}
// @lc code=end
