/*
 * @lc app=leetcode id=12 lang=typescript
 *
 * [12] Integer to Roman
 */

// @lc code=start
function intToRoman(num: number): string {
  const values: number[] = [
    1000, 900, 500, 400, 100, 90, 50, 40, 10, 9, 5, 4, 1,
  ];
  const symbols: string[] = [
    "M",
    "CM",
    "D",
    "CD",
    "C",
    "XC",
    "L",
    "XL",
    "X",
    "IX",
    "V",
    "IV",
    "I",
  ];

  let roman = "";

  for (let i = 0; i < values.length && num > 0; i++) {
    while (values[i] <= num) {
      num -= values[i];
      roman += symbols[i];
    }
  }

  return roman;
}
// @lc code=end
