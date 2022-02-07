/*
 * @lc app=leetcode id=50 lang=golang
 *
 * [50] Pow(x, n)
 */

package leetcode

// @lc code=start
func myPow(x float64, n int) float64 {
	var res float64 = 1
	if n < 0 {
		x = 1 / x
		n = -n
	}

	for i := n; i != 0; i = i / 2 {
		if i%2 == 1 {
			res *= x
		}
		x *= x
	}

	return res
}

// @lc code=end
