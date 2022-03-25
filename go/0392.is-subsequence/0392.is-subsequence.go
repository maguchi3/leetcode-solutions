/*
 * @lc app=leetcode id=392 lang=golang
 *
 * [392] Is Subsequence
 */

package isSubsequence

// @lc code=start

// two pointers
func isSubsequence(s string, t string) bool {
	lb, rb := len(s), len(t)

	var l, r int

	for l < lb && r < rb {
		if s[l] == t[r] {
			l++
		}
		r++
	}

	return l == lb
}

// @lc code=end
