/*
 * @lc app=leetcode id=22 lang=golang
 *
 * [22] Generate Parentheses
 */

package leetcode

// @lc code=start

func generateParenthesis(n int) []string {
	res := make([]string, 0)
	backTrack(&res, "", 0, 0, n)

	return res
}

func backTrack(res *[]string, str string, open, close int, n int) {
	if len(str) == n*2 {
		*res = append(*res, str)
		return
	}

	if open < n {
		backTrack(res, str+"(", open+1, close, n)
	}

	if close < open {
		backTrack(res, str+")", open, close+1, n)
	}
}

// @lc code=end
