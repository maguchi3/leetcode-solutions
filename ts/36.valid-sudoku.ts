/*
 * @lc app=leetcode id=36 lang=typescript
 *
 * [36] Valid Sudoku
 */

// @lc code=start
function isValidSudoku(board: string[][]): boolean {
  const N = 9;
  const rows: Set<string>[] = new Array(N);
  const cols: Set<string>[] = new Array(N);
  const boxes: Set<string>[] = new Array(N);

  for (let r = 0; r < N; r++) {
    rows[r] = new Set<string>();
    cols[r] = new Set<string>();
    boxes[r] = new Set<string>();
  }

  for (let r = 0; r < N; r++) {
    for (let c = 0; c < N; c++) {
      const val = board[r][c];

      if (val === ".") continue;

      if (rows[r].has(val)) return false;
      rows[r].add(val);

      if (cols[c].has(val)) return false;
      cols[c].add(val);

      let b = Math.floor(r / 3) * 3 + Math.floor(c / 3);
      if (boxes[b].has(val)) return false;
      boxes[b].add(val);
    }
  }

  return true;
}
// @lc code=end
