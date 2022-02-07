/*
 * @lc app=leetcode id=200 lang=golang
 *
 * [200] Number of Islands
 */
package leetcode

// @lc code=start

/* DFS(recursion) */
func numIslands(grid [][]byte) int {
	var islands int

	for i, row := range grid {
		for j := range row {
			if grid[i][j] == '1' {
				islands++
				explore(grid, i, j)
			}
		}
	}

	return islands
}

func explore(grid [][]byte, r int, c int) {
	n, m := len(grid), len(grid[0])
	isInvalidPos := r < 0 || r == n || c < 0 || c == m

	if isInvalidPos || grid[r][c] == '0' {
		return
	}

	grid[r][c] = '0'

	explore(grid, r+1, c)
	explore(grid, r, c+1)
	explore(grid, r-1, c)
	explore(grid, r, c-1)
}

// @lc code=end

/* BFS */
func numIslands2(grid [][]byte) int {
	var (
		islands int
		queue   [][2]int
	)
	n, m := len(grid), len(grid[0])

	for r := 0; r < n; r++ {
		for c := 0; c < m; c++ {
			if grid[r][c] == '1' {
				islands++
				queue = append(queue, [2]int{r, c})
				grid[r][c] = 0
				exploreLands(queue, grid)
			}
		}
	}

	return islands
}

func exploreLands(queue [][2]int, grid [][]byte) {
	offset := []int{0, 1, 0, -1, 0}

	for len(queue) > 0 {
		pair := queue[0]
		queue = queue[1:]
		for i := 0; i < 4; i++ {
			r, c := pair[0]+offset[i], pair[1]+offset[i+1]
			isValidRow := r >= 0 && r < len(grid)
			isValidCol := c >= 0 && c < len(grid[0])

			if isValidRow && isValidCol && grid[r][c] == '1' {
				grid[r][c] = 0
				queue = append(queue, [2]int{r, c})
			}
		}
	}
}
