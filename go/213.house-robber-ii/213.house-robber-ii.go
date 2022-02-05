/*
 * @lc app=leetcode id=213 lang=golang
 *
 * [213] House Robber II
 */
package leetcode

import "math"

// @lc code=start

/* Dynamic Programing */
func rob(nums []int) int {
	n := len(nums)
	m1, m2 := maxBooty(nums, 0, n-2), maxBooty(nums, 1, n-1)

	if m1 < m2 {
		return m2
	}

	return m1
}

func maxBooty(houses []int, start, end int) int {
	if len(houses) == 1 {
		return houses[0]
	}

	var prev, max float64

	for i := start; i <= end; i++ {
		tmp := max
		max = math.Max(max, prev+float64(houses[i]))
		prev = tmp
	}

	return int(max)
}

// @lc code=end
