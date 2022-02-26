/*
 * @lc app=leetcode id=283 lang=golang
 *
 * [283] Move Zeroes
 */
package leetcode

// @lc code=start
func moveZeroes(nums []int) {
	// Retaining insert-position of a non-zero value
	var nonZeroPos int

	for i, num := range nums {
		if num != 0 {
			// when non-zero value is found, swapping it for first zero position
			nums[i], nums[nonZeroPos] = nums[nonZeroPos], nums[i]
			nonZeroPos++
		}
	}
}

// @lc code=end
