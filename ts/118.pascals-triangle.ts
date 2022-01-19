/*
 * @lc app=leetcode id=118 lang=typescript
 *
 * [118] Pascal's Triangle
 */

// @lc code=start
function generate(numRows: number): number[][] {
  const triangles: number[][] = [[1]];

  for (let rowNum = 1; rowNum < numRows; rowNum++) {
    const row: number[] = [];
    const prevRow = triangles[rowNum - 1];

    row.push(1);

    for (let i = 1; i < rowNum; i++) {
      row.push(prevRow[i - 1] + prevRow[i]);
    }

    row.push(1);
    triangles.push(row);
  }

  return triangles;
}
// @lc code=end
