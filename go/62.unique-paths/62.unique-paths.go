/*
 * @lc app=leetcode id=62 lang=golang
 *
 * [62] Unique Paths
 */
package leetcode

// @lc code=start
func uniquePaths(m int, n int) int {
	// create row space
	paths := make([][]int, n)

	for i := 0; i < n; i++ {
		// create col space
		paths[i] = make([]int, m)
		// fill leftside with 1
		paths[i][0] = 1
	}

	for j := 0; j < m; j++ {
		// fill upside with 1
		paths[0][j] = 1
	}

	for r := 1; r < n; r++ {
		for c := 1; c < m; c++ {
			// calc ways from up and left
			paths[r][c] = paths[r-1][c] + paths[r][c-1]
		}
	}

	return paths[n-1][m-1]
}

// @lc code=end
