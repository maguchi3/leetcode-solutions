/*
 * @lc app=leetcode id=53 lang=golang
 *
 * [53] Maximum Subarray
 */

package leetcode

// @lc code=start
func maxSubArray(nums []int) int {
	cur, max := nums[0], nums[0]

	for _, v := range nums[1:] {
		if cur < 0 {
			cur = v
		} else {
			cur += v
		}
		if max < cur {
			max = cur
		}
	}

	return max
}

// @lc code=end
