/*
 * @lc app=leetcode id=119 lang=typescript
 *
 * [119] Pascal's Triangle II
 */

// @lc code=start
function getRow(rowIndex: number): number[] {
  const triangles: number[][] = [[1]];

  for (let i = 1; i <= rowIndex; i++) {
    const row: number[] = [];
    const prevRow = triangles[i - 1];

    row.push(1);

    for (let j = 1; j < i; j++) {
      row.push(prevRow[j - 1] + prevRow[j]);
    }

    row.push(1);

    triangles.push(row);
  }

  return triangles[rowIndex];
}
// @lc code=end
