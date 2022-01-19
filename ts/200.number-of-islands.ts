/*
 * @lc app=leetcode id=200 lang=typescript
 *
 * [200] Number of Islands
 */

// @lc code=start
function numIslands(grid: string[][]): number {
  let count = 0;

  for (let r = 0; r < grid.length; r++) {
    for (let c = 0; c < grid[r].length; c++) {
      if (grid[r][c] === "1") {
        count++;
        explore(grid, r, c);
      }
    }
  }

  return count;
}

function explore(grid: string[][], row: number, col: number) {
  const invalidPos =
    row < 0 ||
    col < 0 ||
    row >= grid.length ||
    col >= grid[0].length ||
    grid[row][col] === "0";

  if (invalidPos) return;

  grid[row][col] = "0";
  explore(grid, row - 1, col);
  explore(grid, row, col - 1);
  explore(grid, row + 1, col);
  explore(grid, row, col + 1);
}
// @lc code=end
