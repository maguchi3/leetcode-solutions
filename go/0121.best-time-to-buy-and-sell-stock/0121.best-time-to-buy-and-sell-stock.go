/*
 * @lc app=leetcode id=121 lang=golang
 *
 * [121] Best Time to Buy and Sell Stock
 */

package leetcode

/* iteration (one path) */

// @lc code=start

func maxProfit(prices []int) int {
	min, max := 1<<63-1, 0

	for _, p := range prices {
		if p < min {
			min = p
		} else if p-min > max {
			max = p - min
		}
	}

	return max
}

// @lc code=end

/* recursive */
func _(prices []int) int {
	return buyAndSell(prices, 0, 0)
}

func buyAndSell(prices []int, start int, max int) int {

	for i := start; i < len(prices); i++ {
		diff := prices[i] - prices[start]

		if diff < 0 {
			return buyAndSell(prices, start+1, max)
		}
		if diff > 0 && diff > max {
			max = diff
		}
	}

	return max
}
