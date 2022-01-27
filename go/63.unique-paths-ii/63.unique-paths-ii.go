/*
 * @lc app=leetcode id=63 lang=golang
 *
 * [63] Unique Paths II
 */

package leetcode

// @lc code=start
func uniquePathsWithObstacles(obstacleGrid [][]int) int {
	if obstacleGrid[0][0] == 1 {
		return 0
	}

	m, n := len(obstacleGrid[0]), len(obstacleGrid)

	path := 1

	// fill upside with 1
	for i := range obstacleGrid {
		if obstacleGrid[i][0] == 1 {
			// no path exist in this loop
			path = 0
		}
		obstacleGrid[i][0] = path
	}

	path = 1

	// fill leftside with 1
	for j := 1; j < m; j++ {
		if obstacleGrid[0][j] == 1 {
			path = 0
		}
		obstacleGrid[0][j] = path
	}

	for r := 1; r < n; r++ {
		for c := 1; c < m; c++ {
			// check if obstacle exists
			if obstacleGrid[r][c] == 1 {
				obstacleGrid[r][c] = 0
				continue
			}
			obstacleGrid[r][c] = obstacleGrid[r-1][c] + obstacleGrid[r][c-1]
		}
	}

	return obstacleGrid[n-1][m-1]
}

// @lc code=end
