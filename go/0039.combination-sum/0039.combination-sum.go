/*
 * @lc app=leetcode id=39 lang=golang
 *
 * [39] Combination Sum
 */
package leetcode

import (
	"sort"
)

// @lc code=start
func combinationSum(candidates []int, target int) [][]int {
	sort.Ints(candidates)

	ans, path := [][]int{}, []int{}

	backTrack(candidates, target, 0, path, &ans)

	return ans
}

func backTrack(candidates []int, rest int, pos int, path []int, ans *[][]int) {
	if rest < 0 {
		return
	}

	if rest == 0 {
		// make deep-copy to avoid overriding
		tmp := make([]int, len(path))
		copy(tmp, path)
		*ans = append(*ans, tmp)
		return
	}

	for i := pos; i < len(candidates); i++ {
		cur := candidates[i]

		path = append(path, cur)
		backTrack(candidates, rest-cur, i, path, ans)
		// back tracking
		path = path[:len(path)-1]
	}
}

// @lc code=end
