/*
 * @lc app=leetcode id=63 lang=typescript
 *
 * [63] Unique Paths II
 */

// @lc code=start
function uniquePathsWithObstacles(obstacleGrid: number[][]): number {
  if (obstacleGrid[0][0] === 1) return 0;

  let m = obstacleGrid.length;
  let n = obstacleGrid[0].length;

  const dp = Array.from({ length: m }, () => new Uint32Array(n));
  dp[0][0] = 1;

  for (let i = 0; i < m; i++) {
    for (let j = 0; j < n; j++) {
      if (obstacleGrid[i][j] || (!i && !j)) continue;
      else dp[i][j] = (i ? dp[i - 1][j] : 0) + (j ? dp[i][j - 1] : 0);
    }
  }

  return dp[m - 1][n - 1];
}
// @lc code=end
