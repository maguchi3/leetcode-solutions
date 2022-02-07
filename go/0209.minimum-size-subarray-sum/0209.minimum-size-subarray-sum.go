/*
 * @lc app=leetcode id=209 lang=golang
 *
 * [209] Minimum Size Subarray Sum
 */
package leetcode

import "math"

// @lc code=start

/* Sliding Window O(n) */
func minSubArrayLen(target int, nums []int) int {
	var (
		left, sum int
		min       = math.MaxFloat64
	)

	for i, num := range nums {
		sum += num
		for sum >= target && left < len(nums) {
			min = math.Min(min, float64(i-left+1))
			sum -= nums[left]
			left++
		}
	}

	if min == math.MaxFloat64 {
		return 0
	}

	return int(min)
}

// @lc code=end

/* binary search O(n log n) */
func minSubArrayLen2(target int, nums []int) int {
	n := len(nums)
	if n == 0 {
		return 0
	}

	sums := make([]int, n)
	sums[0] = nums[0]

	for i := 1; i < n; i++ {
		sums[i] = sums[i-1] + nums[i]
	}

	min := float64(n + 1)

	for i := 0; i < n; i++ {
		if sums[i] >= target {
			left := binarySearch(0, i-1, sums[i]-target, sums)
			min = math.Min(min, float64(i-left))
		}
	}

	if min == float64(n+1) {
		return 0
	}

	return int(min)
}

func binarySearch(l, r int, target int, sums []int) int {
	for l < r {
		m := (l + r) / 2
		if sums[m] < target {
			l = m + 1
		} else {
			r = m
		}
	}

	if sums[l] > target {
		return l - 1
	}

	return l
}
