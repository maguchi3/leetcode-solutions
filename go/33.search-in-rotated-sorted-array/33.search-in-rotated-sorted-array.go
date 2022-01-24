/*
 * @lc app=leetcode id=33 lang=golang
 *
 * [33] Search in Rotated Sorted Array
 */

package leetcode

// @lc code=start
func search(nums []int, target int) int {
	n := len(nums) - 1
	left, right := 0, n

	for left <= right {
		mid := (left + right) / 2

		if nums[mid] == target {
			return mid
		}

		// left part is ordered
		if nums[mid] >= nums[left] {
			// target exists between left and mid
			if target >= nums[left] && target < nums[mid] {
				right = mid - 1
			} else {
				left = mid + 1
			}
		} else {
			// target exists between mid and right
			if target <= nums[right] && target > nums[mid] {
				left = mid + 1
			} else {
				right = mid - 1
			}
		}
	}

	return -1
}

// @lc code=end
