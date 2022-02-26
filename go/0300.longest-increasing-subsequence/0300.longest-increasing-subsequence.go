/*
 * @lc app=leetcode id=300 lang=golang
 *
 * [300] Longest Increasing Subsequence
 */

package leetcode

// @lc code=start

// Binary Search
func lengthOfLIS(nums []int) int {
	n := len(nums)
	seq := []int{nums[0]}

	for i := 1; i < n; i++ {
		num := nums[i]
		if num > seq[len(seq)-1] {
			seq = append(seq, num)
		} else {
			pos := binarySearch(seq, num)
			seq[pos] = num
		}
	}

	return len(seq)
}

func binarySearch(seq []int, target int) int {
	left, right := 0, len(seq)-1
	var mid int

	for right > left {
		mid = (left + right) / 2

		if seq[mid] == target {
			return mid
		}

		if seq[mid] > target {
			right = mid
		} else {
			left = mid + 1
		}
	}

	return left
}

// @lc code=end

// Dynamic Programming O(n^2)
func lengthOfLIS2(nums []int) int {
	n := len(nums)
	// lengths[i] = longest subsequence length at index i
	lengths := make([]int, n)

	// fill with 1
	for i := 0; i < n; i++ {
		lengths[i] = 1
	}

	for r := 1; r < len(nums); r++ {
		for l := 0; l < r; l++ {
			if nums[r] > nums[l] {
				lengths[r] = max(lengths[r], lengths[l]+1)
			}
		}
	}

	maxLen := 0

	for i := range lengths {
		maxLen = max(maxLen, lengths[i])
	}

	return maxLen
}

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}
