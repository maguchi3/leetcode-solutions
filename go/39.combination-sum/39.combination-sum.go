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

	ans := [][]int{}

	for s := 0; s < len(candidates); s++ {
		var path []int

		backTrack(candidates, target, s, path, &ans)
	}

	return ans
}

func backTrack(candidates []int, rest int, pos int, path []int, ans *[][]int) {
	if pos >= len(candidates) {
		return
	}

	cur := candidates[pos]

	if rest < cur {
		return
	}

	path = append(path, cur)

	if rest == cur {
		// make deep-copy to avoid overriding
		tmp := make([]int, len(path))
		copy(tmp, path)
		*ans = append(*ans, tmp)
		return
	}

	for i := pos; i < len(candidates); i++ {
		backTrack(candidates, rest-cur, i, path, ans)
	}
}

// @lc code=end
