/*
 * @lc app=leetcode id=122 lang=golang
 *
 * [122] Best Time to Buy and Sell Stock II
 */

package leetcode

// @lc code=start
func maxProfit(prices []int) int {
	total := 0

	for i := 1; i < len(prices); i++ {
		profit := prices[i] - prices[i-1]

		if profit > 0 {
			total += profit
		}
	}

	return total
}

// @lc code=end
