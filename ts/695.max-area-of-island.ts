/*
 * @lc app=leetcode id=695 lang=typescript
 *
 * [695] Max Area of Island
 */

// @lc code=start
function maxAreaOfIsland(grid: number[][]): number {
  const calcAreaFrom = (r: number, c: number): number => {
    if (grid[r]?.[c] !== 1) return 0;

    grid[r][c] = 0;

    const area =
      1 +
      calcAreaFrom(r + 1, c) +
      calcAreaFrom(r, c + 1) +
      calcAreaFrom(r - 1, c) +
      calcAreaFrom(r, c - 1);

    return area;
  };

  let maxArea = 0;

  for (let row = 0; row < grid.length; row++) {
    for (let col = 0; col < grid[0].length; col++) {
      maxArea = Math.max(maxArea, calcAreaFrom(row, col));
    }
  }

  return maxArea;
}
// @lc code=end
