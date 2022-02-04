/*
 * @lc app=leetcode id=153 lang=golang
 *
 * [153] Find Minimum in Rotated Sorted Array
 */
package leetcode

// @lc code=start
func findMin(nums []int) int {
	left, right := 0, len(nums)-1

	// in this case, we have no rotation
	if nums[left] < nums[right] {
		return nums[left]
	}

	for left < right {
		middle := (left + right) / 2

		// in this case, middle is end of the rotation
		if nums[middle] > nums[middle+1] {
			return nums[middle+1]
		}

		// check in which part boundary exists
		if nums[middle] > nums[left] {
			left = middle + 1
		} else {
			right = middle
		}
	}

	return nums[left]
}

// @lc code=end
