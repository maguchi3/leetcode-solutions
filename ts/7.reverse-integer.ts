/*
 * @lc app=leetcode id=7 lang=typescript
 *
 * [7] Reverse Integer
 */

// @lc code=start
function reverse(x: number): number {
  const xArr = x.toString().split("");

  if (xArr.length === 1) return Number(xArr);

  const xArrReversed = xArr.reverse();

  if (xArrReversed[xArr.length - 1] === "-") {
    const negative = xArrReversed.pop();
    xArrReversed.unshift(negative);
  } else if (xArrReversed[0] === "0") {
    xArrReversed.shift();
  }

  const reversedNumber = Number(xArrReversed.join(""));
  const isSafeNumber = (num: number) => num < 2 ** 31 - 1 && num > -(2 ** 31);

  return isSafeNumber(reversedNumber) ? reversedNumber : 0;
}
// @lc code=end
