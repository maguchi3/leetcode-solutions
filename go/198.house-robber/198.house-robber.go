/*
 * @lc app=leetcode id=198 lang=golang
 *
 * [198] House Robber
 */

package leetcode

import "math"

// @lc code=start

/* optimized DP */
func rob(nums []int) int {
	var prev, max float64

	for _, num := range nums {
		tmp := max
		max = math.Max(prev+float64(num), max)
		prev = tmp
	}

	return int(max)
}

// @lc code=end

/* recursion with memo approach */
func rob2(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	if len(nums) == 1 {
		return nums[0]
	}
	memo := make(map[int]int)

	amount := robFrom(0, nums, memo)

	return amount
}

func robFrom(i int, nums []int, memo map[int]int) int {
	if i >= len(nums) {
		return 0
	}

	if _, ok := memo[i]; ok {
		return memo[i]
	}

	m1, m2 := nums[i]+robFrom(i+2, nums, memo), robFrom(i+1, nums, memo)
	max := m1
	if max < m2 {
		max = m2
	}
	memo[i] = max

	return max
}
