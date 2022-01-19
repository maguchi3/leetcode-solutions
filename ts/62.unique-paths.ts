/*
 * @lc app=leetcode id=62 lang=typescript
 *
 * [62] Unique Paths
 */

// @lc code=start
function uniquePaths(m: number, n: number): number {
  const grid: number[][] = new Array(m);
  for (let i = 0; i < m; i++) grid[i] = new Array(n).fill(1);

  for (let c = 1; c < m; c++) {
    for (let r = 1; r < n; r++) {
      grid[c][r] = grid[c - 1][r] + grid[c][r - 1];
    }
  }

  return grid[m - 1][n - 1];
}
// @lc code=end
