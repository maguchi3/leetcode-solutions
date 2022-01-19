/*
 * @lc app=leetcode id=17 lang=typescript
 *
 * [17] Letter Combinations of a Phone Number
 */

// @lc code=start
function letterCombinations(digits: string): string[] {
  if (digits.length === 0) return [];

  const letterMap = new Map<number, string>([
    [2, "abc"],
    [3, "def"],
    [4, "ghi"],
    [5, "jkl"],
    [6, "mno"],
    [7, "pqrs"],
    [8, "tuv"],
    [9, "wxyz"],
  ]);

  const combinations = [];

  const backTrack = (idx: number, path: string[]) => {
    if (path.length === digits.length) {
      combinations.push(path.join(""));
      return;
    }

    let possibleLetters = letterMap.get(Number(digits[idx]));

    for (let i = 0; i < possibleLetters.length; i++) {
      path.push(possibleLetters[i]);
      backTrack(idx + 1, path);
      path.pop();
    }
  };

  backTrack(0, []);
  return combinations;
}
// @lc code=end
